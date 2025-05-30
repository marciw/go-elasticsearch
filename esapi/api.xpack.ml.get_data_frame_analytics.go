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

func newMLGetDataFrameAnalyticsFunc(t Transport) MLGetDataFrameAnalytics {
	return func(o ...func(*MLGetDataFrameAnalyticsRequest)) (*Response, error) {
		var r = MLGetDataFrameAnalyticsRequest{}
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

// MLGetDataFrameAnalytics - Retrieves configuration information for data frame analytics jobs.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/get-dfanalytics.html.
type MLGetDataFrameAnalytics func(o ...func(*MLGetDataFrameAnalyticsRequest)) (*Response, error)

// MLGetDataFrameAnalyticsRequest configures the ML Get Data Frame Analytics API request.
type MLGetDataFrameAnalyticsRequest struct {
	ID string

	AllowNoMatch     *bool
	ExcludeGenerated *bool
	From             *int
	Size             *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MLGetDataFrameAnalyticsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ml.get_data_frame_analytics")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_ml") + 1 + len("data_frame") + 1 + len("analytics") + 1 + len(r.ID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("data_frame")
	path.WriteString("/")
	path.WriteString("analytics")
	if r.ID != "" {
		path.WriteString("/")
		path.WriteString(r.ID)
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", r.ID)
		}
	}

	params = make(map[string]string)

	if r.AllowNoMatch != nil {
		params["allow_no_match"] = strconv.FormatBool(*r.AllowNoMatch)
	}

	if r.ExcludeGenerated != nil {
		params["exclude_generated"] = strconv.FormatBool(*r.ExcludeGenerated)
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
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
		instrument.BeforeRequest(req, "ml.get_data_frame_analytics")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ml.get_data_frame_analytics")
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
func (f MLGetDataFrameAnalytics) WithContext(v context.Context) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.ctx = v
	}
}

// WithID - the ID of the data frame analytics to fetch.
func (f MLGetDataFrameAnalytics) WithID(v string) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.ID = v
	}
}

// WithAllowNoMatch - whether to ignore if a wildcard expression matches no data frame analytics. (this includes `_all` string or when no data frame analytics have been specified).
func (f MLGetDataFrameAnalytics) WithAllowNoMatch(v bool) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.AllowNoMatch = &v
	}
}

// WithExcludeGenerated - omits fields that are illegal to set on data frame analytics put.
func (f MLGetDataFrameAnalytics) WithExcludeGenerated(v bool) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.ExcludeGenerated = &v
	}
}

// WithFrom - skips a number of analytics.
func (f MLGetDataFrameAnalytics) WithFrom(v int) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.From = &v
	}
}

// WithSize - specifies a max number of analytics to get.
func (f MLGetDataFrameAnalytics) WithSize(v int) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.Size = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLGetDataFrameAnalytics) WithPretty() func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLGetDataFrameAnalytics) WithHuman() func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLGetDataFrameAnalytics) WithErrorTrace() func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLGetDataFrameAnalytics) WithFilterPath(v ...string) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLGetDataFrameAnalytics) WithHeader(h map[string]string) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLGetDataFrameAnalytics) WithOpaqueID(s string) func(*MLGetDataFrameAnalyticsRequest) {
	return func(r *MLGetDataFrameAnalyticsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
