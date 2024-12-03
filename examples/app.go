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
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/companyinfo/jsonapi"
)

func main() {
	jsonapi.Instrumentation = func(r *jsonapi.Runtime, eventType jsonapi.Event, callGUID string, dur time.Duration) {
		metricPrefix := r.Value("instrument").(string)

		if eventType == jsonapi.UnmarshalStart {
			fmt.Printf("%s: id, %s, started at %v\n", metricPrefix+".jsonapi_unmarshal_time", callGUID, time.Now())
		}

		if eventType == jsonapi.UnmarshalStop {
			fmt.Printf("%s: id, %s, stopped at, %v , and took %v to unmarshal payload\n", metricPrefix+".jsonapi_unmarshal_time", callGUID, time.Now(), dur)
		}

		if eventType == jsonapi.MarshalStart {
			fmt.Printf("%s: id, %s, started at %v\n", metricPrefix+".jsonapi_marshal_time", callGUID, time.Now())
		}

		if eventType == jsonapi.MarshalStop {
			fmt.Printf("%s: id, %s, stopped at, %v , and took %v to marshal payload\n", metricPrefix+".jsonapi_marshal_time", callGUID, time.Now(), dur)
		}
	}

	exampleHandler := &ExampleHandler{}
	http.HandleFunc("/blogs", exampleHandler.ServeHTTP)
	exerciseHandler()
}

func exerciseHandler() {
	// list
	req, _ := http.NewRequest(http.MethodGet, "/blogs", nil)

	req.Header.Set(headerAccept, jsonapi.MediaType)

	w := httptest.NewRecorder()

	fmt.Println("============ start list ===========")
	http.DefaultServeMux.ServeHTTP(w, req)
	fmt.Println("============ stop list ===========")

	jsonReply, _ := ioutil.ReadAll(w.Body)

	fmt.Println("============ jsonapi response from list ===========")
	fmt.Println(string(jsonReply))
	fmt.Println("============== end raw jsonapi from list =============")

	// show
	req, _ = http.NewRequest(http.MethodGet, "/blogs?id=1", nil)

	req.Header.Set(headerAccept, jsonapi.MediaType)

	w = httptest.NewRecorder()

	fmt.Println("============ start show ===========")
	http.DefaultServeMux.ServeHTTP(w, req)
	fmt.Println("============ stop show ===========")

	jsonReply, _ = ioutil.ReadAll(w.Body)

	fmt.Println("============ jsonapi response from show ===========")
	fmt.Println(string(jsonReply))
	fmt.Println("============== end raw jsonapi from show =============")

	// create
	blog := fixtureBlogCreate(1)
	in := bytes.NewBuffer(nil)
	jsonapi.MarshalOnePayloadEmbedded(in, blog)

	req, _ = http.NewRequest(http.MethodPost, "/blogs", in)

	req.Header.Set(headerAccept, jsonapi.MediaType)

	w = httptest.NewRecorder()

	fmt.Println("============ start create ===========")
	http.DefaultServeMux.ServeHTTP(w, req)
	fmt.Println("============ stop create ===========")

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, w.Body)

	fmt.Println("============ jsonapi response from create ===========")
	fmt.Println(buf.String())
	fmt.Println("============== end raw jsonapi response =============")

	// echo
	blogs := []interface{}{
		fixtureBlogCreate(1),
		fixtureBlogCreate(2),
		fixtureBlogCreate(3),
	}
	in = bytes.NewBuffer(nil)
	jsonapi.MarshalPayload(in, blogs)

	req, _ = http.NewRequest(http.MethodPut, "/blogs", in)

	req.Header.Set(headerAccept, jsonapi.MediaType)

	w = httptest.NewRecorder()

	fmt.Println("============ start echo ===========")
	http.DefaultServeMux.ServeHTTP(w, req)
	fmt.Println("============ stop echo ===========")

	buf = bytes.NewBuffer(nil)
	io.Copy(buf, w.Body)

	fmt.Println("============ jsonapi response from create ===========")
	fmt.Println(buf.String())
	fmt.Println("============== end raw jsonapi response =============")

	responseBlog := new(Blog)

	jsonapi.UnmarshalPayload(buf, responseBlog)

	out := bytes.NewBuffer(nil)
	json.NewEncoder(out).Encode(responseBlog)

	fmt.Println("================ Viola! Converted back our Blog struct =================")
	fmt.Println(string(out.Bytes()))
	fmt.Println("================ end marshal materialized Blog struct =================")
}
