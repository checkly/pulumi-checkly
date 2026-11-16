package checkly

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	_ "embed"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"maps"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	pulumirpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/checkly/pulumi-checkly/provider/pkg/version"
)

//go:embed cmd/pulumi-resource-checkly/schema.json
var testPackageSchema []byte

// stateCompatCase describes a resource as a stack created with pulumi-checkly v2.x stores it,
// and the program a user writes for the same resource after upgrading to v3.
type stateCompatCase struct {
	name  string
	token string
	// v2Inputs are the inputs recorded in the v2.x state, with v2.x property names and shapes.
	v2Inputs map[string]any
	// v2Outputs are the properties the v2.x state holds in addition to its inputs.
	v2Outputs map[string]any
	// v3Inputs are the program inputs for the same resource on v3.
	v3Inputs map[string]any
	// wantNoChanges marks cases whose v2.x state holds every property, as v2.x stores it, so
	// that the matching v3 program must diff as unchanged. Other cases leave out computed
	// properties, which the plan then fills in.
	wantNoChanges bool
	// v3ShapeFrom, when set, names the Terraform resource whose v2.x state upgrade turns the
	// state into the shape v3 writes, at schema version 1. It approximates state that v3 has
	// already refreshed or updated.
	v3ShapeFrom string
}

// rewrittenProperties are the properties that the v2.x state upgrade hooks rewrite. When the
// v2.x state holds one and the program sets the same value, a diff on it would mean the
// rewrite lost or changed the stored value.
var rewrittenProperties = []string{"alertSettings", "enforceAlertSettings", "environmentVariables"}

// TestStateCompatibility upgrades v2.x state of every resource that v3 still provides and
// diffs it against the matching v3 program. It runs offline: the provider is configured
// with dummy credentials and an API URL pointing at a local server that fails the test on
// any request.
func TestStateCompatibility(t *testing.T) {
	setTestVersion()
	for _, tc := range stateCompatCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			runStateCompatCase(t, tc)
		})
	}
}

// TestStateCompatibilityRefresh refreshes v2.x state of every resource, as `pulumi refresh`
// does. The local API server answers every read with 404 Not Found, which the provider
// reports as a deleted resource, so the test checks that the state is read without an error
// before the API is called.
func TestStateCompatibilityRefresh(t *testing.T) {
	setTestVersion()
	for _, tc := range stateCompatCases(t) {
		t.Run(tc.name, func(t *testing.T) {
			reads := 0
			srv := newTestProviderServer(t, func(w http.ResponseWriter, r *http.Request) {
				// Code bundles are read with a POST to the peek endpoint.
				isRead := r.Method == http.MethodGet ||
					(r.Method == http.MethodPost && strings.HasSuffix(r.URL.Path, "/peek-code-bundle"))
				if !isRead {
					t.Errorf("unexpected Checkly API request: %s %s", r.Method, r.URL.Path)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				reads++
				w.WriteHeader(http.StatusNotFound)
			})
			urn := resource.NewURN("test", "statecompat", "", tokens.Type(tc.token), tc.name)
			if _, err := srv.Read(context.Background(), &pulumirpc.ReadRequest{
				Id:         "1234",
				Urn:        string(urn),
				Type:       tc.token,
				Name:       tc.name,
				Properties: marshal(t, tc.v2State()),
				Inputs:     marshal(t, tc.v2Inputs),
			}); err != nil {
				t.Fatalf("Read: %v", err)
			}
			if reads == 0 {
				t.Errorf("Read did not call the API")
			}
		})
	}
}

// setTestVersion sets the provider version, which only the linker sets otherwise. Provider
// derives the Go SDK import path from it.
func setTestVersion() {
	if version.Version == "" {
		version.Version = "3.0.0-alpha.0+dev"
	}
}

// v2State returns the state that v2.x stores for the case.
func (tc stateCompatCase) v2State() map[string]any {
	state := map[string]any{"id": "1234"}
	maps.Copy(state, tc.v2Inputs)
	maps.Copy(state, tc.v2Outputs)
	if tc.v3ShapeFrom == "" {
		return state
	}
	upgraded := upgradeV2State(v2StateShapes[tc.v3ShapeFrom],
		tfbridge.PreStateUpgradeHookArgs{PriorState: resource.NewPropertyMapFromMap(state)})
	upgraded["__meta"] = resource.NewStringProperty(`{"schema_version":"1"}`)
	return upgraded.Mappable()
}

// newTestProviderServer returns a configured provider server whose API URL points at a
// local server with the given handler. The provider gets dummy credentials.
func newTestProviderServer(t *testing.T, handler http.HandlerFunc) pulumirpc.ResourceProviderServer {
	t.Helper()
	api := httptest.NewServer(handler)
	t.Cleanup(api.Close)

	// The provider falls back to these variables for unset configuration, so they must
	// not point anywhere else either.
	t.Setenv("CHECKLY_API_URL", api.URL)
	t.Setenv("CHECKLY_API_KEY", "")
	t.Setenv("CHECKLY_ACCOUNT_ID", "")

	ctx := context.Background()
	srv, err := pfbridge.NewProviderServer(ctx, testLogSink{t}, Provider(),
		pfbridge.ProviderMetadata{PackageSchema: testPackageSchema})
	if err != nil {
		t.Fatalf("creating provider server: %v", err)
	}

	config := marshal(t, map[string]any{
		"apiKey":    "dummy-api-key",
		"accountId": "00000000-0000-4000-8000-000000000000",
		"apiUrl":    api.URL,
	})
	if _, err := srv.Configure(ctx, &pulumirpc.ConfigureRequest{
		Args:            config,
		AcceptSecrets:   true,
		AcceptResources: true,
	}); err != nil {
		t.Fatalf("Configure: %v", err)
	}
	return srv
}

func runStateCompatCase(t *testing.T, tc stateCompatCase) {
	ctx := context.Background()
	srv := newTestProviderServer(t, func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("unexpected Checkly API request: %s %s", r.Method, r.URL.Path)
		w.WriteHeader(http.StatusInternalServerError)
	})

	urn := resource.NewURN("test", "statecompat", "", tokens.Type(tc.token), tc.name)
	oldInputs := marshal(t, tc.v2Inputs)

	check, err := srv.Check(ctx, &pulumirpc.CheckRequest{
		Urn:        string(urn),
		Type:       tc.token,
		Name:       tc.name,
		Olds:       oldInputs,
		News:       marshal(t, tc.v3Inputs),
		RandomSeed: []byte("statecompat"),
	})
	if err != nil {
		t.Fatalf("Check: %v", err)
	}
	for _, f := range check.GetFailures() {
		t.Errorf("Check failure: %s: %s", f.GetProperty(), f.GetReason())
	}
	if t.Failed() {
		return
	}

	resp, err := srv.Diff(ctx, &pulumirpc.DiffRequest{
		Id:        "1234",
		Urn:       string(urn),
		Type:      tc.token,
		Name:      tc.name,
		Olds:      marshal(t, tc.v2State()),
		News:      check.GetInputs(),
		OldInputs: oldInputs,
	})
	if err != nil {
		t.Fatalf("Diff: %v", err)
	}
	if len(resp.GetReplaces()) > 0 || resp.GetDeleteBeforeReplace() {
		t.Errorf("upgrading would replace the resource, because of %v", resp.GetReplaces())
	}
	t.Logf("changes: %v, diffs: %v", resp.GetChanges(), resp.GetDiffs())
	if tc.wantNoChanges && resp.GetChanges() != pulumirpc.DiffResponse_DIFF_NONE {
		t.Errorf("want no changes, got changes to %v", resp.GetDiffs())
	}
	state := tc.v2State()
	for _, p := range rewrittenProperties {
		if _, stored := state[p]; stored && slices.Contains(resp.GetDiffs(), p) {
			t.Errorf("the upgraded state changes %s", p)
		}
	}
}

