package checkly

import (
	"maps"
	"testing"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/reservedkeys"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// v2CheckState returns state of a Check as v2.x stores it.
func v2CheckState() resource.PropertyMap {
	return resource.NewPropertyMapFromMap(map[string]any{
		"id":       "1234",
		"sslCheck": false,
		"alertSettings": map[string]any{
			"escalationType":               "RUN_BASED",
			"runBasedEscalations":          []any{map[string]any{"failedRunThreshold": 1}},
			"timeBasedEscalations":         []any{},
			"reminders":                    []any{map[string]any{"amount": 0, "interval": 5}},
			"parallelRunFailureThresholds": []any{map[string]any{"enabled": false, "percentage": 10}},
			"sslCertificates":              []any{},
		},
		"environmentVariables": map[string]any{"B": "2", "A": "1"},
		"environmentVariable":  []any{},
	})
}

// TestV2StateUpgradeHooksInstalled checks that every resource with a v2.x state shape exists
// in the mapping and gets the TransformFromState, and that only the per-type check resources
// get the PreStateUpgradeHook.
func TestV2StateUpgradeHooksInstalled(t *testing.T) {
	setTestVersion()
	prov := Provider()
	for name := range v2StateShapes {
		res, ok := prov.Resources[name]
		if !ok {
			t.Errorf("%s is not mapped", name)
			continue
		}
		if res.TransformFromState == nil {
			t.Errorf("%s has no TransformFromState", name)
		}
		if hasHook := res.PreStateUpgradeHook != nil; hasHook != v2StateShapes[name].aliasedCheckState {
			t.Errorf("%s: PreStateUpgradeHook installed is %v, want %v", name, hasHook, !hasHook)
		}
	}
}

func TestUpgradeV2State(t *testing.T) {
	shape := v2StateShapes["checkly_check"]
	state := v2CheckState()
	want := resource.NewPropertyMapFromMap(map[string]any{
		"id": "1234",
		"alertSettings": map[string]any{
			"escalationType":              "RUN_BASED",
			"runBasedEscalation":          map[string]any{"failedRunThreshold": 1},
			"timeBasedEscalation":         nil,
			"reminders":                   map[string]any{"amount": 0, "interval": 5},
			"parallelRunFailureThreshold": map[string]any{"enabled": false, "percentage": 10},
		},
		"environmentVariables": []any{
			map[string]any{"key": "A", "value": "1", "locked": false, "secret": false},
			map[string]any{"key": "B", "value": "2", "locked": false, "secret": false},
		},
	})
	before := state.Copy()

	got := upgradeV2State(shape, state)
	if !got.DeepEquals(want) {
		t.Errorf("got %v, want %v", got, want)
	}
	if !state.DeepEquals(before) {
		t.Errorf("the prior state was modified")
	}
}

func TestUpgradeV2StateKeepsEnvironmentVariableList(t *testing.T) {
	list := []any{map[string]any{"key": "A", "value": "1", "locked": true, "secret": false}}
	state := resource.NewPropertyMapFromMap(map[string]any{
		"environmentVariables": map[string]any{},
		"environmentVariable":  list,
	})
	want := resource.NewPropertyMapFromMap(map[string]any{"environmentVariables": list})

	if got := upgradeV2State(v2StateShapes["checkly_check_group"], state); !got.DeepEquals(want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestUpgradeV2StateKeepsSecrets(t *testing.T) {
	state := resource.PropertyMap{
		"environmentVariables": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
			"A": resource.NewStringProperty("1"),
		})),
	}
	got := upgradeV2State(v2StateShapes["checkly_check"], state)
	if v := got["environmentVariables"]; !v.IsSecret() || !v.SecretValue().Element.IsArray() {
		t.Errorf("got %v, want a secret list", v)
	}
}

