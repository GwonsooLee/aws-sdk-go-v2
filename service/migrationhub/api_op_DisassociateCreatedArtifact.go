// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a created artifact of an AWS resource with a migration task
// performed by a migration tool that was previously associated. This API has the
// following traits:
//
// * A migration user can call the DisassociateCreatedArtifacts
// operation to disassociate a created AWS Artifact from a migration task.
//
// * The
// created artifact name must be provided in ARN (Amazon Resource Name) format
// which will contain information about type and region; for example:
// arn:aws:ec2:us-east-1:488216288981:image/ami-6d0ba87b.
//
// * Examples of the AWS
// resource behind the created artifact are, AMI's, EC2 instance, or RDS instance,
// etc.
func (c *Client) DisassociateCreatedArtifact(ctx context.Context, params *DisassociateCreatedArtifactInput, optFns ...func(*Options)) (*DisassociateCreatedArtifactOutput, error) {
	if params == nil {
		params = &DisassociateCreatedArtifactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateCreatedArtifact", params, optFns, addOperationDisassociateCreatedArtifactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateCreatedArtifactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateCreatedArtifactInput struct {

	// An ARN of the AWS resource related to the migration (e.g., AMI, EC2 instance,
	// RDS instance, etc.)
	//
	// This member is required.
	CreatedArtifactName *string

	// Unique identifier that references the migration task to be disassociated with
	// the artifact. Do not store personal data in this field.
	//
	// This member is required.
	MigrationTaskName *string

	// The name of the ProgressUpdateStream.
	//
	// This member is required.
	ProgressUpdateStream *string

	// Optional boolean flag to indicate whether any effect should take place. Used to
	// test if the caller has permission to make the call.
	DryRun *bool
}

type DisassociateCreatedArtifactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateCreatedArtifactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateCreatedArtifact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateCreatedArtifact{}, middleware.After)
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
	if err = addOpDisassociateCreatedArtifactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateCreatedArtifact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateCreatedArtifact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgh",
		OperationName: "DisassociateCreatedArtifact",
	}
}
