// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the virtual interface failover test history.
func (c *Client) ListVirtualInterfaceTestHistory(ctx context.Context, params *ListVirtualInterfaceTestHistoryInput, optFns ...func(*Options)) (*ListVirtualInterfaceTestHistoryOutput, error) {
	if params == nil {
		params = &ListVirtualInterfaceTestHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualInterfaceTestHistory", params, optFns, addOperationListVirtualInterfaceTestHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualInterfaceTestHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListVirtualInterfaceTestHistoryInput struct {

	// The BGP peers that were placed in the DOWN state during the virtual interface
	// failover test.
	BgpPeers []*string

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value. If
	// MaxResults is given a value larger than 100, only 100 results are returned.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The status of the virtual interface failover test.
	Status *string

	// The ID of the virtual interface failover test.
	TestId *string

	// The ID of the virtual interface that was tested.
	VirtualInterfaceId *string
}

type ListVirtualInterfaceTestHistoryOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The ID of the tested virtual interface.
	VirtualInterfaceTestHistory []*types.VirtualInterfaceTestHistory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVirtualInterfaceTestHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListVirtualInterfaceTestHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVirtualInterfaceTestHistory{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListVirtualInterfaceTestHistory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListVirtualInterfaceTestHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "ListVirtualInterfaceTestHistory",
	}
}
