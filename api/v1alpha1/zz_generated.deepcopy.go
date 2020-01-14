// +build !ignore_autogenerated

/*

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
func (in *Hashring) DeepCopyInto(out *Hashring) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tenants != nil {
		in, out := &in.Tenants, &out.Tenants
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hashring.
func (in *Hashring) DeepCopy() *Hashring {
	if in == nil {
		return nil
	}
	out := new(Hashring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observatorium) DeepCopyInto(out *Observatorium) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observatorium.
func (in *Observatorium) DeepCopy() *Observatorium {
	if in == nil {
		return nil
	}
	out := new(Observatorium)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Observatorium) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumList) DeepCopyInto(out *ObservatoriumList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Observatorium, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumList.
func (in *ObservatoriumList) DeepCopy() *ObservatoriumList {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObservatoriumList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumSpec) DeepCopyInto(out *ObservatoriumSpec) {
	*out = *in
	in.Thanos.DeepCopyInto(&out.Thanos)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumSpec.
func (in *ObservatoriumSpec) DeepCopy() *ObservatoriumSpec {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObservatoriumStatus) DeepCopyInto(out *ObservatoriumStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObservatoriumStatus.
func (in *ObservatoriumStatus) DeepCopy() *ObservatoriumStatus {
	if in == nil {
		return nil
	}
	out := new(ObservatoriumStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReceiveController) DeepCopyInto(out *ReceiveController) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.Hashrings != nil {
		in, out := &in.Hashrings, &out.Hashrings
		*out = make([]*Hashring, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Hashring)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReceiveController.
func (in *ReceiveController) DeepCopy() *ReceiveController {
	if in == nil {
		return nil
	}
	out := new(ReceiveController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosComponentSpec) DeepCopyInto(out *ThanosComponentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosComponentSpec.
func (in *ThanosComponentSpec) DeepCopy() *ThanosComponentSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosComponentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosPersistentSpec) DeepCopyInto(out *ThanosPersistentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	if in.PVCSize != nil {
		in, out := &in.PVCSize, &out.PVCSize
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosPersistentSpec.
func (in *ThanosPersistentSpec) DeepCopy() *ThanosPersistentSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosPersistentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosSpec) DeepCopyInto(out *ThanosSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	in.ReceiveControllerSpec.DeepCopyInto(&out.ReceiveControllerSpec)
	in.Receive.DeepCopyInto(&out.Receive)
	in.Querier.DeepCopyInto(&out.Querier)
	in.Store.DeepCopyInto(&out.Store)
	in.Compactor.DeepCopyInto(&out.Compactor)
	if in.ObjectStoreConfigSecret != nil {
		in, out := &in.ObjectStoreConfigSecret, &out.ObjectStoreConfigSecret
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosSpec.
func (in *ThanosSpec) DeepCopy() *ThanosSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosSpec)
	in.DeepCopyInto(out)
	return out
}
