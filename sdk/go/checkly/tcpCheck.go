// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"errors"
	"github.com/checkly/pulumi-checkly/sdk/v2/go/checkly/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// TCP checks allow you to monitor remote endpoints at a lower level.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/checkly/pulumi-checkly/sdk/v2/go/checkly"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Basic TCP Check
//			_, err := checkly.NewTcpCheck(ctx, "example-tcp-check", &checkly.TcpCheckArgs{
//				Name:                   pulumi.String("Example TCP check"),
//				Activated:              pulumi.Bool(true),
//				ShouldFail:             pulumi.Bool(false),
//				Frequency:              pulumi.Int(1),
//				UseGlobalAlertSettings: pulumi.Bool(true),
//				Locations: pulumi.StringArray{
//					pulumi.String("us-west-1"),
//				},
//				Request: &checkly.TcpCheckRequestArgs{
//					Hostname: pulumi.String("api.checklyhq.com"),
//					Port:     pulumi.Int(80),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			// A more complex example using assertions and setting alerts
//			_, err = checkly.NewTcpCheck(ctx, "example-tcp-check-2", &checkly.TcpCheckArgs{
//				Name:                 pulumi.String("Example TCP check 2"),
//				Activated:            pulumi.Bool(true),
//				ShouldFail:           pulumi.Bool(true),
//				Frequency:            pulumi.Int(1),
//				DegradedResponseTime: pulumi.Int(5000),
//				MaxResponseTime:      pulumi.Int(10000),
//				Locations: pulumi.StringArray{
//					pulumi.String("us-west-1"),
//					pulumi.String("ap-northeast-1"),
//					pulumi.String("ap-south-1"),
//				},
//				AlertSettings: &checkly.TcpCheckAlertSettingsArgs{
//					EscalationType: pulumi.String("RUN_BASED"),
//					RunBasedEscalations: checkly.TcpCheckAlertSettingsRunBasedEscalationArray{
//						&checkly.TcpCheckAlertSettingsRunBasedEscalationArgs{
//							FailedRunThreshold: pulumi.Int(1),
//						},
//					},
//					Reminders: checkly.TcpCheckAlertSettingsReminderArray{
//						&checkly.TcpCheckAlertSettingsReminderArgs{
//							Amount: pulumi.Int(1),
//						},
//					},
//				},
//				RetryStrategy: &checkly.TcpCheckRetryStrategyArgs{
//					Type:               pulumi.String("FIXED"),
//					BaseBackoffSeconds: pulumi.Int(60),
//					MaxDurationSeconds: pulumi.Int(600),
//					MaxRetries:         pulumi.Int(3),
//					SameRegion:         pulumi.Bool(false),
//				},
//				Request: &checkly.TcpCheckRequestArgs{
//					Hostname: pulumi.String("api.checklyhq.com"),
//					Port:     pulumi.Int(80),
//					Data:     pulumi.String("hello"),
//					Assertions: checkly.TcpCheckRequestAssertionArray{
//						&checkly.TcpCheckRequestAssertionArgs{
//							Source:     pulumi.String("RESPONSE_DATA"),
//							Property:   pulumi.String(""),
//							Comparison: pulumi.String("CONTAINS"),
//							Target:     pulumi.String("welcome"),
//						},
//						&checkly.TcpCheckRequestAssertionArgs{
//							Source:     pulumi.String("RESPONSE_TIME"),
//							Property:   pulumi.String(""),
//							Comparison: pulumi.String("LESS_THAN"),
//							Target:     pulumi.String("2000"),
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
type TcpCheck struct {
	pulumi.CustomResourceState

	// Determines if the check is running or not. Possible values `true`, and `false`.
	Activated pulumi.BoolOutput `pulumi:"activated"`
	// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
	AlertChannelSubscriptions TcpCheckAlertChannelSubscriptionArrayOutput `pulumi:"alertChannelSubscriptions"`
	AlertSettings             TcpCheckAlertSettingsOutput                 `pulumi:"alertSettings"`
	// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 5000. (Default `4000`).
	DegradedResponseTime pulumi.IntPtrOutput `pulumi:"degradedResponseTime"`
	// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
	Frequency pulumi.IntOutput `pulumi:"frequency"`
	// To create a high frequency check, the property `frequency` must be `0` and `frequencyOffset` can be `10`, `20` or `30`.
	FrequencyOffset pulumi.IntPtrOutput `pulumi:"frequencyOffset"`
	// The id of the check group this check is part of.
	GroupId pulumi.IntPtrOutput `pulumi:"groupId"`
	// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
	GroupOrder pulumi.IntPtrOutput `pulumi:"groupOrder"`
	// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 5000. (Default `5000`).
	MaxResponseTime pulumi.IntPtrOutput `pulumi:"maxResponseTime"`
	// Determines if any notifications will be sent out when a check fails/degrades/recovers.
	Muted pulumi.BoolPtrOutput `pulumi:"muted"`
	// The name of the check.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of one or more private locations slugs.
	PrivateLocations pulumi.StringArrayOutput `pulumi:"privateLocations"`
	// The parameters for the TCP connection.
	Request TcpCheckRequestOutput `pulumi:"request"`
	// A strategy for retrying failed check runs.
	RetryStrategy TcpCheckRetryStrategyOutput `pulumi:"retryStrategy"`
	// Determines if the check should run in all selected locations in parallel or round-robin.
	RunParallel pulumi.BoolPtrOutput `pulumi:"runParallel"`
	// The ID of the runtime to use for this check.
	RuntimeId pulumi.StringPtrOutput `pulumi:"runtimeId"`
	// Allows to invert the behaviour of when a check is considered to fail.
	ShouldFail pulumi.BoolPtrOutput `pulumi:"shouldFail"`
	// A list of tags for organizing and filtering checks.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// When true, the account level alert settings will be used, not the alert setting defined on this check.
	UseGlobalAlertSettings pulumi.BoolPtrOutput `pulumi:"useGlobalAlertSettings"`
}

