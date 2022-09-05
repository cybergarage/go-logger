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
	"errors"
	"testing"
)

const (
	testLogMessage         = "hello"
	nullOutputErrorMessage = "Shared Logger is null, but message is output"
	outputErrorMessage     = "Message can't be output"
)

func TestNullLogger(t *testing.T) {
	SetSharedLogger(nil)

	nOutput := Tracef(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = Infof(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = Errorf(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = Warnf(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}

	nOutput = Fatalf(testLogMessage)
	if 0 < nOutput {
		t.Error(errors.New(nullOutputErrorMessage))
	}
}

func TestStdoutLogger(t *testing.T) {
	SetSharedLogger(NewStdoutLogger(LevelDebug))
	defer SetSharedLogger(nil)

	nOutput := Debugf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = Tracef(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = Infof(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = Errorf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = Warnf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}

	nOutput = Fatalf(testLogMessage)
	if nOutput <= 0 {
		t.Error(errors.New(outputErrorMessage))
	}
}
