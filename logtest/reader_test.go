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
	"encoding/hex"
	"testing"

	"github.com/cybergarage/go-logger/log/hexdump"
)

//go:embed data/hexdump001.hex
var hexdump001 string

func DecodeHexString(hexStr string) []byte {
	hexBytes, _ := hex.DecodeString(hexStr)
	return hexBytes
}

func TestHexdumpReader(t *testing.T) {
	for _, test := range []struct {
		name     string
		data     string
		expected []byte
	}{
		{
			"hexdump001",
			hexdump001,
			DecodeHexString("570000000a352e372e392d7669746573732d31322e302e360001000000446e2b0b420e6e03000fa22100003b011500000000000000000000321e670779782618521d0150006d7973716c5f6e61746976655f70617373776f726400"),
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			decodedBytes, err := hexdump.DecodeHexdumpString(test.data)
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
