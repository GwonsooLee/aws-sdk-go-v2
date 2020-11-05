// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns compliance details of a conformance pack for all AWS resources that are
// monitered by conformance pack.
func (c *Client) GetConformancePackComplianceDetails(ctx context.Context, params *GetConformancePackComplianceDetailsInput, optFns ...func(*Options)) (*GetConformancePackComplianceDetailsOutput, error) {
	if params == nil {
		params = &GetConformancePackComplianceDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetConformancePackComplianceDetails", params, optFns, addOperationGetConformancePackComplianceDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetConformancePackComplianceDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetConformancePackComplianceDetailsInput struct {

	// Name of the conformance pack.
	//
	// This member is required.
	ConformancePackName *string

	// A ConformancePackEvaluationFilters object.
	Filters *types.ConformancePackEvaluationFilters

	// The maximum number of evaluation results returned on each page. If you do no
	// specify a number, AWS Config uses the default. The default is 100.
	Limit *int32

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string
}

type GetConformancePackComplianceDetailsOutput struct {

	// Name of the conformance pack.
	//
	// This member is required.
	ConformancePackName *string

	// Returns a list of ConformancePackEvaluationResult objects.
	ConformancePackRuleEvaluationResults []*types.ConformancePackEvaluationResult

	// The nextToken string returned in a previous request that you use to request the
	// next page of results in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetConformancePackComplianceDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetConformancePackComplianceDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetConformancePackComplianceDetails{}, middleware.After)
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
	if err = addOpGetConformancePackComplianceDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetConformancePackComplianceDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetConformancePackComplianceDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetConformancePackComplianceDetails",
	}
}
