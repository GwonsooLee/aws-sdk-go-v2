// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of deployment jobs for a fleet. You can optionally provide
// filters to retrieve specific deployment jobs.
func (c *Client) ListDeploymentJobs(ctx context.Context, params *ListDeploymentJobsInput, optFns ...func(*Options)) (*ListDeploymentJobsOutput, error) {
	if params == nil {
		params = &ListDeploymentJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentJobs", params, optFns, addOperationListDeploymentJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentJobsInput struct {

	// Optional filters to limit results. The filter names status and fleetName are
	// supported. When filtering, you must use the complete value of the filtered item.
	// You can use up to three filters, but they must be for the same named item. For
	// example, if you are looking for items with the status InProgress or the status
	// Pending.
	Filters []*types.Filter

	// When this parameter is used, ListDeploymentJobs only returns maxResults results
	// in a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListDeploymentJobs request
	// with the returned nextToken value. This value can be between 1 and 200. If this
	// parameter is not used, then ListDeploymentJobs returns up to 200 results and a
	// nextToken value if applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListDeploymentJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string
}

type ListDeploymentJobsOutput struct {

	// A list of deployment jobs that meet the criteria of the request.
	DeploymentJobs []*types.DeploymentJob

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListDeploymentJobs again and assign that token to
	// the request object's nextToken parameter. If there are no remaining results, the
	// previous response object's NextToken parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeploymentJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeploymentJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeploymentJobs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentJobs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDeploymentJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListDeploymentJobs",
	}
}
