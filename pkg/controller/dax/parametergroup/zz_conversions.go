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

package parametergroup

import (
	"github.com/aws/aws-sdk-go/aws/awserr"
	svcsdk "github.com/aws/aws-sdk-go/service/dax"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/dax/v1alpha1"
)

// NOTE(muvaf): We return pointers in case the function needs to start with an
// empty object, hence need to return a new pointer.

// GenerateDescribeParameterGroupsInput returns input for read
// operation.
func GenerateDescribeParameterGroupsInput(cr *svcapitypes.ParameterGroup) *svcsdk.DescribeParameterGroupsInput {
	res := &svcsdk.DescribeParameterGroupsInput{}

	if cr.Status.AtProvider.ParameterGroupName != nil {
		f2 := []*string{}
		f2 = append(f2, cr.Status.AtProvider.ParameterGroupName)
		res.SetParameterGroupNames(f2)
	}

	return res
}

// GenerateParameterGroup returns the current state in the form of *svcapitypes.ParameterGroup.
func GenerateParameterGroup(resp *svcsdk.DescribeParameterGroupsOutput) *svcapitypes.ParameterGroup {
	cr := &svcapitypes.ParameterGroup{}

	found := false
	for _, elem := range resp.ParameterGroups {
		if elem.Description != nil {
			cr.Spec.ForProvider.Description = elem.Description
		} else {
			cr.Spec.ForProvider.Description = nil
		}
		if elem.ParameterGroupName != nil {
			cr.Status.AtProvider.ParameterGroupName = elem.ParameterGroupName
		} else {
			cr.Status.AtProvider.ParameterGroupName = nil
		}
		found = true
		break
	}
	if !found {
		return cr
	}

	return cr
}

// GenerateCreateParameterGroupInput returns a create input.
func GenerateCreateParameterGroupInput(cr *svcapitypes.ParameterGroup) *svcsdk.CreateParameterGroupInput {
	res := &svcsdk.CreateParameterGroupInput{}

	if cr.Spec.ForProvider.Description != nil {
		res.SetDescription(*cr.Spec.ForProvider.Description)
	}

	return res
}

// GenerateUpdateParameterGroupInput returns an update input.
func GenerateUpdateParameterGroupInput(cr *svcapitypes.ParameterGroup) *svcsdk.UpdateParameterGroupInput {
	res := &svcsdk.UpdateParameterGroupInput{}

	if cr.Status.AtProvider.ParameterGroupName != nil {
		res.SetParameterGroupName(*cr.Status.AtProvider.ParameterGroupName)
	}

	return res
}

// GenerateDeleteParameterGroupInput returns a deletion input.
func GenerateDeleteParameterGroupInput(cr *svcapitypes.ParameterGroup) *svcsdk.DeleteParameterGroupInput {
	res := &svcsdk.DeleteParameterGroupInput{}

	if cr.Status.AtProvider.ParameterGroupName != nil {
		res.SetParameterGroupName(*cr.Status.AtProvider.ParameterGroupName)
	}

	return res
}

// IsNotFound returns whether the given error is of type NotFound or not.
func IsNotFound(err error) bool {
	awsErr, ok := err.(awserr.Error)
	return ok && awsErr.Code() == "ParameterGroupNotFoundFault"
}
