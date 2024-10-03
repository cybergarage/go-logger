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

// EncodeByteToASCIIString converts the specified bytes to ASCII strings.
func EncodeByteToASCIIString(r rune) string {
	if unicode.IsPrint(r) {
		return fmt.Sprintf("%c", r)
	}
	return "."
}

// EncodeBytesToASCIIString converts the specified bytes to ASCII strings.
func EncodeBytesToASCIIString(bytes []byte) string {
	str := ""
	for _, b := range bytes {
		str += EncodeByteToASCIIString(rune(b))
	}
	return str
}

// encodeBytesToHexDumpStringLines returns the hexadecimal encoding strings of src with the ASCII repsentations.
func encodeBytesToHexDumpStringLines(src []byte) ([]string, []string) {
	srcLen := len(src)

	hexStrs := make([]string, 0)
	ascStrs := make([]string, 0)

	for offset := 0; offset < srcLen; offset += hexdumpLineColums {
		lineLen := hexdumpLineColums
		if srcLen < (offset + hexdumpLineColums) {
			lineLen = srcLen - offset
		}

		// bytes -> hex strings

		lineHexStr := ""
		for n := 0; n < hexdumpLineColums; n++ {
			if n < lineLen {
				lineHexStr += fmt.Sprintf("%02X ", int(src[offset+n]))
				continue
			}
			lineHexStr += hexdumpLineSep
		}

		hexStrs = append(hexStrs, lineHexStr)

		// bytes -> ASCII strings

		lineAscStr := ""
		for n := 0; n < hexdumpLineColums; n++ {
			if n < lineLen {
				lineAscStr += EncodeByteToASCIIString(rune(src[offset+n]))
				continue
			}
			lineAscStr += hexdumpLineSep
		}

		ascStrs = append(ascStrs, lineAscStr)
	}
	return hexStrs, ascStrs
}

// EncodeBytesToStringLines returns the hexadecimal encoding strings of src with the ASCII repsentations.
func EncodeBytesToStringLines(src []byte) []string {
	hexStrs, strs := encodeBytesToHexDumpStringLines(src)

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
