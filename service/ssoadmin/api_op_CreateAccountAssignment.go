// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns access to a principal for a specified AWS account using a specified
// permission set. The term principal here refers to a user or group that is
// defined in AWS SSO. As part of a successful CreateAccountAssignment call, the
// specified permission set will automatically be provisioned to the account in the
// form of an IAM policy attached to the SSO-created IAM role. If the permission
// set is subsequently updated, the corresponding IAM policies attached to roles in
// your accounts will not be updated automatically. In this case, you will need to
// call ProvisionPermissionSet to make these updates.
func (c *Client) CreateAccountAssignment(ctx context.Context, params *CreateAccountAssignmentInput, optFns ...func(*Options)) (*CreateAccountAssignmentOutput, error) {
	if params == nil {
		params = &CreateAccountAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccountAssignment", params, optFns, addOperationCreateAccountAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccountAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccountAssignmentInput struct {

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the permission set that the admin wants to grant the principal access
	// to.
	//
	// This member is required.
	PermissionSetArn *string

	// An identifier for an object in AWS SSO, such as a user or group. PrincipalIds
	// are GUIDs (For example, f81d4fae-7dec-11d0-a765-00a0c91e6bf6). For more
	// information about PrincipalIds in AWS SSO, see the AWS SSO Identity Store API
	// Reference.
	//
	// This member is required.
	PrincipalId *string

	// The entity type for which the assignment will be created.
	//
	// This member is required.
	PrincipalType types.PrincipalType

	// TargetID is an AWS account identifier, typically a 10-12 digit string (For
	// example, 123456789012).
	//
	// This member is required.
	TargetId *string

	// The entity type for which the assignment will be created.
	//
	// This member is required.
	TargetType types.TargetType
}

type CreateAccountAssignmentOutput struct {

	// The status object for the account assignment creation operation.
	AccountAssignmentCreationStatus *types.AccountAssignmentOperationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateAccountAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAccountAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAccountAssignment{}, middleware.After)
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
	if err = addOpCreateAccountAssignmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccountAssignment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAccountAssignment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "CreateAccountAssignment",
	}
}
