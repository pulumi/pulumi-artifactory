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
	"regexp"
	"strconv"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	artifactoryProvider "github.com/jfrog/terraform-provider-artifactory/v11/pkg/artifactory/provider"

	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-artifactory/provider/v7/pkg/version"
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

func computeIDField(field resource.PropertyKey) tfbridge.ComputeID {
	return tfbridge.DelegateIDField(field, "artifactory", "https://github.com/pulumi/pulumi-artifactory")
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
		TFProviderModuleVersion: "v11",
		Version:                 version.Version,
		DocRules:                &tfbridge.DocRuleInfo{EditRules: docEditRules},
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

			"artifactory_backup":                  {ComputeID: computeIDField("key")},
			"artifactory_property_set":            {ComputeID: computeIDField("name")},
			"artifactory_proxy":                   {ComputeID: computeIDField("key")},
			"artifactory_distribution_public_key": {ComputeID: computeIDField("keyId")},
			"artifactory_mail_server": {ComputeID: func(ctx context.Context, state resource.PropertyMap) (resource.ID, error) {
				host, err := computeIDField("host")(ctx, state)
				if err != nil {
					return "", err
				}
				if state["port"].IsNumber() {
					return resource.ID(host.String() + ":" + strconv.FormatFloat(state["port"].NumberValue(), 'f', -1, 64)), nil
				}
				return host, nil
			}},
			"artifactory_repository_layout": {
				ComputeID: computeIDField("name"),
				Docs:      &tfbridge.DocInfo{AllowMissing: true},
			},
			"artifactory_certificate":                {ComputeID: computeIDField("alias")},
			"artifactory_keypair":                    {ComputeID: computeIDField("pairName")},
			"artifactory_package_cleanup_policy":     {ComputeID: computeIDField("key")},
			"artifactory_password_expiration_policy": {ComputeID: computeIDField("name")},
			"artifactory_user_lock_policy":           {ComputeID: computeIDField("name")},
			"artifactory_artifact": {ComputeID: func(ctx context.Context, state resource.PropertyMap) (resource.ID, error) {
				path, err := computeIDField("path")(ctx, state)
				if err != nil {
					return "", err
				}
				repository, err := computeIDField("repository")(ctx, state)
				if err != nil {
					return "", err
				}
				return resource.ID(path.String() + "@" + repository.String()), nil
			}},
			"artifactory_vault_configuration":         {ComputeID: computeIDField("id")},
			"artifactory_release_bundle_v2":           {ComputeID: computeIDField("name")},
			"artifactory_release_bundle_v2_promotion": {ComputeID: computeIDField("name")},
			"artifactory_item_properties":             {ComputeID: computeIDField("repoKey")},

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

func docEditRules(defaults []tfbridge.DocsEdit) []tfbridge.DocsEdit {
	return append(
		defaults,
		removeTFCloudOIDC,
	)
}

var TFCloudOIDCRegexp = regexp.MustCompile(`\* Terraform Cloud OIDC provider`)
var removeTFCloudOIDC = tfbridge.DocsEdit{
	Path: "index.md",
	Edit: func(_ string, content []byte) ([]byte, error) {
		content = TFCloudOIDCRegexp.ReplaceAllLiteral(content, nil)
		return content, nil
	},
}
