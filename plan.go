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

const PlansUrl string = DefaultRestUrl + "/executables/plans"

type PlanQuery struct {
	Name          string `url:"name,omitempty"`
	TaskId        string `url:"taskId,omitempty"`
	EnvironmentId string `url:"environmentId,omitempty"`
	WorkspaceId   string `url:"workspaceId,omitempty"`
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
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WorkspaceId string `json:"workspaceId"`
	//	Artifact     ArtifactCreated `json:"artifact"`
}

type PlanAvailable struct {
	Items []struct {
		Executable string    `json:"executable"`
		Name       string    `json:"name"`
		Workspace  Workspace `json:"workspace"`
		ArtifactId string    `json:"artifactId"`
		Runtime    struct {
			Type string `json:"type"`
			Id   string `json:"id"`
		} `json:"runtime"`
	} `json:"items"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type PlanById struct {
	Executable string `json:"executable"`
	Name       string `json:"name"`
	Workspace  struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Owner       string `json:"owner,omitempty"`
		Type        string `json:"type"`
		Environment struct {
			Id          string `json:"id"`
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
	StepId          string         `json:"stepId,omitempty"`
	StepName        string         `json:"stepName,omitempty"`
	Status          string         `json:"status,omitempty"`
	StepOnException *ExecutionStep `json:"stepOnException,omitempty"`
	NextStep        *ExecutionStep `json:"nextStep,omitempty"`
	Flows           []struct {
		Id              string `json:"id"`
		Name            string `json:"name"`
		Version         string `json:"version"`
		Description     string `json:"description"`
		Destination     string `json:"destination"`
		WorkspaceId     string `json:"workspaceId"`
		Plan            bool   `json:"plan"`
		ArtifactVersion struct {
			Id                 string `json:"id"`
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
				Id          string `json:"id"`
				Name        string `json:"name"`
				Description string `json:"description"`
				Owner       string `json:"owner"`
				Type        string `json:"type"`
				Environment struct {
					Id      string `json:"id"`
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
	WorkspaceId string            `json:"workspaceId"`
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
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(PlansUrl+"?%s", queryParms.Encode()), nil)
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

func (c *Client) GetPlanById(id string) (*PlanById, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(PlansUrl+"/%s", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var plan PlanById

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
	return c.CreatePlanFromPlainJson(text)
}

func (c *Client) CreatePlanFromPlainJson(jsonPlan string) (*PlanCreated, error) {

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
	req, err := http.NewRequest(http.MethodPost, PlansUrl, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dst := &bytes.Buffer{}
	if err := json.Indent(dst, res, "", "  "); err != nil {
		panic(err)
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
	req, err := http.NewRequest(http.MethodDelete, PlansUrl+"/"+id, nil)
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
	if len(workspaces) > 1 {
		return nil, fmt.Errorf("workspace not unique")
	}
	jplan.WorkspaceId = workspaces[0].Id

	jplan.Steps = make([]jPlanCreateStep, len(p.Steps))
	for i := range p.Steps {
		jplan.Steps[i].Name = p.Steps[i].Name
		jplan.Steps[i].Condition = p.Steps[i].Condition

		jplan.Steps[i].TaskIds = make([]string, len(p.Steps[i].TaskIds))
		for j := range p.Steps[i].TaskIds {
			task, err := c.GetTasks(TaskQuery{
				Name:        p.Steps[i].TaskIds[j],
				WorkspaceId: workspaces[0].Id,
				Limit:       100, Offset: 0})

			if err != nil {
				return nil, err
			}

			if len(task.Items) > 1 {
				return nil, fmt.Errorf("task not unique")
			}

			jplan.Steps[i].TaskIds[j] = task.Items[0].Executable
		}

		jplan.Steps[i].HandlerOnFailure.TaskIds = make([]string, len(p.Steps[i].HandlerOnFailure.TaskIds))
		for k := range p.Steps[i].HandlerOnFailure.TaskIds {
			task, err := c.GetTasks(TaskQuery{
				Name:        p.Steps[i].HandlerOnFailure.TaskIds[k],
				WorkspaceId: workspaces[0].Id,
				Limit:       100, Offset: 0})

			if err != nil {
				return nil, err
			}

			if len(task.Items) > 1 {
				return nil, fmt.Errorf("task not unique")
			}

			jplan.Steps[i].HandlerOnFailure.TaskIds[k] = task.Items[0].Executable
		}
	}

	return &jplan, nil
}

func (c *Client) UpdatePlanFromRawJson(planId string, jsonRequest string) (*PlanCreated, error) {
	req, err := http.NewRequest(http.MethodPut, PlansUrl+"/"+planId, strings.NewReader(jsonRequest))
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
