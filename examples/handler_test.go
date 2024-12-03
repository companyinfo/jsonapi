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

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/companyinfo/jsonapi"
)

func TestExampleHandler_post(t *testing.T) {
	blog := fixtureBlogCreate(1)
	requestBody := bytes.NewBuffer(nil)
	jsonapi.MarshalOnePayloadEmbedded(requestBody, blog)

	r, err := http.NewRequest(http.MethodPost, "/blogs?id=1", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set(headerAccept, jsonapi.MediaType)

	rr := httptest.NewRecorder()
	handler := &ExampleHandler{}
	handler.ServeHTTP(rr, r)

	if e, a := http.StatusCreated, rr.Code; e != a {
		t.Fatalf("Expected a status of %d, got %d", e, a)
	}
}

func TestExampleHandler_put(t *testing.T) {
	blogs := []interface{}{
		fixtureBlogCreate(1),
		fixtureBlogCreate(2),
		fixtureBlogCreate(3),
	}
	requestBody := bytes.NewBuffer(nil)
	jsonapi.MarshalPayload(requestBody, blogs)

	r, err := http.NewRequest(http.MethodPut, "/blogs", requestBody)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set(headerAccept, jsonapi.MediaType)

	rr := httptest.NewRecorder()
	handler := &ExampleHandler{}
	handler.ServeHTTP(rr, r)

	if e, a := http.StatusOK, rr.Code; e != a {
		t.Fatalf("Expected a status of %d, got %d", e, a)
	}
}

func TestExampleHandler_get_show(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/blogs?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set(headerAccept, jsonapi.MediaType)

	rr := httptest.NewRecorder()
	handler := &ExampleHandler{}
	handler.ServeHTTP(rr, r)

	if e, a := http.StatusOK, rr.Code; e != a {
		t.Fatalf("Expected a status of %d, got %d", e, a)
	}
}

func TestExampleHandler_get_list(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set(headerAccept, jsonapi.MediaType)

	rr := httptest.NewRecorder()
	handler := &ExampleHandler{}
	handler.ServeHTTP(rr, r)

	if e, a := http.StatusOK, rr.Code; e != a {
		t.Fatalf("Expected a status of %d, got %d", e, a)
	}
}

func TestHttpErrorWhenHeaderDoesNotMatch(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set(headerAccept, "application/xml")

	rr := httptest.NewRecorder()
	handler := &ExampleHandler{}
	handler.ServeHTTP(rr, r)

	if rr.Code != http.StatusUnsupportedMediaType {
		t.Fatal("expected Unsupported Media Type staus error")
	}
}

func TestHttpErrorWhenMethodDoesNotMatch(t *testing.T) {
	r, err := http.NewRequest(http.MethodPatch, "/blogs", nil)
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set(headerAccept, jsonapi.MediaType)

	rr := httptest.NewRecorder()
	handler := &ExampleHandler{}
	handler.ServeHTTP(rr, r)

	if rr.Code != http.StatusNotFound {
		t.Fatal("expected HTTP Status Not Found status error")
	}
}
