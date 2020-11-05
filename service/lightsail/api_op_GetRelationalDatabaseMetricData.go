// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the data points of the specified metric for a database in Amazon
// Lightsail. Metrics report the utilization of your resources, and the error
// counts generated by them. Monitor and collect metric data regularly to maintain
// the reliability, availability, and performance of your resources.
func (c *Client) GetRelationalDatabaseMetricData(ctx context.Context, params *GetRelationalDatabaseMetricDataInput, optFns ...func(*Options)) (*GetRelationalDatabaseMetricDataOutput, error) {
	if params == nil {
		params = &GetRelationalDatabaseMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabaseMetricData", params, optFns, addOperationGetRelationalDatabaseMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabaseMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseMetricDataInput struct {

	// The end of the time interval from which to get metric data. Constraints:
	//
	// *
	// Specified in Coordinated Universal Time (UTC).
	//
	// * Specified in the Unix time
	// format. For example, if you wish to use an end time of October 1, 2018, at 8 PM
	// UTC, then you input 1538424000 as the end time.
	//
	// This member is required.
	EndTime *time.Time

	// The metric for which you want to return information. Valid relational database
	// metric names are listed below, along with the most useful statistics to include
	// in your request, and the published unit value. All relational database metric
	// data is available in 1-minute (60 seconds) granularity.
	//
	// * CPUUtilization - The
	// percentage of CPU utilization currently in use on the database. Statistics: The
	// most useful statistics are Maximum and Average. Unit: The published unit is
	// Percent.
	//
	// * DatabaseConnections - The number of database connections in use.
	// Statistics: The most useful statistics are Maximum and Sum. Unit: The published
	// unit is Count.
	//
	// * DiskQueueDepth - The number of outstanding IOs (read/write
	// requests) that are waiting to access the disk. Statistics: The most useful
	// statistic is Sum. Unit: The published unit is Count.
	//
	// * FreeStorageSpace - The
	// amount of available storage space. Statistics: The most useful statistic is Sum.
	// Unit: The published unit is Bytes.
	//
	// * NetworkReceiveThroughput - The incoming
	// (Receive) network traffic on the database, including both customer database
	// traffic and AWS traffic used for monitoring and replication. Statistics: The
	// most useful statistic is Average. Unit: The published unit is Bytes/Second.
	//
	// *
	// NetworkTransmitThroughput - The outgoing (Transmit) network traffic on the
	// database, including both customer database traffic and AWS traffic used for
	// monitoring and replication. Statistics: The most useful statistic is Average.
	// Unit: The published unit is Bytes/Second.
	//
	// This member is required.
	MetricName types.RelationalDatabaseMetricName

	// The granularity, in seconds, of the returned data points. All relational
	// database metric data is available in 1-minute (60 seconds) granularity.
	//
	// This member is required.
	Period *int32

	// The name of your database from which to get metric data.
	//
	// This member is required.
	RelationalDatabaseName *string

	// The start of the time interval from which to get metric data. Constraints:
	//
	// *
	// Specified in Coordinated Universal Time (UTC).
	//
	// * Specified in the Unix time
	// format. For example, if you wish to use a start time of October 1, 2018, at 8 PM
	// UTC, then you input 1538424000 as the start time.
	//
	// This member is required.
	StartTime *time.Time

	// The statistic for the metric. The following statistics are available:
	//
	// * Minimum
	// - The lowest value observed during the specified period. Use this value to
	// determine low volumes of activity for your application.
	//
	// * Maximum - The highest
	// value observed during the specified period. Use this value to determine high
	// volumes of activity for your application.
	//
	// * Sum - All values submitted for the
	// matching metric added together. You can use this statistic to determine the
	// total volume of a metric.
	//
	// * Average - The value of Sum / SampleCount during the
	// specified period. By comparing this statistic with the Minimum and Maximum
	// values, you can determine the full scope of a metric and how close the average
	// use is to the Minimum and Maximum values. This comparison helps you to know when
	// to increase or decrease your resources.
	//
	// * SampleCount - The count, or number,
	// of data points used for the statistical calculation.
	//
	// This member is required.
	Statistics []types.MetricStatistic

	// The unit for the metric data request. Valid units depend on the metric data
	// being requested. For the valid units with each available metric, see the
	// metricName parameter.
	//
	// This member is required.
	Unit types.MetricUnit
}

type GetRelationalDatabaseMetricDataOutput struct {

	// An array of objects that describe the metric data returned.
	MetricData []*types.MetricDatapoint

	// The name of the metric returned.
	MetricName types.RelationalDatabaseMetricName

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabaseMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseMetricData{}, middleware.After)
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
	if err = addOpGetRelationalDatabaseMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabaseMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseMetricData",
	}
}
