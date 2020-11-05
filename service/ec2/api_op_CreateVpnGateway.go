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

// Creates a virtual private gateway. A virtual private gateway is the endpoint on
// the VPC side of your VPN connection. You can create a virtual private gateway
// before creating the VPC itself. For more information, see AWS Site-to-Site VPN
// (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in the AWS
// Site-to-Site VPN User Guide.
func (c *Client) CreateVpnGateway(ctx context.Context, params *CreateVpnGatewayInput, optFns ...func(*Options)) (*CreateVpnGatewayOutput, error) {
	if params == nil {
		params = &CreateVpnGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVpnGateway", params, optFns, addOperationCreateVpnGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVpnGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for CreateVpnGateway.
type CreateVpnGatewayInput struct {

	// The type of VPN connection this virtual private gateway supports.
	//
	// This member is required.
	Type types.GatewayType

	// A private Autonomous System Number (ASN) for the Amazon side of a BGP session.
	// If you're using a 16-bit ASN, it must be in the 64512 to 65534 range. If you're
	// using a 32-bit ASN, it must be in the 4200000000 to 4294967294 range. Default:
	// 64512
	AmazonSideAsn *int64

	// The Availability Zone for the virtual private gateway.
	AvailabilityZone *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The tags to apply to the virtual private gateway.
	TagSpecifications []*types.TagSpecification
}

// Contains the output of CreateVpnGateway.
type CreateVpnGatewayOutput struct {

	// Information about the virtual private gateway.
	VpnGateway *types.VpnGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVpnGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpCreateVpnGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpCreateVpnGateway{}, middleware.After)
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
	if err = addOpCreateVpnGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVpnGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVpnGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateVpnGateway",
	}
}