// TestUpgradeV2StateKeepsV3State checks that state in the current shape, as v3 writes it,
// passes through every hook unchanged.
func TestUpgradeV2StateKeepsV3State(t *testing.T) {
	alertSettings := map[string]any{
		"escalationType":              "RUN_BASED",
		"runBasedEscalation":          map[string]any{"failedRunThreshold": 1},
		"timeBasedEscalation":         nil,
		"reminders":                   map[string]any{"amount": 0, "interval": 5},
		"parallelRunFailureThreshold": map[string]any{"enabled": false, "percentage": 10},
	}
	state := resource.NewPropertyMapFromMap(map[string]any{
		"id":            "1234",
		"alertSettings": alertSettings,
		"enforceAlertSettings": map[string]any{
			"enabled":       true,
			"alertSettings": alertSettings,
		},
		"environmentVariables": []any{
			map[string]any{"key": "A", "value": "1", "locked": false, "secret": false},
		},
		"retryStrategy": map[string]any{"type": "FIXED", "maxRetries": 2},
		"shouldFail":    false,
	})
	for name, shape := range v2StateShapes {
		// shouldFail is current on resources that do not list it as removed.
		input := state.Copy()
		for _, path := range shape.removed {
			if len(path) == 1 {
				delete(input, path[0])
			}
		}
		if got := upgradeV2State(shape, input); !got.DeepEquals(input) {
			t.Errorf("%s: got %v, want %v", name, got, input)
		}
	}
}

// TestUpgradeV2StateKeepsNewerState checks that state with a newer schema version or a raw
// state delta, which v2.x never writes, is returned as it is, even in the v2.x shape.
func TestUpgradeV2StateKeepsNewerState(t *testing.T) {
	shape := v2StateShapes["checkly_check"]

	newer := v2CheckState()
	newer[reservedkeys.Meta] = resource.NewStringProperty(`{"schema_version":"1"}`)
	if got := upgradeV2State(shape, newer); !got.DeepEquals(newer) {
		t.Errorf("schema version 1: got %v, want %v", got, newer)
	}

	unreadable := v2CheckState()
	unreadable[reservedkeys.Meta] = resource.NewStringProperty(`not json`)
	if got := upgradeV2State(shape, unreadable); !got.DeepEquals(unreadable) {
		t.Errorf("unreadable __meta: got %v, want %v", got, unreadable)
	}

	// v2.x can record version 0 itself, which is rewritten like state without __meta.
	version0 := v2CheckState()
	version0[reservedkeys.Meta] = resource.NewStringProperty(`{"schema_version":"0"}`)
	if got := upgradeV2State(shape, version0); got.DeepEquals(version0) {
		t.Errorf("schema version 0 was not rewritten")
	}

	withDelta := v2CheckState()
	withDelta[reservedkeys.RawStateDelta] = resource.NewObjectProperty(resource.PropertyMap{})
	if got := upgradeV2State(shape, withDelta); !got.DeepEquals(withDelta) {
		t.Errorf("raw state delta: got %v, want %v", got, withDelta)
	}
}

// TestPerTypeCheckHookReadsCheckState checks that the per-type check resources read the state
// that v3 writes for a Check, at schema version 1, as their own version 0, and leave other
// versions alone.
func TestPerTypeCheckHookReadsCheckState(t *testing.T) {
	setTestVersion()
	prov := Provider()
	state := resource.PropertyMap{"id": resource.NewStringProperty("1234")}
	for _, tc := range []struct {
		name                   string
		prior, resource, wantV int64
	}{
		{"checkly_api_check", 1, 0, 0},
		{"checkly_browser_check", 1, 0, 0},
		{"checkly_multistep_check", 1, 0, 0},
		{"checkly_api_check", 0, 0, 0},
		{"checkly_api_check", 2, 0, 2},
	} {
		name, want := tc.name, tc.wantV
		version, got, err := prov.Resources[name].PreStateUpgradeHook(tfbridge.PreStateUpgradeHookArgs{
			PriorState:              state,
			PriorStateSchemaVersion: tc.prior,
			ResourceSchemaVersion:   tc.resource,
		})
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if version != want || !got.DeepEquals(state) {
			t.Errorf("%s: got version %d and %v, want version %d and the state unchanged", name, version, got, want)
		}
	}
}

