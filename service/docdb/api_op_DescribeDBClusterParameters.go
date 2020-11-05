// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the detailed parameter list for a particular cluster parameter group.
func (c *Client) DescribeDBClusterParameters(ctx context.Context, params *DescribeDBClusterParametersInput, optFns ...func(*Options)) (*DescribeDBClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeDBClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBClusterParameters", params, optFns, addOperationDescribeDBClusterParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to DescribeDBClusterParameters.
type DescribeDBClusterParametersInput struct {

	// The name of a specific cluster parameter group to return parameter details for.
	// Constraints:
	//
	// * If provided, must match the name of an existing
	// DBClusterParameterGroup.
	//
	// This member is required.
	DBClusterParameterGroupName *string

	// This parameter is not currently supported.
	Filters []*types.Filter

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token (marker) is included in
	// the response so that the remaining results can be retrieved. Default: 100
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// A value that indicates to return only parameters for a specific source.
	// Parameter sources can be engine, service, or customer.
	Source *string
}

// Represents the output of DBClusterParameterGroup.
type DescribeDBClusterParametersOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// Provides a list of parameters for the cluster parameter group.
	Parameters []*types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDBClusterParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBClusterParameters{}, middleware.After)
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
	if err = addOpDescribeDBClusterParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBClusterParameters(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeDBClusterParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBClusterParameters",
	}
}
