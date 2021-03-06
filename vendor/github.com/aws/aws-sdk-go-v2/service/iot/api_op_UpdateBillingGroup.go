// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateBillingGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the billing group.
	//
	// BillingGroupName is a required field
	BillingGroupName *string `location:"uri" locationName:"billingGroupName" min:"1" type:"string" required:"true"`

	// The properties of the billing group.
	//
	// BillingGroupProperties is a required field
	BillingGroupProperties *BillingGroupProperties `locationName:"billingGroupProperties" type:"structure" required:"true"`

	// The expected version of the billing group. If the version of the billing
	// group does not match the expected version specified in the request, the UpdateBillingGroup
	// request is rejected with a VersionConflictException.
	ExpectedVersion *int64 `locationName:"expectedVersion" type:"long"`
}

// String returns the string representation
func (s UpdateBillingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateBillingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateBillingGroupInput"}

	if s.BillingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("BillingGroupName"))
	}
	if s.BillingGroupName != nil && len(*s.BillingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("BillingGroupName", 1))
	}

	if s.BillingGroupProperties == nil {
		invalidParams.Add(aws.NewErrParamRequired("BillingGroupProperties"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBillingGroupInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.BillingGroupProperties != nil {
		v := s.BillingGroupProperties

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "billingGroupProperties", v, metadata)
	}
	if s.ExpectedVersion != nil {
		v := *s.ExpectedVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "expectedVersion", protocol.Int64Value(v), metadata)
	}
	if s.BillingGroupName != nil {
		v := *s.BillingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "billingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateBillingGroupOutput struct {
	_ struct{} `type:"structure"`

	// The latest version of the billing group.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s UpdateBillingGroupOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateBillingGroupOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Version != nil {
		v := *s.Version

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "version", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opUpdateBillingGroup = "UpdateBillingGroup"

// UpdateBillingGroupRequest returns a request value for making API operation for
// AWS IoT.
//
// Updates information about the billing group.
//
//    // Example sending a request using UpdateBillingGroupRequest.
//    req := client.UpdateBillingGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) UpdateBillingGroupRequest(input *UpdateBillingGroupInput) UpdateBillingGroupRequest {
	op := &aws.Operation{
		Name:       opUpdateBillingGroup,
		HTTPMethod: "PATCH",
		HTTPPath:   "/billing-groups/{billingGroupName}",
	}

	if input == nil {
		input = &UpdateBillingGroupInput{}
	}

	req := c.newRequest(op, input, &UpdateBillingGroupOutput{})
	return UpdateBillingGroupRequest{Request: req, Input: input, Copy: c.UpdateBillingGroupRequest}
}

// UpdateBillingGroupRequest is the request type for the
// UpdateBillingGroup API operation.
type UpdateBillingGroupRequest struct {
	*aws.Request
	Input *UpdateBillingGroupInput
	Copy  func(*UpdateBillingGroupInput) UpdateBillingGroupRequest
}

// Send marshals and sends the UpdateBillingGroup API request.
func (r UpdateBillingGroupRequest) Send(ctx context.Context) (*UpdateBillingGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateBillingGroupResponse{
		UpdateBillingGroupOutput: r.Request.Data.(*UpdateBillingGroupOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateBillingGroupResponse is the response type for the
// UpdateBillingGroup API operation.
type UpdateBillingGroupResponse struct {
	*UpdateBillingGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateBillingGroup request.
func (r *UpdateBillingGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