// NewTcpCheck registers a new resource with the given unique name, arguments, and options.
func NewTcpCheck(ctx *pulumi.Context,
	name string, args *TcpCheckArgs, opts ...pulumi.ResourceOption) (*TcpCheck, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Activated == nil {
		return nil, errors.New("invalid value for required argument 'Activated'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Request == nil {
		return nil, errors.New("invalid value for required argument 'Request'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TcpCheck
	err := ctx.RegisterResource("checkly:index/tcpCheck:TcpCheck", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTcpCheck gets an existing TcpCheck resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTcpCheck(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TcpCheckState, opts ...pulumi.ResourceOption) (*TcpCheck, error) {
	var resource TcpCheck
	err := ctx.ReadResource("checkly:index/tcpCheck:TcpCheck", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TcpCheck resources.
type tcpCheckState struct {
	// Determines if the check is running or not. Possible values `true`, and `false`.
	Activated *bool `pulumi:"activated"`
	// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
	AlertChannelSubscriptions []TcpCheckAlertChannelSubscription `pulumi:"alertChannelSubscriptions"`
	AlertSettings             *TcpCheckAlertSettings             `pulumi:"alertSettings"`
	// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 5000. (Default `4000`).
	DegradedResponseTime *int `pulumi:"degradedResponseTime"`
	// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
	Frequency *int `pulumi:"frequency"`
	// To create a high frequency check, the property `frequency` must be `0` and `frequencyOffset` can be `10`, `20` or `30`.
	FrequencyOffset *int `pulumi:"frequencyOffset"`
	// The id of the check group this check is part of.
	GroupId *int `pulumi:"groupId"`
	// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
	GroupOrder *int `pulumi:"groupOrder"`
	// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
	Locations []string `pulumi:"locations"`
	// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 5000. (Default `5000`).
	MaxResponseTime *int `pulumi:"maxResponseTime"`
	// Determines if any notifications will be sent out when a check fails/degrades/recovers.
	Muted *bool `pulumi:"muted"`
	// The name of the check.
	Name *string `pulumi:"name"`
	// An array of one or more private locations slugs.
	PrivateLocations []string `pulumi:"privateLocations"`
	// The parameters for the TCP connection.
	Request *TcpCheckRequest `pulumi:"request"`
	// A strategy for retrying failed check runs.
	RetryStrategy *TcpCheckRetryStrategy `pulumi:"retryStrategy"`
	// Determines if the check should run in all selected locations in parallel or round-robin.
	RunParallel *bool `pulumi:"runParallel"`
	// The ID of the runtime to use for this check.
	RuntimeId *string `pulumi:"runtimeId"`
	// Allows to invert the behaviour of when a check is considered to fail.
	ShouldFail *bool `pulumi:"shouldFail"`
	// A list of tags for organizing and filtering checks.
	Tags []string `pulumi:"tags"`
	// When true, the account level alert settings will be used, not the alert setting defined on this check.
	UseGlobalAlertSettings *bool `pulumi:"useGlobalAlertSettings"`
}

type TcpCheckState struct {
	// Determines if the check is running or not. Possible values `true`, and `false`.
	Activated pulumi.BoolPtrInput
	// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
	AlertChannelSubscriptions TcpCheckAlertChannelSubscriptionArrayInput
	AlertSettings             TcpCheckAlertSettingsPtrInput
	// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 5000. (Default `4000`).
	DegradedResponseTime pulumi.IntPtrInput
	// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
	Frequency pulumi.IntPtrInput
	// To create a high frequency check, the property `frequency` must be `0` and `frequencyOffset` can be `10`, `20` or `30`.
	FrequencyOffset pulumi.IntPtrInput
	// The id of the check group this check is part of.
	GroupId pulumi.IntPtrInput
	// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
	GroupOrder pulumi.IntPtrInput
	// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
	Locations pulumi.StringArrayInput
	// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 5000. (Default `5000`).
	MaxResponseTime pulumi.IntPtrInput
	// Determines if any notifications will be sent out when a check fails/degrades/recovers.
	Muted pulumi.BoolPtrInput
	// The name of the check.
	Name pulumi.StringPtrInput
	// An array of one or more private locations slugs.
	PrivateLocations pulumi.StringArrayInput
	// The parameters for the TCP connection.
	Request TcpCheckRequestPtrInput
	// A strategy for retrying failed check runs.
	RetryStrategy TcpCheckRetryStrategyPtrInput
	// Determines if the check should run in all selected locations in parallel or round-robin.
	RunParallel pulumi.BoolPtrInput
	// The ID of the runtime to use for this check.
	RuntimeId pulumi.StringPtrInput
	// Allows to invert the behaviour of when a check is considered to fail.
	ShouldFail pulumi.BoolPtrInput
	// A list of tags for organizing and filtering checks.
	Tags pulumi.StringArrayInput
	// When true, the account level alert settings will be used, not the alert setting defined on this check.
	UseGlobalAlertSettings pulumi.BoolPtrInput
}

func (TcpCheckState) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpCheckState)(nil)).Elem()
}

type tcpCheckArgs struct {
	// Determines if the check is running or not. Possible values `true`, and `false`.
	Activated bool `pulumi:"activated"`
	// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
	AlertChannelSubscriptions []TcpCheckAlertChannelSubscription `pulumi:"alertChannelSubscriptions"`
	AlertSettings             *TcpCheckAlertSettings             `pulumi:"alertSettings"`
	// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 5000. (Default `4000`).
	DegradedResponseTime *int `pulumi:"degradedResponseTime"`
	// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
	Frequency int `pulumi:"frequency"`
	// To create a high frequency check, the property `frequency` must be `0` and `frequencyOffset` can be `10`, `20` or `30`.
	FrequencyOffset *int `pulumi:"frequencyOffset"`
	// The id of the check group this check is part of.
	GroupId *int `pulumi:"groupId"`
	// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
	GroupOrder *int `pulumi:"groupOrder"`
	// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
	Locations []string `pulumi:"locations"`
	// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 5000. (Default `5000`).
	MaxResponseTime *int `pulumi:"maxResponseTime"`
	// Determines if any notifications will be sent out when a check fails/degrades/recovers.
	Muted *bool `pulumi:"muted"`
	// The name of the check.
	Name *string `pulumi:"name"`
	// An array of one or more private locations slugs.
	PrivateLocations []string `pulumi:"privateLocations"`
	// The parameters for the TCP connection.
	Request TcpCheckRequest `pulumi:"request"`
	// A strategy for retrying failed check runs.
	RetryStrategy *TcpCheckRetryStrategy `pulumi:"retryStrategy"`
	// Determines if the check should run in all selected locations in parallel or round-robin.
	RunParallel *bool `pulumi:"runParallel"`
	// The ID of the runtime to use for this check.
	RuntimeId *string `pulumi:"runtimeId"`
	// Allows to invert the behaviour of when a check is considered to fail.
	ShouldFail *bool `pulumi:"shouldFail"`
	// A list of tags for organizing and filtering checks.
	Tags []string `pulumi:"tags"`
	// When true, the account level alert settings will be used, not the alert setting defined on this check.
	UseGlobalAlertSettings *bool `pulumi:"useGlobalAlertSettings"`
}

// The set of arguments for constructing a TcpCheck resource.
type TcpCheckArgs struct {
	// Determines if the check is running or not. Possible values `true`, and `false`.
	Activated pulumi.BoolInput
	// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
	AlertChannelSubscriptions TcpCheckAlertChannelSubscriptionArrayInput
	AlertSettings             TcpCheckAlertSettingsPtrInput
	// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 5000. (Default `4000`).
	DegradedResponseTime pulumi.IntPtrInput
	// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
	Frequency pulumi.IntInput
	// To create a high frequency check, the property `frequency` must be `0` and `frequencyOffset` can be `10`, `20` or `30`.
	FrequencyOffset pulumi.IntPtrInput
	// The id of the check group this check is part of.
	GroupId pulumi.IntPtrInput
	// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
	GroupOrder pulumi.IntPtrInput
	// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
	Locations pulumi.StringArrayInput
	// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 5000. (Default `5000`).
	MaxResponseTime pulumi.IntPtrInput
	// Determines if any notifications will be sent out when a check fails/degrades/recovers.
	Muted pulumi.BoolPtrInput
	// The name of the check.
	Name pulumi.StringPtrInput
	// An array of one or more private locations slugs.
	PrivateLocations pulumi.StringArrayInput
	// The parameters for the TCP connection.
	Request TcpCheckRequestInput
	// A strategy for retrying failed check runs.
	RetryStrategy TcpCheckRetryStrategyPtrInput
	// Determines if the check should run in all selected locations in parallel or round-robin.
	RunParallel pulumi.BoolPtrInput
	// The ID of the runtime to use for this check.
	RuntimeId pulumi.StringPtrInput
	// Allows to invert the behaviour of when a check is considered to fail.
	ShouldFail pulumi.BoolPtrInput
	// A list of tags for organizing and filtering checks.
	Tags pulumi.StringArrayInput
	// When true, the account level alert settings will be used, not the alert setting defined on this check.
	UseGlobalAlertSettings pulumi.BoolPtrInput
}

func (TcpCheckArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tcpCheckArgs)(nil)).Elem()
}

type TcpCheckInput interface {
	pulumi.Input

	ToTcpCheckOutput() TcpCheckOutput
	ToTcpCheckOutputWithContext(ctx context.Context) TcpCheckOutput
}

func (*TcpCheck) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpCheck)(nil)).Elem()
}

