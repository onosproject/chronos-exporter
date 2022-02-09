// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package aether

import "time"

func (s *EnterprisesEnterpriseSite) collect() {

	for _, sl := range *s.Slice {
		sl.collect()
	}

	for _, d := range *s.Device {
		d.collect(time.Second*10, s.SiteId)
	}
}
