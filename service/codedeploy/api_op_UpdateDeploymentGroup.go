// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes information about a deployment group.
func (c *Client) UpdateDeploymentGroup(ctx context.Context, params *UpdateDeploymentGroupInput, optFns ...func(*Options)) (*UpdateDeploymentGroupOutput, error) {
	if params == nil {
		params = &UpdateDeploymentGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDeploymentGroup", params, optFns, addOperationUpdateDeploymentGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDeploymentGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateDeploymentGroup operation.
type UpdateDeploymentGroupInput struct {

	// The application name that corresponds to the deployment group to update.
	//
	// This member is required.
	ApplicationName *string

	// The current name of the deployment group.
	//
	// This member is required.
	CurrentDeploymentGroupName *string

	// Information to add or change about Amazon CloudWatch alarms when the deployment
	// group is updated.
	AlarmConfiguration *types.AlarmConfiguration

	// Information for an automatic rollback configuration that is added or changed
	// when a deployment group is updated.
	AutoRollbackConfiguration *types.AutoRollbackConfiguration

	// The replacement list of Auto Scaling groups to be included in the deployment
	// group, if you want to change them. To keep the Auto Scaling groups, enter their
	// names. To remove Auto Scaling groups, do not enter any Auto Scaling group names.
	AutoScalingGroups []*string

	// Information about blue/green deployment options for a deployment group.
	BlueGreenDeploymentConfiguration *types.BlueGreenDeploymentConfiguration

	// The replacement deployment configuration name to use, if you want to change it.
	DeploymentConfigName *string

	// Information about the type of deployment, either in-place or blue/green, you
	// want to run and whether to route deployment traffic behind a load balancer.
	DeploymentStyle *types.DeploymentStyle

	// The replacement set of Amazon EC2 tags on which to filter, if you want to change
	// them. To keep the existing tags, enter their names. To remove tags, do not enter
	// any tag names.
	Ec2TagFilters []*types.EC2TagFilter

	// Information about groups of tags applied to on-premises instances. The
	// deployment group includes only EC2 instances identified by all the tag groups.
	Ec2TagSet *types.EC2TagSet

	// The target Amazon ECS services in the deployment group. This applies only to
	// deployment groups that use the Amazon ECS compute platform. A target Amazon ECS
	// service is specified as an Amazon ECS cluster and service name pair using the
	// format :.
	EcsServices []*types.ECSService

	// Information about the load balancer used in a deployment.
	LoadBalancerInfo *types.LoadBalancerInfo

	// The new name of the deployment group, if you want to change it.
	NewDeploymentGroupName *string

	// The replacement set of on-premises instance tags on which to filter, if you want
	// to change them. To keep the existing tags, enter their names. To remove tags, do
	// not enter any tag names.
	OnPremisesInstanceTagFilters []*types.TagFilter

	// Information about an on-premises instance tag set. The deployment group includes
	// only on-premises instances identified by all the tag groups.
	OnPremisesTagSet *types.OnPremisesTagSet

	// A replacement ARN for the service role, if you want to change it.
	ServiceRoleArn *string

	// Information about triggers to change when the deployment group is updated. For
	// examples, see Edit a Trigger in a CodeDeploy Deployment Group
	// (https://docs.aws.amazon.com/codedeploy/latest/userguide/how-to-notify-edit.html)
	// in the AWS CodeDeploy User Guide.
	TriggerConfigurations []*types.TriggerConfig
}

// Represents the output of an UpdateDeploymentGroup operation.
type UpdateDeploymentGroupOutput struct {

	// If the output contains no data, and the corresponding deployment group contained
	// at least one Auto Scaling group, AWS CodeDeploy successfully removed all
	// corresponding Auto Scaling lifecycle event hooks from the AWS account. If the
	// output contains data, AWS CodeDeploy could not remove some Auto Scaling
	// lifecycle event hooks from the AWS account.
	HooksNotCleanedUp []*types.AutoScalingGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDeploymentGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDeploymentGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDeploymentGroup{}, middleware.After)
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
	if err = addOpUpdateDeploymentGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDeploymentGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDeploymentGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "UpdateDeploymentGroup",
	}
}
