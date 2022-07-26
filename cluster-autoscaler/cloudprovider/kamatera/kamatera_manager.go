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
	"io"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"

	klog "k8s.io/klog/v2"
)

const (
	clusterServerTagPrefix string = "k8sca-"
	nodeGroupTagPrefix string = "k8scang-"
)

// manager handles Kamatera communication and holds information about
// the node groups
type manager struct {
	client     kamateraAPIClient
	config     *kamateraConfig
	nodeGroups map[string]*NodeGroup // key: NodeGroup.id
}

func newManager(config io.Reader) (*manager, error) {
	cfg, err := buildCloudConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}
	client := buildKamateraAPIClient(cfg.apiClientId, cfg.apiSecret)
	m := &manager{
		client:     client,
		config:     cfg,
		nodeGroups: make(map[string]*NodeGroup),
	}
	return m, nil
}

func (m *manager) refresh() error {
	servers, error := m.client.ListServersByTag(context.Background(),
		fmt.Sprintf("%s%s", clusterServerTagPrefix, m.config.clusterName))
	if error != nil {
		return fmt.Errorf("failed to get list of Kamatera servers from Kamatera API: %v", error)
	}
	nodeGroups := make(map[string]*NodeGroup)
	for nodeGroupName, nodeGroupCfg := range m.config.nodeGroupCfg {
		nodeGroup, err := m.buildNodeGroup(nodeGroupName, nodeGroupCfg, servers)
		if err != nil {
			return fmt.Errorf("failed to build node group %s: %v", nodeGroupName, err)
		}
		nodeGroups[nodeGroupName] = nodeGroup
	}

	// show some debug info
	klog.V(2).Infof("Kamatera node groups after refresh:")
	for _, ng := range nodeGroups {
		klog.V(2).Infof("%s", ng.extendedDebug())
	}

	m.nodeGroups = nodeGroups
	return nil
}

func (m *manager) buildNodeGroup(name string, cfg *nodeGroupConfig, servers []Server) (*NodeGroup, error) {
	instances, err := m.getNodeGroupInstances(name, servers)
	if err != nil {
		return nil, fmt.Errorf("failed to get instances for node group %s: %v", name, err)
	}
	ng := &NodeGroup{
		id: name,
		manager: m,
		minSize: cfg.minSize,
		maxSize: cfg.maxSize,
		instances: instances,
	}
	return ng, nil
}

func (m *manager) getNodeGroupInstances(name string, servers []Server) (map[string]*Instance, error) {
	nodeGroupTag := fmt.Sprintf("%s%s", nodeGroupTagPrefix, name)
	instances := make(map[string]*Instance)
	for _, server := range servers {
		hasNodeGroupTag := false
		for _, tag := range server.Tags {
			if tag == nodeGroupTag {
				hasNodeGroupTag = true
				break
			}
		}
		if hasNodeGroupTag {
			instances[server.Name] = &Instance{
				Id:     server.Name,
				Status: &cloudprovider.InstanceStatus{State: cloudprovider.InstanceRunning},
			}
		}
	}
	return instances, nil
}
