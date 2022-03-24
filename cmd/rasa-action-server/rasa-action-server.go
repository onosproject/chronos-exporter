// SPDX-FileCopyrightText: 2022-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"github.com/onosproject/chronos-exporter/pkg/rasa_action"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("main")

func main() {
	smtpServer := flag.String("smtpServer", "smtp.gmail.com", "smtp server url")
	smtpServerPort := flag.Int("smtpSerPort", 587, "smtp server port")
	port := flag.String("port", "8081", "rasa action server port")
	ready := make(chan bool)
	flag.Parse()

	log.Infof("Starting Rasa Action Server")
	rs := rasa_action.NewRasaAction(*smtpServer, *smtpServerPort, *port)
	rs.Run()

	<-ready
}
