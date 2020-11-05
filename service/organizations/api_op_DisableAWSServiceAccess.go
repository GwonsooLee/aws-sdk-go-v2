// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the integration of an AWS service (the service that is specified by
// ServicePrincipal) with AWS Organizations. When you disable integration, the
// specified service no longer can create a service-linked role
// (http://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html)
// in new accounts in your organization. This means the service can't perform
// operations on your behalf on any new accounts in your organization. The service
// can still perform operations in older accounts until the service completes its
// clean-up from AWS Organizations. We recommend that you disable integration
// between AWS Organizations and the specified AWS service by using the console or
// commands that are provided by the specified service. Doing so ensures that the
// other service is aware that it can clean up any resources that are required only
// for the integration. How the service cleans up its resources in the
// organization's accounts depends on that service. For more information, see the
// documentation for the other AWS service. After you perform the
// DisableAWSServiceAccess operation, the specified service can no longer perform
// operations in your organization's accounts unless the operations are explicitly
// permitted by the IAM policies that are attached to your roles. For more
// information about integrating other services with AWS Organizations, including
// the list of services that work with Organizations, see Integrating AWS
// Organizations with Other AWS Services
// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's management account.
func (c *Client) DisableAWSServiceAccess(ctx context.Context, params *DisableAWSServiceAccessInput, optFns ...func(*Options)) (*DisableAWSServiceAccessOutput, error) {
	if params == nil {
		params = &DisableAWSServiceAccessInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisableAWSServiceAccess", params, optFns, addOperationDisableAWSServiceAccessMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisableAWSServiceAccessOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisableAWSServiceAccessInput struct {

	// The service principal name of the AWS service for which you want to disable
	// integration with your organization. This is typically in the form of a URL, such
	// as  service-abbreviation.amazonaws.com.
	//
	// This member is required.
	ServicePrincipal *string
}

type DisableAWSServiceAccessOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisableAWSServiceAccessMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisableAWSServiceAccess{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisableAWSServiceAccess{}, middleware.After)
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
	if err = addOpDisableAWSServiceAccessValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisableAWSServiceAccess(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisableAWSServiceAccess(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DisableAWSServiceAccess",
	}
}
