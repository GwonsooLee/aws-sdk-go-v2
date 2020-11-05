// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of TrainingJobSummary objects that describe the training jobs that a
// hyperparameter tuning job launched.
func (c *Client) ListTrainingJobsForHyperParameterTuningJob(ctx context.Context, params *ListTrainingJobsForHyperParameterTuningJobInput, optFns ...func(*Options)) (*ListTrainingJobsForHyperParameterTuningJobOutput, error) {
	if params == nil {
		params = &ListTrainingJobsForHyperParameterTuningJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrainingJobsForHyperParameterTuningJob", params, optFns, addOperationListTrainingJobsForHyperParameterTuningJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrainingJobsForHyperParameterTuningJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrainingJobsForHyperParameterTuningJobInput struct {

	// The name of the tuning job whose training jobs you want to list.
	//
	// This member is required.
	HyperParameterTuningJobName *string

	// The maximum number of training jobs to return. The default value is 10.
	MaxResults *int32

	// If the result of the previous ListTrainingJobsForHyperParameterTuningJob request
	// was truncated, the response includes a NextToken. To retrieve the next set of
	// training jobs, use the token in the next request.
	NextToken *string

	// The field to sort results by. The default is Name. If the value of this field is
	// FinalObjectiveMetricValue, any training jobs that did not return an objective
	// metric are not listed.
	SortBy types.TrainingJobSortByOptions

	// The sort order for results. The default is Ascending.
	SortOrder types.SortOrder

	// A filter that returns only training jobs with the specified status.
	StatusEquals types.TrainingJobStatus
}

type ListTrainingJobsForHyperParameterTuningJobOutput struct {

	// A list of TrainingJobSummary objects that describe the training jobs that the
	// ListTrainingJobsForHyperParameterTuningJob request returned.
	//
	// This member is required.
	TrainingJobSummaries []*types.HyperParameterTrainingJobSummary

	// If the result of this ListTrainingJobsForHyperParameterTuningJob request was
	// truncated, the response includes a NextToken. To retrieve the next set of
	// training jobs, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTrainingJobsForHyperParameterTuningJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListTrainingJobsForHyperParameterTuningJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListTrainingJobsForHyperParameterTuningJob{}, middleware.After)
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
	if err = addOpListTrainingJobsForHyperParameterTuningJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrainingJobsForHyperParameterTuningJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTrainingJobsForHyperParameterTuningJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListTrainingJobsForHyperParameterTuningJob",
	}
}
