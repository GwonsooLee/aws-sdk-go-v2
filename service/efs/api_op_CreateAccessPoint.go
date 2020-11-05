// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an EFS access point. An access point is an application-specific view
// into an EFS file system that applies an operating system user and group, and a
// file system path, to any file system request made through the access point. The
// operating system user and group override any identity information provided by
// the NFS client. The file system path is exposed as the access point's root
// directory. Applications using the access point can only access data in its own
// directory and below. To learn more, see Mounting a File System Using EFS Access
// Points (https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html). This
// operation requires permissions for the elasticfilesystem:CreateAccessPoint
// action.
func (c *Client) CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) {
	if params == nil {
		params = &CreateAccessPointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessPoint", params, optFns, addOperationCreateAccessPointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPointInput struct {

	// A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent
	// creation.
	//
	// This member is required.
	ClientToken *string

	// The ID of the EFS file system that the access point provides access to.
	//
	// This member is required.
	FileSystemId *string

	// The operating system user and group applied to all file system requests made
	// using the access point.
	PosixUser *types.PosixUser

	// Specifies the directory on the Amazon EFS file system that the access point
	// exposes as the root directory of your file system to NFS clients using the
	// access point. The clients using the access point can only access the root
	// directory and below. If the RootDirectory > Path specified does not exist, EFS
	// creates it and applies the CreationInfo settings when a client connects to an
	// access point. When specifying a RootDirectory, you need to provide the Path, and
	// the CreationInfo is optional.
	RootDirectory *types.RootDirectory

	// Creates tags associated with the access point. Each tag is a key-value pair.
	Tags []*types.Tag
}

// Provides a description of an EFS file system access point.
type CreateAccessPointOutput struct {

	// The unique Amazon Resource Name (ARN) associated with the access point.
	AccessPointArn *string

	// The ID of the access point, assigned by Amazon EFS.
	AccessPointId *string

	// The opaque string specified in the request to ensure idempotent creation.
	ClientToken *string

	// The ID of the EFS file system that the access point applies to.
	FileSystemId *string

	// Identifies the lifecycle phase of the access point.
	LifeCycleState types.LifeCycleState

	// The name of the access point. This is the value of the Name tag.
	Name *string

	// Identified the AWS account that owns the access point resource.
	OwnerId *string

	// The full POSIX identity, including the user ID, group ID, and secondary group
	// IDs on the access point that is used for all file operations by NFS clients
	// using the access point.
	PosixUser *types.PosixUser

	// The directory on the Amazon EFS file system that the access point exposes as the
	// root directory to NFS clients using the access point.
	RootDirectory *types.RootDirectory

	// The tags associated with the access point, presented as an array of Tag objects.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAccessPointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessPoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessPoint{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAccessPointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAccessPointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPoint(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAccessPoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessPoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessPointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAccessPointMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessPoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessPoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "CreateAccessPoint",
	}
}
