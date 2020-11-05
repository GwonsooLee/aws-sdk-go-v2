// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a cached volume on a specified cached volume gateway. This operation is
// only supported in the cached volume gateway type. Cache storage must be
// allocated to the gateway before you can create a cached volume. Use the AddCache
// operation to add cache storage to a gateway. In the request, you must specify
// the gateway, size of the volume in bytes, the iSCSI target name, an IP address
// on which to expose the target, and a unique client token. In response, the
// gateway creates the volume and returns information about it. This information
// includes the volume Amazon Resource Name (ARN), its size, and the iSCSI target
// ARN that initiators can use to connect to the volume target. Optionally, you can
// provide the ARN for an existing volume as the SourceVolumeARN for this cached
// volume, which creates an exact copy of the existing volume’s latest recovery
// point. The VolumeSizeInBytes value must be equal to or larger than the size of
// the copied volume, in bytes.
func (c *Client) CreateCachediSCSIVolume(ctx context.Context, params *CreateCachediSCSIVolumeInput, optFns ...func(*Options)) (*CreateCachediSCSIVolumeOutput, error) {
	if params == nil {
		params = &CreateCachediSCSIVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCachediSCSIVolume", params, optFns, addOperationCreateCachediSCSIVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateCachediSCSIVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateCachediSCSIVolumeInput struct {

	// A unique identifier that you use to retry a request. If you retry a request, use
	// the same ClientToken you specified in the initial request.
	//
	// This member is required.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// The network interface of the gateway on which to expose the iSCSI target. Only
	// IPv4 addresses are accepted. Use DescribeGatewayInformation to get a list of the
	// network interfaces available on a gateway. Valid Values: A valid IP address.
	//
	// This member is required.
	NetworkInterfaceId *string

	// The name of the iSCSI target used by an initiator to connect to a volume and
	// used as a suffix for the target ARN. For example, specifying TargetName as
	// myvolume results in the target ARN of
	// arn:aws:storagegateway:us-east-2:111122223333:gateway/sgw-12A3456B/target/iqn.1997-05.com.amazon:myvolume.
	// The target name must be unique across all volumes on a gateway. If you don't
	// specify a value, Storage Gateway uses the value that was previously used for
	// this volume as the new target name.
	//
	// This member is required.
	TargetName *string

	// The size of the volume in bytes.
	//
	// This member is required.
	VolumeSizeInBytes *int64

	// Set to true to use Amazon S3 server-side encryption with your own AWS KMS key,
	// or false to use a key managed by Amazon S3. Optional. Valid Values: true | false
	KMSEncrypted *bool

	// The Amazon Resource Name (ARN) of a symmetric customer master key (CMK) used for
	// Amazon S3 server-side encryption. Storage Gateway does not support asymmetric
	// CMKs. This value can only be set when KMSEncrypted is true. Optional.
	KMSKey *string

	// The snapshot ID (e.g. "snap-1122aabb") of the snapshot to restore as the new
	// cached volume. Specify this field if you want to create the iSCSI storage volume
	// from a snapshot; otherwise, do not include this field. To list snapshots for
	// your account use DescribeSnapshots
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/ApiReference-query-DescribeSnapshots.html)
	// in the Amazon Elastic Compute Cloud API Reference.
	SnapshotId *string

	// The ARN for an existing volume. Specifying this ARN makes the new volume into an
	// exact copy of the specified existing volume's latest recovery point. The
	// VolumeSizeInBytes value for this new volume must be equal to or larger than the
	// size of the existing volume, in bytes.
	SourceVolumeARN *string

	// A list of up to 50 tags that you can assign to a cached volume. Each tag is a
	// key-value pair. Valid characters for key and value are letters, spaces, and
	// numbers that you can represent in UTF-8 format, and the following special
	// characters: + - = . _ : / @. The maximum length of a tag's key is 128
	// characters, and the maximum length for a tag's value is 256 characters.
	Tags []*types.Tag
}

type CreateCachediSCSIVolumeOutput struct {

	// The Amazon Resource Name (ARN) of the volume target, which includes the iSCSI
	// name that initiators can use to connect to the target.
	TargetARN *string

	// The Amazon Resource Name (ARN) of the configured volume.
	VolumeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateCachediSCSIVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCachediSCSIVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCachediSCSIVolume{}, middleware.After)
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
	if err = addOpCreateCachediSCSIVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCachediSCSIVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCachediSCSIVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "CreateCachediSCSIVolume",
	}
}
