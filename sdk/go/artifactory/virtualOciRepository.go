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

// Creates a virtual OCI repository.
//
// Official documentation can be found [here](https://jfrog.com/help/r/jfrog-artifactory-documentation/set-up-virtual-oci-repositories).
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
//			_, err := artifactory.NewVirtualOciRepository(ctx, "my-oci-virtual", &artifactory.VirtualOciRepositoryArgs{
//				Key: pulumi.String("my-oci-virtual"),
//				Repositories: pulumi.StringArray{
//					pulumi.String("my-oci-local"),
//					pulumi.String("my-oci-remote"),
//				},
//				Description:               pulumi.String("A test virtual OCI repo"),
//				Notes:                     pulumi.String("Internal description"),
//				ResolveOciTagsByTimestamp: pulumi.Bool(true),
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
// Virtual OCI repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/virtualOciRepository:VirtualOciRepository my-oci-virtual my-oci-virtual
// ```
type VirtualOciRepository struct {
	pulumi.CustomResourceState

	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrOutput `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrOutput `pulumi:"defaultDeploymentRepo"`
	// Public description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrOutput `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrOutput `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringOutput `pulumi:"key"`
	// Internal description.
	Notes               pulumi.StringPtrOutput   `pulumi:"notes"`
	PackageType         pulumi.StringOutput      `pulumi:"packageType"`
	ProjectEnvironments pulumi.StringArrayOutput `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrOutput `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayOutput `pulumi:"repositories"`
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrOutput `pulumi:"resolveOciTagsByTimestamp"`
}

// NewVirtualOciRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualOciRepository(ctx *pulumi.Context,
	name string, args *VirtualOciRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualOciRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualOciRepository
	err := ctx.RegisterResource("artifactory:index/virtualOciRepository:VirtualOciRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualOciRepository gets an existing VirtualOciRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualOciRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualOciRepositoryState, opts ...pulumi.ResourceOption) (*VirtualOciRepository, error) {
	var resource VirtualOciRepository
	err := ctx.ReadResource("artifactory:index/virtualOciRepository:VirtualOciRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualOciRepository resources.
type virtualOciRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// Public description.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key *string `pulumi:"key"`
	// Internal description.
	Notes               *string  `pulumi:"notes"`
	PackageType         *string  `pulumi:"packageType"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

type VirtualOciRepositoryState struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringPtrInput
	// Internal description.
	Notes               pulumi.StringPtrInput
	PackageType         pulumi.StringPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrInput
}

func (VirtualOciRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualOciRepositoryState)(nil)).Elem()
}

type virtualOciRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts *bool `pulumi:"artifactoryRequestsCanRetrieveRemoteArtifacts"`
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo *string `pulumi:"defaultDeploymentRepo"`
	// Public description.
	Description *string `pulumi:"description"`
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern *string `pulumi:"excludesPattern"`
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern *string `pulumi:"includesPattern"`
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key string `pulumi:"key"`
	// Internal description.
	Notes               *string  `pulumi:"notes"`
	ProjectEnvironments []string `pulumi:"projectEnvironments"`
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey *string `pulumi:"projectKey"`
	// Repository layout key for the virtual repository
	RepoLayoutRef *string `pulumi:"repoLayoutRef"`
	// The effective list of actual repositories included in this virtual repository.
	Repositories []string `pulumi:"repositories"`
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp *bool `pulumi:"resolveOciTagsByTimestamp"`
}

// The set of arguments for constructing a VirtualOciRepository resource.
type VirtualOciRepositoryArgs struct {
	// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
	// another Artifactory instance.
	ArtifactoryRequestsCanRetrieveRemoteArtifacts pulumi.BoolPtrInput
	// Default repository to deploy artifacts.
	DefaultDeploymentRepo pulumi.StringPtrInput
	// Public description.
	Description pulumi.StringPtrInput
	// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
	// artifacts are excluded.
	ExcludesPattern pulumi.StringPtrInput
	// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
	// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
	IncludesPattern pulumi.StringPtrInput
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
	Key pulumi.StringInput
	// Internal description.
	Notes               pulumi.StringPtrInput
	ProjectEnvironments pulumi.StringArrayInput
	// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
	// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
	ProjectKey pulumi.StringPtrInput
	// Repository layout key for the virtual repository
	RepoLayoutRef pulumi.StringPtrInput
	// The effective list of actual repositories included in this virtual repository.
	Repositories pulumi.StringArrayInput
	// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
	ResolveOciTagsByTimestamp pulumi.BoolPtrInput
}

func (VirtualOciRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualOciRepositoryArgs)(nil)).Elem()
}

type VirtualOciRepositoryInput interface {
	pulumi.Input

	ToVirtualOciRepositoryOutput() VirtualOciRepositoryOutput
	ToVirtualOciRepositoryOutputWithContext(ctx context.Context) VirtualOciRepositoryOutput
}

func (*VirtualOciRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualOciRepository)(nil)).Elem()
}

