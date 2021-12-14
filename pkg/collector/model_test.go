// SPDX-FileCopyrightText: 2021-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-Only-1.0

package collector

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_LoadModel(t *testing.T) {
	testConfigYAML, err := ioutil.ReadFile("testdata/test-config.yaml")
	assert.NoError(t, err, "error loading testdata file")

	model, err := LoadModel(testConfigYAML)
	assert.NoError(t, err)
	assert.NotNil(t, model)
	assert.Equal(t, 1, len(model.Sites))
	assert.Equal(t, 2, len(model.Sites[0].Slices))
	assert.Equal(t, "slice2", model.Sites[0].Slices[1].SliceId)
	assert.Equal(t, 2, len(model.Sites[0].DeviceGroups))
	devices := model.Sites[0].DeviceGroups[0].Devices
	assert.NotNil(t, devices)
	assert.Equal(t, 2, len(*devices))
	assert.Equal(t, "752365A", (*devices)[0])
}