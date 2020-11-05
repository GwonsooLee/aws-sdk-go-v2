// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard
// (PutDeliverabilityDashboardOption operation) for the domain.
func (c *Client) ListDomainDeliverabilityCampaigns(ctx context.Context, params *ListDomainDeliverabilityCampaignsInput, optFns ...func(*Options)) (*ListDomainDeliverabilityCampaignsOutput, error) {
	if params == nil {
		params = &ListDomainDeliverabilityCampaignsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainDeliverabilityCampaigns", params, optFns, addOperationListDomainDeliverabilityCampaignsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainDeliverabilityCampaignsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard
// (PutDeliverabilityDashboardOption operation) for the domain.
type ListDomainDeliverabilityCampaignsInput struct {

	// The last day, in Unix time format, that you want to obtain deliverability data
	// for. This value has to be less than or equal to 30 days after the value of the
	// StartDate parameter.
	//
	// This member is required.
	EndDate *time.Time

	// The first day, in Unix time format, that you want to obtain deliverability data
	// for.
	//
	// This member is required.
	StartDate *time.Time

	// The domain to obtain deliverability data for.
	//
	// This member is required.
	SubscribedDomain *string

	// A token that’s returned from a previous call to the
	// ListDomainDeliverabilityCampaigns operation. This token indicates the position
	// of a campaign in the list of campaigns.
	NextToken *string

	// The maximum number of results to include in response to a single call to the
	// ListDomainDeliverabilityCampaigns operation. If the number of results is larger
	// than the number that you specify in this parameter, the response includes a
	// NextToken element, which you can use to obtain additional results.
	PageSize *int32
}

// An array of objects that provide deliverability data for all the campaigns that
// used a specific domain to send email during a specified time range. This data is
// available for a domain only if you enabled the Deliverability dashboard
// (PutDeliverabilityDashboardOption operation) for the domain.
type ListDomainDeliverabilityCampaignsOutput struct {

	// An array of responses, one for each campaign that used the domain to send email
	// during the specified time range.
	//
	// This member is required.
	DomainDeliverabilityCampaigns []*types.DomainDeliverabilityCampaign

	// A token that’s returned from a previous call to the
	// ListDomainDeliverabilityCampaigns operation. This token indicates the position
	// of the campaign in the list of campaigns.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDomainDeliverabilityCampaignsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainDeliverabilityCampaigns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainDeliverabilityCampaigns{}, middleware.After)
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
	if err = addOpListDomainDeliverabilityCampaignsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainDeliverabilityCampaigns(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDomainDeliverabilityCampaigns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListDomainDeliverabilityCampaigns",
	}
}
