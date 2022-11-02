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

package userpool

import (
	"context"

	svcapi "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	svcsdk "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	svcsdkapi "github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	cpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/cognitoidentityprovider/v1alpha1"
	awsclient "github.com/crossplane-contrib/provider-aws/pkg/clients"
)

const (
	errUnexpectedObject = "managed resource is not an UserPool resource"

	errCreateSession = "cannot create a new session"
	errCreate        = "cannot create UserPool in AWS"
	errUpdate        = "cannot update UserPool in AWS"
	errDescribe      = "failed to describe UserPool"
	errDelete        = "failed to delete UserPool"
)

type connector struct {
	kube client.Client
	opts []option
}

func (c *connector) Connect(ctx context.Context, mg cpresource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*svcapitypes.UserPool)
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
	cr, ok := mg.(*svcapitypes.UserPool)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errUnexpectedObject)
	}
	if meta.GetExternalName(cr) == "" {
		return managed.ExternalObservation{
			ResourceExists: false,
		}, nil
	}
	input := GenerateDescribeUserPoolInput(cr)
	if err := e.preObserve(ctx, cr, input); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "pre-observe failed")
	}
	resp, err := e.client.DescribeUserPoolWithContext(ctx, input)
	if err != nil {
		return managed.ExternalObservation{ResourceExists: false}, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDescribe)
	}
	currentSpec := cr.Spec.ForProvider.DeepCopy()
	if err := e.lateInitialize(&cr.Spec.ForProvider, resp); err != nil {
		return managed.ExternalObservation{}, errors.Wrap(err, "late-init failed")
	}
	GenerateUserPool(resp).Status.AtProvider.DeepCopyInto(&cr.Status.AtProvider)

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
	cr, ok := mg.(*svcapitypes.UserPool)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Creating())
	input := GenerateCreateUserPoolInput(cr)
	if err := e.preCreate(ctx, cr, input); err != nil {
		return managed.ExternalCreation{}, errors.Wrap(err, "pre-create failed")
	}
	resp, err := e.client.CreateUserPoolWithContext(ctx, input)
	if err != nil {
		return managed.ExternalCreation{}, awsclient.Wrap(err, errCreate)
	}

	if resp.UserPool.AccountRecoverySetting != nil {
		f0 := &svcapitypes.AccountRecoverySettingType{}
		if resp.UserPool.AccountRecoverySetting.RecoveryMechanisms != nil {
			f0f0 := []*svcapitypes.RecoveryOptionType{}
			for _, f0f0iter := range resp.UserPool.AccountRecoverySetting.RecoveryMechanisms {
				f0f0elem := &svcapitypes.RecoveryOptionType{}
				if f0f0iter.Name != nil {
					f0f0elem.Name = f0f0iter.Name
				}
				if f0f0iter.Priority != nil {
					f0f0elem.Priority = f0f0iter.Priority
				}
				f0f0 = append(f0f0, f0f0elem)
			}
			f0.RecoveryMechanisms = f0f0
		}
		cr.Spec.ForProvider.AccountRecoverySetting = f0
	} else {
		cr.Spec.ForProvider.AccountRecoverySetting = nil
	}
	if resp.UserPool.AdminCreateUserConfig != nil {
		f1 := &svcapitypes.AdminCreateUserConfigType{}
		if resp.UserPool.AdminCreateUserConfig.AllowAdminCreateUserOnly != nil {
			f1.AllowAdminCreateUserOnly = resp.UserPool.AdminCreateUserConfig.AllowAdminCreateUserOnly
		}
		if resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate != nil {
			f1f1 := &svcapitypes.MessageTemplateType{}
			if resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate.EmailMessage != nil {
				f1f1.EmailMessage = resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate.EmailMessage
			}
			if resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate.EmailSubject != nil {
				f1f1.EmailSubject = resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate.EmailSubject
			}
			if resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate.SMSMessage != nil {
				f1f1.SMSMessage = resp.UserPool.AdminCreateUserConfig.InviteMessageTemplate.SMSMessage
			}
			f1.InviteMessageTemplate = f1f1
		}
		if resp.UserPool.AdminCreateUserConfig.UnusedAccountValidityDays != nil {
			f1.UnusedAccountValidityDays = resp.UserPool.AdminCreateUserConfig.UnusedAccountValidityDays
		}
		cr.Spec.ForProvider.AdminCreateUserConfig = f1
	} else {
		cr.Spec.ForProvider.AdminCreateUserConfig = nil
	}
	if resp.UserPool.AliasAttributes != nil {
		f2 := []*string{}
		for _, f2iter := range resp.UserPool.AliasAttributes {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		cr.Spec.ForProvider.AliasAttributes = f2
	} else {
		cr.Spec.ForProvider.AliasAttributes = nil
	}
	if resp.UserPool.Arn != nil {
		cr.Status.AtProvider.ARN = resp.UserPool.Arn
	} else {
		cr.Status.AtProvider.ARN = nil
	}
	if resp.UserPool.AutoVerifiedAttributes != nil {
		f4 := []*string{}
		for _, f4iter := range resp.UserPool.AutoVerifiedAttributes {
			var f4elem string
			f4elem = *f4iter
			f4 = append(f4, &f4elem)
		}
		cr.Spec.ForProvider.AutoVerifiedAttributes = f4
	} else {
		cr.Spec.ForProvider.AutoVerifiedAttributes = nil
	}
	if resp.UserPool.CreationDate != nil {
		cr.Status.AtProvider.CreationDate = &metav1.Time{*resp.UserPool.CreationDate}
	} else {
		cr.Status.AtProvider.CreationDate = nil
	}
	if resp.UserPool.CustomDomain != nil {
		cr.Status.AtProvider.CustomDomain = resp.UserPool.CustomDomain
	} else {
		cr.Status.AtProvider.CustomDomain = nil
	}
	if resp.UserPool.DeviceConfiguration != nil {
		f7 := &svcapitypes.DeviceConfigurationType{}
		if resp.UserPool.DeviceConfiguration.ChallengeRequiredOnNewDevice != nil {
			f7.ChallengeRequiredOnNewDevice = resp.UserPool.DeviceConfiguration.ChallengeRequiredOnNewDevice
		}
		if resp.UserPool.DeviceConfiguration.DeviceOnlyRememberedOnUserPrompt != nil {
			f7.DeviceOnlyRememberedOnUserPrompt = resp.UserPool.DeviceConfiguration.DeviceOnlyRememberedOnUserPrompt
		}
		cr.Spec.ForProvider.DeviceConfiguration = f7
	} else {
		cr.Spec.ForProvider.DeviceConfiguration = nil
	}
	if resp.UserPool.Domain != nil {
		cr.Status.AtProvider.Domain = resp.UserPool.Domain
	} else {
		cr.Status.AtProvider.Domain = nil
	}
	if resp.UserPool.EmailConfiguration != nil {
		f9 := &svcapitypes.EmailConfigurationType{}
		if resp.UserPool.EmailConfiguration.ConfigurationSet != nil {
			f9.ConfigurationSet = resp.UserPool.EmailConfiguration.ConfigurationSet
		}
		if resp.UserPool.EmailConfiguration.EmailSendingAccount != nil {
			f9.EmailSendingAccount = resp.UserPool.EmailConfiguration.EmailSendingAccount
		}
		if resp.UserPool.EmailConfiguration.From != nil {
			f9.From = resp.UserPool.EmailConfiguration.From
		}
		if resp.UserPool.EmailConfiguration.ReplyToEmailAddress != nil {
			f9.ReplyToEmailAddress = resp.UserPool.EmailConfiguration.ReplyToEmailAddress
		}
		if resp.UserPool.EmailConfiguration.SourceArn != nil {
			f9.SourceARN = resp.UserPool.EmailConfiguration.SourceArn
		}
		cr.Spec.ForProvider.EmailConfiguration = f9
	} else {
		cr.Spec.ForProvider.EmailConfiguration = nil
	}
	if resp.UserPool.EmailConfigurationFailure != nil {
		cr.Status.AtProvider.EmailConfigurationFailure = resp.UserPool.EmailConfigurationFailure
	} else {
		cr.Status.AtProvider.EmailConfigurationFailure = nil
	}
	if resp.UserPool.EmailVerificationMessage != nil {
		cr.Spec.ForProvider.EmailVerificationMessage = resp.UserPool.EmailVerificationMessage
	} else {
		cr.Spec.ForProvider.EmailVerificationMessage = nil
	}
	if resp.UserPool.EmailVerificationSubject != nil {
		cr.Spec.ForProvider.EmailVerificationSubject = resp.UserPool.EmailVerificationSubject
	} else {
		cr.Spec.ForProvider.EmailVerificationSubject = nil
	}
	if resp.UserPool.EstimatedNumberOfUsers != nil {
		cr.Status.AtProvider.EstimatedNumberOfUsers = resp.UserPool.EstimatedNumberOfUsers
	} else {
		cr.Status.AtProvider.EstimatedNumberOfUsers = nil
	}
	if resp.UserPool.Id != nil {
		cr.Status.AtProvider.ID = resp.UserPool.Id
	} else {
		cr.Status.AtProvider.ID = nil
	}
	if resp.UserPool.LambdaConfig != nil {
		f15 := &svcapitypes.LambdaConfigType{}
		if resp.UserPool.LambdaConfig.CreateAuthChallenge != nil {
			f15.CreateAuthChallenge = resp.UserPool.LambdaConfig.CreateAuthChallenge
		}
		if resp.UserPool.LambdaConfig.CustomEmailSender != nil {
			f15f1 := &svcapitypes.CustomEmailLambdaVersionConfigType{}
			if resp.UserPool.LambdaConfig.CustomEmailSender.LambdaArn != nil {
				f15f1.LambdaARN = resp.UserPool.LambdaConfig.CustomEmailSender.LambdaArn
			}
			if resp.UserPool.LambdaConfig.CustomEmailSender.LambdaVersion != nil {
				f15f1.LambdaVersion = resp.UserPool.LambdaConfig.CustomEmailSender.LambdaVersion
			}
			f15.CustomEmailSender = f15f1
		}
		if resp.UserPool.LambdaConfig.CustomMessage != nil {
			f15.CustomMessage = resp.UserPool.LambdaConfig.CustomMessage
		}
		if resp.UserPool.LambdaConfig.CustomSMSSender != nil {
			f15f3 := &svcapitypes.CustomSMSLambdaVersionConfigType{}
			if resp.UserPool.LambdaConfig.CustomSMSSender.LambdaArn != nil {
				f15f3.LambdaARN = resp.UserPool.LambdaConfig.CustomSMSSender.LambdaArn
			}
			if resp.UserPool.LambdaConfig.CustomSMSSender.LambdaVersion != nil {
				f15f3.LambdaVersion = resp.UserPool.LambdaConfig.CustomSMSSender.LambdaVersion
			}
			f15.CustomSMSSender = f15f3
		}
		if resp.UserPool.LambdaConfig.DefineAuthChallenge != nil {
			f15.DefineAuthChallenge = resp.UserPool.LambdaConfig.DefineAuthChallenge
		}
		if resp.UserPool.LambdaConfig.KMSKeyID != nil {
			f15.KMSKeyID = resp.UserPool.LambdaConfig.KMSKeyID
		}
		if resp.UserPool.LambdaConfig.PostAuthentication != nil {
			f15.PostAuthentication = resp.UserPool.LambdaConfig.PostAuthentication
		}
		if resp.UserPool.LambdaConfig.PostConfirmation != nil {
			f15.PostConfirmation = resp.UserPool.LambdaConfig.PostConfirmation
		}
		if resp.UserPool.LambdaConfig.PreAuthentication != nil {
			f15.PreAuthentication = resp.UserPool.LambdaConfig.PreAuthentication
		}
		if resp.UserPool.LambdaConfig.PreSignUp != nil {
			f15.PreSignUp = resp.UserPool.LambdaConfig.PreSignUp
		}
		if resp.UserPool.LambdaConfig.PreTokenGeneration != nil {
			f15.PreTokenGeneration = resp.UserPool.LambdaConfig.PreTokenGeneration
		}
		if resp.UserPool.LambdaConfig.UserMigration != nil {
			f15.UserMigration = resp.UserPool.LambdaConfig.UserMigration
		}
		if resp.UserPool.LambdaConfig.VerifyAuthChallengeResponse != nil {
			f15.VerifyAuthChallengeResponse = resp.UserPool.LambdaConfig.VerifyAuthChallengeResponse
		}
		cr.Spec.ForProvider.LambdaConfig = f15
	} else {
		cr.Spec.ForProvider.LambdaConfig = nil
	}
	if resp.UserPool.LastModifiedDate != nil {
		cr.Status.AtProvider.LastModifiedDate = &metav1.Time{*resp.UserPool.LastModifiedDate}
	} else {
		cr.Status.AtProvider.LastModifiedDate = nil
	}
	if resp.UserPool.MfaConfiguration != nil {
		cr.Spec.ForProvider.MFAConfiguration = resp.UserPool.MfaConfiguration
	} else {
		cr.Spec.ForProvider.MFAConfiguration = nil
	}
	if resp.UserPool.Name != nil {
		cr.Status.AtProvider.Name = resp.UserPool.Name
	} else {
		cr.Status.AtProvider.Name = nil
	}
	if resp.UserPool.Policies != nil {
		f19 := &svcapitypes.UserPoolPolicyType{}
		if resp.UserPool.Policies.PasswordPolicy != nil {
			f19f0 := &svcapitypes.PasswordPolicyType{}
			if resp.UserPool.Policies.PasswordPolicy.MinimumLength != nil {
				f19f0.MinimumLength = resp.UserPool.Policies.PasswordPolicy.MinimumLength
			}
			if resp.UserPool.Policies.PasswordPolicy.RequireLowercase != nil {
				f19f0.RequireLowercase = resp.UserPool.Policies.PasswordPolicy.RequireLowercase
			}
			if resp.UserPool.Policies.PasswordPolicy.RequireNumbers != nil {
				f19f0.RequireNumbers = resp.UserPool.Policies.PasswordPolicy.RequireNumbers
			}
			if resp.UserPool.Policies.PasswordPolicy.RequireSymbols != nil {
				f19f0.RequireSymbols = resp.UserPool.Policies.PasswordPolicy.RequireSymbols
			}
			if resp.UserPool.Policies.PasswordPolicy.RequireUppercase != nil {
				f19f0.RequireUppercase = resp.UserPool.Policies.PasswordPolicy.RequireUppercase
			}
			if resp.UserPool.Policies.PasswordPolicy.TemporaryPasswordValidityDays != nil {
				f19f0.TemporaryPasswordValidityDays = resp.UserPool.Policies.PasswordPolicy.TemporaryPasswordValidityDays
			}
			f19.PasswordPolicy = f19f0
		}
		cr.Spec.ForProvider.Policies = f19
	} else {
		cr.Spec.ForProvider.Policies = nil
	}
	if resp.UserPool.SchemaAttributes != nil {
		f20 := []*svcapitypes.SchemaAttributeType{}
		for _, f20iter := range resp.UserPool.SchemaAttributes {
			f20elem := &svcapitypes.SchemaAttributeType{}
			if f20iter.AttributeDataType != nil {
				f20elem.AttributeDataType = f20iter.AttributeDataType
			}
			if f20iter.DeveloperOnlyAttribute != nil {
				f20elem.DeveloperOnlyAttribute = f20iter.DeveloperOnlyAttribute
			}
			if f20iter.Mutable != nil {
				f20elem.Mutable = f20iter.Mutable
			}
			if f20iter.Name != nil {
				f20elem.Name = f20iter.Name
			}
			if f20iter.NumberAttributeConstraints != nil {
				f20elemf4 := &svcapitypes.NumberAttributeConstraintsType{}
				if f20iter.NumberAttributeConstraints.MaxValue != nil {
					f20elemf4.MaxValue = f20iter.NumberAttributeConstraints.MaxValue
				}
				if f20iter.NumberAttributeConstraints.MinValue != nil {
					f20elemf4.MinValue = f20iter.NumberAttributeConstraints.MinValue
				}
				f20elem.NumberAttributeConstraints = f20elemf4
			}
			if f20iter.Required != nil {
				f20elem.Required = f20iter.Required
			}
			if f20iter.StringAttributeConstraints != nil {
				f20elemf6 := &svcapitypes.StringAttributeConstraintsType{}
				if f20iter.StringAttributeConstraints.MaxLength != nil {
					f20elemf6.MaxLength = f20iter.StringAttributeConstraints.MaxLength
				}
				if f20iter.StringAttributeConstraints.MinLength != nil {
					f20elemf6.MinLength = f20iter.StringAttributeConstraints.MinLength
				}
				f20elem.StringAttributeConstraints = f20elemf6
			}
			f20 = append(f20, f20elem)
		}
		cr.Status.AtProvider.SchemaAttributes = f20
	} else {
		cr.Status.AtProvider.SchemaAttributes = nil
	}
	if resp.UserPool.SmsAuthenticationMessage != nil {
		cr.Spec.ForProvider.SmsAuthenticationMessage = resp.UserPool.SmsAuthenticationMessage
	} else {
		cr.Spec.ForProvider.SmsAuthenticationMessage = nil
	}
	if resp.UserPool.SmsConfiguration != nil {
		f22 := &svcapitypes.SmsConfigurationType{}
		if resp.UserPool.SmsConfiguration.ExternalId != nil {
			f22.ExternalID = resp.UserPool.SmsConfiguration.ExternalId
		}
		if resp.UserPool.SmsConfiguration.SnsCallerArn != nil {
			f22.SNSCallerARN = resp.UserPool.SmsConfiguration.SnsCallerArn
		}
		cr.Spec.ForProvider.SmsConfiguration = f22
	} else {
		cr.Spec.ForProvider.SmsConfiguration = nil
	}
	if resp.UserPool.SmsConfigurationFailure != nil {
		cr.Status.AtProvider.SmsConfigurationFailure = resp.UserPool.SmsConfigurationFailure
	} else {
		cr.Status.AtProvider.SmsConfigurationFailure = nil
	}
	if resp.UserPool.SmsVerificationMessage != nil {
		cr.Spec.ForProvider.SmsVerificationMessage = resp.UserPool.SmsVerificationMessage
	} else {
		cr.Spec.ForProvider.SmsVerificationMessage = nil
	}
	if resp.UserPool.Status != nil {
		cr.Status.AtProvider.Status = resp.UserPool.Status
	} else {
		cr.Status.AtProvider.Status = nil
	}
	if resp.UserPool.UserPoolAddOns != nil {
		f26 := &svcapitypes.UserPoolAddOnsType{}
		if resp.UserPool.UserPoolAddOns.AdvancedSecurityMode != nil {
			f26.AdvancedSecurityMode = resp.UserPool.UserPoolAddOns.AdvancedSecurityMode
		}
		cr.Spec.ForProvider.UserPoolAddOns = f26
	} else {
		cr.Spec.ForProvider.UserPoolAddOns = nil
	}
	if resp.UserPool.UserPoolTags != nil {
		f27 := map[string]*string{}
		for f27key, f27valiter := range resp.UserPool.UserPoolTags {
			var f27val string
			f27val = *f27valiter
			f27[f27key] = &f27val
		}
		cr.Spec.ForProvider.UserPoolTags = f27
	} else {
		cr.Spec.ForProvider.UserPoolTags = nil
	}
	if resp.UserPool.UsernameAttributes != nil {
		f28 := []*string{}
		for _, f28iter := range resp.UserPool.UsernameAttributes {
			var f28elem string
			f28elem = *f28iter
			f28 = append(f28, &f28elem)
		}
		cr.Spec.ForProvider.UsernameAttributes = f28
	} else {
		cr.Spec.ForProvider.UsernameAttributes = nil
	}
	if resp.UserPool.UsernameConfiguration != nil {
		f29 := &svcapitypes.UsernameConfigurationType{}
		if resp.UserPool.UsernameConfiguration.CaseSensitive != nil {
			f29.CaseSensitive = resp.UserPool.UsernameConfiguration.CaseSensitive
		}
		cr.Spec.ForProvider.UsernameConfiguration = f29
	} else {
		cr.Spec.ForProvider.UsernameConfiguration = nil
	}
	if resp.UserPool.VerificationMessageTemplate != nil {
		f30 := &svcapitypes.VerificationMessageTemplateType{}
		if resp.UserPool.VerificationMessageTemplate.DefaultEmailOption != nil {
			f30.DefaultEmailOption = resp.UserPool.VerificationMessageTemplate.DefaultEmailOption
		}
		if resp.UserPool.VerificationMessageTemplate.EmailMessage != nil {
			f30.EmailMessage = resp.UserPool.VerificationMessageTemplate.EmailMessage
		}
		if resp.UserPool.VerificationMessageTemplate.EmailMessageByLink != nil {
			f30.EmailMessageByLink = resp.UserPool.VerificationMessageTemplate.EmailMessageByLink
		}
		if resp.UserPool.VerificationMessageTemplate.EmailSubject != nil {
			f30.EmailSubject = resp.UserPool.VerificationMessageTemplate.EmailSubject
		}
		if resp.UserPool.VerificationMessageTemplate.EmailSubjectByLink != nil {
			f30.EmailSubjectByLink = resp.UserPool.VerificationMessageTemplate.EmailSubjectByLink
		}
		if resp.UserPool.VerificationMessageTemplate.SmsMessage != nil {
			f30.SmsMessage = resp.UserPool.VerificationMessageTemplate.SmsMessage
		}
		cr.Spec.ForProvider.VerificationMessageTemplate = f30
	} else {
		cr.Spec.ForProvider.VerificationMessageTemplate = nil
	}

	return e.postCreate(ctx, cr, resp, managed.ExternalCreation{}, err)
}

