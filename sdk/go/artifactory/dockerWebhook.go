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
//			my_docker_local, err := artifactory.NewDockerV2Repository(ctx, "my-docker-local", &artifactory.DockerV2RepositoryArgs{
//				Key: pulumi.String("my-docker-local"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = artifactory.NewDockerWebhook(ctx, "docker-webhook", &artifactory.DockerWebhookArgs{
//				Key: pulumi.String("docker-webhook"),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("pushed"),
//					pulumi.String("deleted"),
//					pulumi.String("promoted"),
//				},
//				Criteria: &artifactory.DockerWebhookCriteriaArgs{
//					AnyLocal:     pulumi.Bool(true),
//					AnyRemote:    pulumi.Bool(false),
//					AnyFederated: pulumi.Bool(false),
//					RepoKeys: pulumi.StringArray{
//						my_docker_local.Key,
//					},
//					IncludePatterns: pulumi.StringArray{
//						pulumi.String("foo/**"),
//					},
//					ExcludePatterns: pulumi.StringArray{
//						pulumi.String("bar/**"),
//					},
//				},
//				Handlers: artifactory.DockerWebhookHandlerArray{
//					&artifactory.DockerWebhookHandlerArgs{
//						Url:    pulumi.String("http://tempurl.org/webhook"),
//						Secret: pulumi.String("some-secret"),
//						Proxy:  pulumi.String("proxy-key"),
//						CustomHttpHeaders: pulumi.StringMap{
//							"header-1": pulumi.String("value-1"),
//							"header-2": pulumi.String("value-2"),
//						},
//					},
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				my_docker_local,
//			}))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DockerWebhook struct {
	pulumi.CustomResourceState

	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerWebhookCriteriaPtrOutput `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// At least one is required.
	Handlers DockerWebhookHandlerArrayOutput `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringOutput `pulumi:"key"`
}

// NewDockerWebhook registers a new resource with the given unique name, arguments, and options.
func NewDockerWebhook(ctx *pulumi.Context,
	name string, args *DockerWebhookArgs, opts ...pulumi.ResourceOption) (*DockerWebhook, error) {
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
	var resource DockerWebhook
	err := ctx.RegisterResource("artifactory:index/dockerWebhook:DockerWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDockerWebhook gets an existing DockerWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDockerWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DockerWebhookState, opts ...pulumi.ResourceOption) (*DockerWebhook, error) {
	var resource DockerWebhook
	err := ctx.ReadResource("artifactory:index/dockerWebhook:DockerWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DockerWebhook resources.
type dockerWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *DockerWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DockerWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key *string `pulumi:"key"`
}

type DockerWebhookState struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DockerWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringPtrInput
}

func (DockerWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerWebhookState)(nil)).Elem()
}

type dockerWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria *DockerWebhookCriteria `pulumi:"criteria"`
	// Webhook description. Max length 1000 characters.
	Description *string `pulumi:"description"`
	// Status of webhook. Default to `true`.
	Enabled *bool `pulumi:"enabled"`
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes []string `pulumi:"eventTypes"`
	// At least one is required.
	Handlers []DockerWebhookHandler `pulumi:"handlers"`
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key string `pulumi:"key"`
}

// The set of arguments for constructing a DockerWebhook resource.
type DockerWebhookArgs struct {
	// Specifies where the webhook will be applied on which repositories.
	Criteria DockerWebhookCriteriaPtrInput
	// Webhook description. Max length 1000 characters.
	Description pulumi.StringPtrInput
	// Status of webhook. Default to `true`.
	Enabled pulumi.BoolPtrInput
	// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
	EventTypes pulumi.StringArrayInput
	// At least one is required.
	Handlers DockerWebhookHandlerArrayInput
	// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
	Key pulumi.StringInput
}

func (DockerWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dockerWebhookArgs)(nil)).Elem()
}

type DockerWebhookInput interface {
	pulumi.Input

	ToDockerWebhookOutput() DockerWebhookOutput
	ToDockerWebhookOutputWithContext(ctx context.Context) DockerWebhookOutput
}

func (*DockerWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerWebhook)(nil)).Elem()
}

func (i *DockerWebhook) ToDockerWebhookOutput() DockerWebhookOutput {
	return i.ToDockerWebhookOutputWithContext(context.Background())
}

func (i *DockerWebhook) ToDockerWebhookOutputWithContext(ctx context.Context) DockerWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerWebhookOutput)
}

// DockerWebhookArrayInput is an input type that accepts DockerWebhookArray and DockerWebhookArrayOutput values.
// You can construct a concrete instance of `DockerWebhookArrayInput` via:
//
//	DockerWebhookArray{ DockerWebhookArgs{...} }
type DockerWebhookArrayInput interface {
	pulumi.Input

	ToDockerWebhookArrayOutput() DockerWebhookArrayOutput
	ToDockerWebhookArrayOutputWithContext(context.Context) DockerWebhookArrayOutput
}

