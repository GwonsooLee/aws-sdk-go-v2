// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an event destination in a configuration set.
func (c *Client) DeleteConfigurationSetEventDestination(ctx context.Context, params *DeleteConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*DeleteConfigurationSetEventDestinationOutput, error) {
	if params == nil {
		params = &DeleteConfigurationSetEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConfigurationSetEventDestination", params, optFns, addOperationDeleteConfigurationSetEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConfigurationSetEventDestinationInput struct {

	// ConfigurationSetName
	//
	// This member is required.
	ConfigurationSetName *string

	// EventDestinationName
	//
	// This member is required.
	EventDestinationName *string
}

// An empty object that indicates that the event destination was deleted
// successfully.
type DeleteConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConfigurationSetEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConfigurationSetEventDestination{}, middleware.After)
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
	if err = addOpDeleteConfigurationSetEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConfigurationSetEventDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConfigurationSetEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DeleteConfigurationSetEventDestination",
	}
}
