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

// Searches one or more transit gateway multicast groups and returns the group
// membership information.
func (c *Client) SearchTransitGatewayMulticastGroups(ctx context.Context, params *SearchTransitGatewayMulticastGroupsInput, optFns ...func(*Options)) (*SearchTransitGatewayMulticastGroupsOutput, error) {
	if params == nil {
		params = &SearchTransitGatewayMulticastGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchTransitGatewayMulticastGroups", params, optFns, addOperationSearchTransitGatewayMulticastGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchTransitGatewayMulticastGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchTransitGatewayMulticastGroupsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	// * group-ip-address - The IP
	// address of the transit gateway multicast group.
	//
	// * is-group-member - The
	// resource is a group member. Valid values are true | false.
	//
	// * is-group-source -
	// The resource is a group source. Valid values are true | false.
	//
	// * member-type -
	// The member type. Valid values are igmp | static.
	//
	// * resource-id - The ID of the
	// resource.
	//
	// * resource-type - The type of resource. Valid values are vpc | vpn |
	// direct-connect-gateway | tgw-peering.
	//
	// * source-type - The source type. Valid
	// values are igmp | static.
	//
	// * state - The state of the subnet association. Valid
	// values are associated | associated | disassociated | disassociating.
	//
	// *
	// subnet-id - The ID of the subnet.
	//
	// * transit-gateway-attachment-id - The id of
	// the transit gateway attachment.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string
}

type SearchTransitGatewayMulticastGroupsOutput struct {

	// Information about the transit gateway multicast group.
	MulticastGroups []*types.TransitGatewayMulticastGroup

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchTransitGatewayMulticastGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpSearchTransitGatewayMulticastGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpSearchTransitGatewayMulticastGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchTransitGatewayMulticastGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchTransitGatewayMulticastGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "SearchTransitGatewayMulticastGroups",
	}
}
