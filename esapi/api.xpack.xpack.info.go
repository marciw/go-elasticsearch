// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated from specification version 9.1.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newXPackInfoFunc(t Transport) XPackInfo {
	return func(o ...func(*XPackInfoRequest)) (*Response, error) {
		var r = XPackInfoRequest{}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// XPackInfo - Retrieves information about the installed X-Pack features.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/info-api.html.
type XPackInfo func(o ...func(*XPackInfoRequest)) (*Response, error)

// XPackInfoRequest configures the X Pack Info API request.
type XPackInfoRequest struct {
	AcceptEnterprise *bool
	Categories       []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r XPackInfoRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "xpack.info")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_xpack"))
	path.WriteString("http://")
	path.WriteString("/_xpack")

	params = make(map[string]string)

	if r.AcceptEnterprise != nil {
		params["accept_enterprise"] = strconv.FormatBool(*r.AcceptEnterprise)
	}

	if len(r.Categories) > 0 {
		params["categories"] = strings.Join(r.Categories, ",")
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
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "xpack.info")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "xpack.info")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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
func (f XPackInfo) WithContext(v context.Context) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.ctx = v
	}
}

// WithAcceptEnterprise - if this param is used it must be set to true.
func (f XPackInfo) WithAcceptEnterprise(v bool) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.AcceptEnterprise = &v
	}
}

// WithCategories - comma-separated list of info categories. can be any of: build, license, features.
func (f XPackInfo) WithCategories(v ...string) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.Categories = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f XPackInfo) WithPretty() func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f XPackInfo) WithHuman() func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f XPackInfo) WithErrorTrace() func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f XPackInfo) WithFilterPath(v ...string) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f XPackInfo) WithHeader(h map[string]string) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f XPackInfo) WithOpaqueID(s string) func(*XPackInfoRequest) {
	return func(r *XPackInfoRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
