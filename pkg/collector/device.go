// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"time"
)

func (d *Device) collect(period time.Duration, site string) {
	log.Infof("Starting collector for Device %s", d.SerialNumber)
	go func() {
			if d.Sim != nil {
				for {
					count := float64(rand.Intn(100))
					if count != 5 {
						deviceConnectedStatus.WithLabelValues("Active", site, *d.Sim, d.SerialNumber).Set(1)
					} else {
						deviceConnectedStatus.WithLabelValues("Active", site, *d.Sim, d.SerialNumber).Set(0)
					}
					time.Sleep(period * 3)
				}
			}
	}()

	go func() {
			if d.Sim != nil {
				for {
					randNum := rand.Intn(100)
					if randNum  > 20 && randNum < 30 {
						deviceConnectionEventCore.WithLabelValues("some core event", "Red", site, *d.Sim, d.SerialNumber).Set(1)
					}
				}
			}
	}()

	go func() {
		if d.Sim != nil {
			for {
				randNum := rand.Intn(100)
				if randNum > 10 && randNum < 15 {
					deviceConnectionEventRan.WithLabelValues("some ran event", "Blue", site, *d.Sim, d.SerialNumber).Set(1)
				}
			}
		}
	}()

	go func() {
		if d.Sim != nil {
			for {
				randNum := rand.Intn(100)
				if randNum > 60 && randNum < 70 {
					deviceConnectionEventFabric.WithLabelValues("some fabric event", "Green", site, *d.Sim, d.SerialNumber).Set(1)
				}
			}
		}
	}()
}

var (
	deviceConnectedStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connected_status",
		Help: "Device Status",
	}, []string{"device_status", "site", "iccid", "serial_number"})

	deviceConnectionEventCore = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connection_event_core",
		Help: "Device Connection Event Core",
	}, []string{"msg", "colour", "site", "iccid", "serial_number"})

	deviceConnectionEventRan = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connection_event_ran",
		Help: "Device Connection Event Ran",
	}, []string{"msg", "colour", "site", "iccid", "serial_number"})

	deviceConnectionEventFabric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connection_event_fabric",
		Help: "Device Connection Event Fabric",
	}, []string{"msg", "colour", "site", "iccid", "serial_number"})
)
