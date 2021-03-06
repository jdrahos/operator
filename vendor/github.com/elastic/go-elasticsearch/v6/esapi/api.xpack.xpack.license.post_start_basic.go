// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated from specification version 6.8.8: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newXPackLicensePostStartBasicFunc(t Transport) XPackLicensePostStartBasic {
	return func(o ...func(*XPackLicensePostStartBasicRequest)) (*Response, error) {
		var r = XPackLicensePostStartBasicRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackLicensePostStartBasic - https://www.elastic.co/guide/en/elasticsearch/reference/6.7/start-basic.html
//
type XPackLicensePostStartBasic func(o ...func(*XPackLicensePostStartBasicRequest)) (*Response, error)

// XPackLicensePostStartBasicRequest configures the X Pack License Post Start Basic API request.
//
type XPackLicensePostStartBasicRequest struct {
	Acknowledge *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r XPackLicensePostStartBasicRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(len("/_xpack/license/start_basic"))
	path.WriteString("/_xpack/license/start_basic")

	params = make(map[string]string)

	if r.Acknowledge != nil {
		params["acknowledge"] = strconv.FormatBool(*r.Acknowledge)
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
//
func (f XPackLicensePostStartBasic) WithContext(v context.Context) func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		r.ctx = v
	}
}

// WithAcknowledge - whether the user has acknowledged acknowledge messages (default: false).
//
func (f XPackLicensePostStartBasic) WithAcknowledge(v bool) func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		r.Acknowledge = &v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f XPackLicensePostStartBasic) WithPretty() func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f XPackLicensePostStartBasic) WithHuman() func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f XPackLicensePostStartBasic) WithErrorTrace() func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f XPackLicensePostStartBasic) WithFilterPath(v ...string) func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f XPackLicensePostStartBasic) WithHeader(h map[string]string) func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f XPackLicensePostStartBasic) WithOpaqueID(s string) func(*XPackLicensePostStartBasicRequest) {
	return func(r *XPackLicensePostStartBasicRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
