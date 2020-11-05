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

// Creates and updates the configuration aggregator with the selected source
// accounts and regions. The source account can be individual account(s) or an
// organization. AWS Config should be enabled in source accounts and regions you
// want to aggregate. If your source type is an organization, you must be signed in
// to the master account and all features must be enabled in your organization. AWS
// Config calls EnableAwsServiceAccess API to enable integration between AWS Config
// and AWS Organizations.
func (c *Client) PutConfigurationAggregator(ctx context.Context, params *PutConfigurationAggregatorInput, optFns ...func(*Options)) (*PutConfigurationAggregatorOutput, error) {
	if params == nil {
		params = &PutConfigurationAggregatorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutConfigurationAggregator", params, optFns, addOperationPutConfigurationAggregatorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutConfigurationAggregatorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutConfigurationAggregatorInput struct {

	// The name of the configuration aggregator.
	//
	// This member is required.
	ConfigurationAggregatorName *string

	// A list of AccountAggregationSource object.
	AccountAggregationSources []*types.AccountAggregationSource

	// An OrganizationAggregationSource object.
	OrganizationAggregationSource *types.OrganizationAggregationSource

	// An array of tag object.
	Tags []*types.Tag
}

type PutConfigurationAggregatorOutput struct {

	// Returns a ConfigurationAggregator object.
	ConfigurationAggregator *types.ConfigurationAggregator

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutConfigurationAggregatorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutConfigurationAggregator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutConfigurationAggregator{}, middleware.After)
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
	if err = addOpPutConfigurationAggregatorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutConfigurationAggregator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutConfigurationAggregator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutConfigurationAggregator",
	}
}
