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

package artifactory

import (
	"context"
	"fmt"
	"path/filepath"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	artifactoryProvider "github.com/jfrog/terraform-provider-artifactory/v10/pkg/artifactory/provider"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-artifactory/provider/v6/pkg/version"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "artifactory"
	// modules:
	mainMod = "index" // the artifactory module
)

//go:embed cmd/pulumi-resource-artifactory/bridge-metadata.json
var bridgeMetadata []byte

type computeIDFunc = func(ctx context.Context, state resource.PropertyMap) (resource.ID, error)

func computeIDField(field resource.PropertyKey) computeIDFunc {
	return func(ctx context.Context, state resource.PropertyMap) (resource.ID, error) {
		fieldValue, ok := state[field]
		if !ok {
			return "", fmt.Errorf("id: could not find required property '%s'", field)
		}

		// ComputeID is only called during when preview=false, so we don't need to
		// deal with computed properties.

		if fieldValue.IsSecret() || (fieldValue.IsOutput() && fieldValue.OutputValue().Secret) {
			msg := fmt.Sprintf("Setting non-secret resource ID as '%s' (which is secret)", field)
			tfbridge.GetLogger(ctx).Warn(msg)
			if fieldValue.IsSecret() {
				fieldValue = fieldValue.SecretValue().Element
			} else {
				fieldValue = fieldValue.OutputValue().Element
			}
		}

		if !fieldValue.IsString() {
			return "", fmt.Errorf("expected '%s' to be of type string, found %s",
				field, fieldValue.TypeString())
		}

		return resource.ID(fieldValue.StringValue()), nil
	}
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pfbridge.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(artifactoryProvider.SdkV2()),
		artifactoryProvider.Framework()(),
	)

	makeResource := func(res string) tokens.Type { return tfbridge.MakeResource(mainPkg, mainMod, res) }

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "artifactory",
		Description:             "A Pulumi package for creating and managing artifactory cloud resources.",
		Keywords:                []string{"pulumi", "artifactory"},
		License:                 "Apache-2.0",
		Homepage:                "https://pulumi.io",
		Repository:              "https://github.com/pulumi/pulumi-artifactory",
		GitHubOrg:               "jfrog",
		TFProviderModuleVersion: "v10",
		Version:                 version.Version,
		Config: map[string]*tfbridge.SchemaInfo{
			"check_license": {
				Default: &tfbridge.DefaultInfo{
					Value: false,
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"artifactory_access_token": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"access_token": {CSharpName: "Details"},
				},
			},
			"artifactory_api_key": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"api_key": {CSharpName: "Key"},
				},
			},

			"artifactory_backup": {ComputeID: computeIDField("key")},

			// Old Manual Mappings.
			//
			// We leave these as they are to avoid churn.
			"artifactory_federated_gitlfs_repository":       {Tok: makeResource("FederatedGitltfsRepository")},
			"artifactory_local_gitlfs_repository":           {Tok: makeResource("LocalGitltfsRepository")},
			"artifactory_local_alpine_repository":           {Tok: makeResource("AlpineRepository")},
			"artifactory_local_debian_repository":           {Tok: makeResource("DebianRepository")},
			"artifactory_local_docker_v1_repository":        {Tok: makeResource("DockerV1Repository")},
			"artifactory_local_docker_v2_repository":        {Tok: makeResource("DockerV2Repository")},
			"artifactory_local_bower_repository":            {Tok: makeResource("LocalBowerRepository")},
			"artifactory_virtual_go_repository":             {Tok: makeResource("GoRepository")},
			"artifactory_virtual_maven_repository":          {Tok: makeResource("MavenRepository")},
			"artifactory_local_terraformbackend_repository": {Tok: makeResource("LocalTerraformBackendRepository")},

			"artifactory_repository_layout": {Docs: &tfbridge.DocInfo{AllowMissing: true}},
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
			RespectSchemaVersion: true,
		},
		Python: &tfbridge.PythonInfo{
			RespectSchemaVersion: true,
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
	}

	prov.MustComputeTokens(tks.SingleModule("artifactory_", mainMod, tks.MakeStandard(mainPkg)))

	// The provider doesn't have docs for the following data sources, so we exclude them.
	docLessDataSources := []string{
		"artifactory_federated_docker_repository",
		"artifactory_local_alpine_repository",
		"artifactory_local_bower_repository",
		"artifactory_local_cargo_repository",
		"artifactory_local_chef_repository",
		"artifactory_local_cocoapods_repository",
		"artifactory_local_composer_repository",
		"artifactory_local_conan_repository",
		"artifactory_local_conda_repository",
		"artifactory_local_cran_repository",
		"artifactory_local_debian_repository",
		"artifactory_local_docker_v1_repository",
		"artifactory_local_docker_v2_repository",
		"artifactory_local_gems_repository",
		"artifactory_local_generic_repository",
		"artifactory_local_gitlfs_repository",
		"artifactory_local_go_repository",
		"artifactory_local_gradle_repository",
		"artifactory_local_helm_repository",
		"artifactory_local_huggingfaceml_repository",
		"artifactory_local_ivy_repository",
		"artifactory_local_maven_repository",
		"artifactory_local_npm_repository",
		"artifactory_local_nuget_repository",
		"artifactory_local_opkg_repository",
		"artifactory_local_pub_repository",
		"artifactory_local_puppet_repository",
		"artifactory_local_pypi_repository",
		"artifactory_local_rpm_repository",
		"artifactory_local_sbt_repository",
		"artifactory_local_swift_repository",
		"artifactory_local_terraform_module_repository",
		"artifactory_local_terraform_provider_repository",
		"artifactory_local_terraformbackend_repository",
		"artifactory_local_vagrant_repository",
	}

	missingDocInfo := &tfbridge.DocInfo{AllowMissing: true}
	for _, s := range docLessDataSources {
		prov.DataSources[s].Docs = missingDocInfo
	}

	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
