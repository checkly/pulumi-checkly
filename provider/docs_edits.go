package checkly

import (
	"bytes"
	"regexp"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
)

// upgradeGuideURL is the upgrade guide of this package, which explains how Pulumi users
// move resources between types.
const upgradeGuideURL = "https://github.com/checkly/pulumi-checkly#upgrading-to-v3"

// secretAdvice tells Pulumi users how to keep a credential secret. Pulumi keeps a value
// marked as a secret secret end to end: it shows as [secret] in preview output and is
// encrypted in the state, both as an input and as an output.
const secretAdvice = "mark it as a secret, for example with a secret configuration value or your language's " +
	"Pulumi secret helper. Pulumi then masks it in preview output and encrypts it in the state."

// docsReplacements rewrite upstream sentences that describe Terraform's behaviour, so that
// the generated documentation describes Pulumi's. Each pair is an exact phrase of the
// upstream docs and its Pulumi equivalent. No phrase overlaps another, so the order in which
// they apply doesn't matter.
var docsReplacements = [][2]string{
	{
		"so Terraform keeps the last applied value and does not detect changes made outside Terraform",
		"so Pulumi keeps the last applied value and does not detect changes made outside Pulumi",
	},
	{
		"the Checkly CLI or Terraform are not subscribed",
		"the Checkly CLI or Pulumi are not subscribed",
	},
	{
		"the first apply after an import records the configured key and passphrase in Terraform state only",
		"the first update after an import records the configured key and passphrase in the Pulumi state only",
	},
	{"and for the first apply after an import.", "and for the first update after an import."},
	{
		"that apply records the configured values in Terraform state only",
		"that update records the configured values in the Pulumi state only",
	},
	{
		"replace the resource, for example with `terraform apply -replace`.",
		"replace the resource, for example with `pulumi up --replace <urn>`.",
	},
	{
		"a change made outside Terraform, for example in the web app, doesn't show in the plan. " +
			"An imported secret variable has a null value, so the first plan after the import shows an update, " +
			"and applying it overwrites the value in Checkly with the configured one.",
		"a change made outside Pulumi, for example in the web app, doesn't show in the preview. " +
			"An imported secret variable has a null value, so the first preview after the import shows an update, " +
			"and running it overwrites the value in Checkly with the configured one.",
	},
	{"and load it with Terraform's `filebase64()` function", "and pass the file's contents base64-encoded"},
	{
		"An existing check can be switched over to its per-type resource with a `moved` block, without recreating it. " +
			"Before moving a check, see",
		"An existing check can be moved to its per-type resource with a Pulumi alias, without recreating it. " +
			"Before moving a check, see",
	},
	{
		"so an existing group can be switched over with a `moved` block to a `checkly_check_group_v2` resource.",
		"so an existing group can be moved to a `checkly_check_group_v2` resource with a Pulumi alias.",
	},
	{
		"resources can be migrated to `checkly_heartbeat_monitor` using a `moved` block;",
		"resources can be moved to `checkly_heartbeat_monitor` with a Pulumi alias;",
	},
	{
		"resources can be migrated to `checkly_tcp_monitor` using a `moved` block;",
		"resources can be moved to `checkly_tcp_monitor` with a Pulumi alias;",
	},
	{
		"# A group already managed as checkly_check_group switches to checkly_check_group_v2 with a moved block; " +
			"see the v2 upgrade guide.",
		"# A group already managed as checkly_check_group moves to checkly_check_group_v2 with a Pulumi alias; " +
			"see Upgrading to v3 in the pulumi-checkly README.",
	},
	{
		"a plan that changes it fails, also with `terraform apply -replace` or `replace_triggered_by`. " +
			"To use a different slug, add a new `checkly_private_location` resource block with it, " +
			"move your checks and agents to the new location, then remove the old block.",
		"a preview that changes it fails. `pulumi up --replace` skips this check and replaces the location, which " +
			"the API refuses for a location that checks or groups use, and which disconnects its agents, so don't use " +
			"it to change the slug. To use a different slug, add a new `checkly_private_location` resource with it, " +
			"move your checks and agents to the new location, then remove the old resource.",
	},
	{
		"The values are not sensitive and show in plan output. To keep a credential in a header out of plan output, " +
			"take it from a variable declared with `sensitive = true` or wrap it in `sensitive(...)`.",
		"The values are not secret by default: they show in preview output and are stored unencrypted in the state. " +
			"To keep a credential in a header secret, " + secretAdvice,
	},
	{
		"The values are not sensitive and show in plan output. To keep a credential in a parameter out of plan output, " +
			"take it from a variable declared with `sensitive = true` or wrap it in `sensitive(...)`.",
		"The values are not secret by default: they show in preview output and are stored unencrypted in the state. " +
			"To keep a credential in a parameter secret, " + secretAdvice,
	},
	{
		"The value is not sensitive and shows in plan output. To keep a credential out of plan output, " +
			"take it from a variable declared with `sensitive = true` or wrap it in `sensitive(...)`.",
		"The value is not secret by default: it shows in preview output and is stored unencrypted in the state. " +
			"To keep a credential secret, " + secretAdvice,
	},
	{
		"the next apply restores the configured script unless you copy the repaired script into the configuration",
		"the next update restores the configured script unless you copy the repaired script into your program",
	},
	{
		// The bridge renders ignore_changes per language, for example as ignoreChanges.
		"to `ignore_changes`.",
		"to the `ignore_changes` resource option.",
	},
	{
		"absent from the configuration is removed by the next apply, every time.",
		"absent from your program is removed by the next update, every time.",
	},
	{
		"after an import the first plan can show the dependencies reordered to match the configuration; " +
			"applying it once settles the order.",
		"after an import the first preview can show the dependencies reordered to match your program; " +
			"running that update once settles the order.",
	},
	{"so an alias shows as a change on every plan.", "so an alias shows as a change in every preview."},
	{"resets it to its default on the next apply.", "resets it to its default on the next update."},
	{
		"These rules are enforced by the Checkly API at apply time.",
		"These rules are enforced by the Checkly API when the resource is created or updated.",
	},
	{"rules are checked at plan time;", "rules are checked during preview;"},
	{
		// In the Playwright check suite example, which needs the archive provider to convert.
		"# within Terraform, the simplest way is to use the archive_file Data Source\n# from the archive provider.",
		"# in your program, one way is the archive_file data source of the archive\n# provider.",
	},
	{
		"# The proxy URL can hold credentials, so it comes from a sensitive variable.",
		"# The proxy URL can hold credentials, so set it as a secret configuration value.",
	},
	{
		"# metadata, so the first plan after the import shows an in-place update of\n" +
			"# `bundle`, and the apply sends the bundle from the configuration.",
		"# metadata, so the first preview after the import shows an in-place update of\n" +
			"# `bundle`, and the update sends the bundle from your program.",
	},
}

