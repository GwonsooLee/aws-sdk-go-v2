// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancel an input device transfer that you have requested.
func (c *Client) CancelInputDeviceTransfer(ctx context.Context, params *CancelInputDeviceTransferInput, optFns ...func(*Options)) (*CancelInputDeviceTransferOutput, error) {
	if params == nil {
		params = &CancelInputDeviceTransferInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelInputDeviceTransfer", params, optFns, addOperationCancelInputDeviceTransferMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelInputDeviceTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for CancelInputDeviceTransferRequest
type CancelInputDeviceTransferInput struct {

	// The unique ID of the input device to cancel. For example, hd-123456789abcdef.
	//
	// This member is required.
	InputDeviceId *string
}

// Placeholder documentation for CancelInputDeviceTransferResponse
type CancelInputDeviceTransferOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelInputDeviceTransferMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelInputDeviceTransfer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelInputDeviceTransfer{}, middleware.After)
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
	if err = addOpCancelInputDeviceTransferValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCancelInputDeviceTransfer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCancelInputDeviceTransfer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "CancelInputDeviceTransfer",
	}
}
