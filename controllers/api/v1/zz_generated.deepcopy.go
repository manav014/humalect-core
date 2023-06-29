//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcrCredentials) DeepCopyInto(out *AcrCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcrCredentials.
func (in *AcrCredentials) DeepCopy() *AcrCredentials {
	if in == nil {
		return nil
	}
	out := new(AcrCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Application) DeepCopyInto(out *Application) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Application.
func (in *Application) DeepCopy() *Application {
	if in == nil {
		return nil
	}
	out := new(Application)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Application) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationList) DeepCopyInto(out *ApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Application, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationList.
func (in *ApplicationList) DeepCopy() *ApplicationList {
	if in == nil {
		return nil
	}
	out := new(ApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	in.DeploymentYamlManifest.DeepCopyInto(&out.DeploymentYamlManifest)
	in.ServiceYamlManifest.DeepCopyInto(&out.ServiceYamlManifest)
	in.IngressYamlManifest.DeepCopyInto(&out.IngressYamlManifest)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationStatus) DeepCopyInto(out *ApplicationStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationStatus.
func (in *ApplicationStatus) DeepCopy() *ApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(ApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsSecretCredentials) DeepCopyInto(out *AwsSecretCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsSecretCredentials.
func (in *AwsSecretCredentials) DeepCopy() *AwsSecretCredentials {
	if in == nil {
		return nil
	}
	out := new(AwsSecretCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureSecretCredentials) DeepCopyInto(out *AzureSecretCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureSecretCredentials.
func (in *AzureSecretCredentials) DeepCopy() *AzureSecretCredentials {
	if in == nil {
		return nil
	}
	out := new(AzureSecretCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSet) DeepCopyInto(out *DeploymentSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSet.
func (in *DeploymentSet) DeepCopy() *DeploymentSet {
	if in == nil {
		return nil
	}
	out := new(DeploymentSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSetList) DeepCopyInto(out *DeploymentSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DeploymentSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSetList.
func (in *DeploymentSetList) DeepCopy() *DeploymentSetList {
	if in == nil {
		return nil
	}
	out := new(DeploymentSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeploymentSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSetSpec) DeepCopyInto(out *DeploymentSetSpec) {
	*out = *in
	out.EcrCredentials = in.EcrCredentials
	out.DockerHubCredentials = in.DockerHubCredentials
	out.AcrCredentials = in.AcrCredentials
	out.AwsSecretCredentials = in.AwsSecretCredentials
	out.AzureSecretCredentials = in.AzureSecretCredentials
	in.DeploymentYamlManifest.DeepCopyInto(&out.DeploymentYamlManifest)
	in.ServiceYamlManifest.DeepCopyInto(&out.ServiceYamlManifest)
	in.IngressYamlManifest.DeepCopyInto(&out.IngressYamlManifest)
	if in.DockerManifest != nil {
		in, out := &in.DockerManifest, &out.DockerManifest
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSetSpec.
func (in *DeploymentSetSpec) DeepCopy() *DeploymentSetSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSetStatus) DeepCopyInto(out *DeploymentSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSetStatus.
func (in *DeploymentSetStatus) DeepCopy() *DeploymentSetStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentYamlManifestType) DeepCopyInto(out *DeploymentYamlManifestType) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentYamlManifestType.
func (in *DeploymentYamlManifestType) DeepCopy() *DeploymentYamlManifestType {
	if in == nil {
		return nil
	}
	out := new(DeploymentYamlManifestType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerHubCredentials) DeepCopyInto(out *DockerHubCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerHubCredentials.
func (in *DockerHubCredentials) DeepCopy() *DockerHubCredentials {
	if in == nil {
		return nil
	}
	out := new(DockerHubCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EcrCredentials) DeepCopyInto(out *EcrCredentials) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EcrCredentials.
func (in *EcrCredentials) DeepCopy() *EcrCredentials {
	if in == nil {
		return nil
	}
	out := new(EcrCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressYamlManifestType) DeepCopyInto(out *IngressYamlManifestType) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressYamlManifestType.
func (in *IngressYamlManifestType) DeepCopy() *IngressYamlManifestType {
	if in == nil {
		return nil
	}
	out := new(IngressYamlManifestType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceYamlManifestType) DeepCopyInto(out *ServiceYamlManifestType) {
	*out = *in
	in.Metadata.DeepCopyInto(&out.Metadata)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceYamlManifestType.
func (in *ServiceYamlManifestType) DeepCopy() *ServiceYamlManifestType {
	if in == nil {
		return nil
	}
	out := new(ServiceYamlManifestType)
	in.DeepCopyInto(out)
	return out
}