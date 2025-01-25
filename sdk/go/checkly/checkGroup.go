// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"errors"
	"github.com/checkly/pulumi-checkly/sdk/go/checkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Check groups allow  you to group together a set of related checks, which can also share default settings for various attributes.
type CheckGroup struct {
	pulumi.CustomResourceState

	// Determines if the checks in the group are running or not.
	Activated                 pulumi.BoolOutput                             `pulumi:"activated"`
	AlertChannelSubscriptions CheckGroupAlertChannelSubscriptionArrayOutput `pulumi:"alertChannelSubscriptions"`
	AlertSettings             CheckGroupAlertSettingsOutput                 `pulumi:"alertSettings"`
	ApiCheckDefaults          CheckGroupApiCheckDefaultsOutput              `pulumi:"apiCheckDefaults"`
	// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
	Concurrency pulumi.IntOutput `pulumi:"concurrency"`
	// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
	//
	// Deprecated: The property `doubleCheck` is deprecated and will be removed in a future version. To enable retries for failed check runs, use the `retryStrategy` property instead.
	DoubleCheck pulumi.BoolPtrOutput `pulumi:"doubleCheck"`
	// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden, add secret = true to create a secret variable. These are only relevant for browser checks. Use global environment variables whenever possible.
	EnvironmentVariable CheckGroupEnvironmentVariableArrayOutput `pulumi:"environmentVariable"`
	// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
	//
	// Deprecated: The property `environmentVariables` is deprecated and will be removed in a future version. Consider using the new `environmentVariable` list.
	EnvironmentVariables pulumi.StringMapOutput `pulumi:"environmentVariables"`
	// A valid piece of Node.js code to run in the setup phase of an API check in this group.
	LocalSetupScript pulumi.StringPtrOutput `pulumi:"localSetupScript"`
	// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
	LocalTeardownScript pulumi.StringPtrOutput `pulumi:"localTeardownScript"`
	// An array of one or more data center locations where to run the checks.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
	Muted pulumi.BoolPtrOutput `pulumi:"muted"`
	// The name of the check group.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of one or more private locations slugs.
	PrivateLocations pulumi.StringArrayOutput `pulumi:"privateLocations"`
	// A strategy for retrying failed check runs.
	RetryStrategy CheckGroupRetryStrategyOutput `pulumi:"retryStrategy"`
	// Determines if the checks in the group should run in all selected locations in parallel or round-robin.
	RunParallel pulumi.BoolPtrOutput `pulumi:"runParallel"`
	// The id of the runtime to use for this group.
	RuntimeId pulumi.StringPtrOutput `pulumi:"runtimeId"`
	// An ID reference to a snippet to use in the setup phase of an API check.
	SetupSnippetId pulumi.IntPtrOutput `pulumi:"setupSnippetId"`
	// Tags for organizing and filtering checks.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// An ID reference to a snippet to use in the teardown phase of an API check.
	TeardownSnippetId pulumi.IntPtrOutput `pulumi:"teardownSnippetId"`
	// When true, the account level alert settings will be used, not the alert setting defined on this check group.
	UseGlobalAlertSettings pulumi.BoolPtrOutput `pulumi:"useGlobalAlertSettings"`
}

