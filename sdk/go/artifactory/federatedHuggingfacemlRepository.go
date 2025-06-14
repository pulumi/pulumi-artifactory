// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v8/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a federated Hugging Face ML repository.
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
//			_, err := artifactory.NewFederatedHuggingfacemlRepository(ctx, "my-huggingfaceml-federated", &artifactory.FederatedHuggingfacemlRepositoryArgs{
//				Key: pulumi.String("my-huggingfaceml-federated"),
//				Members: artifactory.FederatedHuggingfacemlRepositoryMemberArray{
//					&artifactory.FederatedHuggingfacemlRepositoryMemberArgs{
//						Url:     pulumi.String("http://tempurl.org/artifactory/my-huggingfaceml-federated"),
//						Enabled: pulumi.Bool(true),
//					},
//					&artifactory.FederatedHuggingfacemlRepositoryMemberArgs{
//						Url:     pulumi.String("http://tempurl2.org/artifactory/my-huggingfaceml-federated-2"),
//						Enabled: pulumi.Bool(true),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Federated repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/federatedHuggingfacemlRepository:FederatedHuggingfacemlRepository my-huggingfaceml-federated my-huggingfaceml-federated
// ```
type FederatedHuggingfacemlRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     pulumi.BoolPtrOutput `pulumi:"cdnRedirect"`
	CleanupOnDelete pulumi.BoolPtrOutput `pulumi:"cleanupOnDelete"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy pulumi.BoolPtrOutput `pulumi:"disableProxy"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedHuggingfacemlRepositoryMemberArrayOutput `pulumi:"members"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewFederatedHuggingfacemlRepository registers a new resource with the given unique name, arguments, and options.
func NewFederatedHuggingfacemlRepository(ctx *pulumi.Context,
	name string, args *FederatedHuggingfacemlRepositoryArgs, opts ...pulumi.ResourceOption) (*FederatedHuggingfacemlRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FederatedHuggingfacemlRepository
	err := ctx.RegisterResource("artifactory:index/federatedHuggingfacemlRepository:FederatedHuggingfacemlRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFederatedHuggingfacemlRepository gets an existing FederatedHuggingfacemlRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFederatedHuggingfacemlRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FederatedHuggingfacemlRepositoryState, opts ...pulumi.ResourceOption) (*FederatedHuggingfacemlRepository, error) {
	var resource FederatedHuggingfacemlRepository
	err := ctx.ReadResource("artifactory:index/federatedHuggingfacemlRepository:FederatedHuggingfacemlRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FederatedHuggingfacemlRepository resources.
type federatedHuggingfacemlRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     *bool `pulumi:"cdnRedirect"`
	CleanupOnDelete *bool `pulumi:"cleanupOnDelete"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy *bool `pulumi:"disableProxy"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedHuggingfacemlRepositoryMember `pulumi:"members"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy *string `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type FederatedHuggingfacemlRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     pulumi.BoolPtrInput
	CleanupOnDelete pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedHuggingfacemlRepositoryMemberArrayInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrInput
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedHuggingfacemlRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedHuggingfacemlRepositoryState)(nil)).Elem()
}

type federatedHuggingfacemlRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     *bool `pulumi:"cdnRedirect"`
	CleanupOnDelete *bool `pulumi:"cleanupOnDelete"`
	// Public description.
	Description *string `pulumi:"description"`
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy *bool `pulumi:"disableProxy"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect *bool `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members []FederatedHuggingfacemlRepositoryMember `pulumi:"members"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy *string `pulumi:"proxy"`
	// Repository layout key for the federated repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a FederatedHuggingfacemlRepository resource.
type FederatedHuggingfacemlRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect     pulumi.BoolPtrInput
	CleanupOnDelete pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
	DisableProxy pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// The list of Federated members and must contain this repository URL (configured base URL
	// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
	// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
	// to set up Federated repositories correctly.
	Members FederatedHuggingfacemlRepositoryMemberArrayInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
	Proxy pulumi.StringPtrInput
	// Repository layout key for the federated repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (FederatedHuggingfacemlRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*federatedHuggingfacemlRepositoryArgs)(nil)).Elem()
}

type FederatedHuggingfacemlRepositoryInput interface {
	pulumi.Input

	ToFederatedHuggingfacemlRepositoryOutput() FederatedHuggingfacemlRepositoryOutput
	ToFederatedHuggingfacemlRepositoryOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryOutput
}

func (*FederatedHuggingfacemlRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedHuggingfacemlRepository)(nil)).Elem()
}

func (i *FederatedHuggingfacemlRepository) ToFederatedHuggingfacemlRepositoryOutput() FederatedHuggingfacemlRepositoryOutput {
	return i.ToFederatedHuggingfacemlRepositoryOutputWithContext(context.Background())
}

func (i *FederatedHuggingfacemlRepository) ToFederatedHuggingfacemlRepositoryOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedHuggingfacemlRepositoryOutput)
}

// FederatedHuggingfacemlRepositoryArrayInput is an input type that accepts FederatedHuggingfacemlRepositoryArray and FederatedHuggingfacemlRepositoryArrayOutput values.
// You can construct a concrete instance of `FederatedHuggingfacemlRepositoryArrayInput` via:
//
//	FederatedHuggingfacemlRepositoryArray{ FederatedHuggingfacemlRepositoryArgs{...} }
type FederatedHuggingfacemlRepositoryArrayInput interface {
	pulumi.Input

	ToFederatedHuggingfacemlRepositoryArrayOutput() FederatedHuggingfacemlRepositoryArrayOutput
	ToFederatedHuggingfacemlRepositoryArrayOutputWithContext(context.Context) FederatedHuggingfacemlRepositoryArrayOutput
}

type FederatedHuggingfacemlRepositoryArray []FederatedHuggingfacemlRepositoryInput

func (FederatedHuggingfacemlRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedHuggingfacemlRepository)(nil)).Elem()
}

func (i FederatedHuggingfacemlRepositoryArray) ToFederatedHuggingfacemlRepositoryArrayOutput() FederatedHuggingfacemlRepositoryArrayOutput {
	return i.ToFederatedHuggingfacemlRepositoryArrayOutputWithContext(context.Background())
}

func (i FederatedHuggingfacemlRepositoryArray) ToFederatedHuggingfacemlRepositoryArrayOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedHuggingfacemlRepositoryArrayOutput)
}

// FederatedHuggingfacemlRepositoryMapInput is an input type that accepts FederatedHuggingfacemlRepositoryMap and FederatedHuggingfacemlRepositoryMapOutput values.
// You can construct a concrete instance of `FederatedHuggingfacemlRepositoryMapInput` via:
//
//	FederatedHuggingfacemlRepositoryMap{ "key": FederatedHuggingfacemlRepositoryArgs{...} }
type FederatedHuggingfacemlRepositoryMapInput interface {
	pulumi.Input

	ToFederatedHuggingfacemlRepositoryMapOutput() FederatedHuggingfacemlRepositoryMapOutput
	ToFederatedHuggingfacemlRepositoryMapOutputWithContext(context.Context) FederatedHuggingfacemlRepositoryMapOutput
}

type FederatedHuggingfacemlRepositoryMap map[string]FederatedHuggingfacemlRepositoryInput

func (FederatedHuggingfacemlRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedHuggingfacemlRepository)(nil)).Elem()
}

func (i FederatedHuggingfacemlRepositoryMap) ToFederatedHuggingfacemlRepositoryMapOutput() FederatedHuggingfacemlRepositoryMapOutput {
	return i.ToFederatedHuggingfacemlRepositoryMapOutputWithContext(context.Background())
}

func (i FederatedHuggingfacemlRepositoryMap) ToFederatedHuggingfacemlRepositoryMapOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FederatedHuggingfacemlRepositoryMapOutput)
}

type FederatedHuggingfacemlRepositoryOutput struct{ *pulumi.OutputState }

func (FederatedHuggingfacemlRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FederatedHuggingfacemlRepository)(nil)).Elem()
}

func (o FederatedHuggingfacemlRepositoryOutput) ToFederatedHuggingfacemlRepositoryOutput() FederatedHuggingfacemlRepositoryOutput {
	return o
}

func (o FederatedHuggingfacemlRepositoryOutput) ToFederatedHuggingfacemlRepositoryOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o FederatedHuggingfacemlRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o FederatedHuggingfacemlRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o FederatedHuggingfacemlRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

