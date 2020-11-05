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

// Starts an on-demand backup job for the specified resource.
func (c *Client) StartBackupJob(ctx context.Context, params *StartBackupJobInput, optFns ...func(*Options)) (*StartBackupJobOutput, error) {
	if params == nil {
		params = &StartBackupJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartBackupJob", params, optFns, addOperationStartBackupJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartBackupJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartBackupJobInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	//
	// This member is required.
	IamRoleArn *string

	// An Amazon Resource Name (ARN) that uniquely identifies a resource. The format of
	// the ARN depends on the resource type.
	//
	// This member is required.
	ResourceArn *string

	// Specifies the backup option for a selected resource. This option is only
	// available for Windows VSS backup jobs. Valid values: Set to
	// "WindowsVSS”:“enabled" to enable WindowsVSS backup option and create a VSS
	// Windows backup. Set to “WindowsVSS”:”disabled” to create a regular backup. The
	// WindowsVSS option is not enabled by default.
	BackupOptions map[string]*string

	// A value in minutes after a backup job is successfully started before it must be
	// completed or it will be canceled by AWS Backup. This value is optional.
	CompleteWindowMinutes *int64

	// A customer chosen string that can be used to distinguish between calls to
	// StartBackupJob.
	IdempotencyToken *string

	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup will transition and expire backups automatically
	// according to the lifecycle that you define. Backups transitioned to cold storage
	// must be stored in cold storage for a minimum of 90 days. Therefore, the “expire
	// after days” setting must be 90 days greater than the “transition to cold after
	// days” setting. The “transition to cold after days” setting cannot be changed
	// after a backup has been transitioned to cold.
	Lifecycle *types.Lifecycle

	// To help organize your resources, you can assign your own metadata to the
	// resources that you create. Each tag is a key-value pair.
	RecoveryPointTags map[string]*string

	// A value in minutes after a backup is scheduled before a job will be canceled if
	// it doesn't start successfully. This value is optional.
	StartWindowMinutes *int64
}

type StartBackupJobOutput struct {

	// Uniquely identifies a request to AWS Backup to back up a resource.
	BackupJobId *string

	// The date and time that a backup job is started, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartBackupJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartBackupJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartBackupJob{}, middleware.After)
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
	if err = addOpStartBackupJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartBackupJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartBackupJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "StartBackupJob",
	}
}
