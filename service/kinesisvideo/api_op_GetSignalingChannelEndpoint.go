// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides an endpoint for the specified signaling channel to send and receive
// messages. This API uses the SingleMasterChannelEndpointConfiguration input
// parameter, which consists of the Protocols and Role properties. Protocols is
// used to determine the communication mechanism. For example, if you specify WSS
// as the protocol, this API produces a secure websocket endpoint. If you specify
// HTTPS as the protocol, this API generates an HTTPS endpoint. Role determines the
// messaging permissions. A MASTER role results in this API generating an endpoint
// that a client can use to communicate with any of the viewers on the channel. A
// VIEWER role results in this API generating an endpoint that a client can use to
// communicate only with a MASTER.
func (c *Client) GetSignalingChannelEndpoint(ctx context.Context, params *GetSignalingChannelEndpointInput, optFns ...func(*Options)) (*GetSignalingChannelEndpointOutput, error) {
	if params == nil {
		params = &GetSignalingChannelEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSignalingChannelEndpoint", params, optFns, addOperationGetSignalingChannelEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSignalingChannelEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSignalingChannelEndpointInput struct {

	// The Amazon Resource Name (ARN) of the signalling channel for which you want to
	// get an endpoint.
	//
	// This member is required.
	ChannelARN *string

	// A structure containing the endpoint configuration for the SINGLE_MASTER channel
	// type.
	SingleMasterChannelEndpointConfiguration *types.SingleMasterChannelEndpointConfiguration
}

type GetSignalingChannelEndpointOutput struct {

	// A list of endpoints for the specified signaling channel.
	ResourceEndpointList []*types.ResourceEndpointListItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSignalingChannelEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSignalingChannelEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSignalingChannelEndpoint{}, middleware.After)
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
	if err = addOpGetSignalingChannelEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSignalingChannelEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSignalingChannelEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "GetSignalingChannelEndpoint",
	}
}
