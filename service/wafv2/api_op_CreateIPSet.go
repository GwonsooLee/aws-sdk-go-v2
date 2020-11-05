// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
// Creates an IPSet, which you use to identify web requests that originate from
// specific IP addresses or ranges of IP addresses. For example, if you're
// receiving a lot of requests from a ranges of IP addresses, you can configure AWS
// WAF to block them using an IPSet that lists those IP addresses.
func (c *Client) CreateIPSet(ctx context.Context, params *CreateIPSetInput, optFns ...func(*Options)) (*CreateIPSetOutput, error) {
	if params == nil {
		params = &CreateIPSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIPSet", params, optFns, addOperationCreateIPSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIPSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIPSetInput struct {

	// Contains an array of strings that specify one or more IP addresses or blocks of
	// IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF supports
	// all address ranges for IP versions IPv4 and IPv6. Examples:
	//
	// * To configure AWS
	// WAF to allow, block, or count requests that originated from the IP address
	// 192.0.2.44, specify 192.0.2.44/32.
	//
	// * To configure AWS WAF to allow, block, or
	// count requests that originated from IP addresses from 192.0.2.0 to 192.0.2.255,
	// specify 192.0.2.0/24.
	//
	// * To configure AWS WAF to allow, block, or count requests
	// that originated from the IP address 1111:0000:0000:0000:0000:0000:0000:0111,
	// specify 1111:0000:0000:0000:0000:0000:0000:0111/128.
	//
	// * To configure AWS WAF to
	// allow, block, or count requests that originated from IP addresses
	// 1111:0000:0000:0000:0000:0000:0000:0000 to
	// 1111:0000:0000:0000:ffff:ffff:ffff:ffff, specify
	// 1111:0000:0000:0000:0000:0000:0000:0000/64.
	//
	// For more information about CIDR
	// notation, see the Wikipedia entry Classless Inter-Domain Routing
	// (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
	//
	// This member is required.
	Addresses []*string

	// Specify IPV4 or IPV6.
	//
	// This member is required.
	IPAddressVersion types.IPAddressVersion

	// The name of the IP set. You cannot change the name of an IPSet after you create
	// it.
	//
	// This member is required.
	Name *string

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	// * CLI - Specify
	// the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	// * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// A description of the IP set that helps with identification. You cannot change
	// the description of an IP set after you create it.
	Description *string

	// An array of key:value pairs to associate with the resource.
	Tags []*types.Tag
}

type CreateIPSetOutput struct {

	// High-level information about an IPSet, returned by operations like create and
	// list. This provides information like the ID, that you can use to retrieve and
	// manage an IPSet, and the ARN, that you provide to the IPSetReferenceStatement to
	// use the address set in a Rule.
	Summary *types.IPSetSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateIPSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateIPSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateIPSet{}, middleware.After)
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
	if err = addOpCreateIPSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIPSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIPSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "CreateIPSet",
	}
}
