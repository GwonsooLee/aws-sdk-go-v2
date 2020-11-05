// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all of the properties of a topic. Topic properties returned might differ
// based on the authorization of the user.
func (c *Client) GetTopicAttributes(ctx context.Context, params *GetTopicAttributesInput, optFns ...func(*Options)) (*GetTopicAttributesOutput, error) {
	if params == nil {
		params = &GetTopicAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTopicAttributes", params, optFns, addOperationGetTopicAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTopicAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for GetTopicAttributes action.
type GetTopicAttributesInput struct {

	// The ARN of the topic whose properties you want to get.
	//
	// This member is required.
	TopicArn *string
}

// Response for GetTopicAttributes action.
type GetTopicAttributesOutput struct {

	// A map of the topic's attributes. Attributes in this map include the
	// following:
	//
	// * DeliveryPolicy – The JSON serialization of the topic's delivery
	// policy.
	//
	// * DisplayName – The human-readable name used in the From field for
	// notifications to email and email-json endpoints.
	//
	// * Owner – The AWS account ID
	// of the topic's owner.
	//
	// * Policy – The JSON serialization of the topic's access
	// control policy.
	//
	// * SubscriptionsConfirmed – The number of confirmed
	// subscriptions for the topic.
	//
	// * SubscriptionsDeleted – The number of deleted
	// subscriptions for the topic.
	//
	// * SubscriptionsPending – The number of
	// subscriptions pending confirmation for the topic.
	//
	// * TopicArn – The topic's
	// ARN.
	//
	// * EffectiveDeliveryPolicy – The JSON serialization of the effective
	// delivery policy, taking system defaults into account.
	//
	// The following attribute
	// applies only to server-side-encryption
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html):
	//
	// *
	// KmsMasterKeyId - The ID of an AWS-managed customer master key (CMK) for Amazon
	// SNS or a custom CMK. For more information, see Key Terms
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms).
	// For more examples, see KeyId
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	// in the AWS Key Management Service API Reference.
	//
	// The following attributes apply
	// only to FIFO topics
	// (https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html):
	//
	// * FifoTopic –
	// When this is set to true, a FIFO topic is created.
	//
	// * ContentBasedDeduplication
	// – Enables content-based deduplication for FIFO topics.
	//
	// * By default,
	// ContentBasedDeduplication is set to false. If you create a FIFO topic and this
	// attribute is false, you must specify a value for the MessageDeduplicationId
	// parameter for the Publish
	// (https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	//
	// * When
	// you set ContentBasedDeduplication to true, Amazon SNS uses a SHA-256 hash to
	// generate the MessageDeduplicationId using the body of the message (but not the
	// attributes of the message). (Optional) To override the generated value, you can
	// specify a value for the the MessageDeduplicationId parameter for the Publish
	// action.
	Attributes map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTopicAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetTopicAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetTopicAttributes{}, middleware.After)
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
	if err = addOpGetTopicAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTopicAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTopicAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "GetTopicAttributes",
	}
}
