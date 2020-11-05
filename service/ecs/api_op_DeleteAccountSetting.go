// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables an account setting for a specified IAM user, IAM role, or the root user
// for an account.
func (c *Client) DeleteAccountSetting(ctx context.Context, params *DeleteAccountSettingInput, optFns ...func(*Options)) (*DeleteAccountSettingOutput, error) {
	if params == nil {
		params = &DeleteAccountSettingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccountSetting", params, optFns, addOperationDeleteAccountSettingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAccountSettingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAccountSettingInput struct {

	// The resource name for which to disable the account setting. If
	// serviceLongArnFormat is specified, the ARN for your Amazon ECS services is
	// affected. If taskLongArnFormat is specified, the ARN and resource ID for your
	// Amazon ECS tasks is affected. If containerInstanceLongArnFormat is specified,
	// the ARN and resource ID for your Amazon ECS container instances is affected. If
	// awsvpcTrunking is specified, the ENI limit for your Amazon ECS container
	// instances is affected.
	//
	// This member is required.
	Name types.SettingName

	// The ARN of the principal, which can be an IAM user, IAM role, or the root user.
	// If you specify the root user, it disables the account setting for all IAM users,
	// IAM roles, and the root user of the account unless an IAM user or role
	// explicitly overrides these settings. If this field is omitted, the setting is
	// changed only for the authenticated user.
	PrincipalArn *string
}

type DeleteAccountSettingOutput struct {

	// The account setting for the specified principal ARN.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAccountSettingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAccountSetting{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAccountSetting{}, middleware.After)
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
	if err = addOpDeleteAccountSettingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccountSetting(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteAccountSetting(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DeleteAccountSetting",
	}
}
