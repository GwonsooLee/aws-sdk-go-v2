// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeNodeConfigurationOptionsInput struct {
	_ struct{} `type:"structure"`

	// The action type to evaluate for possible node configurations. Specify "restore-cluster"
	// to get configuration combinations based on an existing snapshot. Specify
	// "recommend-node-config" to get configuration recommendations based on an
	// existing cluster or snapshot.
	//
	// ActionType is a required field
	ActionType ActionType `type:"string" required:"true" enum:"true"`

	// The identifier of the cluster to evaluate for possible node configurations.
	ClusterIdentifier *string `type:"string"`

	// A set of name, operator, and value items to filter the results.
	Filters []NodeConfigurationOptionsFilter `locationName:"Filter" locationNameList:"NodeConfigurationOptionsFilter" type:"list"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeNodeConfigurationOptions
	// request exceed the value specified in MaxRecords, AWS returns a value in
	// the Marker field of the response. You can retrieve the next set of response
	// records by providing the returned marker value in the Marker parameter and
	// retrying the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 500
	//
	// Constraints: minimum 100, maximum 500.
	MaxRecords *int64 `type:"integer"`

	// The AWS customer account used to create or copy the snapshot. Required if
	// you are restoring a snapshot you do not own, optional if you own the snapshot.
	OwnerAccount *string `type:"string"`

	// The identifier of the snapshot to evaluate for possible node configurations.
	SnapshotIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DescribeNodeConfigurationOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNodeConfigurationOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeNodeConfigurationOptionsInput"}
	if len(s.ActionType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ActionType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeNodeConfigurationOptionsOutput struct {
	_ struct{} `type:"structure"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`

	// A list of valid node configurations.
	NodeConfigurationOptionList []NodeConfigurationOption `locationNameList:"NodeConfigurationOption" type:"list"`
}

// String returns the string representation
func (s DescribeNodeConfigurationOptionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeNodeConfigurationOptions = "DescribeNodeConfigurationOptions"

// DescribeNodeConfigurationOptionsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns properties of possible node configurations such as node type, number
// of nodes, and disk usage for the specified action type.
//
//    // Example sending a request using DescribeNodeConfigurationOptionsRequest.
//    req := client.DescribeNodeConfigurationOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeNodeConfigurationOptions
func (c *Client) DescribeNodeConfigurationOptionsRequest(input *DescribeNodeConfigurationOptionsInput) DescribeNodeConfigurationOptionsRequest {
	op := &aws.Operation{
		Name:       opDescribeNodeConfigurationOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeNodeConfigurationOptionsInput{}
	}

	req := c.newRequest(op, input, &DescribeNodeConfigurationOptionsOutput{})
	return DescribeNodeConfigurationOptionsRequest{Request: req, Input: input, Copy: c.DescribeNodeConfigurationOptionsRequest}
}

// DescribeNodeConfigurationOptionsRequest is the request type for the
// DescribeNodeConfigurationOptions API operation.
type DescribeNodeConfigurationOptionsRequest struct {
	*aws.Request
	Input *DescribeNodeConfigurationOptionsInput
	Copy  func(*DescribeNodeConfigurationOptionsInput) DescribeNodeConfigurationOptionsRequest
}

// Send marshals and sends the DescribeNodeConfigurationOptions API request.
func (r DescribeNodeConfigurationOptionsRequest) Send(ctx context.Context) (*DescribeNodeConfigurationOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeNodeConfigurationOptionsResponse{
		DescribeNodeConfigurationOptionsOutput: r.Request.Data.(*DescribeNodeConfigurationOptionsOutput),
		response:                               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeNodeConfigurationOptionsRequestPaginator returns a paginator for DescribeNodeConfigurationOptions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeNodeConfigurationOptionsRequest(input)
//   p := redshift.NewDescribeNodeConfigurationOptionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeNodeConfigurationOptionsPaginator(req DescribeNodeConfigurationOptionsRequest) DescribeNodeConfigurationOptionsPaginator {
	return DescribeNodeConfigurationOptionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeNodeConfigurationOptionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeNodeConfigurationOptionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeNodeConfigurationOptionsPaginator struct {
	aws.Pager
}

func (p *DescribeNodeConfigurationOptionsPaginator) CurrentPage() *DescribeNodeConfigurationOptionsOutput {
	return p.Pager.CurrentPage().(*DescribeNodeConfigurationOptionsOutput)
}

// DescribeNodeConfigurationOptionsResponse is the response type for the
// DescribeNodeConfigurationOptions API operation.
type DescribeNodeConfigurationOptionsResponse struct {
	*DescribeNodeConfigurationOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeNodeConfigurationOptions request.
func (r *DescribeNodeConfigurationOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}