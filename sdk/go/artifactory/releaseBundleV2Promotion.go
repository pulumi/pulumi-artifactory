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

// This resource enables you to promote Release Bundle V2 version. For more information, see [JFrog documentation](https://jfrog.com/help/r/jfrog-artifactory-documentation/promote-a-release-bundle-v2-to-a-target-environment).
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
//			_, err := artifactory.NewReleaseBundleV2Promotion(ctx, "my-release-bundle-v2-promotion", &artifactory.ReleaseBundleV2PromotionArgs{
//				Name:        pulumi.String("my-release-bundle-v2-artifacts"),
//				Version:     pulumi.String("1.0.0"),
//				KeypairName: pulumi.String("my-keypair-name"),
//				Environment: pulumi.String("DEV"),
//				IncludedRepositoryKeys: pulumi.StringArray{
//					pulumi.String("commons-qa-maven-local"),
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
type ReleaseBundleV2Promotion struct {
	pulumi.CustomResourceState

	// Timestamp when the new version was created (ISO 8601 standard).
	Created pulumi.StringOutput `pulumi:"created"`
	// Timestamp when the new version was created (in milliseconds).
	CreatedMillis pulumi.IntOutput `pulumi:"createdMillis"`
	// Target environment
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Defines specific repositories to exclude from the promotion.
	ExcludedRepositoryKeys pulumi.StringArrayOutput `pulumi:"excludedRepositoryKeys"`
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	IncludedRepositoryKeys pulumi.StringArrayOutput `pulumi:"includedRepositoryKeys"`
	// Key-pair name to use for signature creation
	KeypairName pulumi.StringOutput `pulumi:"keypairName"`
	// Name of Release Bundle
	Name pulumi.StringOutput `pulumi:"name"`
	// Project key the Release Bundle belongs to
	ProjectKey pulumi.StringPtrOutput `pulumi:"projectKey"`
	// Version to promote
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewReleaseBundleV2Promotion registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleV2Promotion(ctx *pulumi.Context,
	name string, args *ReleaseBundleV2PromotionArgs, opts ...pulumi.ResourceOption) (*ReleaseBundleV2Promotion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	if args.KeypairName == nil {
		return nil, errors.New("invalid value for required argument 'KeypairName'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReleaseBundleV2Promotion
	err := ctx.RegisterResource("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseBundleV2Promotion gets an existing ReleaseBundleV2Promotion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseBundleV2Promotion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseBundleV2PromotionState, opts ...pulumi.ResourceOption) (*ReleaseBundleV2Promotion, error) {
	var resource ReleaseBundleV2Promotion
	err := ctx.ReadResource("artifactory:index/releaseBundleV2Promotion:ReleaseBundleV2Promotion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseBundleV2Promotion resources.
type releaseBundleV2PromotionState struct {
	// Timestamp when the new version was created (ISO 8601 standard).
	Created *string `pulumi:"created"`
	// Timestamp when the new version was created (in milliseconds).
	CreatedMillis *int `pulumi:"createdMillis"`
	// Target environment
	Environment *string `pulumi:"environment"`
	// Defines specific repositories to exclude from the promotion.
	ExcludedRepositoryKeys []string `pulumi:"excludedRepositoryKeys"`
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	IncludedRepositoryKeys []string `pulumi:"includedRepositoryKeys"`
	// Key-pair name to use for signature creation
	KeypairName *string `pulumi:"keypairName"`
	// Name of Release Bundle
	Name *string `pulumi:"name"`
	// Project key the Release Bundle belongs to
	ProjectKey *string `pulumi:"projectKey"`
	// Version to promote
	Version *string `pulumi:"version"`
}

type ReleaseBundleV2PromotionState struct {
	// Timestamp when the new version was created (ISO 8601 standard).
	Created pulumi.StringPtrInput
	// Timestamp when the new version was created (in milliseconds).
	CreatedMillis pulumi.IntPtrInput
	// Target environment
	Environment pulumi.StringPtrInput
	// Defines specific repositories to exclude from the promotion.
	ExcludedRepositoryKeys pulumi.StringArrayInput
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	IncludedRepositoryKeys pulumi.StringArrayInput
	// Key-pair name to use for signature creation
	KeypairName pulumi.StringPtrInput
	// Name of Release Bundle
	Name pulumi.StringPtrInput
	// Project key the Release Bundle belongs to
	ProjectKey pulumi.StringPtrInput
	// Version to promote
	Version pulumi.StringPtrInput
}

func (ReleaseBundleV2PromotionState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2PromotionState)(nil)).Elem()
}

type releaseBundleV2PromotionArgs struct {
	// Target environment
	Environment string `pulumi:"environment"`
	// Defines specific repositories to exclude from the promotion.
	ExcludedRepositoryKeys []string `pulumi:"excludedRepositoryKeys"`
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	IncludedRepositoryKeys []string `pulumi:"includedRepositoryKeys"`
	// Key-pair name to use for signature creation
	KeypairName string `pulumi:"keypairName"`
	// Name of Release Bundle
	Name *string `pulumi:"name"`
	// Project key the Release Bundle belongs to
	ProjectKey *string `pulumi:"projectKey"`
	// Version to promote
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a ReleaseBundleV2Promotion resource.
type ReleaseBundleV2PromotionArgs struct {
	// Target environment
	Environment pulumi.StringInput
	// Defines specific repositories to exclude from the promotion.
	ExcludedRepositoryKeys pulumi.StringArrayInput
	// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
	IncludedRepositoryKeys pulumi.StringArrayInput
	// Key-pair name to use for signature creation
	KeypairName pulumi.StringInput
	// Name of Release Bundle
	Name pulumi.StringPtrInput
	// Project key the Release Bundle belongs to
	ProjectKey pulumi.StringPtrInput
	// Version to promote
	Version pulumi.StringInput
}

func (ReleaseBundleV2PromotionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2PromotionArgs)(nil)).Elem()
}

type ReleaseBundleV2PromotionInput interface {
	pulumi.Input

	ToReleaseBundleV2PromotionOutput() ReleaseBundleV2PromotionOutput
	ToReleaseBundleV2PromotionOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionOutput
}

func (*ReleaseBundleV2Promotion) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2Promotion)(nil)).Elem()
}

