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

// Creates a local Hugging Face repository.
//
// Official documentation can be found [here](https://jfrog.com/help/r/jfrog-artifactory-documentation/set-up-local-hugging-face-repositories).
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
//			_, err := artifactory.NewLocalHuggingfacemlRepository(ctx, "local-huggingfaceml-repo", &artifactory.LocalHuggingfacemlRepositoryArgs{
//				Key: pulumi.String("local-huggingfaceml-repo"),
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
// Local repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/localHuggingfacemlRepository:LocalHuggingfacemlRepository local-huggingfaceml-repo local-huggingfaceml-repo
// ```
type LocalHuggingfacemlRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolOutput `pulumi:"cdnRedirect"`
	// Public description.
	Description pulumi.StringOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes pulumi.StringOutput `pulumi:"notes"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolOutput        `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolOutput `pulumi:"xrayIndex"`
}

// NewLocalHuggingfacemlRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalHuggingfacemlRepository(ctx *pulumi.Context,
	name string, args *LocalHuggingfacemlRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalHuggingfacemlRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalHuggingfacemlRepository
	err := ctx.RegisterResource("artifactory:index/localHuggingfacemlRepository:LocalHuggingfacemlRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalHuggingfacemlRepository gets an existing LocalHuggingfacemlRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalHuggingfacemlRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalHuggingfacemlRepositoryState, opts ...pulumi.ResourceOption) (*LocalHuggingfacemlRepository, error) {
	var resource LocalHuggingfacemlRepository
	err := ctx.ReadResource("artifactory:index/localHuggingfacemlRepository:LocalHuggingfacemlRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalHuggingfacemlRepository resources.
type localHuggingfacemlRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalHuggingfacemlRepositoryState struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalHuggingfacemlRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localHuggingfacemlRepositoryState)(nil)).Elem()
}

type localHuggingfacemlRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect *bool `pulumi:"cdnRedirect"`
	// Public description.
	Description *string `pulumi:"description"`
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalHuggingfacemlRepository resource.
type LocalHuggingfacemlRepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrInput
	// Public description.
	Description pulumi.StringPtrInput
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
	// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
	// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalHuggingfacemlRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localHuggingfacemlRepositoryArgs)(nil)).Elem()
}

type LocalHuggingfacemlRepositoryInput interface {
	pulumi.Input

	ToLocalHuggingfacemlRepositoryOutput() LocalHuggingfacemlRepositoryOutput
	ToLocalHuggingfacemlRepositoryOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryOutput
}

func (*LocalHuggingfacemlRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalHuggingfacemlRepository)(nil)).Elem()
}

func (i *LocalHuggingfacemlRepository) ToLocalHuggingfacemlRepositoryOutput() LocalHuggingfacemlRepositoryOutput {
	return i.ToLocalHuggingfacemlRepositoryOutputWithContext(context.Background())
}

func (i *LocalHuggingfacemlRepository) ToLocalHuggingfacemlRepositoryOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHuggingfacemlRepositoryOutput)
}

// LocalHuggingfacemlRepositoryArrayInput is an input type that accepts LocalHuggingfacemlRepositoryArray and LocalHuggingfacemlRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalHuggingfacemlRepositoryArrayInput` via:
//
//	LocalHuggingfacemlRepositoryArray{ LocalHuggingfacemlRepositoryArgs{...} }
type LocalHuggingfacemlRepositoryArrayInput interface {
	pulumi.Input

	ToLocalHuggingfacemlRepositoryArrayOutput() LocalHuggingfacemlRepositoryArrayOutput
	ToLocalHuggingfacemlRepositoryArrayOutputWithContext(context.Context) LocalHuggingfacemlRepositoryArrayOutput
}

type LocalHuggingfacemlRepositoryArray []LocalHuggingfacemlRepositoryInput

func (LocalHuggingfacemlRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalHuggingfacemlRepository)(nil)).Elem()
}

func (i LocalHuggingfacemlRepositoryArray) ToLocalHuggingfacemlRepositoryArrayOutput() LocalHuggingfacemlRepositoryArrayOutput {
	return i.ToLocalHuggingfacemlRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalHuggingfacemlRepositoryArray) ToLocalHuggingfacemlRepositoryArrayOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHuggingfacemlRepositoryArrayOutput)
}

