package talend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const RemoteEngineClusterUrl string = DefaultRestUrl + "/runtimes/remote-engine-clusters"

type RemoteEngineCluster struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Workspace   struct {
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
	} `json:"workspace"`
	CreateDate   int    `json:"createDate"`
	UpdateDate   int    `json:"updateDate"`
	RuntimeId    string `json:"runtimeId"`
	Availability string `json:"availability"`
	Managed      bool   `json:"managed"`
}

type ClusterRunProfile struct {
	Id           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	CreateDate   string   `json:"createDate,omitempty"`
	UpdateDate   string   `json:"updateDate,omitempty"`
	Type         string   `json:"type,omitempty"`
	JvmArguments []string `json:"jvmArguments,omitempty"`
	RuntimeId    string   `json:"runtimeId,omitempty"`
	Version      int      `json:"version,omitempty"`
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

func (c *Client) GetRemoteEngineClustersByClusterId(clusterId string) (*RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineClusterUrl+"/%s", clusterId), nil)
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

func (c *Client) DeleteRemoteEngineClusters(clusterId string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(RemoteEngineClusterUrl+"/%s", clusterId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetRemoteEngineClustersRunProfile(clusterId string) (*[]ClusterRunProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s/run-profiles", clusterId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var crp []ClusterRunProfile

	err = json.Unmarshal(res, &crp)
	if err != nil {
		return nil, err
	}

	return &crp, nil
}

func (c *Client) GetRemoteEngineClustersRunProfileByProfileId(clusterId string, runProfileId string) (*[]ClusterRunProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s/run-profiles/%s", clusterId, runProfileId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var crp []ClusterRunProfile

	err = json.Unmarshal(res, &crp)
	if err != nil {
		return nil, err
	}

	return &crp, nil
}
