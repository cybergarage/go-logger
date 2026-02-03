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

package logtest

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/cybergarage/go-logger/log/hexdump"
)

func TestHexdumpBytes(t *testing.T) {
	for _, test := range []struct {
		name     string
		data     string
		expected []byte
	}{
		{
			"hexdump001",
			hexdump001File,
			hexdump001Bytes,
		},
		{
			"hexdump002",
			hexdump002File,
			hexdump002Bytes,
		},
		{
			"hexdump003",
			hexdump003File,
			hexdump003Bytes,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			decodedBytes, err := hexdump.NewBytesWithHexdumpString(test.data)
			if err != nil {
				t.Error(err)
				return
			}
			if !bytes.Equal(decodedBytes, test.expected) {
				t.Errorf("[%d] %s != [%d] %s", len(decodedBytes), hexdump.EncodeBytesToASCIIString(decodedBytes), len(test.expected), hexdump.EncodeBytesToASCIIString(test.expected))
				return
			}
		})
	}
}
