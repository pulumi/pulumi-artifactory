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

// Provides an Artifactory custom webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
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
//			_, err := artifactory.NewReleaseBundleV2CustomWebhook(ctx, "release-bundle-v2-custom-webhook", &artifactory.ReleaseBundleV2CustomWebhookArgs{
//				Key: pulumi.String("release-bundle-v2-custom-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("release_bundle_v2_started"),
//					pulumi.String("release_bundle_v2_failed"),
//					pulumi.String("release_bundle_v2_completed"),
//				},
//				Criteria: &artifactory.ReleaseBundleV2CustomWebhookCriteriaArgs{
//					AnyReleaseBundle: pulumi.Bool(false),
//					SelectedReleaseBundles: pulumi.StringArray{
//						pulumi.String("bundle-name"),
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//				},
//				Handlers: artifactory.ReleaseBundleV2CustomWebhookHandlerArray{
//					&artifactory.ReleaseBundleV2CustomWebhookHandlerArgs{
//						Url: pulumi.String("https://tempurl.org"),
//						Secrets: pulumi.StringMap{
//							"secretName1": pulumi.String("value1"),
//							"secretName2": pulumi.String("value2"),
//						},
//						HttpHeaders: pulumi.StringMap{
//							"headerName1": pulumi.String("value1"),
//							"headerName2": pulumi.String("value2"),
//						},
//						Payload: pulumi.String("{ \"ref\": \"main\" , \"inputs\": { \"artifact_path\": \"test-repo/repo-path\" } }"),
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
type ReleaseBundleV2CustomWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleV2CustomWebhookCriteriaOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `releaseBundleV2Started`, `releaseBundleV2Failed`, `releaseBundleV2Completed`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers ReleaseBundleV2CustomWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewReleaseBundleV2CustomWebhook registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleV2CustomWebhook(ctx *pulumi.Context,
	name string, args *ReleaseBundleV2CustomWebhookArgs, opts ...pulumi.ResourceOption) (*ReleaseBundleV2CustomWebhook, error) {
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
	var resource ReleaseBundleV2CustomWebhook
	err := ctx.RegisterResource("artifactory:index/releaseBundleV2CustomWebhook:ReleaseBundleV2CustomWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetReleaseBundleV2CustomWebhook gets an existing ReleaseBundleV2CustomWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetReleaseBundleV2CustomWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseBundleV2CustomWebhookState, opts ...pulumi.ResourceOption) (*ReleaseBundleV2CustomWebhook, error) {
	var resource ReleaseBundleV2CustomWebhook
	err := ctx.ReadResource("artifactory:index/releaseBundleV2CustomWebhook:ReleaseBundleV2CustomWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ReleaseBundleV2CustomWebhook resources.
type releaseBundleV2CustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *ReleaseBundleV2CustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `releaseBundleV2Started`, `releaseBundleV2Failed`, `releaseBundleV2Completed`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleV2CustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type ReleaseBundleV2CustomWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleV2CustomWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `releaseBundleV2Started`, `releaseBundleV2Failed`, `releaseBundleV2Completed`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleV2CustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (ReleaseBundleV2CustomWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2CustomWebhookState)(nil)).Elem()
}

type releaseBundleV2CustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleV2CustomWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `releaseBundleV2Started`, `releaseBundleV2Failed`, `releaseBundleV2Completed`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []ReleaseBundleV2CustomWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a ReleaseBundleV2CustomWebhook resource.
type ReleaseBundleV2CustomWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleV2CustomWebhookCriteriaInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `releaseBundleV2Started`, `releaseBundleV2Failed`, `releaseBundleV2Completed`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers ReleaseBundleV2CustomWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (ReleaseBundleV2CustomWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleV2CustomWebhookArgs)(nil)).Elem()
}

type ReleaseBundleV2CustomWebhookInput interface {
	pulumi.Input

	ToReleaseBundleV2CustomWebhookOutput() ReleaseBundleV2CustomWebhookOutput
	ToReleaseBundleV2CustomWebhookOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookOutput
}

func (*ReleaseBundleV2CustomWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2CustomWebhook)(nil)).Elem()
}

func (i *ReleaseBundleV2CustomWebhook) ToReleaseBundleV2CustomWebhookOutput() ReleaseBundleV2CustomWebhookOutput {
	return i.ToReleaseBundleV2CustomWebhookOutputWithContext(context.Background())
}

func (i *ReleaseBundleV2CustomWebhook) ToReleaseBundleV2CustomWebhookOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2CustomWebhookOutput)
}

// ReleaseBundleV2CustomWebhookArrayInput is an input type that accepts ReleaseBundleV2CustomWebhookArray and ReleaseBundleV2CustomWebhookArrayOutput values.
// You can construct a concrete instance of `ReleaseBundleV2CustomWebhookArrayInput` via:
//
//	ReleaseBundleV2CustomWebhookArray{ ReleaseBundleV2CustomWebhookArgs{...} }
type ReleaseBundleV2CustomWebhookArrayInput interface {
	pulumi.Input

	ToReleaseBundleV2CustomWebhookArrayOutput() ReleaseBundleV2CustomWebhookArrayOutput
	ToReleaseBundleV2CustomWebhookArrayOutputWithContext(context.Context) ReleaseBundleV2CustomWebhookArrayOutput
}

type ReleaseBundleV2CustomWebhookArray []ReleaseBundleV2CustomWebhookInput

func (ReleaseBundleV2CustomWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2CustomWebhook)(nil)).Elem()
}

func (i ReleaseBundleV2CustomWebhookArray) ToReleaseBundleV2CustomWebhookArrayOutput() ReleaseBundleV2CustomWebhookArrayOutput {
	return i.ToReleaseBundleV2CustomWebhookArrayOutputWithContext(context.Background())
}

