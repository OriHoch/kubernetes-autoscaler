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
	"encoding/base64"
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
	// TODO: do validation of server args with Kamatera api
	instances, err := m.getNodeGroupInstances(name, servers)
	if err != nil {
		return nil, fmt.Errorf("failed to get instances for node group %s: %v", name, err)
	}
	scriptBytes, err := base64.StdEncoding.DecodeString(cfg.ScriptBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode script for node group %s: %v", name, err)
	}
	script := string(scriptBytes)
	if len(script) < 1 {
		return nil, fmt.Errorf("script for node group %s is empty", name)
	}
	if len(cfg.Datacenter) < 1 {
		return nil, fmt.Errorf("datacenter for node group %s is empty", name)
	}
	if len(cfg.Image) < 1 {
		return nil, fmt.Errorf("image for node group %s is empty", name)
	}
	if len(cfg.Cpu) < 1 {
		return nil, fmt.Errorf("cpu for node group %s is empty", name)
	}
	if len(cfg.Ram) < 1 {
		return nil, fmt.Errorf("ram for node group %s is empty", name)
	}
	if len(cfg.Disks) < 1 {
		return nil, fmt.Errorf("no disks for node group %s", name)
	}
	if len(cfg.Networks) < 1 {
		return nil, fmt.Errorf("no networks for node group %s", name)
	}
	billingCycle := cfg.BillingCycle
	if billingCycle == "" {
		billingCycle = "hourly"
	} else if billingCycle != "hourly" && billingCycle != "monthly" {
		return nil, fmt.Errorf("billing cycle for node group %s is invalid", name)
	}
	serverConfig := ServerConfig{
		NamePrefix:     cfg.NamePrefix,
		Password:       cfg.Password,
		SshKey:         cfg.SshKey,
		Datacenter:     cfg.Datacenter,
		Image:          cfg.Image,
		Cpu:            cfg.Cpu,
		Ram:            cfg.Ram,
		Disks:          cfg.Disks,
		Dailybackup:    cfg.Dailybackup,
		Managed:        cfg.Managed,
		Networks:       cfg.Networks,
		BillingCycle:   billingCycle,
		MonthlyPackage: cfg.MonthlyPackage,
		ScriptFile: 	script,
		UserdataFile:   "",
		Tags:           []string{
			fmt.Sprintf("%s%s", clusterServerTagPrefix, m.config.clusterName),
			fmt.Sprintf("%s%s", nodeGroupTagPrefix, name),
		},
	}
	ng := &NodeGroup{
		id: name,
		manager: m,
		minSize: cfg.minSize,
		maxSize: cfg.maxSize,
		instances: instances,
		serverConfig: serverConfig,
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
