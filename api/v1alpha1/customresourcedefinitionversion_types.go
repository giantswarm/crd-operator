package v1alpha1

// CustomResourceDefinitionVersionsConfig defines stored and served API versions for the CRD.
type CustomResourceDefinitionVersionsConfig struct {
	// Storage is the current storage version. This version is also served and not deprecated.
	Storage CustomResourceDefinitionStorageVersion `json:"storage"`

	// Served defines other served versions, besides the storage one.
	Served []CustomResourceDefinitionServedVersion `json:"served,omitempty"`
}

// CustomResourceDefinitionStorageVersion represents the current storage version
// that is served and not deprecated.
type CustomResourceDefinitionStorageVersion string

// CustomResourceDefinitionServedVersion defines other served versions, besides the storage one.
type CustomResourceDefinitionServedVersion struct {
	// Name of the API version, e.g. v1alpha1, v1beta1, v1, v2
	Name string `json:"name"`

	// Deprecated is the flag that defines if the CRD version is deprecated
	// +optional
	Deprecated *bool `json:"deprecated,omitempty"`
}