// docsPatternReplacements rewrite upstream text that varies too much for exact phrases.
var docsPatternReplacements = []struct {
	pattern     *regexp.Regexp
	replacement string
}{
	{
		// Links to the deprecated resource types section of the upstream guide, which
		// describes moving with Terraform moved blocks.
		regexp.MustCompile(`\[Deprecated resource types\]\(https://registry\.terraform\.io/[^)]*` +
			`#deprecated-resource-types\) in the v2 upgrade guide`),
		`"Moving to the replacement resources" in [Upgrading to v3](` + upgradeGuideURL + `)`,
	},
	{
		// Links to sections of the upstream guide whose detail only it has. The bridge drops
		// links to the Terraform registry and the upstream repository, so the guide is named
		// and reached through this package's upgrade guide, which links it.
		regexp.MustCompile(`\[(Check groups|Status pages)\]\(https://registry\.terraform\.io/[^)]*` +
			`#[a-z-]+\) in the v2 upgrade guide`),
		`$1 in the Terraform provider's v2 upgrade guide, linked from [Upgrading to v3](` + upgradeGuideURL + `)`,
	},
	{
		// Examples that move a checkly_check to a per-type resource with a moved block. Without
		// the moved block, the converted example would just create a new check.
		regexp.MustCompile(`(?s)\n# Migrating an existing checkly_check[^\n]*\n(?:#[^\n]*\n)*` +
			`resource "[a-z_]+" "[a-z0-9_]+" \{\n.*?\n\}\n\nmoved \{\n.*?\n\}\n`),
		"\n",
	},
}

