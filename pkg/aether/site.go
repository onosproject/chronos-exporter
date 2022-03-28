// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package aether

func (s *EnterprisesEnterpriseSite) collect(entId string) {

	for _, sl := range *s.Slice {
		sl.collect(entId, s.SiteId)
	}

	for _, d := range *s.Device {
		d.collect(entId, s.SiteId)
	}
}
