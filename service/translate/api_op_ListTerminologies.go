// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a list of custom terminologies associated with your account.
func (c *Client) ListTerminologies(ctx context.Context, params *ListTerminologiesInput, optFns ...func(*Options)) (*ListTerminologiesOutput, error) {
	if params == nil {
		params = &ListTerminologiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTerminologies", params, optFns, addOperationListTerminologiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTerminologiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTerminologiesInput struct {

	// The maximum number of custom terminologies returned per list request.
	MaxResults *int32

	// If the result of the request to ListTerminologies was truncated, include the
	// NextToken to fetch the next group of custom terminologies.
	NextToken *string
}

type ListTerminologiesOutput struct {

	// If the response to the ListTerminologies was truncated, the NextToken fetches
	// the next group of custom terminologies.
	NextToken *string

	// The properties list of the custom terminologies returned on the list request.
	TerminologyPropertiesList []*types.TerminologyProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTerminologiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTerminologies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTerminologies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTerminologies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTerminologies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "ListTerminologies",
	}
}
