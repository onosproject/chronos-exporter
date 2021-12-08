# SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

# If any command in a pipe has nonzero status, return that status
SHELL = bash -o pipefail

export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

KIND_CLUSTER_NAME           ?= kind
DOCKER_REPOSITORY           ?= onosproject/
ONOS_CHRONOS_EXPORTER_VERSION ?= latest
LOCAL_CHRONOS_EXPORTER         ?=

all: build images

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

images: # @HELP build simulators image
images: chronos-exporter-docker

.PHONY: local-chronos-exporter
local-chronos-exporter:
ifdef LOCAL_CHRONOS_EXPORTER
	rm -rf ./local-chronos-exporter
	cp -a ${LOCAL_CHRONOS_EXPORTER} ./local-chronos-exporter
endif

# @HELP build the go binary in the cmd/chronos-exporter package
build: local-chronos-exporter
	go build -o build/_output/chronos-exporter ./cmd/chronos-exporter

test: build deps license_check_member_only linters
	go test -cover -race github.com/onosproject/chronos-exporter/pkg/...
	go test -cover -race github.com/onosproject/chronos-exporter/cmd/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build deps license_check_member_only linters
	TEST_PACKAGES=github.com/onosproject/chronos-exporter/... ./build/build-tools/build/jenkins/make-unit

chronos-exporter-docker: local-chronos-exporter
	docker build . -f Dockerfile \
	--build-arg LOCAL_AETHER_MODELS=${LOCAL_CHRONOS_EXPORTER} \
	-t ${DOCKER_REPOSITORY}chronos-exporter:${ONOS_CHRONOS_EXPORTER_VERSION}

kind: # @HELP build Docker images and add them to the currently configured kind cluster
kind: images kind-only

kind-only: # @HELP deploy the image without rebuilding first
kind-only:
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image --name ${KIND_CLUSTER_NAME} ${DOCKER_REPOSITORY}chronos-exporter:${ONOS_CHRONOS_EXPORTER_VERSION}

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version ${VERSION} onosproject/chronos-exporter

jenkins-publish: jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/bin/push-images
	./build/build-tools/release-merge-commit

clean:: # @HELP remove all the build artifacts
	rm -rf ./build/_output
	rm -rf ./vendor
	rm -rf ./cmd/chronos-exporter/chronos-exporter

