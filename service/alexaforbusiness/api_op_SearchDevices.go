// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches devices and lists the ones that meet a set of filter criteria.
func (c *Client) SearchDevices(ctx context.Context, params *SearchDevicesInput, optFns ...func(*Options)) (*SearchDevicesOutput, error) {
	if params == nil {
		params = &SearchDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchDevices", params, optFns, addOperationSearchDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDevicesInput struct {

	// The filters to use to list a specified set of devices. Supported filter keys are
	// DeviceName, DeviceStatus, DeviceStatusDetailCode, RoomName, DeviceType,
	// DeviceSerialNumber, UnassociatedOnly, ConnectionStatus (ONLINE and OFFLINE),
	// NetworkProfileName, NetworkProfileArn, Feature, and FailureCode.
	Filters []*types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use in listing the specified set of devices. Supported sort
	// keys are DeviceName, DeviceStatus, RoomName, DeviceType, DeviceSerialNumber,
	// ConnectionStatus, NetworkProfileName, NetworkProfileArn, Feature, and
	// FailureCode.
	SortCriteria []*types.Sort
}

type SearchDevicesOutput struct {

	// The devices that meet the specified set of filter criteria, in sort order.
	Devices []*types.DeviceData

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The total number of devices returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchDevices{}, middleware.After)
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
	if err = addOpSearchDevicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDevices(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchDevices",
	}
}
