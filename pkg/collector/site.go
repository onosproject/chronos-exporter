// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import "time"

type Site struct {
	SiteID       string        `json:"site-id"`
	DisplayName  string        `json:"display-name,omitempty"`
	Image        string        `json:"image,omitempty"`
	SmallCells   []SmallCell   `json:"small-cells,omitempty"`
	Slices       []Slice       `json:"slices,omitempty"`
	DeviceGroups []DeviceGroup `json:"device-groups,omitempty"`
	Devices      []Device      `json:"devices,omitempty"`
	Sims         []Sim         `json:"sims,omitempty"`
}

func (s *Site) collect() {
	// Start a go routine for any Site only metrics - if any

	// Then start the Slice's collector
	for _, slice := range s.Slices {
		slice.collect(time.Second * 10, s.SiteID)
	}

	// and each Devices's collector
	for _, device := range s.Devices {
		device.collect(time.Second * 10, s.SiteID)
	}

	// and each Small cell's collector
	for _, sc := range s.SmallCells {
		sc.collect(s.SiteID)
	}
}