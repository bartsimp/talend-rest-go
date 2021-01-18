package talend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const remoteEngineClusterURL string = defaultRestURL + "/runtimes/remote-engine-clusters"

type RemoteEngineCluster struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Workspace   struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Owner       string `json:"workspace"`
		Type        string `json:"type"`
		Environment struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Default     bool   `json:"default"`
		} `json:"environment"`
	} `json:"workspace"`
	CreateDate   int    `json:"createDate"`
	UpdateDate   int    `json:"updateDate"`
	RuntimeID    string `json:"runtimeId"`
	Availability string `json:"availability"`
	Managed      bool   `json:"managed"`
}

type ClusterRunProfile struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	CreateDate   string   `json:"createDate,omitempty"`
	UpdateDate   string   `json:"updateDate,omitempty"`
	Type         string   `json:"type,omitempty"`
	JvmArguments []string `json:"jvmArguments,omitempty"`
	RuntimeID    string   `json:"runtimeId,omitempty"`
	Version      int      `json:"version,omitempty"`
}

// GetRemoteEngineClusters return remote engines cluster details
func (c *Client) GetRemoteEngineClusters(searchQuery string) (*[]RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(remoteEngineClusterURL+"?_s=%s", searchQuery), nil)
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

func (c *Client) GetRemoteEngineClustersByClusterID(clusterID string) (*RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(remoteEngineClusterURL+"/%s", clusterID), nil)
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

func (c *Client) CreateRemoteEngineClustersFromRawJSON(jsonRequest string) (*RemoteEngineCluster, error) {
	req, err := http.NewRequest(http.MethodPost, remoteEngineClusterURL, strings.NewReader(jsonRequest))
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

func (c *Client) DeleteRemoteEngineClusters(clusterID string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(remoteEngineClusterURL+"/%s", clusterID), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetRemoteEngineClustersRunProfile(clusterID string) (*[]ClusterRunProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s/run-profiles", clusterID), nil)
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

func (c *Client) GetRemoteEngineClustersRunProfileByProfileID(clusterID string, runProfileID string) (*[]ClusterRunProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s/run-profiles/%s", clusterID, runProfileID), nil)
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
