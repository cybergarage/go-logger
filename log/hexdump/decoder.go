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

// DecodeLine returns the bytes of the specified string.
func DecodeLine(line string) ([]byte, error) {
	if len(line) == 0 {
		return []byte{}, nil
	}
	line = strings.TrimSpace(line)
	reps := []struct {
		from string
		to   string
	}{
		{"   ", " "},
	}
	for _, rep := range reps {
		line = strings.ReplaceAll(line, rep.from, rep.to)
	}
	splitHexes := strings.Split(line, " ")
	offset := 1
	lineHexes := splitHexes[offset : hexdumpTwoColumnBytes+offset+1]
	var bytes []byte
	for _, s := range lineHexes {
		if len(s) == 0 {
			break
		}
		hexByte, err := hex.DecodeString(s)
		if err != nil {
			return bytes, err
		}
		bytes = append(bytes, hexByte...)
	}
	return bytes, nil
}

// DecodeLinesToBytes returns the bytes of the specified string lines.
func DecodeLinesToBytes(lines []string) ([]byte, error) {
	var bytes []byte
	for _, line := range lines {
		lineBytes, err := DecodeLine(line)
		if err != nil {
			return bytes, err
		}
		bytes = append(bytes, lineBytes...)
	}
	return bytes, nil
}

// DecodeLogs decodes the specified hex log.
func DecodeLogs(logs []string) ([]byte, error) {
	logPrefixReg := regexp.MustCompile(fmt.LogPrefixRegex)
	for n, log := range logs {
		if !logPrefixReg.MatchString(log) {
			continue
		}
		logs[n] = logPrefixReg.ReplaceAllString(log, "")
	}
	return DecodeLinesToBytes(logs)
}

// DecodeLogString decodes the specified hex log string.
func DecodeLogString(str string) ([]byte, error) {
	lines := make([]string, 0)
	lines = append(lines, strings.Split(str, "\n")...)
	return DecodeLogs(lines)
}

// DecodeLogBytes decodes the specified hex log bytes.
func DecodeLogBytes(b []byte) ([]byte, error) {
	return DecodeLogString(string(b))
}

// DecodeLogFile decodes the specified hex log file.
func DecodeLogFile(filename string) ([]byte, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return DecodeLogBytes(fileBytes)
}
