# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0.0 as build

ENV ADAPTER_ROOT=$GOPATH/src/github.com/onosproject/chronos-exporter
ENV CGO_ENABLED=0

RUN mkdir -p $ADAPTER_ROOT/

COPY . $ADAPTER_ROOT/

RUN cd $ADAPTER_ROOT && GO111MODULE=on go build -o /go/bin/chronos-exporter ./cmd/chronos-exporter

FROM alpine:3.12
RUN apk add bash openssl curl libc6-compat

ENV HOME=/home/chronos-exporter

RUN mkdir $HOME
WORKDIR $HOME

COPY --from=build /go/bin/chronos-exporter /usr/local/bin/
