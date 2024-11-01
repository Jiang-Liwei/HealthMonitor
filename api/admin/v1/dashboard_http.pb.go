// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.0
// source: api/admin/v1/dashboard.proto

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

const OperationDashboardDashboard = "/api.admin.v1.Dashboard/Dashboard"

type DashboardHTTPServer interface {
	Dashboard(context.Context, *DashboardRequest) (*DashboardReply, error)
}

func RegisterDashboardHTTPServer(s *http.Server, srv DashboardHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/dashboard", _Dashboard_Dashboard0_HTTP_Handler(srv))
}

func _Dashboard_Dashboard0_HTTP_Handler(srv DashboardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DashboardRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDashboardDashboard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Dashboard(ctx, req.(*DashboardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DashboardReply)
		return ctx.Result(200, reply)
	}
}

type DashboardHTTPClient interface {
	Dashboard(ctx context.Context, req *DashboardRequest, opts ...http.CallOption) (rsp *DashboardReply, err error)
}

type DashboardHTTPClientImpl struct {
	cc *http.Client
}

func NewDashboardHTTPClient(client *http.Client) DashboardHTTPClient {
	return &DashboardHTTPClientImpl{client}
}

func (c *DashboardHTTPClientImpl) Dashboard(ctx context.Context, in *DashboardRequest, opts ...http.CallOption) (*DashboardReply, error) {
	var out DashboardReply
	pattern := "/admin/v1/dashboard"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDashboardDashboard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}