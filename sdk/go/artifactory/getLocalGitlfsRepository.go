// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLocalGitlfsRepository(ctx *pulumi.Context, args *GetLocalGitlfsRepositoryArgs, opts ...pulumi.InvokeOption) (*GetLocalGitlfsRepositoryResult, error) {
	var rv GetLocalGitlfsRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalGitlfsRepository:getLocalGitlfsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalGitlfsRepository.
type GetLocalGitlfsRepositoryArgs struct {
	ArchiveBrowsingEnabled *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool    `pulumi:"blackedOut"`
	CdnRedirect            *bool    `pulumi:"cdnRedirect"`
	Description            *string  `pulumi:"description"`
	DownloadDirect         *bool    `pulumi:"downloadDirect"`
	ExcludesPattern        *string  `pulumi:"excludesPattern"`
	IncludesPattern        *string  `pulumi:"includesPattern"`
	Key                    string   `pulumi:"key"`
	Notes                  *string  `pulumi:"notes"`
	PriorityResolution     *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments    []string `pulumi:"projectEnvironments"`
	ProjectKey             *string  `pulumi:"projectKey"`
	PropertySets           []string `pulumi:"propertySets"`
	RepoLayoutRef          *string  `pulumi:"repoLayoutRef"`
	XrayIndex              *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalGitlfsRepository.
type GetLocalGitlfsRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	IncludesPattern     string   `pulumi:"includesPattern"`
	Key                 string   `pulumi:"key"`
	Notes               *string  `pulumi:"notes"`
	PackageType         string   `pulumi:"packageType"`
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	ProjectKey          *string  `pulumi:"projectKey"`
	PropertySets        []string `pulumi:"propertySets"`
	RepoLayoutRef       *string  `pulumi:"repoLayoutRef"`
	XrayIndex           *bool    `pulumi:"xrayIndex"`
}

func GetLocalGitlfsRepositoryOutput(ctx *pulumi.Context, args GetLocalGitlfsRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetLocalGitlfsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalGitlfsRepositoryResult, error) {
			args := v.(GetLocalGitlfsRepositoryArgs)
			r, err := GetLocalGitlfsRepository(ctx, &args, opts...)
			var s GetLocalGitlfsRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalGitlfsRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalGitlfsRepository.
type GetLocalGitlfsRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect            pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern        pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern        pulumi.StringPtrInput   `pulumi:"includesPattern"`
	Key                    pulumi.StringInput      `pulumi:"key"`
	Notes                  pulumi.StringPtrInput   `pulumi:"notes"`
	PriorityResolution     pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments    pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey             pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets           pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef          pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex              pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (GetLocalGitlfsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGitlfsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalGitlfsRepository.
type GetLocalGitlfsRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetLocalGitlfsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalGitlfsRepositoryResult)(nil)).Elem()
}

func (o GetLocalGitlfsRepositoryResultOutput) ToGetLocalGitlfsRepositoryResultOutput() GetLocalGitlfsRepositoryResultOutput {
	return o
}

func (o GetLocalGitlfsRepositoryResultOutput) ToGetLocalGitlfsRepositoryResultOutputWithContext(ctx context.Context) GetLocalGitlfsRepositoryResultOutput {
	return o
}

func (o GetLocalGitlfsRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalGitlfsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetLocalGitlfsRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalGitlfsRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalGitlfsRepositoryResultOutput{})
}