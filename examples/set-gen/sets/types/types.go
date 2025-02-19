/*
Copyright 2015 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package types just provides input types to the set generator. It also
// contains a "go generate" block.
// (You must first `go install github.com/dgodyna/gengo/examples/set-gen`)
package types

//go:generate set-gen -i github.com/dgodyna/gengo/examples/set-gen/sets/types -o github.com/dgodyna/gengo/examples/set-gen/sets

type ReferenceSetTypes struct {
	// These types all cause files to be generated
	a int64
	b int
	c byte
	d string
}
