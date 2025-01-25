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

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/checkly/pulumi-checkly/sdk/go/checkly"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := checkly.NewMaintenanceWindow(ctx, "maintenance-1", &checkly.MaintenanceWindowArgs{
//				Name:           pulumi.String("Maintenance Windows"),
//				StartsAt:       pulumi.String("2014-08-24T00:00:00.000Z"),
//				EndsAt:         pulumi.String("2014-08-25T00:00:00.000Z"),
//				RepeatUnit:     pulumi.String("MONTH"),
//				RepeatEndsAt:   pulumi.String("2014-08-24T00:00:00.000Z"),
//				RepeatInterval: pulumi.Int(1),
//				Tags: pulumi.StringArray{
//					pulumi.String("production"),
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
type MaintenanceWindow struct {
	pulumi.CustomResourceState

	// The end date of the maintenance window.
	EndsAt pulumi.StringOutput `pulumi:"endsAt"`
	// The maintenance window name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The date on which the maintenance window should stop repeating.
	RepeatEndsAt pulumi.StringPtrOutput `pulumi:"repeatEndsAt"`
	// The repeat interval of the maintenance window from the first occurrence.
	RepeatInterval pulumi.IntPtrOutput `pulumi:"repeatInterval"`
	// The repeat cadence for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit pulumi.StringPtrOutput `pulumi:"repeatUnit"`
	// The start date of the maintenance window.
	StartsAt pulumi.StringOutput `pulumi:"startsAt"`
	// The names of the checks and groups maintenance window should apply to.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewMaintenanceWindow registers a new resource with the given unique name, arguments, and options.
func NewMaintenanceWindow(ctx *pulumi.Context,
	name string, args *MaintenanceWindowArgs, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndsAt == nil {
		return nil, errors.New("invalid value for required argument 'EndsAt'")
	}
	if args.StartsAt == nil {
		return nil, errors.New("invalid value for required argument 'StartsAt'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MaintenanceWindow
	err := ctx.RegisterResource("checkly:index/maintenanceWindow:MaintenanceWindow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMaintenanceWindow gets an existing MaintenanceWindow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMaintenanceWindow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MaintenanceWindowState, opts ...pulumi.ResourceOption) (*MaintenanceWindow, error) {
	var resource MaintenanceWindow
	err := ctx.ReadResource("checkly:index/maintenanceWindow:MaintenanceWindow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MaintenanceWindow resources.
type maintenanceWindowState struct {
	// The end date of the maintenance window.
	EndsAt *string `pulumi:"endsAt"`
	// The maintenance window name.
	Name *string `pulumi:"name"`
	// The date on which the maintenance window should stop repeating.
	RepeatEndsAt *string `pulumi:"repeatEndsAt"`
	// The repeat interval of the maintenance window from the first occurrence.
	RepeatInterval *int `pulumi:"repeatInterval"`
	// The repeat cadence for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit *string `pulumi:"repeatUnit"`
	// The start date of the maintenance window.
	StartsAt *string `pulumi:"startsAt"`
	// The names of the checks and groups maintenance window should apply to.
	Tags []string `pulumi:"tags"`
}

type MaintenanceWindowState struct {
	// The end date of the maintenance window.
	EndsAt pulumi.StringPtrInput
	// The maintenance window name.
	Name pulumi.StringPtrInput
	// The date on which the maintenance window should stop repeating.
	RepeatEndsAt pulumi.StringPtrInput
	// The repeat interval of the maintenance window from the first occurrence.
	RepeatInterval pulumi.IntPtrInput
	// The repeat cadence for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit pulumi.StringPtrInput
	// The start date of the maintenance window.
	StartsAt pulumi.StringPtrInput
	// The names of the checks and groups maintenance window should apply to.
	Tags pulumi.StringArrayInput
}

func (MaintenanceWindowState) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowState)(nil)).Elem()
}

type maintenanceWindowArgs struct {
	// The end date of the maintenance window.
	EndsAt string `pulumi:"endsAt"`
	// The maintenance window name.
	Name *string `pulumi:"name"`
	// The date on which the maintenance window should stop repeating.
	RepeatEndsAt *string `pulumi:"repeatEndsAt"`
	// The repeat interval of the maintenance window from the first occurrence.
	RepeatInterval *int `pulumi:"repeatInterval"`
	// The repeat cadence for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit *string `pulumi:"repeatUnit"`
	// The start date of the maintenance window.
	StartsAt string `pulumi:"startsAt"`
	// The names of the checks and groups maintenance window should apply to.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a MaintenanceWindow resource.
type MaintenanceWindowArgs struct {
	// The end date of the maintenance window.
	EndsAt pulumi.StringInput
	// The maintenance window name.
	Name pulumi.StringPtrInput
	// The date on which the maintenance window should stop repeating.
	RepeatEndsAt pulumi.StringPtrInput
	// The repeat interval of the maintenance window from the first occurrence.
	RepeatInterval pulumi.IntPtrInput
	// The repeat cadence for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
	RepeatUnit pulumi.StringPtrInput
	// The start date of the maintenance window.
	StartsAt pulumi.StringInput
	// The names of the checks and groups maintenance window should apply to.
	Tags pulumi.StringArrayInput
}

func (MaintenanceWindowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*maintenanceWindowArgs)(nil)).Elem()
}

type MaintenanceWindowInput interface {
	pulumi.Input

	ToMaintenanceWindowOutput() MaintenanceWindowOutput
	ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput
}

func (*MaintenanceWindow) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return i.ToMaintenanceWindowOutputWithContext(context.Background())
}

func (i *MaintenanceWindow) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowOutput)
}

