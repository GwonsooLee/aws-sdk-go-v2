// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the databases in a cluster. A token is returned to page through the
// database list. Depending on the authorization method, use one of the following
// combinations of request parameters:
//
// * AWS Secrets Manager - specify the Amazon
// Resource Name (ARN) of the secret and the cluster identifier that matches the
// cluster in the secret.
//
// * Temporary credentials - specify the cluster
// identifier, the database name, and the database user name. Permission to call
// the redshift:GetClusterCredentials operation is required to use this method.
func (c *Client) ListDatabases(ctx context.Context, params *ListDatabasesInput, optFns ...func(*Options)) (*ListDatabasesOutput, error) {
	if params == nil {
		params = &ListDatabasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatabases", params, optFns, addOperationListDatabasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatabasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatabasesInput struct {

	// The cluster identifier. This parameter is required when authenticating using
	// either AWS Secrets Manager or temporary credentials.
	//
	// This member is required.
	ClusterIdentifier *string

	// The name of the database. This parameter is required when authenticating using
	// temporary credentials.
	Database *string

	// The database user name. This parameter is required when authenticating using
	// temporary credentials.
	DbUser *string

	// The maximum number of databases to return in the response. If more databases
	// exist than fit in one response, then NextToken is returned to page through the
	// results.
	MaxResults *int32

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using AWS Secrets Manager.
	SecretArn *string
}

type ListDatabasesOutput struct {

	// The names of databases.
	Databases []*string

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned NextToken value in the next
	// NextToken parameter and retrying the command. If the NextToken field is empty,
	// all response records have been retrieved for the request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDatabasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatabases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatabases{}, middleware.After)
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
	if err = addOpListDatabasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDatabases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDatabases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "ListDatabases",
	}
}