// NewCheckGroup registers a new resource with the given unique name, arguments, and options.
func NewCheckGroup(ctx *pulumi.Context,
	name string, args *CheckGroupArgs, opts ...pulumi.ResourceOption) (*CheckGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Activated == nil {
		return nil, errors.New("invalid value for required argument 'Activated'")
	}
	if args.Concurrency == nil {
		return nil, errors.New("invalid value for required argument 'Concurrency'")
	}
	if args.ApiCheckDefaults != nil {
		args.ApiCheckDefaults = args.ApiCheckDefaults.ToCheckGroupApiCheckDefaultsPtrOutput().ApplyT(func(v *CheckGroupApiCheckDefaults) *CheckGroupApiCheckDefaults { return v.Defaults() }).(CheckGroupApiCheckDefaultsPtrOutput)
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CheckGroup
	err := ctx.RegisterResource("checkly:index/checkGroup:CheckGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCheckGroup gets an existing CheckGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCheckGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CheckGroupState, opts ...pulumi.ResourceOption) (*CheckGroup, error) {
	var resource CheckGroup
	err := ctx.ReadResource("checkly:index/checkGroup:CheckGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CheckGroup resources.
type checkGroupState struct {
	// Determines if the checks in the group are running or not.
	Activated                 *bool                                `pulumi:"activated"`
	AlertChannelSubscriptions []CheckGroupAlertChannelSubscription `pulumi:"alertChannelSubscriptions"`
	AlertSettings             *CheckGroupAlertSettings             `pulumi:"alertSettings"`
	ApiCheckDefaults          *CheckGroupApiCheckDefaults          `pulumi:"apiCheckDefaults"`
	// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
	Concurrency *int `pulumi:"concurrency"`
	// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
	//
	// Deprecated: The property `doubleCheck` is deprecated and will be removed in a future version. To enable retries for failed check runs, use the `retryStrategy` property instead.
	DoubleCheck *bool `pulumi:"doubleCheck"`
	// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden, add secret = true to create a secret variable. These are only relevant for browser checks. Use global environment variables whenever possible.
	EnvironmentVariable []CheckGroupEnvironmentVariable `pulumi:"environmentVariable"`
	// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
	//
	// Deprecated: The property `environmentVariables` is deprecated and will be removed in a future version. Consider using the new `environmentVariable` list.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// A valid piece of Node.js code to run in the setup phase of an API check in this group.
	LocalSetupScript *string `pulumi:"localSetupScript"`
	// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
	LocalTeardownScript *string `pulumi:"localTeardownScript"`
	// An array of one or more data center locations where to run the checks.
	Locations []string `pulumi:"locations"`
	// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
	Muted *bool `pulumi:"muted"`
	// The name of the check group.
	Name *string `pulumi:"name"`
	// An array of one or more private locations slugs.
	PrivateLocations []string `pulumi:"privateLocations"`
	// A strategy for retrying failed check runs.
	RetryStrategy *CheckGroupRetryStrategy `pulumi:"retryStrategy"`
	// Determines if the checks in the group should run in all selected locations in parallel or round-robin.
	RunParallel *bool `pulumi:"runParallel"`
	// The id of the runtime to use for this group.
	RuntimeId *string `pulumi:"runtimeId"`
	// An ID reference to a snippet to use in the setup phase of an API check.
	SetupSnippetId *int `pulumi:"setupSnippetId"`
	// Tags for organizing and filtering checks.
	Tags []string `pulumi:"tags"`
	// An ID reference to a snippet to use in the teardown phase of an API check.
	TeardownSnippetId *int `pulumi:"teardownSnippetId"`
	// When true, the account level alert settings will be used, not the alert setting defined on this check group.
	UseGlobalAlertSettings *bool `pulumi:"useGlobalAlertSettings"`
}

type CheckGroupState struct {
	// Determines if the checks in the group are running or not.
	Activated                 pulumi.BoolPtrInput
	AlertChannelSubscriptions CheckGroupAlertChannelSubscriptionArrayInput
	AlertSettings             CheckGroupAlertSettingsPtrInput
	ApiCheckDefaults          CheckGroupApiCheckDefaultsPtrInput
	// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
	Concurrency pulumi.IntPtrInput
	// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
	//
	// Deprecated: The property `doubleCheck` is deprecated and will be removed in a future version. To enable retries for failed check runs, use the `retryStrategy` property instead.
	DoubleCheck pulumi.BoolPtrInput
	// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden, add secret = true to create a secret variable. These are only relevant for browser checks. Use global environment variables whenever possible.
	EnvironmentVariable CheckGroupEnvironmentVariableArrayInput
	// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
	//
	// Deprecated: The property `environmentVariables` is deprecated and will be removed in a future version. Consider using the new `environmentVariable` list.
	EnvironmentVariables pulumi.StringMapInput
	// A valid piece of Node.js code to run in the setup phase of an API check in this group.
	LocalSetupScript pulumi.StringPtrInput
	// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
	LocalTeardownScript pulumi.StringPtrInput
	// An array of one or more data center locations where to run the checks.
	Locations pulumi.StringArrayInput
	// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
	Muted pulumi.BoolPtrInput
	// The name of the check group.
	Name pulumi.StringPtrInput
	// An array of one or more private locations slugs.
	PrivateLocations pulumi.StringArrayInput
	// A strategy for retrying failed check runs.
	RetryStrategy CheckGroupRetryStrategyPtrInput
	// Determines if the checks in the group should run in all selected locations in parallel or round-robin.
	RunParallel pulumi.BoolPtrInput
	// The id of the runtime to use for this group.
	RuntimeId pulumi.StringPtrInput
	// An ID reference to a snippet to use in the setup phase of an API check.
	SetupSnippetId pulumi.IntPtrInput
	// Tags for organizing and filtering checks.
	Tags pulumi.StringArrayInput
	// An ID reference to a snippet to use in the teardown phase of an API check.
	TeardownSnippetId pulumi.IntPtrInput
	// When true, the account level alert settings will be used, not the alert setting defined on this check group.
	UseGlobalAlertSettings pulumi.BoolPtrInput
}

func (CheckGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*checkGroupState)(nil)).Elem()
}

type checkGroupArgs struct {
	// Determines if the checks in the group are running or not.
	Activated                 bool                                 `pulumi:"activated"`
	AlertChannelSubscriptions []CheckGroupAlertChannelSubscription `pulumi:"alertChannelSubscriptions"`
	AlertSettings             *CheckGroupAlertSettings             `pulumi:"alertSettings"`
	ApiCheckDefaults          *CheckGroupApiCheckDefaults          `pulumi:"apiCheckDefaults"`
	// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
	Concurrency int `pulumi:"concurrency"`
	// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
	//
	// Deprecated: The property `doubleCheck` is deprecated and will be removed in a future version. To enable retries for failed check runs, use the `retryStrategy` property instead.
	DoubleCheck *bool `pulumi:"doubleCheck"`
	// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden, add secret = true to create a secret variable. These are only relevant for browser checks. Use global environment variables whenever possible.
	EnvironmentVariable []CheckGroupEnvironmentVariable `pulumi:"environmentVariable"`
	// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
	//
	// Deprecated: The property `environmentVariables` is deprecated and will be removed in a future version. Consider using the new `environmentVariable` list.
	EnvironmentVariables map[string]string `pulumi:"environmentVariables"`
	// A valid piece of Node.js code to run in the setup phase of an API check in this group.
	LocalSetupScript *string `pulumi:"localSetupScript"`
	// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
	LocalTeardownScript *string `pulumi:"localTeardownScript"`
	// An array of one or more data center locations where to run the checks.
	Locations []string `pulumi:"locations"`
	// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
	Muted *bool `pulumi:"muted"`
	// The name of the check group.
	Name *string `pulumi:"name"`
	// An array of one or more private locations slugs.
	PrivateLocations []string `pulumi:"privateLocations"`
	// A strategy for retrying failed check runs.
	RetryStrategy *CheckGroupRetryStrategy `pulumi:"retryStrategy"`
	// Determines if the checks in the group should run in all selected locations in parallel or round-robin.
	RunParallel *bool `pulumi:"runParallel"`
	// The id of the runtime to use for this group.
	RuntimeId *string `pulumi:"runtimeId"`
	// An ID reference to a snippet to use in the setup phase of an API check.
	SetupSnippetId *int `pulumi:"setupSnippetId"`
	// Tags for organizing and filtering checks.
	Tags []string `pulumi:"tags"`
	// An ID reference to a snippet to use in the teardown phase of an API check.
	TeardownSnippetId *int `pulumi:"teardownSnippetId"`
	// When true, the account level alert settings will be used, not the alert setting defined on this check group.
	UseGlobalAlertSettings *bool `pulumi:"useGlobalAlertSettings"`
}

// The set of arguments for constructing a CheckGroup resource.
type CheckGroupArgs struct {
	// Determines if the checks in the group are running or not.
	Activated                 pulumi.BoolInput
	AlertChannelSubscriptions CheckGroupAlertChannelSubscriptionArrayInput
	AlertSettings             CheckGroupAlertSettingsPtrInput
	ApiCheckDefaults          CheckGroupApiCheckDefaultsPtrInput
	// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
	Concurrency pulumi.IntInput
	// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
	//
	// Deprecated: The property `doubleCheck` is deprecated and will be removed in a future version. To enable retries for failed check runs, use the `retryStrategy` property instead.
	DoubleCheck pulumi.BoolPtrInput
	// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden, add secret = true to create a secret variable. These are only relevant for browser checks. Use global environment variables whenever possible.
	EnvironmentVariable CheckGroupEnvironmentVariableArrayInput
	// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
	//
	// Deprecated: The property `environmentVariables` is deprecated and will be removed in a future version. Consider using the new `environmentVariable` list.
	EnvironmentVariables pulumi.StringMapInput
	// A valid piece of Node.js code to run in the setup phase of an API check in this group.
	LocalSetupScript pulumi.StringPtrInput
	// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
	LocalTeardownScript pulumi.StringPtrInput
	// An array of one or more data center locations where to run the checks.
	Locations pulumi.StringArrayInput
	// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
	Muted pulumi.BoolPtrInput
	// The name of the check group.
	Name pulumi.StringPtrInput
	// An array of one or more private locations slugs.
	PrivateLocations pulumi.StringArrayInput
	// A strategy for retrying failed check runs.
	RetryStrategy CheckGroupRetryStrategyPtrInput
	// Determines if the checks in the group should run in all selected locations in parallel or round-robin.
	RunParallel pulumi.BoolPtrInput
	// The id of the runtime to use for this group.
	RuntimeId pulumi.StringPtrInput
	// An ID reference to a snippet to use in the setup phase of an API check.
	SetupSnippetId pulumi.IntPtrInput
	// Tags for organizing and filtering checks.
	Tags pulumi.StringArrayInput
	// An ID reference to a snippet to use in the teardown phase of an API check.
	TeardownSnippetId pulumi.IntPtrInput
	// When true, the account level alert settings will be used, not the alert setting defined on this check group.
	UseGlobalAlertSettings pulumi.BoolPtrInput
}

func (CheckGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*checkGroupArgs)(nil)).Elem()
}

