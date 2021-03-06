// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"time"
)

func (s *Slice) collect(period time.Duration, site string) {
	log.Infof("Starting collector for Slice %s", s.SliceId)
	sliceId := s.SliceId
	go func() {
		minThroughput := 100000
		maxThroughput := 1000000
		for {
			throughput := float64(rand.Intn(maxThroughput-minThroughput) + minThroughput)
			sliceThroughputBytes.WithLabelValues(sliceId, site).Set(throughput)
			time.Sleep(period / 10)
		}
	}()

	go func() {
		minPacket := 100
		maxPacket := 800
		for {
			packets := float64(rand.Intn(maxPacket-minPacket) + minPacket)
			sliceThroughputPackets.WithLabelValues(sliceId, site).Set(packets)
			time.Sleep(period / 10)
		}
	}()

	go func() {
		minLatency := 15
		maxLatency := 25
		for {
			latency := float64(rand.Intn(maxLatency-minLatency) + minLatency)
			randomNumber := rand.Intn(100)
			if randomNumber != 17 {
				sliceLatencyEndToEnd.WithLabelValues(sliceId, site).Set(latency)
				time.Sleep(period / 10)
			} else {
				sliceLatencyEndToEnd.WithLabelValues(sliceId, site).Set(100)
				time.Sleep(period / 10)
			}
		}
	}()
}

func (s *Slice) collectPerDevicePerApplication(period time.Duration, site string, app string, devices []string) {
	slice := s.SliceId
	for _, device := range devices {
		serialNumber := device
		go func() {
			minThroughput := 100000
			maxThroughput := 1000000
			for {
				throughput := float64(rand.Intn(maxThroughput-minThroughput) + minThroughput)
				perDevicePerApplicationThroughputBytes.WithLabelValues(site, slice, app, serialNumber).Set(throughput)
				time.Sleep(period / 10)
			}
		}()

		go func() {
			minPacket := 100
			maxPacket := 800
			for {
				throughput := float64(rand.Intn(maxPacket-minPacket) + minPacket)
				perDevicePerApplicationThroughputPacket.WithLabelValues(site, slice, app, serialNumber).Set(throughput)
				time.Sleep(period / 10)
			}
		}()

		go func() {
			minPacket := 1
			maxPacket := 8
			for {
				throughput := float64(rand.Intn(maxPacket-minPacket) + minPacket)
				perDevicePerApplicationLossPacket.WithLabelValues(site, slice, app, serialNumber).Set(throughput)
				time.Sleep(period / 10)
			}
		}()
	}
}

var (
	sliceThroughputBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "slice_throughput_bytes",
		Help: "Slice Throughput",
	}, []string{"slice", "site"})

	sliceThroughputPackets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "slice_throughput_packets",
		Help: "Slice Throughput Packets",
	}, []string{"slice", "site"})

	sliceLatencyEndToEnd = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "slice_latency_end_to_end_milliseconds",
		Help: "Slice Latency End-to-End",
	}, []string{"slice", "site"})

	perDevicePerApplicationThroughputBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "per_device_per_application_throughput_bytes",
		Help: "Per Device Per Application Throughput Bytes",
	}, []string{"site", "slice", "application", "serial_number"})

	perDevicePerApplicationThroughputPacket = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "per_device_per_application_throughput_packets",
		Help: "Per Device Per Application Throughput Packets",
	}, []string{"site", "slice", "application", "serial_number"})

	perDevicePerApplicationLossPacket = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "per_device_per_application_loss_packets",
		Help: "Per Device Per Application Loss Packet",
	}, []string{"site", "slice", "application", "serial_number"})
)
