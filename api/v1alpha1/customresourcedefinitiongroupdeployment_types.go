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

// CustomResourceDefinitionGroupDeploymentSpec defines the desired state of CustomResourceDefinitionGroupDeployment
type CustomResourceDefinitionGroupDeploymentSpec struct {
	// Group is the CRD API group
	Group string `json:"group"`

	// Versions defines stored and served API versions for this CRD
	Versions CustomResourceDefinitionVersionsConfig `json:"versions"`

	// Kinds are the CRD kind names
	Kinds []string `json:"kinds"`

	// Source where the CRDs can be found
	Source CustomResourceDefinitionSource `json:"source"`

	// WebhooksConfig specifies configuration for mutating, validating and conversion
	// webhooks, and certificate used by the webhook service.
	// +optional
	WebhooksConfig *WebhooksConfig `json:"webhooksConfig,omitempty"`
}

// CustomResourceDefinitionGroupDeploymentStatus defines the observed state of CustomResourceDefinitionGroupDeployment
type CustomResourceDefinitionGroupDeploymentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CustomResourceDefinitionGroupDeployment is the Schema for the customresourcedefinitiongroupdeployments API
type CustomResourceDefinitionGroupDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomResourceDefinitionGroupDeploymentSpec   `json:"spec,omitempty"`
	Status CustomResourceDefinitionGroupDeploymentStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CustomResourceDefinitionGroupDeploymentList contains a list of CustomResourceDefinitionGroupDeployment
type CustomResourceDefinitionGroupDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomResourceDefinitionGroupDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomResourceDefinitionGroupDeployment{}, &CustomResourceDefinitionGroupDeploymentList{})
}
