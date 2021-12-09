// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import (
	"encoding/json"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("collector")

type DeviceGroup struct {
	DeviceGroupID string   `json:"device-group-id"`
	DisplayName   string   `json:"display-name,omitempty"`
	Devices       []string `json:"devices,omitempty"` // This is a cross reference to Devices
}

type Sim struct {
	IccID       string `json:"iccid"`
	DisplayName string `json:"display-name,omitempty"`
}

type AetherModel struct {
	Sites []Site `json:"sites,omitempty"`
}

func LoadModel(modelData []byte) (*AetherModel, error) {
	if !json.Valid(modelData) {
		return nil, errors.NewInvalid("JSON is not valid: %s", string(modelData))
	}

	model := new(AetherModel)

	err := json.Unmarshal(modelData, model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (m *AetherModel) Collect() {
	for _, s := range m.Sites {
		log.Infof("Starting collector for Site %s", s.SiteID)
		s.collect()
	}
}
