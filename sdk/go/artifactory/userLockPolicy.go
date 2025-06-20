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

// Provides an Artifactory User Lock Policy resource.
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
//			_, err := artifactory.NewUserLockPolicy(ctx, "my-user-lock-policy", &artifactory.UserLockPolicyArgs{
//				Name:          pulumi.String("my-user-lock-policy"),
//				Enabled:       pulumi.Bool(true),
//				LoginAttempts: pulumi.Int(10),
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
// ```sh
// $ pulumi import artifactory:index/userLockPolicy:UserLockPolicy my-user-lock-policy my-user-lock-policy
// ```
type UserLockPolicy struct {
	pulumi.CustomResourceState

	// Enable User Lock Policy. Lock user after exceeding max failed login attempts.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Max failed login attempts.
	LoginAttempts pulumi.IntOutput `pulumi:"loginAttempts"`
	// Name of the resource. Only used for importing.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewUserLockPolicy registers a new resource with the given unique name, arguments, and options.
func NewUserLockPolicy(ctx *pulumi.Context,
	name string, args *UserLockPolicyArgs, opts ...pulumi.ResourceOption) (*UserLockPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.LoginAttempts == nil {
		return nil, errors.New("invalid value for required argument 'LoginAttempts'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource UserLockPolicy
	err := ctx.RegisterResource("artifactory:index/userLockPolicy:UserLockPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserLockPolicy gets an existing UserLockPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserLockPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserLockPolicyState, opts ...pulumi.ResourceOption) (*UserLockPolicy, error) {
	var resource UserLockPolicy
	err := ctx.ReadResource("artifactory:index/userLockPolicy:UserLockPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserLockPolicy resources.
type userLockPolicyState struct {
	// Enable User Lock Policy. Lock user after exceeding max failed login attempts.
	Enabled *bool `pulumi:"enabled"`
	// Max failed login attempts.
	LoginAttempts *int `pulumi:"loginAttempts"`
	// Name of the resource. Only used for importing.
	Name *string `pulumi:"name"`
}

type UserLockPolicyState struct {
	// Enable User Lock Policy. Lock user after exceeding max failed login attempts.
	Enabled pulumi.BoolPtrInput
	// Max failed login attempts.
	LoginAttempts pulumi.IntPtrInput
	// Name of the resource. Only used for importing.
	Name pulumi.StringPtrInput
}

func (UserLockPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*userLockPolicyState)(nil)).Elem()
}

type userLockPolicyArgs struct {
	// Enable User Lock Policy. Lock user after exceeding max failed login attempts.
	Enabled bool `pulumi:"enabled"`
	// Max failed login attempts.
	LoginAttempts int `pulumi:"loginAttempts"`
	// Name of the resource. Only used for importing.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a UserLockPolicy resource.
type UserLockPolicyArgs struct {
	// Enable User Lock Policy. Lock user after exceeding max failed login attempts.
	Enabled pulumi.BoolInput
	// Max failed login attempts.
	LoginAttempts pulumi.IntInput
	// Name of the resource. Only used for importing.
	Name pulumi.StringPtrInput
}

func (UserLockPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userLockPolicyArgs)(nil)).Elem()
}

type UserLockPolicyInput interface {
	pulumi.Input

	ToUserLockPolicyOutput() UserLockPolicyOutput
	ToUserLockPolicyOutputWithContext(ctx context.Context) UserLockPolicyOutput
}

func (*UserLockPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**UserLockPolicy)(nil)).Elem()
}

func (i *UserLockPolicy) ToUserLockPolicyOutput() UserLockPolicyOutput {
	return i.ToUserLockPolicyOutputWithContext(context.Background())
}

func (i *UserLockPolicy) ToUserLockPolicyOutputWithContext(ctx context.Context) UserLockPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLockPolicyOutput)
}

// UserLockPolicyArrayInput is an input type that accepts UserLockPolicyArray and UserLockPolicyArrayOutput values.
// You can construct a concrete instance of `UserLockPolicyArrayInput` via:
//
//	UserLockPolicyArray{ UserLockPolicyArgs{...} }
type UserLockPolicyArrayInput interface {
	pulumi.Input

	ToUserLockPolicyArrayOutput() UserLockPolicyArrayOutput
	ToUserLockPolicyArrayOutputWithContext(context.Context) UserLockPolicyArrayOutput
}

