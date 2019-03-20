// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr     = flag.String("listen-address", ":80", "The address to listen on for HTTP requests.")
	redirect = flag.String("redirect-address", "https://example.com", "The URL to redirect to.")
)

func main() {
	flag.Parse()

	log.Fatal(http.ListenAndServe(*addr, http.RedirectHandler(*redirect, http.StatusMovedPermanently)))
}
