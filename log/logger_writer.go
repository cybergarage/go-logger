// Copyright (C) 2026 The go-logger Authors. All rights reserved.
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
	"io"
)

// NewWriterLogger creates a logger which writes each line, ended by LF, to
// w, such as an io.MultiWriter of a file and a terminal pane. Lines are
// written one at a time, so w needs no locking of its own.
func NewWriterLogger(w io.Writer, level Level, opts ...Option) *Logger {
	return newLogger(level, outputWriter, w, opts...)
}

func outputWriter(logger *Logger, _ Level, msg string) (int, error) {
	w, ok := logger.data.(io.Writer)
	if !ok {
		return 0, nil
	}
	return io.WriteString(w, msg+LF)
}
