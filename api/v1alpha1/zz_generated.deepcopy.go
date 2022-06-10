//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionWebhook) DeepCopyInto(out *ConversionWebhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionWebhook.
func (in *ConversionWebhook) DeepCopy() *ConversionWebhook {
	if in == nil {
		return nil
	}
	out := new(ConversionWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConversionWebhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionWebhookList) DeepCopyInto(out *ConversionWebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConversionWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionWebhookList.
func (in *ConversionWebhookList) DeepCopy() *ConversionWebhookList {
	if in == nil {
		return nil
	}
	out := new(ConversionWebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConversionWebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionWebhookSpec) DeepCopyInto(out *ConversionWebhookSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionWebhookSpec.
func (in *ConversionWebhookSpec) DeepCopy() *ConversionWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(ConversionWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConversionWebhookStatus) DeepCopyInto(out *ConversionWebhookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConversionWebhookStatus.
func (in *ConversionWebhookStatus) DeepCopy() *ConversionWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(ConversionWebhookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionDeployment) DeepCopyInto(out *CustomResourceDefinitionDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionDeployment.
func (in *CustomResourceDefinitionDeployment) DeepCopy() *CustomResourceDefinitionDeployment {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDefinitionDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionDeploymentList) DeepCopyInto(out *CustomResourceDefinitionDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomResourceDefinitionDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionDeploymentList.
func (in *CustomResourceDefinitionDeploymentList) DeepCopy() *CustomResourceDefinitionDeploymentList {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDefinitionDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionDeploymentSpec) DeepCopyInto(out *CustomResourceDefinitionDeploymentSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionDeploymentSpec.
func (in *CustomResourceDefinitionDeploymentSpec) DeepCopy() *CustomResourceDefinitionDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionDeploymentStatus) DeepCopyInto(out *CustomResourceDefinitionDeploymentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionDeploymentStatus.
func (in *CustomResourceDefinitionDeploymentStatus) DeepCopy() *CustomResourceDefinitionDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSource) DeepCopyInto(out *CustomResourceDefinitionSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSource.
func (in *CustomResourceDefinitionSource) DeepCopy() *CustomResourceDefinitionSource {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDefinitionSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSourceList) DeepCopyInto(out *CustomResourceDefinitionSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomResourceDefinitionSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSourceList.
func (in *CustomResourceDefinitionSourceList) DeepCopy() *CustomResourceDefinitionSourceList {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomResourceDefinitionSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSourceSpec) DeepCopyInto(out *CustomResourceDefinitionSourceSpec) {
	*out = *in
	if in.CRDNames != nil {
		in, out := &in.CRDNames, &out.CRDNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSourceSpec.
func (in *CustomResourceDefinitionSourceSpec) DeepCopy() *CustomResourceDefinitionSourceSpec {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomResourceDefinitionSourceStatus) DeepCopyInto(out *CustomResourceDefinitionSourceStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomResourceDefinitionSourceStatus.
func (in *CustomResourceDefinitionSourceStatus) DeepCopy() *CustomResourceDefinitionSourceStatus {
	if in == nil {
		return nil
	}
	out := new(CustomResourceDefinitionSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhook) DeepCopyInto(out *MutatingWebhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhook.
func (in *MutatingWebhook) DeepCopy() *MutatingWebhook {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MutatingWebhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookList) DeepCopyInto(out *MutatingWebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MutatingWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookList.
func (in *MutatingWebhookList) DeepCopy() *MutatingWebhookList {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MutatingWebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookSpec) DeepCopyInto(out *MutatingWebhookSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookSpec.
func (in *MutatingWebhookSpec) DeepCopy() *MutatingWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MutatingWebhookStatus) DeepCopyInto(out *MutatingWebhookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MutatingWebhookStatus.
func (in *MutatingWebhookStatus) DeepCopy() *MutatingWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(MutatingWebhookStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhook) DeepCopyInto(out *ValidatingWebhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhook.
func (in *ValidatingWebhook) DeepCopy() *ValidatingWebhook {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidatingWebhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookList) DeepCopyInto(out *ValidatingWebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ValidatingWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookList.
func (in *ValidatingWebhookList) DeepCopy() *ValidatingWebhookList {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValidatingWebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookSpec) DeepCopyInto(out *ValidatingWebhookSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookSpec.
func (in *ValidatingWebhookSpec) DeepCopy() *ValidatingWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidatingWebhookStatus) DeepCopyInto(out *ValidatingWebhookStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidatingWebhookStatus.
func (in *ValidatingWebhookStatus) DeepCopy() *ValidatingWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(ValidatingWebhookStatus)
	in.DeepCopyInto(out)
	return out
}
