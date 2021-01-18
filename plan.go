package talend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/google/go-querystring/query"
)

const plansURL string = defaultRestURL + "/executables/plans"

type PlanQuery struct {
	Name          string `url:"name,omitempty"`
	TaskID        string `url:"taskId,omitempty"`
	EnvironmentID string `url:"environmentId,omitempty"`
	WorkspaceID   string `url:"workspaceId,omitempty"`
	Limit         int    `url:"limit,omitempty"`
	Offset        int    `url:"offset,omitempty"`
}

type PlanCreate struct {
	Name      string `json:"name"`
	Workspace struct {
		Name        string `json:"name"`
		Environment string `json:"environment"`
	} `json:"workspace"`
	Description string `json:"description"`
	Steps       []struct {
		Name             string   `json:"name"`
		Condition        string   `json:"condition"`
		TaskIds          []string `json:"taskIds"`
		HandlerOnFailure struct {
			TaskIds []string `json:"taskIds,omitempty"`
			PlanIds []string `json:"planIds,omitempty"`
		} `json:"handlerOnFailure,omitempty"`
	} `json:"steps"`
}

type PlanCreated struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WorkspaceID string `json:"workspaceId"`
	//	Artifact     ArtifactCreated `json:"artifact"`
}

type PlanRunConfigRequest struct {
	Name      string `json:"name"`
	Workspace struct {
		Name        string `json:"name"`
		Environment string `json:"environment"`
	} `json:"workspace"`
	PlanRunConfig PlanRunConfig `json:"planRunConfig"`
}

type PlanRunConfigResponse struct {
	PlanID        string        `json:"planId"`
	PlanRunConfig PlanRunConfig `json:"planRunConfig"`
}

type PlanRunConfig struct {
	Trigger struct {
		Type      string `json:"type"`
		Interval  int    `json:"interval,omitempty"`
		StartDate string `json:"startDate"`
		TimeZone  string `json:"timeZone"`
		AtTimes   struct {
			Type      string   `json:"type"`
			Times     []string `json:"times,omitempty"`
			Time      string   `json:"time"`
			StartTime string   `json:"startTime"`
			EndTime   string   `json:"endTime"`
			Interval  string   `json:"interval"`
		} `json:"atTimes,omitempty"`
		AtDays struct {
			Type string   `json:"type"`
			Day  int      `json:"day,omitempty"`
			Days []string `json:"days,omitempty"`
		} `json:"atDays,omitempty"`
		Webhook struct {
			Name           string `json:"name"`
			Description    string `json:"description,omitempty"`
			TriggerCalls   int    `json:"triggerCalls,omitempty"`
			TriggerTimeout int    `json:"triggerTimeout,omitempty"`
			RunAsUser      string `json:"runAsUser,omitempty"`
			NewURL         bool   `json:"newUrl,omitempty"`
			URL            string `json:"url,omitempty"`
		} `json:"webhook,omitempty"`
	} `json:"trigger,omitempty"`
	Runtime struct {
		Type         string `json:"type"`         // "type": "CLOUD",
		ID           string `json:"id"`           // "id": "5c9a51dc7b353e43c7fc787c",
		EngineID     string `json:"engineId"`     // "engineId": "5c9a51dc7b353e43c7fc787c",
		ClusterID    string `json:"clusterId"`    // "clusterId": "5c9a51dc7b353e43c7fc787c",
		RunProfileID string `json:"runProfileId"` // "runProfileId": "5c9a51dc7b353e43c7fc783c"
	} `json:"runtime,omitempty"`
	ParallelExecutionAllowed bool `json:"parallelExecutionAllowed,omitempty"`
}

