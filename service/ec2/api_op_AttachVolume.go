// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Attaches an EBS volume to a running or stopped instance and exposes it to the
// instance with the specified device name. Encrypted EBS volumes must be attached
// to instances that support Amazon EBS encryption. For more information, see
// Amazon EBS Encryption
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html) in the
// Amazon Elastic Compute Cloud User Guide. After you attach an EBS volume, you
// must make it available. For more information, see Making an EBS volume available
// for use
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-using-volumes.html). If
// a volume has an AWS Marketplace product code:
//
// * The volume can be attached only
// to a stopped instance.
//
// * AWS Marketplace product codes are copied from the
// volume to the instance.
//
// * You must be subscribed to the product.
//
// * The
// instance type and operating system of the instance must support the product. For
// example, you can't detach a volume from a Windows instance and attach it to a
// Linux instance.
//
// For more information, see Attaching Amazon EBS volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-attaching-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) AttachVolume(ctx context.Context, params *AttachVolumeInput, optFns ...func(*Options)) (*AttachVolumeOutput, error) {
	if params == nil {
		params = &AttachVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachVolume", params, optFns, addOperationAttachVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachVolumeInput struct {

	// The device name (for example, /dev/sdh or xvdh).
	//
	// This member is required.
	Device *string

	// The ID of the instance.
	//
	// This member is required.
	InstanceId *string

	// The ID of the EBS volume. The volume and instance must be within the same
	// Availability Zone.
	//
	// This member is required.
	VolumeId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

// Describes volume attachment details.
type AttachVolumeOutput struct {

	// The time stamp when the attachment initiated.
	AttachTime *time.Time

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool

	// The device name.
	Device *string

	// The ID of the instance.
	InstanceId *string

	// The attachment state of the volume.
	State types.VolumeAttachmentState

	// The ID of the volume.
	VolumeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAttachVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpAttachVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpAttachVolume{}, middleware.After)
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
	if err = addOpAttachVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AttachVolume",
	}
}
