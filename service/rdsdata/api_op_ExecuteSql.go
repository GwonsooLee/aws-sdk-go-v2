// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs one or more SQL statements. This operation is deprecated. Use the
// BatchExecuteStatement or ExecuteStatement operation.
func (c *Client) ExecuteSql(ctx context.Context, params *ExecuteSqlInput, optFns ...func(*Options)) (*ExecuteSqlOutput, error) {
	if params == nil {
		params = &ExecuteSqlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExecuteSql", params, optFns, addOperationExecuteSqlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExecuteSqlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a request to run one or more SQL
// statements.
type ExecuteSqlInput struct {

	// The Amazon Resource Name (ARN) of the secret that enables access to the DB
	// cluster.
	//
	// This member is required.
	AwsSecretStoreArn *string

	// The ARN of the Aurora Serverless DB cluster.
	//
	// This member is required.
	DbClusterOrInstanceArn *string

	// One or more SQL statements to run on the DB cluster. You can separate SQL
	// statements from each other with a semicolon (;). Any valid SQL statement is
	// permitted, including data definition, data manipulation, and commit statements.
	//
	// This member is required.
	SqlStatements *string

	// The name of the database.
	Database *string

	// The name of the database schema.
	Schema *string
}

// The response elements represent the output of a request to run one or more SQL
// statements.
type ExecuteSqlOutput struct {

	// The results of the SQL statement or statements.
	SqlStatementResults []*types.SqlStatementResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationExecuteSqlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExecuteSql{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExecuteSql{}, middleware.After)
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
	if err = addOpExecuteSqlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExecuteSql(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExecuteSql(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "ExecuteSql",
	}
}
