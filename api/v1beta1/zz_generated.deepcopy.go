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

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalHostStatus) DeepCopyInto(out *BaremetalHostStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalHostStatus.
func (in *BaremetalHostStatus) DeepCopy() *BaremetalHostStatus {
	if in == nil {
		return nil
	}
	out := new(BaremetalHostStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalSet) DeepCopyInto(out *BaremetalSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalSet.
func (in *BaremetalSet) DeepCopy() *BaremetalSet {
	if in == nil {
		return nil
	}
	out := new(BaremetalSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BaremetalSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalSetList) DeepCopyInto(out *BaremetalSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BaremetalSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalSetList.
func (in *BaremetalSetList) DeepCopy() *BaremetalSetList {
	if in == nil {
		return nil
	}
	out := new(BaremetalSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BaremetalSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalSetSpec) DeepCopyInto(out *BaremetalSetSpec) {
	*out = *in
	if in.BmhLabelSelector != nil {
		in, out := &in.BmhLabelSelector, &out.BmhLabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.HardwareReqs = in.HardwareReqs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalSetSpec.
func (in *BaremetalSetSpec) DeepCopy() *BaremetalSetSpec {
	if in == nil {
		return nil
	}
	out := new(BaremetalSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaremetalSetStatus) DeepCopyInto(out *BaremetalSetStatus) {
	*out = *in
	if in.BaremetalHosts != nil {
		in, out := &in.BaremetalHosts, &out.BaremetalHosts
		*out = make(map[string]BaremetalHostStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaremetalSetStatus.
func (in *BaremetalSetStatus) DeepCopy() *BaremetalSetStatus {
	if in == nil {
		return nil
	}
	out := new(BaremetalSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUCountReq) DeepCopyInto(out *CPUCountReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUCountReq.
func (in *CPUCountReq) DeepCopy() *CPUCountReq {
	if in == nil {
		return nil
	}
	out := new(CPUCountReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUMhzReq) DeepCopyInto(out *CPUMhzReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUMhzReq.
func (in *CPUMhzReq) DeepCopy() *CPUMhzReq {
	if in == nil {
		return nil
	}
	out := new(CPUMhzReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUReqs) DeepCopyInto(out *CPUReqs) {
	*out = *in
	out.CountReq = in.CountReq
	out.MhzReq = in.MhzReq
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUReqs.
func (in *CPUReqs) DeepCopy() *CPUReqs {
	if in == nil {
		return nil
	}
	out := new(CPUReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlane) DeepCopyInto(out *ControlPlane) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlane.
func (in *ControlPlane) DeepCopy() *ControlPlane {
	if in == nil {
		return nil
	}
	out := new(ControlPlane)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlane) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneList) DeepCopyInto(out *ControlPlaneList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControlPlane, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneList.
func (in *ControlPlaneList) DeepCopy() *ControlPlaneList {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneSpec) DeepCopyInto(out *ControlPlaneSpec) {
	*out = *in
	in.Controller.DeepCopyInto(&out.Controller)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneSpec.
func (in *ControlPlaneSpec) DeepCopy() *ControlPlaneSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneStatus) DeepCopyInto(out *ControlPlaneStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneStatus.
func (in *ControlPlaneStatus) DeepCopy() *ControlPlaneStatus {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerSpec) DeepCopyInto(out *ControllerSpec) {
	*out = *in
	in.OSPNetwork.DeepCopyInto(&out.OSPNetwork)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerSpec.
func (in *ControllerSpec) DeepCopy() *ControllerSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerVM) DeepCopyInto(out *ControllerVM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerVM.
func (in *ControllerVM) DeepCopy() *ControllerVM {
	if in == nil {
		return nil
	}
	out := new(ControllerVM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerVM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerVMList) DeepCopyInto(out *ControllerVMList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControllerVM, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerVMList.
func (in *ControllerVMList) DeepCopy() *ControllerVMList {
	if in == nil {
		return nil
	}
	out := new(ControllerVMList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerVMList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerVMSpec) DeepCopyInto(out *ControllerVMSpec) {
	*out = *in
	in.OSPNetwork.DeepCopyInto(&out.OSPNetwork)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerVMSpec.
func (in *ControllerVMSpec) DeepCopy() *ControllerVMSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerVMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerVMStatus) DeepCopyInto(out *ControllerVMStatus) {
	*out = *in
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerVMStatus.
func (in *ControllerVMStatus) DeepCopy() *ControllerVMStatus {
	if in == nil {
		return nil
	}
	out := new(ControllerVMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskGbReq) DeepCopyInto(out *DiskGbReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskGbReq.
func (in *DiskGbReq) DeepCopy() *DiskGbReq {
	if in == nil {
		return nil
	}
	out := new(DiskGbReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskReqs) DeepCopyInto(out *DiskReqs) {
	*out = *in
	out.GbReq = in.GbReq
	out.SSDReq = in.SSDReq
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskReqs.
func (in *DiskReqs) DeepCopy() *DiskReqs {
	if in == nil {
		return nil
	}
	out := new(DiskReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSSDReq) DeepCopyInto(out *DiskSSDReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSSDReq.
func (in *DiskSSDReq) DeepCopy() *DiskSSDReq {
	if in == nil {
		return nil
	}
	out := new(DiskSSDReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HardwareReqs) DeepCopyInto(out *HardwareReqs) {
	*out = *in
	out.CPUReqs = in.CPUReqs
	out.MemReqs = in.MemReqs
	out.DiskReqs = in.DiskReqs
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HardwareReqs.
func (in *HardwareReqs) DeepCopy() *HardwareReqs {
	if in == nil {
		return nil
	}
	out := new(HardwareReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hash) DeepCopyInto(out *Hash) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hash.
func (in *Hash) DeepCopy() *Hash {
	if in == nil {
		return nil
	}
	out := new(Hash)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Host) DeepCopyInto(out *Host) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Host.
func (in *Host) DeepCopy() *Host {
	if in == nil {
		return nil
	}
	out := new(Host)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPReservation) DeepCopyInto(out *IPReservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPReservation.
func (in *IPReservation) DeepCopy() *IPReservation {
	if in == nil {
		return nil
	}
	out := new(IPReservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemGbReq) DeepCopyInto(out *MemGbReq) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemGbReq.
func (in *MemGbReq) DeepCopy() *MemGbReq {
	if in == nil {
		return nil
	}
	out := new(MemGbReq)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemReqs) DeepCopyInto(out *MemReqs) {
	*out = *in
	out.GbReq = in.GbReq
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemReqs.
func (in *MemReqs) DeepCopy() *MemReqs {
	if in == nil {
		return nil
	}
	out := new(MemReqs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	in.DesiredState.DeepCopyInto(&out.DesiredState)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClient) DeepCopyInto(out *OpenStackClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClient.
func (in *OpenStackClient) DeepCopy() *OpenStackClient {
	if in == nil {
		return nil
	}
	out := new(OpenStackClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientList) DeepCopyInto(out *OpenStackClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenStackClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientList.
func (in *OpenStackClientList) DeepCopy() *OpenStackClientList {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenStackClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientSpec) DeepCopyInto(out *OpenStackClientSpec) {
	*out = *in
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]v1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientSpec.
func (in *OpenStackClientSpec) DeepCopy() *OpenStackClientSpec {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenStackClientStatus) DeepCopyInto(out *OpenStackClientStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenStackClientStatus.
func (in *OpenStackClientStatus) DeepCopy() *OpenStackClientStatus {
	if in == nil {
		return nil
	}
	out := new(OpenStackClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudIPSet) DeepCopyInto(out *OvercloudIPSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudIPSet.
func (in *OvercloudIPSet) DeepCopy() *OvercloudIPSet {
	if in == nil {
		return nil
	}
	out := new(OvercloudIPSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvercloudIPSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudIPSetList) DeepCopyInto(out *OvercloudIPSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OvercloudIPSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudIPSetList.
func (in *OvercloudIPSetList) DeepCopy() *OvercloudIPSetList {
	if in == nil {
		return nil
	}
	out := new(OvercloudIPSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvercloudIPSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudIPSetSpec) DeepCopyInto(out *OvercloudIPSetSpec) {
	*out = *in
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudIPSetSpec.
func (in *OvercloudIPSetSpec) DeepCopy() *OvercloudIPSetSpec {
	if in == nil {
		return nil
	}
	out := new(OvercloudIPSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudIPSetStatus) DeepCopyInto(out *OvercloudIPSetStatus) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudIPSetStatus.
func (in *OvercloudIPSetStatus) DeepCopy() *OvercloudIPSetStatus {
	if in == nil {
		return nil
	}
	out := new(OvercloudIPSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudNet) DeepCopyInto(out *OvercloudNet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudNet.
func (in *OvercloudNet) DeepCopy() *OvercloudNet {
	if in == nil {
		return nil
	}
	out := new(OvercloudNet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvercloudNet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudNetList) DeepCopyInto(out *OvercloudNetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OvercloudNet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudNetList.
func (in *OvercloudNetList) DeepCopy() *OvercloudNetList {
	if in == nil {
		return nil
	}
	out := new(OvercloudNetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OvercloudNetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudNetSpec) DeepCopyInto(out *OvercloudNetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudNetSpec.
func (in *OvercloudNetSpec) DeepCopy() *OvercloudNetSpec {
	if in == nil {
		return nil
	}
	out := new(OvercloudNetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OvercloudNetStatus) DeepCopyInto(out *OvercloudNetStatus) {
	*out = *in
	if in.Reservations != nil {
		in, out := &in.Reservations, &out.Reservations
		*out = make([]IPReservation, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OvercloudNetStatus.
func (in *OvercloudNetStatus) DeepCopy() *OvercloudNetStatus {
	if in == nil {
		return nil
	}
	out := new(OvercloudNetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionServer) DeepCopyInto(out *ProvisionServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionServer.
func (in *ProvisionServer) DeepCopy() *ProvisionServer {
	if in == nil {
		return nil
	}
	out := new(ProvisionServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionServerList) DeepCopyInto(out *ProvisionServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProvisionServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionServerList.
func (in *ProvisionServerList) DeepCopy() *ProvisionServerList {
	if in == nil {
		return nil
	}
	out := new(ProvisionServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProvisionServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionServerSpec) DeepCopyInto(out *ProvisionServerSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionServerSpec.
func (in *ProvisionServerSpec) DeepCopy() *ProvisionServerSpec {
	if in == nil {
		return nil
	}
	out := new(ProvisionServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvisionServerStatus) DeepCopyInto(out *ProvisionServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvisionServerStatus.
func (in *ProvisionServerStatus) DeepCopy() *ProvisionServerStatus {
	if in == nil {
		return nil
	}
	out := new(ProvisionServerStatus)
	in.DeepCopyInto(out)
	return out
}
