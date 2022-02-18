// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicDashboard struct {
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

// NewPublicDashboard registers a new resource with the given unique name, arguments, and options.
func NewPublicDashboard(ctx *pulumi.Context,
	name string, args *PublicDashboardArgs, opts ...pulumi.ResourceOption) (*PublicDashboard, error) {
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
	var resource PublicDashboard
	err := ctx.RegisterResource("checkly:index/publicDashboard:PublicDashboard", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPublicDashboard gets an existing PublicDashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PublicDashboardState, opts ...pulumi.ResourceOption) (*PublicDashboard, error) {
	var resource PublicDashboard
	err := ctx.ReadResource("checkly:index/publicDashboard:PublicDashboard", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PublicDashboard resources.
type publicDashboardState struct {
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

type PublicDashboardState struct {
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

func (PublicDashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDashboardState)(nil)).Elem()
}

type publicDashboardArgs struct {
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

// The set of arguments for constructing a PublicDashboard resource.
type PublicDashboardArgs struct {
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

func (PublicDashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*publicDashboardArgs)(nil)).Elem()
}

type PublicDashboardInput interface {
	pulumi.Input

	ToPublicDashboardOutput() PublicDashboardOutput
	ToPublicDashboardOutputWithContext(ctx context.Context) PublicDashboardOutput
}

func (*PublicDashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDashboard)(nil)).Elem()
}

func (i *PublicDashboard) ToPublicDashboardOutput() PublicDashboardOutput {
	return i.ToPublicDashboardOutputWithContext(context.Background())
}

func (i *PublicDashboard) ToPublicDashboardOutputWithContext(ctx context.Context) PublicDashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDashboardOutput)
}

// PublicDashboardArrayInput is an input type that accepts PublicDashboardArray and PublicDashboardArrayOutput values.
// You can construct a concrete instance of `PublicDashboardArrayInput` via:
//
//          PublicDashboardArray{ PublicDashboardArgs{...} }
type PublicDashboardArrayInput interface {
	pulumi.Input

	ToPublicDashboardArrayOutput() PublicDashboardArrayOutput
	ToPublicDashboardArrayOutputWithContext(context.Context) PublicDashboardArrayOutput
}

type PublicDashboardArray []PublicDashboardInput

func (PublicDashboardArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicDashboard)(nil)).Elem()
}

func (i PublicDashboardArray) ToPublicDashboardArrayOutput() PublicDashboardArrayOutput {
	return i.ToPublicDashboardArrayOutputWithContext(context.Background())
}

func (i PublicDashboardArray) ToPublicDashboardArrayOutputWithContext(ctx context.Context) PublicDashboardArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDashboardArrayOutput)
}

// PublicDashboardMapInput is an input type that accepts PublicDashboardMap and PublicDashboardMapOutput values.
// You can construct a concrete instance of `PublicDashboardMapInput` via:
//
//          PublicDashboardMap{ "key": PublicDashboardArgs{...} }
type PublicDashboardMapInput interface {
	pulumi.Input

	ToPublicDashboardMapOutput() PublicDashboardMapOutput
	ToPublicDashboardMapOutputWithContext(context.Context) PublicDashboardMapOutput
}

type PublicDashboardMap map[string]PublicDashboardInput

func (PublicDashboardMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicDashboard)(nil)).Elem()
}

func (i PublicDashboardMap) ToPublicDashboardMapOutput() PublicDashboardMapOutput {
	return i.ToPublicDashboardMapOutputWithContext(context.Background())
}

func (i PublicDashboardMap) ToPublicDashboardMapOutputWithContext(ctx context.Context) PublicDashboardMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PublicDashboardMapOutput)
}

type PublicDashboardOutput struct{ *pulumi.OutputState }

func (PublicDashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicDashboard)(nil)).Elem()
}

func (o PublicDashboardOutput) ToPublicDashboardOutput() PublicDashboardOutput {
	return o
}

func (o PublicDashboardOutput) ToPublicDashboardOutputWithContext(ctx context.Context) PublicDashboardOutput {
	return o
}

type PublicDashboardArrayOutput struct{ *pulumi.OutputState }

func (PublicDashboardArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PublicDashboard)(nil)).Elem()
}

func (o PublicDashboardArrayOutput) ToPublicDashboardArrayOutput() PublicDashboardArrayOutput {
	return o
}

func (o PublicDashboardArrayOutput) ToPublicDashboardArrayOutputWithContext(ctx context.Context) PublicDashboardArrayOutput {
	return o
}

func (o PublicDashboardArrayOutput) Index(i pulumi.IntInput) PublicDashboardOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PublicDashboard {
		return vs[0].([]*PublicDashboard)[vs[1].(int)]
	}).(PublicDashboardOutput)
}

type PublicDashboardMapOutput struct{ *pulumi.OutputState }

func (PublicDashboardMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PublicDashboard)(nil)).Elem()
}

func (o PublicDashboardMapOutput) ToPublicDashboardMapOutput() PublicDashboardMapOutput {
	return o
}

func (o PublicDashboardMapOutput) ToPublicDashboardMapOutputWithContext(ctx context.Context) PublicDashboardMapOutput {
	return o
}

func (o PublicDashboardMapOutput) MapIndex(k pulumi.StringInput) PublicDashboardOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PublicDashboard {
		return vs[0].(map[string]*PublicDashboard)[vs[1].(string)]
	}).(PublicDashboardOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PublicDashboardInput)(nil)).Elem(), &PublicDashboard{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicDashboardArrayInput)(nil)).Elem(), PublicDashboardArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PublicDashboardMapInput)(nil)).Elem(), PublicDashboardMap{})
	pulumi.RegisterOutputType(PublicDashboardOutput{})
	pulumi.RegisterOutputType(PublicDashboardArrayOutput{})
	pulumi.RegisterOutputType(PublicDashboardMapOutput{})
}
