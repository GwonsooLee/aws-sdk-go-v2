// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/braket/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the devices available in Amazon Braket.
func (c *Client) GetDevice(ctx context.Context, params *GetDeviceInput, optFns ...func(*Options)) (*GetDeviceOutput, error) {
	if params == nil {
		params = &GetDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDevice", params, optFns, addOperationGetDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeviceInput struct {

	// The ARN of the device to retrieve.
	//
	// This member is required.
	DeviceArn *string
}

type GetDeviceOutput struct {

	// The ARN of the device.
	//
	// This member is required.
	DeviceArn *string

	// Details about the capabilities of the device.
	// This value conforms to the media type: application/json
	//
	// This member is required.
	DeviceCapabilities *string

	// The name of the device.
	//
	// This member is required.
	DeviceName *string

	// The status of the device.
	//
	// This member is required.
	DeviceStatus types.DeviceStatus

	// The type of the device.
	//
	// This member is required.
	DeviceType types.DeviceType

	// The name of the partner company for the device.
	//
	// This member is required.
	ProviderName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDevice{}, middleware.After)
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
	if err = addOpGetDeviceValidationMiddleware(stack); err != nil {
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
