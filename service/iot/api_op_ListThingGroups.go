// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the thing groups in your account.
func (c *Client) ListThingGroups(ctx context.Context, params *ListThingGroupsInput, optFns ...func(*Options)) (*ListThingGroupsOutput, error) {
	if params == nil {
		params = &ListThingGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingGroups", params, optFns, addOperationListThingGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingGroupsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// A filter that limits the results to those with the specified name prefix.
	NamePrefixFilter *string

	// The token to retrieve the next set of results.
	NextToken *string

	// A filter that limits the results to those with the specified parent group.
	ParentGroup *string

	// If true, return child groups as well.
	Recursive *bool
}

type ListThingGroupsOutput struct {

	// The token used to get the next set of results. Will not be returned if operation
	// has returned all results.
	NextToken *string

	// The thing groups.
	ThingGroups []*types.GroupNameAndArn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListThingGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListThingGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingGroups",
	}
}
