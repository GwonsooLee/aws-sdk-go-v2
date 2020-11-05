// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the shadow for the specified thing. For more information, see
// UpdateThingShadow
// (http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *Client) UpdateThingShadow(ctx context.Context, params *UpdateThingShadowInput, optFns ...func(*Options)) (*UpdateThingShadowOutput, error) {
	if params == nil {
		params = &UpdateThingShadowInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateThingShadow", params, optFns, addOperationUpdateThingShadowMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateThingShadowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the UpdateThingShadow operation.
type UpdateThingShadowInput struct {

	// The state information, in JSON format.
	//
	// This member is required.
	Payload []byte

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	// The name of the shadow.
	ShadowName *string
}

// The output from the UpdateThingShadow operation.
type UpdateThingShadowOutput struct {

	// The state information, in JSON format.
	Payload []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateThingShadowMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThingShadow{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThingShadow{}, middleware.After)
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
	if err = addOpUpdateThingShadowValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThingShadow(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateThingShadow(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "UpdateThingShadow",
	}
}
