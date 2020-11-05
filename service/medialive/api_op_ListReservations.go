// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List purchased reservations.
func (c *Client) ListReservations(ctx context.Context, params *ListReservationsInput, optFns ...func(*Options)) (*ListReservationsOutput, error) {
	if params == nil {
		params = &ListReservationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReservations", params, optFns, addOperationListReservationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReservationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Placeholder documentation for ListReservationsRequest
type ListReservationsInput struct {

	// Filter by channel class, 'STANDARD' or 'SINGLE_PIPELINE'
	ChannelClass *string

	// Filter by codec, 'AVC', 'HEVC', 'MPEG2', 'AUDIO', or 'LINK'
	Codec *string

	// Placeholder documentation for MaxResults
	MaxResults *int32

	// Filter by bitrate, 'MAX_10_MBPS', 'MAX_20_MBPS', or 'MAX_50_MBPS'
	MaximumBitrate *string

	// Filter by framerate, 'MAX_30_FPS' or 'MAX_60_FPS'
	MaximumFramerate *string

	// Placeholder documentation for __string
	NextToken *string

	// Filter by resolution, 'SD', 'HD', 'FHD', or 'UHD'
	Resolution *string

	// Filter by resource type, 'INPUT', 'OUTPUT', 'MULTIPLEX', or 'CHANNEL'
	ResourceType *string

	// Filter by special feature, 'ADVANCED_AUDIO' or 'AUDIO_NORMALIZATION'
	SpecialFeature *string

	// Filter by video quality, 'STANDARD', 'ENHANCED', or 'PREMIUM'
	VideoQuality *string
}

// Placeholder documentation for ListReservationsResponse
type ListReservationsOutput struct {

	// Token to retrieve the next page of results
	NextToken *string

	// List of reservations
	Reservations []*types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListReservationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListReservations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListReservations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListReservations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListReservations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "ListReservations",
	}
}
