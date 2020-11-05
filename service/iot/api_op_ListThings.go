// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists your things. Use the attributeName and attributeValue parameters to filter
// your things. For example, calling ListThings with attributeName=Color and
// attributeValue=Red retrieves all things in the registry that contain an
// attribute Color with the value Red. You will not be charged for calling this API
// if an Access denied error is returned. You will also not be charged if no
// attributes or pagination token was provided in request and no pagination token
// and no results were returned.
func (c *Client) ListThings(ctx context.Context, params *ListThingsInput, optFns ...func(*Options)) (*ListThingsOutput, error) {
	if params == nil {
		params = &ListThingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThings", params, optFns, addOperationListThingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the ListThings operation.
type ListThingsInput struct {

	// The attribute name used to search for things.
	AttributeName *string

	// The attribute value used to search for things.
	AttributeValue *string

	// The maximum number of results to return in this operation.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string

	// The name of the thing type used to search for things.
	ThingTypeName *string
}

// The output from the ListThings operation.
type ListThingsOutput struct {

	// The token used to get the next set of results. Will not be returned if operation
	// has returned all results.
	NextToken *string

	// The things.
	Things []*types.ThingAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThings{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListThings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThings",
	}
}
