// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package collector

import (
	"github.com/onosproject/onos-lib-go/pkg/logging"
	"gopkg.in/yaml.v2"
)

var log = logging.GetLogger("collector")

func LoadModel(modelData []byte) (*AetherModel, error) {

	model := new(AetherModel)

	err := yaml.UnmarshalStrict(modelData, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *AetherModel) Collect() {
	for _, s := range m.Sites {
		log.Infof("Starting collector for Site %s", s.SiteId)
		s.collect()
	}
}
