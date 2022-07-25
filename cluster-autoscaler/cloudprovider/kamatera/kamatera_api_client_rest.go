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
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/version"
)

const (
	userAgent = "kubernetes/cluster-autoscaler/" + version.ClusterAutoscalerVersion
)

// NewKamateraApiClientRest factory to create new Rest API Client struct
func NewKamateraApiClientRest(clientId string, secret string) (client KamateraApiClientRest) {
	return KamateraApiClientRest{
		userAgent: userAgent,
		clientId:  clientId,
		secret:    secret,
	}
}

// KamateraApiClientRest is the struct to perform API calls
type KamateraApiClientRest struct {
	userAgent string
	clientId  string
	secret    string
}

func (c *KamateraApiClientRest) ListServersByTag(ctx context.Context, tag string) ([]Server, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (c *KamateraApiClientRest) DeleteServer(ctx context.Context, id string) error {
	// TODO: call the instance delete method, wait for it to be deleted
	return cloudprovider.ErrNotImplemented
}

func (c *KamateraApiClientRest) CreateServers(ctx context.Context, count int) ([]Server, error) {
	// TODO: add server config param, create the servers, wait for them to be created, return Server structs
	return nil, cloudprovider.ErrNotImplemented
}
