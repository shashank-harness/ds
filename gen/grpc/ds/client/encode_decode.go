// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ds gRPC client encoders and decoders
//
// Command:
// $ goa gen ds/design

package client

import (
	"context"
	ds "ds/gen/ds"
	dsviews "ds/gen/ds/views"
	dspb "ds/gen/grpc/ds/pb"

	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildListFunc builds the remote method to invoke for "ds" service "list"
// endpoint.
func BuildListFunc(grpccli dspb.DsClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.List(ctx, reqpb.(*dspb.ListRequest), opts...)
		}
		return grpccli.List(ctx, &dspb.ListRequest{}, opts...)
	}
}

// DecodeListResponse decodes responses from the ds list endpoint.
func DecodeListResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	message, ok := v.(*dspb.AccountMgmtCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ds", "list", "*dspb.AccountMgmtCollection", v)
	}
	res := NewListResult(message)
	vres := dsviews.AccountMgmtCollection{Projected: res, View: view}
	if err := dsviews.ValidateAccountMgmtCollection(vres); err != nil {
		return nil, err
	}
	return ds.NewAccountMgmtCollection(vres), nil
}

// BuildDemoFunc builds the remote method to invoke for "ds" service "demo"
// endpoint.
func BuildDemoFunc(grpccli dspb.DsClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Demo(ctx, reqpb.(*dspb.DemoRequest), opts...)
		}
		return grpccli.Demo(ctx, &dspb.DemoRequest{}, opts...)
	}
}

// EncodeDemoRequest encodes requests sent to ds demo endpoint.
func EncodeDemoRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*ds.DemoPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ds", "demo", "*ds.DemoPayload", v)
	}
	return NewProtoDemoRequest(payload), nil
}

// DecodeDemoResponse decodes responses from the ds demo endpoint.
func DecodeDemoResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*dspb.DemoResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("ds", "demo", "*dspb.DemoResponse", v)
	}
	res := NewDemoResult(message)
	return res, nil
}
