package talend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const RemoteEngineUrl string = DefaultRestUrl + "/runtimes/remote-engine"

type RemoteEngine struct {
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
	CreateDate    int      `json:"createDate"`
	UpdateDate    int      `json:"updateDate"`
	RuntimeId     string   `json:"runtimeId"`
	Availability  string   `json:"availability"`
	Managed       bool     `json:"managed"`
	EnvironmentId string   `json:"environmentId"`
	WorkspaceId   string   `json:"workspaceId"`
	Status        string   `json:"status"`
	RunProfiles   []string `json:"runProfiles"`
	Debug         struct {
		Host string `json:"host"`
	} `json:"debug"`
	ClusterId        string `json:"clusterId"`
	PreAuthorizedKey string `json:"preAuthorizedKey"`
}

func (c *Client) GetRemoteEngines(searchQuery string) (*[]RemoteEngine, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"?query=%s", searchQuery), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var re []RemoteEngine

	err = json.Unmarshal(res, &re)
	if err != nil {
		return nil, err
	}

	return &re, nil
}
