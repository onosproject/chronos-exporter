// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
	"math/rand"
)

// RecordDeviceStatusMetrics records status of the device
func RecordDeviceStatusMetrics(period time.Duration, site string, iccid string) {
	go func() {
		for {
			count := float64(rand.Intn(10))
			if(count != 5) {
				deviceConnectedStatus.WithLabelValues("Active", site, iccid).Set(1)
			} else {
				deviceConnectedStatus.WithLabelValues("Inactive", site, iccid).Set(0)
			}
			time.Sleep(period)
		}
	}()
}

var (
	deviceConnectedStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connected_status",
		Help: "Device Status",
	}, []string{"device_status", "site", "iccid"})
)
