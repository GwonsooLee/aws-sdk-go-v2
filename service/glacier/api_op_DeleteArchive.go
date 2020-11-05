// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation deletes an archive from a vault. Subsequent requests to initiate
// a retrieval of this archive will fail. Archive retrievals that are in progress
// for this archive ID may or may not succeed according to the following
// scenarios:
//
// * If the archive retrieval job is actively preparing the data for
// download when Amazon S3 Glacier receives the delete archive request, the
// archival retrieval operation might fail.
//
// * If the archive retrieval job has
// successfully prepared the archive for download when Amazon S3 Glacier receives
// the delete archive request, you will be able to download the output.
//
// This
// operation is idempotent. Attempting to delete an already-deleted archive does
// not result in an error. An AWS account has full permission to perform all
// operations (actions). However, AWS Identity and Access Management (IAM) users
// don't have any permissions by default. You must grant them explicit permission
// to perform specific actions. For more information, see Access Control Using AWS
// Identity and Access Management (IAM)
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
// For conceptual information and underlying REST API, see Deleting an Archive in
// Amazon Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/deleting-an-archive.html)
// and Delete Archive
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-archive-delete.html)
// in the Amazon Glacier Developer Guide.
func (c *Client) DeleteArchive(ctx context.Context, params *DeleteArchiveInput, optFns ...func(*Options)) (*DeleteArchiveOutput, error) {
	if params == nil {
		params = &DeleteArchiveInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteArchive", params, optFns, addOperationDeleteArchiveMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteArchiveOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for deleting an archive from an Amazon S3 Glacier vault.
type DeleteArchiveInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The ID of the archive to delete.
	//
	// This member is required.
	ArchiveId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

type DeleteArchiveOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteArchiveMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteArchive{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteArchive{}, middleware.After)
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
	if err = addOpDeleteArchiveValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteArchive(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddTreeHashMiddleware(stack); err != nil {
		return err
	}
	if err = glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion); err != nil {
		return err
	}
	if err = glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeleteArchive(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "DeleteArchive",
	}
}
