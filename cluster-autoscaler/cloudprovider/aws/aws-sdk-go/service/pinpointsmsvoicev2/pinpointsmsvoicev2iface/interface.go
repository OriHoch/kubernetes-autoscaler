// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pinpointsmsvoicev2iface provides an interface to enable mocking the Amazon Pinpoint SMS Voice V2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pinpointsmsvoicev2iface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/pinpointsmsvoicev2"
)

// PinpointSMSVoiceV2API provides an interface to enable mocking the
// pinpointsmsvoicev2.PinpointSMSVoiceV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Pinpoint SMS Voice V2.
//    func myFunc(svc pinpointsmsvoicev2iface.PinpointSMSVoiceV2API) bool {
//        // Make svc.AssociateOriginationIdentity request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := pinpointsmsvoicev2.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPinpointSMSVoiceV2Client struct {
//        pinpointsmsvoicev2iface.PinpointSMSVoiceV2API
//    }
//    func (m *mockPinpointSMSVoiceV2Client) AssociateOriginationIdentity(input *pinpointsmsvoicev2.AssociateOriginationIdentityInput) (*pinpointsmsvoicev2.AssociateOriginationIdentityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPinpointSMSVoiceV2Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type PinpointSMSVoiceV2API interface {
	AssociateOriginationIdentity(*pinpointsmsvoicev2.AssociateOriginationIdentityInput) (*pinpointsmsvoicev2.AssociateOriginationIdentityOutput, error)
	AssociateOriginationIdentityWithContext(aws.Context, *pinpointsmsvoicev2.AssociateOriginationIdentityInput, ...request.Option) (*pinpointsmsvoicev2.AssociateOriginationIdentityOutput, error)
	AssociateOriginationIdentityRequest(*pinpointsmsvoicev2.AssociateOriginationIdentityInput) (*request.Request, *pinpointsmsvoicev2.AssociateOriginationIdentityOutput)

	CreateConfigurationSet(*pinpointsmsvoicev2.CreateConfigurationSetInput) (*pinpointsmsvoicev2.CreateConfigurationSetOutput, error)
	CreateConfigurationSetWithContext(aws.Context, *pinpointsmsvoicev2.CreateConfigurationSetInput, ...request.Option) (*pinpointsmsvoicev2.CreateConfigurationSetOutput, error)
	CreateConfigurationSetRequest(*pinpointsmsvoicev2.CreateConfigurationSetInput) (*request.Request, *pinpointsmsvoicev2.CreateConfigurationSetOutput)

	CreateEventDestination(*pinpointsmsvoicev2.CreateEventDestinationInput) (*pinpointsmsvoicev2.CreateEventDestinationOutput, error)
	CreateEventDestinationWithContext(aws.Context, *pinpointsmsvoicev2.CreateEventDestinationInput, ...request.Option) (*pinpointsmsvoicev2.CreateEventDestinationOutput, error)
	CreateEventDestinationRequest(*pinpointsmsvoicev2.CreateEventDestinationInput) (*request.Request, *pinpointsmsvoicev2.CreateEventDestinationOutput)

	CreateOptOutList(*pinpointsmsvoicev2.CreateOptOutListInput) (*pinpointsmsvoicev2.CreateOptOutListOutput, error)
	CreateOptOutListWithContext(aws.Context, *pinpointsmsvoicev2.CreateOptOutListInput, ...request.Option) (*pinpointsmsvoicev2.CreateOptOutListOutput, error)
	CreateOptOutListRequest(*pinpointsmsvoicev2.CreateOptOutListInput) (*request.Request, *pinpointsmsvoicev2.CreateOptOutListOutput)

	CreatePool(*pinpointsmsvoicev2.CreatePoolInput) (*pinpointsmsvoicev2.CreatePoolOutput, error)
	CreatePoolWithContext(aws.Context, *pinpointsmsvoicev2.CreatePoolInput, ...request.Option) (*pinpointsmsvoicev2.CreatePoolOutput, error)
	CreatePoolRequest(*pinpointsmsvoicev2.CreatePoolInput) (*request.Request, *pinpointsmsvoicev2.CreatePoolOutput)

	DeleteConfigurationSet(*pinpointsmsvoicev2.DeleteConfigurationSetInput) (*pinpointsmsvoicev2.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetWithContext(aws.Context, *pinpointsmsvoicev2.DeleteConfigurationSetInput, ...request.Option) (*pinpointsmsvoicev2.DeleteConfigurationSetOutput, error)
	DeleteConfigurationSetRequest(*pinpointsmsvoicev2.DeleteConfigurationSetInput) (*request.Request, *pinpointsmsvoicev2.DeleteConfigurationSetOutput)

	DeleteDefaultMessageType(*pinpointsmsvoicev2.DeleteDefaultMessageTypeInput) (*pinpointsmsvoicev2.DeleteDefaultMessageTypeOutput, error)
	DeleteDefaultMessageTypeWithContext(aws.Context, *pinpointsmsvoicev2.DeleteDefaultMessageTypeInput, ...request.Option) (*pinpointsmsvoicev2.DeleteDefaultMessageTypeOutput, error)
	DeleteDefaultMessageTypeRequest(*pinpointsmsvoicev2.DeleteDefaultMessageTypeInput) (*request.Request, *pinpointsmsvoicev2.DeleteDefaultMessageTypeOutput)

	DeleteDefaultSenderId(*pinpointsmsvoicev2.DeleteDefaultSenderIdInput) (*pinpointsmsvoicev2.DeleteDefaultSenderIdOutput, error)
	DeleteDefaultSenderIdWithContext(aws.Context, *pinpointsmsvoicev2.DeleteDefaultSenderIdInput, ...request.Option) (*pinpointsmsvoicev2.DeleteDefaultSenderIdOutput, error)
	DeleteDefaultSenderIdRequest(*pinpointsmsvoicev2.DeleteDefaultSenderIdInput) (*request.Request, *pinpointsmsvoicev2.DeleteDefaultSenderIdOutput)

	DeleteEventDestination(*pinpointsmsvoicev2.DeleteEventDestinationInput) (*pinpointsmsvoicev2.DeleteEventDestinationOutput, error)
	DeleteEventDestinationWithContext(aws.Context, *pinpointsmsvoicev2.DeleteEventDestinationInput, ...request.Option) (*pinpointsmsvoicev2.DeleteEventDestinationOutput, error)
	DeleteEventDestinationRequest(*pinpointsmsvoicev2.DeleteEventDestinationInput) (*request.Request, *pinpointsmsvoicev2.DeleteEventDestinationOutput)

	DeleteKeyword(*pinpointsmsvoicev2.DeleteKeywordInput) (*pinpointsmsvoicev2.DeleteKeywordOutput, error)
	DeleteKeywordWithContext(aws.Context, *pinpointsmsvoicev2.DeleteKeywordInput, ...request.Option) (*pinpointsmsvoicev2.DeleteKeywordOutput, error)
	DeleteKeywordRequest(*pinpointsmsvoicev2.DeleteKeywordInput) (*request.Request, *pinpointsmsvoicev2.DeleteKeywordOutput)

	DeleteOptOutList(*pinpointsmsvoicev2.DeleteOptOutListInput) (*pinpointsmsvoicev2.DeleteOptOutListOutput, error)
	DeleteOptOutListWithContext(aws.Context, *pinpointsmsvoicev2.DeleteOptOutListInput, ...request.Option) (*pinpointsmsvoicev2.DeleteOptOutListOutput, error)
	DeleteOptOutListRequest(*pinpointsmsvoicev2.DeleteOptOutListInput) (*request.Request, *pinpointsmsvoicev2.DeleteOptOutListOutput)

	DeleteOptedOutNumber(*pinpointsmsvoicev2.DeleteOptedOutNumberInput) (*pinpointsmsvoicev2.DeleteOptedOutNumberOutput, error)
	DeleteOptedOutNumberWithContext(aws.Context, *pinpointsmsvoicev2.DeleteOptedOutNumberInput, ...request.Option) (*pinpointsmsvoicev2.DeleteOptedOutNumberOutput, error)
	DeleteOptedOutNumberRequest(*pinpointsmsvoicev2.DeleteOptedOutNumberInput) (*request.Request, *pinpointsmsvoicev2.DeleteOptedOutNumberOutput)

	DeletePool(*pinpointsmsvoicev2.DeletePoolInput) (*pinpointsmsvoicev2.DeletePoolOutput, error)
	DeletePoolWithContext(aws.Context, *pinpointsmsvoicev2.DeletePoolInput, ...request.Option) (*pinpointsmsvoicev2.DeletePoolOutput, error)
	DeletePoolRequest(*pinpointsmsvoicev2.DeletePoolInput) (*request.Request, *pinpointsmsvoicev2.DeletePoolOutput)

	DeleteTextMessageSpendLimitOverride(*pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideInput) (*pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideOutput, error)
	DeleteTextMessageSpendLimitOverrideWithContext(aws.Context, *pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideInput, ...request.Option) (*pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideOutput, error)
	DeleteTextMessageSpendLimitOverrideRequest(*pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideInput) (*request.Request, *pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideOutput)

	DeleteVoiceMessageSpendLimitOverride(*pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideInput) (*pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideOutput, error)
	DeleteVoiceMessageSpendLimitOverrideWithContext(aws.Context, *pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideInput, ...request.Option) (*pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideOutput, error)
	DeleteVoiceMessageSpendLimitOverrideRequest(*pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideInput) (*request.Request, *pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideOutput)

	DescribeAccountAttributes(*pinpointsmsvoicev2.DescribeAccountAttributesInput) (*pinpointsmsvoicev2.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeAccountAttributesInput, ...request.Option) (*pinpointsmsvoicev2.DescribeAccountAttributesOutput, error)
	DescribeAccountAttributesRequest(*pinpointsmsvoicev2.DescribeAccountAttributesInput) (*request.Request, *pinpointsmsvoicev2.DescribeAccountAttributesOutput)

	DescribeAccountAttributesPages(*pinpointsmsvoicev2.DescribeAccountAttributesInput, func(*pinpointsmsvoicev2.DescribeAccountAttributesOutput, bool) bool) error
	DescribeAccountAttributesPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeAccountAttributesInput, func(*pinpointsmsvoicev2.DescribeAccountAttributesOutput, bool) bool, ...request.Option) error

	DescribeAccountLimits(*pinpointsmsvoicev2.DescribeAccountLimitsInput) (*pinpointsmsvoicev2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsWithContext(aws.Context, *pinpointsmsvoicev2.DescribeAccountLimitsInput, ...request.Option) (*pinpointsmsvoicev2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsRequest(*pinpointsmsvoicev2.DescribeAccountLimitsInput) (*request.Request, *pinpointsmsvoicev2.DescribeAccountLimitsOutput)

	DescribeAccountLimitsPages(*pinpointsmsvoicev2.DescribeAccountLimitsInput, func(*pinpointsmsvoicev2.DescribeAccountLimitsOutput, bool) bool) error
	DescribeAccountLimitsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeAccountLimitsInput, func(*pinpointsmsvoicev2.DescribeAccountLimitsOutput, bool) bool, ...request.Option) error

	DescribeConfigurationSets(*pinpointsmsvoicev2.DescribeConfigurationSetsInput) (*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, error)
	DescribeConfigurationSetsWithContext(aws.Context, *pinpointsmsvoicev2.DescribeConfigurationSetsInput, ...request.Option) (*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, error)
	DescribeConfigurationSetsRequest(*pinpointsmsvoicev2.DescribeConfigurationSetsInput) (*request.Request, *pinpointsmsvoicev2.DescribeConfigurationSetsOutput)

	DescribeConfigurationSetsPages(*pinpointsmsvoicev2.DescribeConfigurationSetsInput, func(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, bool) bool) error
	DescribeConfigurationSetsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeConfigurationSetsInput, func(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, bool) bool, ...request.Option) error

	DescribeKeywords(*pinpointsmsvoicev2.DescribeKeywordsInput) (*pinpointsmsvoicev2.DescribeKeywordsOutput, error)
	DescribeKeywordsWithContext(aws.Context, *pinpointsmsvoicev2.DescribeKeywordsInput, ...request.Option) (*pinpointsmsvoicev2.DescribeKeywordsOutput, error)
	DescribeKeywordsRequest(*pinpointsmsvoicev2.DescribeKeywordsInput) (*request.Request, *pinpointsmsvoicev2.DescribeKeywordsOutput)

	DescribeKeywordsPages(*pinpointsmsvoicev2.DescribeKeywordsInput, func(*pinpointsmsvoicev2.DescribeKeywordsOutput, bool) bool) error
	DescribeKeywordsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeKeywordsInput, func(*pinpointsmsvoicev2.DescribeKeywordsOutput, bool) bool, ...request.Option) error

	DescribeOptOutLists(*pinpointsmsvoicev2.DescribeOptOutListsInput) (*pinpointsmsvoicev2.DescribeOptOutListsOutput, error)
	DescribeOptOutListsWithContext(aws.Context, *pinpointsmsvoicev2.DescribeOptOutListsInput, ...request.Option) (*pinpointsmsvoicev2.DescribeOptOutListsOutput, error)
	DescribeOptOutListsRequest(*pinpointsmsvoicev2.DescribeOptOutListsInput) (*request.Request, *pinpointsmsvoicev2.DescribeOptOutListsOutput)

	DescribeOptOutListsPages(*pinpointsmsvoicev2.DescribeOptOutListsInput, func(*pinpointsmsvoicev2.DescribeOptOutListsOutput, bool) bool) error
	DescribeOptOutListsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeOptOutListsInput, func(*pinpointsmsvoicev2.DescribeOptOutListsOutput, bool) bool, ...request.Option) error

	DescribeOptedOutNumbers(*pinpointsmsvoicev2.DescribeOptedOutNumbersInput) (*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, error)
	DescribeOptedOutNumbersWithContext(aws.Context, *pinpointsmsvoicev2.DescribeOptedOutNumbersInput, ...request.Option) (*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, error)
	DescribeOptedOutNumbersRequest(*pinpointsmsvoicev2.DescribeOptedOutNumbersInput) (*request.Request, *pinpointsmsvoicev2.DescribeOptedOutNumbersOutput)

	DescribeOptedOutNumbersPages(*pinpointsmsvoicev2.DescribeOptedOutNumbersInput, func(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, bool) bool) error
	DescribeOptedOutNumbersPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeOptedOutNumbersInput, func(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, bool) bool, ...request.Option) error

	DescribePhoneNumbers(*pinpointsmsvoicev2.DescribePhoneNumbersInput) (*pinpointsmsvoicev2.DescribePhoneNumbersOutput, error)
	DescribePhoneNumbersWithContext(aws.Context, *pinpointsmsvoicev2.DescribePhoneNumbersInput, ...request.Option) (*pinpointsmsvoicev2.DescribePhoneNumbersOutput, error)
	DescribePhoneNumbersRequest(*pinpointsmsvoicev2.DescribePhoneNumbersInput) (*request.Request, *pinpointsmsvoicev2.DescribePhoneNumbersOutput)

	DescribePhoneNumbersPages(*pinpointsmsvoicev2.DescribePhoneNumbersInput, func(*pinpointsmsvoicev2.DescribePhoneNumbersOutput, bool) bool) error
	DescribePhoneNumbersPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribePhoneNumbersInput, func(*pinpointsmsvoicev2.DescribePhoneNumbersOutput, bool) bool, ...request.Option) error

	DescribePools(*pinpointsmsvoicev2.DescribePoolsInput) (*pinpointsmsvoicev2.DescribePoolsOutput, error)
	DescribePoolsWithContext(aws.Context, *pinpointsmsvoicev2.DescribePoolsInput, ...request.Option) (*pinpointsmsvoicev2.DescribePoolsOutput, error)
	DescribePoolsRequest(*pinpointsmsvoicev2.DescribePoolsInput) (*request.Request, *pinpointsmsvoicev2.DescribePoolsOutput)

	DescribePoolsPages(*pinpointsmsvoicev2.DescribePoolsInput, func(*pinpointsmsvoicev2.DescribePoolsOutput, bool) bool) error
	DescribePoolsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribePoolsInput, func(*pinpointsmsvoicev2.DescribePoolsOutput, bool) bool, ...request.Option) error

	DescribeSenderIds(*pinpointsmsvoicev2.DescribeSenderIdsInput) (*pinpointsmsvoicev2.DescribeSenderIdsOutput, error)
	DescribeSenderIdsWithContext(aws.Context, *pinpointsmsvoicev2.DescribeSenderIdsInput, ...request.Option) (*pinpointsmsvoicev2.DescribeSenderIdsOutput, error)
	DescribeSenderIdsRequest(*pinpointsmsvoicev2.DescribeSenderIdsInput) (*request.Request, *pinpointsmsvoicev2.DescribeSenderIdsOutput)

	DescribeSenderIdsPages(*pinpointsmsvoicev2.DescribeSenderIdsInput, func(*pinpointsmsvoicev2.DescribeSenderIdsOutput, bool) bool) error
	DescribeSenderIdsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeSenderIdsInput, func(*pinpointsmsvoicev2.DescribeSenderIdsOutput, bool) bool, ...request.Option) error

	DescribeSpendLimits(*pinpointsmsvoicev2.DescribeSpendLimitsInput) (*pinpointsmsvoicev2.DescribeSpendLimitsOutput, error)
	DescribeSpendLimitsWithContext(aws.Context, *pinpointsmsvoicev2.DescribeSpendLimitsInput, ...request.Option) (*pinpointsmsvoicev2.DescribeSpendLimitsOutput, error)
	DescribeSpendLimitsRequest(*pinpointsmsvoicev2.DescribeSpendLimitsInput) (*request.Request, *pinpointsmsvoicev2.DescribeSpendLimitsOutput)

	DescribeSpendLimitsPages(*pinpointsmsvoicev2.DescribeSpendLimitsInput, func(*pinpointsmsvoicev2.DescribeSpendLimitsOutput, bool) bool) error
	DescribeSpendLimitsPagesWithContext(aws.Context, *pinpointsmsvoicev2.DescribeSpendLimitsInput, func(*pinpointsmsvoicev2.DescribeSpendLimitsOutput, bool) bool, ...request.Option) error

	DisassociateOriginationIdentity(*pinpointsmsvoicev2.DisassociateOriginationIdentityInput) (*pinpointsmsvoicev2.DisassociateOriginationIdentityOutput, error)
	DisassociateOriginationIdentityWithContext(aws.Context, *pinpointsmsvoicev2.DisassociateOriginationIdentityInput, ...request.Option) (*pinpointsmsvoicev2.DisassociateOriginationIdentityOutput, error)
	DisassociateOriginationIdentityRequest(*pinpointsmsvoicev2.DisassociateOriginationIdentityInput) (*request.Request, *pinpointsmsvoicev2.DisassociateOriginationIdentityOutput)

	ListPoolOriginationIdentities(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput) (*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, error)
	ListPoolOriginationIdentitiesWithContext(aws.Context, *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, ...request.Option) (*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, error)
	ListPoolOriginationIdentitiesRequest(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput) (*request.Request, *pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput)

	ListPoolOriginationIdentitiesPages(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, func(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, bool) bool) error
	ListPoolOriginationIdentitiesPagesWithContext(aws.Context, *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, func(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*pinpointsmsvoicev2.ListTagsForResourceInput) (*pinpointsmsvoicev2.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *pinpointsmsvoicev2.ListTagsForResourceInput, ...request.Option) (*pinpointsmsvoicev2.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*pinpointsmsvoicev2.ListTagsForResourceInput) (*request.Request, *pinpointsmsvoicev2.ListTagsForResourceOutput)

	PutKeyword(*pinpointsmsvoicev2.PutKeywordInput) (*pinpointsmsvoicev2.PutKeywordOutput, error)
	PutKeywordWithContext(aws.Context, *pinpointsmsvoicev2.PutKeywordInput, ...request.Option) (*pinpointsmsvoicev2.PutKeywordOutput, error)
	PutKeywordRequest(*pinpointsmsvoicev2.PutKeywordInput) (*request.Request, *pinpointsmsvoicev2.PutKeywordOutput)

	PutOptedOutNumber(*pinpointsmsvoicev2.PutOptedOutNumberInput) (*pinpointsmsvoicev2.PutOptedOutNumberOutput, error)
	PutOptedOutNumberWithContext(aws.Context, *pinpointsmsvoicev2.PutOptedOutNumberInput, ...request.Option) (*pinpointsmsvoicev2.PutOptedOutNumberOutput, error)
	PutOptedOutNumberRequest(*pinpointsmsvoicev2.PutOptedOutNumberInput) (*request.Request, *pinpointsmsvoicev2.PutOptedOutNumberOutput)

	ReleasePhoneNumber(*pinpointsmsvoicev2.ReleasePhoneNumberInput) (*pinpointsmsvoicev2.ReleasePhoneNumberOutput, error)
	ReleasePhoneNumberWithContext(aws.Context, *pinpointsmsvoicev2.ReleasePhoneNumberInput, ...request.Option) (*pinpointsmsvoicev2.ReleasePhoneNumberOutput, error)
	ReleasePhoneNumberRequest(*pinpointsmsvoicev2.ReleasePhoneNumberInput) (*request.Request, *pinpointsmsvoicev2.ReleasePhoneNumberOutput)

	RequestPhoneNumber(*pinpointsmsvoicev2.RequestPhoneNumberInput) (*pinpointsmsvoicev2.RequestPhoneNumberOutput, error)
	RequestPhoneNumberWithContext(aws.Context, *pinpointsmsvoicev2.RequestPhoneNumberInput, ...request.Option) (*pinpointsmsvoicev2.RequestPhoneNumberOutput, error)
	RequestPhoneNumberRequest(*pinpointsmsvoicev2.RequestPhoneNumberInput) (*request.Request, *pinpointsmsvoicev2.RequestPhoneNumberOutput)

	SendTextMessage(*pinpointsmsvoicev2.SendTextMessageInput) (*pinpointsmsvoicev2.SendTextMessageOutput, error)
	SendTextMessageWithContext(aws.Context, *pinpointsmsvoicev2.SendTextMessageInput, ...request.Option) (*pinpointsmsvoicev2.SendTextMessageOutput, error)
	SendTextMessageRequest(*pinpointsmsvoicev2.SendTextMessageInput) (*request.Request, *pinpointsmsvoicev2.SendTextMessageOutput)

	SendVoiceMessage(*pinpointsmsvoicev2.SendVoiceMessageInput) (*pinpointsmsvoicev2.SendVoiceMessageOutput, error)
	SendVoiceMessageWithContext(aws.Context, *pinpointsmsvoicev2.SendVoiceMessageInput, ...request.Option) (*pinpointsmsvoicev2.SendVoiceMessageOutput, error)
	SendVoiceMessageRequest(*pinpointsmsvoicev2.SendVoiceMessageInput) (*request.Request, *pinpointsmsvoicev2.SendVoiceMessageOutput)

	SetDefaultMessageType(*pinpointsmsvoicev2.SetDefaultMessageTypeInput) (*pinpointsmsvoicev2.SetDefaultMessageTypeOutput, error)
	SetDefaultMessageTypeWithContext(aws.Context, *pinpointsmsvoicev2.SetDefaultMessageTypeInput, ...request.Option) (*pinpointsmsvoicev2.SetDefaultMessageTypeOutput, error)
	SetDefaultMessageTypeRequest(*pinpointsmsvoicev2.SetDefaultMessageTypeInput) (*request.Request, *pinpointsmsvoicev2.SetDefaultMessageTypeOutput)

	SetDefaultSenderId(*pinpointsmsvoicev2.SetDefaultSenderIdInput) (*pinpointsmsvoicev2.SetDefaultSenderIdOutput, error)
	SetDefaultSenderIdWithContext(aws.Context, *pinpointsmsvoicev2.SetDefaultSenderIdInput, ...request.Option) (*pinpointsmsvoicev2.SetDefaultSenderIdOutput, error)
	SetDefaultSenderIdRequest(*pinpointsmsvoicev2.SetDefaultSenderIdInput) (*request.Request, *pinpointsmsvoicev2.SetDefaultSenderIdOutput)

	SetTextMessageSpendLimitOverride(*pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideInput) (*pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideOutput, error)
	SetTextMessageSpendLimitOverrideWithContext(aws.Context, *pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideInput, ...request.Option) (*pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideOutput, error)
	SetTextMessageSpendLimitOverrideRequest(*pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideInput) (*request.Request, *pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideOutput)

	SetVoiceMessageSpendLimitOverride(*pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideInput) (*pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideOutput, error)
	SetVoiceMessageSpendLimitOverrideWithContext(aws.Context, *pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideInput, ...request.Option) (*pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideOutput, error)
	SetVoiceMessageSpendLimitOverrideRequest(*pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideInput) (*request.Request, *pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideOutput)

	TagResource(*pinpointsmsvoicev2.TagResourceInput) (*pinpointsmsvoicev2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *pinpointsmsvoicev2.TagResourceInput, ...request.Option) (*pinpointsmsvoicev2.TagResourceOutput, error)
	TagResourceRequest(*pinpointsmsvoicev2.TagResourceInput) (*request.Request, *pinpointsmsvoicev2.TagResourceOutput)

	UntagResource(*pinpointsmsvoicev2.UntagResourceInput) (*pinpointsmsvoicev2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *pinpointsmsvoicev2.UntagResourceInput, ...request.Option) (*pinpointsmsvoicev2.UntagResourceOutput, error)
	UntagResourceRequest(*pinpointsmsvoicev2.UntagResourceInput) (*request.Request, *pinpointsmsvoicev2.UntagResourceOutput)

	UpdateEventDestination(*pinpointsmsvoicev2.UpdateEventDestinationInput) (*pinpointsmsvoicev2.UpdateEventDestinationOutput, error)
	UpdateEventDestinationWithContext(aws.Context, *pinpointsmsvoicev2.UpdateEventDestinationInput, ...request.Option) (*pinpointsmsvoicev2.UpdateEventDestinationOutput, error)
	UpdateEventDestinationRequest(*pinpointsmsvoicev2.UpdateEventDestinationInput) (*request.Request, *pinpointsmsvoicev2.UpdateEventDestinationOutput)

	UpdatePhoneNumber(*pinpointsmsvoicev2.UpdatePhoneNumberInput) (*pinpointsmsvoicev2.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberWithContext(aws.Context, *pinpointsmsvoicev2.UpdatePhoneNumberInput, ...request.Option) (*pinpointsmsvoicev2.UpdatePhoneNumberOutput, error)
	UpdatePhoneNumberRequest(*pinpointsmsvoicev2.UpdatePhoneNumberInput) (*request.Request, *pinpointsmsvoicev2.UpdatePhoneNumberOutput)

	UpdatePool(*pinpointsmsvoicev2.UpdatePoolInput) (*pinpointsmsvoicev2.UpdatePoolOutput, error)
	UpdatePoolWithContext(aws.Context, *pinpointsmsvoicev2.UpdatePoolInput, ...request.Option) (*pinpointsmsvoicev2.UpdatePoolOutput, error)
	UpdatePoolRequest(*pinpointsmsvoicev2.UpdatePoolInput) (*request.Request, *pinpointsmsvoicev2.UpdatePoolOutput)
}

var _ PinpointSMSVoiceV2API = (*pinpointsmsvoicev2.PinpointSMSVoiceV2)(nil)
