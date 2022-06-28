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

// ServiceCertificateTemplateSpec defines the desired state of ServiceCertificateTemplate
type ServiceCertificateTemplateSpec struct {
	// SecretName is a name of the secret where the certificate is stored.
	SecretName string `json:"secretName"`

	// IssuerRef is a reference to Issuer or ClusterIssuer resource.
	IssuerRef corev1.TypedLocalObjectReference `json:"issuerRef"`
}

//+kubebuilder:object:root=true

// ServiceCertificateTemplate is the Schema for the servicecertificatetemplates API
type ServiceCertificateTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ServiceCertificateTemplateSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// ServiceCertificateTemplateList contains a list of ServiceCertificateTemplate
type ServiceCertificateTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceCertificateTemplate `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ServiceCertificateTemplate{}, &ServiceCertificateTemplateList{})
}

type CertificateConfig struct {
	// +kubebuilder:default=true
	Enabled bool `json:"enabled"`

	TemplateRef corev1.TypedLocalObjectReference `json:"templateRef"`
}