func marshal(t *testing.T, m map[string]any) *structpb.Struct {
	t.Helper()
	s, err := plugin.MarshalProperties(resource.NewPropertyMapFromMap(m), plugin.MarshalOptions{
		KeepUnknowns: true,
		KeepSecrets:  true,
	})
	if err != nil {
		t.Fatalf("marshalling properties: %v", err)
	}
	return s
}

// without returns a copy of m without the given keys.
func without(m map[string]any, keys ...string) map[string]any {
	out := maps.Clone(m)
	maps.DeleteFunc(out, func(k string, _ any) bool { return slices.Contains(keys, k) })
	return out
}

// testLogSink routes provider log messages to the test log.
type testLogSink struct{ t *testing.T }

func (s testLogSink) Log(_ context.Context, sev diag.Severity, urn resource.URN, msg string) error {
	s.t.Logf("[%s] %s: %s", sev, urn, msg)
	return nil
}

func (s testLogSink) LogStatus(_ context.Context, sev diag.Severity, urn resource.URN, msg string) error {
	s.t.Logf("[%s status] %s: %s", sev, urn, msg)
	return nil
}

// testCertificate returns a self-signed certificate and its private key, PEM encoded.
func testCertificate(t *testing.T) (string, string) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Fatal(err)
	}
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "statecompat.example.com"},
		NotBefore:    time.Unix(0, 0),
		NotAfter:     time.Unix(0, 0).AddDate(100, 0, 0),
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, tmpl, &key.PublicKey, key)
	if err != nil {
		t.Fatal(err)
	}
	keyDER, err := x509.MarshalPKCS8PrivateKey(key)
	if err != nil {
		t.Fatal(err)
	}
	return string(pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})),
		string(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}))
}

// testCodeBundle writes a code bundle archive and returns its path, together with the
// metadata that the provider stores for it.
func testCodeBundle(t *testing.T) (string, string) {
	t.Helper()
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gz)
	content := []byte("{}\n")
	if err := tw.WriteHeader(&tar.Header{Name: "package.json", Mode: 0o644, Size: int64(len(content))}); err != nil {
		t.Fatal(err)
	}
	if _, err := tw.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(t.TempDir(), "bundle.tar.gz")
	if err := os.WriteFile(file, buf.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256(buf.Bytes())
	metadata := `{"v":7,"s256":"` + hex.EncodeToString(sum[:]) + `","pwv":"1.50.0","pm":"npm"}`
	return file, base64.StdEncoding.EncodeToString([]byte(metadata))
}

