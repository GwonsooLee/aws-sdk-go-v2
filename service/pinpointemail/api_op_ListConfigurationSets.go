// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all of the configuration sets associated with your Amazon Pinpoint account
// in the current region. In Amazon Pinpoint, configuration sets are groups of
// rules that you can apply to the emails you send. You apply a configuration set
// to an email by including a reference to the configuration set in the headers of
// the email. When you apply a configuration set to an email, all of the rules in
// that configuration set are applied to the email.
func (c *Client) ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) {
	if params == nil {
		params = &ListConfigurationSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConfigurationSets", params, optFns, addOperationListConfigurationSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConfigurationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain a list of configuration sets for your Amazon Pinpoint
// account in the current AWS Region.
type ListConfigurationSetsInput struct {

	// A token returned from a previous call to ListConfigurationSets to indicate the
	// position in the list of configuration sets.
	NextToken *string

	// The number of results to show in a single call to ListConfigurationSets. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results.
	PageSize *int32
}

// A list of configuration sets in your Amazon Pinpoint account in the current AWS
// Region.
type ListConfigurationSetsOutput struct {

	// An array that contains all of the configuration sets in your Amazon Pinpoint
	// account in the current AWS Region.
	ConfigurationSets []*string

	// A token that indicates that there are additional configuration sets to list. To
	// view additional configuration sets, issue another request to
	// ListConfigurationSets, and pass this token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListConfigurationSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConfigurationSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConfigurationSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationSets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConfigurationSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListConfigurationSets",
	}
}