type CheckGroupInput interface {
	pulumi.Input

	ToCheckGroupOutput() CheckGroupOutput
	ToCheckGroupOutputWithContext(ctx context.Context) CheckGroupOutput
}

func (*CheckGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckGroup)(nil)).Elem()
}

func (i *CheckGroup) ToCheckGroupOutput() CheckGroupOutput {
	return i.ToCheckGroupOutputWithContext(context.Background())
}

func (i *CheckGroup) ToCheckGroupOutputWithContext(ctx context.Context) CheckGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckGroupOutput)
}

// CheckGroupArrayInput is an input type that accepts CheckGroupArray and CheckGroupArrayOutput values.
// You can construct a concrete instance of `CheckGroupArrayInput` via:
//
//	CheckGroupArray{ CheckGroupArgs{...} }
type CheckGroupArrayInput interface {
	pulumi.Input

	ToCheckGroupArrayOutput() CheckGroupArrayOutput
	ToCheckGroupArrayOutputWithContext(context.Context) CheckGroupArrayOutput
}

type CheckGroupArray []CheckGroupInput

func (CheckGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckGroup)(nil)).Elem()
}

func (i CheckGroupArray) ToCheckGroupArrayOutput() CheckGroupArrayOutput {
	return i.ToCheckGroupArrayOutputWithContext(context.Background())
}

