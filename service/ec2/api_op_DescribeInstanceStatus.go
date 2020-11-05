// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the status of the specified instances or all of your instances. By
// default, only running instances are described, unless you specifically indicate
// to return the status of all instances. Instance status includes the following
// components:
//
// * Status checks - Amazon EC2 performs status checks on running EC2
// instances to identify hardware and software issues. For more information, see
// Status checks for your instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-system-instance-status-check.html)
// and Troubleshooting instances with failed status checks
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/TroubleshootingInstances.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// * Scheduled events - Amazon EC2
// can schedule events (such as reboot, stop, or terminate) for your instances
// related to hardware issues, software updates, or system maintenance. For more
// information, see Scheduled events for your instances
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-instances-status-check_sched.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// * Instance state - You can
// manage your instances from the moment you launch them through their termination.
// For more information, see Instance lifecycle
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-lifecycle.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeInstanceStatus(ctx context.Context, params *DescribeInstanceStatusInput, optFns ...func(*Options)) (*DescribeInstanceStatusOutput, error) {
	if params == nil {
		params = &DescribeInstanceStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstanceStatus", params, optFns, addOperationDescribeInstanceStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstanceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstanceStatusInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	// * availability-zone - The Availability Zone of the instance.
	//
	// *
	// event.code - The code for the scheduled event (instance-reboot | system-reboot |
	// system-maintenance | instance-retirement | instance-stop).
	//
	// * event.description
	// - A description of the event.
	//
	// * event.instance-event-id - The ID of the event
	// whose date and time you are modifying.
	//
	// * event.not-after - The latest end time
	// for the scheduled event (for example, 2014-09-15T17:15:20.000Z).
	//
	// *
	// event.not-before - The earliest start time for the scheduled event (for example,
	// 2014-09-15T17:15:20.000Z).
	//
	// * event.not-before-deadline - The deadline for
	// starting the event (for example, 2014-09-15T17:15:20.000Z).
	//
	// *
	// instance-state-code - The code for the instance state, as a 16-bit unsigned
	// integer. The high byte is used for internal purposes and should be ignored. The
	// low byte is set based on the state represented. The valid values are 0
	// (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64 (stopping), and
	// 80 (stopped).
	//
	// * instance-state-name - The state of the instance (pending |
	// running | shutting-down | terminated | stopping | stopped).
	//
	// *
	// instance-status.reachability - Filters on instance status where the name is
	// reachability (passed | failed | initializing | insufficient-data).
	//
	// *
	// instance-status.status - The status of the instance (ok | impaired |
	// initializing | insufficient-data | not-applicable).
	//
	// *
	// system-status.reachability - Filters on system status where the name is
	// reachability (passed | failed | initializing | insufficient-data).
	//
	// *
	// system-status.status - The system status of the instance (ok | impaired |
	// initializing | insufficient-data | not-applicable).
	Filters []*types.Filter

	// When true, includes the health status for all instances. When false, includes
	// the health status for running instances only. Default: false
	IncludeAllInstances *bool

	// The instance IDs. Default: Describes all your instances. Constraints: Maximum
	// 100 explicitly specified instance IDs.
	InstanceIds []*string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter in the same call.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeInstanceStatusOutput struct {

	// Information about the status of the instances.
	InstanceStatuses []*types.InstanceStatus

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstanceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstanceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstanceStatus{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstanceStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeInstanceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstanceStatus",
	}
}
