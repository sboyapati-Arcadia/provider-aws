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

package filesystem

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/efs"
	svcsdk "github.com/aws/aws-sdk-go/service/efs"
	svcsdkapi "github.com/aws/aws-sdk-go/service/efs/efsiface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/efs/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an FileSystem resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create FileSystem in AWS"
	errUpdate        = "cannot update FileSystem in AWS"
	errDescribe      = "failed to describe FileSystem"
	errDelete        = "failed to delete FileSystem"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.FileSystem)
	if !ok {
		return nil, errors.New(errUnexpectedObject)
	}
	sess, err := awsclient.GetConfigV1(ctx, c.kube, mg, cr.Spec.ForProvider.Region)
	if err != nil {
		return nil, errors.Wrap(err, errCreateSession)
	}
	return newExternal(c.kube, svcapi.New(sess), c.opts), nil
}

func (e *external) Observe(ctx context.Context, mg cpresource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*svcapitypes.FileSystem)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeFileSystemsInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeFileSystemsWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	resp = e.filterList(cr, resp)
	if len(resp.FileSystems) == 0 {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateFileSystem(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

	upToDate, err := e.isUpToDate(cr, resp)
	if err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "isUpToDate check failed")
	}
	return e.postObserve(ctx, cr, resp, managed.ExternalObservation{
		ResourceExists:          true,
		ResourceUpToDate:        upToDate,
		ResourceLateInitialized: !cmp.Equal(&cr.Spec.ForProvider, currentSpec),
	}, nil)
}

func (e *external) Create(ctx context.Context, mg cpresource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*svcapitypes.FileSystem)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateFileSystemInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateFileSystemWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.AvailabilityZoneId != nil {
		cr.Status.AtProvider.AvailabilityZoneID = resp.AvailabilityZoneId
	} else {
		cr.Status.AtProvider.AvailabilityZoneID = nil
	}
	if resp.AvailabilityZoneName != nil {
		cr.Spec.ForProvider.AvailabilityZoneName = resp.AvailabilityZoneName
	} else {
		cr.Spec.ForProvider.AvailabilityZoneName = nil
	}
	if resp.CreationTime != nil {
		cr.Status.AtProvider.CreationTime = &metav1.Time{*resp.CreationTime}
	} else {
		cr.Status.AtProvider.CreationTime = nil
	}
	if resp.CreationToken != nil {
		cr.Status.AtProvider.CreationToken = resp.CreationToken
	} else {
		cr.Status.AtProvider.CreationToken = nil
	}
	if resp.Encrypted != nil {
		cr.Spec.ForProvider.Encrypted = resp.Encrypted
	} else {
		cr.Spec.ForProvider.Encrypted = nil
	}
	if resp.FileSystemArn != nil {
		cr.Status.AtProvider.FileSystemARN = resp.FileSystemArn
	} else {
		cr.Status.AtProvider.FileSystemARN = nil
	}
	if resp.FileSystemId != nil {
		cr.Status.AtProvider.FileSystemID = resp.FileSystemId
	} else {
		cr.Status.AtProvider.FileSystemID = nil
	}
	if resp.KmsKeyId != nil {
		cr.Spec.ForProvider.KMSKeyID = resp.KmsKeyId
	} else {
		cr.Spec.ForProvider.KMSKeyID = nil
	}
	if resp.LifeCycleState != nil {
		cr.Status.AtProvider.LifeCycleState = resp.LifeCycleState
	} else {
		cr.Status.AtProvider.LifeCycleState = nil
	}
	if resp.Name != nil {
		cr.Status.AtProvider.Name = resp.Name
	} else {
		cr.Status.AtProvider.Name = nil
	}
	if resp.NumberOfMountTargets != nil {
		cr.Status.AtProvider.NumberOfMountTargets = resp.NumberOfMountTargets
	} else {
		cr.Status.AtProvider.NumberOfMountTargets = nil
	}
	if resp.OwnerId != nil {
		cr.Status.AtProvider.OwnerID = resp.OwnerId
	} else {
		cr.Status.AtProvider.OwnerID = nil
	}
	if resp.PerformanceMode != nil {
		cr.Spec.ForProvider.PerformanceMode = resp.PerformanceMode
	} else {
		cr.Spec.ForProvider.PerformanceMode = nil
	}
	if resp.SizeInBytes != nil {
		f13 := &svcapitypes.FileSystemSize{}
		if resp.SizeInBytes.Timestamp != nil {
			f13.Timestamp = &metav1.Time{*resp.SizeInBytes.Timestamp}
		}
		if resp.SizeInBytes.Value != nil {
			f13.Value = resp.SizeInBytes.Value
		}
		if resp.SizeInBytes.ValueInIA != nil {
			f13.ValueInIA = resp.SizeInBytes.ValueInIA
		}
		if resp.SizeInBytes.ValueInStandard != nil {
			f13.ValueInStandard = resp.SizeInBytes.ValueInStandard
		}
		cr.Status.AtProvider.SizeInBytes = f13
	} else {
		cr.Status.AtProvider.SizeInBytes = nil
	}
	if resp.Tags != nil {
		f14 := []*svcapitypes.Tag{}
		for _, f14iter := range resp.Tags {
			f14elem := &svcapitypes.Tag{}
			if f14iter.Key != nil {
				f14elem.Key = f14iter.Key
			}
			if f14iter.Value != nil {
				f14elem.Value = f14iter.Value
			}
			f14 = append(f14, f14elem)
		}
		cr.Spec.ForProvider.Tags = f14
	} else {
		cr.Spec.ForProvider.Tags = nil
	}
	if resp.ThroughputMode != nil {
		cr.Spec.ForProvider.ThroughputMode = resp.ThroughputMode
	} else {
		cr.Spec.ForProvider.ThroughputMode = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.FileSystem)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateFileSystemInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateFileSystemWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.FileSystem)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteFileSystemInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteFileSystemWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.EFSAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
		filterList:     nopFilterList,
		preCreate:      nopPreCreate,
		postCreate:     nopPostCreate,
		preDelete:      nopPreDelete,
		postDelete:     nopPostDelete,
		preUpdate:      nopPreUpdate,
		postUpdate:     nopPostUpdate,
	}
	for _, f := range opts {
		f(e)
	}
	return e
}

