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

type ResourceOwner string

const (
	ResourceOwner_SELF           ResourceOwner = "SELF"
	ResourceOwner_OTHER_ACCOUNTS ResourceOwner = "OTHER-ACCOUNTS"
)

type ResourceShareAssociationStatus string

const (
	ResourceShareAssociationStatus_ASSOCIATING    ResourceShareAssociationStatus = "ASSOCIATING"
	ResourceShareAssociationStatus_ASSOCIATED     ResourceShareAssociationStatus = "ASSOCIATED"
	ResourceShareAssociationStatus_FAILED         ResourceShareAssociationStatus = "FAILED"
	ResourceShareAssociationStatus_DISASSOCIATING ResourceShareAssociationStatus = "DISASSOCIATING"
	ResourceShareAssociationStatus_DISASSOCIATED  ResourceShareAssociationStatus = "DISASSOCIATED"
)

type ResourceShareAssociationType string

const (
	ResourceShareAssociationType_PRINCIPAL ResourceShareAssociationType = "PRINCIPAL"
	ResourceShareAssociationType_RESOURCE  ResourceShareAssociationType = "RESOURCE"
)

type ResourceShareFeatureSet string

const (
	ResourceShareFeatureSet_CREATED_FROM_POLICY   ResourceShareFeatureSet = "CREATED_FROM_POLICY"
	ResourceShareFeatureSet_PROMOTING_TO_STANDARD ResourceShareFeatureSet = "PROMOTING_TO_STANDARD"
	ResourceShareFeatureSet_STANDARD              ResourceShareFeatureSet = "STANDARD"
)

type ResourceShareInvitationStatus string

const (
	ResourceShareInvitationStatus_PENDING  ResourceShareInvitationStatus = "PENDING"
	ResourceShareInvitationStatus_ACCEPTED ResourceShareInvitationStatus = "ACCEPTED"
	ResourceShareInvitationStatus_REJECTED ResourceShareInvitationStatus = "REJECTED"
	ResourceShareInvitationStatus_EXPIRED  ResourceShareInvitationStatus = "EXPIRED"
)

type ResourceShareStatus_SDK string

const (
	ResourceShareStatus_SDK_PENDING  ResourceShareStatus_SDK = "PENDING"
	ResourceShareStatus_SDK_ACTIVE   ResourceShareStatus_SDK = "ACTIVE"
	ResourceShareStatus_SDK_FAILED   ResourceShareStatus_SDK = "FAILED"
	ResourceShareStatus_SDK_DELETING ResourceShareStatus_SDK = "DELETING"
	ResourceShareStatus_SDK_DELETED  ResourceShareStatus_SDK = "DELETED"
)

type ResourceStatus string

const (
	ResourceStatus_AVAILABLE                   ResourceStatus = "AVAILABLE"
	ResourceStatus_ZONAL_RESOURCE_INACCESSIBLE ResourceStatus = "ZONAL_RESOURCE_INACCESSIBLE"
	ResourceStatus_LIMIT_EXCEEDED              ResourceStatus = "LIMIT_EXCEEDED"
	ResourceStatus_UNAVAILABLE                 ResourceStatus = "UNAVAILABLE"
	ResourceStatus_PENDING                     ResourceStatus = "PENDING"
)
