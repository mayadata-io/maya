// +build !ignore_autogenerated

/*
Copyright 2017 The OpenEBS Authors

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorPool).DeepCopyInto(out.(*CStorPool))
			return nil
		}, InType: reflect.TypeOf(&CStorPool{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorPoolAttr).DeepCopyInto(out.(*CStorPoolAttr))
			return nil
		}, InType: reflect.TypeOf(&CStorPoolAttr{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorPoolList).DeepCopyInto(out.(*CStorPoolList))
			return nil
		}, InType: reflect.TypeOf(&CStorPoolList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorPoolSpec).DeepCopyInto(out.(*CStorPoolSpec))
			return nil
		}, InType: reflect.TypeOf(&CStorPoolSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorVolumeReplica).DeepCopyInto(out.(*CStorVolumeReplica))
			return nil
		}, InType: reflect.TypeOf(&CStorVolumeReplica{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorVolumeReplicaList).DeepCopyInto(out.(*CStorVolumeReplicaList))
			return nil
		}, InType: reflect.TypeOf(&CStorVolumeReplicaList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*CStorVolumeReplicaSpec).DeepCopyInto(out.(*CStorVolumeReplicaSpec))
			return nil
		}, InType: reflect.TypeOf(&CStorVolumeReplicaSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*DiskAttr).DeepCopyInto(out.(*DiskAttr))
			return nil
		}, InType: reflect.TypeOf(&DiskAttr{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Policy).DeepCopyInto(out.(*Policy))
			return nil
		}, InType: reflect.TypeOf(&Policy{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*RunTasks).DeepCopyInto(out.(*RunTasks))
			return nil
		}, InType: reflect.TypeOf(&RunTasks{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StoragePool).DeepCopyInto(out.(*StoragePool))
			return nil
		}, InType: reflect.TypeOf(&StoragePool{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StoragePoolClaim).DeepCopyInto(out.(*StoragePoolClaim))
			return nil
		}, InType: reflect.TypeOf(&StoragePoolClaim{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StoragePoolClaimList).DeepCopyInto(out.(*StoragePoolClaimList))
			return nil
		}, InType: reflect.TypeOf(&StoragePoolClaimList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StoragePoolClaimSpec).DeepCopyInto(out.(*StoragePoolClaimSpec))
			return nil
		}, InType: reflect.TypeOf(&StoragePoolClaimSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StoragePoolList).DeepCopyInto(out.(*StoragePoolList))
			return nil
		}, InType: reflect.TypeOf(&StoragePoolList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*StoragePoolSpec).DeepCopyInto(out.(*StoragePoolSpec))
			return nil
		}, InType: reflect.TypeOf(&StoragePoolSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Task).DeepCopyInto(out.(*Task))
			return nil
		}, InType: reflect.TypeOf(&Task{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeParameterGroup).DeepCopyInto(out.(*VolumeParameterGroup))
			return nil
		}, InType: reflect.TypeOf(&VolumeParameterGroup{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeParameterGroupList).DeepCopyInto(out.(*VolumeParameterGroupList))
			return nil
		}, InType: reflect.TypeOf(&VolumeParameterGroupList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*VolumeParameterGroupSpec).DeepCopyInto(out.(*VolumeParameterGroupSpec))
			return nil
		}, InType: reflect.TypeOf(&VolumeParameterGroupSpec{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPool) DeepCopyInto(out *CStorPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPool.
func (in *CStorPool) DeepCopy() *CStorPool {
	if in == nil {
		return nil
	}
	out := new(CStorPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolAttr) DeepCopyInto(out *CStorPoolAttr) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolAttr.
func (in *CStorPoolAttr) DeepCopy() *CStorPoolAttr {
	if in == nil {
		return nil
	}
	out := new(CStorPoolAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolList) DeepCopyInto(out *CStorPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolList.
func (in *CStorPoolList) DeepCopy() *CStorPoolList {
	if in == nil {
		return nil
	}
	out := new(CStorPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPoolSpec) DeepCopyInto(out *CStorPoolSpec) {
	*out = *in
	in.Disks.DeepCopyInto(&out.Disks)
	out.PoolSpec = in.PoolSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPoolSpec.
func (in *CStorPoolSpec) DeepCopy() *CStorPoolSpec {
	if in == nil {
		return nil
	}
	out := new(CStorPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplica) DeepCopyInto(out *CStorVolumeReplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplica.
func (in *CStorVolumeReplica) DeepCopy() *CStorVolumeReplica {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeReplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaList) DeepCopyInto(out *CStorVolumeReplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CStorVolumeReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaList.
func (in *CStorVolumeReplicaList) DeepCopy() *CStorVolumeReplicaList {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CStorVolumeReplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolumeReplicaSpec) DeepCopyInto(out *CStorVolumeReplicaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolumeReplicaSpec.
func (in *CStorVolumeReplicaSpec) DeepCopy() *CStorVolumeReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(CStorVolumeReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskAttr) DeepCopyInto(out *DiskAttr) {
	*out = *in
	if in.DiskList != nil {
		in, out := &in.DiskList, &out.DiskList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskAttr.
func (in *DiskAttr) DeepCopy() *DiskAttr {
	if in == nil {
		return nil
	}
	out := new(DiskAttr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Policy) DeepCopyInto(out *Policy) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Policy.
func (in *Policy) DeepCopy() *Policy {
	if in == nil {
		return nil
	}
	out := new(Policy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunTasks) DeepCopyInto(out *RunTasks) {
	*out = *in
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]Task, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunTasks.
func (in *RunTasks) DeepCopy() *RunTasks {
	if in == nil {
		return nil
	}
	out := new(RunTasks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePool) DeepCopyInto(out *StoragePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePool.
func (in *StoragePool) DeepCopy() *StoragePool {
	if in == nil {
		return nil
	}
	out := new(StoragePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaim) DeepCopyInto(out *StoragePoolClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaim.
func (in *StoragePoolClaim) DeepCopy() *StoragePoolClaim {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePoolClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaimList) DeepCopyInto(out *StoragePoolClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePoolClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaimList.
func (in *StoragePoolClaimList) DeepCopy() *StoragePoolClaimList {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePoolClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaimSpec) DeepCopyInto(out *StoragePoolClaimSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaimSpec.
func (in *StoragePoolClaimSpec) DeepCopy() *StoragePoolClaimSpec {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolList) DeepCopyInto(out *StoragePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoragePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolList.
func (in *StoragePoolList) DeepCopy() *StoragePoolList {
	if in == nil {
		return nil
	}
	out := new(StoragePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoragePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolSpec) DeepCopyInto(out *StoragePoolSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolSpec.
func (in *StoragePoolSpec) DeepCopy() *StoragePoolSpec {
	if in == nil {
		return nil
	}
	out := new(StoragePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Task) DeepCopyInto(out *Task) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Task.
func (in *Task) DeepCopy() *Task {
	if in == nil {
		return nil
	}
	out := new(Task)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeParameterGroup) DeepCopyInto(out *VolumeParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeParameterGroup.
func (in *VolumeParameterGroup) DeepCopy() *VolumeParameterGroup {
	if in == nil {
		return nil
	}
	out := new(VolumeParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeParameterGroupList) DeepCopyInto(out *VolumeParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VolumeParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeParameterGroupList.
func (in *VolumeParameterGroupList) DeepCopy() *VolumeParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(VolumeParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VolumeParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeParameterGroupSpec) DeepCopyInto(out *VolumeParameterGroupSpec) {
	*out = *in
	if in.Policies != nil {
		in, out := &in.Policies, &out.Policies
		*out = make([]Policy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.RunTasks.DeepCopyInto(&out.RunTasks)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeParameterGroupSpec.
func (in *VolumeParameterGroupSpec) DeepCopy() *VolumeParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(VolumeParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}