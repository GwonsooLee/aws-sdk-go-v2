// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a given object's attributes.
func (c *Client) UpdateObjectAttributes(ctx context.Context, params *UpdateObjectAttributesInput, optFns ...func(*Options)) (*UpdateObjectAttributesOutput, error) {
	if params == nil {
		params = &UpdateObjectAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateObjectAttributes", params, optFns, addOperationUpdateObjectAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateObjectAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateObjectAttributesInput struct {

	// The attributes update structure.
	//
	// This member is required.
	AttributeUpdates []*types.ObjectAttributeUpdate

	// The Amazon Resource Name (ARN) that is associated with the Directory where the
	// object resides. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// The reference that identifies the object.
	//
	// This member is required.
	ObjectReference *types.ObjectReference
}

type UpdateObjectAttributesOutput struct {

	// The ObjectIdentifier of the updated object.
	ObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateObjectAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateObjectAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateObjectAttributes{}, middleware.After)
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
	if err = addOpUpdateObjectAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateObjectAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateObjectAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "UpdateObjectAttributes",
	}
}
