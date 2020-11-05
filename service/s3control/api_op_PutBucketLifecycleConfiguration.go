// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API action puts a lifecycle configuration to an Amazon S3 on Outposts
// bucket. To put a lifecycle configuration to an S3 bucket, see
// PutBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html)
// in the Amazon Simple Storage Service API. Creates a new lifecycle configuration
// for the Outposts bucket or replaces an existing lifecycle configuration.
// Outposts buckets can only support a lifecycle that deletes objects after a
// certain period of time. For more information, see Managing Lifecycle Permissions
// for Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html). All Amazon
// S3 on Outposts REST API requests for this action require an additional parameter
// of outpost-id to be passed with the request and an S3 on Outposts endpoint
// hostname prefix instead of s3-control. For an example of the request syntax for
// Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and
// the outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_PutBucketLifecycleConfiguration.html#API_control_PutBucketLifecycleConfiguration_Examples)
// section below. The following actions are related to
// PutBucketLifecycleConfiguration:
//
// * GetBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html)
//
// *
// DeleteBucketLifecycleConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html)
func (c *Client) PutBucketLifecycleConfiguration(ctx context.Context, params *PutBucketLifecycleConfigurationInput, optFns ...func(*Options)) (*PutBucketLifecycleConfigurationOutput, error) {
	if params == nil {
		params = &PutBucketLifecycleConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketLifecycleConfiguration", params, optFns, addOperationPutBucketLifecycleConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketLifecycleConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLifecycleConfigurationInput struct {

	// The AWS account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The name of the bucket for which to set the configuration.
	//
	// This member is required.
	Bucket *string

	// Container for lifecycle rules. You can add as many as 1,000 rules.
	LifecycleConfiguration *types.LifecycleConfiguration
}

type PutBucketLifecycleConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketLifecycleConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLifecycleConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLifecycleConfiguration{}, middleware.After)
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
	if err = addEndpointPrefix_opPutBucketLifecycleConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketLifecycleConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(options.Region), middleware.Before); err != nil {
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
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opPutBucketLifecycleConfigurationMiddleware struct {
}

func (*endpointPrefix_opPutBucketLifecycleConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutBucketLifecycleConfigurationMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*PutBucketLifecycleConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.HostPrefix = prefix.String()

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opPutBucketLifecycleConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opPutBucketLifecycleConfigurationMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opPutBucketLifecycleConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketLifecycleConfiguration",
	}
}
