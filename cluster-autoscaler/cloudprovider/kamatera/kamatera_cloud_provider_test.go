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
	apiv1 "k8s.io/api/core/v1"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
)

func TestCloudProvider_newKamateraCloudProvider(t *testing.T) {
	// test ok on correctly creating a Kamatera provider
	rl := &cloudprovider.ResourceLimiter{}
	cfg := strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
`)
	_, err := newKamateraCloudProvider(cfg, rl)
	assert.NoError(t, err)

	// test error on creating a Kamatera provider when config is bad
	cfg = strings.NewReader(`
[global]
kamatera-api-client-id=
kamatera-api-secret=
cluster-name=
`)
	_, err = newKamateraCloudProvider(cfg, rl)
	assert.Error(t, err)
	assert.Equal(t, "could not create kamatera manager: failed to parse config: cluster name is not set", err.Error())
}


func TestCloudProvider_NodeGroups(t *testing.T) {
	// test ok on getting the correct nodes when calling NodeGroups()
	cfg := strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
`)
	m, _ := newManager(cfg)
	m.nodeGroups = map[string]*NodeGroup{
		"ng1": {id: "ng1"},
		"ng2": {id: "ng2"},
	}
	kcp := &kamateraCloudProvider{manager: m}
	ng := kcp.NodeGroups()
	assert.Equal(t, 2, len(ng))
	assert.Contains(t, ng, m.nodeGroups["ng1"])
	assert.Contains(t, ng, m.nodeGroups["ng2"])
}

func TestCloudProvider_NodeGroupForNode(t *testing.T) {
	cfg := strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
`)
	m, _ := newManager(cfg)
	kamateraServerId1 := mockKamateraServerId()
	kamateraServerId2 := mockKamateraServerId()
	kamateraServerId3 := mockKamateraServerId()
	kamateraServerId4 := mockKamateraServerId()
	ng1 := &NodeGroup{
		id: "ng1",
		instances: map[string]*Instance{
			kamateraServerId1: {Id: kamateraServerId1},
			kamateraServerId2: {Id: kamateraServerId2},
		},
	}
	ng2 := &NodeGroup{
		id: "ng2",
		instances: map[string]*Instance{
			kamateraServerId3: {Id: kamateraServerId3},
			kamateraServerId4: {Id: kamateraServerId4},
		},
	}
	m.nodeGroups = map[string]*NodeGroup{
		"ng1": ng1,
		"ng2": ng2,
	}
	kcp := &kamateraCloudProvider{manager: m}

	// test ok on getting the right node group for an apiv1.Node
	node := &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: kamateraServerId1,
		},
	}
	ng, err := kcp.NodeGroupForNode(node)
	assert.NoError(t, err)
	assert.Equal(t, ng1, ng)

	// test ok on getting the right node group for an apiv1.Node
	node = &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: kamateraServerId4,
		},
	}
	ng, err = kcp.NodeGroupForNode(node)
	assert.NoError(t, err)
	assert.Equal(t, ng2, ng)

	// test ok on getting nil when looking for a apiv1.Node we do not manage
	node = &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: mockKamateraServerId(),
		},
	}
	ng, err = kcp.NodeGroupForNode(node)
	assert.NoError(t, err)
	assert.Nil(t, ng)

	// test error on looking for a apiv1.Node with a bad providerID
	node = &apiv1.Node{
		Spec: apiv1.NodeSpec{
			ProviderID: "",
		},
	}
	ng, err = kcp.NodeGroupForNode(node)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Invalid node ProviderID")
}


func TestCloudProvider_others(t *testing.T) {
	cfg := strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
`)
	m, _ := newManager(cfg)
	resourceLimiter := &cloudprovider.ResourceLimiter{}
	kcp := &kamateraCloudProvider{
		manager:         m,
		resourceLimiter: resourceLimiter,
	}
	assert.Equal(t, cloudprovider.KamateraProviderName, kcp.Name())
	_, err := kcp.GetAvailableMachineTypes()
	assert.Error(t, err)
	_, err = kcp.NewNodeGroup("", nil, nil, nil, nil)
	assert.Error(t, err)
	rl, err := kcp.GetResourceLimiter()
	assert.Equal(t, resourceLimiter, rl)
	assert.Equal(t, "", kcp.GPULabel())
	assert.Nil(t, kcp.GetAvailableGPUTypes())
	assert.Nil(t, kcp.Cleanup())
	_, err2 := kcp.Pricing()
	assert.Error(t, err2)
}
