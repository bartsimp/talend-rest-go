package talend

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const artifactsURL string = defaultRestURL + "/artifacts"

// JArtifact Artifact response details
type JArtifact struct {
	Items []struct {
		ID        string   `json:"id"`
		Name      string   `json:"name"`
		Type      string   `json:"type"`
		Versions  []string `json:"versions"`
		Workspace struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Owner       string `json:"owner"`
			Type        string `json:"type"`
			Environment struct {
				ID          string `json:"id"`
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

// ArtifactQuery Artifact query details
type ArtifactQuery struct {
	Name          string `url:"name,omitempty"`
	WorkspaceID   string `url:"workspaceId,omitempty"`
	EnvironmentID string `url:"environmentId,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Offset        int    `url:"offset,omitempty"`
}

// GetArtifacts Get available Artifacts
func (c *Client) GetArtifacts(artifactQuery ArtifactQuery) (*JArtifact, error) {
	queryParms, _ := query.Values(artifactQuery)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(artifactsURL+"?%s", queryParms.Encode()), nil)

	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var artifacts JArtifact
	err = json.Unmarshal(res, &artifacts)
	if err != nil {
		return nil, err
	}

	return &artifacts, nil
}
