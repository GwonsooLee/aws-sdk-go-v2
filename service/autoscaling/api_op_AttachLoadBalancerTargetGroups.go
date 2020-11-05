// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches one or more target groups to the specified Auto Scaling group. To
// describe the target groups for an Auto Scaling group, call the
// DescribeLoadBalancerTargetGroups API. To detach the target group from the Auto
// Scaling group, call the DetachLoadBalancerTargetGroups API. With Application
// Load Balancers and Network Load Balancers, instances are registered as targets
// with a target group. With Classic Load Balancers, instances are registered with
// the load balancer. For more information, see Attaching a Load Balancer to Your
// Auto Scaling Group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-load-balancer-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) AttachLoadBalancerTargetGroups(ctx context.Context, params *AttachLoadBalancerTargetGroupsInput, optFns ...func(*Options)) (*AttachLoadBalancerTargetGroupsOutput, error) {
	if params == nil {
		params = &AttachLoadBalancerTargetGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachLoadBalancerTargetGroups", params, optFns, addOperationAttachLoadBalancerTargetGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachLoadBalancerTargetGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachLoadBalancerTargetGroupsInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The Amazon Resource Names (ARN) of the target groups. You can specify up to 10
	// target groups.
	//
	// This member is required.
	TargetGroupARNs []*string
}

type AttachLoadBalancerTargetGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachLoadBalancerTargetGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachLoadBalancerTargetGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachLoadBalancerTargetGroups{}, middleware.After)
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
	if err = addOpAttachLoadBalancerTargetGroupsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachLoadBalancerTargetGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachLoadBalancerTargetGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachLoadBalancerTargetGroups",
	}
}
