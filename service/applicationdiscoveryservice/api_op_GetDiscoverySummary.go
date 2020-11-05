// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a short summary of discovered assets. This API operation takes no
// request parameters and is called as is at the command prompt as shown in the
// example.
func (c *Client) GetDiscoverySummary(ctx context.Context, params *GetDiscoverySummaryInput, optFns ...func(*Options)) (*GetDiscoverySummaryOutput, error) {
	if params == nil {
		params = &GetDiscoverySummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDiscoverySummary", params, optFns, addOperationGetDiscoverySummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDiscoverySummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDiscoverySummaryInput struct {
}

type GetDiscoverySummaryOutput struct {

	// Details about discovered agents, including agent status and health.
	AgentSummary *types.CustomerAgentInfo

	// The number of applications discovered.
	Applications *int64

	// Details about discovered connectors, including connector status and health.
	ConnectorSummary *types.CustomerConnectorInfo

	// The number of servers discovered.
	Servers *int64

	// The number of servers mapped to applications.
	ServersMappedToApplications *int64

	// The number of servers mapped to tags.
	ServersMappedtoTags *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDiscoverySummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDiscoverySummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDiscoverySummary{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDiscoverySummary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDiscoverySummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "GetDiscoverySummary",
	}
}