type UserLockPolicyArray []UserLockPolicyInput

func (UserLockPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserLockPolicy)(nil)).Elem()
}

func (i UserLockPolicyArray) ToUserLockPolicyArrayOutput() UserLockPolicyArrayOutput {
	return i.ToUserLockPolicyArrayOutputWithContext(context.Background())
}

func (i UserLockPolicyArray) ToUserLockPolicyArrayOutputWithContext(ctx context.Context) UserLockPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLockPolicyArrayOutput)
}

// UserLockPolicyMapInput is an input type that accepts UserLockPolicyMap and UserLockPolicyMapOutput values.
// You can construct a concrete instance of `UserLockPolicyMapInput` via:
//
//	UserLockPolicyMap{ "key": UserLockPolicyArgs{...} }
type UserLockPolicyMapInput interface {
	pulumi.Input

	ToUserLockPolicyMapOutput() UserLockPolicyMapOutput
	ToUserLockPolicyMapOutputWithContext(context.Context) UserLockPolicyMapOutput
}

type UserLockPolicyMap map[string]UserLockPolicyInput

func (UserLockPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserLockPolicy)(nil)).Elem()
}

func (i UserLockPolicyMap) ToUserLockPolicyMapOutput() UserLockPolicyMapOutput {
	return i.ToUserLockPolicyMapOutputWithContext(context.Background())
}

func (i UserLockPolicyMap) ToUserLockPolicyMapOutputWithContext(ctx context.Context) UserLockPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserLockPolicyMapOutput)
}

type UserLockPolicyOutput struct{ *pulumi.OutputState }

func (UserLockPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserLockPolicy)(nil)).Elem()
}

func (o UserLockPolicyOutput) ToUserLockPolicyOutput() UserLockPolicyOutput {
	return o
}

func (o UserLockPolicyOutput) ToUserLockPolicyOutputWithContext(ctx context.Context) UserLockPolicyOutput {
	return o
}

// Enable User Lock Policy. Lock user after exceeding max failed login attempts.
func (o UserLockPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *UserLockPolicy) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Max failed login attempts.
func (o UserLockPolicyOutput) LoginAttempts() pulumi.IntOutput {
	return o.ApplyT(func(v *UserLockPolicy) pulumi.IntOutput { return v.LoginAttempts }).(pulumi.IntOutput)
}

// Name of the resource. Only used for importing.
func (o UserLockPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *UserLockPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type UserLockPolicyArrayOutput struct{ *pulumi.OutputState }

func (UserLockPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserLockPolicy)(nil)).Elem()
}

func (o UserLockPolicyArrayOutput) ToUserLockPolicyArrayOutput() UserLockPolicyArrayOutput {
	return o
}

func (o UserLockPolicyArrayOutput) ToUserLockPolicyArrayOutputWithContext(ctx context.Context) UserLockPolicyArrayOutput {
	return o
}

func (o UserLockPolicyArrayOutput) Index(i pulumi.IntInput) UserLockPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserLockPolicy {
		return vs[0].([]*UserLockPolicy)[vs[1].(int)]
	}).(UserLockPolicyOutput)
}

type UserLockPolicyMapOutput struct{ *pulumi.OutputState }

func (UserLockPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserLockPolicy)(nil)).Elem()
}

func (o UserLockPolicyMapOutput) ToUserLockPolicyMapOutput() UserLockPolicyMapOutput {
	return o
}

func (o UserLockPolicyMapOutput) ToUserLockPolicyMapOutputWithContext(ctx context.Context) UserLockPolicyMapOutput {
	return o
}

func (o UserLockPolicyMapOutput) MapIndex(k pulumi.StringInput) UserLockPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserLockPolicy {
		return vs[0].(map[string]*UserLockPolicy)[vs[1].(string)]
	}).(UserLockPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserLockPolicyInput)(nil)).Elem(), &UserLockPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserLockPolicyArrayInput)(nil)).Elem(), UserLockPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserLockPolicyMapInput)(nil)).Elem(), UserLockPolicyMap{})
	pulumi.RegisterOutputType(UserLockPolicyOutput{})
	pulumi.RegisterOutputType(UserLockPolicyArrayOutput{})
	pulumi.RegisterOutputType(UserLockPolicyMapOutput{})
}