func (i *ReleaseBundleV2Promotion) ToReleaseBundleV2PromotionOutput() ReleaseBundleV2PromotionOutput {
	return i.ToReleaseBundleV2PromotionOutputWithContext(context.Background())
}

func (i *ReleaseBundleV2Promotion) ToReleaseBundleV2PromotionOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2PromotionOutput)
}

// ReleaseBundleV2PromotionArrayInput is an input type that accepts ReleaseBundleV2PromotionArray and ReleaseBundleV2PromotionArrayOutput values.
// You can construct a concrete instance of `ReleaseBundleV2PromotionArrayInput` via:
//
//	ReleaseBundleV2PromotionArray{ ReleaseBundleV2PromotionArgs{...} }
type ReleaseBundleV2PromotionArrayInput interface {
	pulumi.Input

	ToReleaseBundleV2PromotionArrayOutput() ReleaseBundleV2PromotionArrayOutput
	ToReleaseBundleV2PromotionArrayOutputWithContext(context.Context) ReleaseBundleV2PromotionArrayOutput
}

type ReleaseBundleV2PromotionArray []ReleaseBundleV2PromotionInput

func (ReleaseBundleV2PromotionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2Promotion)(nil)).Elem()
}

func (i ReleaseBundleV2PromotionArray) ToReleaseBundleV2PromotionArrayOutput() ReleaseBundleV2PromotionArrayOutput {
	return i.ToReleaseBundleV2PromotionArrayOutputWithContext(context.Background())
}

func (i ReleaseBundleV2PromotionArray) ToReleaseBundleV2PromotionArrayOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2PromotionArrayOutput)
}

// ReleaseBundleV2PromotionMapInput is an input type that accepts ReleaseBundleV2PromotionMap and ReleaseBundleV2PromotionMapOutput values.
// You can construct a concrete instance of `ReleaseBundleV2PromotionMapInput` via:
//
//	ReleaseBundleV2PromotionMap{ "key": ReleaseBundleV2PromotionArgs{...} }
type ReleaseBundleV2PromotionMapInput interface {
	pulumi.Input

	ToReleaseBundleV2PromotionMapOutput() ReleaseBundleV2PromotionMapOutput
	ToReleaseBundleV2PromotionMapOutputWithContext(context.Context) ReleaseBundleV2PromotionMapOutput
}

type ReleaseBundleV2PromotionMap map[string]ReleaseBundleV2PromotionInput

