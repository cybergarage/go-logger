// Copyright (C) 2018 The go-logger Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"bytes"
)

// NewBufferLogger creates a buffer logger.
func NewBufferLogger(buffer *bytes.Buffer, level Level) *Logger {
	logger := &Logger{
		level:    level,
		outputer: outputBuffer,
		data:     buffer}
	return logger
}

func outputBuffer(logger *Logger, _ Level, msg string) (int, error) {
	buffer, ok := logger.data.(*bytes.Buffer)
	if !ok {
		return 0, nil
	}
	buffer.WriteString(msg + LF)
	return len(msg), nil
}
