// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53domains/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about all of the operations that return an operation ID and
// that have ever been performed on domains that were registered by the current
// account.
func (c *Client) ListOperations(ctx context.Context, params *ListOperationsInput, optFns ...func(*Options)) (*ListOperationsOutput, error) {
	if params == nil {
		params = &ListOperationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOperations", params, optFns, addOperationListOperationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOperationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListOperations request includes the following elements.
type ListOperationsInput struct {

	// For an initial request for a list of operations, omit this element. If the
	// number of operations that are not yet complete is greater than the value that
	// you specified for MaxItems, you can use Marker to return additional operations.
	// Get the value of NextPageMarker from the previous response, and submit another
	// request that includes the value of NextPageMarker in the Marker element.
	Marker *string

	// Number of domains to be returned. Default: 20
	MaxItems *int32

	// An optional parameter that lets you get information about all the operations
	// that you submitted after a specified date and time. Specify the date and time in
	// Unix time format and Coordinated Universal time (UTC).
	SubmittedSince *time.Time
}

// The ListOperations response includes the following elements.
type ListOperationsOutput struct {

	// Lists summaries of the operations.
	//
	// This member is required.
	Operations []*types.OperationSummary

	// If there are more operations than you specified for MaxItems in the request,
	// submit another request and include the value of NextPageMarker in the value of
	// Marker.
	NextPageMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOperationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListOperations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListOperations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOperations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOperations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ListOperations",
	}
}
