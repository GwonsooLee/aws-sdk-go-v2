// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a resource definition which contains a list of resources to be used in a
// group. You can create an initial version of the definition by providing a list
// of resources now, or use ''CreateResourceDefinitionVersion'' later.
func (c *Client) CreateResourceDefinition(ctx context.Context, params *CreateResourceDefinitionInput, optFns ...func(*Options)) (*CreateResourceDefinitionOutput, error) {
	if params == nil {
		params = &CreateResourceDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourceDefinition", params, optFns, addOperationCreateResourceDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceDefinitionInput struct {

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// Information about the initial version of the resource definition.
	InitialVersion *types.ResourceDefinitionVersion

	// The name of the resource definition.
	Name *string

	// Tag(s) to add to the new resource.
	Tags map[string]*string
}

type CreateResourceDefinitionOutput struct {

	// The ARN of the definition.
	Arn *string

	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string

	// The ID of the definition.
	Id *string

	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string

	// The ID of the latest version associated with the definition.
	LatestVersion *string

	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string

	// The name of the definition.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateResourceDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourceDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourceDefinition{}, middleware.After)
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
	if err = addOpCreateResourceDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourceDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateResourceDefinition",
	}
}
