// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Composer repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v3/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.LookupRemoteComposerRepository(ctx, &artifactory.LookupRemoteComposerRepositoryArgs{
//				Key: "remote-composer",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteComposerRepository(ctx *pulumi.Context, args *LookupRemoteComposerRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteComposerRepositoryResult, error) {
	var rv LookupRemoteComposerRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteComposerRepository:getRemoteComposerRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteComposerRepository.
type LookupRemoteComposerRepositoryArgs struct {
	AllowAnyHostAuth          *bool   `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int    `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool   `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool   `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool   `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string `pulumi:"clientTlsCertificate"`
	// (Optional) Proxy remote Composer repository. Default value is `https://packagist.org`.
	ComposerRegistryUrl    *string                                            `pulumi:"composerRegistryUrl"`
	ContentSynchronisation *GetRemoteComposerRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                            `pulumi:"description"`
	DownloadDirect         *bool                                              `pulumi:"downloadDirect"`
	EnableCookieManagement *bool                                              `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string                                            `pulumi:"excludesPattern"`
	HardFail               *bool                                              `pulumi:"hardFail"`
	IncludesPattern        *string                                            `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                               string   `pulumi:"key"`
	ListRemoteFolderItems             *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                      *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                             *string  `pulumi:"notes"`
	Offline                           *bool    `pulumi:"offline"`
	Password                          *string  `pulumi:"password"`
	PriorityResolution                *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments               []string `pulumi:"projectEnvironments"`
	ProjectKey                        *string  `pulumi:"projectKey"`
	PropertySets                      []string `pulumi:"propertySets"`
	Proxy                             *string  `pulumi:"proxy"`
	QueryParams                       *string  `pulumi:"queryParams"`
	RemoteRepoLayoutRef               *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     *string  `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       *int     `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                *bool    `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int     `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool    `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             *bool    `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int     `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string  `pulumi:"url"`
	Username                          *string  `pulumi:"username"`
	// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
	VcsGitDownloadUrl *string `pulumi:"vcsGitDownloadUrl"`
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
	VcsGitProvider *string `pulumi:"vcsGitProvider"`
	XrayIndex      *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteComposerRepository.
type LookupRemoteComposerRepositoryResult struct {
	AllowAnyHostAuth          *bool  `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int   `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool  `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool  `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool  `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool  `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string `pulumi:"clientTlsCertificate"`
	// (Optional) Proxy remote Composer repository. Default value is `https://packagist.org`.
	ComposerRegistryUrl    *string                                           `pulumi:"composerRegistryUrl"`
	ContentSynchronisation GetRemoteComposerRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description            *string                                           `pulumi:"description"`
	DownloadDirect         *bool                                             `pulumi:"downloadDirect"`
	EnableCookieManagement *bool                                             `pulumi:"enableCookieManagement"`
	ExcludesPattern        *string                                           `pulumi:"excludesPattern"`
	HardFail               *bool                                             `pulumi:"hardFail"`
	// The provider-assigned unique ID for this managed resource.
	Id                                string   `pulumi:"id"`
	IncludesPattern                   *string  `pulumi:"includesPattern"`
	Key                               string   `pulumi:"key"`
	ListRemoteFolderItems             *bool    `pulumi:"listRemoteFolderItems"`
	LocalAddress                      *string  `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      *int     `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  *string  `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          *int     `pulumi:"missedCachePeriodSeconds"`
	Notes                             *string  `pulumi:"notes"`
	Offline                           *bool    `pulumi:"offline"`
	PackageType                       string   `pulumi:"packageType"`
	Password                          *string  `pulumi:"password"`
	PriorityResolution                *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments               []string `pulumi:"projectEnvironments"`
	ProjectKey                        *string  `pulumi:"projectKey"`
	PropertySets                      []string `pulumi:"propertySets"`
	Proxy                             *string  `pulumi:"proxy"`
	QueryParams                       *string  `pulumi:"queryParams"`
	RemoteRepoLayoutRef               *string  `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     *string  `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       *int     `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                bool     `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               *int     `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             *bool    `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             *bool    `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours *int     `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               *string  `pulumi:"url"`
	Username                          *string  `pulumi:"username"`
	// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
	VcsGitDownloadUrl *string `pulumi:"vcsGitDownloadUrl"`
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
	VcsGitProvider *string `pulumi:"vcsGitProvider"`
	XrayIndex      *bool   `pulumi:"xrayIndex"`
}

func LookupRemoteComposerRepositoryOutput(ctx *pulumi.Context, args LookupRemoteComposerRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteComposerRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemoteComposerRepositoryResult, error) {
			args := v.(LookupRemoteComposerRepositoryArgs)
			r, err := LookupRemoteComposerRepository(ctx, &args, opts...)
			var s LookupRemoteComposerRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemoteComposerRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteComposerRepository.
type LookupRemoteComposerRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput   `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput    `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput   `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput   `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput   `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput   `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput `pulumi:"clientTlsCertificate"`
	// (Optional) Proxy remote Composer repository. Default value is `https://packagist.org`.
	ComposerRegistryUrl    pulumi.StringPtrInput                                     `pulumi:"composerRegistryUrl"`
	ContentSynchronisation GetRemoteComposerRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description            pulumi.StringPtrInput                                     `pulumi:"description"`
	DownloadDirect         pulumi.BoolPtrInput                                       `pulumi:"downloadDirect"`
	EnableCookieManagement pulumi.BoolPtrInput                                       `pulumi:"enableCookieManagement"`
	ExcludesPattern        pulumi.StringPtrInput                                     `pulumi:"excludesPattern"`
	HardFail               pulumi.BoolPtrInput                                       `pulumi:"hardFail"`
	IncludesPattern        pulumi.StringPtrInput                                     `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                               pulumi.StringInput      `pulumi:"key"`
	ListRemoteFolderItems             pulumi.BoolPtrInput     `pulumi:"listRemoteFolderItems"`
	LocalAddress                      pulumi.StringPtrInput   `pulumi:"localAddress"`
	MetadataRetrievalTimeoutSecs      pulumi.IntPtrInput      `pulumi:"metadataRetrievalTimeoutSecs"`
	MismatchingMimeTypesOverrideList  pulumi.StringPtrInput   `pulumi:"mismatchingMimeTypesOverrideList"`
	MissedCachePeriodSeconds          pulumi.IntPtrInput      `pulumi:"missedCachePeriodSeconds"`
	Notes                             pulumi.StringPtrInput   `pulumi:"notes"`
	Offline                           pulumi.BoolPtrInput     `pulumi:"offline"`
	Password                          pulumi.StringPtrInput   `pulumi:"password"`
	PriorityResolution                pulumi.BoolPtrInput     `pulumi:"priorityResolution"`
	ProjectEnvironments               pulumi.StringArrayInput `pulumi:"projectEnvironments"`
	ProjectKey                        pulumi.StringPtrInput   `pulumi:"projectKey"`
	PropertySets                      pulumi.StringArrayInput `pulumi:"propertySets"`
	Proxy                             pulumi.StringPtrInput   `pulumi:"proxy"`
	QueryParams                       pulumi.StringPtrInput   `pulumi:"queryParams"`
	RemoteRepoLayoutRef               pulumi.StringPtrInput   `pulumi:"remoteRepoLayoutRef"`
	RepoLayoutRef                     pulumi.StringPtrInput   `pulumi:"repoLayoutRef"`
	RetrievalCachePeriodSeconds       pulumi.IntPtrInput      `pulumi:"retrievalCachePeriodSeconds"`
	ShareConfiguration                pulumi.BoolPtrInput     `pulumi:"shareConfiguration"`
	SocketTimeoutMillis               pulumi.IntPtrInput      `pulumi:"socketTimeoutMillis"`
	StoreArtifactsLocally             pulumi.BoolPtrInput     `pulumi:"storeArtifactsLocally"`
	SynchronizeProperties             pulumi.BoolPtrInput     `pulumi:"synchronizeProperties"`
	UnusedArtifactsCleanupPeriodHours pulumi.IntPtrInput      `pulumi:"unusedArtifactsCleanupPeriodHours"`
	Url                               pulumi.StringPtrInput   `pulumi:"url"`
	Username                          pulumi.StringPtrInput   `pulumi:"username"`
	// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
	VcsGitDownloadUrl pulumi.StringPtrInput `pulumi:"vcsGitDownloadUrl"`
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
	VcsGitProvider pulumi.StringPtrInput `pulumi:"vcsGitProvider"`
	XrayIndex      pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupRemoteComposerRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteComposerRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteComposerRepository.
type LookupRemoteComposerRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteComposerRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteComposerRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteComposerRepositoryResultOutput) ToLookupRemoteComposerRepositoryResultOutput() LookupRemoteComposerRepositoryResultOutput {
	return o
}

func (o LookupRemoteComposerRepositoryResultOutput) ToLookupRemoteComposerRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteComposerRepositoryResultOutput {
	return o
}

func (o LookupRemoteComposerRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

// (Optional) Proxy remote Composer repository. Default value is `https://packagist.org`.
func (o LookupRemoteComposerRepositoryResultOutput) ComposerRegistryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.ComposerRegistryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ContentSynchronisation() GetRemoteComposerRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) GetRemoteComposerRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteComposerRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteComposerRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
func (o LookupRemoteComposerRepositoryResultOutput) VcsGitDownloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.VcsGitDownloadUrl }).(pulumi.StringPtrOutput)
}

// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub or a remote Artifactory instance. Default value is `GITHUB`. Possible values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`.
func (o LookupRemoteComposerRepositoryResultOutput) VcsGitProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *string { return v.VcsGitProvider }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteComposerRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteComposerRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteComposerRepositoryResultOutput{})
}