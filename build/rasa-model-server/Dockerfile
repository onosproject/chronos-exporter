# SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0.0 as build

ENV ADAPTER_ROOT=$GOPATH/src/github.com/onosproject/chronos-exporter
ENV CGO_ENABLED=0
ENV GO111MODULE=on

RUN mkdir -p $ADAPTER_ROOT/

COPY . $ADAPTER_ROOT/
COPY build/build-tools /go/src/github.com/onosproject/chronos-exporter/build-tools/
#RUN cd $ADAPTER_ROOT && GO111MODULE=on go build -o /go/bin/rasa-model-server ./cmd/rasa-model-server
RUN cd $ADAPTER_ROOT && GOFLAGS=-mod=vendor make build

FROM rasa/rasa:3.0.4-full as rasa-trainer
COPY rasa-models/rasa /app
WORKDIR /app
RUN /opt/venv/bin/rasa train

FROM alpine:3.12

COPY --from=build /go/src/github.com/onosproject/chronos-exporter/build/_output/chronos-exporter /usr/local/bin/

COPY --from=rasa-trainer /app/models/*.tar.gz /opt/rasa-models/20220101-000000-aether-trained.tar.gz

