// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a version of a robot application.
func (c *Client) CreateRobotApplicationVersion(ctx context.Context, params *CreateRobotApplicationVersionInput, optFns ...func(*Options)) (*CreateRobotApplicationVersionOutput, error) {
	if params == nil {
		params = &CreateRobotApplicationVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRobotApplicationVersion", params, optFns, addOperationCreateRobotApplicationVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRobotApplicationVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRobotApplicationVersionInput struct {

	// The application information for the robot application.
	//
	// This member is required.
	Application *string

	// The current revision id for the robot application. If you provide a value and it
	// matches the latest revision ID, a new version will be created.
	CurrentRevisionId *string
}

type CreateRobotApplicationVersionOutput struct {

	// The Amazon Resource Name (ARN) of the robot application.
	Arn *string

	// The time, in milliseconds since the epoch, when the robot application was last
	// updated.
	LastUpdatedAt *time.Time

	// The name of the robot application.
	Name *string

	// The revision id of the robot application.
	RevisionId *string

	// The robot software suite (ROS distribution) used by the robot application.
	RobotSoftwareSuite *types.RobotSoftwareSuite

	// The sources of the robot application.
	Sources []*types.Source

	// The version of the robot application.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRobotApplicationVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRobotApplicationVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRobotApplicationVersion{}, middleware.After)
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
	if err = addOpCreateRobotApplicationVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRobotApplicationVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRobotApplicationVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateRobotApplicationVersion",
	}
}
