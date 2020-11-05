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

// Set the logging parameters for a bucket and to specify permissions for who can
// view and modify the logging parameters. All logs are saved to buckets in the
// same AWS Region as the source bucket. To set the logging status of a bucket, you
// must be the bucket owner. The bucket owner is automatically granted FULL_CONTROL
// to all logs. You use the Grantee request element to grant access to other
// people. The Permissions request element specifies the kind of access the grantee
// has to the logs. Grantee Values You can specify the person (grantee) to whom
// you're assigning access rights (using request elements) in the following
// ways:
//
// * By the person's ID: <>ID<><>GranteesEmail<>  DisplayName is optional
// and ignored in the request.
//
// * By Email address:  <>Grantees@email.com<> The
// grantee is resolved to the CanonicalUser and, in a response to a GET Object acl
// request, appears as the CanonicalUser.
//
// * By URI:
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// To enable
// logging, you use LoggingEnabled and its children request elements. To disable
// logging, you use an empty BucketLoggingStatus request element:  For more
// information about server access logging, see Server Access Logging
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerLogs.html). For more
// information about creating a bucket, see CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html). For
// more information about returning the logging status of a bucket, see
// GetBucketLogging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html). The
// following operations are related to PutBucketLogging:
//
// * PutObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html)
//
// *
// DeleteBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)
//
// *
// CreateBucket
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html)
//
// *
// GetBucketLogging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html)
func (c *Client) PutBucketLogging(ctx context.Context, params *PutBucketLoggingInput, optFns ...func(*Options)) (*PutBucketLoggingOutput, error) {
	if params == nil {
		params = &PutBucketLoggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketLogging", params, optFns, addOperationPutBucketLoggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketLoggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketLoggingInput struct {

	// The name of the bucket for which to set the logging parameters.
	//
	// This member is required.
	Bucket *string

	// Container for logging status information.
	//
	// This member is required.
	BucketLoggingStatus *types.BucketLoggingStatus

	// The MD5 hash of the PutBucketLogging request body.
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutBucketLoggingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketLoggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketLogging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketLogging{}, middleware.After)
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
	if err = addOpPutBucketLoggingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketLogging(options.Region), middleware.Before); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opPutBucketLogging(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketLogging",
	}
}
