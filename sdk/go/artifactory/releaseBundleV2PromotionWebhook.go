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

// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
// .
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-artifactory/sdk/v7/go/artifactory"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := artifactory.NewReleaseBundleV2PromotionWebhook(ctx, "release-bundle-v2-promotion-webhook", &artifactory.ReleaseBundleV2PromotionWebhookArgs{
//				Key: pulumi.String("release-bundle-v2-promotion-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("release_bundle_v2_promotion_completed"),
//					pulumi.String("release_bundle_v2_promotion_failed"),
//					pulumi.String("release_bundle_v2_promotion_started"),
//				},
//				Criteria: &artifactory.ReleaseBundleV2PromotionWebhookCriteriaArgs{
//					SelectedEnvironments: pulumi.StringArray{
//						pulumi.String("PROD"),
//						pulumi.String("DEV"),
//					},
//				},
//				Handlers: artifactory.ReleaseBundleV2PromotionWebhookHandlerArray{
//					&artifactory.ReleaseBundleV2PromotionWebhookHandlerArgs{
//						Url:    pulumi.String("http://tempurl.org/webhook"),
//						Secret: pulumi.String("some-secret"),
//						Proxy:  pulumi.String("proxy-key"),
//						CustomHttpHeaders: pulumi.StringMap{
//							"header-1": pulumi.String("value-1"),
//							"header-2": pulumi.String("value-2"),
//						},
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
type ReleaseBundleV2PromotionWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which enviroments.
	Criteria ReleaseBundleV2PromotionWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of event triggers for the Webhook. Allow values: `releaseBundleV2PromotionStarted`, `releaseBundleV2PromotionCompleted`, `releaseBundleV2PromotionFailed`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ReleaseBundleV2PromotionWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewReleaseBundleV2PromotionWebhook registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleV2PromotionWebhook(ctx *pulumi.Context,
	name string, args *ReleaseBundleV2PromotionWebhookArgs, opts ...pulumi.ResourceOption) (*ReleaseBundleV2PromotionWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.Handlers == nil {
		return nil, errors.New("invalid value for required argument 'Handlers'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReleaseBundleV2PromotionWebhook
	err := ctx.RegisterResource("artifactory:index/releaseBundleV2PromotionWebhook:ReleaseBundleV2PromotionWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseBundleV2PromotionWebhook gets an existing ReleaseBundleV2PromotionWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseBundleV2PromotionWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseBundleV2PromotionWebhookState, opts ...pulumi.ResourceOption) (*ReleaseBundleV2PromotionWebhook, error) {
	var resource ReleaseBundleV2PromotionWebhook
	err := ctx.ReadResource("artifactory:index/releaseBundleV2PromotionWebhook:ReleaseBundleV2PromotionWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseBundleV2PromotionWebhook resources.
type releaseBundleV2PromotionWebhookState struct {
	// Specifies where the webhook will be applied on which enviroments.
	Criteria *ReleaseBundleV2PromotionWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of event triggers for the Webhook. Allow values: `releaseBundleV2PromotionStarted`, `releaseBundleV2PromotionCompleted`, `releaseBundleV2PromotionFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleV2PromotionWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ReleaseBundleV2PromotionWebhookState struct {
	// Specifies where the webhook will be applied on which enviroments.
	Criteria ReleaseBundleV2PromotionWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of event triggers for the Webhook. Allow values: `releaseBundleV2PromotionStarted`, `releaseBundleV2PromotionCompleted`, `releaseBundleV2PromotionFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleV2PromotionWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ReleaseBundleV2PromotionWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2PromotionWebhookState)(nil)).Elem()
}

type releaseBundleV2PromotionWebhookArgs struct {
	// Specifies where the webhook will be applied on which enviroments.
	Criteria ReleaseBundleV2PromotionWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`
	Enabled *bool `pulumi:"enabled"`
	// List of event triggers for the Webhook. Allow values: `releaseBundleV2PromotionStarted`, `releaseBundleV2PromotionCompleted`, `releaseBundleV2PromotionFailed`
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleV2PromotionWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ReleaseBundleV2PromotionWebhook resource.
type ReleaseBundleV2PromotionWebhookArgs struct {
	// Specifies where the webhook will be applied on which enviroments.
	Criteria ReleaseBundleV2PromotionWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`
	Enabled pulumi.BoolPtrInput
	// List of event triggers for the Webhook. Allow values: `releaseBundleV2PromotionStarted`, `releaseBundleV2PromotionCompleted`, `releaseBundleV2PromotionFailed`
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleV2PromotionWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ReleaseBundleV2PromotionWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2PromotionWebhookArgs)(nil)).Elem()
}

type ReleaseBundleV2PromotionWebhookInput interface {
	pulumi.Input

	ToReleaseBundleV2PromotionWebhookOutput() ReleaseBundleV2PromotionWebhookOutput
	ToReleaseBundleV2PromotionWebhookOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookOutput
}

func (*ReleaseBundleV2PromotionWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2PromotionWebhook)(nil)).Elem()
}

