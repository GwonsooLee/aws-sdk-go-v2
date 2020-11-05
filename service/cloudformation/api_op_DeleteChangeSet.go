// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified change set. Deleting change sets ensures that no one
// executes the wrong change set. If the call successfully completes, AWS
// CloudFormation successfully deleted the change set.
func (c *Client) DeleteChangeSet(ctx context.Context, params *DeleteChangeSetInput, optFns ...func(*Options)) (*DeleteChangeSetOutput, error) {
	if params == nil {
		params = &DeleteChangeSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChangeSet", params, optFns, addOperationDeleteChangeSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChangeSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DeleteChangeSet action.
type DeleteChangeSetInput struct {

	// The name or Amazon Resource Name (ARN) of the change set that you want to
	// delete.
	//
	// This member is required.
	ChangeSetName *string

	// If you specified the name of a change set to delete, specify the stack name or
	// ID (ARN) that is associated with it.
	StackName *string
}

// The output for the DeleteChangeSet action.
type DeleteChangeSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteChangeSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteChangeSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteChangeSet{}, middleware.After)
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
	if err = addOpDeleteChangeSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChangeSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteChangeSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudformation",
		OperationName: "DeleteChangeSet",
	}
}
