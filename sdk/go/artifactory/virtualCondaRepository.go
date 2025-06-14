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

// Creates a virtual Conda repository.
// Official documentation can be found [here](https://www.jfrog.com/confluence/display/JFROG/Conda+Repositories#CondaRepositories-VirtualRepositories).
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
//			_, err := artifactory.NewVirtualCondaRepository(ctx, "foo-conda", &artifactory.VirtualCondaRepositoryArgs{
//				Key:             pulumi.String("foo-conda"),
//				Repositories:    pulumi.StringArray{},
//				Description:     pulumi.String("A test virtual repo"),
//				Notes:           pulumi.String("Internal description"),
//				IncludesPattern: pulumi.String("com/jfrog/**,cloud/jfrog/**"),
//				ExcludesPattern: pulumi.String("com/google/**"),
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
// Virtual repositories can be imported using their name, e.g.
//
// ```sh
// $ pulumi import artifactory:index/virtualCondaRepository:VirtualCondaRepository foo-conda foo-conda
// ```
type VirtualCondaRepository struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
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
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrOutput `pulumi:"retrievalCachePeriodSeconds"`
}

// NewVirtualCondaRepository registers a new resource with the given unique name, arguments, and options.
func NewVirtualCondaRepository(ctx *pulumi.Context,
	name string, args *VirtualCondaRepositoryArgs, opts ...pulumi.ResourceOption) (*VirtualCondaRepository, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VirtualCondaRepository
	err := ctx.RegisterResource("artifactory:index/virtualCondaRepository:VirtualCondaRepository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualCondaRepository gets an existing VirtualCondaRepository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualCondaRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualCondaRepositoryState, opts ...pulumi.ResourceOption) (*VirtualCondaRepository, error) {
	var resource VirtualCondaRepository
	err := ctx.ReadResource("artifactory:index/virtualCondaRepository:VirtualCondaRepository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualCondaRepository resources.
type virtualCondaRepositoryState struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
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
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

type VirtualCondaRepositoryState struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
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
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualCondaRepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualCondaRepositoryState)(nil)).Elem()
}

type virtualCondaRepositoryArgs struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
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
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds *int `pulumi:"retrievalCachePeriodSeconds"`
}

// The set of arguments for constructing a VirtualCondaRepository resource.
type VirtualCondaRepositoryArgs struct {
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
	// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
	// contain spaces or special characters.
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
	// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
	RetrievalCachePeriodSeconds pulumi.IntPtrInput
}

func (VirtualCondaRepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualCondaRepositoryArgs)(nil)).Elem()
}

type VirtualCondaRepositoryInput interface {
	pulumi.Input

	ToVirtualCondaRepositoryOutput() VirtualCondaRepositoryOutput
	ToVirtualCondaRepositoryOutputWithContext(ctx context.Context) VirtualCondaRepositoryOutput
}

func (*VirtualCondaRepository) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCondaRepository)(nil)).Elem()
}

func (i *VirtualCondaRepository) ToVirtualCondaRepositoryOutput() VirtualCondaRepositoryOutput {
	return i.ToVirtualCondaRepositoryOutputWithContext(context.Background())
}

func (i *VirtualCondaRepository) ToVirtualCondaRepositoryOutputWithContext(ctx context.Context) VirtualCondaRepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCondaRepositoryOutput)
}

// VirtualCondaRepositoryArrayInput is an input type that accepts VirtualCondaRepositoryArray and VirtualCondaRepositoryArrayOutput values.
// You can construct a concrete instance of `VirtualCondaRepositoryArrayInput` via:
//
//	VirtualCondaRepositoryArray{ VirtualCondaRepositoryArgs{...} }
type VirtualCondaRepositoryArrayInput interface {
	pulumi.Input

	ToVirtualCondaRepositoryArrayOutput() VirtualCondaRepositoryArrayOutput
	ToVirtualCondaRepositoryArrayOutputWithContext(context.Context) VirtualCondaRepositoryArrayOutput
}

type VirtualCondaRepositoryArray []VirtualCondaRepositoryInput

func (VirtualCondaRepositoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualCondaRepository)(nil)).Elem()
}

func (i VirtualCondaRepositoryArray) ToVirtualCondaRepositoryArrayOutput() VirtualCondaRepositoryArrayOutput {
	return i.ToVirtualCondaRepositoryArrayOutputWithContext(context.Background())
}

func (i VirtualCondaRepositoryArray) ToVirtualCondaRepositoryArrayOutputWithContext(ctx context.Context) VirtualCondaRepositoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCondaRepositoryArrayOutput)
}

// VirtualCondaRepositoryMapInput is an input type that accepts VirtualCondaRepositoryMap and VirtualCondaRepositoryMapOutput values.
// You can construct a concrete instance of `VirtualCondaRepositoryMapInput` via:
//
//	VirtualCondaRepositoryMap{ "key": VirtualCondaRepositoryArgs{...} }
type VirtualCondaRepositoryMapInput interface {
	pulumi.Input

	ToVirtualCondaRepositoryMapOutput() VirtualCondaRepositoryMapOutput
	ToVirtualCondaRepositoryMapOutputWithContext(context.Context) VirtualCondaRepositoryMapOutput
}

type VirtualCondaRepositoryMap map[string]VirtualCondaRepositoryInput

func (VirtualCondaRepositoryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualCondaRepository)(nil)).Elem()
}

func (i VirtualCondaRepositoryMap) ToVirtualCondaRepositoryMapOutput() VirtualCondaRepositoryMapOutput {
	return i.ToVirtualCondaRepositoryMapOutputWithContext(context.Background())
}

