// Copyright (C) 2024 The go-logger Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hexdump

import (
	"bytes"
	"io"
)

// NewReaderFromHexdumpFile returns a new Reader reading from the specified file.
func NewReaderFromHexdumpFile(filename string) (io.Reader, error) {
	hexBytes, err := DecodeHexdumpFile(filename)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(hexBytes), nil
}

// NewReaderFromHexdumpBytes returns a new Reader reading from the specified bytes.
func NewReaderFromHexdumpBytes(b []byte) (io.Reader, error) {
	hexBytes, err := DecodeHexdumpBytes(b)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(hexBytes), nil
}

// NewReaderFromHexdumpString returns a new Reader reading from the specified string.
func NewReaderFromHexdumpString(str string) (io.Reader, error) {
	hexBytes, err := DecodeHexdumpString(str)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(hexBytes), nil
}
