// Copyright (C) 2019 The go-logger Authors. All rights reserved.
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

/*
hexdump2bin converts the specified ascii hexdump file to the binary file.

	NAME
	 hexdump2bin

	SYNOPSIS
	 hexdump2bin FILE

	RETURN VALUE
	  Return EXIT_SUCCESS or EXIT_FAILURE
*/
//nolint:forbidigo
package main

import (
	"flag"
	"os"

	"github.com/cybergarage/go-logger/log/hexdump"
)

const (
	ProgramName = "hexdump2bin"
)

func usages() {
	println("Usage:")
	println("  " + ProgramName + " <HEX FILE> <BIN FILE>")
	println("")
	println("Return Value:")
	println("  Return EXIT_SUCCESS or EXIT_FAILURE")
	os.Exit(1)
}

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usages()
	}

	hexFileName := args[0]

	hexBytes, err := hexdump.DecodeHexLogFile(hexFileName)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	binFileName := args[1]
	err = os.WriteFile(binFileName, hexBytes, 0600)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