func (o FederatedHuggingfacemlRepositoryOutput) CleanupOnDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.CleanupOnDelete }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o FederatedHuggingfacemlRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set to `true`, the proxy is disabled, and not returned in the API response body. If there is a default proxy set for the Artifactory instance, it will be ignored, too.
func (o FederatedHuggingfacemlRepositoryOutput) DisableProxy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.DisableProxy }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o FederatedHuggingfacemlRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o FederatedHuggingfacemlRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o FederatedHuggingfacemlRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o FederatedHuggingfacemlRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The list of Federated members and must contain this repository URL (configured base URL
// `/artifactory/` + repo `key`). Note that each of the federated members will need to have a base URL set.
// Please follow the [instruction](https://www.jfrog.com/confluence/display/JFROG/Working+with+Federated+Repositories#WorkingwithFederatedRepositories-SettingUpaFederatedRepository)
// to set up Federated repositories correctly.
func (o FederatedHuggingfacemlRepositoryOutput) Members() FederatedHuggingfacemlRepositoryMemberArrayOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) FederatedHuggingfacemlRepositoryMemberArrayOutput {
		return v.Members
	}).(FederatedHuggingfacemlRepositoryMemberArrayOutput)
}

// Internal description.
func (o FederatedHuggingfacemlRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o FederatedHuggingfacemlRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o FederatedHuggingfacemlRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o FederatedHuggingfacemlRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o FederatedHuggingfacemlRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o FederatedHuggingfacemlRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Proxy key from Artifactory Proxies settings. Default is empty field. Can't be set if `disableProxy = true`.
func (o FederatedHuggingfacemlRepositoryOutput) Proxy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.Proxy }).(pulumi.StringPtrOutput)
}

// Repository layout key for the federated repository
func (o FederatedHuggingfacemlRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o FederatedHuggingfacemlRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FederatedHuggingfacemlRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type FederatedHuggingfacemlRepositoryArrayOutput struct{ *pulumi.OutputState }

func (FederatedHuggingfacemlRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FederatedHuggingfacemlRepository)(nil)).Elem()
}

func (o FederatedHuggingfacemlRepositoryArrayOutput) ToFederatedHuggingfacemlRepositoryArrayOutput() FederatedHuggingfacemlRepositoryArrayOutput {
	return o
}

func (o FederatedHuggingfacemlRepositoryArrayOutput) ToFederatedHuggingfacemlRepositoryArrayOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryArrayOutput {
	return o
}

func (o FederatedHuggingfacemlRepositoryArrayOutput) Index(i pulumi.IntInput) FederatedHuggingfacemlRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FederatedHuggingfacemlRepository {
		return vs[0].([]*FederatedHuggingfacemlRepository)[vs[1].(int)]
	}).(FederatedHuggingfacemlRepositoryOutput)
}

type FederatedHuggingfacemlRepositoryMapOutput struct{ *pulumi.OutputState }

func (FederatedHuggingfacemlRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FederatedHuggingfacemlRepository)(nil)).Elem()
}

func (o FederatedHuggingfacemlRepositoryMapOutput) ToFederatedHuggingfacemlRepositoryMapOutput() FederatedHuggingfacemlRepositoryMapOutput {
	return o
}

func (o FederatedHuggingfacemlRepositoryMapOutput) ToFederatedHuggingfacemlRepositoryMapOutputWithContext(ctx context.Context) FederatedHuggingfacemlRepositoryMapOutput {
	return o
}

func (o FederatedHuggingfacemlRepositoryMapOutput) MapIndex(k pulumi.StringInput) FederatedHuggingfacemlRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FederatedHuggingfacemlRepository {
		return vs[0].(map[string]*FederatedHuggingfacemlRepository)[vs[1].(string)]
	}).(FederatedHuggingfacemlRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedHuggingfacemlRepositoryInput)(nil)).Elem(), &FederatedHuggingfacemlRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedHuggingfacemlRepositoryArrayInput)(nil)).Elem(), FederatedHuggingfacemlRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FederatedHuggingfacemlRepositoryMapInput)(nil)).Elem(), FederatedHuggingfacemlRepositoryMap{})
	pulumi.RegisterOutputType(FederatedHuggingfacemlRepositoryOutput{})
	pulumi.RegisterOutputType(FederatedHuggingfacemlRepositoryArrayOutput{})
	pulumi.RegisterOutputType(FederatedHuggingfacemlRepositoryMapOutput{})
}
