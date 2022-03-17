// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package artifactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # Artifactory Release Bundle Webhook Resource
//
// Provides an Artifactory webhook resource. This can be used to register and manage Artifactory webhook subscription which enables you to be notified or notify other users when such events take place in Artifactory.
//
// ## Example Usage
//
// .
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-artifactory/sdk/go/artifactory"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := artifactory.NewReleaseBundleWebhook(ctx, "release-bundle-webhook", &artifactory.ReleaseBundleWebhookArgs{
// 			Criteria: &ReleaseBundleWebhookCriteriaArgs{
// 				AnyReleaseBundle: pulumi.Bool(false),
// 				ExcludePatterns: pulumi.StringArray{
// 					pulumi.String("bar/**"),
// 				},
// 				IncludePatterns: pulumi.StringArray{
// 					pulumi.String("foo/**"),
// 				},
// 				RegisteredReleaseBundleNames: pulumi.StringArray{
// 					pulumi.String("bundle-name"),
// 				},
// 			},
// 			CustomHttpHeaders: pulumi.StringMap{
// 				"header-1": pulumi.String("value-1"),
// 				"header-2": pulumi.String("value-2"),
// 			},
// 			EventTypes: pulumi.StringArray{
// 				pulumi.String("created"),
// 				pulumi.String("signed"),
// 				pulumi.String("deleted"),
// 			},
// 			Key:    pulumi.String("release-bundle-webhook"),
// 			Proxy:  pulumi.String("proxy-key"),
// 			Secret: pulumi.String("some-secret"),
// 			Url:    pulumi.String("http://tempurl.org/webhook"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ReleaseBundleWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteriaOutput `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapOutput `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrOutput `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewReleaseBundleWebhook registers a new resource with the given unique name, arguments, and options.
func NewReleaseBundleWebhook(ctx *pulumi.Context,
	name string, args *ReleaseBundleWebhookArgs, opts ...pulumi.ResourceOption) (*ReleaseBundleWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Criteria == nil {
		return nil, errors.New("invalid value for required argument 'Criteria'")
	}
	if args.EventTypes == nil {
		return nil, errors.New("invalid value for required argument 'EventTypes'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
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
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders map[string]string `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes []string `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret *string `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url *string `pulumi:"url"`
}

type ReleaseBundleWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteriaPtrInput
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes pulumi.StringArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrInput
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringPtrInput
}

func (ReleaseBundleWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseBundleWebhookState)(nil)).Elem()
}

type releaseBundleWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteria `pulumi:"criteria"`
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders map[string]string `pulumi:"customHttpHeaders"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to 'true'
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes []string `pulumi:"eventTypes"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
	// Proxy key from Artifactory Proxies setting
	Proxy *string `pulumi:"proxy"`
	// Secret authentication token that will be sent to the configured URL
	Secret *string `pulumi:"secret"`
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url string `pulumi:"url"`
}

// The set of arguments for constructing a ReleaseBundleWebhook resource.
type ReleaseBundleWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria ReleaseBundleWebhookCriteriaInput
	// Custom HTTP headers you wish to use to invoke the Webhook, comprise of key/value pair.
	CustomHttpHeaders pulumi.StringMapInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to 'true'
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: "created", "signed", "deleted"
	EventTypes pulumi.StringArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
	// Proxy key from Artifactory Proxies setting
	Proxy pulumi.StringPtrInput
	// Secret authentication token that will be sent to the configured URL
	Secret pulumi.StringPtrInput
	// Specifies the URL that the Webhook invokes. This will be the URL that Artifactory will send an HTTP POST request to.
	Url pulumi.StringInput
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
//          ReleaseBundleWebhookArray{ ReleaseBundleWebhookArgs{...} }
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
//          ReleaseBundleWebhookMap{ "key": ReleaseBundleWebhookArgs{...} }
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