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

// ConversionWebhookSpec defines the desired state of ConversionWebhook
type ConversionWebhookSpec struct {
	// CustomResourceDefinition for which this conversion webhook is used
	CustomResourceDefinition string `json:"customResourceDefinition"`

	// +kubebuilder:default=true
	Enabled bool `json:"enabled"`

	// Service that handles webhook requests.
	Service *ServiceReference `json:"service,omitempty"`

	URL *string `json:"URL,omitempty"`
}

// ConversionWebhookStatus defines the observed state of ConversionWebhook
type ConversionWebhookStatus struct {
	// Conditions represent the current state of the conversion webhook.
	Conditions Conditions `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ConversionWebhook is the Schema for the conversionwebhooks API
type ConversionWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConversionWebhookSpec   `json:"spec,omitempty"`
	Status ConversionWebhookStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ConversionWebhookList contains a list of ConversionWebhook
type ConversionWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConversionWebhook `json:"items"`
}

type ConversionWebhookTemplateSpec struct {
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConversionWebhookSpec `json:"spec,omitempty"`
}

func init() {
	SchemeBuilder.Register(&ConversionWebhook{}, &ConversionWebhookList{})
}
