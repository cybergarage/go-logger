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
	"strings"
)

// Level represents a log level.
type Level int

const (
	LevelDebug = Level(1 << 6)
	LevelTrace = Level(1 << 5)
	LevelInfo  = Level(1 << 4)
	LevelWarn  = Level(1 << 3)
	LevelError = Level(1 << 2)
	LevelFatal = Level(1 << 1)
	LevelAll   = Level(0)
)

var logLevelStrings = map[Level]string{
	LevelDebug: "DEBUG",
	LevelTrace: "TRACE",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

// GetLevelString returns a string of the specified log level.
func GetLevelString(logLevel Level) string {
	logString, hasString := logLevelStrings[logLevel]
	if !hasString {
		return loggerLevelUnknownString
	}
	return logString
}

// GetLevelFromString returns a log level of the specified string.
func GetLevelFromString(logLevel string) Level {
	uppderLogLevel := strings.ToUpper(logLevel)
	for level, levelString := range logLevelStrings {
		if strings.HasPrefix(uppderLogLevel, levelString) {
			return level
		}
	}
	return LevelAll
}
