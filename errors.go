// Copyright 2024 Company.info B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package jsonapi

import (
	"encoding/json"
	"fmt"
	"io"
)

// MarshalErrors writes a JSON API response using the given `[]error`.
//
// For more information on JSON API error payloads, see the spec here:
// http://jsonapi.org/format/#document-top-level
// and here: http://jsonapi.org/format/#error-objects.
func MarshalErrors(w io.Writer, errorObjects []*ErrorObject) error {
	return json.NewEncoder(w).Encode(&ErrorsPayload{Errors: errorObjects})
}

// ErrorsPayload is a serializer struct for representing a valid JSON API errors payload.
type ErrorsPayload struct {
	Errors []*ErrorObject `json:"errors"`
}

// ErrorObject is an `Error` implementation as well as an implementation of the JSON API error object.
//
// The main idea behind this struct is that you can use it directly in your code as an error type
// and pass it directly to `MarshalErrors` to get a valid JSON API errors payload.
// For more information on Golang errors, see: https://golang.org/pkg/errors/
// For more information on the JSON API spec's error objects, see: http://jsonapi.org/format/#error-objects
type ErrorObject struct {
	// ID is a unique identifier for this particular occurrence of a problem.
	ID string `json:"id,omitempty"`

	// Title is a short, human-readable summary of the problem that SHOULD NOT change from occurrence to occurrence of the problem, except for purposes of localization.
	Title string `json:"title,omitempty"`

	// Detail is a human-readable explanation specific to this occurrence of the problem. Like title, this field’s value can be localized.
	Detail string `json:"detail,omitempty"`

	// Status is the HTTP status code applicable to this problem, expressed as a string value.
	Status string `json:"status,omitempty"`

	// Code is an application-specific error code, expressed as a string value.
	Code string `json:"code,omitempty"`

	//Source is an object containing references to the primary source of the error.
	Source *Source `json:"source,omitempty"`

	// Meta is an object containing non-standard meta-information about the error.
	Meta *map[string]interface{} `json:"meta,omitempty"`
}

// Error implements the `Error` interface.
func (e *ErrorObject) Error() string {
	return fmt.Sprintf("Error: %s %s\n", e.Title, e.Detail)
}

// Source is an object containing references to the primary source of the error.
type Source struct {
	// Pointer is a string indicating the value in the request document that caused the error.
	Pointer string `json:"pointer,omitempty"`

	// Parameter is a string indicating which query or path parameter caused the error.
	Parameter string `json:"parameter,omitempty"`

	// Header is a string indicating the name of a single request header which caused the error.
	Header string `json:"header,omitempty"`
}
