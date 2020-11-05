// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// If you created a Shield Advanced policy, returns policy-level attack summary
// information in the event of a potential DDoS attack. Other policy types are
// currently unsupported.
func (c *Client) GetProtectionStatus(ctx context.Context, params *GetProtectionStatusInput, optFns ...func(*Options)) (*GetProtectionStatusOutput, error) {
	if params == nil {
		params = &GetProtectionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetProtectionStatus", params, optFns, addOperationGetProtectionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetProtectionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetProtectionStatusInput struct {

	// The ID of the policy for which you want to get the attack information.
	//
	// This member is required.
	PolicyId *string

	// The end of the time period to query for the attacks. This is a timestamp type.
	// The request syntax listing indicates a number type because the default used by
	// AWS Firewall Manager is Unix time in seconds. However, any valid timestamp
	// format is allowed.
	EndTime *time.Time

	// Specifies the number of objects that you want AWS Firewall Manager to return for
	// this request. If you have more objects than the number that you specify for
	// MaxResults, the response includes a NextToken value that you can use to get
	// another batch of objects.
	MaxResults *int32

	// The AWS account that is in scope of the policy that you want to get the details
	// for.
	MemberAccountId *string

	// If you specify a value for MaxResults and you have more objects than the number
	// that you specify for MaxResults, AWS Firewall Manager returns a NextToken value
	// in the response, which you can use to retrieve another group of objects. For the
	// second and subsequent GetProtectionStatus requests, specify the value of
	// NextToken from the previous response to get information about another batch of
	// objects.
	NextToken *string

	// The start of the time period to query for the attacks. This is a timestamp type.
	// The request syntax listing indicates a number type because the default used by
	// AWS Firewall Manager is Unix time in seconds. However, any valid timestamp
	// format is allowed.
	StartTime *time.Time
}

type GetProtectionStatusOutput struct {

	// The ID of the AWS Firewall administrator account for this policy.
	AdminAccountId *string

	// Details about the attack, including the following:
	//
	// * Attack type
	//
	// * Account
	// ID
	//
	// * ARN of the resource attacked
	//
	// * Start time of the attack
	//
	// * End time of
	// the attack (ongoing attacks will not have an end time)
	//
	// The details are in JSON
	// format.
	Data *string

	// If you have more objects than the number that you specified for MaxResults in
	// the request, the response includes a NextToken value. To list more objects,
	// submit another GetProtectionStatus request, and specify the NextToken value from
	// the response in the NextToken value in the next request. AWS SDKs provide
	// auto-pagination that identify NextToken in a response and make subsequent
	// request calls automatically on your behalf. However, this feature is not
	// supported by GetProtectionStatus. You must submit subsequent requests with
	// NextToken using your own processes.
	NextToken *string

	// The service type that is protected by the policy. Currently, this is always
	// SHIELD_ADVANCED.
	ServiceType types.SecurityServiceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetProtectionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetProtectionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetProtectionStatus{}, middleware.After)
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
	if err = addOpGetProtectionStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetProtectionStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetProtectionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "GetProtectionStatus",
	}
}
