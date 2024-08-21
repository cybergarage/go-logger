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
	"strings"
	"testing"

	"github.com/cybergarage/go-logger/log"
	"github.com/cybergarage/go-logger/log/hexdump"
)

func TestHexDecode(t *testing.T) {
	msg := []byte("abcdefghijklmnopqrstuvwxyz1234567890")
	for n := 1; n < len(msg); n++ {
		buffer := bytes.NewBuffer(nil)
		log.SetSharedLogger(log.NewBufferLogger(buffer, log.LevelInfo))
		msgBytes := msg[0:n]
		log.HexInfo(msgBytes)
		logStr := buffer.String()
		logLines := strings.Split(logStr, log.LF)
		decodedBytes, err := hexdump.DecodeHexLogs(logLines)
		if err != nil {
			t.Error(err)
			break
		}
		if !bytes.Equal(decodedBytes, msgBytes) {
			t.Errorf("%s != %s", decodedBytes, msgBytes)
			break
		}
	}
}
