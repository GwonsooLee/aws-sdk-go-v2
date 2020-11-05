// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays details about an import virtual machine or import snapshot tasks that
// are already created.
func (c *Client) DescribeImportImageTasks(ctx context.Context, params *DescribeImportImageTasksInput, optFns ...func(*Options)) (*DescribeImportImageTasksOutput, error) {
	if params == nil {
		params = &DescribeImportImageTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImportImageTasks", params, optFns, addOperationDescribeImportImageTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImportImageTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImportImageTasksInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Filter tasks using the task-state filter and one of the following values:
	// active, completed, deleting, or deleted.
	Filters []*types.Filter

	// The IDs of the import image tasks.
	ImportTaskIds []*string

	// The maximum number of results to return in a single call.
	MaxResults *int32

	// A token that indicates the next page of results.
	NextToken *string
}

type DescribeImportImageTasksOutput struct {

	// A list of zero or more import image tasks that are currently active or were
	// completed or canceled in the previous 7 days.
	ImportImageTasks []*types.ImportImageTask

	// The token to use to get the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeImportImageTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeImportImageTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImportImageTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImportImageTasks(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeImportImageTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImportImageTasks",
	}
}
