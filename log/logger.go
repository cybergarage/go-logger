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
	"fmt"
	"os"
	"sync"
	"time"
)

type LoggerOutpter func(file string, level Level, msg string) (int, error)

type Level int

type Logger struct {
	File     string
	Level    Level
	outputer LoggerOutpter
}

const (
	LevelDebug = Level(1 << 6)
	LevelTrace = Level(1 << 5)
	LevelInfo  = Level(1 << 4)
	LevelWarn  = Level(1 << 3)
	LevelError = Level(1 << 2)
	LevelFatal = Level(1 << 1)
	LevelAll   = Level(0)
)

const (
	logFormat                = "%s %s %s"
	logFilePerm              = 0644
	LF                       = "\n"
	loggerLevelUnknownString = "UNKNOWN"
	loggerStdout             = "stdout"
)

var sharedLogger *Logger
var sharedLoggerMutex = &sync.Mutex{}

// SetSharedLogger sets a singleton logger.
func SetSharedLogger(logger *Logger) {
	sharedLogger = logger
}

// GetSharedLogger gets a shared singleton logger.
func GetSharedLogger() *Logger {
	return sharedLogger
}

var logLevelStrings = map[Level]string{
	LevelDebug: "DEBUG",
	LevelTrace: "TRACE",
	LevelInfo:  "INFO",
	LevelWarn:  "WARN",
	LevelError: "ERROR",
	LevelFatal: "FATAL",
}

func getLevelString(logLevel Level) string {
	logString, hasString := logLevelStrings[logLevel]
	if !hasString {
		return loggerLevelUnknownString
	}
	return logString
}

// SetLevel sets a output log level.
func (logger *Logger) SetLevel(level Level) {
	logger.Level = level
}

// GetLevel gets the current log level.
func (logger *Logger) GetLevel() Level {
	return logger.Level
}

// IsLevel returns true when the specified log level is enable, otherwise false.
func (logger *Logger) IsLevel(logLevel Level) bool {
	return logLevel >= logger.Level
}

// NewStdoutLogger creates a stdout logger.
func NewStdoutLogger(level Level) *Logger {
	logger := &Logger{
		File:     loggerStdout,
		Level:    level,
		outputer: outputStdout}
	return logger
}

func outputStdout(file string, level Level, msg string) (int, error) {
	fmt.Println(msg)
	return len(msg), nil
}

// NewFileLogger creates a file based logger.
func NewFileLogger(file string, level Level) *Logger {
	logger := &Logger{
		File:     file,
		Level:    level,
		outputer: outputToFile}
	return logger
}

func outputToFile(file string, level Level, msg string) (int, error) {
	msgBytes := []byte(msg + LF)
	fd, err := os.OpenFile(file, (os.O_WRONLY | os.O_CREATE | os.O_APPEND), logFilePerm)
	if err != nil {
		return 0, err
	}

	writer := bufio.NewWriter(fd)
	writer.Write(msgBytes)
	writer.Flush()

	fd.Close()

	return len(msgBytes), nil
}

func output(outputLevel Level, msgFormat string, msgArgs ...interface{}) int {
	if sharedLogger == nil {
		return 0
	}

	logLevel := sharedLogger.GetLevel()
	if logLevel < outputLevel {
		return 0
	}

	t := time.Now()
	logDate := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	headerString := fmt.Sprintf("[%s]", getLevelString(outputLevel))
	logMsg := fmt.Sprintf(logFormat, logDate, headerString, fmt.Sprintf(msgFormat, msgArgs...))
	logMsgLen := len(logMsg)

	if 0 < logMsgLen {
		logMsgLen, _ = sharedLogger.outputer(sharedLogger.File, logLevel, logMsg)
	}

	return logMsgLen
}

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
