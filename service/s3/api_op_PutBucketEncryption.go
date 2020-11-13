// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This implementation of the PUT operation uses the encryption subresource to set
// the default encryption state of an existing bucket. This implementation of the
// PUT operation sets default encryption for a bucket using server-side encryption
// with Amazon S3-managed keys SSE-S3 or AWS KMS customer master keys (CMKs)
// (SSE-KMS). For information about the Amazon S3 default encryption feature, see
// Amazon S3 Default Bucket Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html). This
// operation requires AWS Signature Version 4. For more information, see
// Authenticating Requests (AWS Signature Version 4). To use this operation, you
// must have permissions to perform the s3:PutEncryptionConfiguration action. The
// bucket owner has this permission by default. The bucket owner can grant this
// permission to others. For more information about permissions, see Permissions
// Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide. Related Resources
//
// *
// GetBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html)
//
// *
// DeleteBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html)
func (c *Client) PutBucketEncryption(ctx context.Context, params *PutBucketEncryptionInput, optFns ...func(*Options)) (*PutBucketEncryptionOutput, error) {
	if params == nil {
		params = &PutBucketEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketEncryption", params, optFns, addOperationPutBucketEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketEncryptionInput struct {

	// Specifies default encryption for a bucket using server-side encryption with
	// Amazon S3-managed keys (SSE-S3) or customer master keys stored in AWS KMS
	// (SSE-KMS). For information about the Amazon S3 default encryption feature, see
	// Amazon S3 Default Bucket Encryption
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the
	// Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Specifies the default server-side-encryption configuration.
	//
	// This member is required.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The base64-encoded 128-bit MD5 digest of the server-side encryption
	// configuration. This parameter is auto-populated when using the command from the
	// CLI.
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutBucketEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketEncryption{}, middleware.After)
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
	if err = addOpPutBucketEncryptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketEncryption(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addPutBucketEncryptionUpdateEndpoint(stack, options); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketEncryption(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketEncryption",
	}
}

// getPutBucketEncryptionBucketMember returns a pointer to string denoting a
// provided bucket member valueand a boolean indicating if the input has a modeled
// bucket name,
func getPutBucketEncryptionBucketMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketEncryptionInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addPutBucketEncryptionUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getPutBucketEncryptionBucketMember,
			SupportsAccelerate: true,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
