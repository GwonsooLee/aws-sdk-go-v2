// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// This API operation deletes an Amazon S3 on Outposts bucket policy. To delete an
// S3 bucket policy, see DeleteBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html) in
// the Amazon Simple Storage Service API. This implementation of the DELETE
// operation uses the policy subresource to delete the policy of a specified Amazon
// S3 on Outposts bucket. If you are using an identity other than the root user of
// the AWS account that owns the bucket, the calling identity must have the
// s3outposts:DeleteBucketPolicy permissions on the specified Outposts bucket and
// belong to the bucket owner's account to use this operation. For more
// information, see Using Amazon S3 on Outposts
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in Amazon
// Simple Storage Service Developer Guide. If you don't have DeleteBucketPolicy
// permissions, Amazon S3 returns a 403 Access Denied error. If you have the
// correct permissions, but you're not using an identity that belongs to the bucket
// owner's account, Amazon S3 returns a 405 Method Not Allowed error. As a security
// precaution, the root user of the AWS account that owns a bucket can always use
// this operation, even if the policy explicitly denies the root user the ability
// to perform this action. For more information about bucket policies, see Using
// Bucket Policies and User Policies
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). All
// Amazon S3 on Outposts REST API requests for this action require an additional
// parameter of outpost-id to be passed with the request and an S3 on Outposts
// endpoint hostname prefix instead of s3-control. For an example of the request
// syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname
// prefix and the outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteBucketPolicy.html#API_control_DeleteBucketPolicy_Examples)
// section below. The following actions are related to DeleteBucketPolicy:
//
// *
// GetBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html)
//
// *
// PutBucketPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_PutBucketPolicy.html)
func (c *Client) DeleteBucketPolicy(ctx context.Context, params *DeleteBucketPolicyInput, optFns ...func(*Options)) (*DeleteBucketPolicyOutput, error) {
	if params == nil {
		params = &DeleteBucketPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketPolicy", params, optFns, addOperationDeleteBucketPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketPolicyInput struct {

	// The account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// The ARN of the bucket. For Amazon S3 on Outposts specify the ARN of the bucket
	// accessed in the format arn:aws:s3-outposts:::outpost//bucket/. For example, to
	// access the bucket reports through outpost my-outpost owned by account
	// 123456789012 in Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports.
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string
}

type DeleteBucketPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBucketPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketPolicy{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteBucketPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteBucketPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketPolicy(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addDeleteBucketPolicyUpdateEndpoint(stack, options); err != nil {
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
	return nil
}

type endpointPrefix_opDeleteBucketPolicyMiddleware struct {
}

func (*endpointPrefix_opDeleteBucketPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteBucketPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteBucketPolicyInput)
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
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteBucketPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteBucketPolicyMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteBucketPolicy",
	}
}

func copyDeleteBucketPolicyInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteBucketPolicyInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteBucketPolicyInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func getDeleteBucketPolicyARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteBucketPolicyInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setDeleteBucketPolicyARNMember(input interface{}, v string) error {
	in := input.(*DeleteBucketPolicyInput)
	in.Bucket = &v
	return nil
}
func backFillDeleteBucketPolicyAccountID(input interface{}, v string) error {
	in := input.(*DeleteBucketPolicyInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteBucketPolicyUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteBucketPolicyARNMember,
			BackfillAccountID: backFillDeleteBucketPolicyAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteBucketPolicyARNMember,
			CopyInput:         copyDeleteBucketPolicyInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
