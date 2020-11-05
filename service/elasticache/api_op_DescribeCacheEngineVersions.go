// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the available cache engines and their versions.
func (c *Client) DescribeCacheEngineVersions(ctx context.Context, params *DescribeCacheEngineVersionsInput, optFns ...func(*Options)) (*DescribeCacheEngineVersionsOutput, error) {
	if params == nil {
		params = &DescribeCacheEngineVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheEngineVersions", params, optFns, addOperationDescribeCacheEngineVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheEngineVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheEngineVersions operation.
type DescribeCacheEngineVersionsInput struct {

	// The name of a specific cache parameter group family to return details for. Valid
	// values are: memcached1.4 | memcached1.5 | redis2.6 | redis2.8 | redis3.2 |
	// redis4.0 | redis5.0 | redis6.0 | Constraints:
	//
	// * Must be 1 to 255 alphanumeric
	// characters
	//
	// * First character must be a letter
	//
	// * Cannot end with a hyphen or
	// contain two consecutive hyphens
	CacheParameterGroupFamily *string

	// If true, specifies that only the default version of the specified engine or
	// engine and major version combination is to be returned.
	DefaultOnly *bool

	// The cache engine to return. Valid values: memcached | redis
	Engine *string

	// The cache engine version to return. Example: 1.4.14
	EngineVersion *string

	// An optional marker returned from a prior request. Use this marker for pagination
	// of results from this operation. If this parameter is specified, the response
	// includes only records beyond the marker, up to the value specified by
	// MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32
}

// Represents the output of a DescribeCacheEngineVersions operation.
type DescribeCacheEngineVersionsOutput struct {

	// A list of cache engine version details. Each element in the list contains
	// detailed information about one cache engine version.
	CacheEngineVersions []*types.CacheEngineVersion

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCacheEngineVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheEngineVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheEngineVersions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheEngineVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCacheEngineVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheEngineVersions",
	}
}
