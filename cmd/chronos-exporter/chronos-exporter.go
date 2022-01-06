// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package main

import (
	"flag"
	"github.com/onosproject/chronos-exporter/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io/ioutil"
)

var log = logging.GetLogger("main")

func main() {
	configPath := flag.String("config", "", "path to configuration file")
	imagePath := flag.String("imagePath", "/opt/images/", "path to the images")
	sitePlanPath := flag.String("sitePlanPath", "/opt/site-plans/", "path to the site plans")
	ready := make(chan bool)
	flag.Parse()

	if configPath == nil || *configPath == "" {
		log.Fatal("No config specified. Cannot start")
	}

	log.Infof("Starting chronos-exporter with config: %s", *configPath)
	configJSON, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Cannot start without a valid config %v", err)
	}

	mgr := manager.NewManager(configJSON, *imagePath, *sitePlanPath)
	mgr.Run()

	<-ready
}
