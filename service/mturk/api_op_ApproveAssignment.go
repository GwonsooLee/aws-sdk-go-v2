// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ApproveAssignment operation approves the results of a completed assignment.
// Approving an assignment initiates two payments from the Requester's Amazon.com
// account
//
// * The Worker who submitted the results is paid the reward specified in
// the HIT.
//
// * Amazon Mechanical Turk fees are debited.
//
// If the Requester's account
// does not have adequate funds for these payments, the call to ApproveAssignment
// returns an exception, and the approval is not processed. You can include an
// optional feedback message with the approval, which the Worker can see in the
// Status section of the web site. You can also call this operation for assignments
// that were previous rejected and approve them by explicitly overriding the
// previous rejection. This only works on rejected assignments that were submitted
// within the previous 30 days and only if the assignment's related HIT has not
// been deleted.
func (c *Client) ApproveAssignment(ctx context.Context, params *ApproveAssignmentInput, optFns ...func(*Options)) (*ApproveAssignmentOutput, error) {
	if params == nil {
		params = &ApproveAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApproveAssignment", params, optFns, addOperationApproveAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApproveAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApproveAssignmentInput struct {

	// The ID of the assignment. The assignment must correspond to a HIT created by the
	// Requester.
	//
	// This member is required.
	AssignmentId *string

	// A flag indicating that an assignment should be approved even if it was
	// previously rejected. Defaults to False.
	OverrideRejection *bool

	// A message for the Worker, which the Worker can see in the Status section of the
	// web site.
	RequesterFeedback *string
}

type ApproveAssignmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationApproveAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpApproveAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpApproveAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpApproveAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApproveAssignment(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opApproveAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ApproveAssignment",
	}
}
