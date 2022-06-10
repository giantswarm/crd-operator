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

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type CustomResourceDefinitionSourceType string

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

// CustomResourceDefinitionSourceSpec defines the desired state of CustomResourceDefinitionSource
type CustomResourceDefinitionSourceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Type is a CustomResourceDefinitionSource (possible values are File and GitRepoDirectory)
	// +kubebuilder:default=File
	Type CustomResourceDefinitionSourceType `json:"type"`

	// URL pointing to the CRD source
	URL string `json:"url"`

	// CRDNames are the names of CRDs to look for in this CRD source
	CRDNames []string `json:"crd_names"`
}

// CustomResourceDefinitionSourceStatus defines the observed state of CustomResourceDefinitionSource
type CustomResourceDefinitionSourceStatus struct {
	// Conditions define the current state of the CRD source
	// +optional
	Conditions Conditions `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomResourceDefinitionSource is the Schema for the customresourcedefinitionsources API
type CustomResourceDefinitionSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomResourceDefinitionSourceSpec   `json:"spec,omitempty"`
	Status CustomResourceDefinitionSourceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomResourceDefinitionSourceList contains a list of CustomResourceDefinitionSource
type CustomResourceDefinitionSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomResourceDefinitionSource `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomResourceDefinitionSource{}, &CustomResourceDefinitionSourceList{})
}