type external struct {
	kube           client.Client
	client         svcsdkapi.EFSAPI
	preObserve     func(context.Context, *svcapitypes.FileSystem, *svcsdk.DescribeFileSystemsInput) error
	postObserve    func(context.Context, *svcapitypes.FileSystem, *svcsdk.DescribeFileSystemsOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	filterList     func(*svcapitypes.FileSystem, *svcsdk.DescribeFileSystemsOutput) *svcsdk.DescribeFileSystemsOutput
	lateInitialize func(*svcapitypes.FileSystemParameters, *svcsdk.DescribeFileSystemsOutput) error
	isUpToDate     func(*svcapitypes.FileSystem, *svcsdk.DescribeFileSystemsOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.FileSystem, *svcsdk.CreateFileSystemInput) error
	postCreate     func(context.Context, *svcapitypes.FileSystem, *svcsdk.FileSystemDescription, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.FileSystem, *svcsdk.DeleteFileSystemInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.FileSystem, *svcsdk.DeleteFileSystemOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.FileSystem, *svcsdk.UpdateFileSystemInput) error
	postUpdate     func(context.Context, *svcapitypes.FileSystem, *svcsdk.UpdateFileSystemOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.FileSystem, *svcsdk.DescribeFileSystemsInput) error {
	return nil
}
func nopPostObserve(_ context.Context, _ *svcapitypes.FileSystem, _ *svcsdk.DescribeFileSystemsOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopFilterList(_ *svcapitypes.FileSystem, list *svcsdk.DescribeFileSystemsOutput) *svcsdk.DescribeFileSystemsOutput {
	return list
}

func nopLateInitialize(*svcapitypes.FileSystemParameters, *svcsdk.DescribeFileSystemsOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.FileSystem, *svcsdk.DescribeFileSystemsOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.FileSystem, *svcsdk.CreateFileSystemInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.FileSystem, _ *svcsdk.FileSystemDescription, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.FileSystem, *svcsdk.DeleteFileSystemInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.FileSystem, _ *svcsdk.DeleteFileSystemOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.FileSystem, *svcsdk.UpdateFileSystemInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.FileSystem, _ *svcsdk.UpdateFileSystemOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
