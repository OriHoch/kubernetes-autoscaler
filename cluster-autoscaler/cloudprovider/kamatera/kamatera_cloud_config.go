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
	"strconv"

	"gopkg.in/gcfg.v1"
)

const (
	defaultMinSize int = 1
	defaultMaxSize int = 254
)

// nodeGroupConfig is the configuration for a specific node group.
type nodeGroupConfig struct {
	minSize int
	maxSize int
	// TODO: add server configs
}

// kamateraConfig holds the configuration for the Kamatera provider.
type kamateraConfig struct {
	apiClientId     string
	apiSecret       string
	clusterName     string
	defaultMinSize  int
	defaultMaxSize  int
	nodeGroupCfg    map[string]*nodeGroupConfig // key is the node group name
	// TODO: add default server configs
}

// GcfgGlobalConfig is the gcfg representation of the global section in the cloud config file for Kamatera.
type GcfgGlobalConfig struct {
	KamateraApiClientId string `gcfg:"kamatera-api-client-id"`
	KamateraApiSecret   string `gcfg:"kamatera-api-secret"`
	ClusterName 	    string `gcfg:"cluster-name"`
	DefaultMinSize      string   `gcfg:"defaut-min-size"`
	DefaultMaxSize      string   `gcfg:"defaut-max-size"`
	// TODO: add default server configs
}

// GcfgNodeGroupConfig is the gcfg representation of the section in the cloud config file to change defaults for a node group.
type GcfgNodeGroupConfig struct {
	MinSize string `gcfg:"min-size"`
	MaxSize string `gcfg:"max-size"`
	// TODO: add server configs
}

// gcfgCloudConfig is the gcfg representation of the cloud config file for Kamatera.
type gcfgCloudConfig struct {
	Global     GcfgGlobalConfig                `gcfg:"global"`
	NodeGroups map[string]*GcfgNodeGroupConfig `gcfg:"nodegroup"` // key is the node group name
}

// buildCloudConfig creates the configuration struct for the provider.
func buildCloudConfig(config io.Reader) (*kamateraConfig, error) {

	// read the config and get the gcfg struct
	var gcfgCloudConfig gcfgCloudConfig
	if err := gcfg.ReadInto(&gcfgCloudConfig, config); err != nil {
		return nil, err
	}

	// get the clusterName and Kamatera tokens
	clusterName := gcfgCloudConfig.Global.ClusterName
	if len(clusterName) == 0 {
		return nil, fmt.Errorf("cluster name is not set")
	}
	apiClientId := gcfgCloudConfig.Global.KamateraApiClientId
	if len(apiClientId) == 0 {
		return nil, fmt.Errorf("kamatera api client id is not set")
	}
	apiSecret := gcfgCloudConfig.Global.KamateraApiSecret
	if len(apiSecret) == 0 {
		return nil, fmt.Errorf("kamatera api secret is not set")
	}

	// get the default min and max size as defined in the global section of the config file
	defaultMinSize, defaultMaxSize, err := getSizeLimits(
		gcfgCloudConfig.Global.DefaultMinSize,
		gcfgCloudConfig.Global.DefaultMaxSize,
		defaultMinSize,
		defaultMaxSize)
	if err != nil {
		return nil, fmt.Errorf("cannot get default size values in global section: %v", err)
	}

	// get the specific configuration of a node group
	nodeGroupCfg := make(map[string]*nodeGroupConfig)
	for nodeGroupName, gcfgNodeGroup := range gcfgCloudConfig.NodeGroups {
		minSize, maxSize, err := getSizeLimits(gcfgNodeGroup.MinSize, gcfgNodeGroup.MaxSize, defaultMinSize, defaultMaxSize)
		if err != nil {
			return nil, fmt.Errorf("cannot get size values for node group %s: %v", nodeGroupName, err)
		}
		ngc := &nodeGroupConfig{
			maxSize: maxSize,
			minSize: minSize,
		}
		nodeGroupCfg[nodeGroupName] = ngc
	}

	return &kamateraConfig{
		clusterName:     clusterName,
		apiClientId:     apiClientId,
		apiSecret:       apiSecret,
		defaultMinSize:  defaultMinSize,
		defaultMaxSize:  defaultMaxSize,
		nodeGroupCfg:    nodeGroupCfg,
	}, nil
}

// getSizeLimits takes the max, min size of a node group as strings (empty if no values are provided)
// and default sizes, validates them and returns them as integer, or an error if such occurred
func getSizeLimits(minStr string, maxStr string, defaultMin int, defaultMax int) (int, int, error) {
	var err error
	min := defaultMin
	if len(minStr) != 0 {
		min, err = strconv.Atoi(minStr)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse min size for node group: %v", err)
		}
	}
	if min < 1 {
		return 0, 0, fmt.Errorf("min size for node group cannot be < 1")
	}
	max := defaultMax
	if len(maxStr) != 0 {
		max, err = strconv.Atoi(maxStr)
		if err != nil {
			return 0, 0, fmt.Errorf("could not parse max size for node group: %v", err)
		}
	}
	if min > max {
		return 0, 0, fmt.Errorf("min size for a node group must be less than its max size (got min: %d, max: %d)",
			min, max)
	}
	return min, max, nil
}
