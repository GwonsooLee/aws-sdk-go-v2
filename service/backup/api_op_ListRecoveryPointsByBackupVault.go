// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns detailed information about the recovery points stored in a backup vault.
func (c *Client) ListRecoveryPointsByBackupVault(ctx context.Context, params *ListRecoveryPointsByBackupVaultInput, optFns ...func(*Options)) (*ListRecoveryPointsByBackupVaultOutput, error) {
	if params == nil {
		params = &ListRecoveryPointsByBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecoveryPointsByBackupVault", params, optFns, addOperationListRecoveryPointsByBackupVaultMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecoveryPointsByBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecoveryPointsByBackupVaultInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// Returns only recovery points that match the specified backup plan ID.
	ByBackupPlanId *string

	// Returns only recovery points that were created after the specified timestamp.
	ByCreatedAfter *time.Time

	// Returns only recovery points that were created before the specified timestamp.
	ByCreatedBefore *time.Time

	// Returns only recovery points that match the specified resource Amazon Resource
	// Name (ARN).
	ByResourceArn *string

	// Returns only recovery points that match the specified resource type.
	ByResourceType *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
}

type ListRecoveryPointsByBackupVaultOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// An array of objects that contain detailed information about recovery points
	// saved in a backup vault.
	RecoveryPoints []*types.RecoveryPointByBackupVault

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRecoveryPointsByBackupVaultMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecoveryPointsByBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecoveryPointsByBackupVault{}, middleware.After)
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
	if err = addOpListRecoveryPointsByBackupVaultValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecoveryPointsByBackupVault(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRecoveryPointsByBackupVault(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListRecoveryPointsByBackupVault",
	}
}
