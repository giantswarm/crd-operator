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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

const (
	// File is a CRD source type where CRDs are stored in a YAML file that can has a URL.
	// An example use case for this CustomResourceDefinitionSourceType is a GitHub release asset
	// that contains CRDs.
	File CustomResourceDefinitionSourceType = "File"

	// GitRepoDirectory is a CRD source type where CRDs are stored in a git repository.
	//
	// Format of the URL is https://$gitHost/$user/$repo/tree/branch/path/to/api/v1something1,
	// for example https://github.com/kubernetes-sigs/cluster-api/tree/main/controlplane/kubeadm/api/v1beta1
	GitRepoDirectory CustomResourceDefinitionSourceType = "GitRepoDirectory"
)

type CustomResourceDefinitionSourceType string

// CustomResourceDefinitionSource defines the desired state of CustomResourceDefinitionSource
type CustomResourceDefinitionSource struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Type is a CustomResourceDefinitionSource (possible values are File and GitRepoDirectory)
	// +kubebuilder:default=File
	Type CustomResourceDefinitionSourceType `json:"type"`

	// URL pointing to the CRD source
	URL string `json:"url"`
}

type CustomResourceDefinitionVersionsConfig struct {
	// Storage is the current storage version. This version is also served and not deprecated.
	Storage CustomResourceDefinitionStorageVersion `json:"storage"`

	// Served defines all stored versions, besides the storage one.
	Served []CustomResourceDefinitionServedVersion `json:"served,omitempty"`
}

// CustomResourceDefinitionStorageVersion represents the current storage version
// that is served and not deprecated.
type CustomResourceDefinitionStorageVersion string

// CustomResourceDefinitionServedVersion defines all stored versions, besides the storage one.
type CustomResourceDefinitionServedVersion struct {
	// Name of the API version, e.g. v1alpha1, v1beta1, v1, v2
	Name string `json:"name"`

	// Deprecated is the flag that defines if the CRD version is deprecated
	// +kubebuilder:default=false
	Deprecated bool `json:"deprecated"`
}

type CustomResourceDefinitionWebhooks struct {
	Conversion *ConversionWebhookTemplateSpec `json:"conversion,omitempty"`
}

// CustomResourceDefinitionDeploymentSpec defines the desired state of CustomResourceDefinitionDeployment
type CustomResourceDefinitionDeploymentSpec struct {
	// Group is the CRD API group
	Group string `json:"group"`

	// Versions defines stored and served API versions for this CRD
	Versions CustomResourceDefinitionVersionsConfig `json:"versions"`

	// Kind is the CRD kind name
	Kind string `json:"kind"`

	// Source where the CRDs can be found
	Source CustomResourceDefinitionSource `json:"source"`

	// Webhooks is the configuration for mutating, validating and conversion webhooks.
	Webhooks *CustomResourceDefinitionWebhooks `json:"webhooks,omitempty"`
}

// CustomResourceDefinitionDeploymentStatus defines the observed state of CustomResourceDefinitionDeployment
type CustomResourceDefinitionDeploymentStatus struct {
	// Conditions represent the current state of the CRD deployment
	Conditions Conditions `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomResourceDefinitionDeployment is the Schema for the customresourcedefinitiondeployments API
type CustomResourceDefinitionDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomResourceDefinitionDeploymentSpec   `json:"spec,omitempty"`
	Status CustomResourceDefinitionDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomResourceDefinitionDeploymentList contains a list of CustomResourceDefinitionDeployment
type CustomResourceDefinitionDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomResourceDefinitionDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomResourceDefinitionDeployment{}, &CustomResourceDefinitionDeploymentList{})
}
