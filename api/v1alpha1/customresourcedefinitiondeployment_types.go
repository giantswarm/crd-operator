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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CustomResourceDefinitionDeploymentSpec defines the desired state of CustomResourceDefinitionDeployment
type CustomResourceDefinitionDeploymentSpec struct {
	// Group is the CRD API group
	Group string `json:"group"`

	// Versions defines stored and served API versions for this CRD
	Versions CustomResourceDefinitionVersionsConfig `json:"versions"`

	// Kinds is the list of CRD kind names
	Kinds []string `json:"kinds"`

	// Source where the CRDs can be found
	Source CustomResourceDefinitionSource `json:"source"`

	// +optional
	MutatingWebhook *CustomResourceDefinitionWebhookConfig `json:"mutatingWebhook,omitempty"`

	// +optional
	ValidatingWebhook *CustomResourceDefinitionWebhookConfig `json:"validatingWebhook,omitempty"`
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

type CustomResourceDefinitionWebhookConfig struct {
	// +kubebuilder:default=true
	Enabled bool `json:"enabled"`

	TemplateRef corev1.TypedLocalObjectReference `json:"templateRef"`
}

func init() {
	SchemeBuilder.Register(&CustomResourceDefinitionDeployment{}, &CustomResourceDefinitionDeploymentList{})
}
