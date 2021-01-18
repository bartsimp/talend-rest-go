package talend

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const WorkspacesUrl string = defaultRestURL + "/workspaces"

type Workspace struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Type        string `json:"type"`
	Environment struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Default     bool   `json:"default"`
	} `json:"environment"`
}

func (c *Client) GetWorkspaces(query string) ([]Workspace, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(WorkspacesUrl+"?query=%s", query), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var workspaces []Workspace

	err = json.Unmarshal(res, &workspaces)
	if err != nil {
		return nil, err
	}

	return workspaces, nil
}