// movingSection matches the section of a deprecated resource's page that shows how to move
// it to its replacement with a Terraform moved block, up to the generated schema.
var movingSection = regexp.MustCompile(`(?s)## Moving to [^\n]*\n.*?(<!-- schema generated by tfplugindocs -->)`)

// movingSections replace the moved block instructions of the deprecated resources with the
// Pulumi alias that does the same, keyed by upstream docs file.
var movingSections = map[string]string{
	"check.md": "## Moving to the per-type resources\n\n" +
		"An existing check can be moved to the resource of its type, `checkly_api_check`, `checkly_browser_check` " +
		"or `checkly_multistep_check`, without recreating it: declare the new resource with an alias to the old type, " +
		"`checkly:index/check:Check`, and run `pulumi up`. Refresh the stack on v3 before you add the alias. " +
		"The provider can't verify the check type during the move, so pick the resource that matches the check's `type`. " +
		"The per-type resources have no `double_check`, so replace it with the equivalent `retry_strategy` first. " +
		"For a browser or multi-step check whose `should_fail` is `true`, remove `should_fail` and run `pulumi up` " +
		"before the move, which sets it to `false`. " +
		"See \"Moving to the replacement resources\" in [Upgrading to v3](" + upgradeGuideURL + ") for the details.",
	"check_group.md": "## Moving to `checkly_check_group_v2`\n\n" +
		"An existing group can be moved to `checkly_check_group_v2` without recreating it or its checks: declare the new " +
		"resource with an alias to the old type, `checkly:index/checkGroup:CheckGroup`, and run `pulumi up`. Refresh the " +
		"stack on v3 before you add the alias. Its properties have to be translated; see " +
		"[Upgrading to v3](" + upgradeGuideURL + ").",
	"heartbeat.md": "## Moving to `checkly_heartbeat_monitor`\n\n" +
		"An existing `checkly_heartbeat` resource can be moved to `checkly_heartbeat_monitor` without recreating it: " +
		"declare the new resource with the same properties and an alias to the old type, " +
		"`checkly:index/heartbeatCheck:HeartbeatCheck`, and run `pulumi up`. Refresh the stack on v3 before you add " +
		"the alias; see [Upgrading to v3](" + upgradeGuideURL + ").",
	"tcp_check.md": "## Moving to `checkly_tcp_monitor`\n\n" +
		"An existing `checkly_tcp_check` resource can be moved to `checkly_tcp_monitor` without recreating it: " +
		"declare the new resource with the same properties and an alias to the old type, " +
		"`checkly:index/tcpCheck:TcpCheck`, and run `pulumi up`. Refresh the stack on v3 before you add the alias; " +
		"see [Upgrading to v3](" + upgradeGuideURL + ").",
}

// docsEditRules returns the rules above followed by the bridge's default edit rules.
func docsEditRules(defaults []info.DocsEdit) []info.DocsEdit {
	edits := []info.DocsEdit{{
		Path: "*",
		Edit: func(_ string, content []byte) ([]byte, error) {
			for _, r := range docsReplacements {
				content = bytes.ReplaceAll(content, []byte(r[0]), []byte(r[1]))
			}
			for _, r := range docsPatternReplacements {
				content = r.pattern.ReplaceAll(content, []byte(r.replacement))
			}
			return content, nil
		},
	}}
	for file, section := range movingSections {
		edits = append(edits, info.DocsEdit{
			Path: file,
			Edit: func(_ string, content []byte) ([]byte, error) {
				return movingSection.ReplaceAll(content, []byte(section+"\n\n$1")), nil
			},
		})
	}
	// The rules above run first, so that they match the upstream wording before the
	// default rules rewrite it.
	return append(edits, defaults...)
}
