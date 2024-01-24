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
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	src := []byte("abcdefghijklmnopqrstuvwxyz1234567890")
	for i := 0; i < len(src); i++ {
		srcBytes := src[0:i]
		encodedLines := EncodeBytesToStringLines(srcBytes)
		decodedBytes := DecodeStringLinesToBytes(encodedLines)
		if !bytes.Equal(decodedBytes, srcBytes) {
			t.Errorf("%s != %s", decodedBytes, srcBytes)
			break
		}
	}
}