func stateCompatCases(t *testing.T) []stateCompatCase {
	certificate, privateKey := testCertificate(t)
	bundleFile, bundleMetadata := testCodeBundle(t)
	locations := []any{"eu-west-1"}

	// v2.x stores alert settings with a list for each nested block; v3 uses objects.
	v2AlertSettings := map[string]any{
		"escalationType":               "RUN_BASED",
		"runBasedEscalations":          []any{map[string]any{"failedRunThreshold": 1}},
		"timeBasedEscalations":         []any{},
		"reminders":                    []any{map[string]any{"amount": 0, "interval": 5}},
		"parallelRunFailureThresholds": []any{map[string]any{"enabled": false, "percentage": 10}},
		"sslCertificates":              []any{},
	}
	v3AlertSettings := map[string]any{
		"escalationType":              "RUN_BASED",
		"runBasedEscalation":          map[string]any{"failedRunThreshold": 1},
		"reminders":                   map[string]any{"amount": 0, "interval": 5},
		"parallelRunFailureThreshold": map[string]any{"enabled": false, "percentage": 10},
	}
	retryStrategy := map[string]any{
		"type":               "FIXED",
		"baseBackoffSeconds": 60,
		"maxDurationSeconds": 600,
		"maxRetries":         2,
		"sameRegion":         true,
	}
	subscriptions := []any{map[string]any{"channelId": 42, "activated": true}}

	// v2.x refreshes alert settings from the API into the state of every check and
	// monitor, even when the program does not set them.
	v2MonitorOutputs := map[string]any{
		"alertSettings": map[string]any{
			"escalationType":               "RUN_BASED",
			"runBasedEscalations":          []any{map[string]any{"failedRunThreshold": 1}},
			"timeBasedEscalations":         []any{},
			"reminders":                    []any{map[string]any{"amount": 0, "interval": 5}},
			"parallelRunFailureThresholds": []any{map[string]any{"enabled": false, "percentage": 10}},
		},
		"useGlobalAlertSettings": true,
	}

	cases := []stateCompatCase{
		{
			name:  "AlertChannel",
			token: "checkly:index/alertChannel:AlertChannel",
			v2Inputs: map[string]any{
				"email":              map[string]any{"address": "alerts@example.com"},
				"sendRecovery":       true,
				"sendFailure":        true,
				"sendDegraded":       false,
				"sslExpiry":          false,
				"sslExpiryThreshold": 30,
			},
			v3Inputs: map[string]any{
				"email":              map[string]any{"address": "alerts@example.com"},
				"sendRecovery":       true,
				"sendFailure":        true,
				"sendDegraded":       false,
				"sslExpiry":          false,
				"sslExpiryThreshold": 30,
			},
		},
		{
			name:  "AlertChannelWebhook",
			token: "checkly:index/alertChannel:AlertChannel",
			v2Inputs: map[string]any{
				"webhook": map[string]any{
					"name":            "hook",
					"url":             "https://hooks.example.com/checkly",
					"method":          "POST",
					"headers":         map[string]any{"X-Source": "checkly"},
					"queryParameters": map[string]any{},
					"template":        "",
					"webhookSecret":   "",
					"webhookType":     "",
				},
				"sendRecovery":       true,
				"sendFailure":        true,
				"sendDegraded":       false,
				"sslExpiry":          false,
				"sslExpiryThreshold": 30,
			},
			v3Inputs: map[string]any{
				"webhook": map[string]any{
					"name":    "hook",
					"url":     "https://hooks.example.com/checkly",
					"method":  "POST",
					"headers": map[string]any{"X-Source": "checkly"},
				},
				"sendRecovery":       true,
				"sendFailure":        true,
				"sendDegraded":       false,
				"sslExpiry":          false,
				"sslExpiryThreshold": 30,
			},
		},
		{
			name:  "CheckAPI",
			token: "checkly:index/check:Check",
			v2Inputs: map[string]any{
				"name":                 "api check",
				"type":                 "API",
				"activated":            true,
				"frequency":            10,
				"frequencyOffset":      0,
				"locations":            locations,
				"privateLocations":     []any{},
				"tags":                 []any{"pulumi"},
				"muted":                false,
				"shouldFail":           false,
				"runParallel":          false,
				"doubleCheck":          false,
				"degradedResponseTime": 10000,
				"maxResponseTime":      20000,
				"groupId":              0,
				"groupOrder":           0,
				"setupSnippetId":       0,
				"teardownSnippetId":    0,
				"localSetupScript":     "",
				"localTeardownScript":  "",
				"runtimeId":            "",
				"description":          "",
				"sslCheck":             false,
				"request": map[string]any{
					"method":          "GET",
					"url":             "https://api.example.com/health",
					"followRedirects": true,
					"skipSsl":         false,
					"headers":         map[string]any{"Accept": "application/json"},
					"queryParameters": map[string]any{},
					"body":            "",
					"bodyType":        "NONE",
					"ipFamily":        "IPv4",
					"assertions": []any{map[string]any{
						"source":     "STATUS_CODE",
						"comparison": "EQUALS",
						"target":     "200",
						"property":   "",
					}},
				},
				"alertSettings":             v2AlertSettings,
				"retryStrategy":             retryStrategy,
				"useGlobalAlertSettings":    false,
				"alertChannelSubscriptions": subscriptions,
				"environmentVariables":      map[string]any{},
				"environmentVariable":       []any{},
			},
			v2Outputs: map[string]any{
				"script": "",
			},
			v3Inputs: map[string]any{
				"name":                 "api check",
				"type":                 "API",
				"activated":            true,
				"frequency":            10,
				"locations":            locations,
				"tags":                 []any{"pulumi"},
				"degradedResponseTime": 10000,
				"maxResponseTime":      20000,
				"request": map[string]any{
					"method":          "GET",
					"url":             "https://api.example.com/health",
					"followRedirects": true,
					"headers":         map[string]any{"Accept": "application/json"},
					"bodyType":        "NONE",
					"ipFamily":        "IPv4",
					"assertions": []any{map[string]any{
						"source":     "STATUS_CODE",
						"comparison": "EQUALS",
						"target":     "200",
					}},
				},
				"alertSettings":             v3AlertSettings,
				"retryStrategy":             retryStrategy,
				"useGlobalAlertSettings":    false,
				"alertChannelSubscriptions": subscriptions,
			},
			wantNoChanges: true,
		},
		{
			name:  "CheckBrowser",
			token: "checkly:index/check:Check",
			v2Inputs: map[string]any{
				"name":                   "browser check",
				"type":                   "BROWSER",
				"activated":              true,
				"frequency":              10,
				"locations":              locations,
				"privateLocations":       []any{},
				"tags":                   []any{},
				"muted":                  false,
				"shouldFail":             false,
				"runParallel":            false,
				"doubleCheck":            false,
				"script":                 "console.log('hello')",
				"runtimeId":              "",
				"sslCheckDomain":         "",
				"groupId":                0,
				"groupOrder":             0,
				"alertSettings":          v2AlertSettings,
				"useGlobalAlertSettings": true,
				"environmentVariables":   map[string]any{},
				"environmentVariable": []any{map[string]any{
					"key": "BASE_URL", "value": "https://example.com", "locked": false, "secret": false,
				}},
			},
			v2Outputs: map[string]any{
				"degradedResponseTime": 15000,
				"maxResponseTime":      30000,
				"description":          "",
				"localSetupScript":     "",
				"localTeardownScript":  "",
				"retryStrategy": map[string]any{
					"type": "NO_RETRIES", "baseBackoffSeconds": 0, "maxRetries": 0,
					"maxDurationSeconds": 0, "sameRegion": false,
				},
			},
			v3Inputs: map[string]any{
				"name":                   "browser check",
				"type":                   "BROWSER",
				"activated":              true,
				"frequency":              10,
				"locations":              locations,
				"script":                 "console.log('hello')",
				"alertSettings":          v3AlertSettings,
				"useGlobalAlertSettings": true,
				"environmentVariables": []any{map[string]any{
					"key": "BASE_URL", "value": "https://example.com", "locked": false, "secret": false,
				}},
			},
			wantNoChanges: true,
		},
		{
			name:  "CheckEnvironmentVariablesMap",
			token: "checkly:index/check:Check",
			v2Inputs: map[string]any{
				"name":                 "env map check",
				"type":                 "BROWSER",
				"activated":            true,
				"frequency":            10,
				"locations":            locations,
				"script":               "console.log('hello')",
				"environmentVariables": map[string]any{"B_VAR": "2", "A_VAR": "1"},
				"environmentVariable":  []any{},
			},
			v2Outputs: map[string]any{
				"alertSettings":          v2MonitorOutputs["alertSettings"],
				"useGlobalAlertSettings": true,
				"degradedResponseTime":   15000,
				"maxResponseTime":        30000,
				"description":            "",
				"doubleCheck":            false,
				"localSetupScript":       "",
				"localTeardownScript":    "",
				"muted":                  false,
				"runParallel":            false,
				"shouldFail":             false,
				"retryStrategy": map[string]any{
					"type": "NO_RETRIES", "baseBackoffSeconds": 0, "maxRetries": 0,
					"maxDurationSeconds": 0, "sameRegion": false,
				},
			},
			v3Inputs: map[string]any{
				"name":      "env map check",
				"type":      "BROWSER",
				"activated": true,
				"frequency": 10,
				"locations": locations,
				"script":    "console.log('hello')",
				"environmentVariables": []any{
					map[string]any{"key": "A_VAR", "value": "1", "locked": false, "secret": false},
					map[string]any{"key": "B_VAR", "value": "2", "locked": false, "secret": false},
				},
			},
			wantNoChanges: true,
		},
		{
			name:  "CheckGroup",
			token: "checkly:index/checkGroup:CheckGroup",
			v2Inputs: map[string]any{
				"name":             "group",
				"activated":        true,
				"muted":            false,
				"concurrency":      3,
				"locations":        locations,
				"privateLocations": []any{},
				"tags":             []any{"pulumi"},
				"runParallel":      false,
				"doubleCheck":      false,
				"runtimeId":        "",
				"setupSnippetId":   0,
				"apiCheckDefaults": map[string]any{
					"url":             "",
					"headers":         map[string]any{},
					"queryParameters": map[string]any{},
					"assertions":      []any{},
				},
				"alertSettings":             v2AlertSettings,
				"retryStrategy":             retryStrategy,
				"useGlobalAlertSettings":    false,
				"alertChannelSubscriptions": subscriptions,
				"environmentVariables":      map[string]any{},
				"environmentVariable": []any{map[string]any{
					"key": "TOKEN", "value": "secret", "locked": true, "secret": false,
				}},
			},
			v2Outputs: map[string]any{
				"localSetupScript":    "",
				"localTeardownScript": "",
			},
			v3Inputs: map[string]any{
				"name":                      "group",
				"activated":                 true,
				"concurrency":               3,
				"locations":                 locations,
				"tags":                      []any{"pulumi"},
				"alertSettings":             v3AlertSettings,
				"retryStrategy":             retryStrategy,
				"useGlobalAlertSettings":    false,
				"alertChannelSubscriptions": subscriptions,
				"environmentVariables": []any{map[string]any{
					"key": "TOKEN", "value": "secret", "locked": true, "secret": false,
				}},
			},
			wantNoChanges: true,
		},
		{
			name:  "CheckGroupV2",
			token: "checkly:index/checkGroupV2:CheckGroupV2",
			v2Inputs: map[string]any{
				"name":        "group v2",
				"activated":   true,
				"muted":       false,
				"concurrency": 2,
				"tags":        []any{"pulumi"},
				"enforceLocations": map[string]any{
					"enabled":          true,
					"locations":        locations,
					"privateLocations": []any{},
				},
				"enforceAlertSettings": map[string]any{
					"enabled":                   true,
					"useGlobalAlertSettings":    true,
					"alertChannelSubscriptions": subscriptions,
				},
				"enforceSchedulingStrategy": map[string]any{
					"enabled":     true,
					"runParallel": true,
				},
				"apiCheckDefaults": map[string]any{
					"url":             "https://api.example.com",
					"headers":         map[string]any{"Accept": "application/json"},
					"queryParameters": map[string]any{},
					"assertions":      []any{},
				},
				"environmentVariables": []any{map[string]any{
					"key": "BASE_URL", "value": "https://example.com", "locked": false, "secret": false,
				}},
			},
			v2Outputs: map[string]any{
				"enforceRetryStrategy": map[string]any{"enabled": false},
			},
			v3Inputs: map[string]any{
				"name":        "group v2",
				"activated":   true,
				"concurrency": 2,
				"tags":        []any{"pulumi"},
				"enforceLocations": map[string]any{
					"enabled":   true,
					"locations": locations,
				},
				"enforceAlertSettings": map[string]any{
					"enabled":                   true,
					"useGlobalAlertSettings":    true,
					"alertChannelSubscriptions": subscriptions,
				},
				"enforceSchedulingStrategy": map[string]any{
					"enabled":     true,
					"runParallel": true,
				},
				"apiCheckDefaults": map[string]any{
					"url":     "https://api.example.com",
					"headers": map[string]any{"Accept": "application/json"},
				},
				"environmentVariables": []any{map[string]any{
					"key": "BASE_URL", "value": "https://example.com", "locked": false, "secret": false,
				}},
			},
			wantNoChanges: true,
		},
		{
			name:  "CheckGroupV2AlertSettings",
			token: "checkly:index/checkGroupV2:CheckGroupV2",
			v2Inputs: map[string]any{
				"name": "group v2 alerts",
				"enforceAlertSettings": map[string]any{
					"enabled":                   true,
					"useGlobalAlertSettings":    false,
					"alertChannelSubscriptions": subscriptions,
					"alertSettings":             v2AlertSettings,
				},
			},
			v2Outputs: map[string]any{
				"activated":   true,
				"muted":       false,
				"concurrency": 1,
				"tags":        []any{},
				"apiCheckDefaults": map[string]any{
					"url": "", "headers": map[string]any{}, "queryParameters": map[string]any{},
					"assertions": []any{},
				},
				"enforceLocations":          map[string]any{"enabled": false},
				"enforceRetryStrategy":      map[string]any{"enabled": false},
				"enforceSchedulingStrategy": map[string]any{"enabled": false, "runParallel": false},
			},
			v3Inputs: map[string]any{
				"name": "group v2 alerts",
				"enforceAlertSettings": map[string]any{
					"enabled":                   true,
					"useGlobalAlertSettings":    false,
					"alertChannelSubscriptions": subscriptions,
					"alertSettings":             v3AlertSettings,
				},
			},
			wantNoChanges: true,
		},
		{
			name:  "ClientCertificate",
			token: "checkly:index/clientCertificate:ClientCertificate",
			v2Inputs: map[string]any{
				"host":        "*.example.com",
				"certificate": certificate,
				"privateKey":  privateKey,
			},
			v3Inputs: map[string]any{
				"host":        "*.example.com",
				"certificate": certificate,
				"privateKey":  privateKey,
			},
			wantNoChanges: true,
		},
		{
			name:     "Dashboard",
			token:    "checkly:index/dashboard:Dashboard",
			v2Inputs: map[string]any{"customUrl": "statecompat", "header": "Status"},
			v3Inputs: map[string]any{"customUrl": "statecompat", "header": "Status"},
		},
		{
			name:  "DnsMonitor",
			token: "checkly:index/dnsMonitor:DnsMonitor",
			v2Inputs: map[string]any{
				"name": "dns", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"query": "example.com", "recordType": "A"},
			},
			v2Outputs: v2MonitorOutputs,
			v3Inputs: map[string]any{
				"name": "dns", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"query": "example.com", "recordType": "A"},
			},
		},
		{
			name:     "EnvironmentVariable",
			token:    "checkly:index/environmentVariable:EnvironmentVariable",
			v2Inputs: map[string]any{"key": "STATECOMPAT", "value": "value"},
			v3Inputs: map[string]any{"key": "STATECOMPAT", "value": "value"},
		},
		{
			name:  "GrpcMonitor",
			token: "checkly:index/grpcMonitor:GrpcMonitor",
			v2Inputs: map[string]any{
				"name": "grpc", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"host": "grpc.example.com", "port": 443, "grpcMode": "HEALTH"},
			},
			v2Outputs: map[string]any{
				"alertSettings":          v2MonitorOutputs["alertSettings"],
				"useGlobalAlertSettings": true,
				"shouldFail":             false,
				"retryStrategy": map[string]any{
					"type": "NO_RETRIES", "baseBackoffSeconds": 0, "maxDurationSeconds": 0,
					"maxRetries": 0, "sameRegion": false, "onlyOn": nil,
				},
			},
			v3Inputs: map[string]any{
				"name": "grpc", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"host": "grpc.example.com", "port": 443, "grpcMode": "HEALTH"},
			},
		},
		{
			name:  "HeartbeatCheck",
			token: "checkly:index/heartbeatCheck:HeartbeatCheck",
			v2Inputs: map[string]any{
				"name": "heartbeat", "activated": true,
				"heartbeat": map[string]any{"period": 7, "periodUnit": "days", "grace": 1, "graceUnit": "days"},
			},
			v2Outputs: map[string]any{
				"description":            "",
				"muted":                  false,
				"alertSettings":          v2MonitorOutputs["alertSettings"],
				"useGlobalAlertSettings": true,
				"heartbeat": map[string]any{
					"period": 7, "periodUnit": "days", "grace": 1, "graceUnit": "days", "pingToken": "token",
				},
			},
			v3Inputs: map[string]any{
				"name": "heartbeat", "activated": true,
				"heartbeat": map[string]any{"period": 7, "periodUnit": "days", "grace": 1, "graceUnit": "days"},
			},
			wantNoChanges: true,
		},
		{
			name:  "HeartbeatMonitor",
			token: "checkly:index/heartbeatMonitor:HeartbeatMonitor",
			v2Inputs: map[string]any{
				"name": "heartbeat monitor", "activated": true,
				"heartbeat": map[string]any{"period": 7, "periodUnit": "days", "grace": 1, "graceUnit": "days"},
			},
			v2Outputs: map[string]any{
				"description":            "",
				"muted":                  false,
				"alertSettings":          v2MonitorOutputs["alertSettings"],
				"useGlobalAlertSettings": true,
				"heartbeat": map[string]any{
					"period": 7, "periodUnit": "days", "grace": 1, "graceUnit": "days", "pingToken": "token",
				},
			},
			v3Inputs: map[string]any{
				"name": "heartbeat monitor", "activated": true,
				"heartbeat": map[string]any{"period": 7, "periodUnit": "days", "grace": 1, "graceUnit": "days"},
			},
			wantNoChanges: true,
		},
		{
			name:  "IcmpMonitor",
			token: "checkly:index/icmpMonitor:IcmpMonitor",
			v2Inputs: map[string]any{
				"name": "icmp", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com"},
			},
			v2Outputs: v2MonitorOutputs,
			v3Inputs: map[string]any{
				"name": "icmp", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com"},
			},
		},
		{
			name:  "MaintenanceWindow",
			token: "checkly:index/maintenanceWindow:MaintenanceWindow",
			v2Inputs: map[string]any{
				"name":              "maintenance",
				"startsAt":          "2030-01-01T00:00:00.000Z",
				"endsAt":            "2030-01-01T01:00:00.000Z",
				"repeatUnit":        "DAY",
				"repeatInterval":    1,
				"repeatEndsAt":      "2030-03-01T00:00:00.000Z",
				"tags":              []any{"pulumi"},
				"silenceAlertsTags": []any{},
				"pauseAllChecks":    false,
				"silenceAllAlerts":  false,
				"timezone":          "",
				"description":       "",
			},
			v3Inputs: map[string]any{
				"name":           "maintenance",
				"startsAt":       "2030-01-01T00:00:00.000Z",
				"endsAt":         "2030-01-01T01:00:00.000Z",
				"repeatUnit":     "DAY",
				"repeatInterval": 1,
				"repeatEndsAt":   "2030-03-01T00:00:00.000Z",
				"tags":           []any{"pulumi"},
			},
			wantNoChanges: true,
		},
		{
			name:  "PlaywrightCheckSuite",
			token: "checkly:index/playwrightCheckSuite:PlaywrightCheckSuite",
			v2Inputs: map[string]any{
				"name": "suite", "activated": true, "frequency": 10, "locations": locations,
				"bundle": map[string]any{"id": "bundle-id", "metadata": bundleMetadata},
			},
			v2Outputs: v2MonitorOutputs,
			v3Inputs: map[string]any{
				"name": "suite", "activated": true, "frequency": 10, "locations": locations,
				"bundle": map[string]any{"id": "bundle-id", "metadata": bundleMetadata},
			},
		},
		{
			name:          "PlaywrightCodeBundle",
			token:         "checkly:index/playwrightCodeBundle:PlaywrightCodeBundle",
			v2Inputs:      map[string]any{"prebuiltArchive": map[string]any{"file": bundleFile}},
			v2Outputs:     map[string]any{"metadata": bundleMetadata},
			v3Inputs:      map[string]any{"prebuiltArchive": map[string]any{"file": bundleFile}},
			wantNoChanges: true,
		},
		{
			name:     "PrivateLocation",
			token:    "checkly:index/privateLocation:PrivateLocation",
			v2Inputs: map[string]any{"name": "location", "slugName": "statecompat"},
			v3Inputs: map[string]any{"name": "location", "slugName": "statecompat"},
		},
		{
			name:          "Snippet",
			token:         "checkly:index/snippet:Snippet",
			v2Inputs:      map[string]any{"name": "snippet", "script": "console.log('hi')"},
			v3Inputs:      map[string]any{"name": "snippet", "script": "console.log('hi')"},
			wantNoChanges: true,
		},
		{
			name:  "SslMonitor",
			token: "checkly:index/sslMonitor:SslMonitor",
			v2Inputs: map[string]any{
				"name": "ssl", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com"},
			},
			v2Outputs: v2MonitorOutputs,
			v3Inputs: map[string]any{
				"name": "ssl", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com"},
			},
		},
		{
			name:  "StatusPage",
			token: "checkly:index/statusPage:StatusPage",
			v2Inputs: map[string]any{
				"name": "status", "url": "statecompat",
				"cards": []any{map[string]any{
					"name":               "card",
					"serviceAttachments": []any{map[string]any{"serviceId": "service-id"}},
				}},
			},
			v3Inputs: map[string]any{
				"name": "status", "url": "statecompat",
				"cards": []any{map[string]any{
					"name":               "card",
					"serviceAttachments": []any{map[string]any{"serviceId": "service-id"}},
				}},
			},
		},
		{
			name:          "StatusPageService",
			token:         "checkly:index/statusPageService:StatusPageService",
			v2Inputs:      map[string]any{"name": "service"},
			v3Inputs:      map[string]any{"name": "service"},
			wantNoChanges: true,
		},
		{
			name:     "StatusPageV3",
			token:    "checkly:index/statusPageV3:StatusPageV3",
			v2Inputs: map[string]any{"name": "status v3", "url": "statecompat-v3"},
			v3Inputs: map[string]any{"name": "status v3", "url": "statecompat-v3"},
		},
		{
			name:  "StatusPageV3AutomationRule",
			token: "checkly:index/statusPageV3AutomationRule:StatusPageV3AutomationRule",
			v2Inputs: map[string]any{
				"statusPageId": "page-id", "tags": []any{"pulumi"},
				"firstUpdate": "Investigating", "lastUpdate": "Resolved",
			},
			v3Inputs: map[string]any{
				"statusPageId": "page-id", "tags": []any{"pulumi"},
				"firstUpdate": "Investigating", "lastUpdate": "Resolved",
			},
		},
		{
			name:     "StatusPageV3Component",
			token:    "checkly:index/statusPageV3Component:StatusPageV3Component",
			v2Inputs: map[string]any{"name": "component", "statusPageId": "page-id", "displayOrder": 1},
			v3Inputs: map[string]any{"name": "component", "statusPageId": "page-id", "displayOrder": 1},
		},
		{
			name:  "TcpCheck",
			token: "checkly:index/tcpCheck:TcpCheck",
			v2Inputs: map[string]any{
				"name": "tcp", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com", "port": 443},
			},
			v2Outputs: v2MonitorOutputs,
			v3Inputs: map[string]any{
				"name": "tcp", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com", "port": 443},
			},
		},
		{
			name:  "TcpMonitor",
			token: "checkly:index/tcpMonitor:TcpMonitor",
			v2Inputs: map[string]any{
				"name": "tcp monitor", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com", "port": 443},
			},
			v2Outputs: map[string]any{
				"alertSettings":          v2MonitorOutputs["alertSettings"],
				"useGlobalAlertSettings": true,
				// v2.x had an onlyOn object without conditions, which the current schema dropped.
				"retryStrategy": map[string]any{
					"type": "FIXED", "baseBackoffSeconds": 60, "maxRetries": 2,
					"maxDurationSeconds": 600, "sameRegion": true, "onlyOn": map[string]any{},
				},
			},
			v3Inputs: map[string]any{
				"name": "tcp monitor", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"hostname": "example.com", "port": 443},
			},
		},
		{
			name:  "TracerouteMonitor",
			token: "checkly:index/tracerouteMonitor:TracerouteMonitor",
			v2Inputs: map[string]any{
				"name": "traceroute", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"url": "example.com"},
			},
			v2Outputs: v2MonitorOutputs,
			v3Inputs: map[string]any{
				"name": "traceroute", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"url": "example.com"},
			},
		},
		{
			name:  "UrlMonitor",
			token: "checkly:index/urlMonitor:UrlMonitor",
			v2Inputs: map[string]any{
				"name": "url", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"url": "https://example.com"},
			},
			v2Outputs: map[string]any{
				"alertSettings":          v2MonitorOutputs["alertSettings"],
				"useGlobalAlertSettings": true,
				"description":            "",
				"degradedResponseTime":   3000,
				"maxResponseTime":        5000,
				"muted":                  false,
				"runParallel":            false,
				"shouldFail":             false,
				"privateLocations":       []any{},
				"tags":                   []any{},
				"groupId":                0,
				"groupOrder":             0,
				"request": map[string]any{
					"url": "https://example.com", "followRedirects": true, "skipSsl": false,
					"ipFamily": "IPv4", "assertions": []any{},
				},
				"retryStrategy": map[string]any{
					"type": "NO_RETRIES", "baseBackoffSeconds": 0, "maxRetries": 0,
					"maxDurationSeconds": 0, "sameRegion": false,
				},
				"alertChannelSubscriptions": []any{},
			},
			v3Inputs: map[string]any{
				"name": "url", "activated": true, "frequency": 10, "locations": locations,
				"request": map[string]any{"url": "https://example.com"},
			},
			wantNoChanges: true,
		},
	}
	return append(cases, aliasCases(t, cases)...)
}

