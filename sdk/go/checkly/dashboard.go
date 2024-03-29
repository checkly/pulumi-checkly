// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/checkly/pulumi-checkly/sdk/go/checkly"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := checkly.NewDashboard(ctx, "dashboard1", &checkly.DashboardArgs{
// 			CustomDomain:   pulumi.String("status.example.com"),
// 			CustomUrl:      pulumi.String("checkly"),
// 			Header:         pulumi.String("Public dashboard"),
// 			HideTags:       pulumi.Bool(false),
// 			Logo:           pulumi.String("https://www.checklyhq.com/logo.png"),
// 			Paginate:       pulumi.Bool(false),
// 			PaginationRate: pulumi.Int(30),
// 			RefreshRate:    pulumi.Int(60),
// 			Tags: pulumi.StringArray{
// 				pulumi.String("production"),
// 			},
// 			Width: pulumi.String("FULL"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Dashboard struct {
	pulumi.CustomResourceState

	// Determines how many checks to show per page.
	ChecksPerPage pulumi.IntPtrOutput `pulumi:"checksPerPage"`
	// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
	CustomDomain pulumi.StringPtrOutput `pulumi:"customDomain"`
	// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
	CustomUrl pulumi.StringOutput `pulumi:"customUrl"`
	// HTML <meta> description for the dashboard.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A URL pointing to an image file to use as browser favicon.
	Favicon pulumi.StringPtrOutput `pulumi:"favicon"`
	// A piece of text displayed at the top of your dashboard.
	Header pulumi.StringPtrOutput `pulumi:"header"`
	// Show or hide the tags on the dashboard.
	HideTags pulumi.BoolPtrOutput `pulumi:"hideTags"`
	// A link to for the dashboard logo.
	Link pulumi.StringPtrOutput `pulumi:"link"`
	// A URL pointing to an image file to use for the dashboard logo.
	Logo pulumi.StringPtrOutput `pulumi:"logo"`
	// Determines if pagination is on or off.
	Paginate pulumi.BoolPtrOutput `pulumi:"paginate"`
	// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
	PaginationRate pulumi.IntPtrOutput `pulumi:"paginationRate"`
	// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
	RefreshRate pulumi.IntPtrOutput `pulumi:"refreshRate"`
	// A list of one or more tags that filter which checks to display on the dashboard.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// Set when to use AND operator for fetching dashboard tags.
	UseTagsAndOperator pulumi.BoolPtrOutput `pulumi:"useTagsAndOperator"`
	// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
	Width pulumi.StringPtrOutput `pulumi:"width"`
}

// NewDashboard registers a new resource with the given unique name, arguments, and options.
func NewDashboard(ctx *pulumi.Context,
	name string, args *DashboardArgs, opts ...pulumi.ResourceOption) (*Dashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CustomUrl == nil {
		return nil, errors.New("invalid value for required argument 'CustomUrl'")
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
	// Determines how many checks to show per page.
	ChecksPerPage *int `pulumi:"checksPerPage"`
	// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
	CustomDomain *string `pulumi:"customDomain"`
	// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
	CustomUrl *string `pulumi:"customUrl"`
	// HTML <meta> description for the dashboard.
	Description *string `pulumi:"description"`
	// A URL pointing to an image file to use as browser favicon.
	Favicon *string `pulumi:"favicon"`
	// A piece of text displayed at the top of your dashboard.
	Header *string `pulumi:"header"`
	// Show or hide the tags on the dashboard.
	HideTags *bool `pulumi:"hideTags"`
	// A link to for the dashboard logo.
	Link *string `pulumi:"link"`
	// A URL pointing to an image file to use for the dashboard logo.
	Logo *string `pulumi:"logo"`
	// Determines if pagination is on or off.
	Paginate *bool `pulumi:"paginate"`
	// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
	PaginationRate *int `pulumi:"paginationRate"`
	// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
	RefreshRate *int `pulumi:"refreshRate"`
	// A list of one or more tags that filter which checks to display on the dashboard.
	Tags []string `pulumi:"tags"`
	// Set when to use AND operator for fetching dashboard tags.
	UseTagsAndOperator *bool `pulumi:"useTagsAndOperator"`
	// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
	Width *string `pulumi:"width"`
}

type DashboardState struct {
	// Determines how many checks to show per page.
	ChecksPerPage pulumi.IntPtrInput
	// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
	CustomDomain pulumi.StringPtrInput
	// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
	CustomUrl pulumi.StringPtrInput
	// HTML <meta> description for the dashboard.
	Description pulumi.StringPtrInput
	// A URL pointing to an image file to use as browser favicon.
	Favicon pulumi.StringPtrInput
	// A piece of text displayed at the top of your dashboard.
	Header pulumi.StringPtrInput
	// Show or hide the tags on the dashboard.
	HideTags pulumi.BoolPtrInput
	// A link to for the dashboard logo.
	Link pulumi.StringPtrInput
	// A URL pointing to an image file to use for the dashboard logo.
	Logo pulumi.StringPtrInput
	// Determines if pagination is on or off.
	Paginate pulumi.BoolPtrInput
	// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
	PaginationRate pulumi.IntPtrInput
	// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
	RefreshRate pulumi.IntPtrInput
	// A list of one or more tags that filter which checks to display on the dashboard.
	Tags pulumi.StringArrayInput
	// Set when to use AND operator for fetching dashboard tags.
	UseTagsAndOperator pulumi.BoolPtrInput
	// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
	Width pulumi.StringPtrInput
}

func (DashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardState)(nil)).Elem()
}

