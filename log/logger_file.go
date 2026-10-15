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
	"path/filepath"
)

// NewFileLogger creates a logger which appends each line to file, opening
// and closing it for every line, so that the file can be rotated or
// removed while the logger is in use. It creates the file with mode 0644
// unless WithFileMode says otherwise.
func NewFileLogger(file string, level Level, opts ...Option) *Logger {
	return newLogger(level, outputToFile, file, opts...)
}

func outputToFile(logger *Logger, _ Level, msg string) (int, error) {
	filename, ok := logger.data.(string)
	if !ok {
		return 0, nil
	}

	msgBytes := []byte(msg + LF)
	fd, err := os.OpenFile(filepath.Clean(filename), (os.O_WRONLY | os.O_CREATE | os.O_APPEND), os.FileMode(logger.fileMode))
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(fd)
	writer.Write(msgBytes)
	writer.Flush()

	fd.Close()

	return len(msgBytes), nil
}

// OpenFileLogger opens file once, appending to it and creating it if
// needed, and returns a logger which writes each line to it until Close.
// It sets the file's permissions to 0600, or to those of WithFileMode, even
// when the file exists, since a log may hold what others should not read.
func OpenFileLogger(file string, level Level, opts ...Option) (*Logger, error) {
	logger := newLogger(level, outputWriter, nil, append([]Option{WithFileMode(openFileLoggerPerm)}, opts...)...)
	mode := os.FileMode(logger.fileMode)
	fd, err := os.OpenFile(filepath.Clean(file), os.O_WRONLY|os.O_CREATE|os.O_APPEND, mode)
	if err != nil {
		return nil, err
	}
	if err := fd.Chmod(mode); err != nil {
		_ = fd.Close()
		return nil, err
	}
	logger.data = fd
	logger.closer = fd.Close
	return logger, nil
}