func (i *TcpCheck) ToTcpCheckOutput() TcpCheckOutput {
	return i.ToTcpCheckOutputWithContext(context.Background())
}

func (i *TcpCheck) ToTcpCheckOutputWithContext(ctx context.Context) TcpCheckOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpCheckOutput)
}

// TcpCheckArrayInput is an input type that accepts TcpCheckArray and TcpCheckArrayOutput values.
// You can construct a concrete instance of `TcpCheckArrayInput` via:
//
//	TcpCheckArray{ TcpCheckArgs{...} }
type TcpCheckArrayInput interface {
	pulumi.Input

	ToTcpCheckArrayOutput() TcpCheckArrayOutput
	ToTcpCheckArrayOutputWithContext(context.Context) TcpCheckArrayOutput
}

type TcpCheckArray []TcpCheckInput

func (TcpCheckArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpCheck)(nil)).Elem()
}

func (i TcpCheckArray) ToTcpCheckArrayOutput() TcpCheckArrayOutput {
	return i.ToTcpCheckArrayOutputWithContext(context.Background())
}

func (i TcpCheckArray) ToTcpCheckArrayOutputWithContext(ctx context.Context) TcpCheckArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpCheckArrayOutput)
}

// TcpCheckMapInput is an input type that accepts TcpCheckMap and TcpCheckMapOutput values.
// You can construct a concrete instance of `TcpCheckMapInput` via:
//
//	TcpCheckMap{ "key": TcpCheckArgs{...} }
type TcpCheckMapInput interface {
	pulumi.Input

	ToTcpCheckMapOutput() TcpCheckMapOutput
	ToTcpCheckMapOutputWithContext(context.Context) TcpCheckMapOutput
}

