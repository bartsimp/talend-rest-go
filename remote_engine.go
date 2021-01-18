package talend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const RemoteEngineUrl string = defaultRestURL + "/runtimes/remote-engines"

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

type RuntimeRunProfile struct {
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

func (c *Client) GetRemoteEnginesByEngineId(engineId string) (*RemoteEngine, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s", engineId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var re RemoteEngine

	err = json.Unmarshal(res, &re)
	if err != nil {
		return nil, err
	}

	return &re, nil
}

func (c *Client) DeleteRemoteEngines(engineId string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(RemoteEngineUrl+"/%s", engineId), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetRemoteEnginesRunProfile(engineId string) (*[]RuntimeRunProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s/run-profiles", engineId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rerp []RuntimeRunProfile

	err = json.Unmarshal(res, &rerp)
	if err != nil {
		return nil, err
	}

	return &rerp, nil
}

func (c *Client) GetRemoteEnginesRunProfileByProfileId(engineId string, runProfileId string) (*RuntimeRunProfile, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(RemoteEngineUrl+"/%s/run-profiles/%s", engineId, runProfileId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var rerp RuntimeRunProfile

	err = json.Unmarshal(res, &rerp)
	if err != nil {
		return nil, err
	}

	return &rerp, nil
}