func (i ReleaseBundleV2CustomWebhookArray) ToReleaseBundleV2CustomWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2CustomWebhookArrayOutput)
}

// ReleaseBundleV2CustomWebhookMapInput is an input type that accepts ReleaseBundleV2CustomWebhookMap and ReleaseBundleV2CustomWebhookMapOutput values.
// You can construct a concrete instance of `ReleaseBundleV2CustomWebhookMapInput` via:
//
//	ReleaseBundleV2CustomWebhookMap{ "key": ReleaseBundleV2CustomWebhookArgs{...} }
type ReleaseBundleV2CustomWebhookMapInput interface {
	pulumi.Input

	ToReleaseBundleV2CustomWebhookMapOutput() ReleaseBundleV2CustomWebhookMapOutput
	ToReleaseBundleV2CustomWebhookMapOutputWithContext(context.Context) ReleaseBundleV2CustomWebhookMapOutput
}

type ReleaseBundleV2CustomWebhookMap map[string]ReleaseBundleV2CustomWebhookInput

func (ReleaseBundleV2CustomWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2CustomWebhook)(nil)).Elem()
}

func (i ReleaseBundleV2CustomWebhookMap) ToReleaseBundleV2CustomWebhookMapOutput() ReleaseBundleV2CustomWebhookMapOutput {
	return i.ToReleaseBundleV2CustomWebhookMapOutputWithContext(context.Background())
}

func (i ReleaseBundleV2CustomWebhookMap) ToReleaseBundleV2CustomWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseBundleV2CustomWebhookMapOutput)
}

type ReleaseBundleV2CustomWebhookOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2CustomWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReleaseBundleV2CustomWebhook)(nil)).Elem()
}

func (o ReleaseBundleV2CustomWebhookOutput) ToReleaseBundleV2CustomWebhookOutput() ReleaseBundleV2CustomWebhookOutput {
	return o
}

func (o ReleaseBundleV2CustomWebhookOutput) ToReleaseBundleV2CustomWebhookOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o ReleaseBundleV2CustomWebhookOutput) Criteria() ReleaseBundleV2CustomWebhookCriteriaOutput {
	return o.ApplyT(func(v *ReleaseBundleV2CustomWebhook) ReleaseBundleV2CustomWebhookCriteriaOutput { return v.Criteria }).(ReleaseBundleV2CustomWebhookCriteriaOutput)
}

// Webhook description. Max length 1000 characters.
func (o ReleaseBundleV2CustomWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleV2CustomWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o ReleaseBundleV2CustomWebhookOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ReleaseBundleV2CustomWebhook) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `releaseBundleV2Started`, `releaseBundleV2Failed`, `releaseBundleV2Completed`.
func (o ReleaseBundleV2CustomWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleV2CustomWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o ReleaseBundleV2CustomWebhookOutput) Handlers() ReleaseBundleV2CustomWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *ReleaseBundleV2CustomWebhook) ReleaseBundleV2CustomWebhookHandlerArrayOutput {
		return v.Handlers
	}).(ReleaseBundleV2CustomWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o ReleaseBundleV2CustomWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *ReleaseBundleV2CustomWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type ReleaseBundleV2CustomWebhookArrayOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2CustomWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ReleaseBundleV2CustomWebhook)(nil)).Elem()
}

func (o ReleaseBundleV2CustomWebhookArrayOutput) ToReleaseBundleV2CustomWebhookArrayOutput() ReleaseBundleV2CustomWebhookArrayOutput {
	return o
}

func (o ReleaseBundleV2CustomWebhookArrayOutput) ToReleaseBundleV2CustomWebhookArrayOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookArrayOutput {
	return o
}

func (o ReleaseBundleV2CustomWebhookArrayOutput) Index(i pulumi.IntInput) ReleaseBundleV2CustomWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ReleaseBundleV2CustomWebhook {
		return vs[0].([]*ReleaseBundleV2CustomWebhook)[vs[1].(int)]
	}).(ReleaseBundleV2CustomWebhookOutput)
}

type ReleaseBundleV2CustomWebhookMapOutput struct{ *pulumi.OutputState }

func (ReleaseBundleV2CustomWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ReleaseBundleV2CustomWebhook)(nil)).Elem()
}

func (o ReleaseBundleV2CustomWebhookMapOutput) ToReleaseBundleV2CustomWebhookMapOutput() ReleaseBundleV2CustomWebhookMapOutput {
	return o
}

func (o ReleaseBundleV2CustomWebhookMapOutput) ToReleaseBundleV2CustomWebhookMapOutputWithContext(ctx context.Context) ReleaseBundleV2CustomWebhookMapOutput {
	return o
}

func (o ReleaseBundleV2CustomWebhookMapOutput) MapIndex(k pulumi.StringInput) ReleaseBundleV2CustomWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ReleaseBundleV2CustomWebhook {
		return vs[0].(map[string]*ReleaseBundleV2CustomWebhook)[vs[1].(string)]
	}).(ReleaseBundleV2CustomWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2CustomWebhookInput)(nil)).Elem(), &ReleaseBundleV2CustomWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2CustomWebhookArrayInput)(nil)).Elem(), ReleaseBundleV2CustomWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseBundleV2CustomWebhookMapInput)(nil)).Elem(), ReleaseBundleV2CustomWebhookMap{})
	pulumi.RegisterOutputType(ReleaseBundleV2CustomWebhookOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2CustomWebhookArrayOutput{})
	pulumi.RegisterOutputType(ReleaseBundleV2CustomWebhookMapOutput{})
}