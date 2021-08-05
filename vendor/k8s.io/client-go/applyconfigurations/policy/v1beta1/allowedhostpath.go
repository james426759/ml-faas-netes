/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// AllowedHostPathApplyConfiguration represents an declarative configuration of the AllowedHostPath type for use
// with apply.
type AllowedHostPathApplyConfiguration struct {
	PathPrefix *string `json:"pathPrefix,omitempty"`
	ReadOnly   *bool   `json:"readOnly,omitempty"`
}

// AllowedHostPathApplyConfiguration constructs an declarative configuration of the AllowedHostPath type for use with
// apply.
func AllowedHostPath() *AllowedHostPathApplyConfiguration {
	return &AllowedHostPathApplyConfiguration{}
}

// WithPathPrefix sets the PathPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PathPrefix field is set to the value of the last call.
func (b *AllowedHostPathApplyConfiguration) WithPathPrefix(value string) *AllowedHostPathApplyConfiguration {
	b.PathPrefix = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *AllowedHostPathApplyConfiguration) WithReadOnly(value bool) *AllowedHostPathApplyConfiguration {
	b.ReadOnly = &value
	return b
}
