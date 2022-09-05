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
	"fmt"
	"unicode"
)

const (
	hexdumpLineBytes             = 8
	hexdumpTwoColumnBytes        = hexdumpLineBytes * 2
	hexdumpTwoColumnASCIIByteLen = 77
)

// encodeToStringsWithStrings returns the hexadecimal encoding strings of src with the ASCII repsentations.
func encodeToStringsWithStrings(src []byte) ([]string, []string) {
	srcLen := len(src)

	hexStrs := make([]string, 0)
	strs := make([]string, 0)

	for offset := 0; offset < srcLen; offset += hexdumpLineBytes {
		lineLen := hexdumpLineBytes
		if srcLen < (offset + hexdumpLineBytes) {
			lineLen = srcLen - offset
		}

		// bytes -> hex strings

		lineHexStr := ""
		for n := 0; n < hexdumpLineBytes; n++ {
			if n < lineLen {
				lineHexStr += fmt.Sprintf("%02X ", int(src[offset+n]))
				continue
			}
			lineHexStr += "   "
		}

		hexStrs = append(hexStrs, lineHexStr)

		// bytes -> ASCII strings

		lineStr := ""
		for n := 0; n < hexdumpLineBytes; n++ {
			if n < lineLen {
				r := rune(src[offset+n])
				if unicode.IsPrint(r) {
					lineStr += fmt.Sprintf("%c", r)
				} else {
					lineStr += "."
				}
				continue
			}
			lineStr += "   "
		}

		strs = append(strs, lineStr)
	}
	return hexStrs, strs
}

// EncodeToStrings returns the hexadecimal encoding strings of src with the ASCII repsentations.
func EncodeToStrings(src []byte) []string {
	hexStrs, strs := encodeToStringsWithStrings(src)

	lineStrs := make([]string, 0)

	nLines := len(hexStrs)
	for n := 0; n < nLines; n += 2 {
		// line header

		lineStr := fmt.Sprintf("%04X ", n*hexdumpLineBytes)

		// two hex string lines

		lineStr += hexStrs[n] + "  "
		if (n + 1) < nLines {
			lineStr += hexStrs[n+1]
		} else {
			lineStr += "                        "
		}

		lineStr += "    "

		// two string lines

		lineStr += strs[n] + " "
		if (n + 1) < nLines {
			lineStr += strs[n+1] + " "
		} else {
			lineStr += "        "
		}

		lineStrs = append(lineStrs, lineStr)
	}

	return lineStrs
}
