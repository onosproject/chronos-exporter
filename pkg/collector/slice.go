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

type Slice struct {
	SliceID      string   `json:"slice-id"`
	DisplayName  string   `json:"display-name,omitempty"`
	DeviceGroups []string `json:"device-groups,omitempty"` // This is a cross reference to DeviceGroup
}

func (s *Slice) collect(period time.Duration, site string) {

	// TODO: add in slice metrics generation
	go func() {
		minThroughput := 100000
		maxThroughput := 1000000
		for {
			throughput := float64(rand.Intn(maxThroughput - minThroughput) + minThroughput)
			sliceThroughputBytes.WithLabelValues(s.DisplayName, site).Set(throughput)
			time.Sleep(period/10)
		}
	}()

	go func() {
		minPacket := 100
		maxPacket := 800
		for {
			packets := float64(rand.Intn(maxPacket - minPacket) + minPacket)
			sliceThroughputPackets.WithLabelValues(s.DisplayName, site).Set(packets)
			time.Sleep(period/10)
		}
	}()

	go func() {
		minLatency := 10
		maxLatency := 25
		for {
			latency := float64(rand.Intn(maxLatency - minLatency) + minLatency)
			if latency != 17 {
				sliceLatencyEndToEnd.WithLabelValues(s.DisplayName, site).Set(latency)
				time.Sleep(period/10)
			} else {
				sliceLatencyEndToEnd.WithLabelValues(s.DisplayName, site).Set(100)
				randomNum := rand.Intn(10)
				time.Sleep(time.Duration(randomNum) * period)
			}
		}
	}()
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
)