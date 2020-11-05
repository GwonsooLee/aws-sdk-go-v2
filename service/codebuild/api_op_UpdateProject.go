// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Changes the settings of a build project.
func (c *Client) UpdateProject(ctx context.Context, params *UpdateProjectInput, optFns ...func(*Options)) (*UpdateProjectOutput, error) {
	if params == nil {
		params = &UpdateProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProject", params, optFns, addOperationUpdateProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProjectInput struct {

	// The name of the build project. You cannot change a build project's name.
	//
	// This member is required.
	Name *string

	// Information to be changed about the build output artifacts for the build
	// project.
	Artifacts *types.ProjectArtifacts

	// Set this to true to generate a publicly accessible URL for your project's build
	// badge.
	BadgeEnabled *bool

	// Contains configuration information about a batch build project.
	BuildBatchConfig *types.ProjectBuildBatchConfig

	// Stores recently used information so that it can be quickly accessed at a later
	// time.
	Cache *types.ProjectCache

	// A new or replacement description of the build project.
	Description *string

	// The AWS Key Management Service (AWS KMS) customer master key (CMK) to be used
	// for encrypting the build output artifacts. You can use a cross-account KMS key
	// to encrypt the build output artifacts if your service role has permission to
	// that key. You can specify either the Amazon Resource Name (ARN) of the CMK or,
	// if available, the CMK's alias (using the format alias/).
	EncryptionKey *string

	// Information to be changed about the build environment for the build project.
	Environment *types.ProjectEnvironment

	// An array of ProjectFileSystemLocation objects for a CodeBuild build project. A
	// ProjectFileSystemLocation object specifies the identifier, location,
	// mountOptions, mountPoint, and type of a file system created using Amazon Elastic
	// File System.
	FileSystemLocations []*types.ProjectFileSystemLocation

	// Information about logs for the build project. A project can create logs in
	// Amazon CloudWatch Logs, logs in an S3 bucket, or both.
	LogsConfig *types.LogsConfig

	// The number of minutes a build is allowed to be queued before it times out.
	QueuedTimeoutInMinutes *int32

	// An array of ProjectSource objects.
	SecondaryArtifacts []*types.ProjectArtifacts

	// An array of ProjectSourceVersion objects. If secondarySourceVersions is
	// specified at the build level, then they take over these secondarySourceVersions
	// (at the project level).
	SecondarySourceVersions []*types.ProjectSourceVersion

	// An array of ProjectSource objects.
	SecondarySources []*types.ProjectSource

	// The replacement ARN of the AWS Identity and Access Management (IAM) role that
	// enables AWS CodeBuild to interact with dependent AWS services on behalf of the
	// AWS account.
	ServiceRole *string

	// Information to be changed about the build input source code for the build
	// project.
	Source *types.ProjectSource

	// A version of the build input to be built for this project. If not specified, the
	// latest version is used. If specified, it must be one of:
	//
	// * For AWS CodeCommit:
	// the commit ID, branch, or Git tag to use.
	//
	// * For GitHub: the commit ID, pull
	// request ID, branch name, or tag name that corresponds to the version of the
	// source code you want to build. If a pull request ID is specified, it must use
	// the format pr/pull-request-ID (for example pr/25). If a branch name is
	// specified, the branch's HEAD commit ID is used. If not specified, the default
	// branch's HEAD commit ID is used.
	//
	// * For Bitbucket: the commit ID, branch name,
	// or tag name that corresponds to the version of the source code you want to
	// build. If a branch name is specified, the branch's HEAD commit ID is used. If
	// not specified, the default branch's HEAD commit ID is used.
	//
	// * For Amazon Simple
	// Storage Service (Amazon S3): the version ID of the object that represents the
	// build input ZIP file to use.
	//
	// If sourceVersion is specified at the build level,
	// then that version takes precedence over this sourceVersion (at the project
	// level). For more information, see Source Version Sample with CodeBuild
	// (https://docs.aws.amazon.com/codebuild/latest/userguide/sample-source-version.html)
	// in the AWS CodeBuild User Guide.
	SourceVersion *string

	// An updated list of tag key and value pairs associated with this build project.
	// These tags are available for use by AWS services that support AWS CodeBuild
	// build project tags.
	Tags []*types.Tag

	// The replacement value in minutes, from 5 to 480 (8 hours), for AWS CodeBuild to
	// wait before timing out any related build that did not get marked as completed.
	TimeoutInMinutes *int32

	// VpcConfig enables AWS CodeBuild to access resources in an Amazon VPC.
	VpcConfig *types.VpcConfig
}

type UpdateProjectOutput struct {

	// Information about the build project that was changed.
	Project *types.Project

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProject{}, middleware.After)
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
	if err = addOpUpdateProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProject(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "UpdateProject",
	}
}
