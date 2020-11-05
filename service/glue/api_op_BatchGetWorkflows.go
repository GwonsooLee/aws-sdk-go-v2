// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of resource metadata for a given list of workflow names. After
// calling the ListWorkflows operation, you can call this operation to access the
// data to which you have been granted permissions. This operation supports all IAM
// permissions, including permission conditions that uses tags.
func (c *Client) BatchGetWorkflows(ctx context.Context, params *BatchGetWorkflowsInput, optFns ...func(*Options)) (*BatchGetWorkflowsOutput, error) {
	if params == nil {
		params = &BatchGetWorkflowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetWorkflows", params, optFns, addOperationBatchGetWorkflowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetWorkflowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetWorkflowsInput struct {

	// A list of workflow names, which may be the names returned from the ListWorkflows
	// operation.
	//
	// This member is required.
	Names []*string

	// Specifies whether to include a graph when returning the workflow resource
	// metadata.
	IncludeGraph *bool
}

type BatchGetWorkflowsOutput struct {

	// A list of names of workflows not found.
	MissingWorkflows []*string

	// A list of workflow resource metadata.
	Workflows []*types.Workflow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchGetWorkflowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetWorkflows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetWorkflows{}, middleware.After)
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
	if err = addOpBatchGetWorkflowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetWorkflows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetWorkflows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchGetWorkflows",
	}
}
