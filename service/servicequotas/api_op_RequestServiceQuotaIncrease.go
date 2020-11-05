// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicequotas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the details of a service quota increase request. The response to this
// command provides the details in the RequestedServiceQuotaChange object.
func (c *Client) RequestServiceQuotaIncrease(ctx context.Context, params *RequestServiceQuotaIncreaseInput, optFns ...func(*Options)) (*RequestServiceQuotaIncreaseOutput, error) {
	if params == nil {
		params = &RequestServiceQuotaIncreaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestServiceQuotaIncrease", params, optFns, addOperationRequestServiceQuotaIncreaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestServiceQuotaIncreaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestServiceQuotaIncreaseInput struct {

	// Specifies the value submitted in the service quota increase request.
	//
	// This member is required.
	DesiredValue *float64

	// Specifies the service quota that you want to use.
	//
	// This member is required.
	QuotaCode *string

	// Specifies the service that you want to use.
	//
	// This member is required.
	ServiceCode *string
}

type RequestServiceQuotaIncreaseOutput struct {

	// Returns a list of service quota requests.
	RequestedQuota *types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRequestServiceQuotaIncreaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRequestServiceQuotaIncrease{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRequestServiceQuotaIncrease{}, middleware.After)
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
	if err = addOpRequestServiceQuotaIncreaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestServiceQuotaIncrease(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRequestServiceQuotaIncrease(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "RequestServiceQuotaIncrease",
	}
}
