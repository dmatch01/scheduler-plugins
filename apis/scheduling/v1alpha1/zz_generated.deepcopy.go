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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticQuota) DeepCopyInto(out *ElasticQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticQuota.
func (in *ElasticQuota) DeepCopy() *ElasticQuota {
	if in == nil {
		return nil
	}
	out := new(ElasticQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticQuotaList) DeepCopyInto(out *ElasticQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ElasticQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticQuotaList.
func (in *ElasticQuotaList) DeepCopy() *ElasticQuotaList {
	if in == nil {
		return nil
	}
	out := new(ElasticQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ElasticQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticQuotaSpec) DeepCopyInto(out *ElasticQuotaSpec) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticQuotaSpec.
func (in *ElasticQuotaSpec) DeepCopy() *ElasticQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ElasticQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticQuotaStatus) DeepCopyInto(out *ElasticQuotaStatus) {
	*out = *in
	if in.Used != nil {
		in, out := &in.Used, &out.Used
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticQuotaStatus.
func (in *ElasticQuotaStatus) DeepCopy() *ElasticQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ElasticQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroup) DeepCopyInto(out *PodGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroup.
func (in *PodGroup) DeepCopy() *PodGroup {
	if in == nil {
		return nil
	}
	out := new(PodGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupList) DeepCopyInto(out *PodGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupList.
func (in *PodGroupList) DeepCopy() *PodGroupList {
	if in == nil {
		return nil
	}
	out := new(PodGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupSpec) DeepCopyInto(out *PodGroupSpec) {
	*out = *in
	if in.MinResources != nil {
		in, out := &in.MinResources, &out.MinResources
		*out = new(map[v1.ResourceName]resource.Quantity)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[v1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.ScheduleTimeoutSeconds != nil {
		in, out := &in.ScheduleTimeoutSeconds, &out.ScheduleTimeoutSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupSpec.
func (in *PodGroupSpec) DeepCopy() *PodGroupSpec {
	if in == nil {
		return nil
	}
	out := new(PodGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupStatus) DeepCopyInto(out *PodGroupStatus) {
	*out = *in
	in.ScheduleStartTime.DeepCopyInto(&out.ScheduleStartTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupStatus.
func (in *PodGroupStatus) DeepCopy() *PodGroupStatus {
	if in == nil {
		return nil
	}
	out := new(PodGroupStatus)
	in.DeepCopyInto(out)
	return out
}