func (i *ReleaseBundleV2PromotionWebhook) ToReleaseBundleV2PromotionWebhookOutput() ReleaseBundleV2PromotionWebhookOutput {
	return i.ToReleaseBundleV2PromotionWebhookOutputWithContext(context.Background())
}

func (i *ReleaseBundleV2PromotionWebhook) ToReleaseBundleV2PromotionWebhookOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2PromotionWebhookOutput)
}

// ReleaseBundleV2PromotionWebhookArrayInput is an input type that accepts ReleaseBundleV2PromotionWebhookArray and ReleaseBundleV2PromotionWebhookArrayOutput values.
// You can construct a concrete instance of `ReleaseBundleV2PromotionWebhookArrayInput` via:
//
//	ReleaseBundleV2PromotionWebhookArray{ ReleaseBundleV2PromotionWebhookArgs{...} }
type ReleaseBundleV2PromotionWebhookArrayInput interface {
	pulumi.Input

	ToReleaseBundleV2PromotionWebhookArrayOutput() ReleaseBundleV2PromotionWebhookArrayOutput
	ToReleaseBundleV2PromotionWebhookArrayOutputWithContext(context.Context) ReleaseBundleV2PromotionWebhookArrayOutput
}

type ReleaseBundleV2PromotionWebhookArray []ReleaseBundleV2PromotionWebhookInput

func (ReleaseBundleV2PromotionWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2PromotionWebhook)(nil)).Elem()
}

func (i ReleaseBundleV2PromotionWebhookArray) ToReleaseBundleV2PromotionWebhookArrayOutput() ReleaseBundleV2PromotionWebhookArrayOutput {
	return i.ToReleaseBundleV2PromotionWebhookArrayOutputWithContext(context.Background())
}

func (i ReleaseBundleV2PromotionWebhookArray) ToReleaseBundleV2PromotionWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2PromotionWebhookArrayOutput)
}

// ReleaseBundleV2PromotionWebhookMapInput is an input type that accepts ReleaseBundleV2PromotionWebhookMap and ReleaseBundleV2PromotionWebhookMapOutput values.
// You can construct a concrete instance of `ReleaseBundleV2PromotionWebhookMapInput` via:
//
//	ReleaseBundleV2PromotionWebhookMap{ "key": ReleaseBundleV2PromotionWebhookArgs{...} }
type ReleaseBundleV2PromotionWebhookMapInput interface {
	pulumi.Input

	ToReleaseBundleV2PromotionWebhookMapOutput() ReleaseBundleV2PromotionWebhookMapOutput
	ToReleaseBundleV2PromotionWebhookMapOutputWithContext(context.Context) ReleaseBundleV2PromotionWebhookMapOutput
}

type ReleaseBundleV2PromotionWebhookMap map[string]ReleaseBundleV2PromotionWebhookInput

func (ReleaseBundleV2PromotionWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2PromotionWebhook)(nil)).Elem()
}

func (i ReleaseBundleV2PromotionWebhookMap) ToReleaseBundleV2PromotionWebhookMapOutput() ReleaseBundleV2PromotionWebhookMapOutput {
	return i.ToReleaseBundleV2PromotionWebhookMapOutputWithContext(context.Background())
}

