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

// Provides an Artifactory Package Cleanup Policy resource. This resource enable system administrators to define and customize policies based on specific criteria for removing unused binaries from across their JFrog platform. See [Cleanup Policies](https://jfrog.com/help/r/jfrog-platform-administration-documentation/cleanup-policies) for more details.
//
// ~>Currently in beta and will be globally available in v7.98.x.
//
// ## Import
//
// ```sh
// $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy
// ```
//
// ```sh
// $ pulumi import artifactory:index/packageCleanupPolicy:PackageCleanupPolicy my-cleanup-policy my-policy:myproj
// ```
type PackageCleanupPolicy struct {
	pulumi.CustomResourceState

	// The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
	CronExpression pulumi.StringPtrOutput `pulumi:"cronExpression"`
	Description    pulumi.StringPtrOutput `pulumi:"description"`
	// Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
	DurationInMinutes pulumi.IntPtrOutput `pulumi:"durationInMinutes"`
	// Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
	Key pulumi.StringOutput `pulumi:"key"`
	// This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
	ProjectKey     pulumi.StringPtrOutput                   `pulumi:"projectKey"`
	SearchCriteria PackageCleanupPolicySearchCriteriaOutput `pulumi:"searchCriteria"`
	// Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
	SkipTrashcan pulumi.BoolOutput `pulumi:"skipTrashcan"`
}

// NewPackageCleanupPolicy registers a new resource with the given unique name, arguments, and options.
func NewPackageCleanupPolicy(ctx *pulumi.Context,
	name string, args *PackageCleanupPolicyArgs, opts ...pulumi.ResourceOption) (*PackageCleanupPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.SearchCriteria == nil {
		return nil, errors.New("invalid value for required argument 'SearchCriteria'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PackageCleanupPolicy
	err := ctx.RegisterResource("artifactory:index/packageCleanupPolicy:PackageCleanupPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPackageCleanupPolicy gets an existing PackageCleanupPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPackageCleanupPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PackageCleanupPolicyState, opts ...pulumi.ResourceOption) (*PackageCleanupPolicy, error) {
	var resource PackageCleanupPolicy
	err := ctx.ReadResource("artifactory:index/packageCleanupPolicy:PackageCleanupPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PackageCleanupPolicy resources.
type packageCleanupPolicyState struct {
	// The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
	CronExpression *string `pulumi:"cronExpression"`
	Description    *string `pulumi:"description"`
	// Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
	DurationInMinutes *int `pulumi:"durationInMinutes"`
	// Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
	Enabled *bool `pulumi:"enabled"`
	// Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
	Key *string `pulumi:"key"`
	// This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
	ProjectKey     *string                             `pulumi:"projectKey"`
	SearchCriteria *PackageCleanupPolicySearchCriteria `pulumi:"searchCriteria"`
	// Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
	SkipTrashcan *bool `pulumi:"skipTrashcan"`
}

type PackageCleanupPolicyState struct {
	// The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
	CronExpression pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	// Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
	DurationInMinutes pulumi.IntPtrInput
	// Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
	Enabled pulumi.BoolPtrInput
	// Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
	Key pulumi.StringPtrInput
	// This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
	ProjectKey     pulumi.StringPtrInput
	SearchCriteria PackageCleanupPolicySearchCriteriaPtrInput
	// Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
	SkipTrashcan pulumi.BoolPtrInput
}

func (PackageCleanupPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*packageCleanupPolicyState)(nil)).Elem()
}

type packageCleanupPolicyArgs struct {
	// The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
	CronExpression *string `pulumi:"cronExpression"`
	Description    *string `pulumi:"description"`
	// Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
	DurationInMinutes *int `pulumi:"durationInMinutes"`
	// Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
	Enabled *bool `pulumi:"enabled"`
	// Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
	Key string `pulumi:"key"`
	// This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
	ProjectKey     *string                            `pulumi:"projectKey"`
	SearchCriteria PackageCleanupPolicySearchCriteria `pulumi:"searchCriteria"`
	// Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
	SkipTrashcan *bool `pulumi:"skipTrashcan"`
}

// The set of arguments for constructing a PackageCleanupPolicy resource.
type PackageCleanupPolicyArgs struct {
	// The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
	CronExpression pulumi.StringPtrInput
	Description    pulumi.StringPtrInput
	// Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
	DurationInMinutes pulumi.IntPtrInput
	// Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
	Enabled pulumi.BoolPtrInput
	// Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
	Key pulumi.StringInput
	// This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
	ProjectKey     pulumi.StringPtrInput
	SearchCriteria PackageCleanupPolicySearchCriteriaInput
	// Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
	SkipTrashcan pulumi.BoolPtrInput
}

func (PackageCleanupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*packageCleanupPolicyArgs)(nil)).Elem()
}

type PackageCleanupPolicyInput interface {
	pulumi.Input

	ToPackageCleanupPolicyOutput() PackageCleanupPolicyOutput
	ToPackageCleanupPolicyOutputWithContext(ctx context.Context) PackageCleanupPolicyOutput
}

func (*PackageCleanupPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageCleanupPolicy)(nil)).Elem()
}

