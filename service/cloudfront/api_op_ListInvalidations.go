// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists invalidation batches.
func (c *Client) ListInvalidations(ctx context.Context, params *ListInvalidationsInput, optFns ...func(*Options)) (*ListInvalidationsOutput, error) {
	if params == nil {
		params = &ListInvalidationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInvalidations", params, optFns, addOperationListInvalidationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInvalidationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request to list invalidations.
type ListInvalidationsInput struct {

	// The distribution's ID.
	//
	// This member is required.
	DistributionId *string

	// Use this parameter when paginating results to indicate where to begin in your
	// list of invalidation batches. Because the results are returned in decreasing
	// order from most recent to oldest, the most recent results are on the first page,
	// the second page will contain earlier results, and so on. To get the next page of
	// results, set Marker to the value of the NextMarker from the current page's
	// response. This value is the same as the ID of the last invalidation batch on
	// that page.
	Marker *string

	// The maximum number of invalidation batches that you want in the response body.
	MaxItems *string
}

// The returned result of the corresponding request.
type ListInvalidationsOutput struct {

	// Information about invalidation batches.
	InvalidationList *types.InvalidationList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInvalidationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListInvalidations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListInvalidations{}, middleware.After)
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
	if err = addOpListInvalidationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInvalidations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListInvalidations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListInvalidations",
	}
}
