/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kamatera

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_newManager(t *testing.T) {
	cfg := strings.NewReader(`
[globalxxx]
`)
	_, err := newManager(cfg)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "can't store data at section \"globalxxx\"")

	cfg = strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
`)
	_, err = newManager(cfg)
	assert.NoError(t, err)
}


func TestManager_refresh(t *testing.T) {
	cfg := strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb

[nodegroup "ng1"]
min-size=1
max-size=2

[nodegroup "ng2"]
min-size=4
max-size=5
`)
	m, err := newManager(cfg)
	assert.NoError(t, err)

	client := kamateraClientMock{}
	m.client = &client
	ctx := context.Background()

	serverName1 := mockKamateraServerName()
	serverName2 := mockKamateraServerName()
	serverName3 := mockKamateraServerName()
	serverName4 := mockKamateraServerName()
	client.On(
		"ListServersByTag", ctx, fmt.Sprintf("%s%s", clusterServerTagPrefix, "aaabbb"),
	).Return(
		[]Server{
			{Name: serverName1, Tags: []string{fmt.Sprintf("%s%s", clusterServerTagPrefix, "aaabbb"), fmt.Sprintf("%s%s", nodeGroupTagPrefix, "ng1")}},
			{Name: serverName2, Tags: []string{fmt.Sprintf("%s%s", nodeGroupTagPrefix, "ng1"), fmt.Sprintf("%s%s", clusterServerTagPrefix, "aaabbb")}},
			{Name: serverName3, Tags: []string{fmt.Sprintf("%s%s", nodeGroupTagPrefix, "ng1"), fmt.Sprintf("%s%s", clusterServerTagPrefix, "aaabbb")}},
			{Name: serverName4, Tags: []string{fmt.Sprintf("%s%s", nodeGroupTagPrefix, "ng2"), fmt.Sprintf("%s%s", clusterServerTagPrefix, "aaabbb")}},
		},
		nil,
	).Once()
	err = m.refresh()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(m.nodeGroups))
	assert.Equal(t, 3, len(m.nodeGroups["ng1"].instances))
	assert.Equal(t, 1, len(m.nodeGroups["ng2"].instances))

	// test api error
	client.On(
		"ListServersByTag", ctx, fmt.Sprintf("%s%s", clusterServerTagPrefix, "aaabbb"),
	).Return(
		[]Server{},
		fmt.Errorf("error on API call"),
	).Once()
	err = m.refresh()
	assert.Error(t, err)
	assert.Equal(t, "failed to get list of Kamatera servers from Kamatera API: error on API call", err.Error())
}
