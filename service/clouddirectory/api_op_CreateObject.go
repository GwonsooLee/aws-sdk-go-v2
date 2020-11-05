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

// Creates an object in a Directory. Additionally attaches the object to a parent,
// if a parent reference and LinkName is specified. An object is simply a
// collection of Facet attributes. You can also use this API call to create a
// policy object, if the facet from which you create the object is a policy facet.
func (c *Client) CreateObject(ctx context.Context, params *CreateObjectInput, optFns ...func(*Options)) (*CreateObjectOutput, error) {
	if params == nil {
		params = &CreateObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateObject", params, optFns, addOperationCreateObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateObjectInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory in which
	// the object will be created. For more information, see arns.
	//
	// This member is required.
	DirectoryArn *string

	// A list of schema facets to be associated with the object. Do not provide minor
	// version components. See SchemaFacet for details.
	//
	// This member is required.
	SchemaFacets []*types.SchemaFacet

	// The name of link that is used to attach this object to a parent.
	LinkName *string

	// The attribute map whose attribute ARN contains the key and attribute value as
	// the map value.
	ObjectAttributeList []*types.AttributeKeyAndValue

	// If specified, the parent reference to which this object will be attached.
	ParentReference *types.ObjectReference
}

type CreateObjectOutput struct {

	// The identifier that is associated with the object.
	ObjectIdentifier *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateObject{}, middleware.After)
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
	if err = addOpCreateObjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateObject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateObject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "CreateObject",
	}
}
