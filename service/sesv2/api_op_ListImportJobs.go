// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all of the import jobs.
func (c *Client) ListImportJobs(ctx context.Context, params *ListImportJobsInput, optFns ...func(*Options)) (*ListImportJobsOutput, error) {
	if params == nil {
		params = &ListImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListImportJobs", params, optFns, addOperationListImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list all of the import jobs for a data destination
// within the specified maximum number of import jobs.
type ListImportJobsInput struct {

	// The destination of the import job, which can be used to list import jobs that
	// have a certain ImportDestinationType.
	ImportDestinationType types.ImportDestinationType

	// A string token indicating that there might be additional import jobs available
	// to be listed. Copy this token to a subsequent call to ListImportJobs with the
	// same parameters to retrieve the next page of import jobs.
	NextToken *string

	// Maximum number of import jobs to return at once. Use this parameter to paginate
	// results. If additional import jobs exist beyond the specified limit, the
	// NextToken element is sent in the response. Use the NextToken value in subsequent
	// requests to retrieve additional addresses.
	PageSize *int32
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type ListImportJobsOutput struct {

	// A list of the import job summaries.
	ImportJobs []*types.ImportJobSummary

	// A string token indicating that there might be additional import jobs available
	// to be listed. Copy this token to a subsequent call to ListImportJobs with the
	// same parameters to retrieve the next page of import jobs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListImportJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListImportJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListImportJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListImportJobs",
	}
}
