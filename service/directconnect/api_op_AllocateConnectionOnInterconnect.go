// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Deprecated. Use AllocateHostedConnection instead. Creates a hosted connection on
// an interconnect. Allocates a VLAN number and a specified amount of bandwidth for
// use by a hosted connection on the specified interconnect. Intended for use by
// AWS Direct Connect Partners only.
func (c *Client) AllocateConnectionOnInterconnect(ctx context.Context, params *AllocateConnectionOnInterconnectInput, optFns ...func(*Options)) (*AllocateConnectionOnInterconnectOutput, error) {
	if params == nil {
		params = &AllocateConnectionOnInterconnectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllocateConnectionOnInterconnect", params, optFns, addOperationAllocateConnectionOnInterconnectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllocateConnectionOnInterconnectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllocateConnectionOnInterconnectInput struct {

	// The bandwidth of the connection. The possible values are 50Mbps, 100Mbps,
	// 200Mbps, 300Mbps, 400Mbps, 500Mbps, 1Gbps, 2Gbps, 5Gbps, and 10Gbps. Note that
	// only those AWS Direct Connect Partners who have met specific requirements are
	// allowed to create a 1Gbps, 2Gbps, 5Gbps or 10Gbps hosted connection.
	//
	// This member is required.
	Bandwidth *string

	// The name of the provisioned connection.
	//
	// This member is required.
	ConnectionName *string

	// The ID of the interconnect on which the connection will be provisioned.
	//
	// This member is required.
	InterconnectId *string

	// The ID of the AWS account of the customer for whom the connection will be
	// provisioned.
	//
	// This member is required.
	OwnerAccount *string

	// The dedicated VLAN provisioned to the connection.
	//
	// This member is required.
	Vlan *int32
}

// Information about an AWS Direct Connect connection.
type AllocateConnectionOnInterconnectOutput struct {

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDevice *string

	// The Direct Connect endpoint on which the physical connection terminates.
	AwsDeviceV2 *string

	// The bandwidth of the connection.
	Bandwidth *string

	// The ID of the connection.
	ConnectionId *string

	// The name of the connection.
	ConnectionName *string

	// The state of the connection. The following are the possible values:
	//
	// * ordering:
	// The initial state of a hosted connection provisioned on an interconnect. The
	// connection stays in the ordering state until the owner of the hosted connection
	// confirms or declines the connection order.
	//
	// * requested: The initial state of a
	// standard connection. The connection stays in the requested state until the
	// Letter of Authorization (LOA) is sent to the customer.
	//
	// * pending: The
	// connection has been approved and is being initialized.
	//
	// * available: The network
	// link is up and the connection is ready for use.
	//
	// * down: The network link is
	// down.
	//
	// * deleting: The connection is being deleted.
	//
	// * deleted: The connection
	// has been deleted.
	//
	// * rejected: A hosted connection in the ordering state enters
	// the rejected state if it is deleted by the customer.
	//
	// * unknown: The state of
	// the connection is not available.
	ConnectionState types.ConnectionState

	// Indicates whether the connection supports a secondary BGP peer in the same
	// address family (IPv4/IPv6).
	HasLogicalRedundancy types.HasLogicalRedundancy

	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool

	// The ID of the LAG.
	LagId *string

	// The time of the most recent call to DescribeLoa for this connection.
	LoaIssueTime *time.Time

	// The location of the connection.
	Location *string

	// The ID of the AWS account that owns the connection.
	OwnerAccount *string

	// The name of the AWS Direct Connect service provider associated with the
	// connection.
	PartnerName *string

	// The name of the service provider associated with the connection.
	ProviderName *string

	// The AWS Region where the connection is located.
	Region *string

	// The tags associated with the connection.
	Tags []*types.Tag

	// The ID of the VLAN.
	Vlan *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAllocateConnectionOnInterconnectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAllocateConnectionOnInterconnect{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAllocateConnectionOnInterconnect{}, middleware.After)
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
	if err = addOpAllocateConnectionOnInterconnectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllocateConnectionOnInterconnect(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAllocateConnectionOnInterconnect(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "AllocateConnectionOnInterconnect",
	}
}