// MaintenanceWindowArrayInput is an input type that accepts MaintenanceWindowArray and MaintenanceWindowArrayOutput values.
// You can construct a concrete instance of `MaintenanceWindowArrayInput` via:
//
//	MaintenanceWindowArray{ MaintenanceWindowArgs{...} }
type MaintenanceWindowArrayInput interface {
	pulumi.Input

	ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput
	ToMaintenanceWindowArrayOutputWithContext(context.Context) MaintenanceWindowArrayOutput
}

type MaintenanceWindowArray []MaintenanceWindowInput

func (MaintenanceWindowArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return i.ToMaintenanceWindowArrayOutputWithContext(context.Background())
}

func (i MaintenanceWindowArray) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowArrayOutput)
}

// MaintenanceWindowMapInput is an input type that accepts MaintenanceWindowMap and MaintenanceWindowMapOutput values.
// You can construct a concrete instance of `MaintenanceWindowMapInput` via:
//
//	MaintenanceWindowMap{ "key": MaintenanceWindowArgs{...} }
type MaintenanceWindowMapInput interface {
	pulumi.Input

	ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput
	ToMaintenanceWindowMapOutputWithContext(context.Context) MaintenanceWindowMapOutput
}

type MaintenanceWindowMap map[string]MaintenanceWindowInput

func (MaintenanceWindowMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return i.ToMaintenanceWindowMapOutputWithContext(context.Background())
}

func (i MaintenanceWindowMap) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MaintenanceWindowMapOutput)
}

type MaintenanceWindowOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutput() MaintenanceWindowOutput {
	return o
}

func (o MaintenanceWindowOutput) ToMaintenanceWindowOutputWithContext(ctx context.Context) MaintenanceWindowOutput {
	return o
}

// The end date of the maintenance window.
func (o MaintenanceWindowOutput) EndsAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.EndsAt }).(pulumi.StringOutput)
}

// The maintenance window name.
func (o MaintenanceWindowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The date on which the maintenance window should stop repeating.
func (o MaintenanceWindowOutput) RepeatEndsAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.RepeatEndsAt }).(pulumi.StringPtrOutput)
}

// The repeat interval of the maintenance window from the first occurrence.
func (o MaintenanceWindowOutput) RepeatInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.IntPtrOutput { return v.RepeatInterval }).(pulumi.IntPtrOutput)
}

// The repeat cadence for the maintenance window. Possible values `DAY`, `WEEK` and `MONTH`.
func (o MaintenanceWindowOutput) RepeatUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringPtrOutput { return v.RepeatUnit }).(pulumi.StringPtrOutput)
}

// The start date of the maintenance window.
func (o MaintenanceWindowOutput) StartsAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringOutput { return v.StartsAt }).(pulumi.StringOutput)
}

// The names of the checks and groups maintenance window should apply to.
func (o MaintenanceWindowOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MaintenanceWindow) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type MaintenanceWindowArrayOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutput() MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) ToMaintenanceWindowArrayOutputWithContext(ctx context.Context) MaintenanceWindowArrayOutput {
	return o
}

func (o MaintenanceWindowArrayOutput) Index(i pulumi.IntInput) MaintenanceWindowOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].([]*MaintenanceWindow)[vs[1].(int)]
	}).(MaintenanceWindowOutput)
}

type MaintenanceWindowMapOutput struct{ *pulumi.OutputState }

func (MaintenanceWindowMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MaintenanceWindow)(nil)).Elem()
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutput() MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) ToMaintenanceWindowMapOutputWithContext(ctx context.Context) MaintenanceWindowMapOutput {
	return o
}

func (o MaintenanceWindowMapOutput) MapIndex(k pulumi.StringInput) MaintenanceWindowOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MaintenanceWindow {
		return vs[0].(map[string]*MaintenanceWindow)[vs[1].(string)]
	}).(MaintenanceWindowOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowInput)(nil)).Elem(), &MaintenanceWindow{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowArrayInput)(nil)).Elem(), MaintenanceWindowArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MaintenanceWindowMapInput)(nil)).Elem(), MaintenanceWindowMap{})
	pulumi.RegisterOutputType(MaintenanceWindowOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowArrayOutput{})
	pulumi.RegisterOutputType(MaintenanceWindowMapOutput{})
}
