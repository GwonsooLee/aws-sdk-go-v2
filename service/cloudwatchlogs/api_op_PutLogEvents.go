// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Uploads a batch of log events to the specified log stream. You must include the
// sequence token obtained from the response of the previous call. An upload in a
// newly created log stream does not require a sequence token. You can also get the
// sequence token in the expectedSequenceToken field from
// InvalidSequenceTokenException. If you call PutLogEvents twice within a narrow
// time period using the same value for sequenceToken, both calls might be
// successful or one might be rejected. The batch of events must satisfy the
// following constraints:
//
// * The maximum batch size is 1,048,576 bytes. This size
// is calculated as the sum of all event messages in UTF-8, plus 26 bytes for each
// log event.
//
// * None of the log events in the batch can be more than 2 hours in
// the future.
//
// * None of the log events in the batch can be older than 14 days or
// older than the retention period of the log group.
//
// * The log events in the batch
// must be in chronological order by their timestamp. The timestamp is the time the
// event occurred, expressed as the number of milliseconds after Jan 1, 1970
// 00:00:00 UTC. (In AWS Tools for PowerShell and the AWS SDK for .NET, the
// timestamp is specified in .NET format: yyyy-mm-ddThh:mm:ss. For example,
// 2017-09-15T13:45:30.)
//
// * A batch of log events in a single request cannot span
// more than 24 hours. Otherwise, the operation fails.
//
// * The maximum number of log
// events in a batch is 10,000.
//
// * There is a quota of 5 requests per second per
// log stream. Additional requests are throttled. This quota can't be changed.
//
// If
// a call to PutLogEvents returns "UnrecognizedClientException" the most likely
// cause is an invalid AWS access key ID or secret key.
func (c *Client) PutLogEvents(ctx context.Context, params *PutLogEventsInput, optFns ...func(*Options)) (*PutLogEventsOutput, error) {
	if params == nil {
		params = &PutLogEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutLogEvents", params, optFns, addOperationPutLogEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutLogEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutLogEventsInput struct {

	// The log events.
	//
	// This member is required.
	LogEvents []*types.InputLogEvent

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// The name of the log stream.
	//
	// This member is required.
	LogStreamName *string

	// The sequence token obtained from the response of the previous PutLogEvents call.
	// An upload in a newly created log stream does not require a sequence token. You
	// can also get the sequence token using DescribeLogStreams
	// (https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeLogStreams.html).
	// If you call PutLogEvents twice within a narrow time period using the same value
	// for sequenceToken, both calls might be successful or one might be rejected.
	SequenceToken *string
}

type PutLogEventsOutput struct {

	// The next sequence token.
	NextSequenceToken *string

	// The rejected events.
	RejectedLogEventsInfo *types.RejectedLogEventsInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutLogEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutLogEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutLogEvents{}, middleware.After)
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
	if err = addOpPutLogEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutLogEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutLogEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "PutLogEvents",
	}
}
