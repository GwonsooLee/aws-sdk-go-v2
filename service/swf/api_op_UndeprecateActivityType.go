// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Undeprecates a previously deprecated activity type. After an activity type has
// been undeprecated, you can create new tasks of that activity type. This
// operation is eventually consistent. The results are best effort and may not
// exactly reflect recent updates and changes. Access Control You can use IAM
// policies to control this action's access to Amazon SWF resources as follows:
//
// *
// Use a Resource element with the domain name to limit the action to only
// specified domains.
//
// * Use an Action element to allow or deny permission to call
// this action.
//
// * Constrain the following parameters by using a Condition element
// with the appropriate keys.
//
// * activityType.name: String constraint. The key is
// swf:activityType.name.
//
// * activityType.version: String constraint. The key is
// swf:activityType.version.
//
// If the caller doesn't have sufficient permissions to
// invoke the action, or the parameter values fall outside the specified
// constraints, the action fails. The associated event attribute's cause parameter
// is set to OPERATION_NOT_PERMITTED. For details and example IAM policies, see
// Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) UndeprecateActivityType(ctx context.Context, params *UndeprecateActivityTypeInput, optFns ...func(*Options)) (*UndeprecateActivityTypeOutput, error) {
	if params == nil {
		params = &UndeprecateActivityTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UndeprecateActivityType", params, optFns, addOperationUndeprecateActivityTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UndeprecateActivityTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UndeprecateActivityTypeInput struct {

	// The activity type to undeprecate.
	//
	// This member is required.
	ActivityType *types.ActivityType

	// The name of the domain of the deprecated activity type.
	//
	// This member is required.
	Domain *string
}

type UndeprecateActivityTypeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUndeprecateActivityTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpUndeprecateActivityType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpUndeprecateActivityType{}, middleware.After)
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
	if err = addOpUndeprecateActivityTypeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUndeprecateActivityType(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUndeprecateActivityType(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "UndeprecateActivityType",
	}
}
