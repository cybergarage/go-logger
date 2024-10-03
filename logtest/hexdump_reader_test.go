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
	"io"
	"testing"

	"github.com/cybergarage/go-logger/log/hexdump"
)

func TestHexdumpReader(t *testing.T) {
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
		{
			"mysqldump001",
			mysqldump001File,
			mysqldump001Bytes,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			reader, err := hexdump.NewReaderFromHexdumpString(test.data)
			if err != nil {
				t.Error(err)
				return
			}
			decodedBytes, err := io.ReadAll(reader)
			if err != nil {
				t.Error(err)
				return
			}
			if !bytes.Equal(decodedBytes, test.expected) {
				t.Error(test.name)
				t.Error("expectedBytes")
				for _, b := range hexdump.EncodeBytesToOffsetHexdumpLines(test.expected) {
					t.Error(b)
				}
				t.Error("decodedBytes")
				for _, b := range hexdump.EncodeBytesToOffsetHexdumpLines(decodedBytes) {
					t.Error(b)
				}
				return
			}
		})
	}
}
