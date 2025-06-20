// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a local Docker (V2) repository resource
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
//			_, err := artifactory.GetLocalDockerV2Repository(ctx, &artifactory.GetLocalDockerV2RepositoryArgs{
//				Key: "artifactory_local_test_docker_v2_repository",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLocalDockerV2Repository(ctx *pulumi.Context, args *GetLocalDockerV2RepositoryArgs, opts ...pulumi.InvokeOption) (*GetLocalDockerV2RepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocalDockerV2RepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalDockerV2Repository:getLocalDockerV2Repository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalDockerV2Repository.
type GetLocalDockerV2RepositoryArgs struct {
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2
	// schema 1 to this repository.
	BlockPushingSchema1 *bool   `pulumi:"blockPushingSchema1"`
	CdnRedirect         *bool   `pulumi:"cdnRedirect"`
	Description         *string `pulumi:"description"`
	DownloadDirect      *bool   `pulumi:"downloadDirect"`
	ExcludesPattern     *string `pulumi:"excludesPattern"`
	IncludesPattern     *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository.
	// Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
	// is no limit. This only applies to manifest v2.
	MaxUniqueTags       *int     `pulumi:"maxUniqueTags"`
	Notes               *string  `pulumi:"notes"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention *int  `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalDockerV2Repository.
type GetLocalDockerV2RepositoryResult struct {
	// "The Docker API version to use. This cannot be set"
	ApiVersion             string `pulumi:"apiVersion"`
	ArchiveBrowsingEnabled *bool  `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool  `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2
	// schema 1 to this repository.
	BlockPushingSchema1 bool    `pulumi:"blockPushingSchema1"`
	CdnRedirect         *bool   `pulumi:"cdnRedirect"`
	Description         *string `pulumi:"description"`
	DownloadDirect      *bool   `pulumi:"downloadDirect"`
	ExcludesPattern     *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository.
	// Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
	// is no limit. This only applies to manifest v2.
	MaxUniqueTags       *int     `pulumi:"maxUniqueTags"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention *int  `pulumi:"tagRetention"`
	XrayIndex    *bool `pulumi:"xrayIndex"`
}

func GetLocalDockerV2RepositoryOutput(ctx *pulumi.Context, args GetLocalDockerV2RepositoryOutputArgs, opts ...pulumi.InvokeOption) GetLocalDockerV2RepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLocalDockerV2RepositoryResultOutput, error) {
			args := v.(GetLocalDockerV2RepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getLocalDockerV2Repository:getLocalDockerV2Repository", args, GetLocalDockerV2RepositoryResultOutput{}, options).(GetLocalDockerV2RepositoryResultOutput), nil
		}).(GetLocalDockerV2RepositoryResultOutput)
}

// A collection of arguments for invoking getLocalDockerV2Repository.
type GetLocalDockerV2RepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest v2
	// schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput   `pulumi:"blockPushingSchema1"`
	CdnRedirect         pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	Description         pulumi.StringPtrInput `pulumi:"description"`
	DownloadDirect      pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern     pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern     pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringInput `pulumi:"key"`
	// The maximum number of unique tags of a single Docker image to store in this repository.
	// Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
	// is no limit. This only applies to manifest v2.
	MaxUniqueTags       pulumi.IntPtrInput      `pulumi:"maxUniqueTags"`
	Notes               pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef       pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention pulumi.IntPtrInput  `pulumi:"tagRetention"`
	XrayIndex    pulumi.BoolPtrInput `pulumi:"xrayIndex"`
}

func (GetLocalDockerV2RepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalDockerV2RepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalDockerV2Repository.
type GetLocalDockerV2RepositoryResultOutput struct{ *pulumi.OutputState }

func (GetLocalDockerV2RepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalDockerV2RepositoryResult)(nil)).Elem()
}

func (o GetLocalDockerV2RepositoryResultOutput) ToGetLocalDockerV2RepositoryResultOutput() GetLocalDockerV2RepositoryResultOutput {
	return o
}

func (o GetLocalDockerV2RepositoryResultOutput) ToGetLocalDockerV2RepositoryResultOutputWithContext(ctx context.Context) GetLocalDockerV2RepositoryResultOutput {
	return o
}

// "The Docker API version to use. This cannot be set"
func (o GetLocalDockerV2RepositoryResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, Artifactory will block the pushing of Docker images with manifest v2
// schema 1 to this repository.
func (o GetLocalDockerV2RepositoryResultOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) bool { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalDockerV2RepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique tags of a single Docker image to store in this repository.
// Once the number tags for an image exceeds this setting, older tags are removed. A value of 0 (default) indicates there
// is no limit. This only applies to manifest v2.
func (o GetLocalDockerV2RepositoryResultOutput) MaxUniqueTags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *int { return v.MaxUniqueTags }).(pulumi.IntPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// If greater than 1, overwritten tags will be saved by their digest, up to the set up
// number. This only applies to manifest V2.
func (o GetLocalDockerV2RepositoryResultOutput) TagRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *int { return v.TagRetention }).(pulumi.IntPtrOutput)
}

func (o GetLocalDockerV2RepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalDockerV2RepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalDockerV2RepositoryResultOutput{})
}
