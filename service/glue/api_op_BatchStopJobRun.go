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

// Stops one or more job runs for a specified job definition.
func (c *Client) BatchStopJobRun(ctx context.Context, params *BatchStopJobRunInput, optFns ...func(*Options)) (*BatchStopJobRunOutput, error) {
	if params == nil {
		params = &BatchStopJobRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchStopJobRun", params, optFns, addOperationBatchStopJobRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchStopJobRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchStopJobRunInput struct {

	// The name of the job definition for which to stop job runs.
	//
	// This member is required.
	JobName *string

	// A list of the JobRunIds that should be stopped for that job definition.
	//
	// This member is required.
	JobRunIds []*string
}

type BatchStopJobRunOutput struct {

	// A list of the errors that were encountered in trying to stop JobRuns, including
	// the JobRunId for which each error was encountered and details about the error.
	Errors []*types.BatchStopJobRunError

	// A list of the JobRuns that were successfully submitted for stopping.
	SuccessfulSubmissions []*types.BatchStopJobRunSuccessfulSubmission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchStopJobRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchStopJobRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchStopJobRun{}, middleware.After)
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
	if err = addOpBatchStopJobRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchStopJobRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchStopJobRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "BatchStopJobRun",
	}
}
