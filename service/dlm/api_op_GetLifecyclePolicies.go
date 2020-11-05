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

// Gets summary information about all or the specified data lifecycle policies. To
// get complete information about a policy, use GetLifecyclePolicy.
func (c *Client) GetLifecyclePolicies(ctx context.Context, params *GetLifecyclePoliciesInput, optFns ...func(*Options)) (*GetLifecyclePoliciesOutput, error) {
	if params == nil {
		params = &GetLifecyclePoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLifecyclePolicies", params, optFns, addOperationGetLifecyclePoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLifecyclePoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLifecyclePoliciesInput struct {

	// The identifiers of the data lifecycle policies.
	PolicyIds []*string

	// The resource type.
	ResourceTypes []types.ResourceTypeValues

	// The activation state.
	State types.GettablePolicyStateValues

	// The tags to add to objects created by the policy. Tags are strings in the format
	// key=value. These user-defined tags are added in addition to the AWS-added
	// lifecycle tags.
	TagsToAdd []*string

	// The target tag for a policy. Tags are strings in the format key=value.
	TargetTags []*string
}

type GetLifecyclePoliciesOutput struct {

	// Summary information about the lifecycle policies.
	Policies []*types.LifecyclePolicySummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetLifecyclePoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLifecyclePolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLifecyclePolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLifecyclePolicies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLifecyclePolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dlm",
		OperationName: "GetLifecyclePolicies",
	}
}