func (i ReleaseBundleV2PromotionWebhookMap) ToReleaseBundleV2PromotionWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2PromotionWebhookMapOutput)
}

type ReleaseBundleV2PromotionWebhookOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2PromotionWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2PromotionWebhook)(nil)).Elem()
}

func (o ReleaseBundleV2PromotionWebhookOutput) ToReleaseBundleV2PromotionWebhookOutput() ReleaseBundleV2PromotionWebhookOutput {
	return o
}

func (o ReleaseBundleV2PromotionWebhookOutput) ToReleaseBundleV2PromotionWebhookOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which enviroments.
func (o ReleaseBundleV2PromotionWebhookOutput) Criteria() ReleaseBundleV2PromotionWebhookCriteriaOutput {
	return o.ApplyT(func(v *ReleaseBundleV2PromotionWebhook) ReleaseBundleV2PromotionWebhookCriteriaOutput {
		return v.Criteria
	}).(ReleaseBundleV2PromotionWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o ReleaseBundleV2PromotionWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleV2PromotionWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`
func (o ReleaseBundleV2PromotionWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleV2PromotionWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of event triggers for the Webhook. Allow values: `releaseBundleV2PromotionStarted`, `releaseBundleV2PromotionCompleted`, `releaseBundleV2PromotionFailed`
func (o ReleaseBundleV2PromotionWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleV2PromotionWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ReleaseBundleV2PromotionWebhookOutput) Handlers() ReleaseBundleV2PromotionWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleV2PromotionWebhook) ReleaseBundleV2PromotionWebhookHandlerArrayOutput {
		return v.Handlers
	}).(ReleaseBundleV2PromotionWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ReleaseBundleV2PromotionWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2PromotionWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ReleaseBundleV2PromotionWebhookArrayOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2PromotionWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2PromotionWebhook)(nil)).Elem()
}

func (o ReleaseBundleV2PromotionWebhookArrayOutput) ToReleaseBundleV2PromotionWebhookArrayOutput() ReleaseBundleV2PromotionWebhookArrayOutput {
	return o
}

func (o ReleaseBundleV2PromotionWebhookArrayOutput) ToReleaseBundleV2PromotionWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookArrayOutput {
	return o
}

func (o ReleaseBundleV2PromotionWebhookArrayOutput) Index(i pulumi.IntInput) ReleaseBundleV2PromotionWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseBundleV2PromotionWebhook {
		return vs[0].([]*ReleaseBundleV2PromotionWebhook)[vs[1].(int)]
	}).(ReleaseBundleV2PromotionWebhookOutput)
}

type ReleaseBundleV2PromotionWebhookMapOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2PromotionWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2PromotionWebhook)(nil)).Elem()
}

func (o ReleaseBundleV2PromotionWebhookMapOutput) ToReleaseBundleV2PromotionWebhookMapOutput() ReleaseBundleV2PromotionWebhookMapOutput {
	return o
}

func (o ReleaseBundleV2PromotionWebhookMapOutput) ToReleaseBundleV2PromotionWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleV2PromotionWebhookMapOutput {
	return o
}

func (o ReleaseBundleV2PromotionWebhookMapOutput) MapIndex(k pulumi.StringInput) ReleaseBundleV2PromotionWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseBundleV2PromotionWebhook {
		return vs[0].(map[string]*ReleaseBundleV2PromotionWebhook)[vs[1].(string)]
	}).(ReleaseBundleV2PromotionWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2PromotionWebhookInput)(nil)).Elem(), &ReleaseBundleV2PromotionWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2PromotionWebhookArrayInput)(nil)).Elem(), ReleaseBundleV2PromotionWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2PromotionWebhookMapInput)(nil)).Elem(), ReleaseBundleV2PromotionWebhookMap{})
	pulumi.RegisterOutputType(ReleaseBundleV2PromotionWebhookOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2PromotionWebhookArrayOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2PromotionWebhookMapOutput{})
}