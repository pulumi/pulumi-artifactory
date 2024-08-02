// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a local Ansible repository.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory"
//	"github.com/pulumi/pulumi-std/sdk/go/std"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			invokeFile, err := std.File(ctx, &std.FileArgs{
//				Input: "samples/rsa.priv",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			invokeFile1, err := std.File(ctx, &std.FileArgs{
//				Input: "samples/rsa.pub",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewKeypair(ctx, "some-keypair-RSA", &artifactory.KeypairArgs{
//				PairName:   pulumi.String("some-keypair"),
//				PairType:   pulumi.String("RSA"),
//				Alias:      pulumi.String("foo-alias"),
//				PrivateKey: pulumi.String(invokeFile.Result),
//				PublicKey:  pulumi.String(invokeFile1.Result),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewLocalAnsibleRepository(ctx, "test-ansible-local-repo", &artifactory.LocalAnsibleRepositoryArgs{
//				Key:               pulumi.String("test-ansible-local-repo"),
//				PrimaryKeypairRef: some_keypair_RSA.PairName,
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
// $ pulumi import artifactory:index/localAnsibleRepository:LocalAnsibleRepository test-ansible-local-repo test-ansible-local-repo
// ```
type LocalAnsibleRepository struct {
	pulumi.CustomResourceState

	// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
	// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
	// security (e.g., cross-site scripting attacks).
	ArchiveBrowsingEnabled pulumi.BoolPtrOutput `pulumi:"archiveBrowsingEnabled"`
	// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
	BlackedOut pulumi.BoolPtrOutput `pulumi:"blackedOut"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
	// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
	CdnRedirect pulumi.BoolPtrOutput `pulumi:"cdnRedirect"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
	// storage provider. Available in Enterprise+ and Edge licenses only.
	DownloadDirect pulumi.BoolPtrOutput `pulumi:"downloadDirect"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes       pulumi.StringPtrOutput `pulumi:"notes"`
	PackageType pulumi.StringOutput    `pulumi:"packageType"`
	// Used to sign index files in Alpine Linux repositories. See:
	// https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
	PrimaryKeypairRef pulumi.StringPtrOutput `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrOutput     `pulumi:"priorityResolution"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// List of property set name
	PropertySets pulumi.StringArrayOutput `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrOutput `pulumi:"xrayIndex"`
}

