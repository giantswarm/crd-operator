package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AdmissionWebhookHandler struct {
	// URL gives the location of the webhook, in standard URL form
	// (`scheme://host:port/path`). Exactly one of URL or Service
	// must be specified.
	//
	// The `host` should not refer to a service running in the cluster; use
	// the Service field instead. The host might be resolved via external
	// DNS in some apiservers (e.g., `kube-apiserver` cannot resolve
	// in-cluster DNS as that would be a layering violation). `host` may
	// also be an IP address.
	//
	// Please note that using `localhost` or `127.0.0.1` as a `host` is
	// risky unless you take great care to run this webhook on all hosts
	// which run an apiserver which might need to make calls to this
	// webhook. Such installs are likely to be non-portable, i.e., not easy
	// to turn up in a new cluster.
	//
	// The scheme must be "https"; the URL must begin with "https://".
	//
	// A path is optional, and if present may be any string permissible in
	// a URL. You may use the path to pass an arbitrary string to the
	// webhook, for example, a cluster identifier.
	//
	// Attempting to use a user or basic auth e.g. "user:password@" is not
	// allowed. Fragments ("#...") and query parameters ("?...") are not
	// allowed, either.
	//
	// +optional
	URL *string `json:"url,omitempty"`

	// Service is a reference to the service for this webhook. Either
	// Service or URL must be specified.
	//
	// If the webhook is running within the cluster, then you should use Service.
	//
	// +optional
	Service ServiceReference `json:"service,omitempty"`
}

type WebhookPathStyle string

const (
	KubebuilderWebhookPathStyle WebhookPathStyle = "KubebuilderWebhookPathStyle"
)

type ServiceReference struct {
	Namespace string `json:"namespace"`

	Name string `json:"name"`

	// Path used for the webhook. If Path is not set, then PathStyle must be set.
	// +optional
	Path *string `json:"path,omitempty"`

	// PathStyle is a style of service path used for the webhook, e.g. KubebuilderWebhookPathStyle.
	// If PathStyle is not set, then Path must be set.
	// +optional
	PathStyle *WebhookPathStyle `json:"pathStyle,omitempty"`

	// +optional
	Port *int `json:"port,omitempty"`
}

// AdmissionWebhookTemplateSpec defines the desired state of MutatingWebhook.
//
// Notes:
// - It always uses Create and Update for operations.
type AdmissionWebhookTemplateSpec struct {
	// Handler specifies what handles webhook requests.
	Handler AdmissionWebhookHandler `json:"handler"`

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
