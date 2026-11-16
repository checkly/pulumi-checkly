// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package checkly

import (
	"os"
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	checklyprovider "github.com/checkly/terraform-provider-checkly/v2/provider" // Import the upstream provider

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"

	"github.com/checkly/pulumi-checkly/provider/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "checkly"
	// modules:
	mainMod = "index" // the checkly module
)

//go:embed cmd/pulumi-resource-checkly/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		// Instantiate the upstream provider, which is built on the Terraform Plugin Framework.
		// Upstream only reports the version it is given in its Metadata response, which the
		// bridge does not read, and its API client does not send it, so a fixed value is used.
		P: pfbridge.ShimProvider(checklyprovider.New("dev")()),

		Name:    "checkly",
		Version: version.Version,
		// DisplayName is a way to be able to change the casing of the provider name when being
		// displayed on the Pulumi registry
		DisplayName: "Checkly",
		// Change this to your personal name (or a company name) that you would like to be shown in
		// the Pulumi Registry if this package is published there.
		Publisher: "Checkly",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an PNG logo (100x100) for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/checkly/pulumi-checkly/main/assets/checkly.png",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g. https://github.com/org/pulumi-provider-name/releases/download/v${VERSION}/
		PluginDownloadURL: "github://api.github.com/checkly",
		Description:       "A Pulumi package for creating and managing Checkly monitoring resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "checkly", "category/monitoring"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com/registry/packages/checkly",
		Repository: "https://github.com/checkly/pulumi-checkly",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this should
		// match the TF provider module's require directive, not any replace directives.
		GitHubOrg: "checkly",
		// The upstream Go module path carries a /v2 suffix, which the docs lookup needs.
		TFProviderModuleVersion: "v2",
		// UpstreamRepoPath overrides where tfgen reads the upstream docs. When it is empty,
		// tfgen locates them in the Go module cache.
		UpstreamRepoPath: os.Getenv("CHECKLY_UPSTREAM_REPO_PATH"),
		MetadataInfo:     tfbridge.NewProviderMetadata(metadata),
		// checkly_maintenance_windows is a deprecated alias of checkly_maintenance_window, with
		// the same schema. Up to v2.x, MaintenanceWindow mapped checkly_maintenance_windows;
		// it now maps checkly_maintenance_window, so existing programs and stacks keep working,
		// and no stack can hold a resource of the alias. It is left unmapped rather than added
		// as a new resource that is deprecated from the start.
		IgnoreMappings: []string{"checkly_maintenance_windows"},
		Config:         map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*info.Resource{
			// Automatic token maping will ignore resources that are manually mapped.
			"checkly_alert_channel": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "AlertChannel"),
			},
			"checkly_api_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "ApiCheck"),
			},
			"checkly_browser_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "BrowserCheck"),
			},
			"checkly_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Check"),
				DeprecationMessage: "Use the ApiCheck, BrowserCheck or MultistepCheck resource instead, which expose only the properties of their check type. " +
					"An existing check can be moved to one of them with an alias, without recreating it. When moving an API check, note that ApiCheck " +
					"defaults degradedResponseTime and maxResponseTime to 5000 and 20000 instead of 15000 and 30000, and request.followRedirects " +
					"to true instead of false; set them to keep the previous behaviour.",
			},
			"checkly_check_group": {
				Tok:                tfbridge.MakeResource(mainPkg, mainMod, "CheckGroup"),
				DeprecationMessage: "Use the CheckGroupV2 resource instead.",
			},
			"checkly_check_group_v2": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CheckGroupV2"),
			},
			"checkly_dashboard": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dashboard"),
			},
			"checkly_environment_variable": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariable"),
			},
			"checkly_heartbeat": {
				Tok:                tfbridge.MakeResource(mainPkg, mainMod, "HeartbeatCheck"),
				DeprecationMessage: "Use the HeartbeatMonitor resource instead. Both resources have the same properties.",
				CSharpName:         "HeartbeatCheck",
			},
			"checkly_heartbeat_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "HeartbeatMonitor"),
			},
			"checkly_maintenance_window": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MaintenanceWindow"),
			},
			"checkly_multistep_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MultistepCheck"),
			},
			"checkly_private_location": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "PrivateLocation"),
			},
			"checkly_snippet": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Snippet"),
			},
			"checkly_status_page": {
				Tok:                tfbridge.MakeResource(mainPkg, mainMod, "StatusPage"),
				DeprecationMessage: "Use the StatusPageV3 resource instead. v3 status pages are separate objects, so an existing page cannot be converted.",
			},
			"checkly_status_page_service": {
				Tok:                tfbridge.MakeResource(mainPkg, mainMod, "StatusPageService"),
				DeprecationMessage: "Use StatusPageV3Component resources on a StatusPageV3 page instead.",
			},
			"checkly_status_page_v3": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "StatusPageV3"),
			},
			"checkly_status_page_v3_automation_rule": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "StatusPageV3AutomationRule"),
			},
			"checkly_status_page_v3_component": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "StatusPageV3Component"),
			},
			"checkly_tcp_check": {
				Tok:                tfbridge.MakeResource(mainPkg, mainMod, "TcpCheck"),
				DeprecationMessage: "Use the TcpMonitor resource instead. Both resources have the same properties.",
			},
			"checkly_tcp_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TcpMonitor"),
			},
			"checkly_url_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "UrlMonitor"),
			},
			"checkly_dns_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsMonitor"),
			},
			"checkly_icmp_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "IcmpMonitor"),
			},
			"checkly_playwright_check_suite": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "PlaywrightCheckSuite"),
			},
			"checkly_playwright_code_bundle": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "PlaywrightCodeBundle"),
			},
			"checkly_client_certificate": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "ClientCertificate"),
			},
			"checkly_grpc_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "GrpcMonitor"),
			},
			"checkly_ssl_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "SslMonitor"),
			},
			"checkly_traceroute_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TracerouteMonitor"),
			},
		},
		// If extra types are needed for configuration, they can be added here.
		ExtraTypes: map[string]schema.ComplexTypeSpec{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@checkly/pulumi",
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0", // Expand compatibility range.
			},
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Enable modern PyProject support in the generated Python SDK.
			PyProject: struct{ Enabled bool }{true},
		},
		Golang: &tfbridge.GolangInfo{
			// Set where the SDK is going to be published to.
			ImportBasePath: path.Join(
				"github.com/checkly/pulumi-checkly/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			// Opt in to all available code generation features.
			GenerateResourceContainerTypes: true,
			GenerateExtraInputTypes:        true,
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			// RespectSchemaVersion ensures the SDK is generated linking to the correct version of the provider.
			RespectSchemaVersion: true,
			// Use a wildcard import so NuGet will prefer the latest possible version.
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	// MustComputeTokens maps all resources and datasources from the upstream provider into Pulumi.
	//
	// tokens.SingleModule puts every upstream item into your provider's main module.
	//
	// You shouldn't need to override anything, but if you do, use the [tfbridge.ProviderInfo.Resources]
	// and [tfbridge.ProviderInfo.DataSources].
	prov.MustComputeTokens(tokens.SingleModule("checkly_", mainMod,
		tokens.MakeStandard(mainPkg)))

	// Rewrite state that v2.x stored in the shape of the Plugin SDK based provider.
	applyV2StateUpgradeHooks(&prov)

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
