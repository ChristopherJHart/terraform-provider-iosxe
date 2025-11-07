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

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func LastElement(path string) string {
	pes := strings.Split(path, "/")
	var prefix, element string
	for _, pe := range pes {
		// remove key
		if strings.Contains(pe, "=") {
			pe = pe[:strings.Index(pe, "=")]
		}
		if strings.Contains(pe, ":") {
			prefix = strings.Split(pe, ":")[0]
			element = strings.Split(pe, ":")[1]
		} else {
			element = pe
		}
	}
	return prefix + ":" + element
}

func GetValueSlice(result []gjson.Result) []attr.Value {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return v
}

func GetStringList(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.ListValueMust(types.StringType, v)
}

func GetInt64List(result []gjson.Result) types.List {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.ListValueMust(types.Int64Type, v)
}

func GetStringSet(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.StringValue(result[r].String())
	}
	return types.SetValueMust(types.StringType, v)
}

func GetInt64Set(result []gjson.Result) types.Set {
	v := make([]attr.Value, len(result))
	for r := range result {
		v[r] = types.Int64Value(result[r].Int())
	}
	return types.SetValueMust(types.Int64Type, v)
}

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func RemoveEmptyStrings(s []string) []string {
	var r []string
	for _, v := range s {
		if v != "" {
			r = append(r, v)
		}
	}
	return r
}

// cleanMacAddress removes common delimiters from MAC address and validates format
func cleanMacAddress(mac string) string {
	// Remove common delimiters: colons, dashes, dots
	cleaned := strings.ReplaceAll(mac, ":", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, ".", "")
	cleaned = strings.ToLower(cleaned)

	// Validate that cleaned MAC is exactly 12 hex characters
	matched, _ := regexp.MatchString("^[0-9a-f]{12}$", cleaned)
	if !matched {
		panic(fmt.Errorf("invalid MAC address format: %s (must be 12 hexadecimal characters after removing delimiters)", mac))
	}

	return cleaned
}

// MacAddressToDotted converts MAC address to xxxx.xxxx.xxxx format (Cisco dotted notation)
// Accepts formats: 00:11:22:33:44:55, 00-11-22-33-44-55, 0011.2233.4455
// Returns: 0011.2233.4455
func MacAddressToDotted(mac string) string {
	cleaned := cleanMacAddress(mac)
	return fmt.Sprintf("%s.%s.%s", cleaned[0:4], cleaned[4:8], cleaned[8:12])
}

// MacAddressToColon converts MAC address to xx:xx:xx:xx:xx:xx format (IEEE 802 standard)
// Accepts formats: 00:11:22:33:44:55, 00-11-22-33-44-55, 0011.2233.4455
// Returns: 00:11:22:33:44:55
func MacAddressToColon(mac string) string {
	cleaned := cleanMacAddress(mac)
	parts := make([]string, 6)
	for i := 0; i < 6; i++ {
		parts[i] = cleaned[i*2 : i*2+2]
	}
	return strings.Join(parts, ":")
}

// MacAddressToDash converts MAC address to xx-xx-xx-xx-xx-xx format
// Accepts formats: 00:11:22:33:44:55, 00-11-22-33-44-55, 0011.2233.4455
// Returns: 00-11-22-33-44-55
func MacAddressToDash(mac string) string {
	cleaned := cleanMacAddress(mac)
	parts := make([]string, 6)
	for i := 0; i < 6; i++ {
		parts[i] = cleaned[i*2 : i*2+2]
	}
	return strings.Join(parts, "-")
}
