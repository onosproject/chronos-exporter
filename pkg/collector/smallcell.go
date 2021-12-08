// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package collector

type SmallCell struct {
	SmallCellID string `json:"small-cell-id"`
	DisplayName string `json:"display-name,omitempty"`
}

func (sc *SmallCell) collect(site string) {
	// TODO add in any small cell metrics
}
