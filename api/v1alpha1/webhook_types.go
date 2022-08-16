package v1alpha1

import (
	apps "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

type WebhookHandlerConfig struct {
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
	Service *ServiceReference `json:"service,omitempty"`
}

type WebhooksConfig struct {
	// +optional
	Mutating *WebhookConfig `json:"mutating,omitempty"`

	// +optional
	Validating *WebhookConfig `json:"validating,omitempty"`

	// +optional
	Conversion *WebhookConfig `json:"conversion,omitempty"`

	// Handler
	// +optional
	Handler *WebhookHandlerConfig `json:"handler,omitempty"`

	// Certificate config
	// +optional
	Certificate *CertificateConfig `json:"certificate,omitempty"`

	// Deployment specifies config for the webhook server deployment.
	Deployment *DeploymentConfig `json:"deployment,omitempty"`
}

type WebhookConfig struct {
	// +kubebuilder:default=true
	Enabled bool `json:"enabled"`

	// +optional
	TemplateRef *core.TypedLocalObjectReference `json:"templateRef,omitempty"`
}

type DeploymentConfig struct {
	Name string `json:"name"`

	// Mode specifies how the webhook Deployment is managed. Update for just updating
	// an existing Deployment, CreateOrUpdate for creating and owning a new webhook
	// deployment.
	Mode WebhookDeploymentMode `json:"mode"`

	// Template for creating a new Deployment. This field is set only if Mode is set
	// to CreateOrUpdate.
	// +optional
	Template *DeploymentTemplateSpec `json:"template,omitempty"`
}

type WebhookDeploymentMode string

const (
	Update         WebhookDeploymentMode = "Update"
	CreateOrUpdate WebhookDeploymentMode = "CreateOrUpdate"
)

type DeploymentTemplateSpec struct {
	meta.ObjectMeta `json:",inline"`
	Spec            apps.DeploymentSpec `json:"spec"`
}
