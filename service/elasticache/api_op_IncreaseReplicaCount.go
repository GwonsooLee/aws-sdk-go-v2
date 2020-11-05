// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Dynamically increases the number of replics in a Redis (cluster mode disabled)
// replication group or the number of replica nodes in one or more node groups
// (shards) of a Redis (cluster mode enabled) replication group. This operation is
// performed with no cluster down time.
func (c *Client) IncreaseReplicaCount(ctx context.Context, params *IncreaseReplicaCountInput, optFns ...func(*Options)) (*IncreaseReplicaCountOutput, error) {
	if params == nil {
		params = &IncreaseReplicaCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IncreaseReplicaCount", params, optFns, addOperationIncreaseReplicaCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*IncreaseReplicaCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IncreaseReplicaCountInput struct {

	// If True, the number of replica nodes is increased immediately.
	// ApplyImmediately=False is not currently supported.
	//
	// This member is required.
	ApplyImmediately *bool

	// The id of the replication group to which you want to add replica nodes.
	//
	// This member is required.
	ReplicationGroupId *string

	// The number of read replica nodes you want at the completion of this operation.
	// For Redis (cluster mode disabled) replication groups, this is the number of
	// replica nodes in the replication group. For Redis (cluster mode enabled)
	// replication groups, this is the number of replica nodes in each of the
	// replication group's node groups.
	NewReplicaCount *int32

	// A list of ConfigureShard objects that can be used to configure each shard in a
	// Redis (cluster mode enabled) replication group. The ConfigureShard has three
	// members: NewReplicaCount, NodeGroupId, and PreferredAvailabilityZones.
	ReplicaConfiguration []*types.ConfigureShard
}

type IncreaseReplicaCountOutput struct {

	// Contains all of the attributes of a specific Redis replication group.
	ReplicationGroup *types.ReplicationGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationIncreaseReplicaCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpIncreaseReplicaCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpIncreaseReplicaCount{}, middleware.After)
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
	if err = addOpIncreaseReplicaCountValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opIncreaseReplicaCount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opIncreaseReplicaCount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "IncreaseReplicaCount",
	}
}
