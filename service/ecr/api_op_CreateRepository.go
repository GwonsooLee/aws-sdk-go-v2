// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a repository. For more information, see Amazon ECR Repositories
// (https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html) in
// the Amazon Elastic Container Registry User Guide.
func (c *Client) CreateRepository(ctx context.Context, params *CreateRepositoryInput, optFns ...func(*Options)) (*CreateRepositoryOutput, error) {
	if params == nil {
		params = &CreateRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRepository", params, optFns, addOperationCreateRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRepositoryInput struct {

	// The name to use for the repository. The repository name may be specified on its
	// own (such as nginx-web-app) or it can be prepended with a namespace to group the
	// repository into a category (such as project-a/nginx-web-app).
	//
	// This member is required.
	RepositoryName *string

	// The encryption configuration for the repository. This determines how the
	// contents of your repository are encrypted at rest.
	EncryptionConfiguration *types.EncryptionConfiguration

	// The image scanning configuration for the repository. This determines whether
	// images are scanned for known vulnerabilities after being pushed to the
	// repository.
	ImageScanningConfiguration *types.ImageScanningConfiguration

	// The tag mutability setting for the repository. If this parameter is omitted, the
	// default setting of MUTABLE will be used which will allow image tags to be
	// overwritten. If IMMUTABLE is specified, all image tags within the repository
	// will be immutable which will prevent them from being overwritten.
	ImageTagMutability types.ImageTagMutability

	// The metadata that you apply to the repository to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. Tag keys can have a maximum character length of 128 characters, and
	// tag values can have a maximum length of 256 characters.
	Tags []*types.Tag
}

type CreateRepositoryOutput struct {

	// The repository that was created.
	Repository *types.Repository

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRepository{}, middleware.After)
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
	if err = addOpCreateRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRepository(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "CreateRepository",
	}
}
