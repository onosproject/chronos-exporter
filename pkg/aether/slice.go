// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package aether

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"math/rand"
	"time"
)

func (s *EnterprisesEnterpriseSiteSlice) collect(entId string, siteId string) {
	log.Infof("Starting collector for Slice %s", s.SliceId)
	sliceId := s.SliceId

	go func() {
		minThroughput := 100000
		maxThroughput := 1000000
		for {
			throughput := float64(rand.Intn(maxThroughput-minThroughput) + minThroughput)
			sliceThroughputBytes.WithLabelValues(entId, siteId, sliceId).Set(throughput)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		minPacket := 100
		maxPacket := 800
		for {
			packets := float64(rand.Intn(maxPacket-minPacket) + minPacket)
			sliceThroughputPackets.WithLabelValues(entId, siteId, sliceId).Set(packets)
			time.Sleep(time.Second)
		}
	}()

	go func() {
		minLatency := 15
		maxLatency := 25
		for {
			latency := float64(rand.Intn(maxLatency-minLatency) + minLatency)
			randomNumber := rand.Intn(100)
			if randomNumber != 17 {
				sliceLatencyEndToEnd.WithLabelValues(entId, siteId, sliceId).Set(latency)
				time.Sleep(time.Second)
			} else {
				sliceLatencyEndToEnd.WithLabelValues(entId, siteId, sliceId).Set(100)
				time.Sleep(time.Second)
			}
		}
	}()

}

var (
	sliceThroughputBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_slice_throughput_bytes",
		Help: "Slice Throughput",
	}, []string{"enterprise_id", "site_id", "slice_id"})

	sliceThroughputPackets = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_slice_throughput_packets",
		Help: "Slice Throughput Packets",
	}, []string{"enterprise_id", "site_id", "slice_id"})

	sliceLatencyEndToEnd = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "aether_slice_latency_end_to_end_milliseconds",
		Help: "Slice Latency End-to-End",
	}, []string{"enterprise_id", "site_id", "slice_id"})
)
