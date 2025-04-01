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

package logtest

import (
	"errors"
	"testing"

	"github.com/cybergarage/go-logger/log"
)

const (
	testLogMessage         = "hello"
	nullOutputErrorMessage = "Shared Logger is null, but message is output"
	outputErrorMessage     = "Message can't be output"
)

func TestNullLogger(t *testing.T) {
	log.SetSharedLogger(nil)

	nOutput := log.Tracef(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = log.Infof(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = log.Errorf(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = log.Warnf(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = log.Fatalf(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = log.Error(errors.New(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	// Hex*()

	nOutput = log.HexDebug([]byte(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexTrace([]byte(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexInfo([]byte(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexError([]byte(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexWarn([]byte(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexFatal([]byte(testLogMessage))
	if 0 < nOutput {
		t.Error(errors.New(outputErrorMessage))
	}
}

func TestStdoutLogger(t *testing.T) {
	log.SetSharedLogger(log.NewStdoutLogger(log.LevelTrace))

	nOutput := log.Tracef(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.Debugf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.Infof(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.Errorf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.Warnf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.Fatalf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.Error(errors.New(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	// Hex*()

	nOutput = log.HexTrace([]byte(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexDebug([]byte(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexInfo([]byte(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexError([]byte(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexWarn([]byte(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = log.HexFatal([]byte(testLogMessage))
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}
}
