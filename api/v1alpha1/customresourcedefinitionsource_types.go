package v1alpha1

const (
	// File is a CRD source type where CRDs are stored in a YAML file that can has a URL.
	// An example use case for this CustomResourceDefinitionSourceType is a GitHub release asset
	// that contains CRDs.
	File CustomResourceDefinitionSourceType = "File"

	// GitRepoDirectory is a CRD source type where CRDs are stored in a git repository.
	//
	// Format of the URL is https://$gitHost/$user/$repo/tree/branch/path/to/api/v1something1,
	// for example https://github.com/kubernetes-sigs/cluster-api/tree/main/controlplane/kubeadm/api/v1beta1
	GitRepoDirectory CustomResourceDefinitionSourceType = "GitRepoDirectory"
)

// CustomResourceDefinitionSourceType is a type for defining different types of CRD sources,
// i.e. different places from which the CRDs can be obtained.
type CustomResourceDefinitionSourceType string

// CustomResourceDefinitionSource defines the desired state of CustomResourceDefinitionSource
type CustomResourceDefinitionSource struct {
	// Type is a CustomResourceDefinitionSource (possible values are File and GitRepoDirectory)
	// +kubebuilder:default=File
	Type CustomResourceDefinitionSourceType `json:"type"`

	// URL pointing to the CRD source
	URL string `json:"url"`
}