type TcpCheckMap map[string]TcpCheckInput

func (TcpCheckMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpCheck)(nil)).Elem()
}

func (i TcpCheckMap) ToTcpCheckMapOutput() TcpCheckMapOutput {
	return i.ToTcpCheckMapOutputWithContext(context.Background())
}

func (i TcpCheckMap) ToTcpCheckMapOutputWithContext(ctx context.Context) TcpCheckMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TcpCheckMapOutput)
}

type TcpCheckOutput struct{ *pulumi.OutputState }

func (TcpCheckOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TcpCheck)(nil)).Elem()
}

func (o TcpCheckOutput) ToTcpCheckOutput() TcpCheckOutput {
	return o
}

func (o TcpCheckOutput) ToTcpCheckOutputWithContext(ctx context.Context) TcpCheckOutput {
	return o
}

// Determines if the check is running or not. Possible values `true`, and `false`.
func (o TcpCheckOutput) Activated() pulumi.BoolOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.BoolOutput { return v.Activated }).(pulumi.BoolOutput)
}

// An array of channel IDs and whether they're activated or not. If you don't set at least one alert subscription for your check, we won't be able to alert you in case something goes wrong with it.
func (o TcpCheckOutput) AlertChannelSubscriptions() TcpCheckAlertChannelSubscriptionArrayOutput {
	return o.ApplyT(func(v *TcpCheck) TcpCheckAlertChannelSubscriptionArrayOutput { return v.AlertChannelSubscriptions }).(TcpCheckAlertChannelSubscriptionArrayOutput)
}

