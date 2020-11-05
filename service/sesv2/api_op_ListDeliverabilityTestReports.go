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

// Show a list of the predictive inbox placement tests that you've performed,
// regardless of their statuses. For predictive inbox placement tests that are
// complete, you can use the GetDeliverabilityTestReport operation to view the
// results.
func (c *Client) ListDeliverabilityTestReports(ctx context.Context, params *ListDeliverabilityTestReportsInput, optFns ...func(*Options)) (*ListDeliverabilityTestReportsOutput, error) {
	if params == nil {
		params = &ListDeliverabilityTestReportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeliverabilityTestReports", params, optFns, addOperationListDeliverabilityTestReportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeliverabilityTestReportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list all of the predictive inbox placement tests that you've
// performed.
type ListDeliverabilityTestReportsInput struct {

	// A token returned from a previous call to ListDeliverabilityTestReports to
	// indicate the position in the list of predictive inbox placement tests.
	NextToken *string

	// The number of results to show in a single call to ListDeliverabilityTestReports.
	// If the number of results is larger than the number you specified in this
	// parameter, then the response includes a NextToken element, which you can use to
	// obtain additional results. The value you specify has to be at least 0, and can
	// be no more than 1000.
	PageSize *int32
}

// A list of the predictive inbox placement test reports that are available for
// your account, regardless of whether or not those tests are complete.
type ListDeliverabilityTestReportsOutput struct {

	// An object that contains a lists of predictive inbox placement tests that you've
	// performed.
	//
	// This member is required.
	DeliverabilityTestReports []*types.DeliverabilityTestReport

	// A token that indicates that there are additional predictive inbox placement
	// tests to list. To view additional predictive inbox placement tests, issue
	// another request to ListDeliverabilityTestReports, and pass this token in the
	// NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeliverabilityTestReportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeliverabilityTestReports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeliverabilityTestReports{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeliverabilityTestReports(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDeliverabilityTestReports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDeliverabilityTestReports",
	}
}
