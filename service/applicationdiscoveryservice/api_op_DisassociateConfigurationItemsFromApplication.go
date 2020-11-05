// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates one or more configuration items from an application.
func (c *Client) DisassociateConfigurationItemsFromApplication(ctx context.Context, params *DisassociateConfigurationItemsFromApplicationInput, optFns ...func(*Options)) (*DisassociateConfigurationItemsFromApplicationOutput, error) {
	if params == nil {
		params = &DisassociateConfigurationItemsFromApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateConfigurationItemsFromApplication", params, optFns, addOperationDisassociateConfigurationItemsFromApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateConfigurationItemsFromApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateConfigurationItemsFromApplicationInput struct {

	// Configuration ID of an application from which each item is disassociated.
	//
	// This member is required.
	ApplicationConfigurationId *string

	// Configuration ID of each item to be disassociated from an application.
	//
	// This member is required.
	ConfigurationIds []*string
}

type DisassociateConfigurationItemsFromApplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateConfigurationItemsFromApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateConfigurationItemsFromApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateConfigurationItemsFromApplication{}, middleware.After)
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
	if err = addOpDisassociateConfigurationItemsFromApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateConfigurationItemsFromApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateConfigurationItemsFromApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DisassociateConfigurationItemsFromApplication",
	}
}