func (e *external) Update(ctx context.Context, mg cpresource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*svcapitypes.UserPool)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errUnexpectedObject)
	}
	input := GenerateUpdateUserPoolInput(cr)
	if err := e.preUpdate(ctx, cr, input); err != nil {
		return managed.ExternalUpdate{}, errors.Wrap(err, "pre-update failed")
	}
	resp, err := e.client.UpdateUserPoolWithContext(ctx, input)
	return e.postUpdate(ctx, cr, resp, managed.ExternalUpdate{}, awsclient.Wrap(err, errUpdate))
}

func (e *external) Delete(ctx context.Context, mg cpresource.Managed) error {
	cr, ok := mg.(*svcapitypes.UserPool)
	if !ok {
		return errors.New(errUnexpectedObject)
	}
	cr.Status.SetConditions(xpv1.Deleting())
	input := GenerateDeleteUserPoolInput(cr)
	ignore, err := e.preDelete(ctx, cr, input)
	if err != nil {
		return errors.Wrap(err, "pre-delete failed")
	}
	if ignore {
		return nil
	}
	resp, err := e.client.DeleteUserPoolWithContext(ctx, input)
	return e.postDelete(ctx, cr, resp, awsclient.Wrap(cpresource.Ignore(IsNotFound, err), errDelete))
}

