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
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newBulkFunc(t Transport) Bulk {
	return func(body io.Reader, o ...func(*BulkRequest)) (*Response, error) {
		var r = BulkRequest{Body: body}
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

// Bulk allows to perform multiple index/update/delete operations in a single request.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-bulk.html.
type Bulk func(body io.Reader, o ...func(*BulkRequest)) (*Response, error)

// BulkRequest configures the Bulk API request.
type BulkRequest struct {
	Index string

	Body io.Reader

	IncludeSourceOnError  *bool
	ListExecutedPipelines *bool
	Pipeline              string
	Refresh               string
	RequireAlias          *bool
	RequireDataStream     *bool
	Routing               string
	Source                []string
	SourceExcludes        []string
	SourceIncludes        []string
	Timeout               time.Duration
	WaitForActiveShards   string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r BulkRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "bulk")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len(r.Index) + 1 + len("_bulk"))
	path.WriteString("http://")
	if r.Index != "" {
		path.WriteString("/")
		path.WriteString(r.Index)
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.Index)
		}
	}
	path.WriteString("/")
	path.WriteString("_bulk")

	params = make(map[string]string)

	if r.IncludeSourceOnError != nil {
		params["include_source_on_error"] = strconv.FormatBool(*r.IncludeSourceOnError)
	}

	if r.ListExecutedPipelines != nil {
		params["list_executed_pipelines"] = strconv.FormatBool(*r.ListExecutedPipelines)
	}

	if r.Pipeline != "" {
		params["pipeline"] = r.Pipeline
	}

	if r.Refresh != "" {
		params["refresh"] = r.Refresh
	}

	if r.RequireAlias != nil {
		params["require_alias"] = strconv.FormatBool(*r.RequireAlias)
	}

	if r.RequireDataStream != nil {
		params["require_data_stream"] = strconv.FormatBool(*r.RequireDataStream)
	}

	if r.Routing != "" {
		params["routing"] = r.Routing
	}

	if len(r.Source) > 0 {
		params["_source"] = strings.Join(r.Source, ",")
	}

	if len(r.SourceExcludes) > 0 {
		params["_source_excludes"] = strings.Join(r.SourceExcludes, ",")
	}

	if len(r.SourceIncludes) > 0 {
		params["_source_includes"] = strings.Join(r.SourceIncludes, ",")
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.WaitForActiveShards != "" {
		params["wait_for_active_shards"] = r.WaitForActiveShards
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

	req, err := newRequest(method, path.String(), r.Body)
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

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "bulk")
		if reader := instrument.RecordRequestBody(ctx, "bulk", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "bulk")
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
func (f Bulk) WithContext(v context.Context) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.ctx = v
	}
}

// WithIndex - default index for items which don't provide one.
func (f Bulk) WithIndex(v string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Index = v
	}
}

// WithIncludeSourceOnError - true or false if to include the document source in the error message in case of parsing errors. defaults to true..
func (f Bulk) WithIncludeSourceOnError(v bool) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.IncludeSourceOnError = &v
	}
}

// WithListExecutedPipelines - sets list_executed_pipelines for all incoming documents. defaults to unset (false).
func (f Bulk) WithListExecutedPipelines(v bool) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.ListExecutedPipelines = &v
	}
}

// WithPipeline - the pipeline ID to preprocess incoming documents with.
func (f Bulk) WithPipeline(v string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Pipeline = v
	}
}

// WithRefresh - if `true` then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` (the default) then do nothing with refreshes..
func (f Bulk) WithRefresh(v string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Refresh = v
	}
}

// WithRequireAlias - if true, the request’s actions must target an index alias. defaults to false..
func (f Bulk) WithRequireAlias(v bool) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.RequireAlias = &v
	}
}

// WithRequireDataStream - if true, the request's actions must target a data stream (existing or to-be-created). default to false.
func (f Bulk) WithRequireDataStream(v bool) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.RequireDataStream = &v
	}
}

// WithRouting - specific routing value.
func (f Bulk) WithRouting(v string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Routing = v
	}
}

// WithSource - true or false to return the _source field or not, or default list of fields to return, can be overridden on each sub-request.
func (f Bulk) WithSource(v ...string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Source = v
	}
}

// WithSourceExcludes - default list of fields to exclude from the returned _source field, can be overridden on each sub-request.
func (f Bulk) WithSourceExcludes(v ...string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.SourceExcludes = v
	}
}

// WithSourceIncludes - default list of fields to extract and return from the _source field, can be overridden on each sub-request.
func (f Bulk) WithSourceIncludes(v ...string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.SourceIncludes = v
	}
}

// WithTimeout - explicit operation timeout.
func (f Bulk) WithTimeout(v time.Duration) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Timeout = v
	}
}

// WithWaitForActiveShards - sets the number of shard copies that must be active before proceeding with the bulk operation. defaults to 1, meaning the primary shard only. set to `all` for all shard copies, otherwise set to any non-negative value less than or equal to the total number of copies for the shard (number of replicas + 1).
func (f Bulk) WithWaitForActiveShards(v string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.WaitForActiveShards = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f Bulk) WithPretty() func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f Bulk) WithHuman() func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f Bulk) WithErrorTrace() func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f Bulk) WithFilterPath(v ...string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f Bulk) WithHeader(h map[string]string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f Bulk) WithOpaqueID(s string) func(*BulkRequest) {
	return func(r *BulkRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
