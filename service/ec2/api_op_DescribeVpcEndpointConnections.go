// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the VPC endpoint connections to your VPC endpoint services, including
// any endpoints that are pending your acceptance.
func (c *Client) DescribeVpcEndpointConnections(ctx context.Context, params *DescribeVpcEndpointConnectionsInput, optFns ...func(*Options)) (*DescribeVpcEndpointConnectionsOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpointConnections", params, optFns, addOperationDescribeVpcEndpointConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointConnectionsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	// * service-id - The ID of the service.
	//
	// *
	// vpc-endpoint-owner - The AWS account number of the owner of the endpoint.
	//
	// *
	// vpc-endpoint-state - The state of the endpoint (pendingAcceptance | pending |
	// available | deleting | deleted | rejected | failed).
	//
	// * vpc-endpoint-id - The ID
	// of the endpoint.
	Filters []*types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeVpcEndpointConnectionsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about one or more VPC endpoint connections.
	VpcEndpointConnections []*types.VpcEndpointConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcEndpointConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointConnections(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeVpcEndpointConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcEndpointConnections",
	}
}
