// Copyright 2024 oapi-codegen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package nullable provides support for nullable fields in JSON, indicating whether the field is absent, set to null,
// or set to a value.
// Originally from https://github.com/oapi-codegen/nullable.
//
// Unlike other known implementations, this makes it possible to both marshal and unmarshal the value,
// as well as represent all three states:
// - the field is not set
// - the field is explicitly set to null
// - the field is explicitly set to a given value
//
// And can be embedded in structs, for instance with the following definition:
//
//	obj := struct {
//			// RequiredID is a required, nullable field
//			RequiredID     nullable.Nullable[int]     `json:"id"`
//			// OptionalString is an optional, nullable field
//			// NOTE that no pointer is required, only `omitempty`
//			OptionalString nullable.Nullable[string] `json:"optionalString,omitempty"`
//	}{}
package nullable
