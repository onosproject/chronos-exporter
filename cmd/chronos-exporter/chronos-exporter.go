// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package main

import (
	"github.com/onosproject/chronos-exporter/pkg/collector"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func main() {
	collector.RecordDeviceStatusMetrics(2*time.Second, "starbucks-newyork-cameras", "8991101200003204514")
	collector.RecordDeviceStatusMetrics(2*time.Second, "acme-chicago-robots", "8991101200003204515")

	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":2112", nil); err != nil {
		log.Fatal(err)
	}
}
