// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the header information for the .csv file to be used as input for the user
// import job.
func (c *Client) GetCSVHeader(ctx context.Context, params *GetCSVHeaderInput, optFns ...func(*Options)) (*GetCSVHeaderOutput, error) {
	if params == nil {
		params = &GetCSVHeaderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCSVHeader", params, optFns, addOperationGetCSVHeaderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCSVHeaderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get the header information for the .csv file for the
// user import job.
type GetCSVHeaderInput struct {

	// The user pool ID for the user pool that the users are to be imported into.
	//
	// This member is required.
	UserPoolId *string
}

// Represents the response from the server to the request to get the header
// information for the .csv file for the user import job.
type GetCSVHeaderOutput struct {

	// The header information for the .csv file for the user import job.
	CSVHeader []*string

	// The user pool ID for the user pool that the users are to be imported into.
	UserPoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCSVHeaderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCSVHeader{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCSVHeader{}, middleware.After)
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
	if err = addOpGetCSVHeaderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCSVHeader(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetCSVHeader(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetCSVHeader",
	}
}
