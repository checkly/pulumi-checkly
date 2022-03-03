// Copyright 2016-2018, Pulumi Corporation.
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
	"fmt"
	"os"
	"path/filepath"

	"github.com/checkly/pulumi-checkly/provider/pkg/version"
	"github.com/checkly/terraform-provider-checkly/checkly"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "checkly"
	// modules:
	mainMod = "index" // the checkly module
)

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	os.Setenv("CHECKLY_API_SOURCE", "PULUMI")
	p := shimv2.NewProvider(checkly.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                    p,
		Name:                 "checkly",
		Publisher:            "checkly",
		LogoURL:              "",
		Description:          "A Pulumi package for creating and managing checkly cloud resources.",
		PluginDownloadURL:    "https://github.com/checkly/pulumi-checkly/releases/${VERSION}",
		Keywords:             []string{"pulumi", "checkly", "category/monitoring"},
		License:              "Apache-2.0",
		Homepage:             "https://www.pulumi.com",
		Repository:           "https://github.com/checkly/pulumi-checkly",
		GitHubOrg:            "checkly",
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"checkly_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Check"),
			},
			"checkly_check_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "CheckGroup"),
			},
			"checkly_alert_channel": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "AlertChannel"),
			},
			"checkly_snippet": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "Snippet"),
			},
			"checkly_trigger_check": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TriggerCheck"),
			},
			"checkly_trigger_group": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "TriggerCheckGroup"),
			},
			"checkly_maintenance_windows": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "MaintenanceWindow"),
			},
			"checkly_dashboard": {
				Tok: tfbridge.MakeResource(mainPkg, mainMod, "PublicDashboard"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"header": {
						Default: &tfbridge.DefaultInfo{
							Value: "",
						},
					},
					"customUrl": {
						Default: &tfbridge.DefaultInfo{
							Value: "",
						},
					},
					"logo": {
						Default: &tfbridge.DefaultInfo{
							Value: "",
						},
					},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/checkly/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
