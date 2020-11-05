// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancingv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified listeners or the listeners for the specified Application
// Load Balancer or Network Load Balancer. You must specify either a load balancer
// or one or more listeners. For an HTTPS or TLS listener, the output includes the
// default certificate for the listener. To describe the certificate list for the
// listener, use DescribeListenerCertificates.
func (c *Client) DescribeListeners(ctx context.Context, params *DescribeListenersInput, optFns ...func(*Options)) (*DescribeListenersOutput, error) {
	if params == nil {
		params = &DescribeListenersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeListeners", params, optFns, addOperationDescribeListenersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeListenersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeListenersInput struct {

	// The Amazon Resource Names (ARN) of the listeners.
	ListenerArns []*string

	// The Amazon Resource Name (ARN) of the load balancer.
	LoadBalancerArn *string

	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string

	// The maximum number of results to return with this call.
	PageSize *int32
}

type DescribeListenersOutput struct {

	// Information about the listeners.
	Listeners []*types.Listener

	// If there are additional results, this is the marker for the next set of results.
	// Otherwise, this is null.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeListenersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeListeners{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeListeners{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeListeners(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeListeners(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeListeners",
	}
}
