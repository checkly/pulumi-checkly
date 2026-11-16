package checkly

import (
	"context"
	"io"
	"net/http"
	"slices"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

// apiCheckJSON is the API's representation of the check in the CheckAPI case.
const apiCheckJSON = `{
	"id": "1234",
	"checkType": "API",
	"name": "api check",
	"description": "",
	"frequency": 10,
	"activated": true,
	"muted": false,
	"shouldFail": false,
	"runParallel": false,
	"locations": ["eu-west-1"],
	"privateLocations": [],
	"tags": ["pulumi"],
	"degradedResponseTime": 10000,
	"maxResponseTime": 20000,
	"useGlobalAlertSettings": false,
	"alertSettings": {
		"escalationType": "RUN_BASED",
		"runBasedEscalation": {"failedRunThreshold": 1},
		"reminders": {"amount": 0, "interval": 5},
		"parallelRunFailureThreshold": {"enabled": false, "percentage": 10}
	},
	"retryStrategy": {
		"type": "FIXED", "baseBackoffSeconds": 60, "maxRetries": 2,
		"maxDurationSeconds": 600, "sameRegion": true
	},
	"request": {
		"method": "GET",
		"url": "https://api.example.com/health",
		"followRedirects": true,
		"skipSSL": false,
		"body": "",
		"bodyType": "NONE",
		"ipFamily": "IPv4",
		"headers": [{"key": "Accept", "value": "application/json", "locked": false}],
		"queryParameters": [],
		"assertions": [{"source": "STATUS_CODE", "property": "", "comparison": "EQUALS", "target": "200", "regex": null}],
		"basicAuth": null
	},
	"alertChannelSubscriptions": [{"alertChannelId": 42, "activated": true}],
	"doubleCheck": false,
	"localSetupScript": null,
	"localTearDownScript": null,
	"groupOrder": null
}`

// browserCheckJSON is the API's representation of the check in the CheckBrowser case.
const browserCheckJSON = `{
	"id": "1234",
	"checkType": "BROWSER",
	"name": "browser check",
	"description": "",
	"frequency": 10,
	"activated": true,
	"muted": false,
	"runParallel": false,
	"locations": ["eu-west-1"],
	"privateLocations": [],
	"tags": [],
	"script": "console.log('hello')",
	"useGlobalAlertSettings": true,
	"alertSettings": {
		"escalationType": "RUN_BASED",
		"runBasedEscalation": {"failedRunThreshold": 1},
		"reminders": {"amount": 0, "interval": 5},
		"parallelRunFailureThreshold": {"enabled": false, "percentage": 10}
	},
	"retryStrategy": {
		"type": "NO_RETRIES", "baseBackoffSeconds": 0, "maxRetries": 0,
		"maxDurationSeconds": 0, "sameRegion": false
	},
	"environmentVariables": [
		{"key": "BASE_URL", "value": "https://example.com", "locked": false, "secret": false}
	],
	"alertChannelSubscriptions": [],
	"doubleCheck": false,
	"groupOrder": null
}`

// TestRefreshRewritesV2State refreshes the v2.x state of a Check, as `pulumi refresh` does,
// against a local API server that returns the check. It then checks that the refreshed state
// can be deleted, and that it moves to
// its per-type resource with a Pulumi alias without any change.
func TestRefreshRewritesV2State(t *testing.T) {
	setTestVersion()
	cases := map[string]stateCompatCase{}
	for _, c := range stateCompatCases(t) {
		cases[c.name] = c
	}
	for _, move := range []struct {
		from, json, to string
		// wantDiffs are the properties that the move updates once.
		wantDiffs []string
	}{
		{"CheckAPI", apiCheckJSON, "checkly:index/apiCheck:ApiCheck", nil},
		// Check has no dependencies, so BrowserCheck plans its empty default once.
		{"CheckBrowser", browserCheckJSON, "checkly:index/browserCheck:BrowserCheck", []string{"dependencies"}},
	} {
		t.Run(move.from, func(t *testing.T) {
			testRefreshAndMove(t, cases[move.from], move.json, move.to, move.wantDiffs)
		})
	}
}

func testRefreshAndMove(t *testing.T, tc stateCompatCase, checkJSON, target string, wantDiffs []string) {
	ctx := context.Background()

	deleted := false
	srv := newTestProviderServer(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/v1/checks/1234":
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, checkJSON)
		case r.Method == http.MethodDelete && r.URL.Path == "/v1/checks/1234":
			deleted = true
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected Checkly API request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	checkURN := resource.NewURN("test", "statecompat", "", tokens.Type(tc.resourceType), tc.name)
	read, err := srv.Read(ctx, &pulumirpc.ReadRequest{
		Id:         "1234",
		Urn:        string(checkURN),
		Type:       tc.resourceType,
		Name:       tc.name,
		Properties: marshal(t, tc.v2State()),
		Inputs:     marshal(t, tc.v2Inputs),
	})
	if err != nil {
		t.Fatalf("Read: %v", err)
	}
	refreshed, err := plugin.UnmarshalProperties(read.GetProperties(), plugin.MarshalOptions{
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		t.Fatal(err)
	}
	// The refreshed state is at the current schema version, and the upstream upgrade has
	// normalized the zero values that v2.x stored, such as groupId 0, to null.
	if meta := refreshed["__meta"]; !meta.IsString() || meta.StringValue() != `{"schema_version":"1"}` {
		t.Errorf("refreshed __meta is %v, want schema version 1", meta)
	}
	if groupID := refreshed["groupId"]; !groupID.IsNull() {
		t.Errorf("refreshed groupId is %v, want null", groupID)
	}
	if reminders := refreshed["alertSettings"].ObjectValue()["reminders"]; !reminders.IsObject() {
		t.Errorf("refreshed alertSettings.reminders is %v, want an object", reminders)
	}
	for _, k := range []resource.PropertyKey{"environmentVariable", "sslCheck"} {
		if _, ok := refreshed[k]; ok {
			t.Errorf("refreshed state still has %s", k)
		}
	}

	if _, err := srv.Delete(ctx, &pulumirpc.DeleteRequest{
		Id:         "1234",
		Urn:        string(checkURN),
		Type:       tc.resourceType,
		Name:       tc.name,
		Properties: read.GetProperties(),
		OldInputs:  read.GetInputs(),
	}); err != nil {
		t.Errorf("Delete of the refreshed state: %v", err)
	}
	if !deleted {
		t.Errorf("Delete did not call the API")
	}

	targetToken := target
	inputs := without(tc.v3Inputs, "type")
	targetURN := resource.NewURN("test", "statecompat", "", tokens.Type(targetToken), tc.name)
	check, err := srv.Check(ctx, &pulumirpc.CheckRequest{
		Urn:        string(targetURN),
		Type:       targetToken,
		Name:       tc.name,
		Olds:       read.GetInputs(),
		News:       marshal(t, inputs),
		RandomSeed: []byte("statecompat"),
	})
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	for _, f := range check.GetFailures() {
		t.Errorf("Check failure: %s: %s", f.GetProperty(), f.GetReason())
	}
	diff, err := srv.Diff(ctx, &pulumirpc.DiffRequest{
		Id:        "1234",
		Urn:       string(targetURN),
		Type:      targetToken,
		Name:      tc.name,
		Olds:      read.GetProperties(),
		News:      check.GetInputs(),
		OldInputs: read.GetInputs(),
	})
	if err != nil {
		t.Fatalf("Diff: %v", err)
	}
	if !slices.Equal(diff.GetDiffs(), wantDiffs) || len(diff.GetReplaces()) > 0 {
		t.Errorf("moving the refreshed Check to %s changes %v, want %v, and replaces %v", target,
			diff.GetDiffs(), wantDiffs, diff.GetReplaces())
	}
}
