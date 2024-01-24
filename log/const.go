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

const (
	logPrefixFormat          = "%s [%s] %s"
	logPrefixDateFormat      = "%04d-%02d-%02d %02d:%02d:%02d"
	logPrefixRegex           = `^(\d{4})-\d{2}-\d{2} \d{2}:\d{2}:\d{2} \[(\w+)\] `
	logFilePerm              = 0644
	LF                       = "\n"
	loggerLevelUnknownString = "UNKNOWN"
)
