// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package checkly

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Snippet struct {
	pulumi.CustomResourceState

	// The name of the snippet.
	Name pulumi.StringOutput `pulumi:"name"`
	// Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
	Script pulumi.StringOutput `pulumi:"script"`
}

// NewSnippet registers a new resource with the given unique name, arguments, and options.
func NewSnippet(ctx *pulumi.Context,
	name string, args *SnippetArgs, opts ...pulumi.ResourceOption) (*Snippet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Script == nil {
		return nil, errors.New("invalid value for required argument 'Script'")
	}
	var resource Snippet
	err := ctx.RegisterResource("checkly:index/snippet:Snippet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnippet gets an existing Snippet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnippet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnippetState, opts ...pulumi.ResourceOption) (*Snippet, error) {
	var resource Snippet
	err := ctx.ReadResource("checkly:index/snippet:Snippet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Snippet resources.
type snippetState struct {
	// The name of the snippet.
	Name *string `pulumi:"name"`
	// Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
	Script *string `pulumi:"script"`
}

type SnippetState struct {
	// The name of the snippet.
	Name pulumi.StringPtrInput
	// Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
	Script pulumi.StringPtrInput
}

func (SnippetState) ElementType() reflect.Type {
	return reflect.TypeOf((*snippetState)(nil)).Elem()
}

type snippetArgs struct {
	// The name of the snippet.
	Name *string `pulumi:"name"`
	// Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
	Script string `pulumi:"script"`
}

// The set of arguments for constructing a Snippet resource.
type SnippetArgs struct {
	// The name of the snippet.
	Name pulumi.StringPtrInput
	// Your Node.js code that interacts with the API check lifecycle, or functions as a partial for browser checks.
	Script pulumi.StringInput
}

func (SnippetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snippetArgs)(nil)).Elem()
}

type SnippetInput interface {
	pulumi.Input

	ToSnippetOutput() SnippetOutput
	ToSnippetOutputWithContext(ctx context.Context) SnippetOutput
}

func (*Snippet) ElementType() reflect.Type {
	return reflect.TypeOf((**Snippet)(nil)).Elem()
}

func (i *Snippet) ToSnippetOutput() SnippetOutput {
	return i.ToSnippetOutputWithContext(context.Background())
}

func (i *Snippet) ToSnippetOutputWithContext(ctx context.Context) SnippetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnippetOutput)
}

// SnippetArrayInput is an input type that accepts SnippetArray and SnippetArrayOutput values.
// You can construct a concrete instance of `SnippetArrayInput` via:
//
//          SnippetArray{ SnippetArgs{...} }
type SnippetArrayInput interface {
	pulumi.Input

	ToSnippetArrayOutput() SnippetArrayOutput
	ToSnippetArrayOutputWithContext(context.Context) SnippetArrayOutput
}

type SnippetArray []SnippetInput

func (SnippetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snippet)(nil)).Elem()
}

func (i SnippetArray) ToSnippetArrayOutput() SnippetArrayOutput {
	return i.ToSnippetArrayOutputWithContext(context.Background())
}

func (i SnippetArray) ToSnippetArrayOutputWithContext(ctx context.Context) SnippetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnippetArrayOutput)
}

// SnippetMapInput is an input type that accepts SnippetMap and SnippetMapOutput values.
// You can construct a concrete instance of `SnippetMapInput` via:
//
//          SnippetMap{ "key": SnippetArgs{...} }
type SnippetMapInput interface {
	pulumi.Input

	ToSnippetMapOutput() SnippetMapOutput
	ToSnippetMapOutputWithContext(context.Context) SnippetMapOutput
}

type SnippetMap map[string]SnippetInput

func (SnippetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snippet)(nil)).Elem()
}

func (i SnippetMap) ToSnippetMapOutput() SnippetMapOutput {
	return i.ToSnippetMapOutputWithContext(context.Background())
}

func (i SnippetMap) ToSnippetMapOutputWithContext(ctx context.Context) SnippetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnippetMapOutput)
}

type SnippetOutput struct{ *pulumi.OutputState }

func (SnippetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Snippet)(nil)).Elem()
}

func (o SnippetOutput) ToSnippetOutput() SnippetOutput {
	return o
}

func (o SnippetOutput) ToSnippetOutputWithContext(ctx context.Context) SnippetOutput {
	return o
}

type SnippetArrayOutput struct{ *pulumi.OutputState }

func (SnippetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Snippet)(nil)).Elem()
}

func (o SnippetArrayOutput) ToSnippetArrayOutput() SnippetArrayOutput {
	return o
}

func (o SnippetArrayOutput) ToSnippetArrayOutputWithContext(ctx context.Context) SnippetArrayOutput {
	return o
}

func (o SnippetArrayOutput) Index(i pulumi.IntInput) SnippetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Snippet {
		return vs[0].([]*Snippet)[vs[1].(int)]
	}).(SnippetOutput)
}

type SnippetMapOutput struct{ *pulumi.OutputState }

func (SnippetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Snippet)(nil)).Elem()
}

func (o SnippetMapOutput) ToSnippetMapOutput() SnippetMapOutput {
	return o
}

func (o SnippetMapOutput) ToSnippetMapOutputWithContext(ctx context.Context) SnippetMapOutput {
	return o
}

func (o SnippetMapOutput) MapIndex(k pulumi.StringInput) SnippetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Snippet {
		return vs[0].(map[string]*Snippet)[vs[1].(string)]
	}).(SnippetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnippetInput)(nil)).Elem(), &Snippet{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnippetArrayInput)(nil)).Elem(), SnippetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnippetMapInput)(nil)).Elem(), SnippetMap{})
	pulumi.RegisterOutputType(SnippetOutput{})
	pulumi.RegisterOutputType(SnippetArrayOutput{})
	pulumi.RegisterOutputType(SnippetMapOutput{})
}
