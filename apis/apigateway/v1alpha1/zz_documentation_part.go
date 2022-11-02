/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// DocumentationPartParameters defines the desired state of DocumentationPart
type DocumentationPartParameters struct {
	// Region is which region the DocumentationPart will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// [Required] The location of the targeted API entity of the to-be-created documentation
	// part.
	// +kubebuilder:validation:Required
	Location *DocumentationPartLocation `json:"location"`
	// [Required] The new documentation content map of the targeted API entity.
	// Enclosed key-value pairs are API-specific, but only OpenAPI-compliant key-value
	// pairs can be exported and, hence, published.
	// +kubebuilder:validation:Required
	Properties                        *string `json:"properties"`
	CustomDocumentationPartParameters `json:",inline"`
}

// DocumentationPartSpec defines the desired state of DocumentationPart
type DocumentationPartSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DocumentationPartParameters `json:"forProvider"`
}

// DocumentationPartObservation defines the observed state of DocumentationPart
type DocumentationPartObservation struct {
	// The DocumentationPart identifier, generated by API Gateway when the DocumentationPart
	// is created.
	ID *string `json:"id,omitempty"`
}

// DocumentationPartStatus defines the observed state of DocumentationPart.
type DocumentationPartStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DocumentationPartObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPart is the Schema for the DocumentationParts API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type DocumentationPart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocumentationPartSpec   `json:"spec"`
	Status            DocumentationPartStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DocumentationPartList contains a list of DocumentationParts
type DocumentationPartList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DocumentationPart `json:"items"`
}

// Repository type metadata.
var (
	DocumentationPartKind             = "DocumentationPart"
	DocumentationPartGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DocumentationPartKind}.String()
	DocumentationPartKindAPIVersion   = DocumentationPartKind + "." + GroupVersion.String()
	DocumentationPartGroupVersionKind = GroupVersion.WithKind(DocumentationPartKind)
)

func init() {
	SchemeBuilder.Register(&DocumentationPart{}, &DocumentationPartList{})
}
