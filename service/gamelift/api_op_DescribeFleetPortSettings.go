// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a fleet's inbound connection permissions. Connection permissions
// specify the range of IP addresses and port settings that incoming traffic can
// use to access server processes in the fleet. Game sessions that are running on
// instances in the fleet use connections that fall in this range. To get a fleet's
// inbound connection permissions, specify the fleet's unique identifier. If
// successful, a collection of IpPermission objects is returned for the requested
// fleet ID. If the requested fleet has been deleted, the result set is empty.
// Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
// * CreateFleet
//
// * ListFleets
//
// * DeleteFleet
//
// * Describe
// fleets:
//
// * DescribeFleetAttributes
//
// * DescribeFleetCapacity
//
// *
// DescribeFleetPortSettings
//
// * DescribeFleetUtilization
//
// *
// DescribeRuntimeConfiguration
//
// * DescribeEC2InstanceLimits
//
// *
// DescribeFleetEvents
//
// * UpdateFleetAttributes
//
// * StartFleetActions or
// StopFleetActions
func (c *Client) DescribeFleetPortSettings(ctx context.Context, params *DescribeFleetPortSettingsInput, optFns ...func(*Options)) (*DescribeFleetPortSettingsOutput, error) {
	if params == nil {
		params = &DescribeFleetPortSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetPortSettings", params, optFns, addOperationDescribeFleetPortSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetPortSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeFleetPortSettingsInput struct {

	// A unique identifier for a fleet to retrieve port settings for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string
}

// Represents the returned data in response to a request operation.
type DescribeFleetPortSettingsOutput struct {

	// The port settings for the requested fleet ID.
	InboundPermissions []*types.IpPermission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeFleetPortSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeFleetPortSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeFleetPortSettings{}, middleware.After)
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
	if err = addOpDescribeFleetPortSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetPortSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleetPortSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeFleetPortSettings",
	}
}
