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
	"bufio"
	"os"
)

// NewFileLogger creates a file based logger.
func NewFileLogger(file string, level Level) *Logger {
	logger := &Logger{
		level:    level,
		outputer: outputToFile,
		data:     file}
	return logger
}

func outputToFile(logger *Logger, _ Level, msg string) (int, error) {
	filename, ok := logger.data.(string)
	if !ok {
		return 0, nil
	}

	msgBytes := []byte(msg + LF)
	fd, err := os.OpenFile(filename, (os.O_WRONLY | os.O_CREATE | os.O_APPEND), logFilePerm)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(fd)
	writer.Write(msgBytes)
	writer.Flush()

	fd.Close()

	return len(msgBytes), nil
}
