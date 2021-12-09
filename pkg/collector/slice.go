// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import "time"

type Slice struct {
	SliceID      string   `json:"slice-id"`
	DisplayName  string   `json:"display-name,omitempty"`
	DeviceGroups []string `json:"device-groups,omitempty"` // This is a cross reference to DeviceGroup
}

func (s *Slice) collect(period time.Duration, site string) {

	// TODO: add in slice metrics generation
}
