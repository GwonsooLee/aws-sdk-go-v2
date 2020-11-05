// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures an event selector for your trail. Use event selectors to further
// specify the management and data event settings for your trail. By default,
// trails created without specific event selectors will be configured to log all
// read and write management events, and no data events. When an event occurs in
// your account, CloudTrail evaluates the event selectors in all trails. For each
// trail, if the event matches any event selector, the trail processes and logs the
// event. If the event doesn't match any event selector, the trail doesn't log the
// event. Example
//
// * You create an event selector for a trail and specify that you
// want write-only events.
//
// * The EC2 GetConsoleOutput and RunInstances API
// operations occur in your account.
//
// * CloudTrail evaluates whether the events
// match your event selectors.
//
// * The RunInstances is a write-only event and it
// matches your event selector. The trail logs the event.
//
// * The GetConsoleOutput
// is a read-only event but it doesn't match your event selector. The trail doesn't
// log the event.
//
// The PutEventSelectors operation must be called from the region
// in which the trail was created; otherwise, an InvalidHomeRegionException is
// thrown. You can configure up to five event selectors for each trail. For more
// information, see Logging Data and Management Events for Trails
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-management-and-data-events-with-cloudtrail.html)
// and Limits in AWS CloudTrail
// (https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html)
// in the AWS CloudTrail User Guide.
func (c *Client) PutEventSelectors(ctx context.Context, params *PutEventSelectorsInput, optFns ...func(*Options)) (*PutEventSelectorsOutput, error) {
	if params == nil {
		params = &PutEventSelectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEventSelectors", params, optFns, addOperationPutEventSelectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEventSelectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutEventSelectorsInput struct {

	// Specifies the settings for your event selectors. You can configure up to five
	// event selectors for a trail.
	//
	// This member is required.
	EventSelectors []*types.EventSelector

	// Specifies the name of the trail or trail ARN. If you specify a trail name, the
	// string must meet the following requirements:
	//
	// * Contain only ASCII letters (a-z,
	// A-Z), numbers (0-9), periods (.), underscores (_), or dashes (-)
	//
	// * Start with a
	// letter or number, and end with a letter or number
	//
	// * Be between 3 and 128
	// characters
	//
	// * Have no adjacent periods, underscores or dashes. Names like
	// my-_namespace and my--namespace are invalid.
	//
	// * Not be in IP address format (for
	// example, 192.168.5.4)
	//
	// If you specify a trail ARN, it must be in the format:
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// This member is required.
	TrailName *string
}

type PutEventSelectorsOutput struct {

	// Specifies the event selectors configured for your trail.
	EventSelectors []*types.EventSelector

	// Specifies the ARN of the trail that was updated with event selectors. The format
	// of a trail ARN is: arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	TrailARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutEventSelectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutEventSelectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutEventSelectors{}, middleware.After)
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
	if err = addOpPutEventSelectorsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEventSelectors(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEventSelectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "PutEventSelectors",
	}
}