func (i *PackageCleanupPolicy) ToPackageCleanupPolicyOutput() PackageCleanupPolicyOutput {
	return i.ToPackageCleanupPolicyOutputWithContext(context.Background())
}

func (i *PackageCleanupPolicy) ToPackageCleanupPolicyOutputWithContext(ctx context.Context) PackageCleanupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageCleanupPolicyOutput)
}

// PackageCleanupPolicyArrayInput is an input type that accepts PackageCleanupPolicyArray and PackageCleanupPolicyArrayOutput values.
// You can construct a concrete instance of `PackageCleanupPolicyArrayInput` via:
//
//	PackageCleanupPolicyArray{ PackageCleanupPolicyArgs{...} }
type PackageCleanupPolicyArrayInput interface {
	pulumi.Input

	ToPackageCleanupPolicyArrayOutput() PackageCleanupPolicyArrayOutput
	ToPackageCleanupPolicyArrayOutputWithContext(context.Context) PackageCleanupPolicyArrayOutput
}

type PackageCleanupPolicyArray []PackageCleanupPolicyInput

func (PackageCleanupPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PackageCleanupPolicy)(nil)).Elem()
}

func (i PackageCleanupPolicyArray) ToPackageCleanupPolicyArrayOutput() PackageCleanupPolicyArrayOutput {
	return i.ToPackageCleanupPolicyArrayOutputWithContext(context.Background())
}

func (i PackageCleanupPolicyArray) ToPackageCleanupPolicyArrayOutputWithContext(ctx context.Context) PackageCleanupPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageCleanupPolicyArrayOutput)
}

// PackageCleanupPolicyMapInput is an input type that accepts PackageCleanupPolicyMap and PackageCleanupPolicyMapOutput values.
// You can construct a concrete instance of `PackageCleanupPolicyMapInput` via:
//
//	PackageCleanupPolicyMap{ "key": PackageCleanupPolicyArgs{...} }
type PackageCleanupPolicyMapInput interface {
	pulumi.Input

	ToPackageCleanupPolicyMapOutput() PackageCleanupPolicyMapOutput
	ToPackageCleanupPolicyMapOutputWithContext(context.Context) PackageCleanupPolicyMapOutput
}

type PackageCleanupPolicyMap map[string]PackageCleanupPolicyInput

func (PackageCleanupPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PackageCleanupPolicy)(nil)).Elem()
}

func (i PackageCleanupPolicyMap) ToPackageCleanupPolicyMapOutput() PackageCleanupPolicyMapOutput {
	return i.ToPackageCleanupPolicyMapOutputWithContext(context.Background())
}

func (i PackageCleanupPolicyMap) ToPackageCleanupPolicyMapOutputWithContext(ctx context.Context) PackageCleanupPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PackageCleanupPolicyMapOutput)
}

type PackageCleanupPolicyOutput struct{ *pulumi.OutputState }

func (PackageCleanupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PackageCleanupPolicy)(nil)).Elem()
}

func (o PackageCleanupPolicyOutput) ToPackageCleanupPolicyOutput() PackageCleanupPolicyOutput {
	return o
}

func (o PackageCleanupPolicyOutput) ToPackageCleanupPolicyOutputWithContext(ctx context.Context) PackageCleanupPolicyOutput {
	return o
}

