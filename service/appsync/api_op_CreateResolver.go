// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Resolver object. A resolver converts incoming requests into a format
// that a data source can understand and converts the data source's responses into
// GraphQL.
func (c *Client) CreateResolver(ctx context.Context, params *CreateResolverInput, optFns ...func(*Options)) (*CreateResolverOutput, error) {
	if params == nil {
		params = &CreateResolverInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResolver", params, optFns, addOperationCreateResolverMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResolverOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverInput struct {

	// The ID for the GraphQL API for which the resolver is being created.
	//
	// This member is required.
	ApiId *string

	// The name of the field to attach the resolver to.
	//
	// This member is required.
	FieldName *string

	// The name of the Type.
	//
	// This member is required.
	TypeName *string

	// The caching configuration for the resolver.
	CachingConfig *types.CachingConfig

	// The name of the data source for which the resolver is being created.
	DataSourceName *string

	// The resolver type.
	//
	// * UNIT: A UNIT resolver type. A UNIT resolver is the default
	// resolver type. A UNIT resolver enables you to execute a GraphQL query against a
	// single data source.
	//
	// * PIPELINE: A PIPELINE resolver type. A PIPELINE resolver
	// enables you to execute a series of Function in a serial manner. You can use a
	// pipeline resolver to execute a GraphQL query against multiple data sources.
	Kind types.ResolverKind

	// The PipelineConfig.
	PipelineConfig *types.PipelineConfig

	// The mapping template to be used for requests. A resolver uses a request mapping
	// template to convert a GraphQL expression into a format that a data source can
	// understand. Mapping templates are written in Apache Velocity Template Language
	// (VTL). VTL request mapping templates are optional when using a Lambda data
	// source. For all other data sources, VTL request and response mapping templates
	// are required.
	RequestMappingTemplate *string

	// The mapping template to be used for responses from the data source.
	ResponseMappingTemplate *string

	// The SyncConfig for a resolver attached to a versioned datasource.
	SyncConfig *types.SyncConfig
}

type CreateResolverOutput struct {

	// The Resolver object.
	Resolver *types.Resolver

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateResolverMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResolver{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResolver{}, middleware.After)
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
	if err = addOpCreateResolverValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolver(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResolver(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "CreateResolver",
	}
}
