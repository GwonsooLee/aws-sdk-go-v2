// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudsearch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Tells the search domain to start indexing its documents using the latest
// indexing options. This operation must be invoked to activate options whose
// OptionStatus is RequiresIndexDocuments.
func (c *Client) IndexDocuments(ctx context.Context, params *IndexDocumentsInput, optFns ...func(*Options)) (*IndexDocumentsOutput, error) {
	if params == nil {
		params = &IndexDocumentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IndexDocuments", params, optFns, addOperationIndexDocumentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IndexDocumentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the IndexDocuments operation. Specifies the name
// of the domain you want to re-index.
type IndexDocumentsInput struct {

	// A string that represents the name of a domain. Domain names are unique across
	// the domains owned by an account within an AWS region. Domain names start with a
	// letter or number and can contain the following characters: a-z (lowercase), 0-9,
	// and - (hyphen).
	//
	// This member is required.
	DomainName *string
}

// The result of an IndexDocuments request. Contains the status of the indexing
// operation, including the fields being indexed.
type IndexDocumentsOutput struct {

	// The names of the fields that are currently being indexed.
	FieldNames []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationIndexDocumentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpIndexDocuments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpIndexDocuments{}, middleware.After)
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
	if err = addOpIndexDocumentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIndexDocuments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIndexDocuments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudsearch",
		OperationName: "IndexDocuments",
	}
}
