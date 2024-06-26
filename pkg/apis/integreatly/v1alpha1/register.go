// NOTE: Boilerplate only.  Ignore this file.

// Package v1alpha1 contains API Schema definitions for the integreatly v1alpha1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=integreatly.org
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	// ORIGINAL "sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
	// For match new version "sigs.k8s.io/controller-runtime/pkg/scheme"
	//"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
	"sigs.k8s.io/controller-runtime/pkg/scheme"

)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "integreatly.org", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)
