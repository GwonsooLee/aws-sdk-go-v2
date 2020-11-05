// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Runs an SQL statement, which can be data manipulation language (DML) or data
// definition language (DDL). This statement must be a single SQL statement.
// Depending on the authorization method, use one of the following combinations of
// request parameters:
//
// * AWS Secrets Manager - specify the Amazon Resource Name
// (ARN) of the secret and the cluster identifier that matches the cluster in the
// secret.
//
// * Temporary credentials - specify the cluster identifier, the database
// name, and the database user name. Permission to call the
// redshift:GetClusterCredentials operation is required to use this method.
func (c *Client) ExecuteStatement(ctx context.Context, params *ExecuteStatementInput, optFns ...func(*Options)) (*ExecuteStatementOutput, error) {
	if params == nil {
		params = &ExecuteStatementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteStatement", params, optFns, addOperationExecuteStatementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExecuteStatementInput struct {

	// The cluster identifier. This parameter is required when authenticating using
	// either AWS Secrets Manager or temporary credentials.
	//
	// This member is required.
	ClusterIdentifier *string

	// The SQL statement text to run.
	//
	// This member is required.
	Sql *string

	// The name of the database. This parameter is required when authenticating using
	// temporary credentials.
	Database *string

	// The database user name. This parameter is required when authenticating using
	// temporary credentials.
	DbUser *string

	// The name or ARN of the secret that enables access to the database. This
	// parameter is required when authenticating using AWS Secrets Manager.
	SecretArn *string

	// The name of the SQL statement. You can name the SQL statement when you create it
	// to identify the query.
	StatementName *string

	// A value that indicates whether to send an event to the Amazon EventBridge event
	// bus after the SQL statement runs.
	WithEvent *bool
}

type ExecuteStatementOutput struct {

	// The cluster identifier.
	ClusterIdentifier *string

	// The date and time (UTC) the statement was created.
	CreatedAt *time.Time

	// The name of the database.
	Database *string

	// The database user name.
	DbUser *string

	// The identifier of the statement to be run. This value is a universally unique
	// identifier (UUID) generated by Amazon Redshift Data API.
	Id *string

	// The name or ARN of the secret that enables access to the database.
	SecretArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExecuteStatementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpExecuteStatement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpExecuteStatement{}, middleware.After)
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
	if err = addOpExecuteStatementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteStatement(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecuteStatement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-data",
		OperationName: "ExecuteStatement",
	}
}
