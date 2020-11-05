// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all Systems Manager (SSM) documents in the current AWS account and
// Region. You can limit the results of this request by using a filter.
func (c *Client) ListDocuments(ctx context.Context, params *ListDocumentsInput, optFns ...func(*Options)) (*ListDocumentsOutput, error) {
	if params == nil {
		params = &ListDocumentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDocuments", params, optFns, addOperationListDocumentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentsInput struct {

	// This data type is deprecated. Instead, use Filters.
	DocumentFilterList []*types.DocumentFilter

	// One or more DocumentKeyValuesFilter objects. Use a filter to return a more
	// specific list of results. For keys, you can specify one or more key-value pair
	// tags that have been applied to a document. Other valid keys include Owner, Name,
	// PlatformTypes, DocumentType, and TargetType. For example, to return documents
	// you own use Key=Owner,Values=Self. To specify a custom key-value pair, use the
	// format Key=tag:tagName,Values=valueName.
	Filters []*types.DocumentKeyValuesFilter

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
}

type ListDocumentsOutput struct {

	// The names of the Systems Manager documents.
	DocumentIdentifiers []*types.DocumentIdentifier

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDocumentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDocuments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocuments{}, middleware.After)
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
	if err = addOpListDocumentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDocuments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDocuments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "ListDocuments",
	}
}
