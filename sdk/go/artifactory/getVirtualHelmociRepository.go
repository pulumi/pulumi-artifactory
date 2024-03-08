// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a virtual Helm OCI repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v6/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupVirtualHelmociRepository(ctx, &artifactory.LookupVirtualHelmociRepositoryArgs{
//				Key: "my-helmoci-virtual",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVirtualHelmociRepository(ctx *pulumi.Context, args *LookupVirtualHelmociRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHelmociRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVirtualHelmociRepositoryResult
	err := ctx.Invoke("artifactory:index/getVirtualHelmociRepository:getVirtualHelmociRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVirtualHelmociRepository.
type LookupVirtualHelmociRepositoryArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
	IncludesPattern                               *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	Repositories        []string `pulumi:"repositories"`
	// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

// A collection of values returned by getVirtualHelmociRepository.
type LookupVirtualHelmociRepositoryResult struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         *string `pulumi:"defaultDeploymentRepo"`
	Description                                   *string `pulumi:"description"`
	ExcludesPattern                               *string `pulumi:"excludesPattern"`
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
	// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

func LookupVirtualHelmociRepositoryOutput(ctx *pulumi.Context, args LookupVirtualHelmociRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualHelmociRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualHelmociRepositoryResult, error) {
			args := v.(LookupVirtualHelmociRepositoryArgs)
			r, err := LookupVirtualHelmociRepository(ctx, &args, opts...)
			var s LookupVirtualHelmociRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualHelmociRepositoryResultOutput)
}

// A collection of arguments for invoking getVirtualHelmociRepository.
type LookupVirtualHelmociRepositoryOutputArgs struct {
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput   `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	DefaultDeploymentRepo                         pulumi.StringPtrInput `pulumi:"defaultDeploymentRepo"`
	Description                                   pulumi.StringPtrInput `pulumi:"description"`
	ExcludesPattern                               pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern                               pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                 pulumi.StringInput      `pulumi:"key"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	Repositories        pulumi.StringArrayInput `pulumi:"repositories"`
	// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrInput `pulumi:"resolveOciTagsByTimestamp"`
}

func (LookupVirtualHelmociRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHelmociRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getVirtualHelmociRepository.
type LookupVirtualHelmociRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualHelmociRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualHelmociRepositoryResult)(nil)).Elem()
}

func (o LookupVirtualHelmociRepositoryResultOutput) ToLookupVirtualHelmociRepositoryResultOutput() LookupVirtualHelmociRepositoryResultOutput {
	return o
}

func (o LookupVirtualHelmociRepositoryResultOutput) ToLookupVirtualHelmociRepositoryResultOutputWithContext(ctx context.Context) LookupVirtualHelmociRepositoryResultOutput {
	return o
}

func (o LookupVirtualHelmociRepositoryResultOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *bool {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVirtualHelmociRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualHelmociRepositoryResultOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) []string { return v.Repositories }).(pulumi.StringArrayOutput)
}

// (Optional) When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
func (o LookupVirtualHelmociRepositoryResultOutput) ResolveOciTagsByTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualHelmociRepositoryResult) *bool { return v.ResolveOciTagsByTimestamp }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualHelmociRepositoryResultOutput{})
}