// aliasCases returns cases for moving a resource to its replacement with a Pulumi alias, such
// as `aliases: [{ type: "checkly:index/check:Check" }]` on an ApiCheck. The bridge does not
// call the upstream MoveResourceState, so the new type diffs the stored state of the old one.
// Each case reuses the v2.x state of an existing case under the new type.
func aliasCases(t *testing.T, cases []stateCompatCase) []stateCompatCase {
	t.Helper()
	byName := map[string]stateCompatCase{}
	for _, c := range cases {
		byName[c.name] = c
	}
	withName := func(m map[string]any, name string) map[string]any {
		m["name"] = name
		return m
	}
	alias := func(name, from, token string, v3Inputs map[string]any) stateCompatCase {
		c, ok := byName[from]
		if !ok {
			t.Fatalf("no case %s", from)
		}
		return stateCompatCase{
			name:      name,
			token:     token,
			v2Inputs:  c.v2Inputs,
			v2Outputs: c.v2Outputs,
			v3Inputs:  v3Inputs,
		}
	}

	multistep := byName["CheckBrowser"]
	multistepInputs := maps.Clone(multistep.v2Inputs)
	multistepInputs["type"] = "MULTI_STEP"
	multistepInputs["name"] = "multistep check"
	byName["CheckMultistep"] = stateCompatCase{v2Inputs: multistepInputs, v2Outputs: multistep.v2Outputs}

	aliases := []stateCompatCase{
		alias("AliasCheckToApiCheck", "CheckAPI", "checkly:index/apiCheck:ApiCheck",
			without(byName["CheckAPI"].v3Inputs, "type")),
		alias("AliasCheckToBrowserCheck", "CheckBrowser", "checkly:index/browserCheck:BrowserCheck",
			without(byName["CheckBrowser"].v3Inputs, "type")),
		alias("AliasCheckToMultistepCheck", "CheckMultistep", "checkly:index/multistepCheck:MultistepCheck",
			withName(without(byName["CheckBrowser"].v3Inputs, "type"), "multistep check")),
		alias("AliasCheckGroupToCheckGroupV2", "CheckGroup", "checkly:index/checkGroupV2:CheckGroupV2",
			map[string]any{
				"name":        "group",
				"activated":   true,
				"concurrency": 3,
				"tags":        []any{"pulumi"},
				"environmentVariables": []any{map[string]any{
					"key": "TOKEN", "value": "secret", "locked": true, "secret": false,
				}},
			}),
		alias("AliasHeartbeatCheckToHeartbeatMonitor", "HeartbeatCheck",
			"checkly:index/heartbeatMonitor:HeartbeatMonitor", byName["HeartbeatCheck"].v3Inputs),
		alias("AliasTcpCheckToTcpMonitor", "TcpCheck", "checkly:index/tcpMonitor:TcpMonitor",
			byName["TcpCheck"].v3Inputs),
	}

	// The same moves for state that v3 has already rewritten, approximated by the hook's output
	// at schema version 1. TestRefreshRewritesV2State moves checks from state that a refresh
	// has actually rewritten.
	sources := map[string]string{
		"AliasCheckGroupToCheckGroupV2":         "checkly_check_group",
		"AliasHeartbeatCheckToHeartbeatMonitor": "checkly_heartbeat",
		"AliasTcpCheckToTcpMonitor":             "checkly_tcp_check",
	}
	for _, c := range slices.Clone(aliases) {
		if from, ok := sources[c.name]; ok {
			c.v3ShapeFrom = from
			c.name += "FromV3State"
			aliases = append(aliases, c)
		}
	}
	return aliases
}
