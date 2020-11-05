// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Exports a deployed version of a RestApi in a specified format.
func (c *Client) GetExport(ctx context.Context, params *GetExportInput, optFns ...func(*Options)) (*GetExportOutput, error) {
	if params == nil {
		params = &GetExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExport", params, optFns, addOperationGetExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request a new export of a RestApi for a particular Stage.
type GetExportInput struct {

	// [Required] The type of export. Acceptable values are 'oas30' for OpenAPI 3.0.x
	// and 'swagger' for Swagger/OpenAPI 2.0.
	//
	// This member is required.
	ExportType *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// [Required] The name of the Stage that will be exported.
	//
	// This member is required.
	StageName *string

	// The content-type of the export, for example application/json. Currently
	// application/json and application/yaml are supported for exportType ofoas30 and
	// swagger. This should be specified in the Accept header for direct API requests.
	Accepts *string

	// A key-value map of query string parameters that specify properties of the
	// export, depending on the requested exportType. For exportTypeoas30 and swagger,
	// any combination of the following parameters are supported:
	// extensions='integrations' or extensions='apigateway' will export the API with
	// x-amazon-apigateway-integration extensions. extensions='authorizers' will export
	// the API with x-amazon-apigateway-authorizer extensions. postman will export the
	// API with Postman extensions, allowing for import to the Postman tool
	Parameters map[string]*string
}

// The binary blob response to GetExport, which contains the generated SDK.
type GetExportOutput struct {

	// The binary blob response to GetExport, which contains the export.
	Body []byte

	// The content-disposition header value in the HTTP response.
	ContentDisposition *string

	// The content-type header value in the HTTP response. This will correspond to a
	// valid 'accept' type in the request.
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetExport{}, middleware.After)
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
	if err = addOpGetExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetExport",
	}
}
