package talend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-querystring/query"
)

const TasksUrl string = defaultRestURL + "/executables/tasks"

type TaskQuery struct {
	Name                string `url:"name,omitempty"`
	EnvironmentId       string `url:"environmentId,omitempty"`
	WorkspaceId         string `url:"workspaceId,omitempty"`
	Limit               int    `url:"limit,omitempty"`
	Offset              int    `url:"offset,omitempty"`
	ArtifactId          string `url:"artifactId,omitempty"`
	RuntimeType         string `url:"runtimeType,omitempty"`
	RuntimeId           string `url:"runtimeId,omitempty"`
	RuntimeRunProfileId string `url:"runtimeRunProfileId,omitempty"`
}

type TaskCreate struct {
	Name      string `json:"name"`
	Workspace struct {
		Name        string `json:"name"`
		Environment string `json:"environment"`
	} `json:"workspace"`
	Description string            `json:"description"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	Resources   map[string]string `json:"resources,omitempty"`
	Artifact    struct {
		Name string `json:"name"`
	} `json:"artifact"`
	//	AutoUpgradeInfo struct {
	//		AutoUpgradable                bool
	//		OverrideWithDefaultParameters bool
	//		AutoUpgradeFailed             bool
	//	} `json:"autoUpgradeInfo"`
	//	EnvironmentId string `json:"environmentId,omitempty"`
	Environment string `json:"environment"`
}

type TaskCreated struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	WorkspaceId string `json:"workspaceId"`
	Artifact    struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"artifact"`
}

