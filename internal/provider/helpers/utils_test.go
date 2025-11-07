// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package helpers

import "testing"

func TestMacAddressToDotted(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "colon format",
			input:    "00:11:22:33:44:55",
			expected: "0011.2233.4455",
		},
		{
			name:     "dash format",
			input:    "00-11-22-33-44-55",
			expected: "0011.2233.4455",
		},
		{
			name:     "dotted format",
			input:    "0011.2233.4455",
			expected: "0011.2233.4455",
		},
		{
			name:     "mixed case",
			input:    "AA:BB:CC:DD:EE:FF",
			expected: "aabb.ccdd.eeff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MacAddressToDotted(tt.input)
			if result != tt.expected {
				t.Errorf("MacAddressToDotted(%s) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMacAddressToColon(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "dotted format",
			input:    "0011.2233.4455",
			expected: "00:11:22:33:44:55",
		},
		{
			name:     "colon format",
			input:    "00:11:22:33:44:55",
			expected: "00:11:22:33:44:55",
		},
		{
			name:     "dash format",
			input:    "00-11-22-33-44-55",
			expected: "00:11:22:33:44:55",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MacAddressToColon(tt.input)
			if result != tt.expected {
				t.Errorf("MacAddressToColon(%s) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMacAddressToDash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "dotted format",
			input:    "0011.2233.4455",
			expected: "00-11-22-33-44-55",
		},
		{
			name:     "colon format",
			input:    "00:11:22:33:44:55",
			expected: "00-11-22-33-44-55",
		},
		{
			name:     "dash format",
			input:    "00-11-22-33-44-55",
			expected: "00-11-22-33-44-55",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MacAddressToDash(tt.input)
			if result != tt.expected {
				t.Errorf("MacAddressToDash(%s) = %s, want %s", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMacAddressInvalid(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "too short",
			input: "00:11:22:33:44",
		},
		{
			name:  "too long",
			input: "00:11:22:33:44:55:66",
		},
		{
			name:  "invalid characters",
			input: "00:11:22:33:44:ZZ",
		},
		{
			name:  "not hex",
			input: "GG:HH:II:JJ:KK:LL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("MacAddressToDotted(%s) should have panicked", tt.input)
				}
			}()
			MacAddressToDotted(tt.input)
		})
	}
}
