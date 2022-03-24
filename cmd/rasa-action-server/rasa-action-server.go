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
	smtpServer := flag.String("smtpServer", "docker-compose_mailhog", "smtp server url")
	smtpServerPort := flag.Int("smtpSerPort", 1025, "smtp server port")
	port := flag.String("port", "8081", "rasa action server port")
	smtpUser := flag.String("smtpUser", "", "smtp user")
	smtpUserPass := flag.String("smtpUserPass", "", "smtp user password")
	ready := make(chan bool)
	flag.Parse()

	log.Infof("Starting Rasa Action Server")
	rs := rasa_action.NewRasaAction(*smtpServer, *smtpServerPort, *port, *smtpUser, *smtpUserPass)
	rs.Run()

	<-ready
}
