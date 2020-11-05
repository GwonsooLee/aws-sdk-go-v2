// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns information about a transform job.
func (c *Client) DescribeTransformJob(ctx context.Context, params *DescribeTransformJobInput, optFns ...func(*Options)) (*DescribeTransformJobOutput, error) {
	if params == nil {
		params = &DescribeTransformJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransformJob", params, optFns, addOperationDescribeTransformJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransformJobInput struct {

	// The name of the transform job that you want to view details of.
	//
	// This member is required.
	TransformJobName *string
}

type DescribeTransformJobOutput struct {

	// A timestamp that shows when the transform Job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The name of the model used in the transform job.
	//
	// This member is required.
	ModelName *string

	// Describes the dataset to be transformed and the Amazon S3 location where it is
	// stored.
	//
	// This member is required.
	TransformInput *types.TransformInput

	// The Amazon Resource Name (ARN) of the transform job.
	//
	// This member is required.
	TransformJobArn *string

	// The name of the transform job.
	//
	// This member is required.
	TransformJobName *string

	// The status of the transform job. If the transform job failed, the reason is
	// returned in the FailureReason field.
	//
	// This member is required.
	TransformJobStatus types.TransformJobStatus

	// Describes the resources, including ML instance types and ML instance count, to
	// use for the transform job.
	//
	// This member is required.
	TransformResources *types.TransformResources

	// The Amazon Resource Name (ARN) of the AutoML transform job.
	AutoMLJobArn *string

	// Specifies the number of records to include in a mini-batch for an HTTP inference
	// request. A record is a single unit of input data that inference can be made on.
	// For example, a single line in a CSV file is a record. To enable the batch
	// strategy, you must set SplitType to Line, RecordIO, or TFRecord.
	BatchStrategy types.BatchStrategy

	// The data structure used to specify the data to be used for inference in a batch
	// transform job and to associate the data that is relevant to the prediction
	// results in the output. The input filter provided allows you to exclude input
	// data that is not needed for inference in a batch transform job. The output
	// filter provided allows you to include input data relevant to interpreting the
	// predictions in the output from the job. For more information, see Associate
	// Prediction Results with their Corresponding Input Records
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/batch-transform-data-processing.html).
	DataProcessing *types.DataProcessing

	// The environment variables to set in the Docker container. We support up to 16
	// key and values entries in the map.
	Environment map[string]*string

	// Associates a SageMaker job as a trial component with an experiment and trial.
	// Specified when you call the following APIs:
	//
	// * CreateProcessingJob
	//
	// *
	// CreateTrainingJob
	//
	// * CreateTransformJob
	ExperimentConfig *types.ExperimentConfig

	// If the transform job failed, FailureReason describes why it failed. A transform
	// job creates a log file, which includes error messages, and stores it as an
	// Amazon S3 object. For more information, see Log Amazon SageMaker Events with
	// Amazon CloudWatch
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/logging-cloudwatch.html).
	FailureReason *string

	// The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job
	// that created the transform or training job.
	LabelingJobArn *string

	// The maximum number of parallel requests on each instance node that can be
	// launched in a transform job. The default value is 1.
	MaxConcurrentTransforms *int32

	// The maximum payload size, in MB, used in the transform job.
	MaxPayloadInMB *int32

	// The timeout and maximum number of retries for processing a transform job
	// invocation.
	ModelClientConfig *types.ModelClientConfig

	// Indicates when the transform job has been completed, or has stopped or failed.
	// You are billed for the time interval between this time and the value of
	// TransformStartTime.
	TransformEndTime *time.Time

	// Identifies the Amazon S3 location where you want Amazon SageMaker to save the
	// results from the transform job.
	TransformOutput *types.TransformOutput

	// Indicates when the transform job starts on ML instances. You are billed for the
	// time interval between this time and the value of TransformEndTime.
	TransformStartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransformJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTransformJob{}, middleware.After)
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
	if err = addOpDescribeTransformJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransformJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTransformJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeTransformJob",
	}
}
