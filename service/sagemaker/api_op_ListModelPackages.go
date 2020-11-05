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

// Lists the model packages that have been created.
func (c *Client) ListModelPackages(ctx context.Context, params *ListModelPackagesInput, optFns ...func(*Options)) (*ListModelPackagesOutput, error) {
	if params == nil {
		params = &ListModelPackagesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListModelPackages", params, optFns, addOperationListModelPackagesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListModelPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListModelPackagesInput struct {

	// A filter that returns only model packages created after the specified time
	// (timestamp).
	CreationTimeAfter *time.Time

	// A filter that returns only model packages created before the specified time
	// (timestamp).
	CreationTimeBefore *time.Time

	// The maximum number of model packages to return in the response.
	MaxResults *int32

	// A string in the model package name. This filter returns only model packages
	// whose name contains the specified string.
	NameContains *string

	// If the response to a previous ListModelPackages request was truncated, the
	// response includes a NextToken. To retrieve the next set of model packages, use
	// the token in the next request.
	NextToken *string

	// The parameter by which to sort the results. The default is CreationTime.
	SortBy types.ModelPackageSortBy

	// The sort order for the results. The default is Ascending.
	SortOrder types.SortOrder
}

type ListModelPackagesOutput struct {

	// An array of ModelPackageSummary objects, each of which lists a model package.
	//
	// This member is required.
	ModelPackageSummaryList []*types.ModelPackageSummary

	// If the response is truncated, Amazon SageMaker returns this token. To retrieve
	// the next set of model packages, use it in the subsequent request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListModelPackagesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListModelPackages{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListModelPackages{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListModelPackages(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListModelPackages(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListModelPackages",
	}
}
