// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import (
	"time"
)

func (s *Site) collect() {
	// Start a go routine for any Site only metrics - if any

	// Then start the Slice's collector
	for _, slice := range s.Slices {
		slice.collect(time.Second * 10, s.SiteId)
	}

	// and each Devices's collector
	for _, device := range s.Devices {
		device.collect(time.Second * 10, s.SiteId)
	}

	// and each Small cell's collector
	for _, sc := range s.SmallCells {
		sc.collect(s.SiteId)
	}
}