type jTaskAvailable struct {
	Items []struct {
		Executable string `json:"executable"`
		Name       string `json:"name"`
		Workspace  struct {
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
		ArtifactId string `json:"artifactId"`
		Runtime    struct {
			Type string `json:"type"`
			Id   string `json:"id"`
		} `json:"runtime"`
	} `json:"items"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type TaskById struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Workspace   struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Owner       string `json:"owner"`
		Type        string `json:"type"`
		Environment struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Default     bool   `json:"default"`
		} `json:"environment"`
	} `json:"workspace"`
	Version  string `json:"version"`
	Artifact struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"artifact"`
	Tags            []string          `json:"tags,omitempty"`
	Connections     map[string]string `json:"connections,omitempty"`
	Parameters      map[string]string `json:"parameters,omitempty"`
	Resources       map[string]string `json:"resources,omitempty"`
	AutoUpgradeInfo struct {
		AutoUpgradable                bool `json:"autoUpgradable"`
		OverrideWithDefaultParameters bool `json:"overrideWithDefaultParameters"`
		AutoUpgradeFailed             bool `json:"autoUpgradeFailed"`
	} `json:"autoUpgradeInfo,omitempty"`
}

type jTaskCreate struct {
	WorkspaceId string            `json:"workspaceId"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Tags        []string          `json:"tags,omitempty"`
	Connections map[string]string `json:"connections,omitempty"`
	Parameters  map[string]string `json:"parameters,omitempty"`
	Resources   map[string]string `json:"resources,omitempty"`
	Artifact    struct {
		Id      string `json:"id"`
		Version string `json:"version"`
	} `json:"artifact"`
	//	AutoUpgradeInfo struct {
	//		AutoUpgradable                bool `json:"autoUpgradable"`
	//		OverrideWithDefaultParameters bool `json:"overrideWithDefaultParameters"`
	//		AutoUpgradeFailed             bool `json:"autoUpgradeFailed"`
	//	} `json:"autoUpgradeInfo,omitempty"`
	EnvironmentId string `json:"environmentId,omitempty"`
}

type TaskRunConfigRequest struct {
	Name      string `json:"name"`
	Workspace struct {
		Name        string `json:"name"`
		Environment string `json:"environment"`
	} `json:"workspace"`
	TaskRunConfig TaskRunConfig `json:"taskRunConfig"`
}

type TaskRunConfigResponse struct {
	TaskId        string        `json:"taskId"`
	TaskRunConfig TaskRunConfig `json:"taskRunConfig"`
}

type TaskRunConfig struct {
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
			NewUrl         bool   `json:"newUrl,omitempty"`
			Url            string `json:"url,omitempty"`
		} `json:"webhook,omitempty"`
	} `json:"trigger,omitempty"`
	Runtime struct {
		Type         string `json:"type"`         // "type": "CLOUD",
		Id           string `json:"id"`           // "id": "5c9a51dc7b353e43c7fc787c",
		EngineId     string `json:"engineId"`     // "engineId": "5c9a51dc7b353e43c7fc787c",
		ClusterId    string `json:"clusterId"`    // "clusterId": "5c9a51dc7b353e43c7fc787c",
		RunProfileId string `json:"runProfileId"` // "runProfileId": "5c9a51dc7b353e43c7fc783c"
	} `json:"runtime,omitempty"`
	ParallelExecutionAllowed bool   `json:"parallelExecutionAllowed,omitempty"`
	LogLevel                 string `json:"logLevel,omitempty"`
	RemoteRunAsUser          string `json:"remoteRunAsUser,omitempty"`
}

func (c *Client) GetTasks(taskQuery TaskQuery) (*jTaskAvailable, error) {
	queryParms, _ := query.Values(taskQuery)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(TasksUrl+"?%s", queryParms.Encode()), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var tasks jTaskAvailable

	err = json.Unmarshal(res, &tasks)
	if err != nil {
		return nil, err
	}

	return &tasks, nil
}

func (c *Client) GetTaskById(id string) (*TaskById, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(TasksUrl+"/%s", id), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var task TaskById

	err = json.Unmarshal(res, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *Client) CreateTaskFromRawJson(jsonRequest string) (*TaskCreated, error) {
	req, err := http.NewRequest(http.MethodPost, TasksUrl, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var task TaskCreated

	err = json.Unmarshal(res, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *Client) CreateTask(t *TaskCreate) (*TaskCreated, error) {
	jtask, err := c.ParseTask(t)
	if err != nil {
		return nil, err
	}

	j, err := json.Marshal(jtask)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, j, "", "  ")
	return c.CreateTaskFromRawJson(string(prettyJSON.Bytes()))
}

func (c *Client) ParseTask(t *TaskCreate) (*jTaskCreate, error) {
	var jtask jTaskCreate

	jtask.Name = t.Name
	jtask.Description = t.Description

	workspaceQuery := "name==" + t.Workspace.Name + ";environment.name==" + t.Workspace.Environment
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
	jtask.WorkspaceId = workspaces[0].Id

	jtask.Parameters = t.Parameters

	jtask.Resources = t.Resources

	artifact, err := c.GetArtifacts(
		ArtifactQuery{
			Name:        t.Artifact.Name,
			WorkspaceID: workspaces[0].Id})

	for i := range artifact.Items {
		if artifact.Items[i].Workspace.Environment.Name == t.Workspace.Environment {
			jtask.Artifact.Id = artifact.Items[i].ID
			jtask.Artifact.Version = ""
			for j := range artifact.Items[i].Versions {
				if artifact.Items[i].Versions[j] > jtask.Artifact.Version {
					jtask.Artifact.Version = artifact.Items[i].Versions[j]
				}
			}
			break
		}
	}

	/*
		if t.AutoUpgradeInfo != nil {
			jtask.AutoUpgradeInfo.AutoUpgradable = t.AutoUpgradeInfo.AutoUpgradable
			jtask.AutoUpgradeInfo.OverrideWithDefaultParameters = t.AutoUpgradeInfo.OverrideWithDefaultParameters
			jtask.AutoUpgradeInfo.AutoUpgradeFailed = t.AutoUpgradeInfo.AutoUpgradeFailed
		}
	*/

	// EnvironmentId is required in create not in update
	if t.Environment != "" {
		jtask.EnvironmentId = workspaces[0].Environment.Id
	}

	return &jtask, err
}

func (c *Client) CreateTaskFromRawFile(jsonRawFile string) (*TaskCreated, error) {
	content, err := ioutil.ReadFile(jsonRawFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.CreateTaskFromRawJson(text)
}

func (c *Client) CreateTaskFromPlainFile(jsonTaskFile string) (*TaskCreated, error) {
	configFile, err := os.Open(jsonTaskFile)
	if err != nil {
		return nil, err
	}

	var taskCreate TaskCreate

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&taskCreate); err != nil {
		return nil, err
	}

	return c.CreateTask(&taskCreate)
}

func (c *Client) CreateTaskFromPlainJson(jsonTask string) (*TaskCreated, error) {
	var taskCreate TaskCreate

	err := json.Unmarshal([]byte(jsonTask), &taskCreate)
	if err != nil {
		return nil, err
	}

	return c.CreateTask(&taskCreate)
}

func (c *Client) UpdateTask(taskId string, t *TaskCreate) (*TaskCreated, error) {
	jtask, err := c.ParseTask(t)
	if err != nil {
		return nil, err
	}

	jtask.EnvironmentId = "" // EnvironmentId is required in create not in update

	j, err := json.Marshal(jtask)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, j, "", "  ")
	return c.UpdateTaskFromRawJson(taskId, string(prettyJSON.Bytes()))
}

func (c *Client) UpdateTaskFromPlainJson(taskId string, jsonTask string) (*TaskCreated, error) {
	var taskCreate TaskCreate

	err := json.Unmarshal([]byte(jsonTask), &taskCreate)
	if err != nil {
		return nil, err
	}

	return c.UpdateTask(taskId, &taskCreate)
}

func (c *Client) UpdateTaskFromRawFile(taskId string, jsonRawFile string) (*TaskCreated, error) {
	content, err := ioutil.ReadFile(jsonRawFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.UpdateTaskFromRawJson(taskId, text)
}

func (c *Client) UpdateTaskFromRawJson(taskId string, jsonRequest string) (*TaskCreated, error) {
	req, err := http.NewRequest(http.MethodPut, TasksUrl+"/"+taskId, strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var task TaskCreated

	err = json.Unmarshal(res, &task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (c *Client) DeleteTask(id string) error {
	req, err := http.NewRequest(http.MethodDelete, TasksUrl+"/"+id, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetTaskRunConfigByTaskId(taskId string) (*TaskRunConfigResponse, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(TasksUrl+"/%s/run-config", taskId), nil)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var taskrunconfig TaskRunConfig

	err = json.Unmarshal(res, &taskrunconfig)
	if err != nil {
		return nil, err
	}

	var taskRunConfigResponse TaskRunConfigResponse

	taskRunConfigResponse.TaskId = taskId
	taskRunConfigResponse.TaskRunConfig = taskrunconfig

	return &taskRunConfigResponse, nil
}

func (c *Client) DeleteTaskRunConfigByTaskId(taskId string) error {
	req, err := http.NewRequest(http.MethodDelete, TasksUrl+"/"+taskId+"/run-config", nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UpdateTaskRunConfigFromRawFile(taskId string, jsonRawFile string) (*TaskRunConfigResponse, error) {
	content, err := ioutil.ReadFile(jsonRawFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)
	return c.UpdateTaskRunConfigFromRawJson(taskId, text)
}

func (c *Client) UpdateTaskRunConfigFromRawJson(taskId string, jsonRequest string) (*TaskRunConfigResponse, error) {
	req, err := http.NewRequest(http.MethodPut, TasksUrl+"/"+taskId+"/run-config", strings.NewReader(jsonRequest))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var taskrunconfig TaskRunConfig

	err = json.Unmarshal(res, &taskrunconfig)
	if err != nil {
		return nil, err
	}

	var taskrunconfigResponse TaskRunConfigResponse
	taskrunconfigResponse.TaskId = taskId
	taskrunconfigResponse.TaskRunConfig = taskrunconfig

	return &taskrunconfigResponse, nil
}

func (c *Client) UpdateTaskRunConfigFromPlainFile(jsonTaskRunConfigFile string) (*TaskRunConfigResponse, error) {
	content, err := ioutil.ReadFile(jsonTaskRunConfigFile)
	if err != nil {
		return nil, err
	}
	// Convert []byte to string and print to screen
	text := string(content)

	return c.UpdateTaskRunConfigFromPlainJson(text)
}

func (c *Client) UpdateTaskRunConfigFromPlainJson(jsonTaskRunConfig string) (*TaskRunConfigResponse, error) {

	var taskRunConfigRequest TaskRunConfigRequest

	err := json.Unmarshal([]byte(jsonTaskRunConfig), &taskRunConfigRequest)
	if err != nil {
		return nil, err
	}

	return c.UpdateTaskRunConfig(&taskRunConfigRequest)
}

func (c *Client) UpdateTaskRunConfig(taskRunConfig *TaskRunConfigRequest) (*TaskRunConfigResponse, error) {

	workspaceQuery := "name==" + taskRunConfig.Workspace.Name + ";environment.name==" + taskRunConfig.Workspace.Environment
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
	workspaceId := workspaces[0].Id

	taskQuery := TaskQuery{Name: taskRunConfig.Name, WorkspaceId: workspaceId, Limit: 100, Offset: 0}
	tasks, err := c.GetTasks(taskQuery)
	if err != nil {
		return nil, err
	}
	if len(tasks.Items) == 0 {
		strTaskQuery, _ := json.Marshal(taskQuery)
		return nil, fmt.Errorf("task not found (%s)", strTaskQuery)
	}
	if len(tasks.Items) > 1 {
		strTaskQuery, _ := json.Marshal(taskQuery)
		return nil, fmt.Errorf("task not unique (%s)", strTaskQuery)
	}
	taskId := tasks.Items[0].Executable

	j, err := json.Marshal(taskRunConfig.TaskRunConfig)
	if err != nil {
		return nil, err
	}

	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, j, "", "  ")

	return c.UpdateTaskRunConfigFromRawJson(taskId, string(prettyJSON.Bytes()))
}
