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

// Outputf outputs the specified level message to loggers.
func Outputf(outputLevel Level, format string, args ...interface{}) int {
	sharedLoggerMutex.Lock()
	defer sharedLoggerMutex.Unlock()
	return output(outputLevel, format, args...)
}

// Error outputs a error level message to loggers.
func Error(err error) int {
	return Errorf("%s", err.Error())
}
