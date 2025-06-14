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

// Creates a local Docker v2 repository.
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
//			_, err := artifactory.NewDockerV2Repository(ctx, "foo", &artifactory.DockerV2RepositoryArgs{
//				Key:           pulumi.String("foo"),
//				TagRetention:  pulumi.Int(3),
//				MaxUniqueTags: pulumi.Int(5),
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
// $ pulumi import artifactory:index/dockerV2Repository:DockerV2Repository foo foo
// ```
type DockerV2Repository struct {
	pulumi.CustomResourceState

	// The Docker API version to use.
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolOutput `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest
	// v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolOutput `pulumi:"blockPushingSchema1"`
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
	// The maximum number of unique tags of a single Docker image to store in this
	// repository. Once the number tags for an image exceeds this setting, older tags are removed.
	// A value of 0 (default) indicates there is no limit. This only applies to manifest v2.
	MaxUniqueTags pulumi.IntOutput `pulumi:"maxUniqueTags"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention pulumi.IntOutput `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolOutput `pulumi:"xrayIndex"`
}

// NewDockerV2Repository registers a new resource with the given unique name, arguments, and options.
func NewDockerV2Repository(ctx *pulumi.Context,
	name string, args *DockerV2RepositoryArgs, opts ...pulumi.ResourceOption) (*DockerV2Repository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DockerV2Repository
	err := ctx.RegisterResource("artifactory:index/dockerV2Repository:DockerV2Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerV2Repository gets an existing DockerV2Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerV2Repository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerV2RepositoryState, opts ...pulumi.ResourceOption) (*DockerV2Repository, error) {
	var resource DockerV2Repository
	err := ctx.ReadResource("artifactory:index/dockerV2Repository:DockerV2Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerV2Repository resources.
type dockerV2RepositoryState struct {
	// The Docker API version to use.
	ApiVersion *string `pulumi:"apiVersion"`
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest
	// v2 schema 1 to this repository.
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
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
	// The maximum number of unique tags of a single Docker image to store in this
	// repository. Once the number tags for an image exceeds this setting, older tags are removed.
	// A value of 0 (default) indicates there is no limit. This only applies to manifest v2.
	MaxUniqueTags *int `pulumi:"maxUniqueTags"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type DockerV2RepositoryState struct {
	// The Docker API version to use.
	ApiVersion pulumi.StringPtrInput
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, Artifactory will block the pushing of Docker images with manifest
	// v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
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
	// The maximum number of unique tags of a single Docker image to store in this
	// repository. Once the number tags for an image exceeds this setting, older tags are removed.
	// A value of 0 (default) indicates there is no limit. This only applies to manifest v2.
	MaxUniqueTags pulumi.IntPtrInput
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DockerV2RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV2RepositoryState)(nil)).Elem()
}

type dockerV2RepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled *bool `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut *bool `pulumi:"blackedOut"`
	// When set, Artifactory will block the pushing of Docker images with manifest
	// v2 schema 1 to this repository.
	BlockPushingSchema1 *bool `pulumi:"blockPushingSchema1"`
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
	// The maximum number of unique tags of a single Docker image to store in this
	// repository. Once the number tags for an image exceeds this setting, older tags are removed.
	// A value of 0 (default) indicates there is no limit. This only applies to manifest v2.
	MaxUniqueTags *int `pulumi:"maxUniqueTags"`
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention *int `pulumi:"tagRetention"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a DockerV2Repository resource.
type DockerV2RepositoryArgs struct {
	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrInput
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrInput
	// When set, Artifactory will block the pushing of Docker images with manifest
	// v2 schema 1 to this repository.
	BlockPushingSchema1 pulumi.BoolPtrInput
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
	// The maximum number of unique tags of a single Docker image to store in this
	// repository. Once the number tags for an image exceeds this setting, older tags are removed.
	// A value of 0 (default) indicates there is no limit. This only applies to manifest v2.
	MaxUniqueTags pulumi.IntPtrInput
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
	// If greater than 1, overwritten tags will be saved by their digest, up to the set up
	// number. This only applies to manifest V2.
	TagRetention pulumi.IntPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (DockerV2RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerV2RepositoryArgs)(nil)).Elem()
}

