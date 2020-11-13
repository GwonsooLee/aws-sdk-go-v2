// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Updates the status for the specified job. Use this operation to confirm that you
// want to run a job or to cancel an existing job. For more information, see S3
// Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-basics.html) in the
// Amazon Simple Storage Service Developer Guide. Related actions include:
//
// *
// CreateJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//
// *
// ListJobs
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
//
// *
// DescribeJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
//
// *
// UpdateJobStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
func (c *Client) UpdateJobStatus(ctx context.Context, params *UpdateJobStatusInput, optFns ...func(*Options)) (*UpdateJobStatusOutput, error) {
	if params == nil {
		params = &UpdateJobStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobStatus", params, optFns, addOperationUpdateJobStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobStatusInput struct {

	//
	//
	// This member is required.
	AccountId *string

	// The ID of the job whose status you want to update.
	//
	// This member is required.
	JobId *string

	// The status that you want to move the specified job to.
	//
	// This member is required.
	RequestedJobStatus types.RequestedJobStatus

	// A description of the reason why you want to change the specified job's status.
	// This field can be any string up to the maximum length.
	StatusUpdateReason *string
}

type UpdateJobStatusOutput struct {

	// The ID for the job whose status was updated.
	JobId *string

	// The current status for the specified job.
	Status types.JobStatus

	// The reason that the specified job's status was updated.
	StatusUpdateReason *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateJobStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateJobStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateJobStatus{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateJobStatusMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateJobStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobStatus(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateJobStatusUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opUpdateJobStatusMiddleware struct {
}

func (*endpointPrefix_opUpdateJobStatusMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateJobStatusMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*UpdateJobStatusInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUpdateJobStatusMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateJobStatusMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateJobStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UpdateJobStatus",
	}
}

func copyUpdateJobStatusInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*UpdateJobStatusInput)
	if !ok {
		return nil, fmt.Errorf("expect *UpdateJobStatusInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillUpdateJobStatusAccountID(input interface{}, v string) error {
	in := input.(*UpdateJobStatusInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addUpdateJobStatusUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyUpdateJobStatusInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