// The Cron expression that sets the schedule of policy execution. For example, `0 0 2 * * ?` executes the policy every day at 02:00 AM. The minimum recurrent time for policy execution is 6 hours.
func (o PackageCleanupPolicyOutput) CronExpression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.StringPtrOutput { return v.CronExpression }).(pulumi.StringPtrOutput)
}

func (o PackageCleanupPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Enable and select the maximum duration for policy execution. Note: using this setting can cause the policy to stop before completion.
func (o PackageCleanupPolicyOutput) DurationInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.IntPtrOutput { return v.DurationInMinutes }).(pulumi.IntPtrOutput)
}

// Enables or disabled the package cleanup policy. This allows the user to run the policy manually. If a policy has a valid cron expression, then it will be scheduled for execution based on it. If a policy is disabled, its future executions will be unscheduled. Defaults to `true`
func (o PackageCleanupPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

// Policy key. It has to be unique. It should not be used for other policies and configuration entities like archive policies, key pairs, repo layouts, property sets, backups, proxies, reverse proxies etc. A minimum of three characters is required and can include letters, numbers, underscore and hyphen.
func (o PackageCleanupPolicyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// This attribute is used only for project-level cleanup policies, it is not used for global-level policies.
func (o PackageCleanupPolicyOutput) ProjectKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.StringPtrOutput { return v.ProjectKey }).(pulumi.StringPtrOutput)
}

func (o PackageCleanupPolicyOutput) SearchCriteria() PackageCleanupPolicySearchCriteriaOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) PackageCleanupPolicySearchCriteriaOutput { return v.SearchCriteria }).(PackageCleanupPolicySearchCriteriaOutput)
}

// Enabling this setting results in packages being permanently deleted from Artifactory after the cleanup policy is executed instead of going to the Trash Can repository. Defaults to `false`.
func (o PackageCleanupPolicyOutput) SkipTrashcan() pulumi.BoolOutput {
	return o.ApplyT(func(v *PackageCleanupPolicy) pulumi.BoolOutput { return v.SkipTrashcan }).(pulumi.BoolOutput)
}

type PackageCleanupPolicyArrayOutput struct{ *pulumi.OutputState }

func (PackageCleanupPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PackageCleanupPolicy)(nil)).Elem()
}

func (o PackageCleanupPolicyArrayOutput) ToPackageCleanupPolicyArrayOutput() PackageCleanupPolicyArrayOutput {
	return o
}

func (o PackageCleanupPolicyArrayOutput) ToPackageCleanupPolicyArrayOutputWithContext(ctx context.Context) PackageCleanupPolicyArrayOutput {
	return o
}

func (o PackageCleanupPolicyArrayOutput) Index(i pulumi.IntInput) PackageCleanupPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PackageCleanupPolicy {
		return vs[0].([]*PackageCleanupPolicy)[vs[1].(int)]
	}).(PackageCleanupPolicyOutput)
}

type PackageCleanupPolicyMapOutput struct{ *pulumi.OutputState }

func (PackageCleanupPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PackageCleanupPolicy)(nil)).Elem()
}

func (o PackageCleanupPolicyMapOutput) ToPackageCleanupPolicyMapOutput() PackageCleanupPolicyMapOutput {
	return o
}

func (o PackageCleanupPolicyMapOutput) ToPackageCleanupPolicyMapOutputWithContext(ctx context.Context) PackageCleanupPolicyMapOutput {
	return o
}

func (o PackageCleanupPolicyMapOutput) MapIndex(k pulumi.StringInput) PackageCleanupPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PackageCleanupPolicy {
		return vs[0].(map[string]*PackageCleanupPolicy)[vs[1].(string)]
	}).(PackageCleanupPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PackageCleanupPolicyInput)(nil)).Elem(), &PackageCleanupPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageCleanupPolicyArrayInput)(nil)).Elem(), PackageCleanupPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PackageCleanupPolicyMapInput)(nil)).Elem(), PackageCleanupPolicyMap{})
	pulumi.RegisterOutputType(PackageCleanupPolicyOutput{})
	pulumi.RegisterOutputType(PackageCleanupPolicyArrayOutput{})
	pulumi.RegisterOutputType(PackageCleanupPolicyMapOutput{})
}
