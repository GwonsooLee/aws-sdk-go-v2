// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the action executions that have occurred in a pipeline.
func (c *Client) ListActionExecutions(ctx context.Context, params *ListActionExecutionsInput, optFns ...func(*Options)) (*ListActionExecutionsOutput, error) {
	if params == nil {
		params = &ListActionExecutionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListActionExecutions", params, optFns, addOperationListActionExecutionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListActionExecutionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListActionExecutionsInput struct {

	// The name of the pipeline for which you want to list action execution history.
	//
	// This member is required.
	PipelineName *string

	// Input information used to filter action execution history.
	Filter *types.ActionExecutionFilter

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. Action
	// execution history is retained for up to 12 months, based on action execution
	// start times. Default value is 100. Detailed execution history is available for
	// executions run on or after February 21, 2019.
	MaxResults *int32

	// The token that was returned from the previous ListActionExecutions call, which
	// can be used to return the next set of action executions in the list.
	NextToken *string
}

type ListActionExecutionsOutput struct {

	// The details for a list of recent executions, such as action execution ID.
	ActionExecutionDetails []*types.ActionExecutionDetail

	// If the amount of returned information is significantly large, an identifier is
	// also returned and can be used in a subsequent ListActionExecutions call to
	// return the next set of action executions in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListActionExecutionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListActionExecutions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListActionExecutions{}, middleware.After)
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
	if err = addOpListActionExecutionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListActionExecutions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListActionExecutions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "ListActionExecutions",
	}
}