func (o TcpCheckOutput) AlertSettings() TcpCheckAlertSettingsOutput {
	return o.ApplyT(func(v *TcpCheck) TcpCheckAlertSettingsOutput { return v.AlertSettings }).(TcpCheckAlertSettingsOutput)
}

// The response time in milliseconds starting from which a check should be considered degraded. Possible values are between 0 and 5000. (Default `4000`).
func (o TcpCheckOutput) DegradedResponseTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.IntPtrOutput { return v.DegradedResponseTime }).(pulumi.IntPtrOutput)
}

// The frequency in minutes to run the check. Possible values are `0`, `1`, `2`, `5`, `10`, `15`, `30`, `60`, `120`, `180`, `360`, `720`, and `1440`.
func (o TcpCheckOutput) Frequency() pulumi.IntOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.IntOutput { return v.Frequency }).(pulumi.IntOutput)
}

// To create a high frequency check, the property `frequency` must be `0` and `frequencyOffset` can be `10`, `20` or `30`.
func (o TcpCheckOutput) FrequencyOffset() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.IntPtrOutput { return v.FrequencyOffset }).(pulumi.IntPtrOutput)
}

// The id of the check group this check is part of.
func (o TcpCheckOutput) GroupId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.IntPtrOutput { return v.GroupId }).(pulumi.IntPtrOutput)
}

// The position of this check in a check group. It determines in what order checks are run when a group is triggered from the API or from CI/CD.
func (o TcpCheckOutput) GroupOrder() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.IntPtrOutput { return v.GroupOrder }).(pulumi.IntPtrOutput)
}

// An array of one or more data center locations where to run the this check. (Default ["us-east-1"])
func (o TcpCheckOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.StringArrayOutput { return v.Locations }).(pulumi.StringArrayOutput)
}

