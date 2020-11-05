// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a listener to process inbound connections from clients to an accelerator.
// Connections arrive to assigned static IP addresses on a port, port range, or
// list of port ranges that you specify. To see an AWS CLI example of creating a
// listener, scroll down to Example.
func (c *Client) CreateListener(ctx context.Context, params *CreateListenerInput, optFns ...func(*Options)) (*CreateListenerOutput, error) {
	if params == nil {
		params = &CreateListenerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateListener", params, optFns, addOperationCreateListenerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateListenerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateListenerInput struct {

	// The Amazon Resource Name (ARN) of your accelerator.
	//
	// This member is required.
	AcceleratorArn *string

	// A unique, case-sensitive identifier that you provide to ensure the
	// idempotency—that is, the uniqueness—of the request.
	//
	// This member is required.
	IdempotencyToken *string

	// The list of port ranges to support for connections from clients to your
	// accelerator.
	//
	// This member is required.
	PortRanges []*types.PortRange

	// The protocol for connections from clients to your accelerator.
	//
	// This member is required.
	Protocol types.Protocol

	// Client affinity lets you direct all requests from a user to the same endpoint,
	// if you have stateful applications, regardless of the port and protocol of the
	// client request. Client affinity gives you control over whether to always route
	// each client to the same specific endpoint. AWS Global Accelerator uses a
	// consistent-flow hashing algorithm to choose the optimal endpoint for a
	// connection. If client affinity is NONE, Global Accelerator uses the "five-tuple"
	// (5-tuple) properties—source IP address, source port, destination IP address,
	// destination port, and protocol—to select the hash value, and then chooses the
	// best endpoint. However, with this setting, if someone uses different ports to
	// connect to Global Accelerator, their connections might not be always routed to
	// the same endpoint because the hash value changes. If you want a given client to
	// always be routed to the same endpoint, set client affinity to SOURCE_IP instead.
	// When you use the SOURCE_IP setting, Global Accelerator uses the "two-tuple"
	// (2-tuple) properties— source (client) IP address and destination IP address—to
	// select the hash value. The default value is NONE.
	ClientAffinity types.ClientAffinity
}

type CreateListenerOutput struct {

	// The listener that you've created.
	Listener *types.Listener

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateListenerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateListener{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateListener{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateListenerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateListenerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateListener(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateListener struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateListener) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateListener) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateListenerInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateListenerInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateListenerMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateListener{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateListener(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "CreateListener",
	}
}
