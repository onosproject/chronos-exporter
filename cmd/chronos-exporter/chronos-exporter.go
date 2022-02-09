// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"github.com/onosproject/chronos-exporter/pkg/manager"
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"io/ioutil"
)

var log = logging.GetLogger("main")

func main() {
	alertPath := flag.String("alert", "", "path to alert file")
	configPath := flag.String("config", "", "path to configuration file")
	aetherPath := flag.String("aether", "", "path to aether configuration file")
	imagePath := flag.String("imagePath", "/opt/", "path to the images")
	sitePlanPath := flag.String("sitePlanPath", "/opt/", "path to the site plans")
	ready := make(chan bool)
	flag.Parse()

	if configPath == nil || *configPath == "" {
		log.Fatal("No config specified. Cannot start")
	}

	if alertPath == nil || *alertPath == "" {
		log.Fatal("No alert specified. Cannot start")
	}

	log.Infof("Starting chronos-exporter with config: %s", *configPath)
	configJSON, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Cannot start without a valid config %v", err)
	}

	log.Infof("Starting chronos-exporter with aether: %s", *configPath)
	aetherJSON, err := ioutil.ReadFile(*aetherPath)
	if err != nil {
		log.Fatalf("Cannot start without a valid aether config %v", err)
	}

	log.Infof("Starting chronos-exporter with alert: %s", *alertPath)
	alertJSON, err := ioutil.ReadFile(*alertPath)
	if err != nil {
		log.Fatalf("Cannot start without a alert file %v", err)
	}

	mgr := manager.NewManager(configJSON, aetherJSON, alertJSON, *imagePath, *sitePlanPath,
		[]string{"http://localhost:4200"})
	mgr.Run(2112)

	<-ready
}