func (i CheckGroupArray) ToCheckGroupArrayOutputWithContext(ctx context.Context) CheckGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckGroupArrayOutput)
}

// CheckGroupMapInput is an input type that accepts CheckGroupMap and CheckGroupMapOutput values.
// You can construct a concrete instance of `CheckGroupMapInput` via:
//
//	CheckGroupMap{ "key": CheckGroupArgs{...} }
type CheckGroupMapInput interface {
	pulumi.Input

	ToCheckGroupMapOutput() CheckGroupMapOutput
	ToCheckGroupMapOutputWithContext(context.Context) CheckGroupMapOutput
}

type CheckGroupMap map[string]CheckGroupInput

func (CheckGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckGroup)(nil)).Elem()
}

func (i CheckGroupMap) ToCheckGroupMapOutput() CheckGroupMapOutput {
	return i.ToCheckGroupMapOutputWithContext(context.Background())
}

func (i CheckGroupMap) ToCheckGroupMapOutputWithContext(ctx context.Context) CheckGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CheckGroupMapOutput)
}

type CheckGroupOutput struct{ *pulumi.OutputState }

func (CheckGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CheckGroup)(nil)).Elem()
}

func (o CheckGroupOutput) ToCheckGroupOutput() CheckGroupOutput {
	return o
}

