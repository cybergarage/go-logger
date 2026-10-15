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
	"os"
)

const (
	// TimeLayoutSeconds is the default time layout of a log line.
	TimeLayoutSeconds = "2006-01-02 15:04:05"
	// TimeLayoutMilliseconds adds milliseconds, for following timings
	// shorter than a second, such as retransmissions.
	TimeLayoutMilliseconds = "2006-01-02 15:04:05.000"
)

// Option configures a logger when it is created.
type Option func(*Logger)

// WithTimeLayout sets the time.Format layout of the timestamp which starts
// each line, such as TimeLayoutMilliseconds. The hexdump decoder reads the
// default layout and the ones which add fractional seconds to it.
func WithTimeLayout(layout string) Option {
	return func(logger *Logger) {
		logger.timeLayout = layout
	}
}

// WithFileMode sets the permissions a file logger creates its file with;
// OpenFileLogger also applies them to a file which exists. The default is
// 0644 for NewFileLogger and 0600 for OpenFileLogger.
func WithFileMode(mode os.FileMode) Option {
	return func(logger *Logger) {
		logger.fileMode = uint32(mode.Perm())
	}
}
