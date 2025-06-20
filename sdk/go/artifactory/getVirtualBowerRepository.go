// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Bower repository.
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
//			_, err := artifactory.LookupVirtualBowerRepository(ctx, &artifactory.LookupVirtualBowerRepositoryArgs{
//				Key: "virtual-alpine",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualBowerRepository(ctx *pulumi.Context, args *LookupVirtualBowerRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualBowerRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualBowerRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualBowerRepository:getVirtualBowerRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualBowerRepository.
type LookupVirtualBowerRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// (Optional) When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo *string `pulumi:"externalDependenciesRemoteRepo"`
	IncludesPattern                *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
}

// A collection of values returned by getVirtualBowerRepository.
type LookupVirtualBowerRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	// (Optional) When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled *bool `pulumi:"externalDependenciesEnabled"`
	// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns []string `pulumi:"externalDependenciesPatterns"`
	// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo *string `pulumi:"externalDependenciesRemoteRepo"`
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

func LookupVirtualBowerRepositoryOutput(ctx *pulumi.Context, args LookupVirtualBowerRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualBowerRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVirtualBowerRepositoryResultOutput, error) {
			args := v.(LookupVirtualBowerRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getVirtualBowerRepository:getVirtualBowerRepository", args, LookupVirtualBowerRepositoryResultOutput{}, options).(LookupVirtualBowerRepositoryResultOutput), nil
		}).(LookupVirtualBowerRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualBowerRepository.
type LookupVirtualBowerRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	// (Optional) When set, external dependencies are rewritten. Default value is false.
	ExternalDependenciesEnabled pulumi.BoolPtrInput `pulumi:"externalDependenciesEnabled"`
	// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
	ExternalDependenciesPatterns pulumi.StringArrayInput `pulumi:"externalDependenciesPatterns"`
	// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
	ExternalDependenciesRemoteRepo pulumi.StringPtrInput `pulumi:"externalDependenciesRemoteRepo"`
	IncludesPattern                pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 pulumi.StringInput      `pulumi:"key"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories        pulumi.StringArrayInput `pulumi:"repositories"`
}

func (LookupVirtualBowerRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualBowerRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualBowerRepository.
type LookupVirtualBowerRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualBowerRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualBowerRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualBowerRepositoryResultOutput) ToLookupVirtualBowerRepositoryResultOutput() LookupVirtualBowerRepositoryResultOutput {
	return o
}

func (o LookupVirtualBowerRepositoryResultOutput) ToLookupVirtualBowerRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualBowerRepositoryResultOutput {
	return o
}

func (o LookupVirtualBowerRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// (Optional) When set, external dependencies are rewritten. Default value is false.
func (o LookupVirtualBowerRepositoryResultOutput) ExternalDependenciesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *bool { return v.ExternalDependenciesEnabled }).(pulumi.BoolPtrOutput)
}

// (Optional) An Allow List of Ant-style path expressions that specify where external dependencies may be downloaded from. By default, this is set to ** which means that dependencies may be downloaded from any external source.
func (o LookupVirtualBowerRepositoryResultOutput) ExternalDependenciesPatterns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) []string { return v.ExternalDependenciesPatterns }).(pulumi.StringArrayOutput)
}

// (Optional) The remote repository aggregated by this virtual repository in which the external dependency will be cached.
func (o LookupVirtualBowerRepositoryResultOutput) ExternalDependenciesRemoteRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.ExternalDependenciesRemoteRepo }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualBowerRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualBowerRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualBowerRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualBowerRepositoryResultOutput{})
}
