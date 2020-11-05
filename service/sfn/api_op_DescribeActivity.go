// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes an activity. This operation is eventually consistent. The results are
// best effort and may not reflect very recent updates and changes.
func (c *Client) DescribeActivity(ctx context.Context, params *DescribeActivityInput, optFns ...func(*Options)) (*DescribeActivityOutput, error) {
	if params == nil {
		params = &DescribeActivityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeActivity", params, optFns, addOperationDescribeActivityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeActivityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeActivityInput struct {

	// The Amazon Resource Name (ARN) of the activity to describe.
	//
	// This member is required.
	ActivityArn *string
}

type DescribeActivityOutput struct {

	// The Amazon Resource Name (ARN) that identifies the activity.
	//
	// This member is required.
	ActivityArn *string

	// The date the activity is created.
	//
	// This member is required.
	CreationDate *time.Time

	// The name of the activity. A name must not contain:
	//
	// * white space
	//
	// * brackets <
	// > { } [ ]
	//
	// * wildcard characters ? *
	//
	// * special characters " # % \ ^ | ~ ` $ & ,
	// ; : /
	//
	// * control characters (U+0000-001F, U+007F-009F)
	//
	// To enable logging with
	// CloudWatch Logs, the name should only contain 0-9, A-Z, a-z, - and _.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeActivityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeActivity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeActivity{}, middleware.After)
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
	if err = addOpDescribeActivityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeActivity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeActivity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeActivity",
	}
}
