// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package aether

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"sync"
	"time"
)

var counter int = 0

func (d *EnterprisesEnterpriseSiteDevice) collect(entId string, siteId string) {
	log.Infof("Starting collector for Device %s", d.DeviceId)
	var mu sync.Mutex
	if d.SimCard == nil {
		return
	}
	devId := d.DeviceId
	simCard := *d.SimCard
	imei := fmt.Sprintf("%d", d.Imei)
	go func() {
		for {
			count := float64(rand.Intn(100))
			if count != 5 {
				deviceConnectedStatus.WithLabelValues("Active", entId, siteId, devId, imei, simCard).Set(1)
			} else {
				deviceConnectedStatus.WithLabelValues("Active", entId, siteId, devId, imei, simCard).Set(0)
			}
			time.Sleep(time.Second * 5)
		}
	}()

	go func() {
		for {
			randNum := rand.Intn(100)
			if randNum > 20 && randNum < 24 {
				mu.Lock()
				counter++
				deviceConnectionEventCore.WithLabelValues(fmt.Sprintf("%s: core event number-%d", devId, counter),
					entId, siteId, devId, time.Now().Format(time.RFC3339Nano)).
					Set(float64(rand.Intn(5) + 1))
				mu.Unlock()
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(60)+1))
		}
	}()

	go func() {
		for {
			randNum := rand.Intn(100)
			if randNum > 10 && randNum < 15 {
				mu.Lock()
				counter++
				deviceConnectionEventRan.WithLabelValues(fmt.Sprintf("%s: ran event number-%d", devId, counter),
					entId, siteId, devId, time.Now().Format(time.RFC3339Nano)).Set(float64(rand.Intn(5) + 1))
				mu.Unlock()
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(60)+1))
		}
	}()

	go func() {
		for {
			randNum := rand.Intn(100)
			if randNum > 60 && randNum < 65 {
				mu.Lock()
				counter++
				deviceConnectionEventFabric.WithLabelValues(fmt.Sprintf("%s: fabric event number-%d", devId, counter),
					entId, siteId, devId, time.Now().Format(time.RFC3339Nano)).Set(float64(rand.Intn(5) + 1))
				mu.Unlock()
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(60)+1))
		}
	}()
}

var (
	deviceConnectedStatus = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_device_connected_status",
		Help: "Device Status",
	}, []string{"device_status", "enterprise_id", "site_id", "device_id", "imei", "sim_card"})

	deviceConnectionEventCore = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_device_connection_event_core",
		Help: "Device Connection Event Core",
	}, []string{"msg", "enterprise_id", "site_id", "device_id", "time"})

	deviceConnectionEventRan = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_device_connection_event_ran",
		Help: "Device Connection Event Ran",
	}, []string{"msg", "site", "sim_card", "serial_number", "time"})

	deviceConnectionEventFabric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_device_connection_event_fabric",
		Help: "Device Connection Event Fabric",
	}, []string{"msg", "site", "sim_card", "serial_number", "time"})
)