type DockerV2RepositoryInput interface {
	pulumi.Input

	ToDockerV2RepositoryOutput() DockerV2RepositoryOutput
	ToDockerV2RepositoryOutputWithContext(ctx context.Context) DockerV2RepositoryOutput
}

func (*DockerV2Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV2Repository)(nil)).Elem()
}

func (i *DockerV2Repository) ToDockerV2RepositoryOutput() DockerV2RepositoryOutput {
	return i.ToDockerV2RepositoryOutputWithContext(context.Background())
}

func (i *DockerV2Repository) ToDockerV2RepositoryOutputWithContext(ctx context.Context) DockerV2RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV2RepositoryOutput)
}

// DockerV2RepositoryArrayInput is an input type that accepts DockerV2RepositoryArray and DockerV2RepositoryArrayOutput values.
// You can construct a concrete instance of `DockerV2RepositoryArrayInput` via:
//
//	DockerV2RepositoryArray{ DockerV2RepositoryArgs{...} }
type DockerV2RepositoryArrayInput interface {
	pulumi.Input

	ToDockerV2RepositoryArrayOutput() DockerV2RepositoryArrayOutput
	ToDockerV2RepositoryArrayOutputWithContext(context.Context) DockerV2RepositoryArrayOutput
}

type DockerV2RepositoryArray []DockerV2RepositoryInput

func (DockerV2RepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV2Repository)(nil)).Elem()
}

func (i DockerV2RepositoryArray) ToDockerV2RepositoryArrayOutput() DockerV2RepositoryArrayOutput {
	return i.ToDockerV2RepositoryArrayOutputWithContext(context.Background())
}

func (i DockerV2RepositoryArray) ToDockerV2RepositoryArrayOutputWithContext(ctx context.Context) DockerV2RepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV2RepositoryArrayOutput)
}

// DockerV2RepositoryMapInput is an input type that accepts DockerV2RepositoryMap and DockerV2RepositoryMapOutput values.
// You can construct a concrete instance of `DockerV2RepositoryMapInput` via:
//
//	DockerV2RepositoryMap{ "key": DockerV2RepositoryArgs{...} }
type DockerV2RepositoryMapInput interface {
	pulumi.Input

	ToDockerV2RepositoryMapOutput() DockerV2RepositoryMapOutput
	ToDockerV2RepositoryMapOutputWithContext(context.Context) DockerV2RepositoryMapOutput
}

type DockerV2RepositoryMap map[string]DockerV2RepositoryInput

func (DockerV2RepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV2Repository)(nil)).Elem()
}

func (i DockerV2RepositoryMap) ToDockerV2RepositoryMapOutput() DockerV2RepositoryMapOutput {
	return i.ToDockerV2RepositoryMapOutputWithContext(context.Background())
}

func (i DockerV2RepositoryMap) ToDockerV2RepositoryMapOutputWithContext(ctx context.Context) DockerV2RepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerV2RepositoryMapOutput)
}

type DockerV2RepositoryOutput struct{ *pulumi.OutputState }

func (DockerV2RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerV2Repository)(nil)).Elem()
}

func (o DockerV2RepositoryOutput) ToDockerV2RepositoryOutput() DockerV2RepositoryOutput {
	return o
}

func (o DockerV2RepositoryOutput) ToDockerV2RepositoryOutputWithContext(ctx context.Context) DockerV2RepositoryOutput {
	return o
}

// The Docker API version to use.
func (o DockerV2RepositoryOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o DockerV2RepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o DockerV2RepositoryOutput) BlackedOut() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.BlackedOut }).(pulumi.BoolOutput)
}

// When set, Artifactory will block the pushing of Docker images with manifest
// v2 schema 1 to this repository.
func (o DockerV2RepositoryOutput) BlockPushingSchema1() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.BlockPushingSchema1 }).(pulumi.BoolOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o DockerV2RepositoryOutput) CdnRedirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.CdnRedirect }).(pulumi.BoolOutput)
}

