// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes an existing model defined for a RestApi resource.
func (c *Client) GetModel(ctx context.Context, params *GetModelInput, optFns ...func(*Options)) (*GetModelOutput, error) {
	if params == nil {
		params = &GetModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetModel", params, optFns, addOperationGetModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to list information about a model in an existing RestApi resource.
type GetModelInput struct {

	// [Required] The name of the model as an identifier.
	//
	// This member is required.
	ModelName *string

	// [Required] The RestApi identifier under which the Model exists.
	//
	// This member is required.
	RestApiId *string

	// A query parameter of a Boolean value to resolve (true) all external model
	// references and returns a flattened model schema or not (false) The default is
	// false.
	Flatten *bool
}

// Represents the data structure of a method's request or response payload. A
// request model defines the data structure of the client-supplied request payload.
// A response model defines the data structure of the response payload returned by
// the back end. Although not required, models are useful for mapping payloads
// between the front end and back end. A model is used for generating an API's SDK,
// validating the input request body, and creating a skeletal mapping template.
// Method, MethodResponse, Models and Mappings
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/models-mappings.html)
type GetModelOutput struct {

	// The content-type for the model.
	ContentType *string

	// The description of the model.
	Description *string

	// The identifier for the model resource.
	Id *string

	// The name of the model. Must be an alphanumeric string.
	Name *string

	// The schema for the model. For application/json models, this should be JSON
	// schema draft 4 (https://tools.ietf.org/html/draft-zyp-json-schema-04) model. Do
	// not include "\*/" characters in the description of any properties because such
	// "\*/" characters may be interpreted as the closing marker for comments in some
	// languages, such as Java or JavaScript, causing the installation of your API's
	// SDK generated by API Gateway to fail.
	Schema *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetModel{}, middleware.After)
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
	if err = addOpGetModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetModel(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetModel",
	}
}
