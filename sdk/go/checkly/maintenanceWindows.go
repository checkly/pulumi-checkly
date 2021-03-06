// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # MaintenanceWindows
//
// `MaintenanceWindows` allows users to manage Checkly maintenance windows. Add a `MaintenanceWindows` resource to your resource file.
//
// ## Example Usage
//
// Minimal maintenance windows example
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := checkly.NewMaintenanceWindows(ctx, "maintenance-1", &checkly.MaintenanceWindowsArgs{
// 			EndsAt:     pulumi.String("2014-08-25T00:00:00.000Z"),
// 			RepeatUnit: pulumi.String("MONTH"),
// 			StartsAt:   pulumi.String("2014-08-24T00:00:00.000Z"),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("auto"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Full maintenance windows example (includes optional fields)
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := checkly.NewMaintenanceWindows(ctx, "maintenance-1", &checkly.MaintenanceWindowsArgs{
// 			EndsAt:         pulumi.String("2014-08-25T00:00:00.000Z"),
// 			RepeatEndsAt:   pulumi.String("2014-08-24T00:00:00.000Z"),
// 			RepeatInterval: pulumi.Int(1),
// 			RepeatUnit:     pulumi.String("MONTH"),
// 			StartsAt:       pulumi.String("2014-08-24T00:00:00.000Z"),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("auto"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MaintenanceWindows struct {
	pulumi.CustomResourceState

	// The end date of the maintenance window.
	EndsAt pulumi.StringOutput `pulumi:"endsAt"`
	// The maintenance window name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The end date where the maintenance window should stop repeating.
	RepeatEndsAt pulumi.StringOutput `pulumi:"repeatEndsAt"`
	// The repeat interval of the maintenance window from the first occurance.
	RepeatInterval pulumi.IntOutput `pulumi:"repeatInterval"`
	// The repeat strategy for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit pulumi.StringOutput `pulumi:"repeatUnit"`
	// The start date of the maintenance window.
	StartsAt pulumi.StringOutput `pulumi:"startsAt"`
	// The names of the checks and groups maintenance window should apply to.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewMaintenanceWindows registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindows(ctx *pulumi.Context,
	name string, args *MaintenanceWindowsArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindows, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndsAt == nil {
		return nil, errors.New("invalid value for required argument 'EndsAt'")
	}
	if args.RepeatEndsAt == nil {
		return nil, errors.New("invalid value for required argument 'RepeatEndsAt'")
	}
	if args.RepeatInterval == nil {
		return nil, errors.New("invalid value for required argument 'RepeatInterval'")
	}
	if args.RepeatUnit == nil {
		return nil, errors.New("invalid value for required argument 'RepeatUnit'")
	}
	if args.StartsAt == nil {
		return nil, errors.New("invalid value for required argument 'StartsAt'")
	}
	var resource MaintenanceWindows
	err := ctx.RegisterResource("checkly:index/maintenanceWindows:MaintenanceWindows", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindows gets an existing MaintenanceWindows resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindows(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowsState, opts ...pulumi.ResourceOption) (*MaintenanceWindows, error) {
	var resource MaintenanceWindows
	err := ctx.ReadResource("checkly:index/maintenanceWindows:MaintenanceWindows", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindows resources.
type maintenanceWindowsState struct {
	// The end date of the maintenance window.
	EndsAt *string `pulumi:"endsAt"`
	// The maintenance window name.
	Name *string `pulumi:"name"`
	// The end date where the maintenance window should stop repeating.
	RepeatEndsAt *string `pulumi:"repeatEndsAt"`
	// The repeat interval of the maintenance window from the first occurance.
	RepeatInterval *int `pulumi:"repeatInterval"`
	// The repeat strategy for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit *string `pulumi:"repeatUnit"`
	// The start date of the maintenance window.
	StartsAt *string `pulumi:"startsAt"`
	// The names of the checks and groups maintenance window should apply to.
	Tags []string `pulumi:"tags"`
}

type MaintenanceWindowsState struct {
	// The end date of the maintenance window.
	EndsAt pulumi.StringPtrInput
	// The maintenance window name.
	Name pulumi.StringPtrInput
	// The end date where the maintenance window should stop repeating.
	RepeatEndsAt pulumi.StringPtrInput
	// The repeat interval of the maintenance window from the first occurance.
	RepeatInterval pulumi.IntPtrInput
	// The repeat strategy for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit pulumi.StringPtrInput
	// The start date of the maintenance window.
	StartsAt pulumi.StringPtrInput
	// The names of the checks and groups maintenance window should apply to.
	Tags pulumi.StringArrayInput
}

func (MaintenanceWindowsState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowsState)(nil)).Elem()
}

type maintenanceWindowsArgs struct {
	// The end date of the maintenance window.
	EndsAt string `pulumi:"endsAt"`
	// The maintenance window name.
	Name *string `pulumi:"name"`
	// The end date where the maintenance window should stop repeating.
	RepeatEndsAt string `pulumi:"repeatEndsAt"`
	// The repeat interval of the maintenance window from the first occurance.
	RepeatInterval int `pulumi:"repeatInterval"`
	// The repeat strategy for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit string `pulumi:"repeatUnit"`
	// The start date of the maintenance window.
	StartsAt string `pulumi:"startsAt"`
	// The names of the checks and groups maintenance window should apply to.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a MaintenanceWindows resource.
type MaintenanceWindowsArgs struct {
	// The end date of the maintenance window.
	EndsAt pulumi.StringInput
	// The maintenance window name.
	Name pulumi.StringPtrInput
	// The end date where the maintenance window should stop repeating.
	RepeatEndsAt pulumi.StringInput
	// The repeat interval of the maintenance window from the first occurance.
	RepeatInterval pulumi.IntInput
	// The repeat strategy for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit pulumi.StringInput
	// The start date of the maintenance window.
	StartsAt pulumi.StringInput
	// The names of the checks and groups maintenance window should apply to.
	Tags pulumi.StringArrayInput
}

func (MaintenanceWindowsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowsArgs)(nil)).Elem()
}

type MaintenanceWindowsInput interface {
	pulumi.Input

	ToMaintenanceWindowsOutput() MaintenanceWindowsOutput
	ToMaintenanceWindowsOutputWithContext(ctx context.Context) MaintenanceWindowsOutput
}

func (*MaintenanceWindows) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindows)(nil)).Elem()
}

func (i *MaintenanceWindows) ToMaintenanceWindowsOutput() MaintenanceWindowsOutput {
	return i.ToMaintenanceWindowsOutputWithContext(context.Background())
}

func (i *MaintenanceWindows) ToMaintenanceWindowsOutputWithContext(ctx context.Context) MaintenanceWindowsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowsOutput)
}

