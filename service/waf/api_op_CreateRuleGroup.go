// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/waf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Creates a RuleGroup. A rule group is a collection of predefined
// rules that you add to a web ACL. You use UpdateRuleGroup to add rules to the
// rule group. Rule groups are subject to the following limits:
//
// * Three rule
// groups per account. You can request an increase to this limit by contacting
// customer support.
//
// * One rule group per web ACL.
//
// * Ten rules per rule
// group.
//
// For more information about how to use the AWS WAF API to allow or block
// HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateRuleGroup(ctx context.Context, params *CreateRuleGroupInput, optFns ...func(*Options)) (*CreateRuleGroupOutput, error) {
	if params == nil {
		params = &CreateRuleGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRuleGroup", params, optFns, addOperationCreateRuleGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRuleGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRuleGroupInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// A friendly name or description for the metrics for this RuleGroup. The name can
	// contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128
	// and minimum length one. It can't contain whitespace or metric names reserved for
	// AWS WAF, including "All" and "Default_Action." You can't change the name of the
	// metric after you create the RuleGroup.
	//
	// This member is required.
	MetricName *string

	// A friendly name or description of the RuleGroup. You can't change Name after you
	// create a RuleGroup.
	//
	// This member is required.
	Name *string

	//
	Tags []*types.Tag
}

type CreateRuleGroupOutput struct {

	// The ChangeToken that you used to submit the CreateRuleGroup request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// An empty RuleGroup.
	RuleGroup *types.RuleGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRuleGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRuleGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRuleGroup{}, middleware.After)
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
	if err = addOpCreateRuleGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRuleGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRuleGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateRuleGroup",
	}
}
