// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.0
// source: api/bloodstatus/v1/blood_status.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationBloodStatusCreateBloodStatus = "/api.bloodstatus.v1.BloodStatus/CreateBloodStatus"
const OperationBloodStatusDeleteBloodStatus = "/api.bloodstatus.v1.BloodStatus/DeleteBloodStatus"
const OperationBloodStatusGetBloodStatus = "/api.bloodstatus.v1.BloodStatus/GetBloodStatus"
const OperationBloodStatusListBloodStatus = "/api.bloodstatus.v1.BloodStatus/ListBloodStatus"
const OperationBloodStatusUpdateBloodStatus = "/api.bloodstatus.v1.BloodStatus/UpdateBloodStatus"

type BloodStatusHTTPServer interface {
	CreateBloodStatus(context.Context, *CreateBloodStatusRequest) (*CreateBloodStatusReply, error)
	DeleteBloodStatus(context.Context, *DeleteBloodStatusRequest) (*DeleteBloodStatusReply, error)
	GetBloodStatus(context.Context, *GetBloodStatusRequest) (*GetBloodStatusReply, error)
	ListBloodStatus(context.Context, *ListBloodStatusRequest) (*ListBloodStatusReply, error)
	UpdateBloodStatus(context.Context, *UpdateBloodStatusRequest) (*UpdateBloodStatusReply, error)
}

func RegisterBloodStatusHTTPServer(s *http.Server, srv BloodStatusHTTPServer) {
	r := s.Route("/")
	r.POST("api/v1/blood_status", _BloodStatus_CreateBloodStatus0_HTTP_Handler(srv))
	r.PUT("api/v1/blood_status", _BloodStatus_UpdateBloodStatus0_HTTP_Handler(srv))
	r.PATCH("api/v1/blood_status/{id}", _BloodStatus_DeleteBloodStatus0_HTTP_Handler(srv))
	r.GET("api/v1/blood_status/{id}", _BloodStatus_GetBloodStatus0_HTTP_Handler(srv))
	r.GET("api/v1/blood_status", _BloodStatus_ListBloodStatus0_HTTP_Handler(srv))
}

func _BloodStatus_CreateBloodStatus0_HTTP_Handler(srv BloodStatusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateBloodStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBloodStatusCreateBloodStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateBloodStatus(ctx, req.(*CreateBloodStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateBloodStatusReply)
		return ctx.Result(200, reply)
	}
}

func _BloodStatus_UpdateBloodStatus0_HTTP_Handler(srv BloodStatusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateBloodStatusRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBloodStatusUpdateBloodStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateBloodStatus(ctx, req.(*UpdateBloodStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateBloodStatusReply)
		return ctx.Result(200, reply)
	}
}

func _BloodStatus_DeleteBloodStatus0_HTTP_Handler(srv BloodStatusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteBloodStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBloodStatusDeleteBloodStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteBloodStatus(ctx, req.(*DeleteBloodStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteBloodStatusReply)
		return ctx.Result(200, reply)
	}
}

func _BloodStatus_GetBloodStatus0_HTTP_Handler(srv BloodStatusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetBloodStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBloodStatusGetBloodStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetBloodStatus(ctx, req.(*GetBloodStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetBloodStatusReply)
		return ctx.Result(200, reply)
	}
}

func _BloodStatus_ListBloodStatus0_HTTP_Handler(srv BloodStatusHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListBloodStatusRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationBloodStatusListBloodStatus)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListBloodStatus(ctx, req.(*ListBloodStatusRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListBloodStatusReply)
		return ctx.Result(200, reply)
	}
}

type BloodStatusHTTPClient interface {
	CreateBloodStatus(ctx context.Context, req *CreateBloodStatusRequest, opts ...http.CallOption) (rsp *CreateBloodStatusReply, err error)
	DeleteBloodStatus(ctx context.Context, req *DeleteBloodStatusRequest, opts ...http.CallOption) (rsp *DeleteBloodStatusReply, err error)
	GetBloodStatus(ctx context.Context, req *GetBloodStatusRequest, opts ...http.CallOption) (rsp *GetBloodStatusReply, err error)
	ListBloodStatus(ctx context.Context, req *ListBloodStatusRequest, opts ...http.CallOption) (rsp *ListBloodStatusReply, err error)
	UpdateBloodStatus(ctx context.Context, req *UpdateBloodStatusRequest, opts ...http.CallOption) (rsp *UpdateBloodStatusReply, err error)
}

type BloodStatusHTTPClientImpl struct {
	cc *http.Client
}

func NewBloodStatusHTTPClient(client *http.Client) BloodStatusHTTPClient {
	return &BloodStatusHTTPClientImpl{client}
}

func (c *BloodStatusHTTPClientImpl) CreateBloodStatus(ctx context.Context, in *CreateBloodStatusRequest, opts ...http.CallOption) (*CreateBloodStatusReply, error) {
	var out CreateBloodStatusReply
	pattern := "api/v1/blood_status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBloodStatusCreateBloodStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BloodStatusHTTPClientImpl) DeleteBloodStatus(ctx context.Context, in *DeleteBloodStatusRequest, opts ...http.CallOption) (*DeleteBloodStatusReply, error) {
	var out DeleteBloodStatusReply
	pattern := "api/v1/blood_status/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBloodStatusDeleteBloodStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PATCH", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BloodStatusHTTPClientImpl) GetBloodStatus(ctx context.Context, in *GetBloodStatusRequest, opts ...http.CallOption) (*GetBloodStatusReply, error) {
	var out GetBloodStatusReply
	pattern := "api/v1/blood_status/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBloodStatusGetBloodStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BloodStatusHTTPClientImpl) ListBloodStatus(ctx context.Context, in *ListBloodStatusRequest, opts ...http.CallOption) (*ListBloodStatusReply, error) {
	var out ListBloodStatusReply
	pattern := "api/v1/blood_status"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationBloodStatusListBloodStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *BloodStatusHTTPClientImpl) UpdateBloodStatus(ctx context.Context, in *UpdateBloodStatusRequest, opts ...http.CallOption) (*UpdateBloodStatusReply, error) {
	var out UpdateBloodStatusReply
	pattern := "api/v1/blood_status"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationBloodStatusUpdateBloodStatus))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
