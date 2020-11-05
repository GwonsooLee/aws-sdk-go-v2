// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a computer account in the specified directory, and joins the computer to
// the directory.
func (c *Client) CreateComputer(ctx context.Context, params *CreateComputerInput, optFns ...func(*Options)) (*CreateComputerOutput, error) {
	if params == nil {
		params = &CreateComputerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComputer", params, optFns, addOperationCreateComputerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComputerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the CreateComputer operation.
type CreateComputerInput struct {

	// The name of the computer account.
	//
	// This member is required.
	ComputerName *string

	// The identifier of the directory in which to create the computer account.
	//
	// This member is required.
	DirectoryId *string

	// A one-time password that is used to join the computer to the directory. You
	// should generate a random, strong password to use for this parameter.
	//
	// This member is required.
	Password *string

	// An array of Attribute objects that contain any LDAP attributes to apply to the
	// computer account.
	ComputerAttributes []*types.Attribute

	// The fully-qualified distinguished name of the organizational unit to place the
	// computer account in.
	OrganizationalUnitDistinguishedName *string
}

// Contains the results for the CreateComputer operation.
type CreateComputerOutput struct {

	// A Computer object that represents the computer account.
	Computer *types.Computer

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateComputerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateComputer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateComputer{}, middleware.After)
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
	if err = addOpCreateComputerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComputer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateComputer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "CreateComputer",
	}
}
