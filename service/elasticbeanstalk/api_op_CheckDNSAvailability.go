// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Checks if the specified CNAME is available.
func (c *Client) CheckDNSAvailability(ctx context.Context, params *CheckDNSAvailabilityInput, optFns ...func(*Options)) (*CheckDNSAvailabilityOutput, error) {
	if params == nil {
		params = &CheckDNSAvailabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CheckDNSAvailability", params, optFns, addOperationCheckDNSAvailabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CheckDNSAvailabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Results message indicating whether a CNAME is available.
type CheckDNSAvailabilityInput struct {

	// The prefix used when this CNAME is reserved.
	//
	// This member is required.
	CNAMEPrefix *string
}

// Indicates if the specified CNAME is available.
type CheckDNSAvailabilityOutput struct {

	// Indicates if the specified CNAME is available:
	//
	// * true : The CNAME is
	// available.
	//
	// * false : The CNAME is not available.
	Available *bool

	// The fully qualified CNAME to reserve when CreateEnvironment is called with the
	// provided prefix.
	FullyQualifiedCNAME *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCheckDNSAvailabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCheckDNSAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCheckDNSAvailability{}, middleware.After)
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
	if err = addOpCheckDNSAvailabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCheckDNSAvailability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCheckDNSAvailability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "CheckDNSAvailability",
	}
}
