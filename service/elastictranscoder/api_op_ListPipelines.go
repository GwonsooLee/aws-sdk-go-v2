// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListPipelines operation gets a list of the pipelines associated with the
// current AWS account.
func (c *Client) ListPipelines(ctx context.Context, params *ListPipelinesInput, optFns ...func(*Options)) (*ListPipelinesOutput, error) {
	if params == nil {
		params = &ListPipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPipelines", params, optFns, addOperationListPipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListPipelineRequest structure.
type ListPipelinesInput struct {

	// To list pipelines in chronological order by the date and time that they were
	// created, enter true. To list pipelines in reverse chronological order, enter
	// false.
	Ascending *string

	// When Elastic Transcoder returns more than one page of results, use pageToken in
	// subsequent GET requests to get each successive page of results.
	PageToken *string
}

// A list of the pipelines associated with the current AWS account.
type ListPipelinesOutput struct {

	// A value that you use to access the second and subsequent pages of results, if
	// any. When the pipelines fit on one page or when you've reached the last page of
	// results, the value of NextPageToken is null.
	NextPageToken *string

	// An array of Pipeline objects.
	Pipelines []*types.Pipeline

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPipelines{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPipelines(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "ListPipelines",
	}
}
