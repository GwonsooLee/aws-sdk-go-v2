// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the versions of a connector definition, which are containers for
// connectors. Connectors run on the Greengrass core and contain built-in
// integration with local infrastructure, device protocols, AWS, and other cloud
// services.
func (c *Client) ListConnectorDefinitionVersions(ctx context.Context, params *ListConnectorDefinitionVersionsInput, optFns ...func(*Options)) (*ListConnectorDefinitionVersionsOutput, error) {
	if params == nil {
		params = &ListConnectorDefinitionVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnectorDefinitionVersions", params, optFns, addOperationListConnectorDefinitionVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectorDefinitionVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorDefinitionVersionsInput struct {

	// The ID of the connector definition.
	//
	// This member is required.
	ConnectorDefinitionId *string

	// The maximum number of results to be returned per request.
	MaxResults *string

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string
}

type ListConnectorDefinitionVersionsOutput struct {

	// The token for the next set of results, or ''null'' if there are no additional
	// results.
	NextToken *string

	// Information about a version.
	Versions []*types.VersionInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConnectorDefinitionVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConnectorDefinitionVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConnectorDefinitionVersions{}, middleware.After)
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
	if err = addOpListConnectorDefinitionVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectorDefinitionVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConnectorDefinitionVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListConnectorDefinitionVersions",
	}
}
