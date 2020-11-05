// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets information about all of the versions of an intent. The GetIntentVersions
// operation returns an IntentMetadata object for each version of an intent. For
// example, if an intent has three numbered versions, the GetIntentVersions
// operation returns four IntentMetadata objects in the response, one for each
// numbered version and one for the $LATEST version. The GetIntentVersions
// operation always returns at least one version, the $LATEST version. This
// operation requires permissions for the lex:GetIntentVersions action.
func (c *Client) GetIntentVersions(ctx context.Context, params *GetIntentVersionsInput, optFns ...func(*Options)) (*GetIntentVersionsOutput, error) {
	if params == nil {
		params = &GetIntentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntentVersions", params, optFns, addOperationGetIntentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntentVersionsInput struct {

	// The name of the intent for which versions should be returned.
	//
	// This member is required.
	Name *string

	// The maximum number of intent versions to return in the response. The default is
	// 10.
	MaxResults *int32

	// A pagination token for fetching the next page of intent versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string
}

type GetIntentVersionsOutput struct {

	// An array of IntentMetadata objects, one for each numbered version of the intent
	// plus one for the $LATEST version.
	Intents []*types.IntentMetadata

	// A pagination token for fetching the next page of intent versions. If the
	// response to this call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of versions, specify the pagination token in
	// the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIntentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntentVersions{}, middleware.After)
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
	if err = addOpGetIntentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntentVersions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetIntentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetIntentVersions",
	}
}
