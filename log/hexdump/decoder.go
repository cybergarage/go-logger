// Copyright (C) 2018 The go-logger Authors. All rights reserved.
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
package hexdump

import (
	"encoding/hex"
	"strings"
)

// DecodeStringToBytes returns the bytes of the specified string.
func DecodeStringToBytes(src string) []byte {
	hexes := strings.Split(src, " ")
	var bytes []byte
	for n, s := range hexes {
		if n == 0 {
			continue
		}
		hexByte, err := hex.DecodeString(s)
		if err != nil {
			panic(err)
		}
		bytes = append(bytes, hexByte...)
	}
	return bytes
}

// DecodeStringLinesToBytes returns the bytes of the specified string lines.
func DecodeStringLinesToBytes(lines []string) []byte {
	var bytes []byte
	for _, line := range lines {
		bytes = append(bytes, DecodeStringToBytes(line)...)
	}
	return bytes
}
