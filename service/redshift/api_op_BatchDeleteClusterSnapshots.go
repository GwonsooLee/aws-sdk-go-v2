// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a set of cluster snapshots.
func (c *Client) BatchDeleteClusterSnapshots(ctx context.Context, params *BatchDeleteClusterSnapshotsInput, optFns ...func(*Options)) (*BatchDeleteClusterSnapshotsOutput, error) {
	if params == nil {
		params = &BatchDeleteClusterSnapshotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteClusterSnapshots", params, optFns, addOperationBatchDeleteClusterSnapshotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteClusterSnapshotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteClusterSnapshotsInput struct {

	// A list of identifiers for the snapshots that you want to delete.
	//
	// This member is required.
	Identifiers []*types.DeleteClusterSnapshotMessage
}

type BatchDeleteClusterSnapshotsOutput struct {

	// A list of any errors returned.
	Errors []*types.SnapshotErrorMessage

	// A list of the snapshot identifiers that were deleted.
	Resources []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteClusterSnapshotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpBatchDeleteClusterSnapshots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpBatchDeleteClusterSnapshots{}, middleware.After)
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
	if err = addOpBatchDeleteClusterSnapshotsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteClusterSnapshots(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDeleteClusterSnapshots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "BatchDeleteClusterSnapshots",
	}
}
