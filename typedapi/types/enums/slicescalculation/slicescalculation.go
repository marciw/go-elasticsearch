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

// Package slicescalculation
package slicescalculation

import "strings"

// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/_types/common.ts#L368-L376
type SlicesCalculation struct {
	Name string
}

var (
	Auto = SlicesCalculation{"auto"}
)

func (s SlicesCalculation) MarshalText() (text []byte, err error) {
	return []byte(s.String()), nil
}

func (s *SlicesCalculation) UnmarshalText(text []byte) error {
	switch strings.ReplaceAll(strings.ToLower(string(text)), "\"", "") {

	case "auto":
		*s = Auto
	default:
		*s = SlicesCalculation{string(text)}
	}

	return nil
}

func (s SlicesCalculation) String() string {
	return s.Name
}
