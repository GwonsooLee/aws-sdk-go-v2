// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates one or more Amazon Lightsail instances. The create instances operation
// supports tag-based access control via request tags. For more information, see
// the Lightsail Dev Guide
// (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
func (c *Client) CreateInstances(ctx context.Context, params *CreateInstancesInput, optFns ...func(*Options)) (*CreateInstancesOutput, error) {
	if params == nil {
		params = &CreateInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInstances", params, optFns, addOperationCreateInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInstancesInput struct {

	// The Availability Zone in which to create your instance. Use the following
	// format: us-east-2a (case sensitive). You can get a list of Availability Zones by
	// using the get regions
	// (http://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetRegions.html)
	// operation. Be sure to add the include Availability Zones parameter to your
	// request.
	//
	// This member is required.
	AvailabilityZone *string

	// The ID for a virtual private server image (e.g., app_wordpress_4_4 or
	// app_lamp_7_0). Use the get blueprints operation to return a list of available
	// images (or blueprints). Use active blueprints when creating new instances.
	// Inactive blueprints are listed to support customers with existing instances and
	// are not necessarily available to create new instances. Blueprints are marked
	// inactive when they become outdated due to operating system updates or new
	// application releases.
	//
	// This member is required.
	BlueprintId *string

	// The bundle of specification information for your virtual private server (or
	// instance), including the pricing plan (e.g., micro_1_0).
	//
	// This member is required.
	BundleId *string

	// The names to use for your new Lightsail instances. Separate multiple values
	// using quotation marks and commas, for example:
	// ["MyFirstInstance","MySecondInstance"]
	//
	// This member is required.
	InstanceNames []*string

	// An array of objects representing the add-ons to enable for the new instance.
	AddOns []*types.AddOnRequest

	// (Deprecated) The name for your custom image. In releases prior to June 12, 2017,
	// this parameter was ignored by the API. It is now deprecated.
	CustomImageName *string

	// The name of your key pair.
	KeyPairName *string

	// The tag keys and optional values to add to the resource during create. Use the
	// TagResource action to tag a resource after it's created.
	Tags []*types.Tag

	// A launch script you can create that configures a server with additional user
	// data. For example, you might want to run apt-get -y update. Depending on the
	// machine image you choose, the command to get software on your instance varies.
	// Amazon Linux and CentOS use yum, Debian and Ubuntu use apt-get, and FreeBSD uses
	// pkg. For a complete list, see the Dev Guide
	// (https://lightsail.aws.amazon.com/ls/docs/getting-started/article/compare-options-choose-lightsail-instance-image).
	UserData *string
}

type CreateInstancesOutput struct {

	// An array of objects that describe the result of the action, such as the status
	// of the request, the timestamp of the request, and the resources affected by the
	// request.
	Operations []*types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateInstances{}, middleware.After)
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
	if err = addOpCreateInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "CreateInstances",
	}
}
