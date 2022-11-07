// Code generated by test DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elementtype

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ElementType struct {
	pulumi.CustomResourceState

	ElementType_ ElementTypeTypePtrOutput `pulumi:"elementType"`
}

// NewElementType registers a new resource with the given unique name, arguments, and options.
func NewElementType(ctx *pulumi.Context,
	name string, args *ElementTypeArgs, opts ...pulumi.ResourceOption) (*ElementType, error) {
	if args == nil {
		args = &ElementTypeArgs{}
	}

	var resource ElementType
	err := ctx.RegisterResource("repro:elementType:ElementType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetElementType gets an existing ElementType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetElementType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ElementTypeState, opts ...pulumi.ResourceOption) (*ElementType, error) {
	var resource ElementType
	err := ctx.ReadResource("repro:elementType:ElementType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ElementType resources.
type elementTypeState struct {
}

type ElementTypeState struct {
}

func (ElementTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*elementTypeState)(nil)).Elem()
}

type elementTypeArgs struct {
}

// The set of arguments for constructing a ElementType resource.
type ElementTypeArgs struct {
}

func (ElementTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*elementTypeArgs)(nil)).Elem()
}

type ElementTypeInput interface {
	pulumi.Input

	ToElementTypeOutput() ElementTypeOutput
	ToElementTypeOutputWithContext(ctx context.Context) ElementTypeOutput
}

func (*ElementType) ElementType() reflect.Type {
	return reflect.TypeOf((**ElementType)(nil)).Elem()
}

func (i *ElementType) ToElementTypeOutput() ElementTypeOutput {
	return i.ToElementTypeOutputWithContext(context.Background())
}

func (i *ElementType) ToElementTypeOutputWithContext(ctx context.Context) ElementTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElementTypeOutput)
}

// ElementTypeArrayInput is an input type that accepts ElementTypeArray and ElementTypeArrayOutput values.
// You can construct a concrete instance of `ElementTypeArrayInput` via:
//
//	ElementTypeArray{ ElementTypeArgs{...} }
type ElementTypeArrayInput interface {
	pulumi.Input

	ToElementTypeArrayOutput() ElementTypeArrayOutput
	ToElementTypeArrayOutputWithContext(context.Context) ElementTypeArrayOutput
}

type ElementTypeArray []ElementTypeInput

func (ElementTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElementType)(nil)).Elem()
}

func (i ElementTypeArray) ToElementTypeArrayOutput() ElementTypeArrayOutput {
	return i.ToElementTypeArrayOutputWithContext(context.Background())
}

func (i ElementTypeArray) ToElementTypeArrayOutputWithContext(ctx context.Context) ElementTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElementTypeArrayOutput)
}

// ElementTypeMapInput is an input type that accepts ElementTypeMap and ElementTypeMapOutput values.
// You can construct a concrete instance of `ElementTypeMapInput` via:
//
//	ElementTypeMap{ "key": ElementTypeArgs{...} }
type ElementTypeMapInput interface {
	pulumi.Input

	ToElementTypeMapOutput() ElementTypeMapOutput
	ToElementTypeMapOutputWithContext(context.Context) ElementTypeMapOutput
}

type ElementTypeMap map[string]ElementTypeInput

func (ElementTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElementType)(nil)).Elem()
}

func (i ElementTypeMap) ToElementTypeMapOutput() ElementTypeMapOutput {
	return i.ToElementTypeMapOutputWithContext(context.Background())
}

func (i ElementTypeMap) ToElementTypeMapOutputWithContext(ctx context.Context) ElementTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ElementTypeMapOutput)
}

type ElementTypeOutput struct{ *pulumi.OutputState }

func (ElementTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ElementType)(nil)).Elem()
}

func (o ElementTypeOutput) ToElementTypeOutput() ElementTypeOutput {
	return o
}

func (o ElementTypeOutput) ToElementTypeOutputWithContext(ctx context.Context) ElementTypeOutput {
	return o
}

func (o ElementTypeOutput) GetElementType_() ElementTypeTypePtrOutput {
	return o.ApplyT(func(v *ElementType) ElementTypeTypePtrOutput { return v.ElementType_ }).(ElementTypeTypePtrOutput)
}

type ElementTypeArrayOutput struct{ *pulumi.OutputState }

func (ElementTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ElementType)(nil)).Elem()
}

func (o ElementTypeArrayOutput) ToElementTypeArrayOutput() ElementTypeArrayOutput {
	return o
}

func (o ElementTypeArrayOutput) ToElementTypeArrayOutputWithContext(ctx context.Context) ElementTypeArrayOutput {
	return o
}

func (o ElementTypeArrayOutput) Index(i pulumi.IntInput) ElementTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ElementType {
		return vs[0].([]*ElementType)[vs[1].(int)]
	}).(ElementTypeOutput)
}

type ElementTypeMapOutput struct{ *pulumi.OutputState }

func (ElementTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ElementType)(nil)).Elem()
}

func (o ElementTypeMapOutput) ToElementTypeMapOutput() ElementTypeMapOutput {
	return o
}

func (o ElementTypeMapOutput) ToElementTypeMapOutputWithContext(ctx context.Context) ElementTypeMapOutput {
	return o
}

func (o ElementTypeMapOutput) MapIndex(k pulumi.StringInput) ElementTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ElementType {
		return vs[0].(map[string]*ElementType)[vs[1].(string)]
	}).(ElementTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ElementTypeInput)(nil)).Elem(), &ElementType{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElementTypeArrayInput)(nil)).Elem(), ElementTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ElementTypeMapInput)(nil)).Elem(), ElementTypeMap{})
	pulumi.RegisterOutputType(ElementTypeOutput{})
	pulumi.RegisterOutputType(ElementTypeArrayOutput{})
	pulumi.RegisterOutputType(ElementTypeMapOutput{})
}
