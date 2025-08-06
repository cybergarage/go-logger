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
hexdump2ascii converts the specified ascii hexdump file to the ascii file.

	NAME
	 hexdump2ascii

	SYNOPSIS
	 hexdump2ascii FILE

	RETURN VALUE
	  Return EXIT_SUCCESS or EXIT_FAILURE
*/
//nolint:forbidigo
package main

import (
	"flag"
	"os"
	"unicode"

	"github.com/cybergarage/go-logger/log/hexdump"
)

const (
	ProgramName = "hexdump2ascii"
)

func usages() {
	println("Usage:")
	println("  " + ProgramName + " <HEX FILE> <OUTPUT FILE>")
	println("")
	println("Return Value:")
	println("  Return EXIT_SUCCESS or EXIT_FAILURE")
}

func main() {
	exitCode := 0
	defer func() {
		os.Exit(exitCode)
	}()

	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usages()
		exitCode = 1
		return
	}

	hexFileName := args[0]

	hexBytes, err := hexdump.DecodeHexdumpFile(hexFileName)
	if err != nil {
		println(err.Error())
		exitCode = 1
		return
	}

	asciiFileName := args[1]
	file, err := os.OpenFile(asciiFileName, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		println(err.Error())
		exitCode = 1
		return
	}

	closeFile := func() {
		if err := file.Close(); err != nil {
			println(err.Error())
			exitCode = 1
		}
	}
	defer closeFile()

	for i := range hexBytes {
		b := hexBytes[i]
		if !unicode.IsPrint(rune(b)) && b != '\n' {
			b = '.'
		}
		_, err := file.Write([]byte{b})
		if err != nil {
			println(err.Error())
			exitCode = 1
			return
		}
	}
}