func (o CheckGroupOutput) ToCheckGroupOutputWithContext(ctx context.Context) CheckGroupOutput {
	return o
}

// Determines if the checks in the group are running or not.
func (o CheckGroupOutput) Activated() pulumi.BoolOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.BoolOutput { return v.Activated }).(pulumi.BoolOutput)
}

func (o CheckGroupOutput) AlertChannelSubscriptions() CheckGroupAlertChannelSubscriptionArrayOutput {
	return o.ApplyT(func(v *CheckGroup) CheckGroupAlertChannelSubscriptionArrayOutput { return v.AlertChannelSubscriptions }).(CheckGroupAlertChannelSubscriptionArrayOutput)
}

func (o CheckGroupOutput) AlertSettings() CheckGroupAlertSettingsOutput {
	return o.ApplyT(func(v *CheckGroup) CheckGroupAlertSettingsOutput { return v.AlertSettings }).(CheckGroupAlertSettingsOutput)
}

func (o CheckGroupOutput) ApiCheckDefaults() CheckGroupApiCheckDefaultsOutput {
	return o.ApplyT(func(v *CheckGroup) CheckGroupApiCheckDefaultsOutput { return v.ApiCheckDefaults }).(CheckGroupApiCheckDefaultsOutput)
}

// Determines how many checks are run concurrently when triggering a check group from CI/CD or through the API.
func (o CheckGroupOutput) Concurrency() pulumi.IntOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.IntOutput { return v.Concurrency }).(pulumi.IntOutput)
}

// Setting this to `true` will trigger a retry when a check fails from the failing region and another, randomly selected region before marking the check as failed.
//
// Deprecated: The property `doubleCheck` is deprecated and will be removed in a future version. To enable retries for failed check runs, use the `retryStrategy` property instead.
func (o CheckGroupOutput) DoubleCheck() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.BoolPtrOutput { return v.DoubleCheck }).(pulumi.BoolPtrOutput)
}

// Key/value pairs for setting environment variables during check execution, add locked = true to keep value hidden, add secret = true to create a secret variable. These are only relevant for browser checks. Use global environment variables whenever possible.
func (o CheckGroupOutput) EnvironmentVariable() CheckGroupEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v *CheckGroup) CheckGroupEnvironmentVariableArrayOutput { return v.EnvironmentVariable }).(CheckGroupEnvironmentVariableArrayOutput)
}

// Key/value pairs for setting environment variables during check execution. These are only relevant for browser checks. Use global environment variables whenever possible.
//
// Deprecated: The property `environmentVariables` is deprecated and will be removed in a future version. Consider using the new `environmentVariable` list.
func (o CheckGroupOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringMapOutput { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

// A valid piece of Node.js code to run in the setup phase of an API check in this group.
func (o CheckGroupOutput) LocalSetupScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringPtrOutput { return v.LocalSetupScript }).(pulumi.StringPtrOutput)
}

