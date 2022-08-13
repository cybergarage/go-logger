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
	"github.com/cybergarage/go-logger/log/hexdump"
)

// Debug outputs a debug level message to loggers.
func HexDebug(bytes []byte) int {
	return HexOutput(LevelDebug, bytes)
}

// Trace outputs trace level message to loggers.
func HexTrace(bytes []byte) int {
	return HexOutput(LevelTrace, bytes)
}

// Info outputs a information level message to loggers.
func HexInfo(bytes []byte) int {
	return HexOutput(LevelInfo, bytes)
}

// Warn outputs a warning level message to loggers.
func HexWarn(bytes []byte) int {
	return HexOutput(LevelWarn, bytes)
}

// Error outputs a error level message to loggers.
func HexError(bytes []byte) int {
	return HexOutput(LevelError, bytes)
}

// Fatal outputs a fatal level message to loggers.
func HexFatal(bytes []byte) int {
	return HexOutput(LevelFatal, bytes)
}

// Output outputs the specified level message to loggers.
func HexOutput(outputLevel Level, bytes []byte) int {
	nOutput := 0
	lines := hexdump.EncodeToStrings(bytes)
	for _, line := range lines {
		nOutput += output(outputLevel, line)
	}
	return nOutput
}