func (i VirtualCondaRepositoryMap) ToVirtualCondaRepositoryMapOutputWithContext(ctx context.Context) VirtualCondaRepositoryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualCondaRepositoryMapOutput)
}

type VirtualCondaRepositoryOutput struct{ *pulumi.OutputState }

func (VirtualCondaRepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualCondaRepository)(nil)).Elem()
}

func (o VirtualCondaRepositoryOutput) ToVirtualCondaRepositoryOutput() VirtualCondaRepositoryOutput {
	return o
}

func (o VirtualCondaRepositoryOutput) ToVirtualCondaRepositoryOutputWithContext(ctx context.Context) VirtualCondaRepositoryOutput {
	return o
}

// Whether the virtual repository should search through remote repositories when trying to resolve an artifact requested by
// another Artifactory instance.
func (o VirtualCondaRepositoryOutput) ArtifactoryRequestsCanRetrieveRemoteArtifacts() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.BoolPtrOutput {
		return v.ArtifactoryRequestsCanRetrieveRemoteArtifacts
	}).(pulumi.BoolPtrOutput)
}

// Default repository to deploy artifacts.
func (o VirtualCondaRepositoryOutput) DefaultDeploymentRepo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.DefaultDeploymentRepo }).(pulumi.StringPtrOutput)
}

// Public description.
func (o VirtualCondaRepositoryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// List of artifact patterns to exclude when evaluating artifact requests, in the form of `x/y/**/z/*`.By default no
// artifacts are excluded.
func (o VirtualCondaRepositoryOutput) ExcludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.ExcludesPattern }).(pulumi.StringPtrOutput)
}

// List of comma-separated artifact patterns to include when evaluating artifact requests in the form of `x/y/**/z/*`. When
// used, only artifacts matching one of the include patterns are served. By default, all artifacts are included (`**/*`).
func (o VirtualCondaRepositoryOutput) IncludesPattern() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.IncludesPattern }).(pulumi.StringPtrOutput)
}

// A mandatory identifier for the repository that must be unique. It cannot begin with a number or
// contain spaces or special characters.
func (o VirtualCondaRepositoryOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// Internal description.
func (o VirtualCondaRepositoryOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o VirtualCondaRepositoryOutput) PackageType() pulumi.StringOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringOutput { return v.PackageType }).(pulumi.StringOutput)
}

func (o VirtualCondaRepositoryOutput) ProjectEnvironments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringArrayOutput { return v.ProjectEnvironments }).(pulumi.StringArrayOutput)
}

// Project key for assigning this repository to. Must be 2 - 32 lowercase alphanumeric and hyphen characters. When
// assigning repository to a project, repository key must be prefixed with project key, separated by a dash.
func (o VirtualCondaRepositoryOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Repository layout key for the virtual repository
func (o VirtualCondaRepositoryOutput) RepoLayoutRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringPtrOutput { return v.RepoLayoutRef }).(pulumi.StringPtrOutput)
}

// The effective list of actual repositories included in this virtual repository.
func (o VirtualCondaRepositoryOutput) Repositories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.StringArrayOutput { return v.Repositories }).(pulumi.StringArrayOutput)
}

// This value refers to the number of seconds to cache metadata files before checking for newer versions on aggregated repositories. A value of 0 indicates no caching.
func (o VirtualCondaRepositoryOutput) RetrievalCachePeriodSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *VirtualCondaRepository) pulumi.IntPtrOutput { return v.RetrievalCachePeriodSeconds }).(pulumi.IntPtrOutput)
}

type VirtualCondaRepositoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualCondaRepositoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VirtualCondaRepository)(nil)).Elem()
}

func (o VirtualCondaRepositoryArrayOutput) ToVirtualCondaRepositoryArrayOutput() VirtualCondaRepositoryArrayOutput {
	return o
}

func (o VirtualCondaRepositoryArrayOutput) ToVirtualCondaRepositoryArrayOutputWithContext(ctx context.Context) VirtualCondaRepositoryArrayOutput {
	return o
}

func (o VirtualCondaRepositoryArrayOutput) Index(i pulumi.IntInput) VirtualCondaRepositoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VirtualCondaRepository {
		return vs[0].([]*VirtualCondaRepository)[vs[1].(int)]
	}).(VirtualCondaRepositoryOutput)
}

type VirtualCondaRepositoryMapOutput struct{ *pulumi.OutputState }

func (VirtualCondaRepositoryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VirtualCondaRepository)(nil)).Elem()
}

func (o VirtualCondaRepositoryMapOutput) ToVirtualCondaRepositoryMapOutput() VirtualCondaRepositoryMapOutput {
	return o
}

func (o VirtualCondaRepositoryMapOutput) ToVirtualCondaRepositoryMapOutputWithContext(ctx context.Context) VirtualCondaRepositoryMapOutput {
	return o
}

func (o VirtualCondaRepositoryMapOutput) MapIndex(k pulumi.StringInput) VirtualCondaRepositoryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VirtualCondaRepository {
		return vs[0].(map[string]*VirtualCondaRepository)[vs[1].(string)]
	}).(VirtualCondaRepositoryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualCondaRepositoryInput)(nil)).Elem(), &VirtualCondaRepository{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualCondaRepositoryArrayInput)(nil)).Elem(), VirtualCondaRepositoryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VirtualCondaRepositoryMapInput)(nil)).Elem(), VirtualCondaRepositoryMap{})
	pulumi.RegisterOutputType(VirtualCondaRepositoryOutput{})
	pulumi.RegisterOutputType(VirtualCondaRepositoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualCondaRepositoryMapOutput{})
}
