# SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0.0 as build

ENV ADAPTER_ROOT=$GOPATH/src/github.com/onosproject/chronos-exporter
ENV CGO_ENABLED=0

RUN mkdir -p $ADAPTER_ROOT/

COPY . $ADAPTER_ROOT/

RUN cd $ADAPTER_ROOT && GO111MODULE=on go build -o /go/bin/rasa-action-server ./cmd/rasa-action-server


FROM alpine:3.12

COPY --from=build /go/bin/rasa-action-server /usr/local/bin/

