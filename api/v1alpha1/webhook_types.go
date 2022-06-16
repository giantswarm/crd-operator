package v1alpha1

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
