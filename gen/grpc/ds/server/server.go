// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ds gRPC server
//
// Command:
// $ goa gen ds/design

package server

import (
	"context"
	ds "ds/gen/ds"
	dspb "ds/gen/grpc/ds/pb"

	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the dspb.DsServer interface.
type Server struct {
	ListH goagrpc.UnaryHandler
	DemoH goagrpc.UnaryHandler
	dspb.UnimplementedDsServer
}

// New instantiates the server struct with the ds service endpoints.
func New(e *ds.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		ListH: NewListHandler(e.List, uh),
		DemoH: NewDemoHandler(e.Demo, uh),
	}
}

// NewListHandler creates a gRPC handler which serves the "ds" service "list"
// endpoint.
func NewListHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, nil, EncodeListResponse)
	}
	return h
}

// List implements the "List" method in dspb.DsServer interface.
func (s *Server) List(ctx context.Context, message *dspb.ListRequest) (*dspb.AccountMgmtCollection, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "list")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ds")
	resp, err := s.ListH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*dspb.AccountMgmtCollection), nil
}

// NewDemoHandler creates a gRPC handler which serves the "ds" service "demo"
// endpoint.
func NewDemoHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDemoRequest, EncodeDemoResponse)
	}
	return h
}

// Demo implements the "Demo" method in dspb.DsServer interface.
func (s *Server) Demo(ctx context.Context, message *dspb.DemoRequest) (*dspb.DemoResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "demo")
	ctx = context.WithValue(ctx, goa.ServiceKey, "ds")
	resp, err := s.DemoH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*dspb.DemoResponse), nil
}
