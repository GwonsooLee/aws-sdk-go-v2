// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists Amazon GuardDuty usage statistics over the last 30 days for the specified
// detector ID. For newly enabled detectors or data sources the cost returned will
// include only the usage so far under 30 days, this may differ from the cost
// metrics in the console, which projects usage over 30 days to provide a monthly
// cost estimate. For more information see Understanding How Usage Costs are
// Calculated
// (https://docs.aws.amazon.com/guardduty/latest/ug/monitoring_costs.html#usage-calculations).
func (c *Client) GetUsageStatistics(ctx context.Context, params *GetUsageStatisticsInput, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) {
	if params == nil {
		params = &GetUsageStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsageStatistics", params, optFns, addOperationGetUsageStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageStatisticsInput struct {

	// The ID of the detector that specifies the GuardDuty service whose usage
	// statistics you want to retrieve.
	//
	// This member is required.
	DetectorId *string

	// Represents the criteria used for querying usage.
	//
	// This member is required.
	UsageCriteria *types.UsageCriteria

	// The type of usage statistics to retrieve.
	//
	// This member is required.
	UsageStatisticType types.UsageStatisticType

	// The maximum number of results to return in the response.
	MaxResults *int32

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// The currency unit you would like to view your usage statistics in. Current valid
	// values are USD.
	Unit *string
}

type GetUsageStatisticsOutput struct {

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// The usage statistics object. If a UsageStatisticType was provided, the objects
	// representing other types will be null.
	UsageStatistics *types.UsageStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetUsageStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsageStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsageStatistics{}, middleware.After)
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
	if err = addOpGetUsageStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUsageStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetUsageStatistics",
	}
}
