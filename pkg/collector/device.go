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
	if d.Sim == nil {
		return
	}
	sn := d.SerialNumber
	sim := *d.Sim
	go func() {
		for {
			count := float64(rand.Intn(100))
			if count != 5 {
				deviceConnectedStatus.WithLabelValues("Active", site, sim, sn).Set(1)
			} else {
				deviceConnectedStatus.WithLabelValues("Active", site, sim, sn).Set(0)
			}
				time.Sleep(period * 3)
		}
	}()

	go func() {
		for {
			randNum := rand.Intn(100)
			if randNum  > 20 && randNum < 30 {
				deviceConnectionEventCore.WithLabelValues("some core event", site, sim, sn).
					Set(float64(rand.Intn(5) +1 ))
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(60) +1))
		}
	}()

	go func() {
		for {
			randNum := rand.Intn(100)
			if randNum > 10 && randNum < 15 {
				deviceConnectionEventRan.WithLabelValues("some ran event", site, sim, sn).
					Set(float64(rand.Intn(5) + 1))
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(60) + 1))
		}
	}()

	go func() {
		for {
			randNum := rand.Intn(100)
			if randNum > 60 && randNum < 70 {
				deviceConnectionEventFabric.WithLabelValues("some fabric event", site, sim, sn).
					Set(float64(rand.Intn(5) + 1))
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(60) + 1))
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
	}, []string{"msg", "site", "iccid", "serial_number"})

	deviceConnectionEventRan = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connection_event_ran",
		Help: "Device Connection Event Ran",
	}, []string{"msg", "site", "iccid", "serial_number"})

	deviceConnectionEventFabric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "device_connection_event_fabric",
		Help: "Device Connection Event Fabric",
	}, []string{"msg", "site", "iccid", "serial_number"})
)
