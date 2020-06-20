# SPDX-license-identifier: Apache-2.0
##############################################################################
# Copyright (c) 2020 Tech Mahindra
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

PWD := $(shell pwd)
PLATFORM := linux
BINARY := inventory

export GO111MODULE=on

all: build
deploy: build

build: clean 
	CGO_ENABLED=0 GOOS=$(PLATFORM) GOARCH=amd64
	go build -a -ldflags '-extldflags "-static"' \
	-o $(PWD)/$(BINARY) controller/main.go

deploy: build

format:
	@go fmt ./...

clean:
	@rm -f $(BINARY)

