// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the configuration information about an origin access identity.
func (c *Client) GetCloudFrontOriginAccessIdentityConfig(ctx context.Context, params *GetCloudFrontOriginAccessIdentityConfigInput, optFns ...func(*Options)) (*GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	if params == nil {
		params = &GetCloudFrontOriginAccessIdentityConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCloudFrontOriginAccessIdentityConfig", params, optFns, addOperationGetCloudFrontOriginAccessIdentityConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCloudFrontOriginAccessIdentityConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The origin access identity's configuration information. For more information,
// see CloudFrontOriginAccessIdentityConfig
// (https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CloudFrontOriginAccessIdentityConfig.html).
type GetCloudFrontOriginAccessIdentityConfigInput struct {

	// The identity's ID.
	//
	// This member is required.
	Id *string
}

// The returned result of the corresponding request.
type GetCloudFrontOriginAccessIdentityConfigOutput struct {

	// The origin access identity's configuration information.
	CloudFrontOriginAccessIdentityConfig *types.CloudFrontOriginAccessIdentityConfig

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCloudFrontOriginAccessIdentityConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetCloudFrontOriginAccessIdentityConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetCloudFrontOriginAccessIdentityConfig{}, middleware.After)
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
	if err = addOpGetCloudFrontOriginAccessIdentityConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentityConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCloudFrontOriginAccessIdentityConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "GetCloudFrontOriginAccessIdentityConfig",
	}
}
