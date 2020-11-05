// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a transcript of the session. Note that ConnectionToken is used for
// invoking this API instead of ParticipantToken.
func (c *Client) GetTranscript(ctx context.Context, params *GetTranscriptInput, optFns ...func(*Options)) (*GetTranscriptOutput, error) {
	if params == nil {
		params = &GetTranscriptInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTranscript", params, optFns, addOperationGetTranscriptMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTranscriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTranscriptInput struct {

	// The authentication token associated with the participant's connection.
	//
	// This member is required.
	ConnectionToken *string

	// The contactId from the current contact chain for which transcript is needed.
	ContactId *string

	// The maximum number of results to return in the page. Default: 10.
	MaxResults *int32

	// The pagination token. Use the value returned previously in the next subsequent
	// request to retrieve the next set of results.
	NextToken *string

	// The direction from StartPosition from which to retrieve message. Default:
	// BACKWARD when no StartPosition is provided, FORWARD with StartPosition.
	ScanDirection types.ScanDirection

	// The sort order for the records. Default: DESCENDING.
	SortOrder types.SortKey

	// A filtering option for where to start.
	StartPosition *types.StartPosition
}

type GetTranscriptOutput struct {

	// The initial contact ID for the contact.
	InitialContactId *string

	// The pagination token. Use the value returned previously in the next subsequent
	// request to retrieve the next set of results.
	NextToken *string

	// The list of messages in the session.
	Transcript []*types.Item

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTranscriptMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetTranscript{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetTranscript{}, middleware.After)
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
	if err = addOpGetTranscriptValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTranscript(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTranscript(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetTranscript",
	}
}
