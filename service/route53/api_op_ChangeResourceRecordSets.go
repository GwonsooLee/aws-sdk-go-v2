// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	route53cust "github.com/aws/aws-sdk-go-v2/service/route53/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates, changes, or deletes a resource record set, which contains authoritative
// DNS information for a specified domain name or subdomain name. For example, you
// can use ChangeResourceRecordSets to create a resource record set that routes
// traffic for test.example.com to a web server that has an IP address of
// 192.0.2.44. Deleting Resource Record Sets To delete a resource record set, you
// must specify all the same values that you specified when you created it. Change
// Batches and Transactional Changes The request body must include a document with
// a ChangeResourceRecordSetsRequest element. The request body contains a list of
// change items, known as a change batch. Change batches are considered
// transactional changes. Route 53 validates the changes in the request and then
// either makes all or none of the changes in the change batch request. This
// ensures that DNS routing isn't adversely affected by partial changes to the
// resource record sets in a hosted zone. For example, suppose a change batch
// request contains two changes: it deletes the CNAME resource record set for
// www.example.com and creates an alias resource record set for www.example.com. If
// validation for both records succeeds, Route 53 deletes the first resource record
// set and creates the second resource record set in a single operation. If
// validation for either the DELETE or the CREATE action fails, then the request is
// canceled, and the original CNAME record continues to exist. If you try to delete
// the same resource record set more than once in a single change batch, Route 53
// returns an InvalidChangeBatch error. Traffic Flow To create resource record sets
// for complex routing configurations, use either the traffic flow visual editor in
// the Route 53 console or the API actions for traffic policies and traffic policy
// instances. Save the configuration as a traffic policy, then associate the
// traffic policy with one or more domain names (such as example.com) or subdomain
// names (such as www.example.com), in the same hosted zone or in multiple hosted
// zones. You can roll back the updates if the new configuration isn't performing
// as expected. For more information, see Using Traffic Flow to Route DNS Traffic
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/traffic-flow.html) in
// the Amazon Route 53 Developer Guide. Create, Delete, and Upsert Use
// ChangeResourceRecordsSetsRequest to perform the following actions:
//
// * CREATE:
// Creates a resource record set that has the specified values.
//
// * DELETE: Deletes
// an existing resource record set that has the specified values.
//
// * UPSERT: If a
// resource record set does not already exist, AWS creates it. If a resource set
// does exist, Route 53 updates it with the values in the request.
//
// Syntaxes for
// Creating, Updating, and Deleting Resource Record Sets The syntax for a request
// depends on the type of resource record set that you want to create, delete, or
// update, such as weighted, alias, or failover. The XML elements in your request
// must appear in the order listed in the syntax. For an example for each type of
// resource record set, see "Examples." Don't refer to the syntax in the "Parameter
// Syntax" section, which includes all of the elements for every kind of resource
// record set that you can create, delete, or update by using
// ChangeResourceRecordSets. Change Propagation to Route 53 DNS Servers When you
// submit a ChangeResourceRecordSets request, Route 53 propagates your changes to
// all of the Route 53 authoritative DNS servers. While your changes are
// propagating, GetChange returns a status of PENDING. When propagation is
// complete, GetChange returns a status of INSYNC. Changes generally propagate to
// all Route 53 name servers within 60 seconds. For more information, see GetChange
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html).
// Limits on ChangeResourceRecordSets Requests For information about the limits on
// a ChangeResourceRecordSets request, see Limits
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html)
// in the Amazon Route 53 Developer Guide.
func (c *Client) ChangeResourceRecordSets(ctx context.Context, params *ChangeResourceRecordSetsInput, optFns ...func(*Options)) (*ChangeResourceRecordSetsOutput, error) {
	if params == nil {
		params = &ChangeResourceRecordSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeResourceRecordSets", params, optFns, addOperationChangeResourceRecordSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeResourceRecordSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains change information for the resource record set.
type ChangeResourceRecordSetsInput struct {

	// A complex type that contains an optional comment and the Changes element.
	//
	// This member is required.
	ChangeBatch *types.ChangeBatch

	// The ID of the hosted zone that contains the resource record sets that you want
	// to change.
	//
	// This member is required.
	HostedZoneId *string
}

// A complex type containing the response for the request.
type ChangeResourceRecordSetsOutput struct {

	// A complex type that contains information about changes made to your hosted zone.
	// This element contains an ID that you use when performing a GetChange
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html)
	// action to get detailed information about the change.
	//
	// This member is required.
	ChangeInfo *types.ChangeInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationChangeResourceRecordSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpChangeResourceRecordSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpChangeResourceRecordSets{}, middleware.After)
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
	if err = addOpChangeResourceRecordSetsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opChangeResourceRecordSets(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = route53cust.HandleCustomErrorDeserialization(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opChangeResourceRecordSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ChangeResourceRecordSets",
	}
}
