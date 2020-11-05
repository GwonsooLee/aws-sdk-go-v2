// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the specified user's password in a user pool as an administrator. Works on
// any user. The password can be temporary or permanent. If it is temporary, the
// user status will be placed into the FORCE_CHANGE_PASSWORD state. When the user
// next tries to sign in, the InitiateAuth/AdminInitiateAuth response will contain
// the NEW_PASSWORD_REQUIRED challenge. If the user does not sign in before it
// expires, the user will not be able to sign in and their password will need to be
// reset by an administrator. Once the user has set a new password, or the password
// is permanent, the user status will be set to Confirmed.
func (c *Client) AdminSetUserPassword(ctx context.Context, params *AdminSetUserPasswordInput, optFns ...func(*Options)) (*AdminSetUserPasswordOutput, error) {
	if params == nil {
		params = &AdminSetUserPasswordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminSetUserPassword", params, optFns, addOperationAdminSetUserPasswordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminSetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminSetUserPasswordInput struct {

	// The password for the user.
	//
	// This member is required.
	Password *string

	// The user pool ID for the user pool where you want to set the user's password.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user whose password you wish to set.
	//
	// This member is required.
	Username *string

	// True if the password is permanent, False if it is temporary.
	Permanent *bool
}

type AdminSetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminSetUserPasswordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserPassword{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserPassword{}, middleware.After)
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
	if err = addOpAdminSetUserPasswordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserPassword(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminSetUserPassword(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminSetUserPassword",
	}
}
