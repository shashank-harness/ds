// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ds endpoints
//
// Command:
// $ goa gen ds/design

package ds

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "ds" service endpoints.
type Endpoints struct {
	List     goa.Endpoint
	Complete goa.Endpoint
	Demo     goa.Endpoint
}

// NewEndpoints wraps the methods of the "ds" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		List:     NewListEndpoint(s),
		Complete: NewCompleteEndpoint(s),
		Demo:     NewDemoEndpoint(s),
	}
}

// Use applies the given middleware to all the "ds" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.List = m(e.List)
	e.Complete = m(e.Complete)
	e.Demo = m(e.Demo)
}

// NewListEndpoint returns an endpoint function that calls the method "list" of
// service "ds".
func NewListEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		res, err := s.List(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedAccountMgmtCollection(res, "default")
		return vres, nil
	}
}

// NewCompleteEndpoint returns an endpoint function that calls the method
// "complete" of service "ds".
func NewCompleteEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CompletePayload)
		res, err := s.Complete(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUserResource(res, "default")
		return vres, nil
	}
}

// NewDemoEndpoint returns an endpoint function that calls the method "demo" of
// service "ds".
func NewDemoEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DemoPayload)
		return s.Demo(ctx, p)
	}
}
