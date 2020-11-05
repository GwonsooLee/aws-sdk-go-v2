// Code generated by smithy-go-codegen DO NOT EDIT.

package sqs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the visibility timeout of a specified message in a queue to a new value.
// The default visibility timeout for a message is 30 seconds. The minimum is 0
// seconds. The maximum is 12 hours. For more information, see Visibility Timeout
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon Simple Queue Service Developer Guide. For example, you have a
// message with a visibility timeout of 5 minutes. After 3 minutes, you call
// ChangeMessageVisibility with a timeout of 10 minutes. You can continue to call
// ChangeMessageVisibility to extend the visibility timeout to the maximum allowed
// time. If you try to extend the visibility timeout beyond the maximum, your
// request is rejected. An Amazon SQS message has three basic states:
//
// * Sent to a
// queue by a producer.
//
// * Received from the queue by a consumer.
//
// * Deleted from
// the queue.
//
// A message is considered to be stored after it is sent to a queue by
// a producer, but not yet received from the queue by a consumer (that is, between
// states 1 and 2). There is no limit to the number of stored messages. A message
// is considered to be in flight after it is received from a queue by a consumer,
// but not yet deleted from the queue (that is, between states 2 and 3). There is a
// limit to the number of inflight messages. Limits that apply to inflight messages
// are unrelated to the unlimited number of stored messages. For most standard
// queues (depending on queue traffic and message backlog), there can be a maximum
// of approximately 120,000 inflight messages (received from a queue by a consumer,
// but not yet deleted from the queue). If you reach this limit, Amazon SQS returns
// the OverLimit error message. To avoid reaching the limit, you should delete
// messages from the queue after they're processed. You can also increase the
// number of queues you use to process your messages. To request a limit increase,
// file a support request
// (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-sqs).
// For FIFO queues, there can be a maximum of 20,000 inflight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns no error messages. If you attempt to set the
// VisibilityTimeout to a value greater than the maximum time left, Amazon SQS
// returns an error. Amazon SQS doesn't automatically recalculate and increase the
// timeout to the maximum remaining time. Unlike with a queue, when you change the
// visibility timeout for a specific message the timeout value is applied
// immediately but isn't saved in memory for that message. If you don't delete a
// message after it is received, the visibility timeout for the message reverts to
// the original timeout value (not to the value you set using the
// ChangeMessageVisibility action) the next time the message is received.
func (c *Client) ChangeMessageVisibility(ctx context.Context, params *ChangeMessageVisibilityInput, optFns ...func(*Options)) (*ChangeMessageVisibilityOutput, error) {
	if params == nil {
		params = &ChangeMessageVisibilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeMessageVisibility", params, optFns, addOperationChangeMessageVisibilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeMessageVisibilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ChangeMessageVisibilityInput struct {

	// The URL of the Amazon SQS queue whose message's visibility is changed. Queue
	// URLs and names are case-sensitive.
	//
	// This member is required.
	QueueUrl *string

	// The receipt handle associated with the message whose visibility timeout is
	// changed. This parameter is returned by the ReceiveMessage action.
	//
	// This member is required.
	ReceiptHandle *string

	// The new value for the message's visibility timeout (in seconds). Values range: 0
	// to 43200. Maximum: 12 hours.
	//
	// This member is required.
	VisibilityTimeout *int32
}

type ChangeMessageVisibilityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationChangeMessageVisibilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpChangeMessageVisibility{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpChangeMessageVisibility{}, middleware.After)
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
	if err = addOpChangeMessageVisibilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangeMessageVisibility(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opChangeMessageVisibility(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sqs",
		OperationName: "ChangeMessageVisibility",
	}
}
