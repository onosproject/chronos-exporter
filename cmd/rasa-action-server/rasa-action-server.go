// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package main

import (
	"github.com/onosproject/chronos-exporter/pkg/rasa_action"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("main")

func main() {
	//modelPath := flag.String("modelPath", "/opt/rasa-models/", "path to the models")
	ready := make(chan bool)
	//flag.Parse()

	log.Infof("Starting Rasa Action Server")
	rs := rasa_action.NewRasaAction()
	rs.Run()

	<-ready
}
