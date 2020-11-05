// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new Authorizer resource to an existing RestApi resource. AWS CLI
// (https://docs.aws.amazon.com/cli/latest/reference/apigateway/create-authorizer.html)
func (c *Client) CreateAuthorizer(ctx context.Context, params *CreateAuthorizerInput, optFns ...func(*Options)) (*CreateAuthorizerOutput, error) {
	if params == nil {
		params = &CreateAuthorizerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAuthorizer", params, optFns, addOperationCreateAuthorizerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to add a new Authorizer to an existing RestApi resource.
type CreateAuthorizerInput struct {

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// [Required] The authorizer type. Valid values are TOKEN for a Lambda function
	// using a single authorization token submitted in a custom header, REQUEST for a
	// Lambda function using incoming request parameters, and COGNITO_USER_POOLS for
	// using an Amazon Cognito user pool.
	//
	// This member is required.
	Type types.AuthorizerType

	// Optional customer-defined field, used in OpenAPI imports and exports without
	// functional impact.
	AuthType *string

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, specify null.
	AuthorizerCredentials *string

	// The TTL in seconds of cached authorizer results. If it equals 0, authorization
	// caching is disabled. If it is greater than 0, API Gateway will cache authorizer
	// responses. If this field is not set, the default value is 300. The maximum value
	// is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *int32

	// Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or
	// REQUEST authorizers, this must be a well-formed Lambda function URI, for
	// example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form
	// arn:aws:apigateway:{region}:lambda:path/{service_api}, where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations.
	AuthorizerUri *string

	// The identity source for which authorization is requested.
	//
	// * For a TOKEN or
	// COGNITO_USER_POOLS authorizer, this is required and specifies the request header
	// mapping expression for the custom header holding the authorization token
	// submitted by the client. For example, if the token header name is Auth, the
	// header mapping expression is method.request.header.Auth.
	//
	// * For the REQUEST
	// authorizer, this is required when authorization caching is enabled. The value is
	// a comma-separated string of one or more mapping expressions of the specified
	// request parameters. For example, if an Auth header, a Name query string
	// parameter are defined as identity sources, this value is
	// method.request.header.Auth, method.request.querystring.Name. These parameters
	// will be used to derive the authorization caching key and to perform runtime
	// validation of the REQUEST authorizer by verifying all of the identity-related
	// request parameters are present, not null and non-empty. Only when this is true
	// does the authorizer invoke the authorizer Lambda function, otherwise, it returns
	// a 401 Unauthorized response without calling the Lambda function. The valid value
	// is a string of comma-separated mapping expressions of the specified request
	// parameters. When the authorization caching is not enabled, this property is
	// optional.
	IdentitySource *string

	// A validation expression for the incoming identity token. For TOKEN authorizers,
	// this value is a regular expression. For COGNITO_USER_POOLS authorizers, API
	// Gateway will match the aud field of the incoming token from the client against
	// the specified regular expression. It will invoke the authorizer's Lambda
	// function when there is a match. Otherwise, it will return a 401 Unauthorized
	// response without calling the Lambda function. The validation expression does not
	// apply to the REQUEST authorizer.
	IdentityValidationExpression *string

	// A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS
	// authorizer. Each element is of this format:
	// arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}. For a TOKEN
	// or REQUEST authorizer, this is not defined.
	ProviderARNs []*string
}

// Represents an authorization layer for methods. If enabled on a method, API
// Gateway will activate the authorizer when a client calls the method. Use Lambda
// Function as Authorizer
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-use-lambda-authorizer.html)Use
// Cognito User Pool as Authorizer
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/apigateway-integrate-with-cognito.html)
type CreateAuthorizerOutput struct {

	// Optional customer-defined field, used in OpenAPI imports and exports without
	// functional impact.
	AuthType *string

	// Specifies the required credentials as an IAM role for API Gateway to invoke the
	// authorizer. To specify an IAM role for API Gateway to assume, use the role's
	// Amazon Resource Name (ARN). To use resource-based permissions on the Lambda
	// function, specify null.
	AuthorizerCredentials *string

	// The TTL in seconds of cached authorizer results. If it equals 0, authorization
	// caching is disabled. If it is greater than 0, API Gateway will cache authorizer
	// responses. If this field is not set, the default value is 300. The maximum value
	// is 3600, or 1 hour.
	AuthorizerResultTtlInSeconds *int32

	// Specifies the authorizer's Uniform Resource Identifier (URI). For TOKEN or
	// REQUEST authorizers, this must be a well-formed Lambda function URI, for
	// example,
	// arn:aws:apigateway:us-west-2:lambda:path/2015-03-31/functions/arn:aws:lambda:us-west-2:{account_id}:function:{lambda_function_name}/invocations.
	// In general, the URI has this form
	// arn:aws:apigateway:{region}:lambda:path/{service_api}, where {region} is the
	// same as the region hosting the Lambda function, path indicates that the
	// remaining substring in the URI should be treated as the path to the resource,
	// including the initial /. For Lambda functions, this is usually of the form
	// /2015-03-31/functions/[FunctionARN]/invocations.
	AuthorizerUri *string

	// The identifier for the authorizer resource.
	Id *string

	// The identity source for which authorization is requested.
	//
	// * For a TOKEN or
	// COGNITO_USER_POOLS authorizer, this is required and specifies the request header
	// mapping expression for the custom header holding the authorization token
	// submitted by the client. For example, if the token header name is Auth, the
	// header mapping expression is method.request.header.Auth.
	//
	// * For the REQUEST
	// authorizer, this is required when authorization caching is enabled. The value is
	// a comma-separated string of one or more mapping expressions of the specified
	// request parameters. For example, if an Auth header, a Name query string
	// parameter are defined as identity sources, this value is
	// method.request.header.Auth, method.request.querystring.Name. These parameters
	// will be used to derive the authorization caching key and to perform runtime
	// validation of the REQUEST authorizer by verifying all of the identity-related
	// request parameters are present, not null and non-empty. Only when this is true
	// does the authorizer invoke the authorizer Lambda function, otherwise, it returns
	// a 401 Unauthorized response without calling the Lambda function. The valid value
	// is a string of comma-separated mapping expressions of the specified request
	// parameters. When the authorization caching is not enabled, this property is
	// optional.
	IdentitySource *string

	// A validation expression for the incoming identity token. For TOKEN authorizers,
	// this value is a regular expression. For COGNITO_USER_POOLS authorizers, API
	// Gateway will match the aud field of the incoming token from the client against
	// the specified regular expression. It will invoke the authorizer's Lambda
	// function when there is a match. Otherwise, it will return a 401 Unauthorized
	// response without calling the Lambda function. The validation expression does not
	// apply to the REQUEST authorizer.
	IdentityValidationExpression *string

	// [Required] The name of the authorizer.
	Name *string

	// A list of the Amazon Cognito user pool ARNs for the COGNITO_USER_POOLS
	// authorizer. Each element is of this format:
	// arn:aws:cognito-idp:{region}:{account_id}:userpool/{user_pool_id}. For a TOKEN
	// or REQUEST authorizer, this is not defined.
	ProviderARNs []*string

	// The authorizer type. Valid values are TOKEN for a Lambda function using a single
	// authorization token submitted in a custom header, REQUEST for a Lambda function
	// using incoming request parameters, and COGNITO_USER_POOLS for using an Amazon
	// Cognito user pool.
	Type types.AuthorizerType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAuthorizerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAuthorizer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAuthorizer{}, middleware.After)
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
	if err = addOpCreateAuthorizerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAuthorizer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAuthorizer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateAuthorizer",
	}
}
