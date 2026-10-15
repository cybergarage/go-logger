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
	std_fmt "fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cybergarage/go-logger/log/fmt"
)

var sharedLogger *Logger
var sharedLoggerMutex = &sync.Mutex{}

type LoggerOutpter func(logger *Logger, level Level, msg string) (int, error)

// Logger writes the messages of the package-level output functions. Its
// level can be changed while other goroutines log.
type Logger struct {
	level      atomic.Int64
	outputer   LoggerOutpter
	data       any
	timeLayout string
	fileMode   uint32
	closer     func() error
}

func newLogger(level Level, outputer LoggerOutpter, data any, opts ...Option) *Logger {
	logger := &Logger{
		outputer:   outputer,
		data:       data,
		timeLayout: TimeLayoutSeconds,
		fileMode:   logFilePerm,
	}
	logger.level.Store(int64(level))
	for _, opt := range opts {
		opt(logger)
	}
	return logger
}

// SetDefault sets the package-level default logger used by package-level output functions.
func SetDefault(logger *Logger) {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	sharedLogger = logger
}

// Default returns the current package-level default logger.
func Default() *Logger {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return sharedLogger
}

// SetSharedLogger sets a singleton logger.
//
// Deprecated: use SetDefault.
func SetSharedLogger(logger *Logger) {
	SetDefault(logger)
}

// GetSharedLogger gets a shared singleton logger.
//
// Deprecated: use Default.
func GetSharedLogger() *Logger {
	return Default()
}

// SetLevel sets a output log level. It may be called while other
// goroutines log.
func (logger *Logger) SetLevel(level Level) {
	logger.level.Store(int64(level))
}

// Level gets the current log level.
func (logger *Logger) Level() Level {
	return Level(logger.level.Load())
}

// IsLevel returns true when the specified log level is enable, otherwise false.
func (logger *Logger) IsLevel(logLevel Level) bool {
	return logLevel >= logger.Level()
}

// Close releases what the logger holds open, such as the file of
// OpenFileLogger, after the message being written. The logger writes
// nothing afterwards. It is a no-op for the other loggers.
func (logger *Logger) Close() error {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	closer := logger.closer
	logger.closer = nil
	if closer == nil {
		return nil
	}
	logger.outputer = outputNothing
	return closer()
}

func outputNothing(_ *Logger, _ Level, _ string) (int, error) {
	return 0, nil
}

func output(outputLevel Level, msgFormat string, msgArgs ...any) int {
	if sharedLogger == nil {
		return 0
	}

	logLevel := sharedLogger.Level()
	if logLevel < outputLevel {
		return 0
	}

	logDate := time.Now().Format(sharedLogger.timeLayout)

	logMsg := std_fmt.Sprintf(msgFormat, msgArgs...)

	outMsgLen := 0
	if 0 < len(logMsg) {
		for lineMsg := range strings.SplitSeq(logMsg, "\n") {
			lineMsg := std_fmt.Sprintf(fmt.LogPrefixFormat, logDate, GetLevelString(outputLevel), lineMsg)
			n, _ := sharedLogger.outputer(sharedLogger, logLevel, lineMsg)
			outMsgLen += n
		}
	}

	return outMsgLen
}
