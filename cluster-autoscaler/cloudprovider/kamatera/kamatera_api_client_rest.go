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
	"encoding/hex"
	"fmt"
	"github.com/satori/go.uuid"
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
		baseURL:   "https://cloudcli.cloudwm.com",
	}
}

type KamateraServerPostRequest struct {
	ServerName string `json:"name"`
}

type KamateraServerTerminatePostRequest struct {
	ServerName string `json:"name"`
	Force bool      `json:"force"`
}

type KamateraServerCreatePostRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
	PasswordValidate string `json:"passwordValidate"`
	SshKey string `json:"ssh-key"`
	Datacenter string `json:"datacenter"`
	Image string `json:"image"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	Disk string `json:"disk"`
	Dailybackup string `json:"dailybackup"`
	Managed string `json:"managed"`
	Network string `json:"network"`
	Quantity int `json:"quantity"`
	BillingCycle string `json:"billingCycle"`
	MonthlyPackage string `json:"monthlypackage"`
	Poweronaftercreate string `json:"poweronaftercreate"`
	ScriptFile string `json:"script-file"`
	UserdataFile string `json:"userdata-file"`
	Tag string `json:"tag"`
}

// KamateraApiClientRest is the struct to perform API calls
type KamateraApiClientRest struct {
	userAgent string
	clientId  string
	secret    string
	baseURL   string
}

func (c *KamateraApiClientRest) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *KamateraApiClientRest) ListServersByTag(ctx context.Context, tag string) ([]Server, error) {
	res, err := request(
		ctx,
		ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
		"GET",
		"/service/servers",
		nil,
	)
	if err != nil {
		return nil, err
	}
	var servers []Server
	for _, server := range res.([]interface{}) {
		server := server.(map[string]interface{})
		serverName := server["name"].(string)
		serverTags, err := c.getServerTags(ctx, serverName)
		if err != nil {
			return nil, err
		}
		hasTag := false
		for _, t := range serverTags {
			if t == tag {
				hasTag = true
				break
			}
		}
		if hasTag {
			servers = append(servers, Server{
				Name: serverName,
				Tags: serverTags,
			})
		}
	}
	return servers, nil
}

func (c *KamateraApiClientRest) DeleteServer(ctx context.Context, name string) error {
	res, err := request(
		ctx,
		ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
		"POST",
		"/service/server/poweroff",
		KamateraServerPostRequest{ServerName: name},
	)
	if err != nil {
		return err
	}
	commandId := res.([]interface{})[0].(string)
	_, err = waitCommand(
		ctx,
		ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
		commandId,
	)
	if err != nil {
		return err
	}
	_, err = request(
		ctx,
		ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
		"POST",
		"/service/server/delete",
		KamateraServerTerminatePostRequest{ServerName: name, Force: true},
	)
	if err != nil {
		return err
	}
	return nil
}

func (c *KamateraApiClientRest) CreateServers(ctx context.Context, count int) ([]Server, error) {
	// TODO: add configurable server params
	// TODO: add configurable server name prefix
	// TODO: add server tags based on node tags
	baseName := fmt.Sprintf("test-%s", hex.EncodeToString(uuid.NewV4().Bytes()))
	serverNameCommandIds := make(map[string]string)
	for i := 0; i < count; i++ {
		serverName := fmt.Sprintf("%s-%d", baseName, i)
		res, err := request(
			ctx,
			ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
			"POST",
			"/service/server",
			KamateraServerCreatePostRequest{
				Name:               serverName,
				Password:           "",
				PasswordValidate:   "",
				SshKey:             "ssh-rsa AAAA== root@localhost",
				Datacenter:         "IL",
				Image:              "ubuntu_server_18.04_64-bit",
				Cpu:                "1A",
				Ram:                "1024",
				Disk:               "size=10",
				Dailybackup:        "no",
				Managed:            "no",
				Network:            "name=wan,ip=auto",
				Quantity:           1,
				BillingCycle:       "hourly",
				MonthlyPackage:     "",
				Poweronaftercreate: "yes",
				ScriptFile:         "",
				UserdataFile:       "",
				Tag:                "",
			},
		)
		if err != nil {
			return nil, err
		}
		serverNameCommandIds[serverName] = res.([]interface{})[0].(string)
	}
	results, err := waitCommands(
		ctx,
		ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
		serverNameCommandIds,
	)
	if err != nil {
		return nil, err
	}
	var servers []Server
	for serverName, _ := range results {
		servers = append(servers, Server{
			Name:   serverName,
			Tags: []string{},
		})
	}
	return servers, nil
}

func (c *KamateraApiClientRest) getServerTags(ctx context.Context, serverName string) ([]string, error) {
	res, err := request(
		ctx,
		ProviderConfig{ApiUrl: c.baseURL, ApiClientID: c.clientId, ApiSecret: c.secret},
		"GET",
		"/server/tags",
		KamateraServerPostRequest{ServerName: serverName},
	)
	if err != nil {
		return nil, err
	}
	var tags []string
	for _, row := range res.([]interface{}) {
		row := row.(map[string]interface{})
		tags = append(tags, row["tag name"].(string))
	}
	return tags, nil
}
