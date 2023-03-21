// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetLocalAlpineRepository(ctx *pulumi.Context, args *GetLocalAlpineRepositoryArgs, opts ...pulumi.InvokeOption) (*GetLocalAlpineRepositoryResult, error) {
	var rv GetLocalAlpineRepositoryResult
	err := ctx.Invoke("artifactory:index/getLocalAlpineRepository:getLocalAlpineRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocalAlpineRepository.
type GetLocalAlpineRepositoryArgs struct {
	ArchiveBrowsingEnabled  *bool    `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              *bool    `pulumi:"blackedOut"`
	CdnRedirect             *bool    `pulumi:"cdnRedirect"`
	Description             *string  `pulumi:"description"`
	DownloadDirect          *bool    `pulumi:"downloadDirect"`
	ExcludesPattern         *string  `pulumi:"excludesPattern"`
	IncludesPattern         *string  `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	Key                     string   `pulumi:"key"`
	Notes                   *string  `pulumi:"notes"`
	PrimaryKeypairRef       *string  `pulumi:"primaryKeypairRef"`
	PriorityResolution      *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments     []string `pulumi:"projectEnvironments"`
	ProjectKey              *string  `pulumi:"projectKey"`
	PropertySets            []string `pulumi:"propertySets"`
	RepoLayoutRef           *string  `pulumi:"repoLayoutRef"`
	XrayIndex               *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getLocalAlpineRepository.
type GetLocalAlpineRepositoryResult struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	Description            *string `pulumi:"description"`
	DownloadDirect         *bool   `pulumi:"downloadDirect"`
	ExcludesPattern        string  `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string   `pulumi:"id"`
	IncludesPattern         string   `pulumi:"includesPattern"`
	IndexCompressionFormats []string `pulumi:"indexCompressionFormats"`
	Key                     string   `pulumi:"key"`
	Notes                   *string  `pulumi:"notes"`
	PackageType             string   `pulumi:"packageType"`
	PrimaryKeypairRef       *string  `pulumi:"primaryKeypairRef"`
	PriorityResolution      *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments     []string `pulumi:"projectEnvironments"`
	ProjectKey              *string  `pulumi:"projectKey"`
	PropertySets            []string `pulumi:"propertySets"`
	RepoLayoutRef           *string  `pulumi:"repoLayoutRef"`
	XrayIndex               *bool    `pulumi:"xrayIndex"`
}

func GetLocalAlpineRepositoryOutput(ctx *pulumi.Context, args GetLocalAlpineRepositoryOutputArgs, opts ...pulumi.InvokeOption) GetLocalAlpineRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocalAlpineRepositoryResult, error) {
			args := v.(GetLocalAlpineRepositoryArgs)
			r, err := GetLocalAlpineRepository(ctx, &args, opts...)
			var s GetLocalAlpineRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocalAlpineRepositoryResultOutput)
}

// A collection of arguments for invoking getLocalAlpineRepository.
type GetLocalAlpineRepositoryOutputArgs struct {
	ArchiveBrowsingEnabled  pulumi.BoolPtrInput     `pulumi:"archiveBrowsingEnabled"`
	BlackedOut              pulumi.BoolPtrInput     `pulumi:"blackedOut"`
	CdnRedirect             pulumi.BoolPtrInput     `pulumi:"cdnRedirect"`
	Description             pulumi.StringPtrInput   `pulumi:"description"`
	DownloadDirect          pulumi.BoolPtrInput     `pulumi:"downloadDirect"`
	ExcludesPattern         pulumi.StringPtrInput   `pulumi:"excludesPattern"`
	IncludesPattern         pulumi.StringPtrInput   `pulumi:"includesPattern"`
	IndexCompressionFormats pulumi.StringArrayInput `pulumi:"indexCompressionFormats"`
	Key                     pulumi.StringInput      `pulumi:"key"`
	Notes                   pulumi.StringPtrInput   `pulumi:"notes"`
	PrimaryKeypairRef       pulumi.StringPtrInput   `pulumi:"primaryKeypairRef"`
	PriorityResolution      pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments     pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey              pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets            pulumi.StringArrayInput `pulumi:"propertySets"`
	RepoLayoutRef           pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	XrayIndex               pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (GetLocalAlpineRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalAlpineRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getLocalAlpineRepository.
type GetLocalAlpineRepositoryResultOutput struct{ *pulumi.OutputState }

func (GetLocalAlpineRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocalAlpineRepositoryResult)(nil)).Elem()
}

func (o GetLocalAlpineRepositoryResultOutput) ToGetLocalAlpineRepositoryResultOutput() GetLocalAlpineRepositoryResultOutput {
	return o
}

func (o GetLocalAlpineRepositoryResultOutput) ToGetLocalAlpineRepositoryResultOutputWithContext(ctx context.Context) GetLocalAlpineRepositoryResultOutput {
	return o
}

func (o GetLocalAlpineRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) string { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocalAlpineRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) string { return v.IncludesPattern }).(pulumi.StringOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) IndexCompressionFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) []string { return v.IndexCompressionFormats }).(pulumi.StringArrayOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *string { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o GetLocalAlpineRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetLocalAlpineRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocalAlpineRepositoryResultOutput{})
}