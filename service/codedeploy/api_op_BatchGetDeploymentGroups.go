// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about one or more deployment groups.
func (c *Client) BatchGetDeploymentGroups(ctx context.Context, params *BatchGetDeploymentGroupsInput, optFns ...func(*Options)) (*BatchGetDeploymentGroupsOutput, error) {
	if params == nil {
		params = &BatchGetDeploymentGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetDeploymentGroups", params, optFns, addOperationBatchGetDeploymentGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetDeploymentGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetDeploymentGroups operation.
type BatchGetDeploymentGroupsInput struct {

	// The name of an AWS CodeDeploy application associated with the applicable IAM
	// user or AWS account.
	//
	// This member is required.
	ApplicationName *string

	// The names of the deployment groups.
	//
	// This member is required.
	DeploymentGroupNames []*string
}

// Represents the output of a BatchGetDeploymentGroups operation.
type BatchGetDeploymentGroupsOutput struct {

	// Information about the deployment groups.
	DeploymentGroupsInfo []*types.DeploymentGroupInfo

	// Information about errors that might have occurred during the API call.
	ErrorMessage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetDeploymentGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetDeploymentGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetDeploymentGroups{}, middleware.After)
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
	if err = addOpBatchGetDeploymentGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDeploymentGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetDeploymentGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "BatchGetDeploymentGroups",
	}
}
