// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"github.com/onosproject/chronos-exporter/pkg/rasa"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("main")

func main() {
	modelPath := flag.String("modelPath", "/opt/rasa-models/", "path to the models")
	ready := make(chan bool)
	flag.Parse()

	log.Infof("Starting Rasa Server")
	rs := rasa.NewRasa(*modelPath)
	rs.Run()

	<-ready
}
