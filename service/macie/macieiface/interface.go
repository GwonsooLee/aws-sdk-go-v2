// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package macieiface provides an interface to enable mocking the Amazon Macie service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package macieiface

import (
	"github.com/aws/aws-sdk-go-v2/service/macie"
)

// MacieAPI provides an interface to enable mocking the
// macie.Macie service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Macie.
//    func myFunc(svc macieiface.MacieAPI) bool {
//        // Make svc.AssociateMemberAccount request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := macie.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMacieClient struct {
//        macieiface.MacieAPI
//    }
//    func (m *mockMacieClient) AssociateMemberAccount(input *macie.AssociateMemberAccountInput) (*macie.AssociateMemberAccountOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMacieClient{}
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
type MacieAPI interface {
	AssociateMemberAccountRequest(*macie.AssociateMemberAccountInput) macie.AssociateMemberAccountRequest

	AssociateS3ResourcesRequest(*macie.AssociateS3ResourcesInput) macie.AssociateS3ResourcesRequest

	DisassociateMemberAccountRequest(*macie.DisassociateMemberAccountInput) macie.DisassociateMemberAccountRequest

	DisassociateS3ResourcesRequest(*macie.DisassociateS3ResourcesInput) macie.DisassociateS3ResourcesRequest

	ListMemberAccountsRequest(*macie.ListMemberAccountsInput) macie.ListMemberAccountsRequest

	ListS3ResourcesRequest(*macie.ListS3ResourcesInput) macie.ListS3ResourcesRequest

	UpdateS3ResourcesRequest(*macie.UpdateS3ResourcesInput) macie.UpdateS3ResourcesRequest
}

var _ MacieAPI = (*macie.Macie)(nil)