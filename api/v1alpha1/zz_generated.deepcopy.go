// +build !ignore_autogenerated

/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CcChannel) DeepCopyInto(out *CcChannel) {
	*out = *in
	if in.Orgs != nil {
		in, out := &in.Orgs, &out.Orgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CcChannel.
func (in *CcChannel) DeepCopy() *CcChannel {
	if in == nil {
		return nil
	}
	out := new(CcChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chaincode) DeepCopyInto(out *Chaincode) {
	*out = *in
	if in.Orgs != nil {
		in, out := &in.Orgs, &out.Orgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CcChannel != nil {
		in, out := &in.CcChannel, &out.CcChannel
		*out = make([]CcChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chaincode.
func (in *Chaincode) DeepCopy() *Chaincode {
	if in == nil {
		return nil
	}
	out := new(Chaincode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channel) DeepCopyInto(out *Channel) {
	*out = *in
	if in.Orgs != nil {
		in, out := &in.Orgs, &out.Orgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channel.
func (in *Channel) DeepCopy() *Channel {
	if in == nil {
		return nil
	}
	out := new(Channel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configtx) DeepCopyInto(out *Configtx) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configtx.
func (in *Configtx) DeepCopy() *Configtx {
	if in == nil {
		return nil
	}
	out := new(Configtx)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricNetwork) DeepCopyInto(out *FabricNetwork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricNetwork.
func (in *FabricNetwork) DeepCopy() *FabricNetwork {
	if in == nil {
		return nil
	}
	out := new(FabricNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricNetwork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricNetworkList) DeepCopyInto(out *FabricNetworkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FabricNetwork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricNetworkList.
func (in *FabricNetworkList) DeepCopy() *FabricNetworkList {
	if in == nil {
		return nil
	}
	out := new(FabricNetworkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FabricNetworkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricNetworkSpec) DeepCopyInto(out *FabricNetworkSpec) {
	*out = *in
	out.Configtx = in.Configtx
	out.Genesis = in.Genesis
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]v1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Topology.DeepCopyInto(&out.Topology)
	in.Network.DeepCopyInto(&out.Network)
	in.HlfKube.DeepCopyInto(&out.HlfKube)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricNetworkSpec.
func (in *FabricNetworkSpec) DeepCopy() *FabricNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(FabricNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FabricNetworkStatus) DeepCopyInto(out *FabricNetworkStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FabricNetworkStatus.
func (in *FabricNetworkStatus) DeepCopy() *FabricNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(FabricNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Genesis) DeepCopyInto(out *Genesis) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Genesis.
func (in *Genesis) DeepCopy() *Genesis {
	if in == nil {
		return nil
	}
	out := new(Genesis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = make([]Channel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Chaincode != nil {
		in, out := &in.Chaincode, &out.Chaincode
		*out = make([]Chaincode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
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
func (in *OrdererOrg) DeepCopyInto(out *OrdererOrg) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrdererOrg.
func (in *OrdererOrg) DeepCopy() *OrdererOrg {
	if in == nil {
		return nil
	}
	out := new(OrdererOrg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeerOrg) DeepCopyInto(out *PeerOrg) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeerOrg.
func (in *PeerOrg) DeepCopy() *PeerOrg {
	if in == nil {
		return nil
	}
	out := new(PeerOrg)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topology) DeepCopyInto(out *Topology) {
	*out = *in
	if in.OrdererOrg != nil {
		in, out := &in.OrdererOrg, &out.OrdererOrg
		*out = make([]OrdererOrg, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PeerOrg != nil {
		in, out := &in.PeerOrg, &out.PeerOrg
		*out = make([]PeerOrg, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topology.
func (in *Topology) DeepCopy() *Topology {
	if in == nil {
		return nil
	}
	out := new(Topology)
	in.DeepCopyInto(out)
	return out
}
