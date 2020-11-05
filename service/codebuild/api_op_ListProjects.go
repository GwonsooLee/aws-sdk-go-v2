// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of build project names, with each build project name representing a
// single build project.
func (c *Client) ListProjects(ctx context.Context, params *ListProjectsInput, optFns ...func(*Options)) (*ListProjectsOutput, error) {
	if params == nil {
		params = &ListProjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProjects", params, optFns, addOperationListProjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListProjectsInput struct {

	// During a previous call, if there are more than 100 items in the list, only the
	// first 100 items are returned, along with a unique string called a nextToken. To
	// get the next batch of items in the list, call this operation again, adding the
	// next token to the call. To get all of the items in the list, keep calling this
	// operation with each subsequent next token that is returned, until no more next
	// tokens are returned.
	NextToken *string

	// The criterion to be used to list build project names. Valid values include:
	//
	// *
	// CREATED_TIME: List based on when each build project was created.
	//
	// *
	// LAST_MODIFIED_TIME: List based on when information about each build project was
	// last changed.
	//
	// * NAME: List based on each build project's name.
	//
	// Use sortOrder
	// to specify in what order to list the build project names based on the preceding
	// criteria.
	SortBy types.ProjectSortByType

	// The order in which to list build projects. Valid values include:
	//
	// * ASCENDING:
	// List in ascending order.
	//
	// * DESCENDING: List in descending order.
	//
	// Use sortBy to
	// specify the criterion to be used to list build project names.
	SortOrder types.SortOrderType
}

type ListProjectsOutput struct {

	// If there are more than 100 items in the list, only the first 100 items are
	// returned, along with a unique string called a nextToken. To get the next batch
	// of items in the list, call this operation again, adding the next token to the
	// call.
	NextToken *string

	// The list of build project names, with each build project name representing a
	// single build project.
	Projects []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListProjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListProjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListProjects{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProjects(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListProjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListProjects",
	}
}
