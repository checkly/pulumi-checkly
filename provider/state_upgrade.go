package checkly

import (
	"context"
	"encoding/json"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/reservedkeys"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Up to v2.x, pulumi-checkly bridged a Terraform provider built on the Terraform Plugin SDK
// v2. Its stacks store Pulumi properties in the shape of that provider's schema, at Terraform
// schema version 0, and without a raw Terraform state. The bridge encodes such properties
// with the current schema before the upstream provider upgrades them, and encoding fails
// wherever a property kept its name but changed its shape.
//
// The rewrites below turn those properties into the current shape. They run as the
// TransformFromState of each resource, which the bridge applies to the stored state before it
// uses it in Check, Diff, Read, Update and Delete, and before the state upgrade. Check passes
// the stored inputs instead, which have the same shape. The new inputs of a Diff are not
// transformed, so a Diff against inputs that v2.x stored, which `pulumi refresh
// --run-program` runs, still fails.
//
// State written by v3 is left alone. State at a newer schema version or with a raw state delta
// is returned as it is. The per-type check resources, which v2.x did not have, write their
// state at schema version 0 too; for them, each rewrite recognises the old shape and leaves
// the current shape untouched.

// v2StateShape lists the properties of one resource that v2.x stored in a different shape.
type v2StateShape struct {
	// alertSettings is the path of the alert settings object, if the resource has one.
	// v2.x stored its nested settings as lists of at most one element, under plural names.
	alertSettings []resource.PropertyKey
	// environmentVariables marks resources that stored environment variables both in an
	// environmentVariable list and in a deprecated environmentVariables map of names to
	// values. The current schema has a single environmentVariables list.
	environmentVariables bool
	// removed are the paths of properties that the current schema no longer has.
	removed [][]resource.PropertyKey
	// aliasedCheckState marks the per-type check resources, at schema version 0, which can
	// receive the state of a Check through a Pulumi alias. The bridge does not call the
	// upstream MoveResourceState, and it refuses state at a newer schema version than the
	// resource has, so the state that v3 writes for a Check, at schema version 1, is read
	// as version 0. Upstream moves a Check by copying its attributes one to one and
	// dropping those that the target does not have, which is what reading it as version 0
	// does.
	//
	// The version is mapped by a PreStateUpgradeHook, which the bridge calls with the
	// schema versions. A Check whose state still has the v2.x shape, at schema version 0,
	// can be aliased too, but then skips the upstream normalization of v2.x state, which
	// keeps zero values such as groupId 0. The upgrade notes therefore ask for a refresh on
	// v3 before an alias is added.
	aliasedCheckState bool
}

// checkSchemaVersion is the schema version of the state that v3 writes for a Check.
const checkSchemaVersion = 1

var (
	alertSettingsPath = []resource.PropertyKey{"alertSettings"}
	onlyOnPath        = []resource.PropertyKey{"retryStrategy", "onlyOn"}
	shouldFailPath    = []resource.PropertyKey{"shouldFail"}

	// checkShape is the v2.x shape of checkly_check.
	checkShape = v2StateShape{
		alertSettings:        alertSettingsPath,
		environmentVariables: true,
		removed:              [][]resource.PropertyKey{{"sslCheck"}},
	}
	// perTypeCheckShape is the shape of the per-type check resources. A Check moved to one
	// of them with a Pulumi alias brings its stored state along.
	perTypeCheckShape = withAliasedCheckState(checkShape)
	monitorShape      = v2StateShape{alertSettings: alertSettingsPath}
	// monitorWithoutOnlyOnShape is the shape of monitors whose retry strategy lost onlyOn.
	monitorWithoutOnlyOnShape = v2StateShape{
		alertSettings: alertSettingsPath,
		removed:       [][]resource.PropertyKey{onlyOnPath},
	}
	// monitorWithoutShouldFailShape is the shape of monitors that also lost shouldFail.
	monitorWithoutShouldFailShape = v2StateShape{
		alertSettings: alertSettingsPath,
		removed:       [][]resource.PropertyKey{onlyOnPath, shouldFailPath},
	}
)

// withAliasedCheckState returns shape with aliasedCheckState set.
func withAliasedCheckState(shape v2StateShape) v2StateShape {
	shape.aliasedCheckState = true
	return shape
}

// v2StateShapes maps Terraform resource names to the properties that need rewriting.
var v2StateShapes = map[string]v2StateShape{
	"checkly_check":           checkShape,
	"checkly_api_check":       perTypeCheckShape,
	"checkly_browser_check":   perTypeCheckShape,
	"checkly_multistep_check": perTypeCheckShape,
	"checkly_check_group": {
		alertSettings:        alertSettingsPath,
		environmentVariables: true,
	},
	// A CheckGroup moved to CheckGroupV2 with a Pulumi alias brings its environment variables
	// along. Its top-level alertSettings is dropped by the bridge, as CheckGroupV2 has none.
	"checkly_check_group_v2": {
		alertSettings:        []resource.PropertyKey{"enforceAlertSettings", "alertSettings"},
		environmentVariables: true,
	},
	"checkly_dns_monitor":            monitorWithoutOnlyOnShape,
	"checkly_grpc_monitor":           monitorWithoutShouldFailShape,
	"checkly_heartbeat":              monitorShape,
	"checkly_heartbeat_monitor":      monitorShape,
	"checkly_icmp_monitor":           monitorWithoutOnlyOnShape,
	"checkly_playwright_check_suite": monitorShape,
	"checkly_ssl_monitor":            monitorWithoutShouldFailShape,
	"checkly_tcp_check":              monitorWithoutOnlyOnShape,
	"checkly_tcp_monitor":            monitorWithoutOnlyOnShape,
	"checkly_traceroute_monitor":     monitorWithoutShouldFailShape,
	"checkly_url_monitor":            monitorShape,
}

// v2AlertSettingsRenames maps the v2.x names of the nested alert settings, each a list of at
// most one element, to their current names, where each is a single object.
var v2AlertSettingsRenames = map[resource.PropertyKey]resource.PropertyKey{
	"reminders":                    "reminders",
	"runBasedEscalations":          "runBasedEscalation",
	"timeBasedEscalations":         "timeBasedEscalation",
	"parallelRunFailureThresholds": "parallelRunFailureThreshold",
}

// applyV2StateUpgradeHooks installs the rewrites on every resource whose v2.x state needs
// them, and the schema version mapping on the per-type check resources.
func applyV2StateUpgradeHooks(prov *tfbridge.ProviderInfo) {
	for name, shape := range v2StateShapes {
		res, ok := prov.Resources[name]
		if !ok {
			continue
		}
		res.TransformFromState = func(_ context.Context, state resource.PropertyMap) (resource.PropertyMap, error) {
			return upgradeV2State(shape, state), nil
		}
		if shape.aliasedCheckState {
			res.PreStateUpgradeHook = readCheckStateAsVersion0
		}
	}
}

// readCheckStateAsVersion0 reads the state that v3 writes for a Check, at checkSchemaVersion,
// as version 0 of a per-type check resource; see aliasedCheckState.
func readCheckStateAsVersion0(args tfbridge.PreStateUpgradeHookArgs) (int64, resource.PropertyMap, error) {
	version := args.PriorStateSchemaVersion
	if version == checkSchemaVersion && args.ResourceSchemaVersion == 0 {
		version = 0
	}
	return version, args.PriorState, nil
}

// upgradeV2State returns state with the properties of shape rewritten from their v2.x shape.
// State that v2.x cannot have written is returned unchanged.
func upgradeV2State(shape v2StateShape, state resource.PropertyMap) resource.PropertyMap {
	if !mayBeV2State(state) {
		return state
	}

	state = state.Copy()
	if shape.alertSettings != nil {
		updateAt(state, shape.alertSettings, upgradeV2AlertSettings)
	}
	if shape.environmentVariables {
		upgradeV2EnvironmentVariables(state)
	}
	for _, path := range shape.removed {
		deleteAt(state, path)
	}
	return state
}

// mayBeV2State reports whether v2.x can have written state: it has no raw state delta, and its
// __meta property records no schema version or version 0. A __meta that can't be read is left
// for the bridge to report.
func mayBeV2State(state resource.PropertyMap) bool {
	if _, ok := state[reservedkeys.RawStateDelta]; ok {
		return false
	}
	meta, ok := state[reservedkeys.Meta]
	if !ok {
		return true
	}
	if !meta.IsString() {
		return false
	}
	var parsed struct {
		SchemaVersion string `json:"schema_version"`
	}
	if err := json.Unmarshal([]byte(meta.StringValue()), &parsed); err != nil {
		return false
	}
	return parsed.SchemaVersion == "" || parsed.SchemaVersion == "0"
}

// upgradeV2AlertSettings unwraps the nested alert settings that v2.x stored as lists and
// moves them to their current names. sslCertificates, which the API ignored, is dropped.
// Settings that are already objects are left as they are.
func upgradeV2AlertSettings(v resource.PropertyValue) resource.PropertyValue {
	return mapObject(v, func(m resource.PropertyMap) {
		for oldKey, newKey := range v2AlertSettingsRenames {
			old, ok := m[oldKey]
			if !ok || !isArray(old) {
				continue
			}
			delete(m, oldKey)
			m[newKey] = firstElement(old)
		}
		delete(m, "sslCertificates")
	})
}

// upgradeV2EnvironmentVariables merges the environment variables of v2.x state into the
// current environmentVariables list. v2.x kept at most one of its environmentVariable list
// and environmentVariables map non-empty. A non-empty map becomes list elements that are
// neither locked nor secret, sorted by name because the map has no order, as the upstream
// provider upgrades its own state. Otherwise the environmentVariable list is used. A list
// already stored under environmentVariables is the current shape and is kept.
func upgradeV2EnvironmentVariables(state resource.PropertyMap) {
	list, hasList := state["environmentVariable"]
	delete(state, "environmentVariable")

	current := state["environmentVariables"]
	if isArray(current) {
		return
	}

	envMap, secret := unwrapSecret(current)
	var elements []resource.PropertyValue
	if envMap.IsObject() {
		obj := envMap.ObjectValue()
		for _, k := range obj.StableKeys() {
			elements = append(elements, resource.NewObjectProperty(resource.PropertyMap{
				"key":    resource.NewStringProperty(string(k)),
				"value":  obj[k],
				"locked": resource.NewBoolProperty(false),
				"secret": resource.NewBoolProperty(false),
			}))
		}
	}

	switch {
	case len(elements) > 0 && secret:
		state["environmentVariables"] = resource.MakeSecret(resource.NewArrayProperty(elements))
	case len(elements) > 0:
		state["environmentVariables"] = resource.NewArrayProperty(elements)
	case hasList:
		state["environmentVariables"] = list
	default:
		delete(state, "environmentVariables")
	}
}

// updateAt replaces the value at path in m with f applied to it, if the path exists. Objects
// along the path are copied, so values shared with other property maps are not modified.
func updateAt(m resource.PropertyMap, path []resource.PropertyKey, f func(resource.PropertyValue) resource.PropertyValue) {
	v, ok := m[path[0]]
	if !ok {
		return
	}
	if len(path) == 1 {
		m[path[0]] = f(v)
		return
	}
	m[path[0]] = mapObject(v, func(inner resource.PropertyMap) {
		updateAt(inner, path[1:], f)
	})
}

// deleteAt removes the value at path in m, if the path exists.
func deleteAt(m resource.PropertyMap, path []resource.PropertyKey) {
	last := len(path) - 1
	if last == 0 {
		delete(m, path[0])
		return
	}
	updateAt(m, path[:last], func(v resource.PropertyValue) resource.PropertyValue {
		return mapObject(v, func(inner resource.PropertyMap) { delete(inner, path[last]) })
	})
}

// mapObject applies f to a copy of the object held by v, keeping v secret if it was. Values
// that are not objects are returned unchanged.
func mapObject(v resource.PropertyValue, f func(resource.PropertyMap)) resource.PropertyValue {
	inner, secret := unwrapSecret(v)
	if !inner.IsObject() {
		return v
	}
	m := inner.ObjectValue().Copy()
	f(m)
	out := resource.NewObjectProperty(m)
	if secret {
		out = resource.MakeSecret(out)
	}
	return out
}

// firstElement returns the first element of a list, or null for an empty list.
func firstElement(v resource.PropertyValue) resource.PropertyValue {
	inner, secret := unwrapSecret(v)
	if a := inner.ArrayValue(); len(a) > 0 {
		if secret {
			return resource.MakeSecret(a[0])
		}
		return a[0]
	}
	return resource.NewNullProperty()
}

func isArray(v resource.PropertyValue) bool {
	inner, _ := unwrapSecret(v)
	return inner.IsArray()
}

func unwrapSecret(v resource.PropertyValue) (resource.PropertyValue, bool) {
	if v.IsSecret() {
		return v.SecretValue().Element, true
	}
	return v, false
}
