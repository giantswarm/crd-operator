package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AdmissionWebhookTemplateSpec defines the desired state of Mutating.
//
// Notes:
// - It always uses Create and Update for operations.
type AdmissionWebhookTemplateSpec struct {
	// Handler specifies what handles webhook requests.
	// +optional
	Handler *WebhookHandlerConfig `json:"handler,omitempty"`

	// ObjectSelector decides whether to run the webhook based on if the
	// object has matching labels. objectSelector is evaluated against both
	// the oldObject and newObject that would be sent to the webhook, and
	// is considered to match if either object matches the selector. A null
	// object (oldObject in the case of create, or newObject in the case of
	// delete) or an object that cannot have labels (like a
	// DeploymentRollback or a PodProxyOptions object) is not considered to
	// match.
	// Use the object selector only if the webhook is opt-in, because end
	// users may skip the admission webhook by setting the labels.
	// Default to the empty LabelSelector, which matches everything.
	// +optional
	ObjectSelector *metav1.LabelSelector `json:"objectSelector,omitempty"`

	// IgnoreErrors is a flag specifying if the webhook errors should be ignored or not.
	// Errors are not ignored by default.
	// +optional
	IgnoreErrors *bool `json:"ignoreErrors,omitempty"`
}
