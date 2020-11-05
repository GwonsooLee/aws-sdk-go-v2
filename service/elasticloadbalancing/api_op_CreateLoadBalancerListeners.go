// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more listeners for the specified load balancer. If a listener
// with the specified port does not already exist, it is created; otherwise, the
// properties of the new listener must match the properties of the existing
// listener. For more information, see Listeners for Your Classic Load Balancer
// (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-listener-config.html)
// in the Classic Load Balancers Guide.
func (c *Client) CreateLoadBalancerListeners(ctx context.Context, params *CreateLoadBalancerListenersInput, optFns ...func(*Options)) (*CreateLoadBalancerListenersOutput, error) {
	if params == nil {
		params = &CreateLoadBalancerListenersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoadBalancerListeners", params, optFns, addOperationCreateLoadBalancerListenersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoadBalancerListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateLoadBalancerListeners.
type CreateLoadBalancerListenersInput struct {

	// The listeners.
	//
	// This member is required.
	Listeners []*types.Listener

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string
}

// Contains the parameters for CreateLoadBalancerListener.
type CreateLoadBalancerListenersOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateLoadBalancerListenersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreateLoadBalancerListeners{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateLoadBalancerListeners{}, middleware.After)
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
	if err = addOpCreateLoadBalancerListenersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoadBalancerListeners(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoadBalancerListeners(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "CreateLoadBalancerListeners",
	}
}
