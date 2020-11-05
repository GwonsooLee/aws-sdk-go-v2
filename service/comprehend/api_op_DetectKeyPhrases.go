// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detects the key noun phrases found in the text.
func (c *Client) DetectKeyPhrases(ctx context.Context, params *DetectKeyPhrasesInput, optFns ...func(*Options)) (*DetectKeyPhrasesOutput, error) {
	if params == nil {
		params = &DetectKeyPhrasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectKeyPhrases", params, optFns, addOperationDetectKeyPhrasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectKeyPhrasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectKeyPhrasesInput struct {

	// The language of the input documents. You can specify any of the primary
	// languages supported by Amazon Comprehend. All documents must be in the same
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A UTF-8 text string. Each string must contain fewer that 5,000 bytes of UTF-8
	// encoded characters.
	//
	// This member is required.
	Text *string
}

type DetectKeyPhrasesOutput struct {

	// A collection of key phrases that Amazon Comprehend identified in the input text.
	// For each key phrase, the response provides the text of the key phrase, where the
	// key phrase begins and ends, and the level of confidence that Amazon Comprehend
	// has in the accuracy of the detection.
	KeyPhrases []*types.KeyPhrase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetectKeyPhrasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectKeyPhrases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectKeyPhrases{}, middleware.After)
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
	if err = addOpDetectKeyPhrasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectKeyPhrases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDetectKeyPhrases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectKeyPhrases",
	}
}
