// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package pinpointemailiface provides an interface to enable mocking the Amazon Pinpoint Email Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package pinpointemailiface

import (
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
)

// PinpointEmailAPI provides an interface to enable mocking the
// pinpointemail.PinpointEmail service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Pinpoint Email Service.
//    func myFunc(svc pinpointemailiface.PinpointEmailAPI) bool {
//        // Make svc.CreateConfigurationSet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := pinpointemail.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPinpointEmailClient struct {
//        pinpointemailiface.PinpointEmailAPI
//    }
//    func (m *mockPinpointEmailClient) CreateConfigurationSet(input *pinpointemail.CreateConfigurationSetInput) (*pinpointemail.CreateConfigurationSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPinpointEmailClient{}
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
type PinpointEmailAPI interface {
	CreateConfigurationSetRequest(*pinpointemail.CreateConfigurationSetInput) pinpointemail.CreateConfigurationSetRequest

	CreateConfigurationSetEventDestinationRequest(*pinpointemail.CreateConfigurationSetEventDestinationInput) pinpointemail.CreateConfigurationSetEventDestinationRequest

	CreateDedicatedIpPoolRequest(*pinpointemail.CreateDedicatedIpPoolInput) pinpointemail.CreateDedicatedIpPoolRequest

	CreateEmailIdentityRequest(*pinpointemail.CreateEmailIdentityInput) pinpointemail.CreateEmailIdentityRequest

	DeleteConfigurationSetRequest(*pinpointemail.DeleteConfigurationSetInput) pinpointemail.DeleteConfigurationSetRequest

	DeleteConfigurationSetEventDestinationRequest(*pinpointemail.DeleteConfigurationSetEventDestinationInput) pinpointemail.DeleteConfigurationSetEventDestinationRequest

	DeleteDedicatedIpPoolRequest(*pinpointemail.DeleteDedicatedIpPoolInput) pinpointemail.DeleteDedicatedIpPoolRequest

	DeleteEmailIdentityRequest(*pinpointemail.DeleteEmailIdentityInput) pinpointemail.DeleteEmailIdentityRequest

	GetAccountRequest(*pinpointemail.GetAccountInput) pinpointemail.GetAccountRequest

	GetConfigurationSetRequest(*pinpointemail.GetConfigurationSetInput) pinpointemail.GetConfigurationSetRequest

	GetConfigurationSetEventDestinationsRequest(*pinpointemail.GetConfigurationSetEventDestinationsInput) pinpointemail.GetConfigurationSetEventDestinationsRequest

	GetDedicatedIpRequest(*pinpointemail.GetDedicatedIpInput) pinpointemail.GetDedicatedIpRequest

	GetDedicatedIpsRequest(*pinpointemail.GetDedicatedIpsInput) pinpointemail.GetDedicatedIpsRequest

	GetEmailIdentityRequest(*pinpointemail.GetEmailIdentityInput) pinpointemail.GetEmailIdentityRequest

	ListConfigurationSetsRequest(*pinpointemail.ListConfigurationSetsInput) pinpointemail.ListConfigurationSetsRequest

	ListDedicatedIpPoolsRequest(*pinpointemail.ListDedicatedIpPoolsInput) pinpointemail.ListDedicatedIpPoolsRequest

	ListEmailIdentitiesRequest(*pinpointemail.ListEmailIdentitiesInput) pinpointemail.ListEmailIdentitiesRequest

	PutAccountDedicatedIpWarmupAttributesRequest(*pinpointemail.PutAccountDedicatedIpWarmupAttributesInput) pinpointemail.PutAccountDedicatedIpWarmupAttributesRequest

	PutAccountSendingAttributesRequest(*pinpointemail.PutAccountSendingAttributesInput) pinpointemail.PutAccountSendingAttributesRequest

	PutConfigurationSetDeliveryOptionsRequest(*pinpointemail.PutConfigurationSetDeliveryOptionsInput) pinpointemail.PutConfigurationSetDeliveryOptionsRequest

	PutConfigurationSetReputationOptionsRequest(*pinpointemail.PutConfigurationSetReputationOptionsInput) pinpointemail.PutConfigurationSetReputationOptionsRequest

	PutConfigurationSetSendingOptionsRequest(*pinpointemail.PutConfigurationSetSendingOptionsInput) pinpointemail.PutConfigurationSetSendingOptionsRequest

	PutConfigurationSetTrackingOptionsRequest(*pinpointemail.PutConfigurationSetTrackingOptionsInput) pinpointemail.PutConfigurationSetTrackingOptionsRequest

	PutDedicatedIpInPoolRequest(*pinpointemail.PutDedicatedIpInPoolInput) pinpointemail.PutDedicatedIpInPoolRequest

	PutDedicatedIpWarmupAttributesRequest(*pinpointemail.PutDedicatedIpWarmupAttributesInput) pinpointemail.PutDedicatedIpWarmupAttributesRequest

	PutEmailIdentityDkimAttributesRequest(*pinpointemail.PutEmailIdentityDkimAttributesInput) pinpointemail.PutEmailIdentityDkimAttributesRequest

	PutEmailIdentityFeedbackAttributesRequest(*pinpointemail.PutEmailIdentityFeedbackAttributesInput) pinpointemail.PutEmailIdentityFeedbackAttributesRequest

	PutEmailIdentityMailFromAttributesRequest(*pinpointemail.PutEmailIdentityMailFromAttributesInput) pinpointemail.PutEmailIdentityMailFromAttributesRequest

	SendEmailRequest(*pinpointemail.SendEmailInput) pinpointemail.SendEmailRequest

	UpdateConfigurationSetEventDestinationRequest(*pinpointemail.UpdateConfigurationSetEventDestinationInput) pinpointemail.UpdateConfigurationSetEventDestinationRequest
}

var _ PinpointEmailAPI = (*pinpointemail.PinpointEmail)(nil)