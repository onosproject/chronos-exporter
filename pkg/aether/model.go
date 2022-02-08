// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package aether

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"gopkg.in/yaml.v2"
)

var log = logging.GetLogger("aether")

func LoadModel(modelData []byte) (*Enterprises, error) {

	model := new(Enterprises)

	err := yaml.UnmarshalStrict(modelData, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (a *Enterprises) Collect() {
	for _, e := range *a.Enterprise {
		log.Infof("Starting collector for Enterprise %s", e.EnterpriseId)
		e.collect()
	}
}