// NewLocalAnsibleRepository registers a new resource with the given unique name, arguments, and options.
func NewLocalAnsibleRepository(ctx *pulumi.Context,
	name string, args *LocalAnsibleRepositoryArgs, opts ...pulumi.ResourceOption) (*LocalAnsibleRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LocalAnsibleRepository
	err := ctx.RegisterResource("artifactory:index/localAnsibleRepository:LocalAnsibleRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalAnsibleRepository gets an existing LocalAnsibleRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalAnsibleRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalAnsibleRepositoryState, opts ...pulumi.ResourceOption) (*LocalAnsibleRepository, error) {
	var resource LocalAnsibleRepository
	err := ctx.ReadResource("artifactory:index/localAnsibleRepository:LocalAnsibleRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalAnsibleRepository resources.
type localAnsibleRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes       *string `pulumi:"notes"`
	PackageType *string `pulumi:"packageType"`
	// Used to sign index files in Alpine Linux repositories. See:
	// https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

type LocalAnsibleRepositoryState struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes       pulumi.StringPtrInput
	PackageType pulumi.StringPtrInput
	// Used to sign index files in Alpine Linux repositories. See:
	// https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalAnsibleRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*localAnsibleRepositoryState)(nil)).Elem()
}

type localAnsibleRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern *string `pulumi:"includesPattern"`
	// the identity key of the repo.
	Key string `pulumi:"key"`
	// Internal description.
	Notes *string `pulumi:"notes"`
	// Used to sign index files in Alpine Linux repositories. See:
	// https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
	PrimaryKeypairRef *string `pulumi:"primaryKeypairRef"`
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  *bool    `pulumi:"priorityResolution"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// List of property set name
	PropertySets []string `pulumi:"propertySets"`
	// Repository layout key for the local repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex *bool `pulumi:"xrayIndex"`
}

// The set of arguments for constructing a LocalAnsibleRepository resource.
type LocalAnsibleRepositoryArgs struct {
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
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
	IncludesPattern pulumi.StringPtrInput
	// the identity key of the repo.
	Key pulumi.StringInput
	// Internal description.
	Notes pulumi.StringPtrInput
	// Used to sign index files in Alpine Linux repositories. See:
	// https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
	PrimaryKeypairRef pulumi.StringPtrInput
	// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
	PriorityResolution  pulumi.BoolPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// List of property set name
	PropertySets pulumi.StringArrayInput
	// Repository layout key for the local repository
	RepoLayoutRef pulumi.StringPtrInput
	// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
	// Xray settings.
	XrayIndex pulumi.BoolPtrInput
}

func (LocalAnsibleRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localAnsibleRepositoryArgs)(nil)).Elem()
}

type LocalAnsibleRepositoryInput interface {
	pulumi.Input

	ToLocalAnsibleRepositoryOutput() LocalAnsibleRepositoryOutput
	ToLocalAnsibleRepositoryOutputWithContext(ctx context.Context) LocalAnsibleRepositoryOutput
}

func (*LocalAnsibleRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalAnsibleRepository)(nil)).Elem()
}

func (i *LocalAnsibleRepository) ToLocalAnsibleRepositoryOutput() LocalAnsibleRepositoryOutput {
	return i.ToLocalAnsibleRepositoryOutputWithContext(context.Background())
}

func (i *LocalAnsibleRepository) ToLocalAnsibleRepositoryOutputWithContext(ctx context.Context) LocalAnsibleRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalAnsibleRepositoryOutput)
}

// LocalAnsibleRepositoryArrayInput is an input type that accepts LocalAnsibleRepositoryArray and LocalAnsibleRepositoryArrayOutput values.
// You can construct a concrete instance of `LocalAnsibleRepositoryArrayInput` via:
//
//	LocalAnsibleRepositoryArray{ LocalAnsibleRepositoryArgs{...} }
type LocalAnsibleRepositoryArrayInput interface {
	pulumi.Input

	ToLocalAnsibleRepositoryArrayOutput() LocalAnsibleRepositoryArrayOutput
	ToLocalAnsibleRepositoryArrayOutputWithContext(context.Context) LocalAnsibleRepositoryArrayOutput
}

type LocalAnsibleRepositoryArray []LocalAnsibleRepositoryInput

func (LocalAnsibleRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalAnsibleRepository)(nil)).Elem()
}

func (i LocalAnsibleRepositoryArray) ToLocalAnsibleRepositoryArrayOutput() LocalAnsibleRepositoryArrayOutput {
	return i.ToLocalAnsibleRepositoryArrayOutputWithContext(context.Background())
}

func (i LocalAnsibleRepositoryArray) ToLocalAnsibleRepositoryArrayOutputWithContext(ctx context.Context) LocalAnsibleRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalAnsibleRepositoryArrayOutput)
}

// LocalAnsibleRepositoryMapInput is an input type that accepts LocalAnsibleRepositoryMap and LocalAnsibleRepositoryMapOutput values.
// You can construct a concrete instance of `LocalAnsibleRepositoryMapInput` via:
//
//	LocalAnsibleRepositoryMap{ "key": LocalAnsibleRepositoryArgs{...} }
type LocalAnsibleRepositoryMapInput interface {
	pulumi.Input

	ToLocalAnsibleRepositoryMapOutput() LocalAnsibleRepositoryMapOutput
	ToLocalAnsibleRepositoryMapOutputWithContext(context.Context) LocalAnsibleRepositoryMapOutput
}

type LocalAnsibleRepositoryMap map[string]LocalAnsibleRepositoryInput

func (LocalAnsibleRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalAnsibleRepository)(nil)).Elem()
}

func (i LocalAnsibleRepositoryMap) ToLocalAnsibleRepositoryMapOutput() LocalAnsibleRepositoryMapOutput {
	return i.ToLocalAnsibleRepositoryMapOutputWithContext(context.Background())
}

func (i LocalAnsibleRepositoryMap) ToLocalAnsibleRepositoryMapOutputWithContext(ctx context.Context) LocalAnsibleRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocalAnsibleRepositoryMapOutput)
}

type LocalAnsibleRepositoryOutput struct{ *pulumi.OutputState }

func (LocalAnsibleRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocalAnsibleRepository)(nil)).Elem()
}

func (o LocalAnsibleRepositoryOutput) ToLocalAnsibleRepositoryOutput() LocalAnsibleRepositoryOutput {
	return o
}

func (o LocalAnsibleRepositoryOutput) ToLocalAnsibleRepositoryOutputWithContext(ctx context.Context) LocalAnsibleRepositoryOutput {
	return o
}

// When set, you may view content such as HTML or Javadoc files directly from Artifactory. This may not be safe and
// therefore requires strict content moderation to prevent malicious users from uploading content that may compromise
// security (e.g., cross-site scripting attacks).
func (o LocalAnsibleRepositoryOutput) ArchiveBrowsingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.BoolPtrOutput { return v.ArchiveBrowsingEnabled }).(pulumi.BoolPtrOutput)
}

// When set, the repository does not participate in artifact resolution and new artifacts cannot be deployed.
func (o LocalAnsibleRepositoryOutput) BlackedOut() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.BoolPtrOutput { return v.BlackedOut }).(pulumi.BoolPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from AWS
// CloudFront. Available in Enterprise+ and Edge licenses only. Default value is 'false'
func (o LocalAnsibleRepositoryOutput) CdnRedirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.BoolPtrOutput { return v.CdnRedirect }).(pulumi.BoolPtrOutput)
}

// Public description.
func (o LocalAnsibleRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// When set, download requests to this repository will redirect the client to download the artifact directly from the cloud
// storage provider. Available in Enterprise+ and Edge licenses only.
func (o LocalAnsibleRepositoryOutput) DownloadDirect() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.BoolPtrOutput { return v.DownloadDirect }).(pulumi.BoolPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of x/y/**/z/*.By default no
// artifacts are excluded.
func (o LocalAnsibleRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of x/y/**/z/*. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (**/*).
func (o LocalAnsibleRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// the identity key of the repo.
func (o LocalAnsibleRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o LocalAnsibleRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o LocalAnsibleRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

// Used to sign index files in Alpine Linux repositories. See:
// https://www.jfrog.com/confluence/display/JFROG/Alpine+Linux+Repositories#AlpineLinuxRepositories-SigningAlpineLinuxIndex
func (o LocalAnsibleRepositoryOutput) PrimaryKeypairRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.PrimaryKeypairRef }).(pulumi.StringPtrOutput)
}

// Setting repositories with priority will cause metadata to be merged only from repositories set with this field
func (o LocalAnsibleRepositoryOutput) PriorityResolution() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.BoolPtrOutput { return v.PriorityResolution }).(pulumi.BoolPtrOutput)
}

func (o LocalAnsibleRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o LocalAnsibleRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// List of property set name
func (o LocalAnsibleRepositoryOutput) PropertySets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringArrayOutput { return v.PropertySets }).(pulumi.StringArrayOutput)
}

// Repository layout key for the local repository
func (o LocalAnsibleRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// Enable Indexing In Xray. Repository will be indexed with the default retention period. You will be able to change it via
// Xray settings.
func (o LocalAnsibleRepositoryOutput) XrayIndex() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LocalAnsibleRepository) pulumi.BoolPtrOutput { return v.XrayIndex }).(pulumi.BoolPtrOutput)
}

type LocalAnsibleRepositoryArrayOutput struct{ *pulumi.OutputState }

func (LocalAnsibleRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LocalAnsibleRepository)(nil)).Elem()
}

func (o LocalAnsibleRepositoryArrayOutput) ToLocalAnsibleRepositoryArrayOutput() LocalAnsibleRepositoryArrayOutput {
	return o
}

func (o LocalAnsibleRepositoryArrayOutput) ToLocalAnsibleRepositoryArrayOutputWithContext(ctx context.Context) LocalAnsibleRepositoryArrayOutput {
	return o
}

func (o LocalAnsibleRepositoryArrayOutput) Index(i pulumi.IntInput) LocalAnsibleRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LocalAnsibleRepository {
		return vs[0].([]*LocalAnsibleRepository)[vs[1].(int)]
	}).(LocalAnsibleRepositoryOutput)
}

type LocalAnsibleRepositoryMapOutput struct{ *pulumi.OutputState }

func (LocalAnsibleRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LocalAnsibleRepository)(nil)).Elem()
}

func (o LocalAnsibleRepositoryMapOutput) ToLocalAnsibleRepositoryMapOutput() LocalAnsibleRepositoryMapOutput {
	return o
}

func (o LocalAnsibleRepositoryMapOutput) ToLocalAnsibleRepositoryMapOutputWithContext(ctx context.Context) LocalAnsibleRepositoryMapOutput {
	return o
}

func (o LocalAnsibleRepositoryMapOutput) MapIndex(k pulumi.StringInput) LocalAnsibleRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LocalAnsibleRepository {
		return vs[0].(map[string]*LocalAnsibleRepository)[vs[1].(string)]
	}).(LocalAnsibleRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LocalAnsibleRepositoryInput)(nil)).Elem(), &LocalAnsibleRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalAnsibleRepositoryArrayInput)(nil)).Elem(), LocalAnsibleRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LocalAnsibleRepositoryMapInput)(nil)).Elem(), LocalAnsibleRepositoryMap{})
	pulumi.RegisterOutputType(LocalAnsibleRepositoryOutput{})
	pulumi.RegisterOutputType(LocalAnsibleRepositoryArrayOutput{})
	pulumi.RegisterOutputType(LocalAnsibleRepositoryMapOutput{})
}