// MaintenanceWindowsArrayInput is an input type that accepts MaintenanceWindowsArray and MaintenanceWindowsArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowsArrayInput` via:
//
//          MaintenanceWindowsArray{ MaintenanceWindowsArgs{...} }
type MaintenanceWindowsArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowsArrayOutput() MaintenanceWindowsArrayOutput
	ToMaintenanceWindowsArrayOutputWithContext(context.Context) MaintenanceWindowsArrayOutput
}

type MaintenanceWindowsArray []MaintenanceWindowsInput

func (MaintenanceWindowsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindows)(nil)).Elem()
}

func (i MaintenanceWindowsArray) ToMaintenanceWindowsArrayOutput() MaintenanceWindowsArrayOutput {
	return i.ToMaintenanceWindowsArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowsArray) ToMaintenanceWindowsArrayOutputWithContext(ctx context.Context) MaintenanceWindowsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowsArrayOutput)
}

// MaintenanceWindowsMapInput is an input type that accepts MaintenanceWindowsMap and MaintenanceWindowsMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowsMapInput` via:
//
//          MaintenanceWindowsMap{ "key": MaintenanceWindowsArgs{...} }
type MaintenanceWindowsMapInput interface {
	pulumi.Input

	ToMaintenanceWindowsMapOutput() MaintenanceWindowsMapOutput
	ToMaintenanceWindowsMapOutputWithContext(context.Context) MaintenanceWindowsMapOutput
}

type MaintenanceWindowsMap map[string]MaintenanceWindowsInput

func (MaintenanceWindowsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindows)(nil)).Elem()
}

func (i MaintenanceWindowsMap) ToMaintenanceWindowsMapOutput() MaintenanceWindowsMapOutput {
	return i.ToMaintenanceWindowsMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowsMap) ToMaintenanceWindowsMapOutputWithContext(ctx context.Context) MaintenanceWindowsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowsMapOutput)
}

type MaintenanceWindowsOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindows)(nil)).Elem()
}

func (o MaintenanceWindowsOutput) ToMaintenanceWindowsOutput() MaintenanceWindowsOutput {
	return o
}

func (o MaintenanceWindowsOutput) ToMaintenanceWindowsOutputWithContext(ctx context.Context) MaintenanceWindowsOutput {
	return o
}

type MaintenanceWindowsArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindows)(nil)).Elem()
}

func (o MaintenanceWindowsArrayOutput) ToMaintenanceWindowsArrayOutput() MaintenanceWindowsArrayOutput {
	return o
}

func (o MaintenanceWindowsArrayOutput) ToMaintenanceWindowsArrayOutputWithContext(ctx context.Context) MaintenanceWindowsArrayOutput {
	return o
}

func (o MaintenanceWindowsArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MaintenanceWindows {
		return vs[0].([]*MaintenanceWindows)[vs[1].(int)]
	}).(MaintenanceWindowsOutput)
}

type MaintenanceWindowsMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindows)(nil)).Elem()
}

func (o MaintenanceWindowsMapOutput) ToMaintenanceWindowsMapOutput() MaintenanceWindowsMapOutput {
	return o
}

func (o MaintenanceWindowsMapOutput) ToMaintenanceWindowsMapOutputWithContext(ctx context.Context) MaintenanceWindowsMapOutput {
	return o
}

func (o MaintenanceWindowsMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MaintenanceWindows {
		return vs[0].(map[string]*MaintenanceWindows)[vs[1].(string)]
	}).(MaintenanceWindowsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowsInput)(nil)).Elem(), &MaintenanceWindows{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowsArrayInput)(nil)).Elem(), MaintenanceWindowsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowsMapInput)(nil)).Elem(), MaintenanceWindowsMap{})
	pulumi.RegisterOutputType(MaintenanceWindowsOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowsArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowsMapOutput{})
}
