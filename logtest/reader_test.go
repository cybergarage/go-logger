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

//go:embed data/hexdump002.hex
var hexdump002 string

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
		{
			"hexdump002",
			hexdump002,
			DecodeHexString("ef0000000300000003000000010000000800000000000000000000000000000001000000cb000000106d617842736f6e4f626a65637453697a650000000001106d61785772697465426174636853697a6500a0860100106d696e5769726556657273696f6e0000000000106d61785769726556657273696f6e0007000000016f6b00000000000000f03f08726561644f6e6c790000096c6f63616c54696d65003f64296400000000106c6f676963616c53657373696f6e54696d656f75744d696e75746573001e0000000869736d61737465720001106d61784d65737361676553697a65427974657300006cdc0200"),
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
