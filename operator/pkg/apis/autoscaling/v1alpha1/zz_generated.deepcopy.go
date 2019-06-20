// +build !ignore_autogenerated

/*
Copyright 2019 The Alameda Authors.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaContainer) DeepCopyInto(out *AlamedaContainer) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaContainer.
func (in *AlamedaContainer) DeepCopy() *AlamedaContainer {
	if in == nil {
		return nil
	}
	out := new(AlamedaContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaController) DeepCopyInto(out *AlamedaController) {
	*out = *in
	if in.Deployments != nil {
		in, out := &in.Deployments, &out.Deployments
		*out = make(map[string]AlamedaResource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.DeploymentConfigs != nil {
		in, out := &in.DeploymentConfigs, &out.DeploymentConfigs
		*out = make(map[string]AlamedaResource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaController.
func (in *AlamedaController) DeepCopy() *AlamedaController {
	if in == nil {
		return nil
	}
	out := new(AlamedaController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaPod) DeepCopyInto(out *AlamedaPod) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]AlamedaContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaPod.
func (in *AlamedaPod) DeepCopy() *AlamedaPod {
	if in == nil {
		return nil
	}
	out := new(AlamedaPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaRecommendation) DeepCopyInto(out *AlamedaRecommendation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaRecommendation.
func (in *AlamedaRecommendation) DeepCopy() *AlamedaRecommendation {
	if in == nil {
		return nil
	}
	out := new(AlamedaRecommendation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaRecommendation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaRecommendationList) DeepCopyInto(out *AlamedaRecommendationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlamedaRecommendation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaRecommendationList.
func (in *AlamedaRecommendationList) DeepCopy() *AlamedaRecommendationList {
	if in == nil {
		return nil
	}
	out := new(AlamedaRecommendationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaRecommendationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaRecommendationSpec) DeepCopyInto(out *AlamedaRecommendationSpec) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]AlamedaContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaRecommendationSpec.
func (in *AlamedaRecommendationSpec) DeepCopy() *AlamedaRecommendationSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaRecommendationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaRecommendationStatus) DeepCopyInto(out *AlamedaRecommendationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaRecommendationStatus.
func (in *AlamedaRecommendationStatus) DeepCopy() *AlamedaRecommendationStatus {
	if in == nil {
		return nil
	}
	out := new(AlamedaRecommendationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaResource) DeepCopyInto(out *AlamedaResource) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make(map[string]AlamedaPod, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaResource.
func (in *AlamedaResource) DeepCopy() *AlamedaResource {
	if in == nil {
		return nil
	}
	out := new(AlamedaResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaScaler) DeepCopyInto(out *AlamedaScaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaScaler.
func (in *AlamedaScaler) DeepCopy() *AlamedaScaler {
	if in == nil {
		return nil
	}
	out := new(AlamedaScaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaScaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaScalerList) DeepCopyInto(out *AlamedaScalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlamedaScaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaScalerList.
func (in *AlamedaScalerList) DeepCopy() *AlamedaScalerList {
	if in == nil {
		return nil
	}
	out := new(AlamedaScalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlamedaScalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaScalerSpec) DeepCopyInto(out *AlamedaScalerSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.ScalingTool.DeepCopyInto(&out.ScalingTool)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaScalerSpec.
func (in *AlamedaScalerSpec) DeepCopy() *AlamedaScalerSpec {
	if in == nil {
		return nil
	}
	out := new(AlamedaScalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlamedaScalerStatus) DeepCopyInto(out *AlamedaScalerStatus) {
	*out = *in
	in.AlamedaController.DeepCopyInto(&out.AlamedaController)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlamedaScalerStatus.
func (in *AlamedaScalerStatus) DeepCopy() *AlamedaScalerStatus {
	if in == nil {
		return nil
	}
	out := new(AlamedaScalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExecutionStrategy) DeepCopyInto(out *ExecutionStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExecutionStrategy.
func (in *ExecutionStrategy) DeepCopy() *ExecutionStrategy {
	if in == nil {
		return nil
	}
	out := new(ExecutionStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingToolSpec) DeepCopyInto(out *ScalingToolSpec) {
	*out = *in
	if in.ExecutionStrategy != nil {
		in, out := &in.ExecutionStrategy, &out.ExecutionStrategy
		*out = new(ExecutionStrategy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingToolSpec.
func (in *ScalingToolSpec) DeepCopy() *ScalingToolSpec {
	if in == nil {
		return nil
	}
	out := new(ScalingToolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingToolstruct) DeepCopyInto(out *ScalingToolstruct) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingToolstruct.
func (in *ScalingToolstruct) DeepCopy() *ScalingToolstruct {
	if in == nil {
		return nil
	}
	out := new(ScalingToolstruct)
	in.DeepCopyInto(out)
	return out
}
