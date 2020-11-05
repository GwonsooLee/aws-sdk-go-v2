// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an inline document as the entire HTTP payload.
func (c *Client) InlineDocumentAsPayload(ctx context.Context, params *InlineDocumentAsPayloadInput, optFns ...func(*Options)) (*InlineDocumentAsPayloadOutput, error) {
	if params == nil {
		params = &InlineDocumentAsPayloadInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "InlineDocumentAsPayload", params, optFns, addOperationInlineDocumentAsPayloadMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*InlineDocumentAsPayloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InlineDocumentAsPayloadInput struct {
	DocumentValue smithy.Document
}

type InlineDocumentAsPayloadOutput struct {
	DocumentValue smithy.Document

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationInlineDocumentAsPayloadMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpInlineDocumentAsPayload{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpInlineDocumentAsPayload{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opInlineDocumentAsPayload(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opInlineDocumentAsPayload(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "InlineDocumentAsPayload",
	}
}
