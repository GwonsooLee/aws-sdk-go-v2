// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of archive rules created for the specified analyzer.
func (c *Client) ListArchiveRules(ctx context.Context, params *ListArchiveRulesInput, optFns ...func(*Options)) (*ListArchiveRulesOutput, error) {
	if params == nil {
		params = &ListArchiveRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListArchiveRules", params, optFns, addOperationListArchiveRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListArchiveRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieves a list of archive rules created for the specified analyzer.
type ListArchiveRulesInput struct {

	// The name of the analyzer to retrieve rules from.
	//
	// This member is required.
	AnalyzerName *string

	// The maximum number of results to return in the request.
	MaxResults *int32

	// A token used for pagination of results returned.
	NextToken *string
}

// The response to the request.
type ListArchiveRulesOutput struct {

	// A list of archive rules created for the specified analyzer.
	//
	// This member is required.
	ArchiveRules []*types.ArchiveRuleSummary

	// A token used for pagination of results returned.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListArchiveRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListArchiveRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListArchiveRules{}, middleware.After)
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
	if err = addOpListArchiveRulesValidationMiddleware(stack); err != nil {
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
