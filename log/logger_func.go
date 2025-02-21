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

// Debugf outputs a debug level message to loggers.
func Debugf(format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelDebug, format, args...)
}

// Tracef outputs trace level message to loggers.
func Tracef(format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelTrace, format, args...)
}

// Infof outputs a information level message to loggers.
func Infof(format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelInfo, format, args...)
}

// Warnf outputs a warning level message to loggers.
func Warnf(format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelWarn, format, args...)
}

// Errorf outputs a error level message to loggers.
func Errorf(format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelError, format, args...)
}

// Fatalf outputs a fatal level message to loggers.
func Fatalf(format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelFatal, format, args...)
}

// Outputf outputs the specified level message to loggers
func Outputf(outputLevel Level, format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(outputLevel, format, args...)
}

// Debug outputs debug level message to loggers.
func Debug(s string) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelDebug, "%s", s)
}

// Trace outputs trace level message to loggers.
func Trace(s string) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelTrace, "%s", s)
}

// Info outputs a information level message to loggers.
func Info(s string) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelInfo, "%s", s)
}

// Warn outputs a warning level message to loggers.
func Warn(s string) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelWarn, "%s", s)
}

// Error outputs a error level message to loggers.
func Error(err error) int {
	return Errorf("%s", err.Error())
}

// Fatal outputs a fatal level message to loggers.
func Fatal(s string) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(LevelFatal, "%s", s)
}

// Output outputs the specified level message to loggers.
func Output(outputLevel Level, s string) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(outputLevel, "%s", s)
}
