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
	apiv1 "k8s.io/api/core/v1"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
)

func TestNodeGroup_IncreaseSize(t *testing.T) {
	client := kamateraClientMock{}
	ctx := context.Background()
	mgr := manager{
		client:     &client,
	}
	serverName1 := mockKamateraServerName()
	serverName2 := mockKamateraServerName()
	serverName3 := mockKamateraServerName()
	serverConfig := mockServerConfig("test", []string{})
	ng := NodeGroup{
		id:        "ng1",
		manager:   &mgr,
		minSize:   1,
		maxSize:   7,
		instances: map[string]*Instance{
			serverName1: {Id: serverName1, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName2: {Id: serverName2, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName3: {Id: serverName3, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
		},
		serverConfig: serverConfig,
	}

	// test error on bad delta values
	err := ng.IncreaseSize(0)
	assert.Error(t, err)
	assert.Equal(t, "delta must be positive, have: 0", err.Error())

	err = ng.IncreaseSize(-1)
	assert.Error(t, err)
	assert.Equal(t, "delta must be positive, have: -1", err.Error())

	// test error on a too large increase of nodes
	err = ng.IncreaseSize(5)
	assert.Error(t, err)
	assert.Equal(t, "size increase is too large. current: 3 desired: 8 max: 7", err.Error())

	// test ok to add a node
	client.On(
		"CreateServers", ctx, 1, serverConfig,
	).Return(
		[]Server{{Name: mockKamateraServerName()}} , nil,
	).Once()
	err = ng.IncreaseSize(1)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(ng.instances))


	// test ok to add multiple nodes
	client.On(
		"CreateServers", ctx, 2, serverConfig,
	).Return(
		[]Server{
			{Name: mockKamateraServerName()},
			{Name: mockKamateraServerName()},
		} , nil,
	).Once()
	err = ng.IncreaseSize(2)
	assert.NoError(t, err)
	assert.Equal(t, 6, len(ng.instances))

	// test error on linode API call error
	client.On(
		"CreateServers", ctx, 1, serverConfig,
	).Return(
		[]Server{}, fmt.Errorf("error on API call"),
	).Once()
	err = ng.IncreaseSize(1)
	assert.Error(t, err, "no error on injected API call error")
	assert.Equal(t, "error on API call", err.Error())
}


func TestNodeGroup_DecreaseTargetSize(t *testing.T) {
	ng := &NodeGroup{}
	err := ng.DecreaseTargetSize(-1)
	assert.Error(t, err)
	assert.Equal(t, "Not implemented", err.Error())
}

func TestNodeGroup_DeleteNodes(t *testing.T) {
	client := kamateraClientMock{}
	ctx := context.Background()
	mgr := manager{
		client:     &client,
	}
	serverName1 := mockKamateraServerName()
	serverName2 := mockKamateraServerName()
	serverName3 := mockKamateraServerName()
	serverName4 := mockKamateraServerName()
	serverName5 := mockKamateraServerName()
	serverName6 := mockKamateraServerName()
	ng := NodeGroup{
		id: "ng1",
		minSize: 1,
		maxSize: 6,
		instances: map[string]*Instance{
			serverName1: {Id: serverName1, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName2: {Id: serverName2, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName3: {Id: serverName3, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName4: {Id: serverName4, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName5: {Id: serverName5, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName6: {Id: serverName6, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
		},
		manager: &mgr,
	}

	// test of deleting nodes
	client.On(
		"DeleteServer", ctx, serverName1,
	).Return(nil).Once().On(
		"DeleteServer", ctx, serverName2,
	).Return(nil).Once().On(
		"DeleteServer", ctx, serverName6,
	).Return(nil).Once()
	err := ng.DeleteNodes([]*apiv1.Node{
		{Spec: apiv1.NodeSpec{ProviderID: serverName1}},
		{Spec: apiv1.NodeSpec{ProviderID: serverName2}},
		{Spec: apiv1.NodeSpec{ProviderID: serverName6}},
	})
	assert.NoError(t, err)
	assert.Equal(t, 3, len(ng.instances))
	assert.Equal(t, serverName3, ng.instances[serverName3].Id)
	assert.Equal(t, serverName4, ng.instances[serverName4].Id)
	assert.Equal(t, serverName5, ng.instances[serverName5].Id)

		// test error on deleting a node with invalid providerID
	err = ng.DeleteNodes([]*apiv1.Node{{Spec: apiv1.NodeSpec{ProviderID: ""}}})
	assert.Error(t, err)
	assert.Equal(t, "Invalid node ProviderID: ", err.Error())


	// test error on deleting a node we are not managing
	err = ng.DeleteNodes([]*apiv1.Node{{Spec: apiv1.NodeSpec{ProviderID: mockKamateraServerName()}}})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "cannot find this node in the node group")

	// test error on deleting a node when the linode API call fails
	client.On(
		"DeleteServer", ctx, serverName4,
	).Return(fmt.Errorf("error on API call")).Once()
	err = ng.DeleteNodes([]*apiv1.Node{
		{Spec: apiv1.NodeSpec{ProviderID: serverName4}},
	})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "error on API call")
}

func TestNodeGroup_Nodes(t *testing.T) {
	client := kamateraClientMock{}
	mgr := manager{
		client:     &client,
	}
	serverName1 := mockKamateraServerName()
	serverName2 := mockKamateraServerName()
	serverName3 := mockKamateraServerName()
	ng := NodeGroup{
		id: "ng1",
		minSize: 1,
		maxSize: 6,
		instances: map[string]*Instance{
			serverName1: {Id: serverName1, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName2: {Id: serverName2, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName3: {Id: serverName3, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
		},
		manager: &mgr,
	}

	// test nodes returned from Nodes() are only the ones we are expecting
	instancesList, err := ng.Nodes()
	assert.NoError(t, err)
	assert.Equal(t, 3, len(instancesList))
	var serverIds []string
	for _, instance := range instancesList {
		serverIds = append(serverIds, instance.Id)
	}
	assert.Equal(t, 3, len(serverIds))
	assert.Contains(t, serverIds, serverName1)
	assert.Contains(t, serverIds, serverName2)
	assert.Contains(t, serverIds, serverName3)
}

func TestNodeGroup_Others(t *testing.T) {
	client := kamateraClientMock{}
	mgr := manager{
		client:     &client,
	}
	serverName1 := mockKamateraServerName()
	serverName2 := mockKamateraServerName()
	serverName3 := mockKamateraServerName()
	ng := NodeGroup{
		id: "ng1",
		minSize: 1,
		maxSize: 7,
		instances: map[string]*Instance{
			serverName1: {Id: serverName1, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName2: {Id: serverName2, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
			serverName3: {Id: serverName3, Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning}},
		},
		manager: &mgr,
	}
	assert.Equal(t, 1, ng.MinSize())
	assert.Equal(t, 7, ng.MaxSize())
	ts, err := ng.TargetSize()
	assert.NoError(t, err)
	assert.Equal(t, 3, ts)
	assert.Equal(t, "ng1", ng.Id())
	assert.Equal(t, "node group ID: ng1 (min:1 max:7)", ng.Debug())
	assert.Equal(t, "node group ID: ng1 (min:1 max:7)", ng.extendedDebug())
	assert.Equal(t, true, ng.Exist())
	assert.Equal(t, false, ng.Autoprovisioned())
	_, err = ng.TemplateNodeInfo()
	assert.Error(t, err)
	assert.Equal(t, "Not implemented", err.Error())
	_, err = ng.Create()
	assert.Error(t, err)
	assert.Equal(t, "Not implemented", err.Error())
	err = ng.Delete()
	assert.Error(t, err)
	assert.Equal(t, "Not implemented", err.Error())
}
