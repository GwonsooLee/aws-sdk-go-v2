// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables Security Hub for your account in the current Region or the Region you
// specify in the request. When you enable Security Hub, you grant to Security Hub
// the permissions necessary to gather findings from other services that are
// integrated with Security Hub. When you use the EnableSecurityHub operation to
// enable Security Hub, you also automatically enable the following standards.
//
// *
// CIS AWS Foundations
//
// * AWS Foundational Security Best Practices
//
// You do not
// enable the Payment Card Industry Data Security Standard (PCI DSS) standard. To
// not enable the automatically enabled standards, set EnableDefaultStandards to
// false. After you enable Security Hub, to enable a standard, use the
// BatchEnableStandards operation. To disable a standard, use the
// BatchDisableStandards operation. To learn more, see Setting Up AWS Security Hub
// (https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-settingup.html)
// in the AWS Security Hub User Guide.
func (c *Client) EnableSecurityHub(ctx context.Context, params *EnableSecurityHubInput, optFns ...func(*Options)) (*EnableSecurityHubOutput, error) {
	if params == nil {
		params = &EnableSecurityHubInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableSecurityHub", params, optFns, addOperationEnableSecurityHubMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableSecurityHubOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableSecurityHubInput struct {

	// Whether to enable the security standards that Security Hub has designated as
	// automatically enabled. If you do not provide a value for EnableDefaultStandards,
	// it is set to true. To not enable the automatically enabled standards, set
	// EnableDefaultStandards to false.
	EnableDefaultStandards *bool

	// The tags to add to the hub resource when you enable Security Hub.
	Tags map[string]*string
}

type EnableSecurityHubOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableSecurityHubMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpEnableSecurityHub{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableSecurityHub{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableSecurityHub(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opEnableSecurityHub(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "EnableSecurityHub",
	}
}
