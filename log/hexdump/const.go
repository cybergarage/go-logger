// Copyright (C) 2018 The go-logger Authors. All rights reserved.
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

const (
	hexdumpLineSep               = "   "
	hexdumpASCIISep              = " "
	hexdumpLineColums            = 8
	hexdumpLineColumnByteLen     = (2 * hexdumpLineColums) + (1 /* space */ * (hexdumpLineColums - 1))
	hexdumpTwoLineColums         = hexdumpLineColums * 2
	hexdumpTwoLineColumnByteLen  = (hexdumpLineColumnByteLen * 2) + len(hexdumpLineSep)
	hexdumpTwoColumnASCIIByteLen = (hexdumpTwoLineColumnByteLen + len(hexdumpLineSep)) + hexdumpLineColums + len(hexdumpASCIISep) + hexdumpLineColums
)
