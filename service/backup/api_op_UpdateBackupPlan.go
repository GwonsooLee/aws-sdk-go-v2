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

// Updates an existing backup plan identified by its backupPlanId with the input
// document in JSON format. The new version is uniquely identified by a VersionId.
func (c *Client) UpdateBackupPlan(ctx context.Context, params *UpdateBackupPlanInput, optFns ...func(*Options)) (*UpdateBackupPlanOutput, error) {
	if params == nil {
		params = &UpdateBackupPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateBackupPlan", params, optFns, addOperationUpdateBackupPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateBackupPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBackupPlanInput struct {

	// Specifies the body of a backup plan. Includes a BackupPlanName and one or more
	// sets of Rules.
	//
	// This member is required.
	BackupPlan *types.BackupPlanInput

	// Uniquely identifies a backup plan.
	//
	// This member is required.
	BackupPlanId *string
}

type UpdateBackupPlanOutput struct {

	// Contains a list of BackupOptions for each resource type.
	AdvancedBackupSettings []*types.AdvancedBackupSetting

	// An Amazon Resource Name (ARN) that uniquely identifies a backup plan; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:plan:8F81F553-3A74-4A3F-B93D-B3360DC80C50.
	BackupPlanArn *string

	// Uniquely identifies a backup plan.
	BackupPlanId *string

	// The date and time a backup plan is updated, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Unique, randomly generated, Unicode, UTF-8 encoded strings that are at most
	// 1,024 bytes long. Version Ids cannot be edited.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateBackupPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBackupPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBackupPlan{}, middleware.After)
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
	if err = addOpUpdateBackupPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBackupPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateBackupPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "UpdateBackupPlan",
	}
}
