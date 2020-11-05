// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the major version families of each published schema. If a major version
// ARN is provided as SchemaArn, the minor version revisions in that family are
// listed instead.
func (c *Client) ListPublishedSchemaArns(ctx context.Context, params *ListPublishedSchemaArnsInput, optFns ...func(*Options)) (*ListPublishedSchemaArnsOutput, error) {
	if params == nil {
		params = &ListPublishedSchemaArnsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublishedSchemaArns", params, optFns, addOperationListPublishedSchemaArnsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublishedSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPublishedSchemaArnsInput struct {

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// The response for ListPublishedSchemaArns when this parameter is used will list
	// all minor version ARNs for a major version.
	SchemaArn *string
}

type ListPublishedSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of published schemas.
	SchemaArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPublishedSchemaArnsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPublishedSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPublishedSchemaArns{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPublishedSchemaArns(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPublishedSchemaArns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListPublishedSchemaArns",
	}
}
