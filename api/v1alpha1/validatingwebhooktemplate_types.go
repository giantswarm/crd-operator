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

// ValidatingWebhookTemplateSpec defines the desired state of ValidatingWebhookTemplate
type ValidatingWebhookTemplateSpec struct {
	WebhookTemplateSpec `json:",inline"`
}

// ValidatingWebhookTemplateStatus defines the observed state of ValidatingWebhookTemplate
type ValidatingWebhookTemplateStatus struct {
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ValidatingWebhookTemplate is the Schema for the validatingwebhooktemplates API
type ValidatingWebhookTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ValidatingWebhookTemplateSpec   `json:"spec,omitempty"`
	Status ValidatingWebhookTemplateStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ValidatingWebhookTemplateList contains a list of ValidatingWebhookTemplate
type ValidatingWebhookTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ValidatingWebhookTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ValidatingWebhookTemplate{}, &ValidatingWebhookTemplateList{})
}
