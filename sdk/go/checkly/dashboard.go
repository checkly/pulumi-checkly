// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Dashboard struct {
	pulumi.CustomResourceState

	CustomDomain   pulumi.StringOutput      `pulumi:"customDomain"`
	CustomUrl      pulumi.StringOutput      `pulumi:"customUrl"`
	Header         pulumi.StringOutput      `pulumi:"header"`
	HideTags       pulumi.BoolOutput        `pulumi:"hideTags"`
	Logo           pulumi.StringOutput      `pulumi:"logo"`
	Paginate       pulumi.BoolOutput        `pulumi:"paginate"`
	PaginationRate pulumi.IntOutput         `pulumi:"paginationRate"`
	RefreshRate    pulumi.IntOutput         `pulumi:"refreshRate"`
	Tags           pulumi.StringArrayOutput `pulumi:"tags"`
	Width          pulumi.StringPtrOutput   `pulumi:"width"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomDomain == nil {
		return nil, errors.New("invalid value for required argument 'CustomDomain'")
	}
	if args.CustomUrl == nil {
		return nil, errors.New("invalid value for required argument 'CustomUrl'")
	}
	if args.Header == nil {
		return nil, errors.New("invalid value for required argument 'Header'")
	}
	if args.HideTags == nil {
		return nil, errors.New("invalid value for required argument 'HideTags'")
	}
	if args.Logo == nil {
		return nil, errors.New("invalid value for required argument 'Logo'")
	}
	if args.Paginate == nil {
		return nil, errors.New("invalid value for required argument 'Paginate'")
	}
	if args.PaginationRate == nil {
		return nil, errors.New("invalid value for required argument 'PaginationRate'")
	}
	if args.RefreshRate == nil {
		return nil, errors.New("invalid value for required argument 'RefreshRate'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Dashboard
	err := ctx.RegisterResource("checkly:index/dashboard:Dashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboard gets an existing Dashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardState, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	var resource Dashboard
	err := ctx.ReadResource("checkly:index/dashboard:Dashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Dashboard resources.
type dashboardState struct {
	CustomDomain   *string  `pulumi:"customDomain"`
	CustomUrl      *string  `pulumi:"customUrl"`
	Header         *string  `pulumi:"header"`
	HideTags       *bool    `pulumi:"hideTags"`
	Logo           *string  `pulumi:"logo"`
	Paginate       *bool    `pulumi:"paginate"`
	PaginationRate *int     `pulumi:"paginationRate"`
	RefreshRate    *int     `pulumi:"refreshRate"`
	Tags           []string `pulumi:"tags"`
	Width          *string  `pulumi:"width"`
}

type DashboardState struct {
	CustomDomain   pulumi.StringPtrInput
	CustomUrl      pulumi.StringPtrInput
	Header         pulumi.StringPtrInput
	HideTags       pulumi.BoolPtrInput
	Logo           pulumi.StringPtrInput
	Paginate       pulumi.BoolPtrInput
	PaginationRate pulumi.IntPtrInput
	RefreshRate    pulumi.IntPtrInput
	Tags           pulumi.StringArrayInput
	Width          pulumi.StringPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	CustomDomain   string   `pulumi:"customDomain"`
	CustomUrl      string   `pulumi:"customUrl"`
	Header         string   `pulumi:"header"`
	HideTags       bool     `pulumi:"hideTags"`
	Logo           string   `pulumi:"logo"`
	Paginate       bool     `pulumi:"paginate"`
	PaginationRate int      `pulumi:"paginationRate"`
	RefreshRate    int      `pulumi:"refreshRate"`
	Tags           []string `pulumi:"tags"`
	Width          *string  `pulumi:"width"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	CustomDomain   pulumi.StringInput
	CustomUrl      pulumi.StringInput
	Header         pulumi.StringInput
	HideTags       pulumi.BoolInput
	Logo           pulumi.StringInput
	Paginate       pulumi.BoolInput
	PaginationRate pulumi.IntInput
	RefreshRate    pulumi.IntInput
	Tags           pulumi.StringArrayInput
	Width          pulumi.StringPtrInput
}

func (DashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardArgs)(nil)).Elem()
}

type DashboardInput interface {
	pulumi.Input

	ToDashboardOutput() DashboardOutput
	ToDashboardOutputWithContext(ctx context.Context) DashboardOutput
}

func (*Dashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (i *Dashboard) ToDashboardOutput() DashboardOutput {
	return i.ToDashboardOutputWithContext(context.Background())
}

func (i *Dashboard) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardOutput)
}

// DashboardArrayInput is an input type that accepts DashboardArray and DashboardArrayOutput values.
// You can construct a concrete instance of `DashboardArrayInput` via:
//
//          DashboardArray{ DashboardArgs{...} }
type DashboardArrayInput interface {
	pulumi.Input

	ToDashboardArrayOutput() DashboardArrayOutput
	ToDashboardArrayOutputWithContext(context.Context) DashboardArrayOutput
}

type DashboardArray []DashboardInput

func (DashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (i DashboardArray) ToDashboardArrayOutput() DashboardArrayOutput {
	return i.ToDashboardArrayOutputWithContext(context.Background())
}

func (i DashboardArray) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardArrayOutput)
}

// DashboardMapInput is an input type that accepts DashboardMap and DashboardMapOutput values.
// You can construct a concrete instance of `DashboardMapInput` via:
//
//          DashboardMap{ "key": DashboardArgs{...} }
type DashboardMapInput interface {
	pulumi.Input

	ToDashboardMapOutput() DashboardMapOutput
	ToDashboardMapOutputWithContext(context.Context) DashboardMapOutput
}

type DashboardMap map[string]DashboardInput

func (DashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (i DashboardMap) ToDashboardMapOutput() DashboardMapOutput {
	return i.ToDashboardMapOutputWithContext(context.Background())
}

func (i DashboardMap) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardMapOutput)
}

type DashboardOutput struct{ *pulumi.OutputState }

func (DashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dashboard)(nil)).Elem()
}

func (o DashboardOutput) ToDashboardOutput() DashboardOutput {
	return o
}

func (o DashboardOutput) ToDashboardOutputWithContext(ctx context.Context) DashboardOutput {
	return o
}

type DashboardArrayOutput struct{ *pulumi.OutputState }

func (DashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Dashboard)(nil)).Elem()
}

func (o DashboardArrayOutput) ToDashboardArrayOutput() DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) ToDashboardArrayOutputWithContext(ctx context.Context) DashboardArrayOutput {
	return o
}

func (o DashboardArrayOutput) Index(i pulumi.IntInput) DashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].([]*Dashboard)[vs[1].(int)]
	}).(DashboardOutput)
}

type DashboardMapOutput struct{ *pulumi.OutputState }

func (DashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Dashboard)(nil)).Elem()
}

func (o DashboardMapOutput) ToDashboardMapOutput() DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) ToDashboardMapOutputWithContext(ctx context.Context) DashboardMapOutput {
	return o
}

func (o DashboardMapOutput) MapIndex(k pulumi.StringInput) DashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Dashboard {
		return vs[0].(map[string]*Dashboard)[vs[1].(string)]
	}).(DashboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardInput)(nil)).Elem(), &Dashboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardArrayInput)(nil)).Elem(), DashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardMapInput)(nil)).Elem(), DashboardMap{})
	pulumi.RegisterOutputType(DashboardOutput{})
	pulumi.RegisterOutputType(DashboardArrayOutput{})
	pulumi.RegisterOutputType(DashboardMapOutput{})
}