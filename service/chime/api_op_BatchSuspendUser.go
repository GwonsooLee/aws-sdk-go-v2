// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Suspends up to 50 users from a Team or EnterpriseLWA Amazon Chime account. For
// more information about different account types, see Managing Your Amazon Chime
// Accounts (https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html)
// in the Amazon Chime Administration Guide. Users suspended from a Team account
// are disassociated from the account, but they can continue to use Amazon Chime as
// free users. To remove the suspension from suspended Team account users, invite
// them to the Team account again. You can use the InviteUsers action to do so.
// Users suspended from an EnterpriseLWA account are immediately signed out of
// Amazon Chime and can no longer sign in. To remove the suspension from suspended
// EnterpriseLWA account users, use the BatchUnsuspendUser action. To sign out
// users without suspending them, use the LogoutUser action.
func (c *Client) BatchSuspendUser(ctx context.Context, params *BatchSuspendUserInput, optFns ...func(*Options)) (*BatchSuspendUserOutput, error) {
	if params == nil {
		params = &BatchSuspendUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchSuspendUser", params, optFns, addOperationBatchSuspendUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchSuspendUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchSuspendUserInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The request containing the user IDs to suspend.
	//
	// This member is required.
	UserIdList []*string
}

type BatchSuspendUserOutput struct {

	// If the BatchSuspendUser action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []*types.UserError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchSuspendUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchSuspendUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchSuspendUser{}, middleware.After)
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
	if err = addOpBatchSuspendUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchSuspendUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchSuspendUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchSuspendUser",
	}
}
