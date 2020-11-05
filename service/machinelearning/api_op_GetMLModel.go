// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/machinelearning/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns an MLModel that includes detailed metadata, data source information, and
// the current status of the MLModel. GetMLModel provides results in normal or
// verbose format.
func (c *Client) GetMLModel(ctx context.Context, params *GetMLModelInput, optFns ...func(*Options)) (*GetMLModelOutput, error) {
	if params == nil {
		params = &GetMLModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLModel", params, optFns, addOperationGetMLModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLModelInput struct {

	// The ID assigned to the MLModel at creation.
	//
	// This member is required.
	MLModelId *string

	// Specifies whether the GetMLModel operation should return Recipe. If true, Recipe
	// is returned. If false, Recipe is not returned.
	Verbose *bool
}

// Represents the output of a GetMLModel operation, and provides detailed
// information about a MLModel.
type GetMLModelOutput struct {

	// The approximate CPU time in milliseconds that Amazon Machine Learning spent
	// processing the MLModel, normalized and scaled on computation resources.
	// ComputeTime is only available if the MLModel is in the COMPLETED state.
	ComputeTime *int64

	// The time that the MLModel was created. The time is expressed in epoch time.
	CreatedAt *time.Time

	// The AWS user account from which the MLModel was created. The account type can be
	// either an AWS root account or an AWS Identity and Access Management (IAM) user
	// account.
	CreatedByIamUser *string

	// The current endpoint of the MLModel
	EndpointInfo *types.RealtimeEndpointInfo

	// The epoch time when Amazon Machine Learning marked the MLModel as COMPLETED or
	// FAILED. FinishedAt is only available when the MLModel is in the COMPLETED or
	// FAILED state.
	FinishedAt *time.Time

	// The location of the data file or directory in Amazon Simple Storage Service
	// (Amazon S3).
	InputDataLocationS3 *string

	// The time of the most recent edit to the MLModel. The time is expressed in epoch
	// time.
	LastUpdatedAt *time.Time

	// A link to the file that contains logs of the CreateMLModel operation.
	LogUri *string

	// The MLModel ID, which is same as the MLModelId in the request.
	MLModelId *string

	// Identifies the MLModel category. The following are the available types:
	//
	// *
	// REGRESSION -- Produces a numeric result. For example, "What price should a house
	// be listed at?"
	//
	// * BINARY -- Produces one of two possible results. For example,
	// "Is this an e-commerce website?"
	//
	// * MULTICLASS -- Produces one of several
	// possible results. For example, "Is this a HIGH, LOW or MEDIUM risk trade?"
	MLModelType types.MLModelType

	// A description of the most recent details about accessing the MLModel.
	Message *string

	// A user-supplied name or description of the MLModel.
	Name *string

	// The recipe to use when training the MLModel. The Recipe provides detailed
	// information about the observation data to use during training, and manipulations
	// to perform on the observation data during training. Note: This parameter is
	// provided as part of the verbose format.
	Recipe *string

	// The schema used by all of the data files referenced by the DataSource. Note:
	// This parameter is provided as part of the verbose format.
	Schema *string

	// The scoring threshold is used in binary classification MLModel models. It marks
	// the boundary between a positive prediction and a negative prediction. Output
	// values greater than or equal to the threshold receive a positive result from the
	// MLModel, such as true. Output values less than the threshold receive a negative
	// response from the MLModel, such as false.
	ScoreThreshold *float32

	// The time of the most recent edit to the ScoreThreshold. The time is expressed in
	// epoch time.
	ScoreThresholdLastUpdatedAt *time.Time

	// Long integer type that is a 64-bit signed number.
	SizeInBytes *int64

	// The epoch time when Amazon Machine Learning marked the MLModel as INPROGRESS.
	// StartedAt isn't available if the MLModel is in the PENDING state.
	StartedAt *time.Time

	// The current status of the MLModel. This element can have one of the following
	// values:
	//
	// * PENDING - Amazon Machine Learning (Amazon ML) submitted a request to
	// describe a MLModel.
	//
	// * INPROGRESS - The request is processing.
	//
	// * FAILED - The
	// request did not run to completion. The ML model isn't usable.
	//
	// * COMPLETED - The
	// request completed successfully.
	//
	// * DELETED - The MLModel is marked as deleted.
	// It isn't usable.
	Status types.EntityStatus

	// The ID of the training DataSource.
	TrainingDataSourceId *string

	// A list of the training parameters in the MLModel. The list is implemented as a
	// map of key-value pairs. The following is the current set of training
	// parameters:
	//
	// * sgd.maxMLModelSizeInBytes - The maximum allowed size of the
	// model. Depending on the input data, the size of the model might affect its
	// performance. The value is an integer that ranges from 100000 to 2147483648. The
	// default value is 33554432.
	//
	// * sgd.maxPasses - The number of times that the
	// training process traverses the observations to build the MLModel. The value is
	// an integer that ranges from 1 to 100. The default value is 10.
	//
	// *
	// sgd.shuffleType - Whether Amazon ML shuffles the training data. Shuffling data
	// improves a model's ability to find the optimal solution for a variety of data
	// types. The valid values are auto and none. The default value is none. We
	// strongly recommend that you shuffle your data.
	//
	// * sgd.l1RegularizationAmount -
	// The coefficient regularization L1 norm. It controls overfitting the data by
	// penalizing large coefficients. This tends to drive coefficients to zero,
	// resulting in a sparse feature set. If you use this parameter, start by
	// specifying a small value, such as 1.0E-08. The value is a double that ranges
	// from 0 to MAX_DOUBLE. The default is to not use L1 normalization. This parameter
	// can't be used when L2 is specified. Use this parameter sparingly.
	//
	// *
	// sgd.l2RegularizationAmount - The coefficient regularization L2 norm. It controls
	// overfitting the data by penalizing large coefficients. This tends to drive
	// coefficients to small, nonzero values. If you use this parameter, start by
	// specifying a small value, such as 1.0E-08. The value is a double that ranges
	// from 0 to MAX_DOUBLE. The default is to not use L2 normalization. This parameter
	// can't be used when L1 is specified. Use this parameter sparingly.
	TrainingParameters map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMLModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMLModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMLModel{}, middleware.After)
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
	if err = addOpGetMLModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMLModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "GetMLModel",
	}
}
