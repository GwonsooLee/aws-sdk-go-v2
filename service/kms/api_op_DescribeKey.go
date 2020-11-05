// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides detailed information about a customer master key (CMK). You can run
// DescribeKey on a customer managed CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk)
// or an AWS managed CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk).
// This detailed information includes the key ARN, creation date (and deletion
// date, if applicable), the key state, and the origin and expiration date (if any)
// of the key material. For CMKs in custom key stores, it includes information
// about the custom key store, such as the key store ID and the AWS CloudHSM
// cluster ID. It includes fields, like KeySpec, that help you distinguish
// symmetric from asymmetric CMKs. It also provides information that is
// particularly important to asymmetric CMKs, such as the key usage (encryption or
// signing) and the encryption algorithms or signing algorithms that the CMK
// supports. DescribeKey does not return the following information:
//
// * Aliases
// associated with the CMK. To get this information, use ListAliases.
//
// * Whether
// automatic key rotation is enabled on the CMK. To get this information, use
// GetKeyRotationStatus. Also, some key states prevent a CMK from being
// automatically rotated. For details, see How Automatic Key Rotation Works
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html#rotate-keys-how-it-works)
// in AWS Key Management Service Developer Guide.
//
// * Tags on the CMK. To get this
// information, use ListResourceTags.
//
// * Key policies and grants on the CMK. To get
// this information, use GetKeyPolicy and ListGrants.
//
// If you call the DescribeKey
// operation on a predefined AWS alias, that is, an AWS alias with no key ID, AWS
// KMS creates an AWS managed CMK
// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys).
// Then, it associates the alias with the new CMK, and returns the KeyId and Arn of
// the new CMK in the response. To perform this operation on a CMK in a different
// AWS account, specify the key ARN or alias ARN in the value of the KeyId
// parameter.
func (c *Client) DescribeKey(ctx context.Context, params *DescribeKeyInput, optFns ...func(*Options)) (*DescribeKeyOutput, error) {
	if params == nil {
		params = &DescribeKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeKey", params, optFns, addOperationDescribeKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeKeyInput struct {

	// Describes the specified customer master key (CMK). If you specify a predefined
	// AWS alias (an AWS alias with no key ID), KMS associates the alias with an AWS
	// managed CMK
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#master_keys)
	// and returns its KeyId and Arn in the response. To specify a CMK, use its key ID,
	// Amazon Resource Name (ARN), alias name, or alias ARN. When using an alias name,
	// prefix it with "alias/". To specify a CMK in a different AWS account, you must
	// use the key ARN or alias ARN. For example:
	//
	// * Key ID:
	// 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// * Key ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// *
	// Alias name: alias/ExampleAlias
	//
	// * Alias ARN:
	// arn:aws:kms:us-east-2:111122223333:alias/ExampleAlias
	//
	// To get the key ID and key
	// ARN for a CMK, use ListKeys or DescribeKey. To get the alias name and alias ARN,
	// use ListAliases.
	//
	// This member is required.
	KeyId *string

	// A list of grant tokens. For more information, see Grant Tokens
	// (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#grant_token)
	// in the AWS Key Management Service Developer Guide.
	GrantTokens []*string
}

type DescribeKeyOutput struct {

	// Metadata associated with the key.
	KeyMetadata *types.KeyMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeKey{}, middleware.After)
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
	if err = addOpDescribeKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "DescribeKey",
	}
}
