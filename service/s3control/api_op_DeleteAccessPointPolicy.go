// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"strings"
)

// Deletes the access point policy for the specified access point. All Amazon S3 on
// Outposts REST API requests for this action require an additional parameter of
// outpost-id to be passed with the request and an S3 on Outposts endpoint hostname
// prefix instead of s3-control. For an example of the request syntax for Amazon S3
// on Outposts that uses the S3 on Outposts endpoint hostname prefix and the
// outpost-id derived using the access point ARN, see the  Example
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API__control_DeleteAccessPointPolicy.html#API_control_DeleteAccessPointPolicy_Examples)
// section below. The following actions are related to DeleteAccessPointPolicy:
//
// *
// PutAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html)
//
// *
// GetAccessPointPolicy
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html)
func (c *Client) DeleteAccessPointPolicy(ctx context.Context, params *DeleteAccessPointPolicyInput, optFns ...func(*Options)) (*DeleteAccessPointPolicyOutput, error) {
	if params == nil {
		params = &DeleteAccessPointPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccessPointPolicy", params, optFns, addOperationDeleteAccessPointPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccessPointPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccessPointPolicyInput struct {

	// The account ID for the account that owns the specified access point.
	//
	// This member is required.
	AccountId *string

	// The name of the access point whose policy you want to delete. For Amazon S3 on
	// Outposts specify the ARN of the access point accessed in the format
	// arn:aws:s3-outposts:::outpost//accesspoint/. For example, to access the access
	// point reports-ap through outpost my-outpost owned by account 123456789012 in
	// Region us-west-2, use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/accesspoint/reports-ap.
	// The value must be URL encoded.
	//
	// This member is required.
	Name *string
}

type DeleteAccessPointPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccessPointPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteAccessPointPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteAccessPointPolicy{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteAccessPointPolicyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteAccessPointPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccessPointPolicy(options.Region), middleware.Before); err != nil {
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
	return nil
}

type endpointPrefix_opDeleteAccessPointPolicyMiddleware struct {
}

func (*endpointPrefix_opDeleteAccessPointPolicyMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteAccessPointPolicyMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*DeleteAccessPointPolicyInput)
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
func addEndpointPrefix_opDeleteAccessPointPolicyMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteAccessPointPolicyMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteAccessPointPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "DeleteAccessPointPolicy",
	}
}