type PlanAvailable struct {
	Items []struct {
		Executable string `json:"executable"`
		Name       string `json:"name"`
		Workspace  struct {
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
		ArtifactID string `json:"artifactId"`
		Runtime    struct {
			Type string `json:"type"`
			ID   string `json:"id"`
		} `json:"runtime"`
	} `json:"items"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type PlanByID struct {
	Executable string `json:"executable"`
	Name       string `json:"name"`
	Workspace  struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Owner       string `json:"owner,omitempty"`
		Type        string `json:"type"`
		Environment struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Default     bool   `json:"default"`
		} `json:"environment,omitempty"`
	} `json:"workspace,omitempty"`
	Description string        `json:"description,omitempty"`
	Chart       ExecutionStep `json:"chart,omitempty"`
	Status      string        `json:"status,omitempty"`
}

type ExecutionStep struct {
	StepID          string         `json:"stepId,omitempty"`
	StepName        string         `json:"stepName,omitempty"`
	Status          string         `json:"status,omitempty"`
	StepOnException *ExecutionStep `json:"stepOnException,omitempty"`
	NextStep        *ExecutionStep `json:"nextStep,omitempty"`
	Flows           []struct {
		ID              string `json:"id"`
		Name            string `json:"name"`
		Version         string `json:"version"`
		Description     string `json:"description"`
		Destination     string `json:"destination"`
		WorkspaceID     string `json:"workspaceId"`
		Plan            bool   `json:"plan"`
		ArtifactVersion struct {
			ID                 string `json:"id"`
			Name               string `json:"name"`
			Version            string `json:"version"`
			Type               string `json:"type"`
			Publisher          string `json:"publisher"`
			MarketplaceProduct string `json:"marketplaceProduct"`
			Parameters         []struct {
				Name     string `json:"name"`
				Value    string `json:"value"`
				Type     string `json:"type"`
				Required string `json:"required"`
				Comment  string `json:"comment"`
			} `json:"parameters"`
			Workspace struct {
				ID          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Owner       string `json:"owner"`
				Type        string `json:"type"`
				Environment struct {
					ID      string `json:"id"`
					Name    string `json:"name"`
					Default bool   `json:"default"`
				} `json:"environment"`
			} `json:"workspace"`
		} `json:"artifactVersion"`
		UpgradeInfo struct {
			Upgradable        bool `json:"upgradable"`
			AutoUpgradeFailed bool `json:"autoUpgradeFailed"`
		} `json:"upgradeInfo"`
	} `json:"flows,omitempty"`
}

type jPlanCreate struct {
	Name        string            `json:"name"`
	WorkspaceID string            `json:"workspaceId"`
	Description string            `json:"description,omitempty"`
	Steps       []jPlanCreateStep `json:"steps,omitempty"`
}

type jPlanCreateStep struct {
	Name             string   `json:"name"`
	Condition        string   `json:"condition"`
	TaskIds          []string `json:"taskIds"`
	HandlerOnFailure struct {
		TaskIds []string `json:"taskIds,omitempty"`
		PlanIds []string `json:"planIds,omitempty"`
	} `json:"handlerOnFailure,omitempty"`
}

func (c *Client) GetPlans(planQuery PlanQuery) (*PlanAvailable, error) {
	queryParms, _ := query.Values(planQuery)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(plansURL+"?%s", queryParms.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var plans PlanAvailable

	err = json.Unmarshal(res, &plans)
	if err != nil {
		return nil, err
	}

	return &plans, nil
}

func (c *Client) GetPlanByID(id string) (*PlanByID, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(plansURL+"/%s", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var plan PlanByID

	err = json.Unmarshal(res, &plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func (c *Client) CreatePlanFromPlainFile(jsonPlanFile string) (*PlanCreated, error) {
	content, err := ioutil.ReadFile(jsonPlanFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.CreatePlanFromPlainJSON(text)
}

func (c *Client) CreatePlanFromPlainJSON(jsonPlan string) (*PlanCreated, error) {

	var planCreate PlanCreate

	err := json.Unmarshal([]byte(jsonPlan), &planCreate)
	if err != nil {
		return nil, err
	}

	return c.CreatePlan(&planCreate)
}

func (c *Client) CreatePlan(p *PlanCreate) (*PlanCreated, error) {
	jplan, err := c.ParsePlan(p)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(jplan)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, j, "", "  ")
	return c.CreatePlanFromRawJson(string(prettyJSON.Bytes()))
}

func (c *Client) CreatePlanFromRawJson(jsonRequest string) (*PlanCreated, error) {
	req, err := http.NewRequest(http.MethodPost, plansURL, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var plan PlanCreated

	err = json.Unmarshal(res, &plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func (c *Client) CreatePlanFromRawFile(jsonRawFile string) (*PlanCreated, error) {
	content, err := ioutil.ReadFile(jsonRawFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.CreatePlanFromRawJson(text)
}

func (c *Client) DeletePlan(id string) error {
	req, err := http.NewRequest(http.MethodDelete, plansURL+"/"+id, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ParsePlan(p *PlanCreate) (*jPlanCreate, error) {

	var jplan jPlanCreate

	jplan.Name = p.Name
	jplan.Description = p.Description

	workspaceQuery := "name==" + p.Workspace.Name + ";environment.name==" + p.Workspace.Environment
	workspaces, err := c.GetWorkspaces(workspaceQuery)
	if err != nil {
		return nil, err
	}
	if len(workspaces) == 0 {
		return nil, fmt.Errorf("workspace not found (%s)", workspaceQuery)
	}
	if len(workspaces) > 1 {
		return nil, fmt.Errorf("workspace not unique (%s)", workspaceQuery)
	}
	jplan.WorkspaceID = workspaces[0].Id

	jplan.Steps = make([]jPlanCreateStep, len(p.Steps))
	for i := range p.Steps {
		jplan.Steps[i].Name = p.Steps[i].Name
		jplan.Steps[i].Condition = p.Steps[i].Condition

		jplan.Steps[i].TaskIds = make([]string, len(p.Steps[i].TaskIds))
		for j := range p.Steps[i].TaskIds {
			tasks, err := c.GetTasks(TaskQuery{
				Name:        p.Steps[i].TaskIds[j],
				WorkspaceId: workspaces[0].Id,
				Limit:       100, Offset: 0})
			if err != nil {
				return nil, err
			}
			if len(tasks.Items) == 0 {
				return nil, fmt.Errorf("task not found (%s)", workspaceQuery)
			}
			if len(tasks.Items) > 1 {
				return nil, fmt.Errorf("task not unique (%s)", workspaceQuery)
			}
			jplan.Steps[i].TaskIds[j] = tasks.Items[0].Executable
		}

		jplan.Steps[i].HandlerOnFailure.TaskIds = make([]string, len(p.Steps[i].HandlerOnFailure.TaskIds))
		for k := range p.Steps[i].HandlerOnFailure.TaskIds {
			tasks, err := c.GetTasks(TaskQuery{
				Name:        p.Steps[i].HandlerOnFailure.TaskIds[k],
				WorkspaceId: workspaces[0].Id,
				Limit:       100, Offset: 0})
			if err != nil {
				return nil, err
			}
			if len(tasks.Items) == 0 {
				return nil, fmt.Errorf("task not found (%s)", workspaceQuery)
			}
			if len(tasks.Items) > 1 {
				return nil, fmt.Errorf("task not unique (%s)", workspaceQuery)
			}
			jplan.Steps[i].HandlerOnFailure.TaskIds[k] = tasks.Items[0].Executable
		}
	}

	return &jplan, nil
}

func (c *Client) UpdatePlanFromRawJson(planId string, jsonRequest string) (*PlanCreated, error) {
	req, err := http.NewRequest(http.MethodPut, plansURL+"/"+planId, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var plan PlanCreated

	err = json.Unmarshal(res, &plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func (c *Client) UpdatePlanFromRawFile(planId string, jsonRawFile string) (*PlanCreated, error) {
	content, err := ioutil.ReadFile(jsonRawFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.UpdatePlanFromRawJson(planId, text)
}

func (c *Client) UpdatePlanFromPlainFile(planId string, jsonPlanFile string) (*PlanCreated, error) {
	content, err := ioutil.ReadFile(jsonPlanFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.UpdatePlanFromPlainJson(planId, text)
}

func (c *Client) UpdatePlanFromPlainJson(planId string, jsonPlan string) (*PlanCreated, error) {

	var planCreate PlanCreate

	err := json.Unmarshal([]byte(jsonPlan), &planCreate)
	if err != nil {
		return nil, err
	}

	return c.UpdatePlan(planId, &planCreate)
}

func (c *Client) UpdatePlan(planId string, p *PlanCreate) (*PlanCreated, error) {
	jplan, err := c.ParsePlan(p)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(jplan)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, j, "", "  ")
	return c.UpdatePlanFromRawJson(planId, string(prettyJSON.Bytes()))
}

func (c *Client) GetPlanRunConfigByPlanId(planId string) (*PlanRunConfigResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(plansURL+"/%s/run-config", planId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var planrunconfig PlanRunConfig

	err = json.Unmarshal(res, &planrunconfig)
	if err != nil {
		return nil, err
	}

	var planRunConfigResponse PlanRunConfigResponse

	planRunConfigResponse.PlanID = planId
	planRunConfigResponse.PlanRunConfig = planrunconfig

	return &planRunConfigResponse, nil
}

func (c *Client) DeletePlanRunConfigByPlanId(planId string) error {
	req, err := http.NewRequest(http.MethodDelete, plansURL+"/"+planId+"/run-config", nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdatePlanRunConfigFromRawFile(planId string, jsonRawFile string) (*PlanRunConfigResponse, error) {
	content, err := ioutil.ReadFile(jsonRawFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.UpdatePlanRunConfigFromRawJson(planId, text)
}

func (c *Client) UpdatePlanRunConfigFromRawJson(planId string, jsonRequest string) (*PlanRunConfigResponse, error) {
	req, err := http.NewRequest(http.MethodPut, plansURL+"/"+planId+"/run-config", strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var planrunconfig PlanRunConfig

	err = json.Unmarshal(res, &planrunconfig)
	if err != nil {
		return nil, err
	}

	var planrunconfigResponse PlanRunConfigResponse
	planrunconfigResponse.PlanID = planId
	planrunconfigResponse.PlanRunConfig = planrunconfig

	return &planrunconfigResponse, nil
}

func (c *Client) UpdatePlanRunConfigFromPlainFile(jsonPlanRunConfigFile string) (*PlanRunConfigResponse, error) {
	content, err := ioutil.ReadFile(jsonPlanRunConfigFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)

	return c.UpdatePlanRunConfigFromPlainJson(text)
}

func (c *Client) UpdatePlanRunConfigFromPlainJson(jsonPlanRunConfig string) (*PlanRunConfigResponse, error) {

	var planRunConfigRequest PlanRunConfigRequest

	err := json.Unmarshal([]byte(jsonPlanRunConfig), &planRunConfigRequest)
	if err != nil {
		return nil, err
	}

	return c.UpdatePlanRunConfig(&planRunConfigRequest)
}

func (c *Client) UpdatePlanRunConfig(planRunConfig *PlanRunConfigRequest) (*PlanRunConfigResponse, error) {

	workspaceQuery := "name==" + planRunConfig.Workspace.Name + ";environment.name==" + planRunConfig.Workspace.Environment
	workspaces, err := c.GetWorkspaces(workspaceQuery)
	if err != nil {
		return nil, err
	}
	if len(workspaces) > 1 {
		return nil, fmt.Errorf("workspace not unique")
	}
	workspaceId := workspaces[0].Id

	plans, err := c.GetPlans(PlanQuery{Name: planRunConfig.Name, WorkspaceID: workspaceId, Limit: 100, Offset: 0})
	if len(plans.Items) > 1 {
		return nil, fmt.Errorf("plan not unique")
	}
	planId := plans.Items[0].Executable

	j, err := json.Marshal(planRunConfig.PlanRunConfig)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, j, "", "  ")

	return c.UpdatePlanRunConfigFromRawJson(planId, string(prettyJSON.Bytes()))
}