// The response time in milliseconds starting from which a check should be considered failing. Possible values are between 0 and 5000. (Default `5000`).
func (o TcpCheckOutput) MaxResponseTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.IntPtrOutput { return v.MaxResponseTime }).(pulumi.IntPtrOutput)
}

// Determines if any notifications will be sent out when a check fails/degrades/recovers.
func (o TcpCheckOutput) Muted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.BoolPtrOutput { return v.Muted }).(pulumi.BoolPtrOutput)
}

// The name of the check.
func (o TcpCheckOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// An array of one or more private locations slugs.
func (o TcpCheckOutput) PrivateLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.StringArrayOutput { return v.PrivateLocations }).(pulumi.StringArrayOutput)
}

// The parameters for the TCP connection.
func (o TcpCheckOutput) Request() TcpCheckRequestOutput {
	return o.ApplyT(func(v *TcpCheck) TcpCheckRequestOutput { return v.Request }).(TcpCheckRequestOutput)
}

// A strategy for retrying failed check runs.
func (o TcpCheckOutput) RetryStrategy() TcpCheckRetryStrategyOutput {
	return o.ApplyT(func(v *TcpCheck) TcpCheckRetryStrategyOutput { return v.RetryStrategy }).(TcpCheckRetryStrategyOutput)
}

// Determines if the check should run in all selected locations in parallel or round-robin.
func (o TcpCheckOutput) RunParallel() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.BoolPtrOutput { return v.RunParallel }).(pulumi.BoolPtrOutput)
}

// The ID of the runtime to use for this check.
func (o TcpCheckOutput) RuntimeId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.StringPtrOutput { return v.RuntimeId }).(pulumi.StringPtrOutput)
}

// Allows to invert the behaviour of when a check is considered to fail.
func (o TcpCheckOutput) ShouldFail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.BoolPtrOutput { return v.ShouldFail }).(pulumi.BoolPtrOutput)
}

// A list of tags for organizing and filtering checks.
func (o TcpCheckOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// When true, the account level alert settings will be used, not the alert setting defined on this check.
func (o TcpCheckOutput) UseGlobalAlertSettings() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TcpCheck) pulumi.BoolPtrOutput { return v.UseGlobalAlertSettings }).(pulumi.BoolPtrOutput)
}

type TcpCheckArrayOutput struct{ *pulumi.OutputState }

func (TcpCheckArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TcpCheck)(nil)).Elem()
}

func (o TcpCheckArrayOutput) ToTcpCheckArrayOutput() TcpCheckArrayOutput {
	return o
}

func (o TcpCheckArrayOutput) ToTcpCheckArrayOutputWithContext(ctx context.Context) TcpCheckArrayOutput {
	return o
}

func (o TcpCheckArrayOutput) Index(i pulumi.IntInput) TcpCheckOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TcpCheck {
		return vs[0].([]*TcpCheck)[vs[1].(int)]
	}).(TcpCheckOutput)
}

type TcpCheckMapOutput struct{ *pulumi.OutputState }

func (TcpCheckMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TcpCheck)(nil)).Elem()
}

func (o TcpCheckMapOutput) ToTcpCheckMapOutput() TcpCheckMapOutput {
	return o
}

func (o TcpCheckMapOutput) ToTcpCheckMapOutputWithContext(ctx context.Context) TcpCheckMapOutput {
	return o
}

func (o TcpCheckMapOutput) MapIndex(k pulumi.StringInput) TcpCheckOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TcpCheck {
		return vs[0].(map[string]*TcpCheck)[vs[1].(string)]
	}).(TcpCheckOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TcpCheckInput)(nil)).Elem(), &TcpCheck{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpCheckArrayInput)(nil)).Elem(), TcpCheckArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TcpCheckMapInput)(nil)).Elem(), TcpCheckMap{})
	pulumi.RegisterOutputType(TcpCheckOutput{})
	pulumi.RegisterOutputType(TcpCheckArrayOutput{})
	pulumi.RegisterOutputType(TcpCheckMapOutput{})
}
