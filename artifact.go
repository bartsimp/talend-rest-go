package talend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const ArtifactsUrl string = DefaultRestUrl + "/artifacts"

type jArtifact struct {
	Items []struct {
		Id        string   `json:"id"`
		Name      string   `json:"name"`
		Type      string   `json:"type"`
		Versions  []string `json:"versions"`
		Workspace struct {
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
		} `json:"workspace"`
	} `json:"items"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ArtifactQuery struct {
	Name          string `url:"name,omitempty"`
	WorkspaceId   string `url:"workspaceId,omitempty"`
	EnvironmentId string `url:"environmentId,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Offset        int    `url:"offset,omitempty"`
}

func (c *Client) GetArtifacts(artifactQuery ArtifactQuery) (*jArtifact, error) {
	queryParms, _ := query.Values(artifactQuery)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(ArtifactsUrl+"?%s", queryParms.Encode()), nil)

	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var artifacts jArtifact
	err = json.Unmarshal(res, &artifacts)
	if err != nil {
		return nil, err
	}

	return &artifacts, nil
}
