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
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MutatingWebhookTemplateSpec defines the desired state of MutatingWebhookTemplate
type MutatingWebhookTemplateSpec struct {
	WebhookTemplateSpec `json:",inline"`
	ReinvocationPolicy  *admissionregistrationv1.ReinvocationPolicyType `json:"reinvocationPolicy,omitempty"`
}

// MutatingWebhookTemplateStatus defines the observed state of MutatingWebhookTemplate
type MutatingWebhookTemplateStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// MutatingWebhookTemplate is the Schema for the mutatingwebhooktemplates API
type MutatingWebhookTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MutatingWebhookTemplateSpec   `json:"spec,omitempty"`
	Status MutatingWebhookTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// MutatingWebhookTemplateList contains a list of MutatingWebhookTemplate
type MutatingWebhookTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MutatingWebhookTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MutatingWebhookTemplate{}, &MutatingWebhookTemplateList{})
}
