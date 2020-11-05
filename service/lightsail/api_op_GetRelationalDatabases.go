// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about all of your databases in Amazon Lightsail.
func (c *Client) GetRelationalDatabases(ctx context.Context, params *GetRelationalDatabasesInput, optFns ...func(*Options)) (*GetRelationalDatabasesOutput, error) {
	if params == nil {
		params = &GetRelationalDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRelationalDatabases", params, optFns, addOperationGetRelationalDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRelationalDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabasesInput struct {

	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabases request. If your results
	// are paginated, the response will return a next page token that you can specify
	// as the page token in a subsequent request.
	PageToken *string
}

type GetRelationalDatabasesOutput struct {

	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabases request and specify the
	// next page token using the pageToken parameter.
	NextPageToken *string

	// An object describing the result of your get relational databases request.
	RelationalDatabases []*types.RelationalDatabase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetRelationalDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabases{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRelationalDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabases",
	}
}
