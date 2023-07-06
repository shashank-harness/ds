// Code generated by goa v3.11.3, DO NOT EDIT.
//
// ds HTTP client encoders and decoders
//
// Command:
// $ goa gen ds/design

package client

import (
	"bytes"
	"context"
	ds "ds/gen/ds"
	dsviews "ds/gen/ds/views"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildListRequest instantiates a HTTP request object with method and path set
// to call the "ds" service "list" endpoint
func (c *Client) BuildListRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ListDsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("ds", "list", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeListResponse returns a decoder for responses returned by the ds list
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeListResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ListResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ds", "list", err)
			}
			p := NewListAccountMgmtCollectionOK(body)
			view := "default"
			vres := dsviews.AccountMgmtCollection{Projected: p, View: view}
			if err = dsviews.ValidateAccountMgmtCollection(vres); err != nil {
				return nil, goahttp.ErrValidationError("ds", "list", err)
			}
			res := ds.NewAccountMgmtCollection(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("ds", "list", resp.StatusCode, string(body))
		}
	}
}

// BuildCompleteRequest instantiates a HTTP request object with method and path
// set to call the "ds" service "complete" endpoint
func (c *Client) BuildCompleteRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		token string
	)
	{
		p, ok := v.(*ds.CompletePayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("ds", "complete", "*ds.CompletePayload", v)
		}
		if p.Token != nil {
			token = *p.Token
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CompleteDsPath(token)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("ds", "complete", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCompleteRequest returns an encoder for requests sent to the ds
// complete server.
func EncodeCompleteRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*ds.CompletePayload)
		if !ok {
			return goahttp.ErrInvalidType("ds", "complete", "*ds.CompletePayload", v)
		}
		body := NewCompleteRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("ds", "complete", err)
		}
		return nil
	}
}

// DecodeCompleteResponse returns a decoder for responses returned by the ds
// complete endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeCompleteResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CompleteResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ds", "complete", err)
			}
			p := NewCompleteUserResourceOK(&body)
			view := "default"
			vres := &dsviews.UserResource{Projected: p, View: view}
			if err = dsviews.ValidateUserResource(vres); err != nil {
				return nil, goahttp.ErrValidationError("ds", "complete", err)
			}
			res := ds.NewUserResource(vres)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("ds", "complete", resp.StatusCode, string(body))
		}
	}
}

// BuildDemoRequest instantiates a HTTP request object with method and path set
// to call the "ds" service "demo" endpoint
func (c *Client) BuildDemoRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*ds.DemoPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("ds", "demo", "*ds.DemoPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DemoDsPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("ds", "demo", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDemoResponse returns a decoder for responses returned by the ds demo
// endpoint. restoreBody controls whether the response body should be restored
// after having been read.
func DecodeDemoResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("ds", "demo", err)
			}
			return body, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("ds", "demo", resp.StatusCode, string(body))
		}
	}
}

// unmarshalAccountMgmtResponseToDsviewsAccountMgmtView builds a value of type
// *dsviews.AccountMgmtView from a value of type *AccountMgmtResponse.
func unmarshalAccountMgmtResponseToDsviewsAccountMgmtView(v *AccountMgmtResponse) *dsviews.AccountMgmtView {
	res := &dsviews.AccountMgmtView{
		ID:          v.ID,
		UUID:        v.UUID,
		Clusterurl:  v.Clusterurl,
		Accountname: v.Accountname,
	}

	return res
}
