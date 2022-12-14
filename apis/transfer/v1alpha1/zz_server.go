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

// ServerParameters defines the desired state of Server
type ServerParameters struct {
	// Region is which region the Server will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The domain of the storage system that is used for file transfers. There are
	// two domains available: Amazon Simple Storage Service (Amazon S3) and Amazon
	// Elastic File System (Amazon EFS). The default value is S3.
	//
	// After the server is created, the domain cannot be changed.
	Domain *string `json:"domain,omitempty"`
	// The type of endpoint that you want your server to use. You can choose to
	// make your server's endpoint publicly accessible (PUBLIC) or host it inside
	// your VPC. With an endpoint that is hosted in a VPC, you can restrict access
	// to your server and resources only within your VPC or choose to make it internet
	// facing by attaching Elastic IP addresses directly to it.
	//
	// After May 19, 2021, you won't be able to create a server using EndpointType=VPC_ENDPOINT
	// in your Amazon Web Services account if your account hasn't already done so
	// before May 19, 2021. If you have already created servers with EndpointType=VPC_ENDPOINT
	// in your Amazon Web Services account on or before May 19, 2021, you will not
	// be affected. After this date, use EndpointType=VPC.
	//
	// For more information, see https://docs.aws.amazon.com/transfer/latest/userguide/create-server-in-vpc.html#deprecate-vpc-endpoint.
	//
	// It is recommended that you use VPC as the EndpointType. With this endpoint
	// type, you have the option to directly associate up to three Elastic IPv4
	// addresses (BYO IP included) with your server's endpoint and use VPC security
	// groups to restrict traffic by the client's public IP address. This is not
	// possible with EndpointType set to VPC_ENDPOINT.
	EndpointType *string `json:"endpointType,omitempty"`
	// The RSA private key as generated by the ssh-keygen -N "" -m PEM -f my-new-server-key
	// command.
	//
	// If you aren't planning to migrate existing users from an existing SFTP-enabled
	// server to a new server, don't update the host key. Accidentally changing
	// a server's host key can be disruptive.
	//
	// For more information, see Change the host key for your SFTP-enabled server
	// (https://docs.aws.amazon.com/transfer/latest/userguide/edit-server-config.html#configuring-servers-change-host-key)
	// in the Amazon Web Services Transfer Family User Guide.
	HostKey *string `json:"hostKey,omitempty"`
	// Required when IdentityProviderType is set to AWS_DIRECTORY_SERVICE or API_GATEWAY.
	// Accepts an array containing all of the information required to use a directory
	// in AWS_DIRECTORY_SERVICE or invoke a customer-supplied authentication API,
	// including the API Gateway URL. Not required when IdentityProviderType is
	// set to SERVICE_MANAGED.
	IdentityProviderDetails *IdentityProviderDetails `json:"identityProviderDetails,omitempty"`
	// Specifies the mode of authentication for a server. The default value is SERVICE_MANAGED,
	// which allows you to store and access user credentials within the Amazon Web
	// Services Transfer Family service.
	//
	// Use AWS_DIRECTORY_SERVICE to provide access to Active Directory groups in
	// Amazon Web Services Managed Active Directory or Microsoft Active Directory
	// in your on-premises environment or in Amazon Web Services using AD Connectors.
	// This option also requires you to provide a Directory ID using the IdentityProviderDetails
	// parameter.
	//
	// Use the API_GATEWAY value to integrate with an identity provider of your
	// choosing. The API_GATEWAY setting requires you to provide an API Gateway
	// endpoint URL to call for authentication using the IdentityProviderDetails
	// parameter.
	IdentityProviderType *string `json:"identityProviderType,omitempty"`
	// Specifies the file transfer protocol or protocols over which your file transfer
	// protocol client can connect to your server's endpoint. The available protocols
	// are:
	//
	//    * SFTP (Secure Shell (SSH) File Transfer Protocol): File transfer over
	//    SSH
	//
	//    * FTPS (File Transfer Protocol Secure): File transfer with TLS encryption
	//
	//    * FTP (File Transfer Protocol): Unencrypted file transfer
	//
	// If you select FTPS, you must choose a certificate stored in Amazon Web Services
	// Certificate Manager (ACM) which is used to identify your server when clients
	// connect to it over FTPS.
	//
	// If Protocol includes either FTP or FTPS, then the EndpointType must be VPC
	// and the IdentityProviderType must be AWS_DIRECTORY_SERVICE or API_GATEWAY.
	//
	// If Protocol includes FTP, then AddressAllocationIds cannot be associated.
	//
	// If Protocol is set only to SFTP, the EndpointType can be set to PUBLIC and
	// the IdentityProviderType can be set to SERVICE_MANAGED.
	Protocols []*string `json:"protocols,omitempty"`
	// Specifies the name of the security policy that is attached to the server.
	SecurityPolicyName *string `json:"securityPolicyName,omitempty"`
	// Key-value pairs that can be used to group and search for servers.
	Tags []*Tag `json:"tags,omitempty"`
	// Specifies the workflow ID for the workflow to assign and the execution role
	// used for executing the workflow.
	WorkflowDetails        *WorkflowDetails `json:"workflowDetails,omitempty"`
	CustomServerParameters `json:",inline"`
}

// ServerSpec defines the desired state of Server
type ServerSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ServerParameters `json:"forProvider"`
}

// ServerObservation defines the observed state of Server
type ServerObservation struct {
	// The service-assigned ID of the server that is created.
	ServerID *string `json:"serverID,omitempty"`
}

// ServerStatus defines the observed state of Server.
type ServerStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ServerObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Server is the Schema for the Servers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Server struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServerSpec   `json:"spec"`
	Status            ServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerList contains a list of Servers
type ServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Server `json:"items"`
}

// Repository type metadata.
var (
	ServerKind             = "Server"
	ServerGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerKind}.String()
	ServerKindAPIVersion   = ServerKind + "." + GroupVersion.String()
	ServerGroupVersionKind = GroupVersion.WithKind(ServerKind)
)

func init() {
	SchemeBuilder.Register(&Server{}, &ServerList{})
}