type DockerWebhookArray []DockerWebhookInput

func (DockerWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerWebhook)(nil)).Elem()
}

func (i DockerWebhookArray) ToDockerWebhookArrayOutput() DockerWebhookArrayOutput {
	return i.ToDockerWebhookArrayOutputWithContext(context.Background())
}

func (i DockerWebhookArray) ToDockerWebhookArrayOutputWithContext(ctx context.Context) DockerWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerWebhookArrayOutput)
}

// DockerWebhookMapInput is an input type that accepts DockerWebhookMap and DockerWebhookMapOutput values.
// You can construct a concrete instance of `DockerWebhookMapInput` via:
//
//	DockerWebhookMap{ "key": DockerWebhookArgs{...} }
type DockerWebhookMapInput interface {
	pulumi.Input

	ToDockerWebhookMapOutput() DockerWebhookMapOutput
	ToDockerWebhookMapOutputWithContext(context.Context) DockerWebhookMapOutput
}

type DockerWebhookMap map[string]DockerWebhookInput

func (DockerWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerWebhook)(nil)).Elem()
}

func (i DockerWebhookMap) ToDockerWebhookMapOutput() DockerWebhookMapOutput {
	return i.ToDockerWebhookMapOutputWithContext(context.Background())
}

func (i DockerWebhookMap) ToDockerWebhookMapOutputWithContext(ctx context.Context) DockerWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DockerWebhookMapOutput)
}

type DockerWebhookOutput struct{ *pulumi.OutputState }

func (DockerWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DockerWebhook)(nil)).Elem()
}

func (o DockerWebhookOutput) ToDockerWebhookOutput() DockerWebhookOutput {
	return o
}

func (o DockerWebhookOutput) ToDockerWebhookOutputWithContext(ctx context.Context) DockerWebhookOutput {
	return o
}

// Specifies where the webhook will be applied on which repositories.
func (o DockerWebhookOutput) Criteria() DockerWebhookCriteriaPtrOutput {
	return o.ApplyT(func(v *DockerWebhook) DockerWebhookCriteriaPtrOutput { return v.Criteria }).(DockerWebhookCriteriaPtrOutput)
}

// Webhook description. Max length 1000 characters.
func (o DockerWebhookOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DockerWebhook) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Status of webhook. Default to `true`.
func (o DockerWebhookOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *DockerWebhook) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// List of Events in Artifactory, Distribution, Release Bundle that function as the event trigger for the Webhook. Allow values: `pushed`, `deleted`, `promoted`.
func (o DockerWebhookOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DockerWebhook) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// At least one is required.
func (o DockerWebhookOutput) Handlers() DockerWebhookHandlerArrayOutput {
	return o.ApplyT(func(v *DockerWebhook) DockerWebhookHandlerArrayOutput { return v.Handlers }).(DockerWebhookHandlerArrayOutput)
}

// The identity key of the webhook. Must be between 2 and 200 characters. Cannot contain spaces.
func (o DockerWebhookOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *DockerWebhook) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

type DockerWebhookArrayOutput struct{ *pulumi.OutputState }

func (DockerWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DockerWebhook)(nil)).Elem()
}

func (o DockerWebhookArrayOutput) ToDockerWebhookArrayOutput() DockerWebhookArrayOutput {
	return o
}

func (o DockerWebhookArrayOutput) ToDockerWebhookArrayOutputWithContext(ctx context.Context) DockerWebhookArrayOutput {
	return o
}

func (o DockerWebhookArrayOutput) Index(i pulumi.IntInput) DockerWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DockerWebhook {
		return vs[0].([]*DockerWebhook)[vs[1].(int)]
	}).(DockerWebhookOutput)
}

type DockerWebhookMapOutput struct{ *pulumi.OutputState }

func (DockerWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DockerWebhook)(nil)).Elem()
}

func (o DockerWebhookMapOutput) ToDockerWebhookMapOutput() DockerWebhookMapOutput {
	return o
}

func (o DockerWebhookMapOutput) ToDockerWebhookMapOutputWithContext(ctx context.Context) DockerWebhookMapOutput {
	return o
}

func (o DockerWebhookMapOutput) MapIndex(k pulumi.StringInput) DockerWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DockerWebhook {
		return vs[0].(map[string]*DockerWebhook)[vs[1].(string)]
	}).(DockerWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DockerWebhookInput)(nil)).Elem(), &DockerWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerWebhookArrayInput)(nil)).Elem(), DockerWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DockerWebhookMapInput)(nil)).Elem(), DockerWebhookMap{})
	pulumi.RegisterOutputType(DockerWebhookOutput{})
	pulumi.RegisterOutputType(DockerWebhookArrayOutput{})
	pulumi.RegisterOutputType(DockerWebhookMapOutput{})
}
