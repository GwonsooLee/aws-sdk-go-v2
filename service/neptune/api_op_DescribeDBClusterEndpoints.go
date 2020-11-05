// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about endpoints for an Amazon Neptune DB cluster. This
// operation can also return information for Amazon RDS clusters and Amazon DocDB
// clusters.
func (c *Client) DescribeDBClusterEndpoints(ctx context.Context, params *DescribeDBClusterEndpointsInput, optFns ...func(*Options)) (*DescribeDBClusterEndpointsOutput, error) {
	if params == nil {
		params = &DescribeDBClusterEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterEndpoints", params, optFns, addOperationDescribeDBClusterEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBClusterEndpointsInput struct {

	// The identifier of the endpoint to describe. This parameter is stored as a
	// lowercase string.
	DBClusterEndpointIdentifier *string

	// The DB cluster identifier of the DB cluster associated with the endpoint. This
	// parameter is stored as a lowercase string.
	DBClusterIdentifier *string

	// A set of name-value pairs that define which endpoints to include in the output.
	// The filters are specified as name-value pairs, in the format
	// Name=endpoint_type,Values=endpoint_type1,endpoint_type2,.... Name can be one of:
	// db-cluster-endpoint-type, db-cluster-endpoint-custom-type,
	// db-cluster-endpoint-id, db-cluster-endpoint-status. Values for the
	// db-cluster-endpoint-type filter can be one or more of: reader, writer, custom.
	// Values for the db-cluster-endpoint-custom-type filter can be one or more of:
	// reader, any. Values for the db-cluster-endpoint-status filter can be one or more
	// of: available, creating, deleting, inactive, modifying.
	Filters []*types.Filter

	// An optional pagination token provided by a previous DescribeDBClusterEndpoints
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so you can retrieve the remaining results. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
}

type DescribeDBClusterEndpointsOutput struct {

	// Contains the details of the endpoints associated with the cluster and matching
	// any filter conditions.
	DBClusterEndpoints []*types.DBClusterEndpoint

	// An optional pagination token provided by a previous DescribeDBClusterEndpoints
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterEndpoints{}, middleware.After)
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
	if err = addOpDescribeDBClusterEndpointsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBClusterEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterEndpoints",
	}
}