type option func(*external)

func newExternal(kube client.Client, client svcsdkapi.CognitoIdentityProviderAPI, opts []option) *external {
	e := &external{
		kube:           kube,
		client:         client,
		preObserve:     nopPreObserve,
		postObserve:    nopPostObserve,
		lateInitialize: nopLateInitialize,
		isUpToDate:     alwaysUpToDate,
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
	client         svcsdkapi.CognitoIdentityProviderAPI
	preObserve     func(context.Context, *svcapitypes.UserPool, *svcsdk.DescribeUserPoolInput) error
	postObserve    func(context.Context, *svcapitypes.UserPool, *svcsdk.DescribeUserPoolOutput, managed.ExternalObservation, error) (managed.ExternalObservation, error)
	lateInitialize func(*svcapitypes.UserPoolParameters, *svcsdk.DescribeUserPoolOutput) error
	isUpToDate     func(*svcapitypes.UserPool, *svcsdk.DescribeUserPoolOutput) (bool, error)
	preCreate      func(context.Context, *svcapitypes.UserPool, *svcsdk.CreateUserPoolInput) error
	postCreate     func(context.Context, *svcapitypes.UserPool, *svcsdk.CreateUserPoolOutput, managed.ExternalCreation, error) (managed.ExternalCreation, error)
	preDelete      func(context.Context, *svcapitypes.UserPool, *svcsdk.DeleteUserPoolInput) (bool, error)
	postDelete     func(context.Context, *svcapitypes.UserPool, *svcsdk.DeleteUserPoolOutput, error) error
	preUpdate      func(context.Context, *svcapitypes.UserPool, *svcsdk.UpdateUserPoolInput) error
	postUpdate     func(context.Context, *svcapitypes.UserPool, *svcsdk.UpdateUserPoolOutput, managed.ExternalUpdate, error) (managed.ExternalUpdate, error)
}

func nopPreObserve(context.Context, *svcapitypes.UserPool, *svcsdk.DescribeUserPoolInput) error {
	return nil
}

func nopPostObserve(_ context.Context, _ *svcapitypes.UserPool, _ *svcsdk.DescribeUserPoolOutput, obs managed.ExternalObservation, err error) (managed.ExternalObservation, error) {
	return obs, err
}
func nopLateInitialize(*svcapitypes.UserPoolParameters, *svcsdk.DescribeUserPoolOutput) error {
	return nil
}
func alwaysUpToDate(*svcapitypes.UserPool, *svcsdk.DescribeUserPoolOutput) (bool, error) {
	return true, nil
}

func nopPreCreate(context.Context, *svcapitypes.UserPool, *svcsdk.CreateUserPoolInput) error {
	return nil
}
func nopPostCreate(_ context.Context, _ *svcapitypes.UserPool, _ *svcsdk.CreateUserPoolOutput, cre managed.ExternalCreation, err error) (managed.ExternalCreation, error) {
	return cre, err
}
func nopPreDelete(context.Context, *svcapitypes.UserPool, *svcsdk.DeleteUserPoolInput) (bool, error) {
	return false, nil
}
func nopPostDelete(_ context.Context, _ *svcapitypes.UserPool, _ *svcsdk.DeleteUserPoolOutput, err error) error {
	return err
}
func nopPreUpdate(context.Context, *svcapitypes.UserPool, *svcsdk.UpdateUserPoolInput) error {
	return nil
}
func nopPostUpdate(_ context.Context, _ *svcapitypes.UserPool, _ *svcsdk.UpdateUserPoolOutput, upd managed.ExternalUpdate, err error) (managed.ExternalUpdate, error) {
	return upd, err
}
