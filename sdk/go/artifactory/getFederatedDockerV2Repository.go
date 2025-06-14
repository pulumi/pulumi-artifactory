// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a federated Docker repository.
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
//			_, err := artifactory.LookupFederatedDockerV2Repository(ctx, &artifactory.LookupFederatedDockerV2RepositoryArgs{
//				Key: "federated-test-docker-repo",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupFederatedDockerV2Repository(ctx *pulumi.Context, args *LookupFederatedDockerV2RepositoryArgs, opts ...pulumi.InvokeOption) (*LookupFederatedDockerV2RepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFederatedDockerV2RepositoryResult
	err := ctx.Invoke("artifactory:index/getFederatedDockerV2Repository:getFederatedDockerV2Repository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFederatedDockerV2Repository.
type LookupFederatedDockerV2RepositoryArgs struct {
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	BlockPushingSchema1    *bool   `pulumi:"blockPushingSchema1"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    *bool   `pulumi:"disableProxy"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           string `pulumi:"key"`
	MaxUniqueTags *int   `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedDockerV2RepositoryMember `pulumi:"members"`
	Notes               *string                                `pulumi:"notes"`
	PriorityResolution  *bool                                  `pulumi:"priorityResolution"`
	ProjectEnvironments []string                               `pulumi:"projectEnvironments"`
	ProjectKey          *string                                `pulumi:"projectKey"`
	PropertySets        []string                               `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         *string `pulumi:"proxy"`
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	TagRetention  *int    `pulumi:"tagRetention"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getFederatedDockerV2Repository.
type LookupFederatedDockerV2RepositoryResult struct {
	ApiVersion             string  `pulumi:"apiVersion"`
	ArchiveBrowsingEnabled *bool   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             *bool   `pulumi:"blackedOut"`
	BlockPushingSchema1    bool    `pulumi:"blockPushingSchema1"`
	CdnRedirect            *bool   `pulumi:"cdnRedirect"`
	CleanupOnDelete        *bool   `pulumi:"cleanupOnDelete"`
	Description            *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    *bool   `pulumi:"disableProxy"`
	DownloadDirect  *bool   `pulumi:"downloadDirect"`
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// The provider-assigned unique ID for this managed resource.
	Id              string  `pulumi:"id"`
	IncludesPattern *string `pulumi:"includesPattern"`
	Key             string  `pulumi:"key"`
	MaxUniqueTags   *int    `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             []GetFederatedDockerV2RepositoryMember `pulumi:"members"`
	Notes               *string                                `pulumi:"notes"`
	PackageType         string                                 `pulumi:"packageType"`
	PriorityResolution  *bool                                  `pulumi:"priorityResolution"`
	ProjectEnvironments []string                               `pulumi:"projectEnvironments"`
	ProjectKey          *string                                `pulumi:"projectKey"`
	PropertySets        []string                               `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         *string `pulumi:"proxy"`
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	TagRetention  *int    `pulumi:"tagRetention"`
	XrayIndex     *bool   `pulumi:"xrayIndex"`
}

func LookupFederatedDockerV2RepositoryOutput(ctx *pulumi.Context, args LookupFederatedDockerV2RepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupFederatedDockerV2RepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFederatedDockerV2RepositoryResultOutput, error) {
			args := v.(LookupFederatedDockerV2RepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getFederatedDockerV2Repository:getFederatedDockerV2Repository", args, LookupFederatedDockerV2RepositoryResultOutput{}, options).(LookupFederatedDockerV2RepositoryResultOutput), nil
		}).(LookupFederatedDockerV2RepositoryResultOutput)
}

// A collection of arguments for invoking getFederatedDockerV2Repository.
type LookupFederatedDockerV2RepositoryOutputArgs struct {
	ArchiveBrowsingEnabled pulumi.BoolPtrInput   `pulumi:"archiveBrowsingEnabled"`
	BlackedOut             pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	BlockPushingSchema1    pulumi.BoolPtrInput   `pulumi:"blockPushingSchema1"`
	CdnRedirect            pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	CleanupOnDelete        pulumi.BoolPtrInput   `pulumi:"cleanupOnDelete"`
	Description            pulumi.StringPtrInput `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy    pulumi.BoolPtrInput   `pulumi:"disableProxy"`
	DownloadDirect  pulumi.BoolPtrInput   `pulumi:"downloadDirect"`
	ExcludesPattern pulumi.StringPtrInput `pulumi:"excludesPattern"`
	IncludesPattern pulumi.StringPtrInput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key           pulumi.StringInput `pulumi:"key"`
	MaxUniqueTags pulumi.IntPtrInput `pulumi:"maxUniqueTags"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members             GetFederatedDockerV2RepositoryMemberArrayInput `pulumi:"members"`
	Notes               pulumi.StringPtrInput                          `pulumi:"notes"`
	PriorityResolution  pulumi.BoolPtrInput                            `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayInput                        `pulumi:"projectEnvironments"`
	ProjectKey          pulumi.StringPtrInput                          `pulumi:"projectKey"`
	PropertySets        pulumi.StringArrayInput                        `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings.
	Proxy         pulumi.StringPtrInput `pulumi:"proxy"`
	RepoLayoutRef pulumi.StringPtrInput `pulumi:"repoLayoutRef"`
	TagRetention  pulumi.IntPtrInput    `pulumi:"tagRetention"`
	XrayIndex     pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupFederatedDockerV2RepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDockerV2RepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getFederatedDockerV2Repository.
type LookupFederatedDockerV2RepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupFederatedDockerV2RepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFederatedDockerV2RepositoryResult)(nil)).Elem()
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ToLookupFederatedDockerV2RepositoryResultOutput() LookupFederatedDockerV2RepositoryResultOutput {
	return o
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ToLookupFederatedDockerV2RepositoryResultOutputWithContext(ctx context.Context) LookupFederatedDockerV2RepositoryResultOutput {
	return o
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) string { return v.ApiVersion }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) bool { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o LookupFederatedDockerV2RepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupFederatedDockerV2RepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) MaxUniqueTags() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *int { return v.MaxUniqueTags }).(pulumi.IntPtrOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o LookupFederatedDockerV2RepositoryResultOutput) Members() GetFederatedDockerV2RepositoryMemberArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) []GetFederatedDockerV2RepositoryMember {
		return v.Members
	}).(GetFederatedDockerV2RepositoryMemberArrayOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings.
func (o LookupFederatedDockerV2RepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) TagRetention() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *int { return v.TagRetention }).(pulumi.IntPtrOutput)
}

func (o LookupFederatedDockerV2RepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFederatedDockerV2RepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFederatedDockerV2RepositoryResultOutput{})
}
