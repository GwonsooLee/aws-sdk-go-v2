// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the specified network interface attribute. You can specify only one
// attribute at a time. You can use this action to attach and detach security
// groups from an existing EC2 instance.
func (c *Client) ModifyNetworkInterfaceAttribute(ctx context.Context, params *ModifyNetworkInterfaceAttributeInput, optFns ...func(*Options)) (*ModifyNetworkInterfaceAttributeOutput, error) {
	if params == nil {
		params = &ModifyNetworkInterfaceAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyNetworkInterfaceAttribute", params, optFns, addOperationModifyNetworkInterfaceAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyNetworkInterfaceAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifyNetworkInterfaceAttribute.
type ModifyNetworkInterfaceAttributeInput struct {

	// The ID of the network interface.
	//
	// This member is required.
	NetworkInterfaceId *string

	// Information about the interface attachment. If modifying the 'delete on
	// termination' attribute, you must specify the ID of the interface attachment.
	Attachment *types.NetworkInterfaceAttachmentChanges

	// A description for the network interface.
	Description *types.AttributeValue

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Changes the security groups for the network interface. The new set of groups you
	// specify replaces the current set. You must specify at least one group, even if
	// it's just the default security group in the VPC. You must specify the ID of the
	// security group, not the name.
	Groups []*string

	// Indicates whether source/destination checking is enabled. A value of true means
	// checking is enabled, and false means checking is disabled. This value must be
	// false for a NAT instance to perform NAT. For more information, see NAT Instances
	// (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_NAT_Instance.html)
	// in the Amazon Virtual Private Cloud User Guide.
	SourceDestCheck *types.AttributeBooleanValue
}

type ModifyNetworkInterfaceAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyNetworkInterfaceAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyNetworkInterfaceAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyNetworkInterfaceAttribute{}, middleware.After)
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
	if err = addOpModifyNetworkInterfaceAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyNetworkInterfaceAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyNetworkInterfaceAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyNetworkInterfaceAttribute",
	}
}
