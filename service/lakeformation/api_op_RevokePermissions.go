// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Revokes permissions to the principal to access metadata in the Data Catalog and
// data organized in underlying data storage such as Amazon S3.
func (c *Client) RevokePermissions(ctx context.Context, params *RevokePermissionsInput, optFns ...func(*Options)) (*RevokePermissionsOutput, error) {
	if params == nil {
		params = &RevokePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokePermissions", params, optFns, addOperationRevokePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokePermissionsInput struct {

	// The permissions revoked to the principal on the resource. For information about
	// permissions, see Security and Access Control to Metadata and Data
	// (https://docs-aws.amazon.com/lake-formation/latest/dg/security-data-access.html).
	//
	// This member is required.
	Permissions []types.Permission

	// The principal to be revoked permissions on the resource.
	//
	// This member is required.
	Principal *types.DataLakePrincipal

	// The resource to which permissions are to be revoked.
	//
	// This member is required.
	Resource *types.Resource

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// Indicates a list of permissions for which to revoke the grant option allowing
	// the principal to pass permissions to other principals.
	PermissionsWithGrantOption []types.Permission
}

type RevokePermissionsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRevokePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRevokePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRevokePermissions{}, middleware.After)
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
	if err = addOpRevokePermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokePermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "RevokePermissions",
	}
}
