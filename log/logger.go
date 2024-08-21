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
	"time"

	"github.com/cybergarage/go-logger/log/fmt"
)

var sharedLogger *Logger
var sharedLoggerMutex = &sync.Mutex{}

type LoggerOutpter func(logger *Logger, level Level, msg string) (int, error)

type Logger struct {
	level    Level
	outputer LoggerOutpter
	data     any
}

// SetSharedLogger sets a singleton logger.
func SetSharedLogger(logger *Logger) {
	sharedLogger = logger
}

// GetSharedLogger gets a shared singleton logger.
func GetSharedLogger() *Logger {
	return sharedLogger
}

// SetLevel sets a output log level.
func (logger *Logger) SetLevel(level Level) {
	logger.level = level
}

// Level gets the current log level.
func (logger *Logger) Level() Level {
	return logger.level
}

// IsLevel returns true when the specified log level is enable, otherwise false.
func (logger *Logger) IsLevel(logLevel Level) bool {
	return logLevel >= logger.level
}

func output(outputLevel Level, msgFormat string, msgArgs ...interface{}) int {
	if sharedLogger == nil {
		return 0
	}

	logLevel := sharedLogger.Level()
	if logLevel < outputLevel {
		return 0
	}

	t := time.Now()
	logDate := std_fmt.Sprintf(fmt.LogPrefixDateFormat,
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	logMsg := std_fmt.Sprintf(msgFormat, msgArgs...)

	outMsgLen := 0
	if 0 < len(logMsg) {
		for _, lineMsg := range strings.Split(logMsg, "\n") {
			lineMsg := std_fmt.Sprintf(fmt.LogPrefixFormat, logDate, GetLevelString(outputLevel), lineMsg)
			n, _ := sharedLogger.outputer(sharedLogger, logLevel, lineMsg)
			outMsgLen += n
		}
	}

	return outMsgLen
}
