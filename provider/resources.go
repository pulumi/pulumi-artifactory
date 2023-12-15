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

//go:embed cmd/pulumi-resource-artifactory/bridge-metadata.json
var bridgeMetadata []byte

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
			"artifactory_certificate":                        {Tok: makeResource("Certificate")},
			"artifactory_general_security":                   {Tok: makeResource("GeneralSecurity")},
			"artifactory_group":                              {Tok: makeResource("Group")},
			"artifactory_keypair":                            {Tok: makeResource("Keypair")},
			"artifactory_local_alpine_repository":            {Tok: makeResource("AlpineRepository")},
			"artifactory_local_debian_repository":            {Tok: makeResource("DebianRepository")},
			"artifactory_local_docker_v1_repository":         {Tok: makeResource("DockerV1Repository")},
			"artifactory_local_docker_v2_repository":         {Tok: makeResource("DockerV2Repository")},
			"artifactory_local_bower_repository":             {Tok: makeResource("LocalBowerRepository")},
			"artifactory_local_chef_repository":              {Tok: makeResource("LocalChefRepository")},
			"artifactory_local_cocoapods_repository":         {Tok: makeResource("LocalCocoapodsRepository")},
			"artifactory_local_composer_repository":          {Tok: makeResource("LocalComposerRepository")},
			"artifactory_local_conan_repository":             {Tok: makeResource("LocalConanRepository")},
			"artifactory_local_cran_repository":              {Tok: makeResource("LocalCranRepository")},
			"artifactory_local_gems_repository":              {Tok: makeResource("LocalGemsRepository")},
			"artifactory_local_generic_repository":           {Tok: makeResource("LocalGenericRepository")},
			"artifactory_local_gitlfs_repository":            {Tok: makeResource("LocalGitltfsRepository")},
			"artifactory_local_go_repository":                {Tok: makeResource("LocalGoRepository")},
			"artifactory_local_gradle_repository":            {Tok: makeResource("LocalGradleRepository")},
			"artifactory_local_helm_repository":              {Tok: makeResource("LocalHelmRepository")},
			"artifactory_local_ivy_repository":               {Tok: makeResource("LocalIvyRepository")},
			"artifactory_local_maven_repository":             {Tok: makeResource("LocalMavenRepository")},
			"artifactory_local_npm_repository":               {Tok: makeResource("LocalNpmRepository")},
			"artifactory_local_nuget_repository":             {Tok: makeResource("LocalNugetRepository")},
			"artifactory_local_opkg_repository":              {Tok: makeResource("LocalOpkgRepository")},
			"artifactory_local_puppet_repository":            {Tok: makeResource("LocalPuppetRepository")},
			"artifactory_local_pypi_repository":              {Tok: makeResource("LocalPypiRepository")},
			"artifactory_local_rpm_repository":               {Tok: makeResource("LocalRpmRepository")},
			"artifactory_local_sbt_repository":               {Tok: makeResource("LocalSbtRepository")},
			"artifactory_local_pub_repository":               {Tok: makeResource("LocalPubRepository")},
			"artifactory_local_vagrant_repository":           {Tok: makeResource("LocalVagrantRepository")},
			"artifactory_oauth_settings":                     {Tok: makeResource("OauthSettings")},
			"artifactory_permission_target":                  {Tok: makeResource("PermissionTarget")},
			"artifactory_property_set":                       {Tok: makeResource("PropertySet")},
			"artifactory_proxy":                              {Tok: makeResource("Proxy")},
			"artifactory_pull_replication":                   {Tok: makeResource("PullReplication")},
			"artifactory_push_replication":                   {Tok: makeResource("PushReplication")},
			"artifactory_remote_cargo_repository":            {Tok: makeResource("RemoteCargoRepository")},
			"artifactory_remote_docker_repository":           {Tok: makeResource("RemoteDockerRepository")},
			"artifactory_remote_helm_repository":             {Tok: makeResource("RemoteHelmRepository")},
			"artifactory_remote_npm_repository":              {Tok: makeResource("RemoteNpmRepository")},
			"artifactory_remote_pub_repository":              {Tok: makeResource("RemotePubRepository")},
			"artifactory_replication_config":                 {Tok: makeResource("ReplicationConfig")},
			"artifactory_saml_settings":                      {Tok: makeResource("SamlSettings")},
			"artifactory_single_replication_config":          {Tok: makeResource("SingleReplicationConfig")},
			"artifactory_user":                               {Tok: makeResource("User")},
			"artifactory_virtual_go_repository":              {Tok: makeResource("GoRepository")},
			"artifactory_virtual_maven_repository":           {Tok: makeResource("MavenRepository")},
			"artifactory_federated_alpine_repository":        {Tok: makeResource("FederatedAlpineRepository")},
			"artifactory_federated_bower_repository":         {Tok: makeResource("FederatedBowerRepository")},
			"artifactory_federated_cargo_repository":         {Tok: makeResource("FederatedCargoRepository")},
			"artifactory_federated_chef_repository":          {Tok: makeResource("FederatedChefRepository")},
			"artifactory_federated_cocoapods_repository":     {Tok: makeResource("FederatedCocoapodsRepository")},
			"artifactory_federated_composer_repository":      {Tok: makeResource("FederatedComposerRepository")},
			"artifactory_federated_conan_repository":         {Tok: makeResource("FederatedConanRepository")},
			"artifactory_federated_conda_repository":         {Tok: makeResource("FederatedCondaRepository")},
			"artifactory_federated_cran_repository":          {Tok: makeResource("FederatedCranRepository")},
			"artifactory_federated_debian_repository":        {Tok: makeResource("FederatedDebianRepository")},
			"artifactory_federated_docker_repository":        {Tok: makeResource("FederatedDockerRepository")},
			"artifactory_federated_docker_v1_repository":     {Tok: makeResource("FederatedDockerV1Repository")},
			"artifactory_federated_docker_v2_repository":     {Tok: makeResource("FederatedDockerV2Repository")},
			"artifactory_federated_gems_repository":          {Tok: makeResource("FederatedGemsRepository")},
			"artifactory_federated_generic_repository":       {Tok: makeResource("FederatedGenericRepository")},
			"artifactory_federated_gitlfs_repository":        {Tok: makeResource("FederatedGitltfsRepository")},
			"artifactory_federated_go_repository":            {Tok: makeResource("FederatedGoRepository")},
			"artifactory_federated_gradle_repository":        {Tok: makeResource("FederatedGradleRepository")},
			"artifactory_federated_helm_repository":          {Tok: makeResource("FederatedHelmRepository")},
			"artifactory_federated_ivy_repository":           {Tok: makeResource("FederatedIvyRepository")},
			"artifactory_federated_maven_repository":         {Tok: makeResource("FederatedMavenRepository")},
			"artifactory_federated_npm_repository":           {Tok: makeResource("FederatedNpmRepository")},
			"artifactory_federated_nuget_repository":         {Tok: makeResource("FederatedNugetRepository")},
			"artifactory_federated_opkg_repository":          {Tok: makeResource("FederatedOpkgRepository")},
			"artifactory_federated_puppet_repository":        {Tok: makeResource("FederatedPuppetRepository")},
			"artifactory_federated_pypi_repository":          {Tok: makeResource("FederatedPypiRepository")},
			"artifactory_federated_rpm_repository":           {Tok: makeResource("FederatedRpmRepository")},
			"artifactory_federated_sbt_repository":           {Tok: makeResource("FederatedSbtRepository")},
			"artifactory_federated_swift_repository":         {Tok: makeResource("FederatedSwiftRepository")},
			"artifactory_federated_vagrant_repository":       {Tok: makeResource("FederatedVagrantRepository")},
			"artifactory_ldap_group_setting":                 {Tok: makeResource("LdapGroupSetting")},
			"artifactory_ldap_setting":                       {Tok: makeResource("LdapSetting")},
			"artifactory_virtual_conan_repository":           {Tok: makeResource("VirtualConanRepository")},
			"artifactory_virtual_generic_repository":         {Tok: makeResource("VirtualGenericRepository")},
			"artifactory_virtual_helm_repository":            {Tok: makeResource("VirtualHelmRepository")},
			"artifactory_virtual_rpm_repository":             {Tok: makeResource("VirtualRpmRepository")},
			"artifactory_remote_pypi_repository":             {Tok: makeResource("RemotePypiRepository")},
			"artifactory_artifact_property_webhook":          {Tok: makeResource("ArtifactPropertyWebhook")},
			"artifactory_artifact_webhook":                   {Tok: makeResource("ArtifactWebhook")},
			"artifactory_artifactory_release_bundle_webhook": {Tok: makeResource("ArtifactoryReleaseBundleWebhook")},
			"artifactory_build_webhook":                      {Tok: makeResource("BuildWebhook")},
			"artifactory_distribution_webhook":               {Tok: makeResource("DistributionWebhook")},
			"artifactory_docker_webhook":                     {Tok: makeResource("DockerWebhook")},
			"artifactory_release_bundle_webhook":             {Tok: makeResource("ReleaseBundleWebhook")},
			"artifactory_backup":                             {Tok: makeResource("Backup")},
			"artifactory_remote_gradle_repository":           {Tok: makeResource("RemoteGradleRepository")},
			"artifactory_remote_maven_repository":            {Tok: makeResource("RemoteMavenRepository")},
			"artifactory_remote_ivy_repository":              {Tok: makeResource("RemoteIvyRepository")},
			"artifactory_remote_sbt_repository":              {Tok: makeResource("RemoteSbtRepository")},
			"artifactory_remote_alpine_repository":           {Tok: makeResource("RemoteAlpineRepository")},
			"artifactory_remote_bower_repository":            {Tok: makeResource("RemoteBowerRepository")},
			"artifactory_remote_chef_repository":             {Tok: makeResource("RemoteChefRepository")},
			"artifactory_remote_cocoapods_repository":        {Tok: makeResource("RemoteCocoapodsRepository")},
			"artifactory_remote_composer_repository":         {Tok: makeResource("RemoteComposerRepository")},
			"artifactory_remote_conan_repository":            {Tok: makeResource("RemoteConanRepository")},
			"artifactory_remote_conda_repository":            {Tok: makeResource("RemoteCondaRepository")},
			"artifactory_remote_cran_repository":             {Tok: makeResource("RemoteCranRepository")},
			"artifactory_remote_debian_repository":           {Tok: makeResource("RemoteDebianRepository")},
			"artifactory_remote_gems_repository":             {Tok: makeResource("RemoteGemsRepository")},
			"artifactory_remote_generic_repository":          {Tok: makeResource("RemoteGenericRepository")},
			"artifactory_remote_gitlfs_repository":           {Tok: makeResource("RemoteGitlfsRepository")},
			"artifactory_remote_go_repository":               {Tok: makeResource("RemoteGoRepository")},
			"artifactory_remote_nuget_repository":            {Tok: makeResource("RemoteNugetRepository")},
			"artifactory_remote_opkg_repository":             {Tok: makeResource("RemoteOpkgRepository")},
			"artifactory_remote_p2_repository":               {Tok: makeResource("RemoteP2Repository")},
			"artifactory_remote_puppet_repository":           {Tok: makeResource("RemotePuppetRepository")},
			"artifactory_remote_rpm_repository":              {Tok: makeResource("RemoteRpmRepository")},
			"artifactory_remote_vcs_repository":              {Tok: makeResource("RemoteVcsRepository")},
			"artifactory_scoped_token":                       {Tok: makeResource("ScopedToken")},
			"artifactory_virtual_alpine_repository":          {Tok: makeResource("VirtualAlpineRepository")},
			"artifactory_virtual_bower_repository":           {Tok: makeResource("VirtualBowerRepository")},
			"artifactory_virtual_chef_repository":            {Tok: makeResource("VirtualChefRepository")},
			"artifactory_virtual_composer_repository":        {Tok: makeResource("VirtualComposerRepository")},
			"artifactory_virtual_conda_repository":           {Tok: makeResource("VirtualCondaRepository")},
			"artifactory_virtual_cran_repository":            {Tok: makeResource("VirtualCranRepository")},
			"artifactory_virtual_debian_repository":          {Tok: makeResource("VirtualDebianRepository")},
			"artifactory_virtual_docker_repository":          {Tok: makeResource("VirtualDockerRepository")},
			"artifactory_virtual_gems_repository":            {Tok: makeResource("VirtualGemsRepository")},
			"artifactory_virtual_gitlfs_repository":          {Tok: makeResource("VirtualGitlfsRepository")},
			"artifactory_virtual_gradle_repository":          {Tok: makeResource("VirtualGradleRepository")},
			"artifactory_virtual_ivy_repository":             {Tok: makeResource("VirtualIvyRepository")},
			"artifactory_virtual_npm_repository":             {Tok: makeResource("VirtualNpmRepository")},
			"artifactory_virtual_nuget_repository":           {Tok: makeResource("VirtualNugetRepository")},
			"artifactory_virtual_p2_repository":              {Tok: makeResource("VirtualP2Repository")},
			"artifactory_virtual_puppet_repository":          {Tok: makeResource("VirtualPuppetRepository")},
			"artifactory_virtual_pypi_repository":            {Tok: makeResource("VirtualPypiRepository")},
			"artifactory_virtual_sbt_repository":             {Tok: makeResource("VirtualSbtRepository")},
			"artifactory_local_cargo_repository":             {Tok: makeResource("LocalCargoRepository")},
			"artifactory_local_conda_repository":             {Tok: makeResource("LocalCondaRepository")},
			"artifactory_anonymous_user":                     {Tok: makeResource("AnonymousUser")},
			"artifactory_managed_user":                       {Tok: makeResource("ManagedUser")},
			"artifactory_unmanaged_user":                     {Tok: makeResource("UnmanagedUser")},
			"artifactory_virtual_pub_repository":             {Tok: makeResource("VirtualPubRepository")},
			"artifactory_federated_terraform_module_repository": {
				Tok: makeResource("FederatedTerraformModuleRepository"),
			},
			"artifactory_federated_terraform_provider_repository": {
				Tok: makeResource("FederatedTerraformProviderRepository"),
			},
			"artifactory_local_terraform_module_repository": {
				Tok: makeResource("LocalTerraformModuleRepository"),
			},
			"artifactory_local_terraform_provider_repository": {
				Tok: makeResource("LocalTerraformProviderRepository"),
			},
			"artifactory_local_terraformbackend_repository": {
				Tok: makeResource("LocalTerraformBackendRepository"),
			},
			"artifactory_remote_terraform_repository":  {Tok: makeResource("RemoteTerraformRepository")},
			"artifactory_virtual_terraform_repository": {Tok: makeResource("VirtualTerraformRepository")},
			"artifactory_local_swift_repository":       {Tok: makeResource("LocalSwiftRepository")},
			"artifactory_remote_swift_repository":      {Tok: makeResource("RemoteSwiftRepository")},
			"artifactory_virtual_swift_repository":     {Tok: makeResource("VirtualSwiftRepository")},
			"artifactory_repository_layout": {
				Tok:  makeResource("RepositoryLayout"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
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
