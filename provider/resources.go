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
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	checkly "github.com/checkly/terraform-provider-checkly/checkly" // Import the upstream provider

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/info"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"

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
		// Instantiate the Terraform provider
		//
		// The [pulumi-terraform-bridge](https://github.com/pulumi/pulumi-terraform-bridge) supports 3
		// types of Terraform providers:
		//
		// 1. Providers written with the terraform-plugin-sdk/v1:
		//
		//    If the provider you are bridging is written with the terraform-plugin-sdk/v1, then you
		//    will need to adapt the boilerplate:
		//
		//    - Change the import "shimv2" to "shimv1" and change the associated import to
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1".
		//
		//    You can then proceed as normal.
		//
		// 2. Providers written with terraform-plugin-sdk/v2:
		//
		//    This boilerplate is already geared towards providers written with the
		//    terraform-plugin-sdk/v2, since it is the most common provider framework used. No
		//    adaptions are needed.
		//
		// 3. Providers written with terraform-plugin-framework:
		//
		//    If the provider you are bridging is written with the terraform-plugin-framework, then
		//    you will need to adapt the boilerplate:
		//
		//    - Remove the `shimv2` import and add:
		//
		//      	pfbridge "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
		//
		//    - Replace `shimv2.NewProvider` with `pfbridge.ShimProvider`.
		//
		//    - In provider/cmd/pulumi-tfgen-checkly/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfgen". Remove the `version.Version`
		//      argument to `tfgen.Main`.
		//
		//    - In provider/cmd/pulumi-resource-checkly/main.go, replace the
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge" import with
		//      "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge". Replace the arguments to the
		//      `tfbridge.Main` so it looks like this:
		//
		//      	tfbridge.Main(context.Background(), "checkly", checkly.Provider(),
		//			tfbridge.ProviderMetadata{PulumiSchema: pulumiSchema})
		//
		//   Detailed instructions can be found at
		//   https://pulumi-developer-docs.readthedocs.io/projects/pulumi-terraform-bridge/en/latest/docs/guides/new-pf-provider.html
		//   After that, you can proceed as normal.
		//
		// This is where you give the bridge a handle to the upstream terraform provider. SDKv2
		// convention is to have a function at "github.com/checkly/terraform-provider-checkly/provider".New
		// which takes a version and produces a factory function. The provider you are bridging may
		// not do that. You will need to find the function (generally called in upstream's main.go)
		// that produces a:
		//
		// - *"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema".Provider (for SDKv2)
		// - *"github.com/hashicorp/terraform-plugin-sdk/v1/helper/schema".Provider (for SDKv1)
		// - "github.com/hashicorp/terraform-plugin-framework/provider".Provider (for plugin-framework)
		//
		//nolint:lll
		P: shimv2.NewProvider(checkly.Provider()),

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
		GitHubOrg:    "checkly",
		MetadataInfo: tfbridge.NewProviderMetadata(metadata),
		Config:       map[string]*tfbridge.SchemaInfo{},
		Resources: map[string]*info.Resource{
			// Automatic token maping will ignore resources that are manually mapped.
			"checkly_alert_channel": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "AlertChannel"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"webhook": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"headers": {
									Transform: handleFaultySliceMapDefault,
								},
								"query_parameters": {
									Transform: handleFaultySliceMapDefault,
								},
							},
						},
					},
				},
			},
			"checkly_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Check"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"request": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"headers": {
									Transform: handleFaultySliceMapDefault,
								},
								"query_parameters": {
									Transform: handleFaultySliceMapDefault,
								},
							},
						},
					},
				},
			},
			"checkly_check_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CheckGroup"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"api_check_defaults": {
						Elem: &tfbridge.SchemaInfo{
							Fields: map[string]*tfbridge.SchemaInfo{
								"url": {
									Default: &tfbridge.DefaultInfo{
										Value: "",
									},
								},
								"headers": {
									Transform: handleFaultySliceMapDefault,
								},
								"query_parameters": {
									Transform: handleFaultySliceMapDefault,
								},
							},
						},
					},
				},
			},
			"checkly_dashboard": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Dashboard"),
			},
			"checkly_environment_variable": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "EnvironmentVariable"),
			},
			"checkly_heartbeat": {
				Tok:        tfbridge.MakeResource(mainPkg, mainMod, "HeartbeatCheck"),
				CSharpName: "HeartbeatCheck",
			},
			"checkly_heartbeat_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "HeartbeatMonitor"),
			},
			"checkly_maintenance_windows": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MaintenanceWindow"),
			},
			"checkly_private_location": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "PrivateLocation"),
			},
			"checkly_snippet": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Snippet"),
			},
			"checkly_status_page": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "StatusPage"),
			},
			"checkly_status_page_service": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "StatusPageService"),
			},
			"checkly_tcp_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TcpCheck"),
			},
			"checkly_tcp_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TcpMonitor"),
			},
			"checkly_trigger_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TriggerCheck"),
			},
			"checkly_trigger_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TriggerCheckGroup"),
			},
			"checkly_url_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "UrlMonitor"),
			},
			"checkly_dns_monitor": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "DnsMonitor"),
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

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}

// Provider 1.8.2 uses DefaultFunc to create a faulty default value of
// `[]tfMap{}`, which Pulumi cannot deal with, so we have have to remap the
// faulty value to an empty object which it should have been.
func handleFaultySliceMapDefault(value resource.PropertyValue) (resource.PropertyValue, error) {
	// If we find a slice, we know we have the faulty default value.
	if value.IsArray() {
		return resource.NewObjectProperty(resource.PropertyMap{}), nil
	}

	return value, nil
}
