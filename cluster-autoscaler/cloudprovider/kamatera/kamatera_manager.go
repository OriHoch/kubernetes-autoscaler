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
	"fmt"
	"io"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"

	klog "k8s.io/klog/v2"
)

// manager handles Kamatera communication and holds information about
// the node groups
type manager struct {
	config     *kamateraConfig
	nodeGroups map[string]*NodeGroup // key: NodeGroup.id
}

func newManager(config io.Reader) (*manager, error) {
	cfg, err := buildCloudConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %v", err)
	}
	m := &manager{
		config:     cfg,
		nodeGroups: make(map[string]*NodeGroup),
	}
	return m, nil
}

func (m *manager) refresh() error {
	nodeGroups := make(map[string]*NodeGroup)
	for nodeGroupName, nodeGroupCfg := range m.config.nodeGroupCfg {
		nodeGroup, err := m.buildNodeGroup(nodeGroupName, nodeGroupCfg)
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

func (m *manager) buildNodeGroup(name string, cfg *nodeGroupConfig) (*NodeGroup, error) {
	instances, err := m.getNodeGroupInstances(name)
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

func (m *manager) getNodeGroupInstances(name string) (map[string]*Instance, error) {
	return nil, cloudprovider.ErrNotImplemented
}
