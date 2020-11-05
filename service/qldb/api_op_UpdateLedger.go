// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Updates properties on a ledger.
func (c *Client) UpdateLedger(ctx context.Context, params *UpdateLedgerInput, optFns ...func(*Options)) (*UpdateLedgerOutput, error) {
	if params == nil {
		params = &UpdateLedgerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLedger", params, optFns, addOperationUpdateLedgerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLedgerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLedgerInput struct {

	// The name of the ledger.
	//
	// This member is required.
	Name *string

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default. If deletion
	// protection is enabled, you must first disable it before you can delete the
	// ledger using the QLDB API or the AWS Command Line Interface (AWS CLI). You can
	// disable it by calling the UpdateLedger operation to set the flag to false. The
	// QLDB console disables deletion protection for you when you use it to delete a
	// ledger.
	DeletionProtection *bool
}

type UpdateLedgerOutput struct {

	// The Amazon Resource Name (ARN) for the ledger.
	Arn *string

	// The date and time, in epoch time format, when the ledger was created. (Epoch
	// time format is the number of seconds elapsed since 12:00:00 AM January 1, 1970
	// UTC.)
	CreationDateTime *time.Time

	// The flag that prevents a ledger from being deleted by any user. If not provided
	// on ledger creation, this feature is enabled (true) by default. If deletion
	// protection is enabled, you must first disable it before you can delete the
	// ledger using the QLDB API or the AWS Command Line Interface (AWS CLI). You can
	// disable it by calling the UpdateLedger operation to set the flag to false. The
	// QLDB console disables deletion protection for you when you use it to delete a
	// ledger.
	DeletionProtection *bool

	// The name of the ledger.
	Name *string

	// The current status of the ledger.
	State types.LedgerState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateLedgerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLedger{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLedger{}, middleware.After)
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
	if err = addOpUpdateLedgerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLedger(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLedger(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "UpdateLedger",
	}
}
