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

// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// !>This resource is being deprecated and replaced by `DestinationWebhook` resource.
//
// ## Example Usage
//
// .
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
//			_, err := artifactory.NewReleaseBundleWebhook(ctx, "release-bundle-webhook", &artifactory.ReleaseBundleWebhookArgs{
//				Key: pulumi.String("release-bundle-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("created"),
//					pulumi.String("signed"),
//					pulumi.String("deleted"),
//				},
//				Criteria: &artifactory.ReleaseBundleWebhookCriteriaArgs{
//					AnyReleaseBundle: pulumi.Bool(false),
//					RegisteredReleaseBundleNames: pulumi.StringArray{
//						pulumi.String("bundle-name"),
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//				},
//				Handlers: artifactory.ReleaseBundleWebhookHandlerArray{
//					&artifactory.ReleaseBundleWebhookHandlerArgs{
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
type ReleaseBundleWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteriaPtrOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ReleaseBundleWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewReleaseBundleWebhook registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleWebhook(ctx *pulumi.Context,
	name string, args *ReleaseBundleWebhookArgs, opts ...pulumi.ResourceOption) (*ReleaseBundleWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ReleaseBundleWebhook
	err := ctx.RegisterResource("artifactory:index/releaseBundleWebhook:ReleaseBundleWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseBundleWebhook gets an existing ReleaseBundleWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseBundleWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseBundleWebhookState, opts ...pulumi.ResourceOption) (*ReleaseBundleWebhook, error) {
	var resource ReleaseBundleWebhook
	err := ctx.ReadResource("artifactory:index/releaseBundleWebhook:ReleaseBundleWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseBundleWebhook resources.
type releaseBundleWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ReleaseBundleWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ReleaseBundleWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ReleaseBundleWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleWebhookState)(nil)).Elem()
}

type releaseBundleWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ReleaseBundleWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ReleaseBundleWebhook resource.
type ReleaseBundleWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ReleaseBundleWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleWebhookArgs)(nil)).Elem()
}

type ReleaseBundleWebhookInput interface {
	pulumi.Input

	ToReleaseBundleWebhookOutput() ReleaseBundleWebhookOutput
	ToReleaseBundleWebhookOutputWithContext(ctx context.Context) ReleaseBundleWebhookOutput
}

func (*ReleaseBundleWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleWebhook)(nil)).Elem()
}

func (i *ReleaseBundleWebhook) ToReleaseBundleWebhookOutput() ReleaseBundleWebhookOutput {
	return i.ToReleaseBundleWebhookOutputWithContext(context.Background())
}

func (i *ReleaseBundleWebhook) ToReleaseBundleWebhookOutputWithContext(ctx context.Context) ReleaseBundleWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleWebhookOutput)
}

// ReleaseBundleWebhookArrayInput is an input type that accepts ReleaseBundleWebhookArray and ReleaseBundleWebhookArrayOutput values.
// You can construct a concrete instance of `ReleaseBundleWebhookArrayInput` via:
//
//	ReleaseBundleWebhookArray{ ReleaseBundleWebhookArgs{...} }
type ReleaseBundleWebhookArrayInput interface {
	pulumi.Input

	ToReleaseBundleWebhookArrayOutput() ReleaseBundleWebhookArrayOutput
	ToReleaseBundleWebhookArrayOutputWithContext(context.Context) ReleaseBundleWebhookArrayOutput
}

type ReleaseBundleWebhookArray []ReleaseBundleWebhookInput

func (ReleaseBundleWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleWebhook)(nil)).Elem()
}

func (i ReleaseBundleWebhookArray) ToReleaseBundleWebhookArrayOutput() ReleaseBundleWebhookArrayOutput {
	return i.ToReleaseBundleWebhookArrayOutputWithContext(context.Background())
}

func (i ReleaseBundleWebhookArray) ToReleaseBundleWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleWebhookArrayOutput)
}

