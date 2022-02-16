// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package aether

func (e *EnterprisesEnterprise) collect() {

	for _, site := range *e.Site {
		site.collect(e.EnterpriseId)
	}

}
