// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of target IDs that are associated a deployment.
func (c *Client) ListDeploymentTargets(ctx context.Context, params *ListDeploymentTargetsInput, optFns ...func(*Options)) (*ListDeploymentTargetsOutput, error) {
	if params == nil {
		params = &ListDeploymentTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeploymentTargets", params, optFns, addOperationListDeploymentTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentTargetsInput struct {

	// The unique ID of a deployment.
	DeploymentId *string

	// A token identifier returned from the previous ListDeploymentTargets call. It can
	// be used to return the next set of deployment targets in the list.
	NextToken *string

	// A key used to filter the returned targets. The two valid values are:
	//
	// *
	// TargetStatus - A TargetStatus filter string can be Failed, InProgress, Pending,
	// Ready, Skipped, Succeeded, or Unknown.
	//
	// * ServerInstanceLabel - A
	// ServerInstanceLabel filter string can be Blue or Green.
	TargetFilters map[string][]*string
}

type ListDeploymentTargetsOutput struct {

	// If a large amount of information is returned, a token identifier is also
	// returned. It can be used in a subsequent ListDeploymentTargets call to return
	// the next set of deployment targets in the list.
	NextToken *string

	// The unique IDs of deployment targets.
	TargetIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeploymentTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentTargets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDeploymentTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListDeploymentTargets",
	}
}
