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
PKG_ID=${MODULE_ROOT}/${PKG_NAME}
PKG_SRC_DIR=${PKG_NAME}
PKG_SRCS=\
        ${PKG_SRC_DIR}
PKGS=\
	${PKG_ID}

.PHONY: format vet lint clean

all: test

format:
	gofmt -w ${PKG_SRC_DIR}

vet: format
	go vet ${PKG_ID}

lint: vet
	golangci-lint run ${PKG_SRCS}

test: lint
	go test -v -cover -timeout 60s ${PKGS}

clean:
	go clean -i ${PKGS}