// ReleaseBundleWebhookMapInput is an input type that accepts ReleaseBundleWebhookMap and ReleaseBundleWebhookMapOutput values.
// You can construct a concrete instance of `ReleaseBundleWebhookMapInput` via:
//
//	ReleaseBundleWebhookMap{ "key": ReleaseBundleWebhookArgs{...} }
type ReleaseBundleWebhookMapInput interface {
	pulumi.Input

	ToReleaseBundleWebhookMapOutput() ReleaseBundleWebhookMapOutput
	ToReleaseBundleWebhookMapOutputWithContext(context.Context) ReleaseBundleWebhookMapOutput
}

type ReleaseBundleWebhookMap map[string]ReleaseBundleWebhookInput

func (ReleaseBundleWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleWebhook)(nil)).Elem()
}

func (i ReleaseBundleWebhookMap) ToReleaseBundleWebhookMapOutput() ReleaseBundleWebhookMapOutput {
	return i.ToReleaseBundleWebhookMapOutputWithContext(context.Background())
}

func (i ReleaseBundleWebhookMap) ToReleaseBundleWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleWebhookMapOutput)
}

type ReleaseBundleWebhookOutput struct{ *pulumi.OutputState }

func (ReleaseBundleWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleWebhook)(nil)).Elem()
}

func (o ReleaseBundleWebhookOutput) ToReleaseBundleWebhookOutput() ReleaseBundleWebhookOutput {
	return o
}

func (o ReleaseBundleWebhookOutput) ToReleaseBundleWebhookOutputWithContext(ctx context.Context) ReleaseBundleWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o ReleaseBundleWebhookOutput) Criteria() ReleaseBundleWebhookCriteriaPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleWebhook) ReleaseBundleWebhookCriteriaPtrOutput { return v.Criteria }).(ReleaseBundleWebhookCriteriaPtrOutput)
}

// Webhook description. Max length 1000 characters.
func (o ReleaseBundleWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o ReleaseBundleWebhookOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ReleaseBundleWebhook) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `created`, `signed`, `deleted`.
func (o ReleaseBundleWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ReleaseBundleWebhookOutput) Handlers() ReleaseBundleWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleWebhook) ReleaseBundleWebhookHandlerArrayOutput { return v.Handlers }).(ReleaseBundleWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ReleaseBundleWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ReleaseBundleWebhookArrayOutput struct{ *pulumi.OutputState }

func (ReleaseBundleWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleWebhook)(nil)).Elem()
}

func (o ReleaseBundleWebhookArrayOutput) ToReleaseBundleWebhookArrayOutput() ReleaseBundleWebhookArrayOutput {
	return o
}

func (o ReleaseBundleWebhookArrayOutput) ToReleaseBundleWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleWebhookArrayOutput {
	return o
}

func (o ReleaseBundleWebhookArrayOutput) Index(i pulumi.IntInput) ReleaseBundleWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseBundleWebhook {
		return vs[0].([]*ReleaseBundleWebhook)[vs[1].(int)]
	}).(ReleaseBundleWebhookOutput)
}

type ReleaseBundleWebhookMapOutput struct{ *pulumi.OutputState }

func (ReleaseBundleWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleWebhook)(nil)).Elem()
}

func (o ReleaseBundleWebhookMapOutput) ToReleaseBundleWebhookMapOutput() ReleaseBundleWebhookMapOutput {
	return o
}

func (o ReleaseBundleWebhookMapOutput) ToReleaseBundleWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleWebhookMapOutput {
	return o
}

func (o ReleaseBundleWebhookMapOutput) MapIndex(k pulumi.StringInput) ReleaseBundleWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseBundleWebhook {
		return vs[0].(map[string]*ReleaseBundleWebhook)[vs[1].(string)]
	}).(ReleaseBundleWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleWebhookInput)(nil)).Elem(), &ReleaseBundleWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleWebhookArrayInput)(nil)).Elem(), ReleaseBundleWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleWebhookMapInput)(nil)).Elem(), ReleaseBundleWebhookMap{})
	pulumi.RegisterOutputType(ReleaseBundleWebhookOutput{})
	pulumi.RegisterOutputType(ReleaseBundleWebhookArrayOutput{})
	pulumi.RegisterOutputType(ReleaseBundleWebhookMapOutput{})
}