// LocalHuggingfacemlRepositoryMapInput is an input type that accepts LocalHuggingfacemlRepositoryMap and LocalHuggingfacemlRepositoryMapOutput values.
// You can construct a concrete instance of `LocalHuggingfacemlRepositoryMapInput` via:
//
//	LocalHuggingfacemlRepositoryMap{ "key": LocalHuggingfacemlRepositoryArgs{...} }
type LocalHuggingfacemlRepositoryMapInput interface {
	pulumi.Input

	ToLocalHuggingfacemlRepositoryMapOutput() LocalHuggingfacemlRepositoryMapOutput
	ToLocalHuggingfacemlRepositoryMapOutputWithContext(context.Context) LocalHuggingfacemlRepositoryMapOutput
}

type LocalHuggingfacemlRepositoryMap map[string]LocalHuggingfacemlRepositoryInput

func (LocalHuggingfacemlRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalHuggingfacemlRepository)(nil)).Elem()
}

func (i LocalHuggingfacemlRepositoryMap) ToLocalHuggingfacemlRepositoryMapOutput() LocalHuggingfacemlRepositoryMapOutput {
	return i.ToLocalHuggingfacemlRepositoryMapOutputWithContext(context.Background())
}

func (i LocalHuggingfacemlRepositoryMap) ToLocalHuggingfacemlRepositoryMapOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalHuggingfacemlRepositoryMapOutput)
}

type LocalHuggingfacemlRepositoryOutput struct{ *pulumi.OutputState }

func (LocalHuggingfacemlRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalHuggingfacemlRepository)(nil)).Elem()
}

func (o LocalHuggingfacemlRepositoryOutput) ToLocalHuggingfacemlRepositoryOutput() LocalHuggingfacemlRepositoryOutput {
	return o
}

func (o LocalHuggingfacemlRepositoryOutput) ToLocalHuggingfacemlRepositoryOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalHuggingfacemlRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalHuggingfacemlRepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalHuggingfacemlRepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o LocalHuggingfacemlRepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalHuggingfacemlRepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o LocalHuggingfacemlRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o LocalHuggingfacemlRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalHuggingfacemlRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalHuggingfacemlRepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalHuggingfacemlRepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o LocalHuggingfacemlRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalHuggingfacemlRepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o LocalHuggingfacemlRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o LocalHuggingfacemlRepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalHuggingfacemlRepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalHuggingfacemlRepository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type LocalHuggingfacemlRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalHuggingfacemlRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalHuggingfacemlRepository)(nil)).Elem()
}

func (o LocalHuggingfacemlRepositoryArrayOutput) ToLocalHuggingfacemlRepositoryArrayOutput() LocalHuggingfacemlRepositoryArrayOutput {
	return o
}

func (o LocalHuggingfacemlRepositoryArrayOutput) ToLocalHuggingfacemlRepositoryArrayOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryArrayOutput {
	return o
}

func (o LocalHuggingfacemlRepositoryArrayOutput) Index(i pulumi.IntInput) LocalHuggingfacemlRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalHuggingfacemlRepository {
		return vs[0].([]*LocalHuggingfacemlRepository)[vs[1].(int)]
	}).(LocalHuggingfacemlRepositoryOutput)
}

type LocalHuggingfacemlRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalHuggingfacemlRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalHuggingfacemlRepository)(nil)).Elem()
}

func (o LocalHuggingfacemlRepositoryMapOutput) ToLocalHuggingfacemlRepositoryMapOutput() LocalHuggingfacemlRepositoryMapOutput {
	return o
}

func (o LocalHuggingfacemlRepositoryMapOutput) ToLocalHuggingfacemlRepositoryMapOutputWithContext(ctx context.Context) LocalHuggingfacemlRepositoryMapOutput {
	return o
}

func (o LocalHuggingfacemlRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalHuggingfacemlRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalHuggingfacemlRepository {
		return vs[0].(map[string]*LocalHuggingfacemlRepository)[vs[1].(string)]
	}).(LocalHuggingfacemlRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHuggingfacemlRepositoryInput)(nil)).Elem(), &LocalHuggingfacemlRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHuggingfacemlRepositoryArrayInput)(nil)).Elem(), LocalHuggingfacemlRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalHuggingfacemlRepositoryMapInput)(nil)).Elem(), LocalHuggingfacemlRepositoryMap{})
	pulumi.RegisterOutputType(LocalHuggingfacemlRepositoryOutput{})
	pulumi.RegisterOutputType(LocalHuggingfacemlRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalHuggingfacemlRepositoryMapOutput{})
}
