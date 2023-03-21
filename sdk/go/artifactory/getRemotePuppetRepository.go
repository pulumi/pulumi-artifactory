// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves a remote Peppet repository.
func LookupRemotePuppetRepository(ctx *pulumi.Context, args *LookupRemotePuppetRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRemotePuppetRepositoryResult, error) {
	var rv LookupRemotePuppetRepositoryResult
	err := ctx.Invoke("artifactory:index/getRemotePuppetRepository:getRemotePuppetRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRemotePuppetRepository.
type LookupRemotePuppetRepositoryArgs struct {
	AllowAnyHostAuth          *bool                                            `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                             `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                            `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                            `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                            `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                            `pulumi:"cdnRedirect"`
	ClientTlsCertificate      *string                                          `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    *GetRemotePuppetRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                          `pulumi:"description"`
	DownloadDirect            *bool                                            `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                            `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                          `pulumi:"excludesPattern"`
	HardFail                  *bool                                            `pulumi:"hardFail"`
	IncludesPattern           *string                                          `pulumi:"includesPattern"`
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
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

// A collection of values returned by getRemotePuppetRepository.
type LookupRemotePuppetRepositoryResult struct {
	AllowAnyHostAuth          *bool                                           `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  *int                                            `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                *bool                                           `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes *bool                                           `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        *bool                                           `pulumi:"bypassHeadRequests"`
	CdnRedirect               *bool                                           `pulumi:"cdnRedirect"`
	ClientTlsCertificate      string                                          `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemotePuppetRepositoryContentSynchronisation `pulumi:"contentSynchronisation"`
	Description               *string                                         `pulumi:"description"`
	DownloadDirect            *bool                                           `pulumi:"downloadDirect"`
	EnableCookieManagement    *bool                                           `pulumi:"enableCookieManagement"`
	ExcludesPattern           *string                                         `pulumi:"excludesPattern"`
	HardFail                  *bool                                           `pulumi:"hardFail"`
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
	XrayIndex                         *bool    `pulumi:"xrayIndex"`
}

func LookupRemotePuppetRepositoryOutput(ctx *pulumi.Context, args LookupRemotePuppetRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRemotePuppetRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRemotePuppetRepositoryResult, error) {
			args := v.(LookupRemotePuppetRepositoryArgs)
			r, err := LookupRemotePuppetRepository(ctx, &args, opts...)
			var s LookupRemotePuppetRepositoryResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRemotePuppetRepositoryResultOutput)
}

// A collection of arguments for invoking getRemotePuppetRepository.
type LookupRemotePuppetRepositoryOutputArgs struct {
	AllowAnyHostAuth          pulumi.BoolPtrInput                                     `pulumi:"allowAnyHostAuth"`
	AssumedOfflinePeriodSecs  pulumi.IntPtrInput                                      `pulumi:"assumedOfflinePeriodSecs"`
	BlackedOut                pulumi.BoolPtrInput                                     `pulumi:"blackedOut"`
	BlockMismatchingMimeTypes pulumi.BoolPtrInput                                     `pulumi:"blockMismatchingMimeTypes"`
	BypassHeadRequests        pulumi.BoolPtrInput                                     `pulumi:"bypassHeadRequests"`
	CdnRedirect               pulumi.BoolPtrInput                                     `pulumi:"cdnRedirect"`
	ClientTlsCertificate      pulumi.StringPtrInput                                   `pulumi:"clientTlsCertificate"`
	ContentSynchronisation    GetRemotePuppetRepositoryContentSynchronisationPtrInput `pulumi:"contentSynchronisation"`
	Description               pulumi.StringPtrInput                                   `pulumi:"description"`
	DownloadDirect            pulumi.BoolPtrInput                                     `pulumi:"downloadDirect"`
	EnableCookieManagement    pulumi.BoolPtrInput                                     `pulumi:"enableCookieManagement"`
	ExcludesPattern           pulumi.StringPtrInput                                   `pulumi:"excludesPattern"`
	HardFail                  pulumi.BoolPtrInput                                     `pulumi:"hardFail"`
	IncludesPattern           pulumi.StringPtrInput                                   `pulumi:"includesPattern"`
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
	XrayIndex                         pulumi.BoolPtrInput     `pulumi:"xrayIndex"`
}

func (LookupRemotePuppetRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemotePuppetRepositoryArgs)(nil)).Elem()
}

// A collection of values returned by getRemotePuppetRepository.
type LookupRemotePuppetRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRemotePuppetRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRemotePuppetRepositoryResult)(nil)).Elem()
}

func (o LookupRemotePuppetRepositoryResultOutput) ToLookupRemotePuppetRepositoryResultOutput() LookupRemotePuppetRepositoryResultOutput {
	return o
}

func (o LookupRemotePuppetRepositoryResultOutput) ToLookupRemotePuppetRepositoryResultOutputWithContext(ctx context.Context) LookupRemotePuppetRepositoryResultOutput {
	return o
}

func (o LookupRemotePuppetRepositoryResultOutput) AllowAnyHostAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.AllowAnyHostAuth }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) AssumedOfflinePeriodSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *int { return v.AssumedOfflinePeriodSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) BlockMismatchingMimeTypes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.BlockMismatchingMimeTypes }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) BypassHeadRequests() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.BypassHeadRequests }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ClientTlsCertificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) string { return v.ClientTlsCertificate }).(pulumi.StringOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ContentSynchronisation() GetRemotePuppetRepositoryContentSynchronisationOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) GetRemotePuppetRepositoryContentSynchronisation {
		return v.ContentSynchronisation
	}).(GetRemotePuppetRepositoryContentSynchronisationOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) EnableCookieManagement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.EnableCookieManagement }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) HardFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.HardFail }).(pulumi.BoolPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRemotePuppetRepositoryResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) string { return v.Key }).(pulumi.StringOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ListRemoteFolderItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.ListRemoteFolderItems }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) LocalAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.LocalAddress }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) MetadataRetrievalTimeoutSecs() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *int { return v.MetadataRetrievalTimeoutSecs }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) MismatchingMimeTypesOverrideList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.MismatchingMimeTypesOverrideList }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) MissedCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *int { return v.MissedCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Offline() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.Offline }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) string { return v.PackageType }).(pulumi.StringOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) []string { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) []string { return v.PropertySets }).(pulumi.StringArrayOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.Proxy }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) QueryParams() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.QueryParams }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) RemoteRepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.RemoteRepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *int { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) ShareConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) bool { return v.ShareConfiguration }).(pulumi.BoolOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) SocketTimeoutMillis() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *int { return v.SocketTimeoutMillis }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) StoreArtifactsLocally() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.StoreArtifactsLocally }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) SynchronizeProperties() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.SynchronizeProperties }).(pulumi.BoolPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) UnusedArtifactsCleanupPeriodHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *int { return v.UnusedArtifactsCleanupPeriodHours }).(pulumi.IntPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *string { return v.Username }).(pulumi.StringPtrOutput)
}

func (o LookupRemotePuppetRepositoryResultOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupRemotePuppetRepositoryResult) *bool { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRemotePuppetRepositoryResultOutput{})
}