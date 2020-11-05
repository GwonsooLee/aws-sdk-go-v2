// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a system instance. This action validates the system instance, prepares
// the deployment-related resources. For Greengrass deployments, it updates the
// Greengrass group that is specified by the greengrassGroupName parameter. It also
// adds a file to the S3 bucket specified by the s3BucketName parameter. You need
// to call DeploySystemInstance after running this action. For Greengrass
// deployments, since this action modifies and adds resources to a Greengrass group
// and an S3 bucket on the caller's behalf, the calling identity must have write
// permissions to both the specified Greengrass group and S3 bucket. Otherwise, the
// call will fail with an authorization error. For cloud deployments, this action
// requires a flowActionsRoleArn value. This is an IAM role that has permissions to
// access AWS services, such as AWS Lambda and AWS IoT, that the flow uses when it
// executes. If the definition document doesn't specify a version of the user's
// namespace, the latest version will be used by default.
func (c *Client) CreateSystemInstance(ctx context.Context, params *CreateSystemInstanceInput, optFns ...func(*Options)) (*CreateSystemInstanceOutput, error) {
	if params == nil {
		params = &CreateSystemInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSystemInstance", params, optFns, addOperationCreateSystemInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSystemInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSystemInstanceInput struct {

	// A document that defines an entity.
	//
	// This member is required.
	Definition *types.DefinitionDocument

	// The target type of the deployment. Valid values are GREENGRASS and CLOUD.
	//
	// This member is required.
	Target types.DeploymentTarget

	// The ARN of the IAM role that AWS IoT Things Graph will assume when it executes
	// the flow. This role must have read and write access to AWS Lambda and AWS IoT
	// and any other AWS services that the flow uses when it executes. This value is
	// required if the value of the target parameter is CLOUD.
	FlowActionsRoleArn *string

	// The name of the Greengrass group where the system instance will be deployed.
	// This value is required if the value of the target parameter is GREENGRASS.
	GreengrassGroupName *string

	// An object that specifies whether cloud metrics are collected in a deployment
	// and, if so, what role is used to collect metrics.
	MetricsConfiguration *types.MetricsConfiguration

	// The name of the Amazon Simple Storage Service bucket that will be used to store
	// and deploy the system instance's resource file. This value is required if the
	// value of the target parameter is GREENGRASS.
	S3BucketName *string

	// Metadata, consisting of key-value pairs, that can be used to categorize your
	// system instances.
	Tags []*types.Tag
}

type CreateSystemInstanceOutput struct {

	// The summary object that describes the new system instance.
	Summary *types.SystemInstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSystemInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateSystemInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateSystemInstance{}, middleware.After)
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
	if err = addOpCreateSystemInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSystemInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSystemInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "CreateSystemInstance",
	}
}
