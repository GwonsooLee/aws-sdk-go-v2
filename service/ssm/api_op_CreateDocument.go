// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Systems Manager (SSM) document. An SSM document defines the actions
// that Systems Manager performs on your managed instances. For more information
// about SSM documents, including information about supported schemas, features,
// and syntax, see AWS Systems Manager Documents
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-ssm-docs.html)
// in the AWS Systems Manager User Guide.
func (c *Client) CreateDocument(ctx context.Context, params *CreateDocumentInput, optFns ...func(*Options)) (*CreateDocumentOutput, error) {
	if params == nil {
		params = &CreateDocumentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDocument", params, optFns, addOperationCreateDocumentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDocumentInput struct {

	// The content for the new SSM document in JSON or YAML format. We recommend
	// storing the contents for your new document in an external JSON or YAML file and
	// referencing the file in a command. For examples, see the following topics in the
	// AWS Systems Manager User Guide.
	//
	// * Create an SSM document (AWS API)
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/create-ssm-document-api.html)
	//
	// *
	// Create an SSM document (AWS CLI)
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/create-ssm-document-cli.html)
	//
	// *
	// Create an SSM document (API)
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/create-ssm-document-api.html)
	//
	// This member is required.
	Content *string

	// A name for the Systems Manager document. You can't use the following strings as
	// document name prefixes. These are reserved by AWS for use as document name
	// prefixes:
	//
	// * aws-
	//
	// * amazon
	//
	// * amzn
	//
	// This member is required.
	Name *string

	// A list of key and value pairs that describe attachments to a version of a
	// document.
	Attachments []*types.AttachmentsSource

	// Specify the document format for the request. The document format can be JSON,
	// YAML, or TEXT. JSON is the default format.
	DocumentFormat types.DocumentFormat

	// The type of document to create.
	DocumentType types.DocumentType

	// A list of SSM documents required by a document. This parameter is used
	// exclusively by AWS AppConfig. When a user creates an AppConfig configuration in
	// an SSM document, the user must also specify a required document for validation
	// purposes. In this case, an ApplicationConfiguration document requires an
	// ApplicationConfigurationSchema document for validation purposes. For more
	// information, see AWS AppConfig
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/appconfig.html) in
	// the AWS Systems Manager User Guide.
	Requires []*types.DocumentRequires

	// Optional metadata that you assign to a resource. Tags enable you to categorize a
	// resource in different ways, such as by purpose, owner, or environment. For
	// example, you might want to tag an SSM document to identify the types of targets
	// or the environment where it will run. In this case, you could specify the
	// following key name/value pairs:
	//
	// * Key=OS,Value=Windows
	//
	// *
	// Key=Environment,Value=Production
	//
	// To add tags to an existing SSM document, use
	// the AddTagsToResource action.
	Tags []*types.Tag

	// Specify a target type to define the kinds of resources the document can run on.
	// For example, to run a document on EC2 instances, specify the following value:
	// /AWS::EC2::Instance. If you specify a value of '/' the document can run on all
	// types of resources. If you don't specify a value, the document can't run on any
	// resources. For a list of valid resource types, see AWS resource and property
	// types reference
	// (http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// in the AWS CloudFormation User Guide.
	TargetType *string

	// An optional field specifying the version of the artifact you are creating with
	// the document. For example, "Release 12, Update 6". This value is unique across
	// all versions of a document, and cannot be changed.
	VersionName *string
}

type CreateDocumentOutput struct {

	// Information about the Systems Manager document.
	DocumentDescription *types.DocumentDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateDocumentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateDocument{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateDocument{}, middleware.After)
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
	if err = addOpCreateDocumentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDocument(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDocument(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateDocument",
	}
}
