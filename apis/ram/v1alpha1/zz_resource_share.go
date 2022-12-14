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

// ResourceShareParameters defines the desired state of ResourceShare
type ResourceShareParameters struct {
	// Region is which region the ResourceShare will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// Indicates whether principals outside your organization in Organizations can
	// be associated with a resource share.
	AllowExternalPrincipals *bool `json:"allowExternalPrincipals,omitempty"`
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string `json:"clientToken,omitempty"`
	// The name of the resource share.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// The Amazon Resource Names (ARNs) of the permissions to associate with the
	// resource share. If you do not specify an ARN for the permission, RAM automatically
	// attaches the default version of the permission for each resource type. Only
	// one permission can be associated with each resource type in a resource share.
	PermissionARNs []*string `json:"permissionARNs,omitempty"`
	// The principals to associate with the resource share. The possible values
	// are:
	//
	//    * An Amazon Web Services account ID
	//
	//    * An Amazon Resource Name (ARN) of an organization in Organizations
	//
	//    * An ARN of an organizational unit (OU) in Organizations
	//
	//    * An ARN of an IAM role
	//
	//    * An ARN of an IAM user
	//
	// Not all resource types can be shared with IAM roles and IAM users. For more
	// information, see Sharing with IAM roles and IAM users (https://docs.aws.amazon.com/ram/latest/userguide/permissions.html#permissions-rbp-supported-resource-types)
	// in the Resource Access Manager User Guide.
	Principals []*string `json:"principals,omitempty"`
	// The ARNs of the resources to associate with the resource share.
	ResourceARNs []*string `json:"resourceARNs,omitempty"`
	// One or more tags.
	Tags                          []*Tag `json:"tags,omitempty"`
	CustomResourceShareParameters `json:",inline"`
}

// ResourceShareSpec defines the desired state of ResourceShare
type ResourceShareSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ResourceShareParameters `json:"forProvider"`
}

// ResourceShareObservation defines the observed state of ResourceShare
type ResourceShareObservation struct {
	// Information about the resource share.
	ResourceShare *ResourceShare_SDK `json:"resourceShare,omitempty"`
}

// ResourceShareStatus defines the observed state of ResourceShare.
type ResourceShareStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ResourceShareObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceShare is the Schema for the ResourceShares API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ResourceShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceShareSpec   `json:"spec"`
	Status            ResourceShareStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceShareList contains a list of ResourceShares
type ResourceShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceShare `json:"items"`
}

// Repository type metadata.
var (
	ResourceShareKind             = "ResourceShare"
	ResourceShareGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceShareKind}.String()
	ResourceShareKindAPIVersion   = ResourceShareKind + "." + GroupVersion.String()
	ResourceShareGroupVersionKind = GroupVersion.WithKind(ResourceShareKind)
)

func init() {
	SchemeBuilder.Register(&ResourceShare{}, &ResourceShareList{})
}