func (i *VirtualOciRepository) ToVirtualOciRepositoryOutput() VirtualOciRepositoryOutput {
	return i.ToVirtualOciRepositoryOutputWithContext(context.Background())
}

func (i *VirtualOciRepository) ToVirtualOciRepositoryOutputWithContext(ctx context.Context) VirtualOciRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualOciRepositoryOutput)
}

// VirtualOciRepositoryArrayInput is an input type that accepts VirtualOciRepositoryArray and VirtualOciRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualOciRepositoryArrayInput` via:
//
//	VirtualOciRepositoryArray{ VirtualOciRepositoryArgs{...} }
type VirtualOciRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualOciRepositoryArrayOutput() VirtualOciRepositoryArrayOutput
	ToVirtualOciRepositoryArrayOutputWithContext(context.Context) VirtualOciRepositoryArrayOutput
}

type VirtualOciRepositoryArray []VirtualOciRepositoryInput

func (VirtualOciRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualOciRepository)(nil)).Elem()
}

func (i VirtualOciRepositoryArray) ToVirtualOciRepositoryArrayOutput() VirtualOciRepositoryArrayOutput {
	return i.ToVirtualOciRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualOciRepositoryArray) ToVirtualOciRepositoryArrayOutputWithContext(ctx context.Context) VirtualOciRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualOciRepositoryArrayOutput)
}

// VirtualOciRepositoryMapInput is an input type that accepts VirtualOciRepositoryMap and VirtualOciRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualOciRepositoryMapInput` via:
//
//	VirtualOciRepositoryMap{ "key": VirtualOciRepositoryArgs{...} }
type VirtualOciRepositoryMapInput interface {
	pulumi.Input

	ToVirtualOciRepositoryMapOutput() VirtualOciRepositoryMapOutput
	ToVirtualOciRepositoryMapOutputWithContext(context.Context) VirtualOciRepositoryMapOutput
}

type VirtualOciRepositoryMap map[string]VirtualOciRepositoryInput

func (VirtualOciRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualOciRepository)(nil)).Elem()
}

func (i VirtualOciRepositoryMap) ToVirtualOciRepositoryMapOutput() VirtualOciRepositoryMapOutput {
	return i.ToVirtualOciRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualOciRepositoryMap) ToVirtualOciRepositoryMapOutputWithContext(ctx context.Context) VirtualOciRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualOciRepositoryMapOutput)
}

type VirtualOciRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualOciRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualOciRepository)(nil)).Elem()
}

func (o VirtualOciRepositoryOutput) ToVirtualOciRepositoryOutput() VirtualOciRepositoryOutput {
	return o
}

func (o VirtualOciRepositoryOutput) ToVirtualOciRepositoryOutputWithContext(ctx context.Context) VirtualOciRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualOciRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualOciRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualOciRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o VirtualOciRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o VirtualOciRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or contain spaces or special characters.
func (o VirtualOciRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualOciRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualOciRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

func (o VirtualOciRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualOciRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualOciRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualOciRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// When enabled, in cases where the same OCI tag exists in two or more of the aggregated repositories, Artifactory will return the tag that has the latest timestamp. Default values is `false`.
func (o VirtualOciRepositoryOutput) ResolveOciTagsByTimestamp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualOciRepository) pulumi.BoolPtrOutput { return v.ResolveOciTagsByTimestamp }).(pulumi.BoolPtrOutput)
}

type VirtualOciRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualOciRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualOciRepository)(nil)).Elem()
}

func (o VirtualOciRepositoryArrayOutput) ToVirtualOciRepositoryArrayOutput() VirtualOciRepositoryArrayOutput {
	return o
}

func (o VirtualOciRepositoryArrayOutput) ToVirtualOciRepositoryArrayOutputWithContext(ctx context.Context) VirtualOciRepositoryArrayOutput {
	return o
}

func (o VirtualOciRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualOciRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualOciRepository {
		return vs[0].([]*VirtualOciRepository)[vs[1].(int)]
	}).(VirtualOciRepositoryOutput)
}

type VirtualOciRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualOciRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualOciRepository)(nil)).Elem()
}

func (o VirtualOciRepositoryMapOutput) ToVirtualOciRepositoryMapOutput() VirtualOciRepositoryMapOutput {
	return o
}

func (o VirtualOciRepositoryMapOutput) ToVirtualOciRepositoryMapOutputWithContext(ctx context.Context) VirtualOciRepositoryMapOutput {
	return o
}

func (o VirtualOciRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualOciRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualOciRepository {
		return vs[0].(map[string]*VirtualOciRepository)[vs[1].(string)]
	}).(VirtualOciRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualOciRepositoryInput)(nil)).Elem(), &VirtualOciRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualOciRepositoryArrayInput)(nil)).Elem(), VirtualOciRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualOciRepositoryMapInput)(nil)).Elem(), VirtualOciRepositoryMap{})
	pulumi.RegisterOutputType(VirtualOciRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualOciRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualOciRepositoryMapOutput{})
}
