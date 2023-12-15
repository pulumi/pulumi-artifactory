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

package artifactory

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-artifactory/provider/v6/pkg/version"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"
	"path/filepath"
	"unicode"

	artifactoryProvider "github.com/jfrog/terraform-provider-artifactory/v10/pkg/artifactory/provider"
	pfbridge "github.com/pulumi/pulumi-terraform-bridge/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	mainPkg = "artifactory"
	// modules:
	mainMod = "index" // the artifactory module
)

func makeToken(mod, name string) string {
	lowerName := string(unicode.ToLower(rune(name[0]))) + name[1:]
	return fmt.Sprintf("%s:%s/%s:%s", mainPkg, mod, lowerName, name)
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, dataSource string) tokens.ModuleMember {
	return tokens.ModuleMember(makeToken(mod, dataSource))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	return tokens.Type(makeToken(mod, res))
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := pfbridge.MuxShimWithPF(context.Background(),
		shimv2.NewProvider(artifactoryProvider.SdkV2()),
		artifactoryProvider.Framework()(),
	)

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
				Tok: makeResource(mainMod, "AccessToken"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"access_token": {CSharpName: "Details"},
				},
			},
			"artifactory_api_key": {
				Tok: makeResource(mainMod, "ApiKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"api_key": {CSharpName: "Key"},
				},
			},
			"artifactory_certificate":                        {Tok: makeResource(mainMod, "Certificate")},
			"artifactory_general_security":                   {Tok: makeResource(mainMod, "GeneralSecurity")},
			"artifactory_group":                              {Tok: makeResource(mainMod, "Group")},
			"artifactory_keypair":                            {Tok: makeResource(mainMod, "Keypair")},
			"artifactory_local_alpine_repository":            {Tok: makeResource(mainMod, "AlpineRepository")},
			"artifactory_local_debian_repository":            {Tok: makeResource(mainMod, "DebianRepository")},
			"artifactory_local_docker_v1_repository":         {Tok: makeResource(mainMod, "DockerV1Repository")},
			"artifactory_local_docker_v2_repository":         {Tok: makeResource(mainMod, "DockerV2Repository")},
			"artifactory_local_bower_repository":             {Tok: makeResource(mainMod, "LocalBowerRepository")},
			"artifactory_local_chef_repository":              {Tok: makeResource(mainMod, "LocalChefRepository")},
			"artifactory_local_cocoapods_repository":         {Tok: makeResource(mainMod, "LocalCocoapodsRepository")},
			"artifactory_local_composer_repository":          {Tok: makeResource(mainMod, "LocalComposerRepository")},
			"artifactory_local_conan_repository":             {Tok: makeResource(mainMod, "LocalConanRepository")},
			"artifactory_local_cran_repository":              {Tok: makeResource(mainMod, "LocalCranRepository")},
			"artifactory_local_gems_repository":              {Tok: makeResource(mainMod, "LocalGemsRepository")},
			"artifactory_local_generic_repository":           {Tok: makeResource(mainMod, "LocalGenericRepository")},
			"artifactory_local_gitlfs_repository":            {Tok: makeResource(mainMod, "LocalGitltfsRepository")},
			"artifactory_local_go_repository":                {Tok: makeResource(mainMod, "LocalGoRepository")},
			"artifactory_local_gradle_repository":            {Tok: makeResource(mainMod, "LocalGradleRepository")},
			"artifactory_local_helm_repository":              {Tok: makeResource(mainMod, "LocalHelmRepository")},
			"artifactory_local_ivy_repository":               {Tok: makeResource(mainMod, "LocalIvyRepository")},
			"artifactory_local_maven_repository":             {Tok: makeResource(mainMod, "LocalMavenRepository")},
			"artifactory_local_npm_repository":               {Tok: makeResource(mainMod, "LocalNpmRepository")},
			"artifactory_local_nuget_repository":             {Tok: makeResource(mainMod, "LocalNugetRepository")},
			"artifactory_local_opkg_repository":              {Tok: makeResource(mainMod, "LocalOpkgRepository")},
			"artifactory_local_puppet_repository":            {Tok: makeResource(mainMod, "LocalPuppetRepository")},
			"artifactory_local_pypi_repository":              {Tok: makeResource(mainMod, "LocalPypiRepository")},
			"artifactory_local_rpm_repository":               {Tok: makeResource(mainMod, "LocalRpmRepository")},
			"artifactory_local_sbt_repository":               {Tok: makeResource(mainMod, "LocalSbtRepository")},
			"artifactory_local_pub_repository":               {Tok: makeResource(mainMod, "LocalPubRepository")},
			"artifactory_local_vagrant_repository":           {Tok: makeResource(mainMod, "LocalVagrantRepository")},
			"artifactory_oauth_settings":                     {Tok: makeResource(mainMod, "OauthSettings")},
			"artifactory_permission_target":                  {Tok: makeResource(mainMod, "PermissionTarget")},
			"artifactory_property_set":                       {Tok: makeResource(mainMod, "PropertySet")},
			"artifactory_proxy":                              {Tok: makeResource(mainMod, "Proxy")},
			"artifactory_pull_replication":                   {Tok: makeResource(mainMod, "PullReplication")},
			"artifactory_push_replication":                   {Tok: makeResource(mainMod, "PushReplication")},
			"artifactory_remote_cargo_repository":            {Tok: makeResource(mainMod, "RemoteCargoRepository")},
			"artifactory_remote_docker_repository":           {Tok: makeResource(mainMod, "RemoteDockerRepository")},
			"artifactory_remote_helm_repository":             {Tok: makeResource(mainMod, "RemoteHelmRepository")},
			"artifactory_remote_npm_repository":              {Tok: makeResource(mainMod, "RemoteNpmRepository")},
			"artifactory_remote_pub_repository":              {Tok: makeResource(mainMod, "RemotePubRepository")},
			"artifactory_replication_config":                 {Tok: makeResource(mainMod, "ReplicationConfig")},
			"artifactory_saml_settings":                      {Tok: makeResource(mainMod, "SamlSettings")},
			"artifactory_single_replication_config":          {Tok: makeResource(mainMod, "SingleReplicationConfig")},
			"artifactory_user":                               {Tok: makeResource(mainMod, "User")},
			"artifactory_virtual_go_repository":              {Tok: makeResource(mainMod, "GoRepository")},
			"artifactory_virtual_maven_repository":           {Tok: makeResource(mainMod, "MavenRepository")},
			"artifactory_federated_alpine_repository":        {Tok: makeResource(mainMod, "FederatedAlpineRepository")},
			"artifactory_federated_bower_repository":         {Tok: makeResource(mainMod, "FederatedBowerRepository")},
			"artifactory_federated_cargo_repository":         {Tok: makeResource(mainMod, "FederatedCargoRepository")},
			"artifactory_federated_chef_repository":          {Tok: makeResource(mainMod, "FederatedChefRepository")},
			"artifactory_federated_cocoapods_repository":     {Tok: makeResource(mainMod, "FederatedCocoapodsRepository")},
			"artifactory_federated_composer_repository":      {Tok: makeResource(mainMod, "FederatedComposerRepository")},
			"artifactory_federated_conan_repository":         {Tok: makeResource(mainMod, "FederatedConanRepository")},
			"artifactory_federated_conda_repository":         {Tok: makeResource(mainMod, "FederatedCondaRepository")},
			"artifactory_federated_cran_repository":          {Tok: makeResource(mainMod, "FederatedCranRepository")},
			"artifactory_federated_debian_repository":        {Tok: makeResource(mainMod, "FederatedDebianRepository")},
			"artifactory_federated_docker_repository":        {Tok: makeResource(mainMod, "FederatedDockerRepository")},
			"artifactory_federated_docker_v1_repository":     {Tok: makeResource(mainMod, "FederatedDockerV1Repository")},
			"artifactory_federated_docker_v2_repository":     {Tok: makeResource(mainMod, "FederatedDockerV2Repository")},
			"artifactory_federated_gems_repository":          {Tok: makeResource(mainMod, "FederatedGemsRepository")},
			"artifactory_federated_generic_repository":       {Tok: makeResource(mainMod, "FederatedGenericRepository")},
			"artifactory_federated_gitlfs_repository":        {Tok: makeResource(mainMod, "FederatedGitltfsRepository")},
			"artifactory_federated_go_repository":            {Tok: makeResource(mainMod, "FederatedGoRepository")},
			"artifactory_federated_gradle_repository":        {Tok: makeResource(mainMod, "FederatedGradleRepository")},
			"artifactory_federated_helm_repository":          {Tok: makeResource(mainMod, "FederatedHelmRepository")},
			"artifactory_federated_ivy_repository":           {Tok: makeResource(mainMod, "FederatedIvyRepository")},
			"artifactory_federated_maven_repository":         {Tok: makeResource(mainMod, "FederatedMavenRepository")},
			"artifactory_federated_npm_repository":           {Tok: makeResource(mainMod, "FederatedNpmRepository")},
			"artifactory_federated_nuget_repository":         {Tok: makeResource(mainMod, "FederatedNugetRepository")},
			"artifactory_federated_opkg_repository":          {Tok: makeResource(mainMod, "FederatedOpkgRepository")},
			"artifactory_federated_puppet_repository":        {Tok: makeResource(mainMod, "FederatedPuppetRepository")},
			"artifactory_federated_pypi_repository":          {Tok: makeResource(mainMod, "FederatedPypiRepository")},
			"artifactory_federated_rpm_repository":           {Tok: makeResource(mainMod, "FederatedRpmRepository")},
			"artifactory_federated_sbt_repository":           {Tok: makeResource(mainMod, "FederatedSbtRepository")},
			"artifactory_federated_swift_repository":         {Tok: makeResource(mainMod, "FederatedSwiftRepository")},
			"artifactory_federated_vagrant_repository":       {Tok: makeResource(mainMod, "FederatedVagrantRepository")},
			"artifactory_ldap_group_setting":                 {Tok: makeResource(mainMod, "LdapGroupSetting")},
			"artifactory_ldap_setting":                       {Tok: makeResource(mainMod, "LdapSetting")},
			"artifactory_virtual_conan_repository":           {Tok: makeResource(mainMod, "VirtualConanRepository")},
			"artifactory_virtual_generic_repository":         {Tok: makeResource(mainMod, "VirtualGenericRepository")},
			"artifactory_virtual_helm_repository":            {Tok: makeResource(mainMod, "VirtualHelmRepository")},
			"artifactory_virtual_rpm_repository":             {Tok: makeResource(mainMod, "VirtualRpmRepository")},
			"artifactory_remote_pypi_repository":             {Tok: makeResource(mainMod, "RemotePypiRepository")},
			"artifactory_artifact_property_webhook":          {Tok: makeResource(mainMod, "ArtifactPropertyWebhook")},
			"artifactory_artifact_webhook":                   {Tok: makeResource(mainMod, "ArtifactWebhook")},
			"artifactory_artifactory_release_bundle_webhook": {Tok: makeResource(mainMod, "ArtifactoryReleaseBundleWebhook")},
			"artifactory_build_webhook":                      {Tok: makeResource(mainMod, "BuildWebhook")},
			"artifactory_distribution_webhook":               {Tok: makeResource(mainMod, "DistributionWebhook")},
			"artifactory_docker_webhook":                     {Tok: makeResource(mainMod, "DockerWebhook")},
			"artifactory_release_bundle_webhook":             {Tok: makeResource(mainMod, "ReleaseBundleWebhook")},
			"artifactory_backup":                             {Tok: makeResource(mainMod, "Backup")},
			"artifactory_remote_gradle_repository":           {Tok: makeResource(mainMod, "RemoteGradleRepository")},
			"artifactory_remote_maven_repository":            {Tok: makeResource(mainMod, "RemoteMavenRepository")},
			"artifactory_remote_ivy_repository":              {Tok: makeResource(mainMod, "RemoteIvyRepository")},
			"artifactory_remote_sbt_repository":              {Tok: makeResource(mainMod, "RemoteSbtRepository")},
			"artifactory_remote_alpine_repository":           {Tok: makeResource(mainMod, "RemoteAlpineRepository")},
			"artifactory_remote_bower_repository":            {Tok: makeResource(mainMod, "RemoteBowerRepository")},
			"artifactory_remote_chef_repository":             {Tok: makeResource(mainMod, "RemoteChefRepository")},
			"artifactory_remote_cocoapods_repository":        {Tok: makeResource(mainMod, "RemoteCocoapodsRepository")},
			"artifactory_remote_composer_repository":         {Tok: makeResource(mainMod, "RemoteComposerRepository")},
			"artifactory_remote_conan_repository":            {Tok: makeResource(mainMod, "RemoteConanRepository")},
			"artifactory_remote_conda_repository":            {Tok: makeResource(mainMod, "RemoteCondaRepository")},
			"artifactory_remote_cran_repository":             {Tok: makeResource(mainMod, "RemoteCranRepository")},
			"artifactory_remote_debian_repository":           {Tok: makeResource(mainMod, "RemoteDebianRepository")},
			"artifactory_remote_gems_repository":             {Tok: makeResource(mainMod, "RemoteGemsRepository")},
			"artifactory_remote_generic_repository":          {Tok: makeResource(mainMod, "RemoteGenericRepository")},
			"artifactory_remote_gitlfs_repository":           {Tok: makeResource(mainMod, "RemoteGitlfsRepository")},
			"artifactory_remote_go_repository":               {Tok: makeResource(mainMod, "RemoteGoRepository")},
			"artifactory_remote_nuget_repository":            {Tok: makeResource(mainMod, "RemoteNugetRepository")},
			"artifactory_remote_opkg_repository":             {Tok: makeResource(mainMod, "RemoteOpkgRepository")},
			"artifactory_remote_p2_repository":               {Tok: makeResource(mainMod, "RemoteP2Repository")},
			"artifactory_remote_puppet_repository":           {Tok: makeResource(mainMod, "RemotePuppetRepository")},
			"artifactory_remote_rpm_repository":              {Tok: makeResource(mainMod, "RemoteRpmRepository")},
			"artifactory_remote_vcs_repository":              {Tok: makeResource(mainMod, "RemoteVcsRepository")},
			"artifactory_scoped_token":                       {Tok: makeResource(mainMod, "ScopedToken")},
			"artifactory_virtual_alpine_repository":          {Tok: makeResource(mainMod, "VirtualAlpineRepository")},
			"artifactory_virtual_bower_repository":           {Tok: makeResource(mainMod, "VirtualBowerRepository")},
			"artifactory_virtual_chef_repository":            {Tok: makeResource(mainMod, "VirtualChefRepository")},
			"artifactory_virtual_composer_repository":        {Tok: makeResource(mainMod, "VirtualComposerRepository")},
			"artifactory_virtual_conda_repository":           {Tok: makeResource(mainMod, "VirtualCondaRepository")},
			"artifactory_virtual_cran_repository":            {Tok: makeResource(mainMod, "VirtualCranRepository")},
			"artifactory_virtual_debian_repository":          {Tok: makeResource(mainMod, "VirtualDebianRepository")},
			"artifactory_virtual_docker_repository":          {Tok: makeResource(mainMod, "VirtualDockerRepository")},
			"artifactory_virtual_gems_repository":            {Tok: makeResource(mainMod, "VirtualGemsRepository")},
			"artifactory_virtual_gitlfs_repository":          {Tok: makeResource(mainMod, "VirtualGitlfsRepository")},
			"artifactory_virtual_gradle_repository":          {Tok: makeResource(mainMod, "VirtualGradleRepository")},
			"artifactory_virtual_ivy_repository":             {Tok: makeResource(mainMod, "VirtualIvyRepository")},
			"artifactory_virtual_npm_repository":             {Tok: makeResource(mainMod, "VirtualNpmRepository")},
			"artifactory_virtual_nuget_repository":           {Tok: makeResource(mainMod, "VirtualNugetRepository")},
			"artifactory_virtual_p2_repository":              {Tok: makeResource(mainMod, "VirtualP2Repository")},
			"artifactory_virtual_puppet_repository":          {Tok: makeResource(mainMod, "VirtualPuppetRepository")},
			"artifactory_virtual_pypi_repository":            {Tok: makeResource(mainMod, "VirtualPypiRepository")},
			"artifactory_virtual_sbt_repository":             {Tok: makeResource(mainMod, "VirtualSbtRepository")},
			"artifactory_local_cargo_repository":             {Tok: makeResource(mainMod, "LocalCargoRepository")},
			"artifactory_local_conda_repository":             {Tok: makeResource(mainMod, "LocalCondaRepository")},
			"artifactory_anonymous_user":                     {Tok: makeResource(mainMod, "AnonymousUser")},
			"artifactory_managed_user":                       {Tok: makeResource(mainMod, "ManagedUser")},
			"artifactory_unmanaged_user":                     {Tok: makeResource(mainMod, "UnmanagedUser")},
			"artifactory_virtual_pub_repository":             {Tok: makeResource(mainMod, "VirtualPubRepository")},
			"artifactory_federated_terraform_module_repository": {
				Tok: makeResource(mainMod, "FederatedTerraformModuleRepository"),
			},
			"artifactory_federated_terraform_provider_repository": {
				Tok: makeResource(mainMod, "FederatedTerraformProviderRepository"),
			},
			"artifactory_local_terraform_module_repository": {
				Tok: makeResource(mainMod, "LocalTerraformModuleRepository"),
			},
			"artifactory_local_terraform_provider_repository": {
				Tok: makeResource(mainMod, "LocalTerraformProviderRepository"),
			},
			"artifactory_local_terraformbackend_repository": {
				Tok: makeResource(mainMod, "LocalTerraformBackendRepository"),
			},
			"artifactory_remote_terraform_repository":  {Tok: makeResource(mainMod, "RemoteTerraformRepository")},
			"artifactory_virtual_terraform_repository": {Tok: makeResource(mainMod, "VirtualTerraformRepository")},
			"artifactory_local_swift_repository":       {Tok: makeResource(mainMod, "LocalSwiftRepository")},
			"artifactory_remote_swift_repository":      {Tok: makeResource(mainMod, "RemoteSwiftRepository")},
			"artifactory_virtual_swift_repository":     {Tok: makeResource(mainMod, "VirtualSwiftRepository")},
			"artifactory_repository_layout": {
				Tok: makeResource(mainMod, "RepositoryLayout"),
				Docs: &tfbridge.DocInfo{
					Markdown: []byte(" "),
				},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"artifactory_file":     {Tok: makeDataSource(mainMod, "getFile")},
			"artifactory_fileinfo": {Tok: makeDataSource(mainMod, "getFileinfo")},
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
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{

				Requires: map[string]string{
					"pulumi": ">=3.0.0,<4.0.0",
				}}
			i.PyProject.Enabled = true
			return i
		})(),

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
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
		MetadataInfo: tfbridge.NewProviderMetadata(bridgeMetadata),
	}

	prov.MustComputeTokens(tks.SingleModule("artifactory_", mainMod,
		tks.MakeStandard(mainPkg)))

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

//go:embed cmd/pulumi-resource-artifactory/bridge-metadata.json
var bridgeMetadata []byte
