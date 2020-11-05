// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the company network configuration for the fleet.
func (c *Client) UpdateCompanyNetworkConfiguration(ctx context.Context, params *UpdateCompanyNetworkConfigurationInput, optFns ...func(*Options)) (*UpdateCompanyNetworkConfigurationOutput, error) {
	if params == nil {
		params = &UpdateCompanyNetworkConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateCompanyNetworkConfiguration", params, optFns, addOperationUpdateCompanyNetworkConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateCompanyNetworkConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateCompanyNetworkConfigurationInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The security groups associated with access to the provided subnets.
	//
	// This member is required.
	SecurityGroupIds []*string

	// The subnets used for X-ENI connections from Amazon WorkLink rendering
	// containers.
	//
	// This member is required.
	SubnetIds []*string

	// The VPC with connectivity to associated websites.
	//
	// This member is required.
	VpcId *string
}

type UpdateCompanyNetworkConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateCompanyNetworkConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateCompanyNetworkConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateCompanyNetworkConfiguration{}, middleware.After)
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
	if err = addOpUpdateCompanyNetworkConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateCompanyNetworkConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateCompanyNetworkConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "UpdateCompanyNetworkConfiguration",
	}
}
