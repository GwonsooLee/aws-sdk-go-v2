// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the patch baseline that should be used for the specified patch group.
func (c *Client) GetPatchBaselineForPatchGroup(ctx context.Context, params *GetPatchBaselineForPatchGroupInput, optFns ...func(*Options)) (*GetPatchBaselineForPatchGroupOutput, error) {
	if params == nil {
		params = &GetPatchBaselineForPatchGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPatchBaselineForPatchGroup", params, optFns, addOperationGetPatchBaselineForPatchGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPatchBaselineForPatchGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPatchBaselineForPatchGroupInput struct {

	// The name of the patch group whose patch baseline should be retrieved.
	//
	// This member is required.
	PatchGroup *string

	// Returns he operating system rule specified for patch groups using the patch
	// baseline.
	OperatingSystem types.OperatingSystem
}

type GetPatchBaselineForPatchGroupOutput struct {

	// The ID of the patch baseline that should be used for the patch group.
	BaselineId *string

	// The operating system rule specified for patch groups using the patch baseline.
	OperatingSystem types.OperatingSystem

	// The name of the patch group.
	PatchGroup *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPatchBaselineForPatchGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPatchBaselineForPatchGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPatchBaselineForPatchGroup{}, middleware.After)
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
	if err = addOpGetPatchBaselineForPatchGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPatchBaselineForPatchGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPatchBaselineForPatchGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetPatchBaselineForPatchGroup",
	}
}
