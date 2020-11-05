// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Produces a summary of an Input Security Group
func (c *Client) DescribeInputSecurityGroup(ctx context.Context, params *DescribeInputSecurityGroupInput, optFns ...func(*Options)) (*DescribeInputSecurityGroupOutput, error) {
	if params == nil {
		params = &DescribeInputSecurityGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInputSecurityGroup", params, optFns, addOperationDescribeInputSecurityGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInputSecurityGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for DescribeInputSecurityGroupRequest
type DescribeInputSecurityGroupInput struct {

	// The id of the Input Security Group to describe
	//
	// This member is required.
	InputSecurityGroupId *string
}

// Placeholder documentation for DescribeInputSecurityGroupResponse
type DescribeInputSecurityGroupOutput struct {

	// Unique ARN of Input Security Group
	Arn *string

	// The Id of the Input Security Group
	Id *string

	// The list of inputs currently using this Input Security Group.
	Inputs []*string

	// The current state of the Input Security Group.
	State types.InputSecurityGroupState

	// A collection of key-value pairs.
	Tags map[string]*string

	// Whitelist rules and their sync status
	WhitelistRules []*types.InputWhitelistRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInputSecurityGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeInputSecurityGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeInputSecurityGroup{}, middleware.After)
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
	if err = addOpDescribeInputSecurityGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInputSecurityGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInputSecurityGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "DescribeInputSecurityGroup",
	}
}