type dashboardArgs struct {
	// Determines how many checks to show per page.
	ChecksPerPage *int `pulumi:"checksPerPage"`
	// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
	CustomDomain *string `pulumi:"customDomain"`
	// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
	CustomUrl string `pulumi:"customUrl"`
	// HTML <meta> description for the dashboard.
	Description *string `pulumi:"description"`
	// A URL pointing to an image file to use as browser favicon.
	Favicon *string `pulumi:"favicon"`
	// A piece of text displayed at the top of your dashboard.
	Header *string `pulumi:"header"`
	// Show or hide the tags on the dashboard.
	HideTags *bool `pulumi:"hideTags"`
	// A link to for the dashboard logo.
	Link *string `pulumi:"link"`
	// A URL pointing to an image file to use for the dashboard logo.
	Logo *string `pulumi:"logo"`
	// Determines if pagination is on or off.
	Paginate *bool `pulumi:"paginate"`
	// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
	PaginationRate *int `pulumi:"paginationRate"`
	// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
	RefreshRate *int `pulumi:"refreshRate"`
	// A list of one or more tags that filter which checks to display on the dashboard.
	Tags []string `pulumi:"tags"`
	// Set when to use AND operator for fetching dashboard tags.
	UseTagsAndOperator *bool `pulumi:"useTagsAndOperator"`
	// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
	Width *string `pulumi:"width"`
}

// The set of arguments for constructing a Dashboard resource.
type DashboardArgs struct {
	// Determines how many checks to show per page.
	ChecksPerPage pulumi.IntPtrInput
	// A custom user domain, e.g. 'status.example.com'. See the docs on updating your DNS and SSL usage.
	CustomDomain pulumi.StringPtrInput
	// A subdomain name under 'checklyhq.com'. Needs to be unique across all users.
	CustomUrl pulumi.StringInput
	// HTML <meta> description for the dashboard.
	Description pulumi.StringPtrInput
	// A URL pointing to an image file to use as browser favicon.
	Favicon pulumi.StringPtrInput
	// A piece of text displayed at the top of your dashboard.
	Header pulumi.StringPtrInput
	// Show or hide the tags on the dashboard.
	HideTags pulumi.BoolPtrInput
	// A link to for the dashboard logo.
	Link pulumi.StringPtrInput
	// A URL pointing to an image file to use for the dashboard logo.
	Logo pulumi.StringPtrInput
	// Determines if pagination is on or off.
	Paginate pulumi.BoolPtrInput
	// How often to trigger pagination in seconds. Possible values `30`, `60` and `300`.
	PaginationRate pulumi.IntPtrInput
	// How often to refresh the dashboard in seconds. Possible values `60`, '300' and `600`.
	RefreshRate pulumi.IntPtrInput
	// A list of one or more tags that filter which checks to display on the dashboard.
	Tags pulumi.StringArrayInput
	// Set when to use AND operator for fetching dashboard tags.
	UseTagsAndOperator pulumi.BoolPtrInput
	// Determines whether to use the full screen or focus in the center. Possible values `FULL` and `960PX`.
	Width pulumi.StringPtrInput
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
