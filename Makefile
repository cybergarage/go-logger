# Copyright (C) 2022 The go-logger Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

SHELL := bash

#PREFIX?=$(shell pwd)
#GOPATH:=$(shell pwd)
#export GOPATH

MODULE_ROOT=github.com/cybergarage/go-logger

PKG_NAME=log
PKG_VER=$(shell git describe --abbrev=0 --tags)
PKG_COVER=${PKG_NAME}-cover

PKG_ID=${MODULE_ROOT}/${PKG_NAME}
PKG_SRC_DIR=${PKG_NAME}
PKG=${MODULE_ROOT}/${PKG_SRC_DIR}

TEST_PKG_NAME=${PKG_NAME}test
TEST_PKG_ID=${MODULE_ROOT}/${TEST_PKG_NAME}
TEST_PKG_DIR=${TEST_PKG_NAME}
TEST_PKG=${MODULE_ROOT}/${TEST_PKG_DIR}

CMD_ROOT=cmd
CMD_PKG_ROOT=${MODULE_ROOT}/${CMD_ROOT}
CMD_SRC_DIR=${CMD_ROOT}
CMD_BINARIES=\
	${CMD_PKG_ROOT}/hexdump2bin \
	${CMD_PKG_ROOT}/hexdump2ascii

BINARIES=\
	${CMD_BINARIES}

.PHONY: format vet lint clean
.IGNORE: lint

all: test

format:
	gofmt -w ${PKG_SRC_DIR} ${TEST_PKG_DIR} ${CMD_SRC_DIR}

vet: format
	go vet ${PKG_ID}

lint: vet
	golangci-lint run ${PKG_SRC_DIR}/... ${TEST_PKG_DIR}/... ${CMD_SRC_DIR}/...

test: lint
	go test -v -p 1 -timeout 10m -cover -coverpkg=${PKG}/... -coverprofile=${PKG_COVER}.out ${PKG}/... ${TEST_PKG}/...
	go tool cover -html=${PKG_COVER}.out -o ${PKG_COVER}.html

install:
	go install -v ${BINARIES}

clean:
	go clean -i ${PKG} ${TEST_PKG}