// A valid piece of Node.js code to run in the teardown phase of an API check in this group.
func (o CheckGroupOutput) LocalTeardownScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringPtrOutput { return v.LocalTeardownScript }).(pulumi.StringPtrOutput)
}

// An array of one or more data center locations where to run the checks.
func (o CheckGroupOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringArrayOutput { return v.Locations }).(pulumi.StringArrayOutput)
}

// Determines if any notifications will be sent out when a check in this group fails and/or recovers.
func (o CheckGroupOutput) Muted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.BoolPtrOutput { return v.Muted }).(pulumi.BoolPtrOutput)
}

// The name of the check group.
func (o CheckGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of one or more private locations slugs.
func (o CheckGroupOutput) PrivateLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringArrayOutput { return v.PrivateLocations }).(pulumi.StringArrayOutput)
}

// A strategy for retrying failed check runs.
func (o CheckGroupOutput) RetryStrategy() CheckGroupRetryStrategyOutput {
	return o.ApplyT(func(v *CheckGroup) CheckGroupRetryStrategyOutput { return v.RetryStrategy }).(CheckGroupRetryStrategyOutput)
}

// Determines if the checks in the group should run in all selected locations in parallel or round-robin.
func (o CheckGroupOutput) RunParallel() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.BoolPtrOutput { return v.RunParallel }).(pulumi.BoolPtrOutput)
}

// The id of the runtime to use for this group.
func (o CheckGroupOutput) RuntimeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringPtrOutput { return v.RuntimeId }).(pulumi.StringPtrOutput)
}

// An ID reference to a snippet to use in the setup phase of an API check.
func (o CheckGroupOutput) SetupSnippetId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.IntPtrOutput { return v.SetupSnippetId }).(pulumi.IntPtrOutput)
}

// Tags for organizing and filtering checks.
func (o CheckGroupOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// An ID reference to a snippet to use in the teardown phase of an API check.
func (o CheckGroupOutput) TeardownSnippetId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.IntPtrOutput { return v.TeardownSnippetId }).(pulumi.IntPtrOutput)
}

// When true, the account level alert settings will be used, not the alert setting defined on this check group.
func (o CheckGroupOutput) UseGlobalAlertSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CheckGroup) pulumi.BoolPtrOutput { return v.UseGlobalAlertSettings }).(pulumi.BoolPtrOutput)
}

type CheckGroupArrayOutput struct{ *pulumi.OutputState }

func (CheckGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CheckGroup)(nil)).Elem()
}

func (o CheckGroupArrayOutput) ToCheckGroupArrayOutput() CheckGroupArrayOutput {
	return o
}

func (o CheckGroupArrayOutput) ToCheckGroupArrayOutputWithContext(ctx context.Context) CheckGroupArrayOutput {
	return o
}

func (o CheckGroupArrayOutput) Index(i pulumi.IntInput) CheckGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CheckGroup {
		return vs[0].([]*CheckGroup)[vs[1].(int)]
	}).(CheckGroupOutput)
}

type CheckGroupMapOutput struct{ *pulumi.OutputState }

func (CheckGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CheckGroup)(nil)).Elem()
}

func (o CheckGroupMapOutput) ToCheckGroupMapOutput() CheckGroupMapOutput {
	return o
}

func (o CheckGroupMapOutput) ToCheckGroupMapOutputWithContext(ctx context.Context) CheckGroupMapOutput {
	return o
}

func (o CheckGroupMapOutput) MapIndex(k pulumi.StringInput) CheckGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CheckGroup {
		return vs[0].(map[string]*CheckGroup)[vs[1].(string)]
	}).(CheckGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CheckGroupInput)(nil)).Elem(), &CheckGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckGroupArrayInput)(nil)).Elem(), CheckGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CheckGroupMapInput)(nil)).Elem(), CheckGroupMap{})
	pulumi.RegisterOutputType(CheckGroupOutput{})
	pulumi.RegisterOutputType(CheckGroupArrayOutput{})
	pulumi.RegisterOutputType(CheckGroupMapOutput{})
}
