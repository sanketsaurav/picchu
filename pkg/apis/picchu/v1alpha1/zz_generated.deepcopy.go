// +build !ignore_autogenerated

// Copyright © 2019 A Medium Corporation.
// Licensed under the Apache License, Version 2.0; see the NOTICE file.

// Code generated by deepcopy. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAWSInfo) DeepCopyInto(out *ClusterAWSInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAWSInfo.
func (in *ClusterAWSInfo) DeepCopy() *ClusterAWSInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterAWSInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConditionStatus) DeepCopyInto(out *ClusterConditionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConditionStatus.
func (in *ClusterConditionStatus) DeepCopy() *ClusterConditionStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterConditionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterDNSGroup) DeepCopyInto(out *ClusterDNSGroup) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterDNSGroup.
func (in *ClusterDNSGroup) DeepCopy() *ClusterDNSGroup {
	if in == nil {
		return nil
	}
	out := new(ClusterDNSGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIngresses) DeepCopyInto(out *ClusterIngresses) {
	*out = *in
	out.Public = in.Public
	out.Private = in.Private
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIngresses.
func (in *ClusterIngresses) DeepCopy() *ClusterIngresses {
	if in == nil {
		return nil
	}
	out := new(ClusterIngresses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterKubernetesStatus) DeepCopyInto(out *ClusterKubernetesStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterKubernetesStatus.
func (in *ClusterKubernetesStatus) DeepCopy() *ClusterKubernetesStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterKubernetesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(ClusterConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(ClusterAWSInfo)
		**out = **in
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]ClusterDNSGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.Ingresses = in.Ingresses
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	out.Kubernetes = in.Kubernetes
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(ClusterAWSInfo)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ClusterConditionStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Incarnation) DeepCopyInto(out *Incarnation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Incarnation.
func (in *Incarnation) DeepCopy() *Incarnation {
	if in == nil {
		return nil
	}
	out := new(Incarnation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Incarnation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationApp) DeepCopyInto(out *IncarnationApp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationApp.
func (in *IncarnationApp) DeepCopy() *IncarnationApp {
	if in == nil {
		return nil
	}
	out := new(IncarnationApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationAssignment) DeepCopyInto(out *IncarnationAssignment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationAssignment.
func (in *IncarnationAssignment) DeepCopy() *IncarnationAssignment {
	if in == nil {
		return nil
	}
	out := new(IncarnationAssignment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationHealthMetricStatus) DeepCopyInto(out *IncarnationHealthMetricStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationHealthMetricStatus.
func (in *IncarnationHealthMetricStatus) DeepCopy() *IncarnationHealthMetricStatus {
	if in == nil {
		return nil
	}
	out := new(IncarnationHealthMetricStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationHealthStatus) DeepCopyInto(out *IncarnationHealthStatus) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]IncarnationHealthMetricStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationHealthStatus.
func (in *IncarnationHealthStatus) DeepCopy() *IncarnationHealthStatus {
	if in == nil {
		return nil
	}
	out := new(IncarnationHealthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationList) DeepCopyInto(out *IncarnationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Incarnation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationList.
func (in *IncarnationList) DeepCopy() *IncarnationList {
	if in == nil {
		return nil
	}
	out := new(IncarnationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IncarnationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationResourceStatus) DeepCopyInto(out *IncarnationResourceStatus) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(types.NamespacedName)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationResourceStatus.
func (in *IncarnationResourceStatus) DeepCopy() *IncarnationResourceStatus {
	if in == nil {
		return nil
	}
	out := new(IncarnationResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationScaleStatus) DeepCopyInto(out *IncarnationScaleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationScaleStatus.
func (in *IncarnationScaleStatus) DeepCopy() *IncarnationScaleStatus {
	if in == nil {
		return nil
	}
	out := new(IncarnationScaleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationSpec) DeepCopyInto(out *IncarnationSpec) {
	*out = *in
	out.App = in.App
	out.Assignment = in.Assignment
	in.Scale.DeepCopyInto(&out.Scale)
	out.Release = in.Release
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigSelector != nil {
		in, out := &in.ConfigSelector, &out.ConfigSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationSpec.
func (in *IncarnationSpec) DeepCopy() *IncarnationSpec {
	if in == nil {
		return nil
	}
	out := new(IncarnationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IncarnationStatus) DeepCopyInto(out *IncarnationStatus) {
	*out = *in
	in.Health.DeepCopyInto(&out.Health)
	out.Scale = in.Scale
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]IncarnationResourceStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IncarnationStatus.
func (in *IncarnationStatus) DeepCopy() *IncarnationStatus {
	if in == nil {
		return nil
	}
	out := new(IncarnationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressInfo) DeepCopyInto(out *IngressInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressInfo.
func (in *IngressInfo) DeepCopy() *IngressInfo {
	if in == nil {
		return nil
	}
	out := new(IngressInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortInfo) DeepCopyInto(out *PortInfo) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortInfo.
func (in *PortInfo) DeepCopy() *PortInfo {
	if in == nil {
		return nil
	}
	out := new(PortInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateInfo) DeepCopyInto(out *RateInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateInfo.
func (in *RateInfo) DeepCopy() *RateInfo {
	if in == nil {
		return nil
	}
	out := new(RateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseInfo) DeepCopyInto(out *ReleaseInfo) {
	*out = *in
	out.Rate = in.Rate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseInfo.
func (in *ReleaseInfo) DeepCopy() *ReleaseInfo {
	if in == nil {
		return nil
	}
	out := new(ReleaseInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManager) DeepCopyInto(out *ReleaseManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManager.
func (in *ReleaseManager) DeepCopy() *ReleaseManager {
	if in == nil {
		return nil
	}
	out := new(ReleaseManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerList) DeepCopyInto(out *ReleaseManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReleaseManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerList.
func (in *ReleaseManagerList) DeepCopy() *ReleaseManagerList {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerReleaseStatus) DeepCopyInto(out *ReleaseManagerReleaseStatus) {
	*out = *in
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerReleaseStatus.
func (in *ReleaseManagerReleaseStatus) DeepCopy() *ReleaseManagerReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerSpec) DeepCopyInto(out *ReleaseManagerSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerSpec.
func (in *ReleaseManagerSpec) DeepCopy() *ReleaseManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseManagerStatus) DeepCopyInto(out *ReleaseManagerStatus) {
	*out = *in
	if in.Releases != nil {
		in, out := &in.Releases, &out.Releases
		*out = make([]ReleaseManagerReleaseStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseManagerStatus.
func (in *ReleaseManagerStatus) DeepCopy() *ReleaseManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceInfo) DeepCopyInto(out *ResourceInfo) {
	*out = *in
	out.CPU = in.CPU.DeepCopy()
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceInfo.
func (in *ResourceInfo) DeepCopy() *ResourceInfo {
	if in == nil {
		return nil
	}
	out := new(ResourceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Revision) DeepCopyInto(out *Revision) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Revision.
func (in *Revision) DeepCopy() *Revision {
	if in == nil {
		return nil
	}
	out := new(Revision)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Revision) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionApp) DeepCopyInto(out *RevisionApp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionApp.
func (in *RevisionApp) DeepCopy() *RevisionApp {
	if in == nil {
		return nil
	}
	out := new(RevisionApp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionList) DeepCopyInto(out *RevisionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Revision, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionList.
func (in *RevisionList) DeepCopy() *RevisionList {
	if in == nil {
		return nil
	}
	out := new(RevisionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RevisionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionRelease) DeepCopyInto(out *RevisionRelease) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionRelease.
func (in *RevisionRelease) DeepCopy() *RevisionRelease {
	if in == nil {
		return nil
	}
	out := new(RevisionRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionSpec) DeepCopyInto(out *RevisionSpec) {
	*out = *in
	out.App = in.App
	out.Release = in.Release
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]RevisionTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionSpec.
func (in *RevisionSpec) DeepCopy() *RevisionSpec {
	if in == nil {
		return nil
	}
	out := new(RevisionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionStatus) DeepCopyInto(out *RevisionStatus) {
	*out = *in
	if in.Incarnations != nil {
		in, out := &in.Incarnations, &out.Incarnations
		*out = make([]RevisionTargetIncarnationStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionStatus.
func (in *RevisionStatus) DeepCopy() *RevisionStatus {
	if in == nil {
		return nil
	}
	out := new(RevisionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTarget) DeepCopyInto(out *RevisionTarget) {
	*out = *in
	in.Scale.DeepCopyInto(&out.Scale)
	out.Release = in.Release
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]RevisionTargetMetric, len(*in))
		copy(*out, *in)
	}
	if in.ConfigSelector != nil {
		in, out := &in.ConfigSelector, &out.ConfigSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTarget.
func (in *RevisionTarget) DeepCopy() *RevisionTarget {
	if in == nil {
		return nil
	}
	out := new(RevisionTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTargetIncarnationStatus) DeepCopyInto(out *RevisionTargetIncarnationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTargetIncarnationStatus.
func (in *RevisionTargetIncarnationStatus) DeepCopy() *RevisionTargetIncarnationStatus {
	if in == nil {
		return nil
	}
	out := new(RevisionTargetIncarnationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTargetMetric) DeepCopyInto(out *RevisionTargetMetric) {
	*out = *in
	out.Queries = in.Queries
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTargetMetric.
func (in *RevisionTargetMetric) DeepCopy() *RevisionTargetMetric {
	if in == nil {
		return nil
	}
	out := new(RevisionTargetMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RevisionTargetMetricQueries) DeepCopyInto(out *RevisionTargetMetricQueries) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RevisionTargetMetricQueries.
func (in *RevisionTargetMetricQueries) DeepCopy() *RevisionTargetMetricQueries {
	if in == nil {
		return nil
	}
	out := new(RevisionTargetMetricQueries)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleInfo) DeepCopyInto(out *ScaleInfo) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleInfo.
func (in *ScaleInfo) DeepCopy() *ScaleInfo {
	if in == nil {
		return nil
	}
	out := new(ScaleInfo)
	in.DeepCopyInto(out)
	return out
}
