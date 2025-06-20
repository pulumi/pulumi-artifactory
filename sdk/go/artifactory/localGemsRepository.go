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

// Creates a local Gems repository.
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
//			_, err := artifactory.NewLocalGemsRepository(ctx, "terraform-local-test-gems-repo", &artifactory.LocalGemsRepositoryArgs{
//				Key: pulumi.String("terraform-local-test-gems-repo"),
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
// $ pulumi import artifactory:index/localGemsRepository:LocalGemsRepository terraform-local-test-gems-repo terraform-local-test-gems-repo
// ```
type LocalGemsRepository struct {
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

// NewLocalGemsRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalGemsRepository(ctx *pulumi.Context,
	name string, args *LocalGemsRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalGemsRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalGemsRepository
	err := ctx.RegisterResource("artifactory:index/localGemsRepository:LocalGemsRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalGemsRepository gets an existing LocalGemsRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalGemsRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalGemsRepositoryState, opts ...pulumi.ResourceOption) (*LocalGemsRepository, error) {
	var resource LocalGemsRepository
	err := ctx.ReadResource("artifactory:index/localGemsRepository:LocalGemsRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalGemsRepository resources.
type localGemsRepositoryState struct {
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

type LocalGemsRepositoryState struct {
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

func (LocalGemsRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localGemsRepositoryState)(nil)).Elem()
}

type localGemsRepositoryArgs struct {
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

// The set of arguments for constructing a LocalGemsRepository resource.
type LocalGemsRepositoryArgs struct {
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

func (LocalGemsRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localGemsRepositoryArgs)(nil)).Elem()
}

type LocalGemsRepositoryInput interface {
	pulumi.Input

	ToLocalGemsRepositoryOutput() LocalGemsRepositoryOutput
	ToLocalGemsRepositoryOutputWithContext(ctx context.Context) LocalGemsRepositoryOutput
}

func (*LocalGemsRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGemsRepository)(nil)).Elem()
}

func (i *LocalGemsRepository) ToLocalGemsRepositoryOutput() LocalGemsRepositoryOutput {
	return i.ToLocalGemsRepositoryOutputWithContext(context.Background())
}

func (i *LocalGemsRepository) ToLocalGemsRepositoryOutputWithContext(ctx context.Context) LocalGemsRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGemsRepositoryOutput)
}

// LocalGemsRepositoryArrayInput is an input type that accepts LocalGemsRepositoryArray and LocalGemsRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalGemsRepositoryArrayInput` via:
//
//	LocalGemsRepositoryArray{ LocalGemsRepositoryArgs{...} }
type LocalGemsRepositoryArrayInput interface {
	pulumi.Input

	ToLocalGemsRepositoryArrayOutput() LocalGemsRepositoryArrayOutput
	ToLocalGemsRepositoryArrayOutputWithContext(context.Context) LocalGemsRepositoryArrayOutput
}

type LocalGemsRepositoryArray []LocalGemsRepositoryInput

func (LocalGemsRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGemsRepository)(nil)).Elem()
}

func (i LocalGemsRepositoryArray) ToLocalGemsRepositoryArrayOutput() LocalGemsRepositoryArrayOutput {
	return i.ToLocalGemsRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalGemsRepositoryArray) ToLocalGemsRepositoryArrayOutputWithContext(ctx context.Context) LocalGemsRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGemsRepositoryArrayOutput)
}

// LocalGemsRepositoryMapInput is an input type that accepts LocalGemsRepositoryMap and LocalGemsRepositoryMapOutput values.
// You can construct a concrete instance of `LocalGemsRepositoryMapInput` via:
//
//	LocalGemsRepositoryMap{ "key": LocalGemsRepositoryArgs{...} }
type LocalGemsRepositoryMapInput interface {
	pulumi.Input

	ToLocalGemsRepositoryMapOutput() LocalGemsRepositoryMapOutput
	ToLocalGemsRepositoryMapOutputWithContext(context.Context) LocalGemsRepositoryMapOutput
}

type LocalGemsRepositoryMap map[string]LocalGemsRepositoryInput

func (LocalGemsRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGemsRepository)(nil)).Elem()
}

func (i LocalGemsRepositoryMap) ToLocalGemsRepositoryMapOutput() LocalGemsRepositoryMapOutput {
	return i.ToLocalGemsRepositoryMapOutputWithContext(context.Background())
}

func (i LocalGemsRepositoryMap) ToLocalGemsRepositoryMapOutputWithContext(ctx context.Context) LocalGemsRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalGemsRepositoryMapOutput)
}

type LocalGemsRepositoryOutput struct{ *pulumi.OutputState }

func (LocalGemsRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalGemsRepository)(nil)).Elem()
}

func (o LocalGemsRepositoryOutput) ToLocalGemsRepositoryOutput() LocalGemsRepositoryOutput {
	return o
}

func (o LocalGemsRepositoryOutput) ToLocalGemsRepositoryOutputWithContext(ctx context.Context) LocalGemsRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalGemsRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalGemsRepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalGemsRepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o LocalGemsRepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalGemsRepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o LocalGemsRepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o LocalGemsRepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o LocalGemsRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalGemsRepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalGemsRepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o LocalGemsRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalGemsRepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o LocalGemsRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o LocalGemsRepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalGemsRepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *LocalGemsRepository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type LocalGemsRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalGemsRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalGemsRepository)(nil)).Elem()
}

func (o LocalGemsRepositoryArrayOutput) ToLocalGemsRepositoryArrayOutput() LocalGemsRepositoryArrayOutput {
	return o
}

func (o LocalGemsRepositoryArrayOutput) ToLocalGemsRepositoryArrayOutputWithContext(ctx context.Context) LocalGemsRepositoryArrayOutput {
	return o
}

func (o LocalGemsRepositoryArrayOutput) Index(i pulumi.IntInput) LocalGemsRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalGemsRepository {
		return vs[0].([]*LocalGemsRepository)[vs[1].(int)]
	}).(LocalGemsRepositoryOutput)
}

type LocalGemsRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalGemsRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalGemsRepository)(nil)).Elem()
}

func (o LocalGemsRepositoryMapOutput) ToLocalGemsRepositoryMapOutput() LocalGemsRepositoryMapOutput {
	return o
}

func (o LocalGemsRepositoryMapOutput) ToLocalGemsRepositoryMapOutputWithContext(ctx context.Context) LocalGemsRepositoryMapOutput {
	return o
}

func (o LocalGemsRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalGemsRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalGemsRepository {
		return vs[0].(map[string]*LocalGemsRepository)[vs[1].(string)]
	}).(LocalGemsRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGemsRepositoryInput)(nil)).Elem(), &LocalGemsRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGemsRepositoryArrayInput)(nil)).Elem(), LocalGemsRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalGemsRepositoryMapInput)(nil)).Elem(), LocalGemsRepositoryMap{})
	pulumi.RegisterOutputType(LocalGemsRepositoryOutput{})
	pulumi.RegisterOutputType(LocalGemsRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalGemsRepositoryMapOutput{})
}
