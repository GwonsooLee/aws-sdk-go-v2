// Code generated by smithy-go-codegen DO NOT EDIT.

package dlm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dlm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a policy to manage the lifecycle of the specified AWS resources. You can
// create up to 100 lifecycle policies.
func (c *Client) CreateLifecyclePolicy(ctx context.Context, params *CreateLifecyclePolicyInput, optFns ...func(*Options)) (*CreateLifecyclePolicyOutput, error) {
	if params == nil {
		params = &CreateLifecyclePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLifecyclePolicy", params, optFns, addOperationCreateLifecyclePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLifecyclePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLifecyclePolicyInput struct {

	// A description of the lifecycle policy. The characters ^[0-9A-Za-z _-]+$ are
	// supported.
	//
	// This member is required.
	Description *string

	// The Amazon Resource Name (ARN) of the IAM role used to run the operations
	// specified by the lifecycle policy.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The configuration details of the lifecycle policy.
	//
	// This member is required.
	PolicyDetails *types.PolicyDetails

	// The desired activation state of the lifecycle policy after creation.
	//
	// This member is required.
	State types.SettablePolicyStateValues

	// The tags to apply to the lifecycle policy during creation.
	Tags map[string]*string
}

type CreateLifecyclePolicyOutput struct {

	// The identifier of the lifecycle policy.
	PolicyId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLifecyclePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLifecyclePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLifecyclePolicy{}, middleware.After)
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
	if err = addOpCreateLifecyclePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLifecyclePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLifecyclePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dlm",
		OperationName: "CreateLifecyclePolicy",
	}
}
