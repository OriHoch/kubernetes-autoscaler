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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloudConfig_getSizeLimits(t *testing.T) {
	_, _, err := getSizeLimits("3", "2", 1, 2)
	assert.Error(t, err, "no errors on minSize > maxSize")

	_, _, err = getSizeLimits("4", "", 2, 3)
	assert.Error(t, err, "no errors on minSize > maxSize using defaults")

	_, _, err = getSizeLimits("", "4", 5, 10)
	assert.Error(t, err, "no errors on minSize > maxSize using defaults")

	_, _, err = getSizeLimits("-1", "4", 5, 10)
	assert.Error(t, err, "no errors on minSize <= 0")

	_, _, err = getSizeLimits("1", "4a", 5, 10)
	assert.Error(t, err, "no error on malformed integer string")

	_, _, err = getSizeLimits("1.0", "4", 5, 10)
	assert.Error(t, err, "no error on malformed integer string")

	min, max, err := getSizeLimits("", "", 1, 2)
	assert.Equal(t, 1, min)
	assert.Equal(t, 2, max)

	min, max, err = getSizeLimits("", "3", 1, 2)
	assert.Equal(t, 1, min)
	assert.Equal(t, 3, max)

	min, max, err = getSizeLimits("6", "8", 1, 2)
	assert.Equal(t, 6, min)
	assert.Equal(t, 8, max)
}

func TestCloudConfig_buildCloudConfig(t *testing.T) {
	cfg := strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
default-min-size=1
default-max-size=10

[nodegroup "default"]

[nodegroup "highcpu"]
min-size=3

[nodegroup "highram"]
max-size=2
`)
	config, err := buildCloudConfig(cfg)
	assert.NoError(t, err)
	assert.Equal(t, "1a222bbb3ccc44d5555e6ff77g88hh9i", config.apiClientId)
	assert.Equal(t, "9ii88h7g6f55555ee4444444dd33eee2", config.apiSecret)
	assert.Equal(t, "aaabbb", config.clusterName)
	assert.Equal(t, 1, config.defaultMinSize)
	assert.Equal(t, 10, config.defaultMaxSize)
	assert.Equal(t, 3, len(config.nodeGroupCfg))
	assert.Equal(t, 1, config.nodeGroupCfg["default"].minSize)
	assert.Equal(t, 10, config.nodeGroupCfg["default"].maxSize)
	assert.Equal(t, 3, config.nodeGroupCfg["highcpu"].minSize)
	assert.Equal(t, 10, config.nodeGroupCfg["highcpu"].maxSize)
	assert.Equal(t, 1, config.nodeGroupCfg["highram"].minSize)
	assert.Equal(t, 2, config.nodeGroupCfg["highram"].maxSize)

	cfg = strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
default-min-size=1
default-max-size=10

[nodegroup "default"]

[nodegroup "highcpu"]
min-size=3

[nodegroup "highram"]
max-size=2a
`)
	config, err = buildCloudConfig(cfg)
	assert.Error(t, err, "no error on size of a specific node group is not an integer string")

	cfg = strings.NewReader(`
[global]
cluster-name=aaabbb
default-min-size=1
default-max-size=10
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i

[nodegroup "default"]

[nodegroup "highcpu"]
min-size=3

[nodegroup "highram"]
max-size=2
`)
	config, err = buildCloudConfig(cfg)
	assert.Error(t, err, "no error on missing kamatera api secret")
	assert.Contains(t, err.Error(), "kamatera api secret is not set")

	cfg = strings.NewReader(`
[global]
cluster-name=aaabbb
default-min-size=1
default-max-size=10

[nodegroup "default"]

[nodegroup "highcpu"]
min-size=3

[nodegroup "highram"]
max-size=2
`)
	config, err = buildCloudConfig(cfg)
	assert.Error(t, err, "no error on missing kamatera api client id")
	assert.Contains(t, err.Error(), "kamatera api client id is not set")

	cfg = strings.NewReader(`
[gglobal]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
default-min-size=1
default-max-size=10

[nodegroup "default"]

[nodegroup "highcpu"]
min-size=3

[nodegroup "highram"]
max-size=2
`)
	config, err = buildCloudConfig(cfg)
	assert.Error(t, err, "no error when config has no global section")
	assert.Contains(t, err.Error(), "can't store data at section \"gglobal\"")

	cfg = strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=aaabbb
default-min-size=1
default-max-size=10

[nodegroup "1234567890123456"]
`)
	config, err = buildCloudConfig(cfg)
	assert.Error(t, err, "no error when nodegroup name is more then 15 characters")
	assert.Contains(t, err.Error(), "node group name must be at most 15 characters long")

	cfg = strings.NewReader(`
[global]
kamatera-api-client-id=1a222bbb3ccc44d5555e6ff77g88hh9i
kamatera-api-secret=9ii88h7g6f55555ee4444444dd33eee2
cluster-name=1234567890123456
default-min-size=1
default-max-size=10

[nodegroup "default"]
`)
	config, err = buildCloudConfig(cfg)
	assert.Error(t, err, "no error when cluster name is more then 15 characters")
	assert.Contains(t, err.Error(), "cluster name must be at most 15 characters long")
}
