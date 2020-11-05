// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// For a specific time period, retrieve the top N dimension keys for a metric.
func (c *Client) DescribeDimensionKeys(ctx context.Context, params *DescribeDimensionKeysInput, optFns ...func(*Options)) (*DescribeDimensionKeysOutput, error) {
	if params == nil {
		params = &DescribeDimensionKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDimensionKeys", params, optFns, addOperationDescribeDimensionKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDimensionKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDimensionKeysInput struct {

	// The date and time specifying the end of the requested time series data. The
	// value specified is exclusive - data points less than (but not equal to) EndTime
	// will be returned. The value for EndTime must be later than the value for
	// StartTime.
	//
	// This member is required.
	EndTime *time.Time

	// A specification for how to aggregate the data points from a query result. You
	// must specify a valid dimension group. Performance Insights will return all of
	// the dimensions within that group, unless you provide the names of specific
	// dimensions within that group. You can also request that Performance Insights
	// return a limited number of values for a dimension.
	//
	// This member is required.
	GroupBy *types.DimensionGroup

	// An immutable, AWS Region-unique identifier for a data source. Performance
	// Insights gathers metrics from this data source. To use an Amazon RDS instance as
	// a data source, you specify its DbiResourceId value - for example:
	// db-FAIHNTYBKTGAUSUZQYPDS2GW4A
	//
	// This member is required.
	Identifier *string

	// The name of a Performance Insights metric to be measured. Valid values for
	// Metric are:
	//
	// * db.load.avg - a scaled representation of the number of active
	// sessions for the database engine.
	//
	// * db.sampledload.avg - the raw number of
	// active sessions for the database engine.
	//
	// This member is required.
	Metric *string

	// The AWS service for which Performance Insights will return metrics. The only
	// valid value for ServiceType is: RDS
	//
	// This member is required.
	ServiceType types.ServiceType

	// The date and time specifying the beginning of the requested time series data.
	// You can't specify a StartTime that's earlier than 7 days ago. The value
	// specified is inclusive - data points equal to or greater than StartTime will be
	// returned. The value for StartTime must be earlier than the value for EndTime.
	//
	// This member is required.
	StartTime *time.Time

	// One or more filters to apply in the request. Restrictions:
	//
	// * Any number of
	// filters by the same dimension, as specified in the GroupBy or Partition
	// parameters.
	//
	// * A single filter for any other dimension in this dimension group.
	Filter map[string]*string

	// The maximum number of items to return in the response. If more items exist than
	// the specified MaxRecords value, a pagination token is included in the response
	// so that the remaining results can be retrieved.
	MaxResults *int32

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords.
	NextToken *string

	// For each dimension specified in GroupBy, specify a secondary dimension to
	// further subdivide the partition keys in the response.
	PartitionBy *types.DimensionGroup

	// The granularity, in seconds, of the data points returned from Performance
	// Insights. A period can be as short as one second, or as long as one day (86400
	// seconds). Valid values are:
	//
	// * 1 (one second)
	//
	// * 60 (one minute)
	//
	// * 300 (five
	// minutes)
	//
	// * 3600 (one hour)
	//
	// * 86400 (twenty-four hours)
	//
	// If you don't specify
	// PeriodInSeconds, then Performance Insights will choose a value for you, with a
	// goal of returning roughly 100-200 data points in the response.
	PeriodInSeconds *int32
}

type DescribeDimensionKeysOutput struct {

	// The end time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedEndTime will be greater than
	// or equal to the value of the user-specified Endtime.
	AlignedEndTime *time.Time

	// The start time for the returned dimension keys, after alignment to a granular
	// boundary (as specified by PeriodInSeconds). AlignedStartTime will be less than
	// or equal to the value of the user-specified StartTime.
	AlignedStartTime *time.Time

	// The dimension keys that were requested.
	Keys []*types.DimensionKeyDescription

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the token, up to the
	// value specified by MaxRecords.
	NextToken *string

	// If PartitionBy was present in the request, PartitionKeys contains the breakdown
	// of dimension keys by the specified partitions.
	PartitionKeys []*types.ResponsePartitionKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDimensionKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDimensionKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDimensionKeys{}, middleware.After)
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
	if err = addOpDescribeDimensionKeysValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDimensionKeys(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDimensionKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pi",
		OperationName: "DescribeDimensionKeys",
	}
}
