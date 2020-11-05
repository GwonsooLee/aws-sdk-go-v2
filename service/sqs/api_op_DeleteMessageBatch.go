// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes up to ten messages from the specified queue. This is a batch version of
// DeleteMessage. The result of the action on each message is reported individually
// in the response. Because the batch request can result in a combination of
// successful and unsuccessful actions, you should check for batch errors even when
// the call returns an HTTP status code of 200. Some actions take lists of
// parameters. These lists are specified using the param.n notation. Values of n
// are integers starting from 1. For example, a parameter list with two elements
// looks like this: &AttributeName.1=first
//     &AttributeName.2=second
func (c *Client) DeleteMessageBatch(ctx context.Context, params *DeleteMessageBatchInput, optFns ...func(*Options)) (*DeleteMessageBatchOutput, error) {
	if params == nil {
		params = &DeleteMessageBatchInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteMessageBatch", params, optFns, addOperationDeleteMessageBatchMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteMessageBatchOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteMessageBatchInput struct {

	// A list of receipt handles for the messages to be deleted.
	//
	// This member is required.
	Entries []*types.DeleteMessageBatchRequestEntry

	// The URL of the Amazon SQS queue from which messages are deleted. Queue URLs and
	// names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string
}

// For each message in the batch, the response contains a
// DeleteMessageBatchResultEntry tag if the message is deleted or a
// BatchResultErrorEntry tag if the message can't be deleted.
type DeleteMessageBatchOutput struct {

	// A list of BatchResultErrorEntry items.
	//
	// This member is required.
	Failed []*types.BatchResultErrorEntry

	// A list of DeleteMessageBatchResultEntry items.
	//
	// This member is required.
	Successful []*types.DeleteMessageBatchResultEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteMessageBatchMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteMessageBatch{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteMessageBatch{}, middleware.After)
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
	if err = addOpDeleteMessageBatchValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteMessageBatch(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteMessageBatch(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "DeleteMessageBatch",
	}
}
