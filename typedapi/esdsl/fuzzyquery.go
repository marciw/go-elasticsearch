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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _fuzzyQuery struct {
	k string
	v *types.FuzzyQuery
}

// Returns documents that contain terms similar to the search term, as measured
// by a Levenshtein edit distance.
func NewFuzzyQuery(field string, value string) *_fuzzyQuery {
	tmp := &_fuzzyQuery{
		k: field,
		v: types.NewFuzzyQuery(),
	}

	tmp.Value(value)
	return tmp
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_fuzzyQuery) Boost(boost float32) *_fuzzyQuery {

	s.v.Boost = &boost

	return s
}

// Maximum edit distance allowed for matching.
func (s *_fuzzyQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_fuzzyQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

// Maximum number of variations created.
func (s *_fuzzyQuery) MaxExpansions(maxexpansions int) *_fuzzyQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

// Number of beginning characters left unchanged when creating expansions.
func (s *_fuzzyQuery) PrefixLength(prefixlength int) *_fuzzyQuery {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_fuzzyQuery) QueryName_(queryname_ string) *_fuzzyQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Number of beginning characters left unchanged when creating expansions.
func (s *_fuzzyQuery) Rewrite(multitermqueryrewrite string) *_fuzzyQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

// Indicates whether edits include transpositions of two adjacent characters
// (for example `ab` to `ba`).
func (s *_fuzzyQuery) Transpositions(transpositions bool) *_fuzzyQuery {

	s.v.Transpositions = &transpositions

	return s
}

// Term you wish to find in the provided field.
func (s *_fuzzyQuery) Value(value string) *_fuzzyQuery {

	s.v.Value = value

	return s
}

func (s *_fuzzyQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Fuzzy = map[string]types.FuzzyQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleFuzzyQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleFuzzyQuery() *_fuzzyQuery {
	return &_fuzzyQuery{
		k: "",
		v: types.NewFuzzyQuery(),
	}
}

func (s *_fuzzyQuery) FuzzyQueryCaster() *types.FuzzyQuery {
	return s.v.FuzzyQueryCaster()
}
