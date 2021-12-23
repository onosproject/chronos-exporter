// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import (
	"time"
)

func (s *Site) collect() {
	// Start a go routine for any Site only metrics - if any

	// getting the list of devices in each device group
	var devices = make(map[string][]string)
	for _, dg := range s.DeviceGroups {
		devices[dg.DeviceGroupId] = *dg.Devices
	}

	// Then start the Slice's collector
	for _, slice := range s.Slices {
		slice.collect(time.Second * 10, s.SiteId)

		if slice.Applications == nil || slice.DeviceGroups == nil {
			return
		}
		for _, app := range *slice.Applications {
			appId := app
			for _, dg := range *slice.DeviceGroups {
				slice.collectPerDevicePerApplication(time.Second * 10, s.SiteId, appId, devices[dg])
			}
		}
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