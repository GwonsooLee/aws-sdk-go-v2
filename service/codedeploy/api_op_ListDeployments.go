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

// Lists the deployments in a deployment group for an application registered with
// the IAM user or AWS account.
func (c *Client) ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeployments", params, optFns, addOperationListDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListDeployments operation.
type ListDeploymentsInput struct {

	// The name of an AWS CodeDeploy application associated with the IAM user or AWS
	// account. If applicationName is specified, then deploymentGroupName must be
	// specified. If it is not specified, then deploymentGroupName must not be
	// specified.
	ApplicationName *string

	// A time range (start and end) for returning a subset of the list of deployments.
	CreateTimeRange *types.TimeRange

	// The name of a deployment group for the specified application. If
	// deploymentGroupName is specified, then applicationName must be specified. If it
	// is not specified, then applicationName must not be specified.
	DeploymentGroupName *string

	// The unique ID of an external resource for returning deployments linked to the
	// external resource.
	ExternalId *string

	// A subset of deployments to list by status:
	//
	// * Created: Include created
	// deployments in the resulting list.
	//
	// * Queued: Include queued deployments in the
	// resulting list.
	//
	// * In Progress: Include in-progress deployments in the resulting
	// list.
	//
	// * Succeeded: Include successful deployments in the resulting list.
	//
	// *
	// Failed: Include failed deployments in the resulting list.
	//
	// * Stopped: Include
	// stopped deployments in the resulting list.
	IncludeOnlyStatuses []types.DeploymentStatus

	// An identifier returned from the previous list deployments call. It can be used
	// to return the next set of deployments in the list.
	NextToken *string
}

// Represents the output of a ListDeployments operation.
type ListDeploymentsOutput struct {

	// A list of deployment IDs.
	Deployments []*string

	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent list deployments call to return the next set of
	// deployments in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeployments{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeployments(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListDeployments",
	}
}
