//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package wholepkg

import (
	otherpkg "github.com/dgodyna/gengo/examples/deepcopy-gen/output_tests/otherpkg"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ManualSlice) DeepCopyInto(out *ManualSlice) {
	{
		in := &in
		*out = in.DeepCopy()
		return
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualStruct) DeepCopyInto(out *ManualStruct) {
	*out = in.DeepCopy()
	return
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManualStruct_Alias) DeepCopyInto(out *ManualStruct_Alias) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManualStruct_Alias.
func (in *ManualStruct_Alias) DeepCopy() *ManualStruct_Alias {
	if in == nil {
		return nil
	}
	out := new(ManualStruct_Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_B) DeepCopyInto(out *Struct_B) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_B.
func (in *Struct_B) DeepCopy() *Struct_B {
	if in == nil {
		return nil
	}
	out := new(Struct_B)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Embed_Int) DeepCopyInto(out *Struct_Embed_Int) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Embed_Int.
func (in *Struct_Embed_Int) DeepCopy() *Struct_Embed_Int {
	if in == nil {
		return nil
	}
	out := new(Struct_Embed_Int)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Embed_ManualStruct) DeepCopyInto(out *Struct_Embed_ManualStruct) {
	*out = *in
	out.ManualStruct = in.ManualStruct.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Embed_ManualStruct.
func (in *Struct_Embed_ManualStruct) DeepCopy() *Struct_Embed_ManualStruct {
	if in == nil {
		return nil
	}
	out := new(Struct_Embed_ManualStruct)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Embed_Pointer) DeepCopyInto(out *Struct_Embed_Pointer) {
	*out = *in
	if in.int != nil {
		in, out := &in.int, &out.int
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Embed_Pointer.
func (in *Struct_Embed_Pointer) DeepCopy() *Struct_Embed_Pointer {
	if in == nil {
		return nil
	}
	out := new(Struct_Embed_Pointer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Embed_Struct_PrimitivePointers) DeepCopyInto(out *Struct_Embed_Struct_PrimitivePointers) {
	*out = *in
	in.Struct_PrimitivePointers.DeepCopyInto(&out.Struct_PrimitivePointers)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Embed_Struct_PrimitivePointers.
func (in *Struct_Embed_Struct_PrimitivePointers) DeepCopy() *Struct_Embed_Struct_PrimitivePointers {
	if in == nil {
		return nil
	}
	out := new(Struct_Embed_Struct_PrimitivePointers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Embed_Struct_Primitives) DeepCopyInto(out *Struct_Embed_Struct_Primitives) {
	*out = *in
	out.Struct_Primitives = in.Struct_Primitives
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Embed_Struct_Primitives.
func (in *Struct_Embed_Struct_Primitives) DeepCopy() *Struct_Embed_Struct_Primitives {
	if in == nil {
		return nil
	}
	out := new(Struct_Embed_Struct_Primitives)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Embed_Struct_Slices) DeepCopyInto(out *Struct_Embed_Struct_Slices) {
	*out = *in
	in.Struct_Slices.DeepCopyInto(&out.Struct_Slices)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Embed_Struct_Slices.
func (in *Struct_Embed_Struct_Slices) DeepCopy() *Struct_Embed_Struct_Slices {
	if in == nil {
		return nil
	}
	out := new(Struct_Embed_Struct_Slices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Empty) DeepCopyInto(out *Struct_Empty) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Empty.
func (in *Struct_Empty) DeepCopy() *Struct_Empty {
	if in == nil {
		return nil
	}
	out := new(Struct_Empty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Everything) DeepCopyInto(out *Struct_Everything) {
	*out = *in
	out.StructField = in.StructField
	out.EmptyStructField = in.EmptyStructField
	out.ManualStructField = in.ManualStructField.DeepCopy()
	out.ManualStructAliasField = in.ManualStructAliasField
	if in.BoolPtrField != nil {
		in, out := &in.BoolPtrField, &out.BoolPtrField
		*out = new(bool)
		**out = **in
	}
	if in.IntPtrField != nil {
		in, out := &in.IntPtrField, &out.IntPtrField
		*out = new(int)
		**out = **in
	}
	if in.StringPtrField != nil {
		in, out := &in.StringPtrField, &out.StringPtrField
		*out = new(string)
		**out = **in
	}
	if in.FloatPtrField != nil {
		in, out := &in.FloatPtrField, &out.FloatPtrField
		*out = new(float64)
		**out = **in
	}
	in.PrimitivePointersField.DeepCopyInto(&out.PrimitivePointersField)
	if in.ManualStructPtrField != nil {
		in, out := &in.ManualStructPtrField, &out.ManualStructPtrField
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.ManualStructAliasPtrField != nil {
		in, out := &in.ManualStructAliasPtrField, &out.ManualStructAliasPtrField
		*out = new(ManualStruct_Alias)
		**out = **in
	}
	if in.SliceBoolField != nil {
		in, out := &in.SliceBoolField, &out.SliceBoolField
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.SliceByteField != nil {
		in, out := &in.SliceByteField, &out.SliceByteField
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SliceIntField != nil {
		in, out := &in.SliceIntField, &out.SliceIntField
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.SliceStringField != nil {
		in, out := &in.SliceStringField, &out.SliceStringField
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SliceFloatField != nil {
		in, out := &in.SliceFloatField, &out.SliceFloatField
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	in.SlicesField.DeepCopyInto(&out.SlicesField)
	if in.SliceManualStructField != nil {
		in, out := &in.SliceManualStructField, &out.SliceManualStructField
		*out = make([]ManualStruct, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ManualSliceField = in.ManualSliceField.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Everything.
func (in *Struct_Everything) DeepCopy() *Struct_Everything {
	if in == nil {
		return nil
	}
	out := new(Struct_Everything)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_ExplicitObject) DeepCopyInto(out *Struct_ExplicitObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_ExplicitObject.
func (in *Struct_ExplicitObject) DeepCopy() *Struct_ExplicitObject {
	if in == nil {
		return nil
	}
	out := new(Struct_ExplicitObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new otherpkg.Object.
func (in *Struct_ExplicitObject) DeepCopyObject() otherpkg.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_ExplicitSelectorExplicitObject) DeepCopyInto(out *Struct_ExplicitSelectorExplicitObject) {
	*out = *in
	out.Struct_TypeMeta = in.Struct_TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_ExplicitSelectorExplicitObject.
func (in *Struct_ExplicitSelectorExplicitObject) DeepCopy() *Struct_ExplicitSelectorExplicitObject {
	if in == nil {
		return nil
	}
	out := new(Struct_ExplicitSelectorExplicitObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new otherpkg.Object.
func (in *Struct_ExplicitSelectorExplicitObject) DeepCopyObject() otherpkg.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopySelector is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Struct_ExplicitSelectorExplicitObject) DeepCopySelector() Selector {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Interfaces) DeepCopyInto(out *Struct_Interfaces) {
	*out = *in
	if in.ObjectField != nil {
		out.ObjectField = in.ObjectField.DeepCopyObject()
	}
	if in.NilObjectField != nil {
		out.NilObjectField = in.NilObjectField.DeepCopyObject()
	}
	if in.SelectorField != nil {
		out.SelectorField = in.SelectorField.DeepCopySelector()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Interfaces.
func (in *Struct_Interfaces) DeepCopy() *Struct_Interfaces {
	if in == nil {
		return nil
	}
	out := new(Struct_Interfaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_NonPointerExplicitObject) DeepCopyInto(out *Struct_NonPointerExplicitObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_NonPointerExplicitObject.
func (in *Struct_NonPointerExplicitObject) DeepCopy() *Struct_NonPointerExplicitObject {
	if in == nil {
		return nil
	}
	out := new(Struct_NonPointerExplicitObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new otherpkg.Object.
func (in Struct_NonPointerExplicitObject) DeepCopyObject() otherpkg.Object {
	return *in.DeepCopy()
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_ObjectAndList) DeepCopyInto(out *Struct_ObjectAndList) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_ObjectAndList.
func (in *Struct_ObjectAndList) DeepCopy() *Struct_ObjectAndList {
	if in == nil {
		return nil
	}
	out := new(Struct_ObjectAndList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyList is an autogenerated deepcopy function, copying the receiver, creating a new otherpkg.List.
func (in *Struct_ObjectAndList) DeepCopyList() otherpkg.List {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new otherpkg.Object.
func (in *Struct_ObjectAndList) DeepCopyObject() otherpkg.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_ObjectAndObject) DeepCopyInto(out *Struct_ObjectAndObject) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_ObjectAndObject.
func (in *Struct_ObjectAndObject) DeepCopy() *Struct_ObjectAndObject {
	if in == nil {
		return nil
	}
	out := new(Struct_ObjectAndObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new otherpkg.Object.
func (in *Struct_ObjectAndObject) DeepCopyObject() otherpkg.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_PrimitivePointers) DeepCopyInto(out *Struct_PrimitivePointers) {
	*out = *in
	if in.BoolPtrField != nil {
		in, out := &in.BoolPtrField, &out.BoolPtrField
		*out = new(bool)
		**out = **in
	}
	if in.IntPtrField != nil {
		in, out := &in.IntPtrField, &out.IntPtrField
		*out = new(int)
		**out = **in
	}
	if in.StringPtrField != nil {
		in, out := &in.StringPtrField, &out.StringPtrField
		*out = new(string)
		**out = **in
	}
	if in.FloatPtrField != nil {
		in, out := &in.FloatPtrField, &out.FloatPtrField
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_PrimitivePointers.
func (in *Struct_PrimitivePointers) DeepCopy() *Struct_PrimitivePointers {
	if in == nil {
		return nil
	}
	out := new(Struct_PrimitivePointers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_PrimitivePointers_Alias) DeepCopyInto(out *Struct_PrimitivePointers_Alias) {
	*out = *in
	if in.BoolPtrField != nil {
		in, out := &in.BoolPtrField, &out.BoolPtrField
		*out = new(bool)
		**out = **in
	}
	if in.IntPtrField != nil {
		in, out := &in.IntPtrField, &out.IntPtrField
		*out = new(int)
		**out = **in
	}
	if in.StringPtrField != nil {
		in, out := &in.StringPtrField, &out.StringPtrField
		*out = new(string)
		**out = **in
	}
	if in.FloatPtrField != nil {
		in, out := &in.FloatPtrField, &out.FloatPtrField
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_PrimitivePointers_Alias.
func (in *Struct_PrimitivePointers_Alias) DeepCopy() *Struct_PrimitivePointers_Alias {
	if in == nil {
		return nil
	}
	out := new(Struct_PrimitivePointers_Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Primitives) DeepCopyInto(out *Struct_Primitives) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Primitives.
func (in *Struct_Primitives) DeepCopy() *Struct_Primitives {
	if in == nil {
		return nil
	}
	out := new(Struct_Primitives)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Primitives_Alias) DeepCopyInto(out *Struct_Primitives_Alias) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Primitives_Alias.
func (in *Struct_Primitives_Alias) DeepCopy() *Struct_Primitives_Alias {
	if in == nil {
		return nil
	}
	out := new(Struct_Primitives_Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Slices) DeepCopyInto(out *Struct_Slices) {
	*out = *in
	if in.SliceBoolField != nil {
		in, out := &in.SliceBoolField, &out.SliceBoolField
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.SliceByteField != nil {
		in, out := &in.SliceByteField, &out.SliceByteField
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SliceIntField != nil {
		in, out := &in.SliceIntField, &out.SliceIntField
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.SliceStringField != nil {
		in, out := &in.SliceStringField, &out.SliceStringField
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SliceFloatField != nil {
		in, out := &in.SliceFloatField, &out.SliceFloatField
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	if in.SliceStructPrimitivesField != nil {
		in, out := &in.SliceStructPrimitivesField, &out.SliceStructPrimitivesField
		*out = make([]Struct_Primitives, len(*in))
		copy(*out, *in)
	}
	if in.SliceStructPrimitivesAliasField != nil {
		in, out := &in.SliceStructPrimitivesAliasField, &out.SliceStructPrimitivesAliasField
		*out = make([]Struct_Primitives_Alias, len(*in))
		copy(*out, *in)
	}
	if in.SliceStructPrimitivePointersField != nil {
		in, out := &in.SliceStructPrimitivePointersField, &out.SliceStructPrimitivePointersField
		*out = make([]Struct_PrimitivePointers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SliceStructPrimitivePointersAliasField != nil {
		in, out := &in.SliceStructPrimitivePointersAliasField, &out.SliceStructPrimitivePointersAliasField
		*out = make([]Struct_PrimitivePointers_Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SliceSliceIntField != nil {
		in, out := &in.SliceSliceIntField, &out.SliceSliceIntField
		*out = make([][]int, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]int, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.SliceManualStructField != nil {
		in, out := &in.SliceManualStructField, &out.SliceManualStructField
		*out = make([]ManualStruct, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ManualSliceField = in.ManualSliceField.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Slices.
func (in *Struct_Slices) DeepCopy() *Struct_Slices {
	if in == nil {
		return nil
	}
	out := new(Struct_Slices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Slices_Alias) DeepCopyInto(out *Struct_Slices_Alias) {
	*out = *in
	if in.SliceBoolField != nil {
		in, out := &in.SliceBoolField, &out.SliceBoolField
		*out = make([]bool, len(*in))
		copy(*out, *in)
	}
	if in.SliceByteField != nil {
		in, out := &in.SliceByteField, &out.SliceByteField
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SliceIntField != nil {
		in, out := &in.SliceIntField, &out.SliceIntField
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.SliceStringField != nil {
		in, out := &in.SliceStringField, &out.SliceStringField
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SliceFloatField != nil {
		in, out := &in.SliceFloatField, &out.SliceFloatField
		*out = make([]float64, len(*in))
		copy(*out, *in)
	}
	if in.SliceStructPrimitivesField != nil {
		in, out := &in.SliceStructPrimitivesField, &out.SliceStructPrimitivesField
		*out = make([]Struct_Primitives, len(*in))
		copy(*out, *in)
	}
	if in.SliceStructPrimitivesAliasField != nil {
		in, out := &in.SliceStructPrimitivesAliasField, &out.SliceStructPrimitivesAliasField
		*out = make([]Struct_Primitives_Alias, len(*in))
		copy(*out, *in)
	}
	if in.SliceStructPrimitivePointersField != nil {
		in, out := &in.SliceStructPrimitivePointersField, &out.SliceStructPrimitivePointersField
		*out = make([]Struct_PrimitivePointers, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SliceStructPrimitivePointersAliasField != nil {
		in, out := &in.SliceStructPrimitivePointersAliasField, &out.SliceStructPrimitivePointersAliasField
		*out = make([]Struct_PrimitivePointers_Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SliceSliceIntField != nil {
		in, out := &in.SliceSliceIntField, &out.SliceSliceIntField
		*out = make([][]int, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make([]int, len(*in))
				copy(*out, *in)
			}
		}
	}
	if in.SliceManualStructField != nil {
		in, out := &in.SliceManualStructField, &out.SliceManualStructField
		*out = make([]ManualStruct, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ManualSliceField = in.ManualSliceField.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Slices_Alias.
func (in *Struct_Slices_Alias) DeepCopy() *Struct_Slices_Alias {
	if in == nil {
		return nil
	}
	out := new(Struct_Slices_Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Struct_PrimitivePointers) DeepCopyInto(out *Struct_Struct_PrimitivePointers) {
	*out = *in
	in.StructField.DeepCopyInto(&out.StructField)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Struct_PrimitivePointers.
func (in *Struct_Struct_PrimitivePointers) DeepCopy() *Struct_Struct_PrimitivePointers {
	if in == nil {
		return nil
	}
	out := new(Struct_Struct_PrimitivePointers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Struct_Primitives) DeepCopyInto(out *Struct_Struct_Primitives) {
	*out = *in
	out.StructField = in.StructField
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Struct_Primitives.
func (in *Struct_Struct_Primitives) DeepCopy() *Struct_Struct_Primitives {
	if in == nil {
		return nil
	}
	out := new(Struct_Struct_Primitives)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Struct_Struct_Slices) DeepCopyInto(out *Struct_Struct_Slices) {
	*out = *in
	in.StructField.DeepCopyInto(&out.StructField)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Struct_Struct_Slices.
func (in *Struct_Struct_Slices) DeepCopy() *Struct_Struct_Slices {
	if in == nil {
		return nil
	}
	out := new(Struct_Struct_Slices)
	in.DeepCopyInto(out)
	return out
}