func (ReleaseBundleV2PromotionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2Promotion)(nil)).Elem()
}

func (i ReleaseBundleV2PromotionMap) ToReleaseBundleV2PromotionMapOutput() ReleaseBundleV2PromotionMapOutput {
	return i.ToReleaseBundleV2PromotionMapOutputWithContext(context.Background())
}

func (i ReleaseBundleV2PromotionMap) ToReleaseBundleV2PromotionMapOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2PromotionMapOutput)
}

type ReleaseBundleV2PromotionOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2PromotionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2Promotion)(nil)).Elem()
}

func (o ReleaseBundleV2PromotionOutput) ToReleaseBundleV2PromotionOutput() ReleaseBundleV2PromotionOutput {
	return o
}

func (o ReleaseBundleV2PromotionOutput) ToReleaseBundleV2PromotionOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionOutput {
	return o
}

// Timestamp when the new version was created (ISO 8601 standard).
func (o ReleaseBundleV2PromotionOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// Timestamp when the new version was created (in milliseconds).
func (o ReleaseBundleV2PromotionOutput) CreatedMillis() pulumi.IntOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.IntOutput { return v.CreatedMillis }).(pulumi.IntOutput)
}

// Target environment
func (o ReleaseBundleV2PromotionOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Defines specific repositories to exclude from the promotion.
func (o ReleaseBundleV2PromotionOutput) ExcludedRepositoryKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringArrayOutput { return v.ExcludedRepositoryKeys }).(pulumi.StringArrayOutput)
}

// Defines specific repositories to include in the promotion. If this property is left undefined, all repositories (except those specifically excluded) are included in the promotion. Important: If one or more repositories are specifically included, all other repositories are excluded (regardless of what is defined in `excludedRepositoryKeys`).
func (o ReleaseBundleV2PromotionOutput) IncludedRepositoryKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringArrayOutput { return v.IncludedRepositoryKeys }).(pulumi.StringArrayOutput)
}

// Key-pair name to use for signature creation
func (o ReleaseBundleV2PromotionOutput) KeypairName() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringOutput { return v.KeypairName }).(pulumi.StringOutput)
}

// Name of Release Bundle
func (o ReleaseBundleV2PromotionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Project key the Release Bundle belongs to
func (o ReleaseBundleV2PromotionOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

// Version to promote
func (o ReleaseBundleV2PromotionOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2Promotion) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ReleaseBundleV2PromotionArrayOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2PromotionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2Promotion)(nil)).Elem()
}

func (o ReleaseBundleV2PromotionArrayOutput) ToReleaseBundleV2PromotionArrayOutput() ReleaseBundleV2PromotionArrayOutput {
	return o
}

func (o ReleaseBundleV2PromotionArrayOutput) ToReleaseBundleV2PromotionArrayOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionArrayOutput {
	return o
}

func (o ReleaseBundleV2PromotionArrayOutput) Index(i pulumi.IntInput) ReleaseBundleV2PromotionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseBundleV2Promotion {
		return vs[0].([]*ReleaseBundleV2Promotion)[vs[1].(int)]
	}).(ReleaseBundleV2PromotionOutput)
}

type ReleaseBundleV2PromotionMapOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2PromotionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2Promotion)(nil)).Elem()
}

func (o ReleaseBundleV2PromotionMapOutput) ToReleaseBundleV2PromotionMapOutput() ReleaseBundleV2PromotionMapOutput {
	return o
}

func (o ReleaseBundleV2PromotionMapOutput) ToReleaseBundleV2PromotionMapOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionMapOutput {
	return o
}

func (o ReleaseBundleV2PromotionMapOutput) MapIndex(k pulumi.StringInput) ReleaseBundleV2PromotionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseBundleV2Promotion {
		return vs[0].(map[string]*ReleaseBundleV2Promotion)[vs[1].(string)]
	}).(ReleaseBundleV2PromotionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2PromotionInput)(nil)).Elem(), &ReleaseBundleV2Promotion{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2PromotionArrayInput)(nil)).Elem(), ReleaseBundleV2PromotionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2PromotionMapInput)(nil)).Elem(), ReleaseBundleV2PromotionMap{})
	pulumi.RegisterOutputType(ReleaseBundleV2PromotionOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2PromotionArrayOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2PromotionMapOutput{})
}
