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

type AuthorizationType string

const (
	AuthorizationType_NONE    AuthorizationType = "NONE"
	AuthorizationType_AWS_IAM AuthorizationType = "AWS_IAM"
	AuthorizationType_CUSTOM  AuthorizationType = "CUSTOM"
	AuthorizationType_JWT     AuthorizationType = "JWT"
)

type AuthorizerType string

const (
	AuthorizerType_REQUEST AuthorizerType = "REQUEST"
	AuthorizerType_JWT     AuthorizerType = "JWT"
)

type ConnectionType string

const (
	ConnectionType_INTERNET ConnectionType = "INTERNET"
	ConnectionType_VPC_LINK ConnectionType = "VPC_LINK"
)

type ContentHandlingStrategy string

const (
	ContentHandlingStrategy_CONVERT_TO_BINARY ContentHandlingStrategy = "CONVERT_TO_BINARY"
	ContentHandlingStrategy_CONVERT_TO_TEXT   ContentHandlingStrategy = "CONVERT_TO_TEXT"
)

type DeploymentStatus_SDK string

const (
	DeploymentStatus_SDK_PENDING  DeploymentStatus_SDK = "PENDING"
	DeploymentStatus_SDK_FAILED   DeploymentStatus_SDK = "FAILED"
	DeploymentStatus_SDK_DEPLOYED DeploymentStatus_SDK = "DEPLOYED"
)

type DomainNameStatus_SDK string

const (
	DomainNameStatus_SDK_AVAILABLE                      DomainNameStatus_SDK = "AVAILABLE"
	DomainNameStatus_SDK_UPDATING                       DomainNameStatus_SDK = "UPDATING"
	DomainNameStatus_SDK_PENDING_CERTIFICATE_REIMPORT   DomainNameStatus_SDK = "PENDING_CERTIFICATE_REIMPORT"
	DomainNameStatus_SDK_PENDING_OWNERSHIP_VERIFICATION DomainNameStatus_SDK = "PENDING_OWNERSHIP_VERIFICATION"
)

type EndpointType string

const (
	EndpointType_REGIONAL EndpointType = "REGIONAL"
	EndpointType_EDGE     EndpointType = "EDGE"
)

type IntegrationType string

const (
	IntegrationType_AWS        IntegrationType = "AWS"
	IntegrationType_HTTP       IntegrationType = "HTTP"
	IntegrationType_MOCK       IntegrationType = "MOCK"
	IntegrationType_HTTP_PROXY IntegrationType = "HTTP_PROXY"
	IntegrationType_AWS_PROXY  IntegrationType = "AWS_PROXY"
)

type LoggingLevel string

const (
	LoggingLevel_ERROR LoggingLevel = "ERROR"
	LoggingLevel_INFO  LoggingLevel = "INFO"
	LoggingLevel_OFF   LoggingLevel = "OFF"
)

type PassthroughBehavior string

const (
	PassthroughBehavior_WHEN_NO_MATCH     PassthroughBehavior = "WHEN_NO_MATCH"
	PassthroughBehavior_NEVER             PassthroughBehavior = "NEVER"
	PassthroughBehavior_WHEN_NO_TEMPLATES PassthroughBehavior = "WHEN_NO_TEMPLATES"
)

type ProtocolType string

const (
	ProtocolType_WEBSOCKET ProtocolType = "WEBSOCKET"
	ProtocolType_HTTP      ProtocolType = "HTTP"
)

type SecurityPolicy string

const (
	SecurityPolicy_TLS_1_0 SecurityPolicy = "TLS_1_0"
	SecurityPolicy_TLS_1_2 SecurityPolicy = "TLS_1_2"
)

type VPCLinkStatus_SDK string

const (
	VPCLinkStatus_SDK_PENDING   VPCLinkStatus_SDK = "PENDING"
	VPCLinkStatus_SDK_AVAILABLE VPCLinkStatus_SDK = "AVAILABLE"
	VPCLinkStatus_SDK_DELETING  VPCLinkStatus_SDK = "DELETING"
	VPCLinkStatus_SDK_FAILED    VPCLinkStatus_SDK = "FAILED"
	VPCLinkStatus_SDK_INACTIVE  VPCLinkStatus_SDK = "INACTIVE"
)

type VPCLinkVersion string

const (
	VPCLinkVersion_V2 VPCLinkVersion = "V2"
)
