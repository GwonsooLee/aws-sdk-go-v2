// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the content of an approval rule template. You can change the number of
// required approvals, the membership of the approval rule, and whether an approval
// pool is defined.
func (c *Client) UpdateApprovalRuleTemplateContent(ctx context.Context, params *UpdateApprovalRuleTemplateContentInput, optFns ...func(*Options)) (*UpdateApprovalRuleTemplateContentOutput, error) {
	if params == nil {
		params = &UpdateApprovalRuleTemplateContentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApprovalRuleTemplateContent", params, optFns, addOperationUpdateApprovalRuleTemplateContentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApprovalRuleTemplateContentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApprovalRuleTemplateContentInput struct {

	// The name of the approval rule template where you want to update the content of
	// the rule.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The content that replaces the existing content of the rule. Content statements
	// must be complete. You cannot provide only the changes.
	//
	// This member is required.
	NewRuleContent *string

	// The SHA-256 hash signature for the content of the approval rule. You can
	// retrieve this information by using GetPullRequest.
	ExistingRuleContentSha256 *string
}

type UpdateApprovalRuleTemplateContentOutput struct {

	// Returns information about an approval rule template.
	//
	// This member is required.
	ApprovalRuleTemplate *types.ApprovalRuleTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateApprovalRuleTemplateContentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApprovalRuleTemplateContent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApprovalRuleTemplateContent{}, middleware.After)
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
	if err = addOpUpdateApprovalRuleTemplateContentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApprovalRuleTemplateContent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApprovalRuleTemplateContent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdateApprovalRuleTemplateContent",
	}
}
