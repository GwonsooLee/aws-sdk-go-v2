// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or updates the remediation configuration with a specific AWS Config rule
// with the selected target or action. The API creates the RemediationConfiguration
// object for the AWS Config rule. The AWS Config rule must already exist for you
// to add a remediation configuration. The target (SSM document) must exist and
// have permissions to use the target. If you make backward incompatible changes to
// the SSM document, you must call this again to ensure the remediations can run.
func (c *Client) PutRemediationConfigurations(ctx context.Context, params *PutRemediationConfigurationsInput, optFns ...func(*Options)) (*PutRemediationConfigurationsOutput, error) {
	if params == nil {
		params = &PutRemediationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutRemediationConfigurations", params, optFns, addOperationPutRemediationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutRemediationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutRemediationConfigurationsInput struct {

	// A list of remediation configuration objects.
	//
	// This member is required.
	RemediationConfigurations []*types.RemediationConfiguration
}

type PutRemediationConfigurationsOutput struct {

	// Returns a list of failed remediation batch objects.
	FailedBatches []*types.FailedRemediationBatch

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutRemediationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutRemediationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutRemediationConfigurations{}, middleware.After)
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
	if err = addOpPutRemediationConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutRemediationConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutRemediationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutRemediationConfigurations",
	}
}