// TestUpgradeV2StateRemovesOnlyOn checks that retryStrategy.onlyOn is dropped from the
// monitors that lost it, whatever its stored value, and that the other retry strategy
// fields are kept.
func TestUpgradeV2StateRemovesOnlyOn(t *testing.T) {
	retryStrategy := map[string]any{
		"type": "FIXED", "baseBackoffSeconds": 60, "maxRetries": 2,
		"maxDurationSeconds": 600, "sameRegion": true,
	}
	want := resource.NewPropertyMapFromMap(map[string]any{"retryStrategy": retryStrategy})
	for name, onlyOn := range map[string]any{
		"empty object": map[string]any{},
		"list":         []any{"NETWORK_ERROR"},
	} {
		stored := maps.Clone(retryStrategy)
		stored["onlyOn"] = onlyOn
		state := resource.NewPropertyMapFromMap(map[string]any{"retryStrategy": stored})

		if got := upgradeV2State(v2StateShapes["checkly_tcp_monitor"], state); !got.DeepEquals(want) {
			t.Errorf("%s: got %v, want %v", name, got, want)
		}
	}
}

// TestUpgradeV2StateKeepsSecretAlertSettings checks that secret alert settings, and secret
// lists inside them, stay secret once rewritten into objects.
func TestUpgradeV2StateKeepsSecretAlertSettings(t *testing.T) {
	escalation := resource.NewObjectProperty(resource.PropertyMap{
		"failedRunThreshold": resource.NewNumberProperty(1),
	})
	state := resource.PropertyMap{
		"alertSettings": resource.MakeSecret(resource.NewObjectProperty(resource.PropertyMap{
			"escalationType":      resource.NewStringProperty("RUN_BASED"),
			"runBasedEscalations": resource.MakeSecret(resource.NewArrayProperty([]resource.PropertyValue{escalation})),
		})),
	}

	got := upgradeV2State(v2StateShapes["checkly_url_monitor"], state)

	settings := got["alertSettings"]
	if !settings.IsSecret() || !settings.SecretValue().Element.IsObject() {
		t.Fatalf("alertSettings is %v, want a secret object", settings)
	}
	inner := settings.SecretValue().Element.ObjectValue()
	if _, ok := inner["runBasedEscalations"]; ok {
		t.Errorf("runBasedEscalations was kept: %v", inner)
	}
	if v := inner["runBasedEscalation"]; !v.IsSecret() || !v.SecretValue().Element.DeepEquals(escalation) {
		t.Errorf("runBasedEscalation is %v, want the secret object %v", v, escalation)
	}
}

// TestAliasedCheckSchemaVersions pins the upstream schema versions that the per-type check
// hooks rely on: they read the state of a Check, at checkSchemaVersion, as their own version
// 0. If upstream changes one of these versions, aliasedCheckState in state_upgrade.go and
// the PreStateUpgradeHook it drives must change with it.
func TestAliasedCheckSchemaVersions(t *testing.T) {
	setTestVersion()
	resources := Provider().P.ResourcesMap()
	for name, want := range map[string]int{
		"checkly_check":           checkSchemaVersion,
		"checkly_api_check":       0,
		"checkly_browser_check":   0,
		"checkly_multistep_check": 0,
	} {
		res, ok := resources.GetOk(name)
		if !ok {
			t.Errorf("upstream has no %s", name)
			continue
		}
		if got := res.SchemaVersion(); got != want {
			t.Errorf("upstream %s is at schema version %d, but the aliasedCheckState hook in state_upgrade.go expects %d", name, got, want)
		}
	}
}
