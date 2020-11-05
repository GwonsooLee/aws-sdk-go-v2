// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enable or disable email sending for messages that use a particular configuration
// set in a specific AWS Region.
func (c *Client) PutConfigurationSetSendingOptions(ctx context.Context, params *PutConfigurationSetSendingOptionsInput, optFns ...func(*Options)) (*PutConfigurationSetSendingOptionsOutput, error) {
	if params == nil {
		params = &PutConfigurationSetSendingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationSetSendingOptions", params, optFns, addOperationPutConfigurationSetSendingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationSetSendingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable the ability of Amazon Pinpoint to send emails
// that use a specific configuration set.
type PutConfigurationSetSendingOptionsInput struct {

	// The name of the configuration set that you want to enable or disable email
	// sending for.
	//
	// This member is required.
	ConfigurationSetName *string

	// If true, email sending is enabled for the configuration set. If false, email
	// sending is disabled for the configuration set.
	SendingEnabled *bool
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutConfigurationSetSendingOptionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutConfigurationSetSendingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutConfigurationSetSendingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutConfigurationSetSendingOptions{}, middleware.After)
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
	if err = addOpPutConfigurationSetSendingOptionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationSetSendingOptions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigurationSetSendingOptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutConfigurationSetSendingOptions",
	}
}
