// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resource shares that were created by attaching a policy to a resource are
// visible only to the resource share owner, and the resource share cannot be
// modified in AWS RAM. Use this API action to promote the resource share. When you
// promote the resource share, it becomes:
//
// * Visible to all principals that it is
// shared with.
//
// * Modifiable in AWS RAM.
func (c *Client) PromoteResourceShareCreatedFromPolicy(ctx context.Context, params *PromoteResourceShareCreatedFromPolicyInput, optFns ...func(*Options)) (*PromoteResourceShareCreatedFromPolicyOutput, error) {
	if params == nil {
		params = &PromoteResourceShareCreatedFromPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PromoteResourceShareCreatedFromPolicy", params, optFns, addOperationPromoteResourceShareCreatedFromPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PromoteResourceShareCreatedFromPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PromoteResourceShareCreatedFromPolicyInput struct {

	// The ARN of the resource share to promote.
	//
	// This member is required.
	ResourceShareArn *string
}

type PromoteResourceShareCreatedFromPolicyOutput struct {

	// Indicates whether the request succeeded.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPromoteResourceShareCreatedFromPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPromoteResourceShareCreatedFromPolicy{}, middleware.After)
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
	if err = addOpPromoteResourceShareCreatedFromPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPromoteResourceShareCreatedFromPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPromoteResourceShareCreatedFromPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "PromoteResourceShareCreatedFromPolicy",
	}
}
