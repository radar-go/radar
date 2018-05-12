# Copyright (C) 2017 Radar team (see AUTHORS)
#
# This file is part of radar.
#
# radar is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# radar is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with radar. If not, see <http://www.gnu.org/licenses/>.

# The binary to build (just the basename).
BIN := radar

# This repo's root import path (under GOPATH).
PKG := github.com/radar-go/radar

# Where to push the docker image.
REGISTRY ?= radar-go

# Which architecture to build - see $(ALL_ARCH) for options.
ARCH ?= amd64

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)

###
### These variables should not need tweaking.
###

SRC_DIRS := . # directories which hold app source (not vendored)

ALL_ARCH := amd64 arm arm64 ppc64le

CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

# Set default base image dynamically for each arch
ifeq ($(ARCH),amd64)
    BASEIMAGE?=alpine
endif
ifeq ($(ARCH),arm)
    BASEIMAGE?=armel/busybox
endif
ifeq ($(ARCH),arm64)
    BASEIMAGE?=aarch64/busybox
endif
ifeq ($(ARCH),ppc64le)
    BASEIMAGE?=ppc64le/busybox
endif

IMAGE := $(REGISTRY)/$(BIN)-$(ARCH)

BUILD_IMAGE ?= golang:1.10-alpine

# If you want to build all binaries, see the 'all-build' rule.
# If you want to build all containers, see the 'all-container' rule.
# If you want to build AND push all containers, see the 'all-push' rule.
all: build

build-%:
	@$(MAKE) --no-print-directory ARCH=$* build

container-%:
	@$(MAKE) --no-print-directory ARCH=$* container

push-%:
	@$(MAKE) --no-print-directory ARCH=$* push

all-build: $(addprefix build-, $(ALL_ARCH))

all-container: $(addprefix container-, $(ALL_ARCH))

all-push: $(addprefix push-, $(ALL_ARCH))

build: bin/$(ARCH)/$(BIN)

bin/$(ARCH)/$(BIN): build-dirs
	@echo "building: $@"
	@docker run                                                            \
	    -ti                                                                \
	    --rm                                                               \
	    -u $$(id -u):$$(id -g)                                             \
	    -v $$(pwd)/.go:/go                                                 \
	    -v $$(pwd):/go/src/$(PKG)                                          \
	    -v $$(pwd)/bin/$(ARCH):/go/bin                                     \
	    -v $$(pwd)/bin/$(ARCH):/go/bin/linux_$(ARCH)                       \
	    -v $$(pwd)/.go/std/$(ARCH):/usr/local/go/pkg/linux_$(ARCH)_static  \
	    -w /go/src/$(PKG)                                                  \
	    $(BUILD_IMAGE)                                                     \
	    /bin/sh -c "                                                       \
	        ARCH=$(ARCH)                                                   \
	        VERSION=$(VERSION)                                             \
	        PKG=$(PKG)                                                     \
	        BIN=$(BIN)                                                     \
	        ./build/build.sh                                               \
	    "

DOTFILE_IMAGE = $(subst /,_,$(IMAGE))-$(VERSION)

container: .container-$(DOTFILE_IMAGE) container-name
.container-$(DOTFILE_IMAGE): bin/$(ARCH)/$(BIN) Dockerfile.in
	@sed \
	    -e 's|ARG_BIN|$(BIN)|g' \
	    -e 's|ARG_ARCH|$(ARCH)|g' \
	    -e 's|ARG_FROM|$(BASEIMAGE)|g' \
	    Dockerfile.in > .dockerfile-$(ARCH)
	@docker build -t $(IMAGE):$(VERSION) -f .dockerfile-$(ARCH) .
	@docker images -q $(IMAGE):$(VERSION) > $@

container-name:
	@echo "container: $(IMAGE):$(VERSION)"

