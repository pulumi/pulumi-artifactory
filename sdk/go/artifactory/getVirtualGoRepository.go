// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Go repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.GetVirtualGoRepository(ctx, &artifactory.GetVirtualGoRepositoryArgs{
//				Key: "virtual-go",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetVirtualGoRepository(ctx *pulumi.Context, args *GetVirtualGoRepositoryArgs, opts ...pulumi.InvokeOption) (*GetVirtualGoRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetVirtualGoRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualGoRepository:getVirtualGoRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualGoRepository.
type GetVirtualGoRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// (Optional) 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	IncludesPattern              *string  `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualGoRepository.
type GetVirtualGoRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// (Optional) 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	IncludesPattern     *string  `pulumi:"includesPattern"`
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
}

func GetVirtualGoRepositoryOutput(ctx *pulumi.Context, args GetVirtualGoRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetVirtualGoRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetVirtualGoRepositoryResultOutput, error) {
			args := v.(GetVirtualGoRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getVirtualGoRepository:getVirtualGoRepository", args, GetVirtualGoRepositoryResultOutput{}, options).(GetVirtualGoRepositoryResultOutput), nil
		}).(GetVirtualGoRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualGoRepository.
type GetVirtualGoRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
	// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
	ExternalDependenciesEnabled pulumi.BoolPtrInput `pulumi:"externalDependenciesEnabled"`
	// (Optional) 'go-import' Allow List on the UI.
	ExternalDependenciesPatterns pulumi.StringArrayInput `pulumi:"externalDependenciesPatterns"`
	IncludesPattern              pulumi.StringPtrInput   `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 pulumi.StringInput      `pulumi:"key"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories        pulumi.StringArrayInput `pulumi:"repositories"`
}

func (GetVirtualGoRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualGoRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualGoRepository.
type GetVirtualGoRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetVirtualGoRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualGoRepositoryResult)(nil)).Elem()
}

func (o GetVirtualGoRepositoryResultOutput) ToGetVirtualGoRepositoryResultOutput() GetVirtualGoRepositoryResultOutput {
	return o
}

func (o GetVirtualGoRepositoryResultOutput) ToGetVirtualGoRepositoryResultOutputWithContext(ctx context.Context) GetVirtualGoRepositoryResultOutput {
	return o
}

func (o GetVirtualGoRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *bool { return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts }).(pulumi.BoolPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional) Shorthand for "Enable 'go-import' Meta Tags" on the UI. This must be set to true in order to use the allow list.
// When checked (default), Artifactory will automatically follow remote VCS roots in 'go-import' meta tags to download remote modules.
func (o GetVirtualGoRepositoryResultOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *bool { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

// (Optional) 'go-import' Allow List on the UI.
func (o GetVirtualGoRepositoryResultOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) []string { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetVirtualGoRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetVirtualGoRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetVirtualGoRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetVirtualGoRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetVirtualGoRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualGoRepositoryResultOutput{})
}
