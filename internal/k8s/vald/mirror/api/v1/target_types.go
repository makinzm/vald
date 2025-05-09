// Copyright (C) 2019-2025 vdaas.org vald team <vald@vdaas.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/scheme"
)

var (
	// GroupVersion is group version used to register these objects.
	GroupVersion = schema.GroupVersion{Group: "vald.vdaas.org", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme.
	SchemeBuilder = &scheme.Builder{GroupVersion: GroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func init() {
	SchemeBuilder.Register(
		&ValdMirrorTarget{},
		&ValdMirrorTargetList{},
	)
}

// ValdMirrorTarget is a mirror information.
type ValdMirrorTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   MirrorTargetSpec   `json:"spec,omitempty"`
	Status MirrorTargetStatus `json:"status,omitempty"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdMirrorTarget) DeepCopyInto(out *ValdMirrorTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValdMirror.
func (in *ValdMirrorTarget) DeepCopy() *ValdMirrorTarget {
	if in == nil {
		return nil
	}
	out := new(ValdMirrorTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdMirrorTarget) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}

// MirrorTargetSpec is a description of a ValdMirrorTarget.
type MirrorTargetSpec struct {
	Colocation string       `json:"colocation,omitempty"`
	Target     MirrorTarget `json:"target,omitempty"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorTargetSpec) DeepCopyInto(out *MirrorTargetSpec) {
	*out = *in
	out.Colocation = in.Colocation
	out.Target = in.Target
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorSpec.
func (in *MirrorTargetSpec) DeepCopy() *MirrorTargetSpec {
	if in == nil {
		return nil
	}
	out := new(MirrorTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// MirrorTarget is a target information.
type MirrorTarget struct {
	Host string `json:"host,omitempty"`
	Port int    `json:"port,omitempty"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirrorTarget) DeepCopyInto(out *MirrorTarget) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirrorTarget.
func (in *MirrorTarget) DeepCopy() *MirrorTarget {
	if in == nil {
		return nil
	}
	out := new(MirrorTarget)
	in.DeepCopyInto(out)
	return out
}

// MirrorTargetStatus is a status of ValdMirrorTarget.
type MirrorTargetStatus struct {
	Phase              MirrorTargetPhase `json:"phase,omitempty"`
	LastTransitionTime metav1.Time       `json:"lastTransitionTime,omitempty"`
}

// MirrorTargetPhase is a label for the condition of a ValdMirrorTarget at the current time.
type MirrorTargetPhase string

const (
	// MirrorTargetConnected means that the ValdMirrorTarget has been accepted by the system.
	MirrorTargetPending = MirrorTargetPhase("Pending")

	// MirrorTargetConnected means that the target was connected.
	MirrorTargetConnected = MirrorTargetPhase("Connected")

	// MirrorTargetDisconnected means that the target was disconnected.
	MirrorTargetDisconnected = MirrorTargetPhase("Disconnected")

	// MirrorTargetUnknown means that for some reason the state of the ValdMirrorTarget could not be obtained.
	MirrorTargetUnknown = MirrorTargetPhase("Unknown")
)

// ValdMirrorTargetList is the whole list of all ValdMirror which have been registered with master.
type ValdMirrorTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []ValdMirrorTarget `json:"items,omitempty"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValdMirrorTargetList) DeepCopyInto(out *ValdMirrorTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if len(in.Items) != 0 {
		out.Items = make([]ValdMirrorTarget, len(in.Items))
		for i := 0; i < len(in.Items); i++ {
			out.Items[i] = *in.Items[i].DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValdMirrorList.
func (in *ValdMirrorTargetList) DeepCopy() *ValdMirrorTargetList {
	if in == nil {
		return nil
	}
	out := new(ValdMirrorTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ValdMirrorTargetList) DeepCopyObject() runtime.Object {
	return in.DeepCopy()
}
