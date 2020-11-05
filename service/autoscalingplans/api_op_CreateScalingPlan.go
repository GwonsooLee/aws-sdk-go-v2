// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a scaling plan.
func (c *Client) CreateScalingPlan(ctx context.Context, params *CreateScalingPlanInput, optFns ...func(*Options)) (*CreateScalingPlanOutput, error) {
	if params == nil {
		params = &CreateScalingPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateScalingPlan", params, optFns, addOperationCreateScalingPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateScalingPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScalingPlanInput struct {

	// A CloudFormation stack or set of tags. You can create one scaling plan per
	// application source.
	//
	// This member is required.
	ApplicationSource *types.ApplicationSource

	// The scaling instructions.
	//
	// This member is required.
	ScalingInstructions []*types.ScalingInstruction

	// The name of the scaling plan. Names cannot contain vertical bars, colons, or
	// forward slashes.
	//
	// This member is required.
	ScalingPlanName *string
}

type CreateScalingPlanOutput struct {

	// The version number of the scaling plan. This value is always 1. Currently, you
	// cannot specify multiple scaling plan versions.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateScalingPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateScalingPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateScalingPlan{}, middleware.After)
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
	if err = addOpCreateScalingPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScalingPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateScalingPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "CreateScalingPlan",
	}
}