push: .push-$(DOTFILE_IMAGE) push-name
.push-$(DOTFILE_IMAGE): .container-$(DOTFILE_IMAGE)
	@gcloud docker push $(IMAGE):$(VERSION)
	@docker images -q $(IMAGE):$(VERSION) > $@

push-name:
	@echo "pushed: $(IMAGE):$(VERSION)"

version:
	@echo $(VERSION)

tests: build-dirs
	@echo "Running tests"
	@docker run                                                            \
	    -ti                                                                \
	    --rm                                                               \
	    -u $$(id -u):$$(id -g)                                             \
	    -v $$(pwd)/.go:/go                                                 \
	    -v $$(pwd):/go/src/$(PKG)                                          \
	    -v $$(pwd)/bin/$(ARCH):/go/bin                                     \
	    -v $$(pwd)/bin/$(ARCH):/go/bin/linux_$(ARCH)                       \
	    -v $$(pwd)/.go/std/$(ARCH):/usr/local/go/pkg/linux_$(ARCH)_static  \
	    -w /go/src/$(PKG)                                                  \
	    $(BUILD_IMAGE)                                                     \
	    /bin/sh -c "                                                       \
	        ./build/test.sh $(SRC_DIRS)                                    \
	    "

build-dirs:
	@mkdir -p bin/$(ARCH)
	@mkdir -p .go/src/$(PKG) .go/pkg .go/bin .go/std/$(ARCH)
	@qtc -dir $(CURRENT_DIR)/ui/web/templates

clean: container-clean bin-clean files-clean

container-clean:
	@rm -rf .container-* .dockerfile-* .push-*

start:
	@if [ ! -f $(CURRENT_DIR)/.radar.pid ]; then \
		echo -n "\\033[1;35m+++ Startng radar\\033[39;0m "; \
		$(CURRENT_DIR)/bin/$(ARCH)/$(BIN) -alsologtostderr=true 1> radar.log < /dev/null 2>&1 & \
		echo $$! > $(CURRENT_DIR)/.radar.pid ; \
		while ! curl localhost:20000/healthcheck > /dev/null 2>&1; do \
			/bin/sleep 1; \
			echo -n "\\033[1;35m.\\033[39;0m"; \
		done; \
		echo; \
	fi

stop:
	@if [ -f $(CURRENT_DIR)/.radar.pid ]; then \
		echo -n "\\033[1;35m+++ Stopping web\\033[39;0m "; \
		kill -s 15 `cat $(CURRENT_DIR)/.radar.pid`; \
		while curl localhost:20000/healthcheck > /dev/null 2>&1; do \
			/bin/sleep 1; \
			echo -n "\\033[1;35m.\\033[39;0m"; \
		done; \
		rm -f $(CURRENT_DIR)/.radar.pid; \
		echo; \
	fi

restart:
	$(MAKE) stop
	$(MAKE) start

build-local: build-dirs
	@go build -o bin/$(ARCH)/$(BIN) cmd/radar/main.go

rpm build-rpm: build-local
	@cp bin/$(ARCH)/$(BIN) rpm/
	@rpmbuild --define "_sourcedir $(CURRENT_DIR)/rpm" -bb rpm/radar.spec

bin-clean:
	@rm -rf .go bin

files-clean:
	@rm -fr cpu-*.log mem-*.log block-*.log *.test
	@rm -fr radar.log

check-cpu-tests:
	@go tool pprof -text -nodecount=10 ./radar.test cpu-*.log

check-block-tests:
	@go tool pprof -text -nodecount=10 ./radar.test block-*.log

check-mem-tests:
	@go tool pprof -text -nodecount=10 ./radar.test mem-*.log

check-tests:
	@echo "CPU tests"
	@+make check-cpu-tests
	@echo "Block tests"
	@+make check-block-tests
	@echo "Mem tests"
	@+make check-mem-tests

update-vendors:
	@dep ensure

list-packages:
	@go list ./...

PACKAGE?=github.com/radar-go/radar
update-golden-files:
	@go test $(PACKAGE) -update
