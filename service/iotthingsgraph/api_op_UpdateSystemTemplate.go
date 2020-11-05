// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the specified system. You don't need to run this action after updating a
// workflow. Any deployment that uses the system will see the changes in the system
// when it is redeployed.
func (c *Client) UpdateSystemTemplate(ctx context.Context, params *UpdateSystemTemplateInput, optFns ...func(*Options)) (*UpdateSystemTemplateOutput, error) {
	if params == nil {
		params = &UpdateSystemTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSystemTemplate", params, optFns, addOperationUpdateSystemTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSystemTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSystemTemplateInput struct {

	// The DefinitionDocument that contains the updated system definition.
	//
	// This member is required.
	Definition *types.DefinitionDocument

	// The ID of the system to be updated. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:system:SYSTEMNAME
	//
	// This member is required.
	Id *string

	// The version of the user's namespace. Defaults to the latest version of the
	// user's namespace. If no value is specified, the latest version is used by
	// default.
	CompatibleNamespaceVersion *int64
}

type UpdateSystemTemplateOutput struct {

	// An object containing summary information about the updated system.
	Summary *types.SystemTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateSystemTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSystemTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSystemTemplate{}, middleware.After)
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
	if err = addOpUpdateSystemTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSystemTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSystemTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "UpdateSystemTemplate",
	}
}