// Public description.
func (o DockerV2RepositoryOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o DockerV2RepositoryOutput) DownloadDirect() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.DownloadDirect }).(pulumi.BoolOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o DockerV2RepositoryOutput) ExcludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.ExcludesPattern }).(pulumi.StringOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o DockerV2RepositoryOutput) IncludesPattern() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.IncludesPattern }).(pulumi.StringOutput)
}

// the identity key of the repo.
func (o DockerV2RepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The maximum number of unique tags of a single Docker image to store in this
// repository. Once the number tags for an image exceeds this setting, older tags are removed.
// A value of 0 (default) indicates there is no limit. This only applies to manifest v2.
func (o DockerV2RepositoryOutput) MaxUniqueTags() pulumi.IntOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.IntOutput { return v.MaxUniqueTags }).(pulumi.IntOutput)
}

// Internal description.
func (o DockerV2RepositoryOutput) Notes() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.Notes }).(pulumi.StringOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o DockerV2RepositoryOutput) PriorityResolution() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.PriorityResolution }).(pulumi.BoolOutput)
}

func (o DockerV2RepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o DockerV2RepositoryOutput) ProjectKey() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.ProjectKey }).(pulumi.StringOutput)
}

// List of property set name
func (o DockerV2RepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Sets the layout that the repository should use for storing and identifying modules. A recommended layout that
// corresponds to the package type defined is suggested, and index packages uploaded and calculate metadata accordingly.
func (o DockerV2RepositoryOutput) RepoLayoutRef() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.StringOutput { return v.RepoLayoutRef }).(pulumi.StringOutput)
}

// If greater than 1, overwritten tags will be saved by their digest, up to the set up
// number. This only applies to manifest V2.
func (o DockerV2RepositoryOutput) TagRetention() pulumi.IntOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.IntOutput { return v.TagRetention }).(pulumi.IntOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o DockerV2RepositoryOutput) XrayIndex() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerV2Repository) pulumi.BoolOutput { return v.XrayIndex }).(pulumi.BoolOutput)
}

type DockerV2RepositoryArrayOutput struct{ *pulumi.OutputState }

func (DockerV2RepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerV2Repository)(nil)).Elem()
}

func (o DockerV2RepositoryArrayOutput) ToDockerV2RepositoryArrayOutput() DockerV2RepositoryArrayOutput {
	return o
}

func (o DockerV2RepositoryArrayOutput) ToDockerV2RepositoryArrayOutputWithContext(ctx context.Context) DockerV2RepositoryArrayOutput {
	return o
}

func (o DockerV2RepositoryArrayOutput) Index(i pulumi.IntInput) DockerV2RepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DockerV2Repository {
		return vs[0].([]*DockerV2Repository)[vs[1].(int)]
	}).(DockerV2RepositoryOutput)
}

type DockerV2RepositoryMapOutput struct{ *pulumi.OutputState }

func (DockerV2RepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerV2Repository)(nil)).Elem()
}

func (o DockerV2RepositoryMapOutput) ToDockerV2RepositoryMapOutput() DockerV2RepositoryMapOutput {
	return o
}

func (o DockerV2RepositoryMapOutput) ToDockerV2RepositoryMapOutputWithContext(ctx context.Context) DockerV2RepositoryMapOutput {
	return o
}

func (o DockerV2RepositoryMapOutput) MapIndex(k pulumi.StringInput) DockerV2RepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DockerV2Repository {
		return vs[0].(map[string]*DockerV2Repository)[vs[1].(string)]
	}).(DockerV2RepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV2RepositoryInput)(nil)).Elem(), &DockerV2Repository{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV2RepositoryArrayInput)(nil)).Elem(), DockerV2RepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerV2RepositoryMapInput)(nil)).Elem(), DockerV2RepositoryMap{})
	pulumi.RegisterOutputType(DockerV2RepositoryOutput{})
	pulumi.RegisterOutputType(DockerV2RepositoryArrayOutput{})
	pulumi.RegisterOutputType(DockerV2RepositoryMapOutput{})
}
