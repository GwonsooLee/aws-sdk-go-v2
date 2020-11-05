// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of connection definitions from the Data Catalog.
func (c *Client) GetConnections(ctx context.Context, params *GetConnectionsInput, optFns ...func(*Options)) (*GetConnectionsOutput, error) {
	if params == nil {
		params = &GetConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConnections", params, optFns, addOperationGetConnectionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConnectionsInput struct {

	// The ID of the Data Catalog in which the connections reside. If none is provided,
	// the AWS account ID is used by default.
	CatalogId *string

	// A filter that controls which connections are returned.
	Filter *types.GetConnectionsFilter

	// Allows you to retrieve the connection metadata without returning the password.
	// For instance, the AWS Glue console uses this flag to retrieve the connection,
	// and does not display the password. Set this parameter when the caller might not
	// have permission to use the AWS KMS key to decrypt the password, but it does have
	// permission to access the rest of the connection properties.
	HidePassword *bool

	// The maximum number of connections to return in one response.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string
}

type GetConnectionsOutput struct {

	// A list of requested connection definitions.
	ConnectionList []*types.Connection

	// A continuation token, if the list of connections returned does not include the
	// last of the filtered connections.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetConnectionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConnections{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConnections(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConnections(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetConnections",
	}
}
