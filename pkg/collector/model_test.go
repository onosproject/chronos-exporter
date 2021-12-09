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
	testConfigJSON, err := ioutil.ReadFile("testdata/test-config.json")
	assert.NoError(t, err, "error loading testdata file")

	model, err := LoadModel(testConfigJSON)
	assert.NoError(t, err)
	assert.NotNil(t, model)
	assert.Equal(t, 1, len(model.Sites))
	assert.Equal(t, 2, len(model.Sites[0].Slices))
	assert.Equal(t, "slice2", model.Sites[0].Slices[1].SliceID)
}