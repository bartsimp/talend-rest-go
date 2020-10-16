package talend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const RemoteEngineClusterUrl string = DefaultRestUrl + "/runtimes/remote-engine-clusters"

type RemoteEngineCluster struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Workspace    Workspaceinfo `json:"workspace"`
	CreateDate   int           `json:"createDate"`
	UpdateDate   int           `json:"updateDate"`
	RuntimeId    string        `json:"runtimeId"`
	Availability string        `json:"availability"`
	Managed      bool          `json:"managed"`
}

type Workspaceinfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       string `json:"workspace"`
	Type        string `json:"type"`
	Environment struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"environment"`
}

func (c *Client) GetRemoteEngineClusters(searchQuery string) (*[]RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineClusterUrl+"?_s=%s", searchQuery), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rec []RemoteEngineCluster

	err = json.Unmarshal(res, &rec)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}

func (c *Client) GetRemoteEngineClustersById(id string) (*RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineClusterUrl+"/%s", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rec RemoteEngineCluster

	err = json.Unmarshal(res, &rec)
	if err != nil {
		return nil, err
	}

	return &rec, nil
}

func (c *Client) CreateRemoteEngineClustersFromRawJson(jsonRequest string) (*RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodPost, RemoteEngineClusterUrl, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var response RemoteEngineCluster

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
