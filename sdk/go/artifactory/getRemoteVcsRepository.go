// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote VCS repository.
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
//			_, err := artifactory.LookupRemoteVcsRepository(ctx, &artifactory.LookupRemoteVcsRepositoryArgs{
//				Key: "remote-vcs",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupRemoteVcsRepository(ctx *pulumi.Context, args *LookupRemoteVcsRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemoteVcsRepositoryResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRemoteVcsRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemoteVcsRepository:getRemoteVcsRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemoteVcsRepository.
type LookupRemoteVcsRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                         `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                         `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                          `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                         `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                         `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                         `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                         `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemoteVcsRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                       `pulumi:"description"`
	DisableProxy              *bool                                         `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                         `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                         `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                         `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                       `pulumi:"excludesPattern"`
	HardFail                  *bool                                         `pulumi:"hardFail"`
	IncludesPattern           *string                                       `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                   string  `pulumi:"key"`
	ListRemoteFolderItems *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress          *string `pulumi:"localAddress"`
	// (Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots                *int     `pulumi:"maxUniqueSnapshots"`
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
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`. Default value is `GITHUB`
	VcsGitProvider *string `pulumi:"vcsGitProvider"`
	XrayIndex      *bool   `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemoteVcsRepository.
type LookupRemoteVcsRepositoryResult struct {
	AllowAnyHostAuth          *bool                                        `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    *bool                                        `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  *int                                         `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                        `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                        `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                        `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                        `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                       `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteVcsRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                      `pulumi:"description"`
	DisableProxy              *bool                                        `pulumi:"disableProxy"`
	DisableUrlNormalization   *bool                                        `pulumi:"disableUrlNormalization"`
	DownloadDirect            *bool                                        `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                        `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                      `pulumi:"excludesPattern"`
	HardFail                  *bool                                        `pulumi:"hardFail"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string  `pulumi:"id"`
	IncludesPattern       *string `pulumi:"includesPattern"`
	Key                   string  `pulumi:"key"`
	ListRemoteFolderItems *bool   `pulumi:"listRemoteFolderItems"`
	LocalAddress          *string `pulumi:"localAddress"`
	// (Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots                *int     `pulumi:"maxUniqueSnapshots"`
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
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`. Default value is `GITHUB`
	VcsGitProvider *string `pulumi:"vcsGitProvider"`
	XrayIndex      *bool   `pulumi:"xrayIndex"`
}

func LookupRemoteVcsRepositoryOutput(ctx *pulumi.Context, args LookupRemoteVcsRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemoteVcsRepositoryResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRemoteVcsRepositoryResultOutput, error) {
			args := v.(LookupRemoteVcsRepositoryArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("artifactory:index/getRemoteVcsRepository:getRemoteVcsRepository", args, LookupRemoteVcsRepositoryResultOutput{}, options).(LookupRemoteVcsRepositoryResultOutput), nil
		}).(LookupRemoteVcsRepositoryResultOutput)
}

// A collection of arguments for invoking getRemoteVcsRepository.
type LookupRemoteVcsRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                  `pulumi:"allowAnyHostAuth"`
	ArchiveBrowsingEnabled    pulumi.BoolPtrInput                                  `pulumi:"archiveBrowsingEnabled"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                   `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                  `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                  `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                  `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                  `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemoteVcsRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                `pulumi:"description"`
	DisableProxy              pulumi.BoolPtrInput                                  `pulumi:"disableProxy"`
	DisableUrlNormalization   pulumi.BoolPtrInput                                  `pulumi:"disableUrlNormalization"`
	DownloadDirect            pulumi.BoolPtrInput                                  `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                  `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                  `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key                   pulumi.StringInput    `pulumi:"key"`
	ListRemoteFolderItems pulumi.BoolPtrInput   `pulumi:"listRemoteFolderItems"`
	LocalAddress          pulumi.StringPtrInput `pulumi:"localAddress"`
	// (Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
	MaxUniqueSnapshots                pulumi.IntPtrInput      `pulumi:"maxUniqueSnapshots"`
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
	// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`. Default value is `GITHUB`
	VcsGitProvider pulumi.StringPtrInput `pulumi:"vcsGitProvider"`
	XrayIndex      pulumi.BoolPtrInput   `pulumi:"xrayIndex"`
}

func (LookupRemoteVcsRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteVcsRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemoteVcsRepository.
type LookupRemoteVcsRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemoteVcsRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemoteVcsRepositoryResult)(nil)).Elem()
}

func (o LookupRemoteVcsRepositoryResultOutput) ToLookupRemoteVcsRepositoryResultOutput() LookupRemoteVcsRepositoryResultOutput {
	return o
}

func (o LookupRemoteVcsRepositoryResultOutput) ToLookupRemoteVcsRepositoryResultOutputWithContext(ctx context.Context) LookupRemoteVcsRepositoryResultOutput {
	return o
}

func (o LookupRemoteVcsRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ContentSynchronisation() GetRemoteVcsRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) GetRemoteVcsRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemoteVcsRepositoryContentSynchronisationOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) DisableUrlNormalization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.DisableUrlNormalization }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemoteVcsRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

// (Optional) The maximum number of unique snapshots of a single artifact to store. Once the number of snapshots exceeds this setting, older versions are removed. A value of 0 (default) indicates there is no limit, and unique snapshots are not cleaned up.
func (o LookupRemoteVcsRepositoryResultOutput) MaxUniqueSnapshots() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.MaxUniqueSnapshots }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

// (Optional) This attribute is used when vcsGitProvider is set to `CUSTOM`. Provided URL will be used as proxy.
func (o LookupRemoteVcsRepositoryResultOutput) VcsGitDownloadUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.VcsGitDownloadUrl }).(pulumi.StringPtrOutput)
}

// (Optional) Artifactory supports proxying the following Git providers out-of-the-box: GitHub, Bitbucket, Stash, a remote Artifactory instance or a custom Git repository. Allowed values are: `GITHUB`, `BITBUCKET`, `OLDSTASH`, `STASH`, `ARTIFACTORY`, `CUSTOM`. Default value is `GITHUB`
func (o LookupRemoteVcsRepositoryResultOutput) VcsGitProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *string { return v.VcsGitProvider }).(pulumi.StringPtrOutput)
}

func (o LookupRemoteVcsRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemoteVcsRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemoteVcsRepositoryResultOutput{})
}
