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
	"strings"
	"unicode"
)

// EncodeByteToASCIIString converts the specified bytes to ASCII strings.
func EncodeByteToASCIIString(r rune) string {
	if unicode.IsPrint(r) {
		return fmt.Sprintf("%c", r)
	}
	return "."
}

// EncodeBytesToASCIIString converts the specified bytes to ASCII strings.
func EncodeBytesToASCIIString(bytes []byte) string {
	var str strings.Builder
	for _, b := range bytes {
		str.WriteString(EncodeByteToASCIIString(rune(b)))
	}
	return str.String()
}

// encodeBytesToHexdumpLines returns the hexadecimal encoding strings of src with the ASCII repsentations.
func encodeBytesToHexdumpLines(src []byte) ([]string, []string) {
	srcLen := len(src)

	hexStrs := make([]string, 0)
	ascStrs := make([]string, 0)

	for offset := 0; offset < srcLen; offset += hexdumpLineColums {
		lineLen := hexdumpLineColums
		if srcLen < (offset + hexdumpLineColums) {
			lineLen = srcLen - offset
		}

		// bytes -> hex strings

		var lineHexStr strings.Builder
		for n := range hexdumpLineColums {
			if n < lineLen {
				lineHexStr.WriteString(fmt.Sprintf("%02X ", int(src[offset+n])))
				continue
			}
			lineHexStr.WriteString(hexdumpLineSep)
		}

		hexStrs = append(hexStrs, lineHexStr.String())

		// bytes -> ASCII strings

		var lineAscStr strings.Builder
		for n := range hexdumpLineColums {
			if n < lineLen {
				lineAscStr.WriteString(EncodeByteToASCIIString(rune(src[offset+n])))
				continue
			}
			lineAscStr.WriteString(hexdumpLineSep)
		}

		ascStrs = append(ascStrs, lineAscStr.String())
	}
	return hexStrs, ascStrs
}

// EncodeBytesToOffsetHexdumpLines returns the hexadecimal encoding strings of src with the ASCII repsentations.
func EncodeBytesToOffsetHexdumpLines(src []byte) []string {
	hexStrs, strs := encodeBytesToHexdumpLines(src)

	lineStrs := make([]string, 0)

	nLines := len(hexStrs)
	for n := 0; n < nLines; n += 2 {
		// line header

		lineStr := fmt.Sprintf("%04X ", n*hexdumpLineColums)

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
