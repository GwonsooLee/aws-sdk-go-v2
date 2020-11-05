// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets a metrics configuration (specified by the metrics configuration ID) for the
// bucket. You can have up to 1,000 metrics configurations per bucket. If you're
// updating an existing metrics configuration, note that this is a full replacement
// of the existing metrics configuration. If you don't include the elements you
// want to keep, they are erased. To use this operation, you must have permissions
// to perform the s3:PutMetricsConfiguration action. The bucket owner has this
// permission by default. The bucket owner can grant this permission to others. For
// more information about permissions, see Permissions Related to Bucket
// Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). For
// information about CloudWatch request metrics for Amazon S3, see Monitoring
// Metrics with Amazon CloudWatch
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html).
// The following operations are related to PutBucketMetricsConfiguration:
//
// *
// DeleteBucketMetricsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html)
//
// *
// PutBucketMetricsConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html)
//
// *
// ListBucketMetricsConfigurations
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html)
//
// GetBucketLifecycle
// has the following special error:
//
// * Error code: TooManyConfigurations
//
// *
// Description: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// * HTTP Status Code: HTTP 400 Bad Request
func (c *Client) PutBucketMetricsConfiguration(ctx context.Context, params *PutBucketMetricsConfigurationInput, optFns ...func(*Options)) (*PutBucketMetricsConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketMetricsConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketMetricsConfiguration", params, optFns, addOperationPutBucketMetricsConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketMetricsConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketMetricsConfigurationInput struct {

	// The name of the bucket for which the metrics configuration is set.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the metrics configuration.
	//
	// This member is required.
	Id *string

	// Specifies the metrics configuration.
	//
	// This member is required.
	MetricsConfiguration *types.MetricsConfiguration

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutBucketMetricsConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketMetricsConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketMetricsConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketMetricsConfiguration{}, middleware.After)
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
	if err = addOpPutBucketMetricsConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketMetricsConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketMetricsConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketMetricsConfiguration",
	}
}
