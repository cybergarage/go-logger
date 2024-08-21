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
	"os"
	"regexp"
	"strings"

	"github.com/cybergarage/go-logger/log/fmt"
)

// DecodeStringLineToBytes returns the bytes of the specified string.
func DecodeStringLineToBytes(line string) ([]byte, error) {
	if len(line) == 0 {
		return []byte{}, nil
	}
	splitHexes := strings.Split(line, " ")
	lineHexes := splitHexes[1 : hexdumpTwoColumnBytes+3]
	var bytes []byte
	for _, s := range lineHexes {
		if len(s) == 0 {
			continue
		}
		hexByte, err := hex.DecodeString(s)
		if err != nil {
			return bytes, err
		}
		bytes = append(bytes, hexByte...)
	}
	return bytes, nil
}

// DecodeStringLinesToBytes returns the bytes of the specified string lines.
func DecodeStringLinesToBytes(lines []string) ([]byte, error) {
	var bytes []byte
	for _, line := range lines {
		lineBytes, err := DecodeStringLineToBytes(line)
		if err != nil {
			return bytes, err
		}
		bytes = append(bytes, lineBytes...)
	}
	return bytes, nil
}

// DecodeHexLogs decodes the specified hex log.
func DecodeHexLogs(logs []string) ([]byte, error) {
	logPrefixReg := regexp.MustCompile(fmt.LogPrefixRegex)
	for n, log := range logs {
		if !logPrefixReg.MatchString(log) {
			continue
		}
		logs[n] = logPrefixReg.ReplaceAllString(log, "")
	}
	return DecodeStringLinesToBytes(logs)
}

// DecodeHexLogFile decodes the specified hex log file.
func DecodeHexLogFile(filename string) ([]byte, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0)
	lines = append(lines, strings.Split(string(fileBytes), "\n")...)
	return DecodeHexLogs(lines)
}
