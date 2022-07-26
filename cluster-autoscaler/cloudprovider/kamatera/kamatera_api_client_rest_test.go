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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	. "k8s.io/autoscaler/cluster-autoscaler/utils/test"
	"testing"
)

const (
	mockKamateraClientId = "mock-client-id"
	mockKamateraSecret = "mock-secret"
)

func TestApiClientRest_ListServersByTag_NoServers(t *testing.T) {
	server := NewHttpServerMock(MockFieldContentType, MockFieldResponse)
	defer server.Close()
	ctx := context.Background()
	client := NewKamateraApiClientRest(mockKamateraClientId, mockKamateraSecret)
	client.SetBaseURL(server.URL)
	server.On("handle", "/service/servers").Return(
		"application/json",
		`[]`,
	).Once()
	servers, err := client.ListServersByTag(ctx, "test-tag")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(servers))
	mock.AssertExpectationsForObjects(t, server)
}

func TestApiClientRest_ListServersByTag_NoMatchingServerTags(t *testing.T) {
	server := NewHttpServerMock(MockFieldContentType, MockFieldResponse)
	defer server.Close()
	ctx := context.Background()
	client := NewKamateraApiClientRest(mockKamateraClientId, mockKamateraSecret)
	client.SetBaseURL(server.URL)
	serverName1 := mockKamateraServerName()
	serverName2 := mockKamateraServerName()
	server.On("handle", "/service/servers").Return(
		"application/json",
		fmt.Sprintf(`[{"name": "%s"}, {"name": "%s"}]`, serverName1, serverName2),
	).On("handle", "/server/tags").Return(
		"application/json",
		`[]`,
	)
	servers, err := client.ListServersByTag(ctx, "test-tag")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(servers))
	mock.AssertExpectationsForObjects(t, server)
}

func TestApiClientRest_ListServersByTag_MatchingServerTags(t *testing.T) {
	server := NewHttpServerMock(MockFieldContentType, MockFieldResponse)
	defer server.Close()
	ctx := context.Background()
	client := NewKamateraApiClientRest(mockKamateraClientId, mockKamateraSecret)
	client.SetBaseURL(server.URL)
	serverName := mockKamateraServerName()
	server.On("handle", "/service/servers").Return(
		"application/json",
		fmt.Sprintf(`[{"name": "%s"}]`, serverName),
	).On("handle", "/server/tags").Return(
		"application/json",
		`[{"tag name": "test-tag"}]`,
	)
	servers, err := client.ListServersByTag(ctx, "test-tag")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(servers))
	assert.Equal(t, servers[0].Name, serverName)
	mock.AssertExpectationsForObjects(t, server)
}

func TestApiClientRest_DeleteServer(t *testing.T) {
	server := NewHttpServerMock(MockFieldContentType, MockFieldResponse)
	defer server.Close()
	ctx := context.Background()
	client := NewKamateraApiClientRest(mockKamateraClientId, mockKamateraSecret)
	client.SetBaseURL(server.URL)
	serverName := mockKamateraServerName()
	commandId := "mock-command-id"
	server.On("handle", "/service/server/poweroff").Return(
		"application/json",
		fmt.Sprintf(`["%s"]`, commandId),
	).Once().On("handle", "/service/queue").Return(
		"application/json",
		`[{"status": "complete"}]`,
	).Once().On("handle", "/service/server/delete").Return(
		"application/json",
		"{}",
	).Once()
	err := client.DeleteServer(ctx, serverName)
	assert.NoError(t, err)
	mock.AssertExpectationsForObjects(t, server)
}

func TestApiClientRest_CreateServers(t *testing.T) {
	server := NewHttpServerMock(MockFieldContentType, MockFieldResponse)
	defer server.Close()
	ctx := context.Background()
	client := NewKamateraApiClientRest(mockKamateraClientId, mockKamateraSecret)
	client.SetBaseURL(server.URL)
	commandId := "command"
	server.On("handle", "/service/server").Return(
		"application/json",
		fmt.Sprintf(`["%s"]`, commandId),
	).Twice().On("handle", "/service/queue").Return(
		"application/json",
		`[{"status": "complete"}]`,
	).Twice()
	servers, err := client.CreateServers(ctx, 2, mockServerConfig("test", []string{"foo", "bar"}))
	assert.NoError(t, err)
	assert.Equal(t, 2, len(servers))
	assert.Less(t, 10, len(servers[0].Name))
	assert.Less(t, 10, len(servers[1].Name))
	assert.Equal(t, servers[0].Tags, []string{"foo", "bar"})
	assert.Equal(t, servers[1].Tags, []string{"foo", "bar"})
	mock.AssertExpectationsForObjects(t, server)
}
