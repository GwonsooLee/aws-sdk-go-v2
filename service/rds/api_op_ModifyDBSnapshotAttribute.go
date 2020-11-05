// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds an attribute and values to, or removes an attribute and values from, a
// manual DB snapshot. To share a manual DB snapshot with other AWS accounts,
// specify restore as the AttributeName and use the ValuesToAdd parameter to add a
// list of IDs of the AWS accounts that are authorized to restore the manual DB
// snapshot. Uses the value all to make the manual DB snapshot public, which means
// it can be copied or restored by all AWS accounts. Don't add the all value for
// any manual DB snapshots that contain private information that you don't want
// available to all AWS accounts. If the manual DB snapshot is encrypted, it can be
// shared, but only by specifying a list of authorized AWS account IDs for the
// ValuesToAdd parameter. You can't use all as a value for that parameter in this
// case. To view which AWS accounts have access to copy or restore a manual DB
// snapshot, or whether a manual DB snapshot public or private, use the
// DescribeDBSnapshotAttributes API action. The accounts are returned as values for
// the restore attribute.
func (c *Client) ModifyDBSnapshotAttribute(ctx context.Context, params *ModifyDBSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBSnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifyDBSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBSnapshotAttribute", params, optFns, addOperationModifyDBSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ModifyDBSnapshotAttributeInput struct {

	// The name of the DB snapshot attribute to modify. To manage authorization for
	// other AWS accounts to copy or restore a manual DB snapshot, set this value to
	// restore. To view the list of attributes available to modify, use the
	// DescribeDBSnapshotAttributes API action.
	//
	// This member is required.
	AttributeName *string

	// The identifier for the DB snapshot to modify the attributes for.
	//
	// This member is required.
	DBSnapshotIdentifier *string

	// A list of DB snapshot attributes to add to the attribute specified by
	// AttributeName. To authorize other AWS accounts to copy or restore a manual
	// snapshot, set this list to include one or more AWS account IDs, or all to make
	// the manual DB snapshot restorable by any AWS account. Do not add the all value
	// for any manual DB snapshots that contain private information that you don't want
	// available to all AWS accounts.
	ValuesToAdd []*string

	// A list of DB snapshot attributes to remove from the attribute specified by
	// AttributeName. To remove authorization for other AWS accounts to copy or restore
	// a manual snapshot, set this list to include one or more AWS account identifiers,
	// or all to remove authorization for any AWS account to copy or restore the DB
	// snapshot. If you specify all, an AWS account whose account ID is explicitly
	// added to the restore attribute can still copy or restore the manual DB snapshot.
	ValuesToRemove []*string
}

type ModifyDBSnapshotAttributeOutput struct {

	// Contains the results of a successful call to the DescribeDBSnapshotAttributes
	// API action. Manual DB snapshot attributes are used to authorize other AWS
	// accounts to copy or restore a manual DB snapshot. For more information, see the
	// ModifyDBSnapshotAttribute API action.
	DBSnapshotAttributesResult *types.DBSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyDBSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBSnapshotAttribute{}, middleware.After)
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
	if err = addOpModifyDBSnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBSnapshotAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBSnapshotAttribute",
	}
}
