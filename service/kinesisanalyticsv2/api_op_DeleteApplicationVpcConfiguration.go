// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a VPC configuration from a Kinesis Data Analytics application.
func (c *Client) DeleteApplicationVpcConfiguration(ctx context.Context, params *DeleteApplicationVpcConfigurationInput, optFns ...func(*Options)) (*DeleteApplicationVpcConfigurationOutput, error) {
	if params == nil {
		params = &DeleteApplicationVpcConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteApplicationVpcConfiguration", params, optFns, addOperationDeleteApplicationVpcConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteApplicationVpcConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteApplicationVpcConfigurationInput struct {

	// The name of an existing application.
	//
	// This member is required.
	ApplicationName *string

	// The current application version ID. You can retrieve the application version ID
	// using DescribeApplication.
	//
	// This member is required.
	CurrentApplicationVersionId *int64

	// The ID of the VPC configuration to delete.
	//
	// This member is required.
	VpcConfigurationId *string
}

type DeleteApplicationVpcConfigurationOutput struct {

	// The ARN of the Kinesis Data Analytics application.
	ApplicationARN *string

	// The updated version ID of the application.
	ApplicationVersionId *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteApplicationVpcConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteApplicationVpcConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteApplicationVpcConfiguration{}, middleware.After)
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
	if err = addOpDeleteApplicationVpcConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteApplicationVpcConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteApplicationVpcConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "DeleteApplicationVpcConfiguration",
	}
}
