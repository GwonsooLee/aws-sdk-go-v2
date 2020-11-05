// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Batch operation to revoke permissions from the principal.
func (c *Client) BatchRevokePermissions(ctx context.Context, params *BatchRevokePermissionsInput, optFns ...func(*Options)) (*BatchRevokePermissionsOutput, error) {
	if params == nil {
		params = &BatchRevokePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchRevokePermissions", params, optFns, addOperationBatchRevokePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchRevokePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchRevokePermissionsInput struct {

	// A list of up to 20 entries for resource permissions to be revoked by batch
	// operation to the principal.
	//
	// This member is required.
	Entries []*types.BatchPermissionsRequestEntry

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string
}

type BatchRevokePermissionsOutput struct {

	// A list of failures to revoke permissions to the resources.
	Failures []*types.BatchPermissionsFailureEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchRevokePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchRevokePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchRevokePermissions{}, middleware.After)
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
	if err = addOpBatchRevokePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchRevokePermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchRevokePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "BatchRevokePermissions",
	}
}
