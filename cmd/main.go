package main

import (
	"encoding/json"
	"fmt"

	"github.com/bartsimp/talend-rest-go/client"
	"github.com/bartsimp/talend-rest-go/client/artifacts"
	"github.com/bartsimp/talend-rest-go/client/environments"
	"github.com/bartsimp/talend-rest-go/client/plans_executables"
	"github.com/bartsimp/talend-rest-go/client/runtime_engines"
	"github.com/bartsimp/talend-rest-go/client/tasks"
	"github.com/bartsimp/talend-rest-go/client/workspaces"
	"github.com/bartsimp/talend-rest-go/models"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

const TMC_API_HOST = "api.eu.cloud.talend.com"
const TMC_API_SCHEMA = "https"

func main() {
	apiKey := "my-api-key"

	c := client.NewHTTPClientWithConfig(
		strfmt.Default,
		client.DefaultTransportConfig().
			WithHost(TMC_API_HOST).
			WithSchemes([]string{TMC_API_SCHEMA}),
	)

	authInfo := runtime.ClientAuthInfoWriterFunc(
		func(cr runtime.ClientRequest, r strfmt.Registry) error {
			cr.SetHeaderParam("Authorization", fmt.Sprint("Bearer ", apiKey))
			return nil
		},
	)

	getEnvironments(c, authInfo, "")

	envName := "myEnvironment"
	wsName := "myWorkspace"
	environment := createEnvironment(c, authInfo, &models.CreateEnvironmentRequest{
		Name:          &envName,
		WorkspaceName: &wsName,
	})
	environmentID := *environment.ID

	getEnvironments(c, authInfo, "")
	workspaces := getWorkspaces(c, authInfo, fmt.Sprintf("environment.id==%s", environmentID))
	// getWorkspaces(c, authInfo, fmt.Sprintf("environment.name==%s", envName))
	var workspaceID string
	for _, workspace := range workspaces {
		if *workspace.Name == wsName && *workspace.Environment.Name == envName {
			workspaceID = *workspace.ID
		}

	}
	updateWorkspace(c, authInfo, workspaceID, &models.UpdateWorkspaceRequest{
		Name:        &wsName,
		Description: "myWorkspaceDescription",
	})
	getWorkspaces(c, authInfo, fmt.Sprintf("environment.id==%s", environmentID))

	wsName = "ws_tmp"
	ws2 := newWorkspace(c, authInfo, &models.CreateWorkspaceRequest{
		EnvironmentID: &environmentID,
		Name:          &wsName,
		Description:   "desc_tmp",
	})
	getWorkspaces(c, authInfo, fmt.Sprintf("environment.id==%s", environmentID))
	updateWorkspace(c, authInfo, ws2.ID, &models.UpdateWorkspaceRequest{
		Name:        &wsName,
		Description: "desc_tmp_updated",
	})
	getWorkspaces(c, authInfo, fmt.Sprintf("environment.id==%s", environmentID))
	deleteWorkspace(c, authInfo, ws2.ID)
	getWorkspaces(c, authInfo, fmt.Sprintf("environment.id==%s", environmentID))

	getAvailablePlans(c, authInfo, environmentID)
	getAvailableTasks(c, authInfo, environmentID)
	getArtifact(c, authInfo, "artifactID")
	getArtifactsAvailable(c, authInfo, environmentID, workspaceID)

	taskName := "myTaskName"
	taskDescription := "myTaskDescription"
	artifactId := ""
	artifactVersion := ""
	createTask(c, authInfo, &models.TaskV21CreateRequest{
		WorkspaceID:   &workspaceID,
		Name:          &taskName,
		Description:   &taskDescription,
		EnvironmentID: &environmentID,
		Artifact: &models.ArtifactRequest{
			ID:      &artifactId,
			Version: &artifactVersion,
		},
	})

	getRemoteEngines(c, authInfo, "")
	reName := "myRemoteEngine"
	re := createRemoteEngine(c, authInfo, &models.EngineRequest{
		Name:          &reName,
		EnvironmentID: &environmentID,
	})
	getRemoteEngines(c, authInfo, "")
	deleteRemoteEngine(c, authInfo, *re.ID)
	getRemoteEngines(c, authInfo, "")

	deleteEnvironment(c, authInfo, environmentID)
	getEnvironments(c, authInfo, "")

	// getResources(c, authInfo)

}

func getAvailablePlans(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, environmentID string) {
	fmt.Println("getAvailablePlans")
	getAvailablePlans, err := c.PlansExecutables.GetAvailablePlans(
		plans_executables.NewGetAvailablePlansParams().WithEnvironmentID(&environmentID),
		authInfo,
	)
	if err != nil {
		// switch err.(type) {
		// case *plans_executables.GetAvailablePlansBadRequest:
		// 	json, err := json.Marshal(err.(*plans_executables.GetAvailablePlansBadRequest).Payload)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Printf("%s\n", json)
		// }
		switch err := err.(type) {
		case *plans_executables.GetAvailablePlansBadRequest:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return

	}
	page := getAvailablePlans.GetPayload()
	json, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)

}

func getAvailableTasks(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, environmentID string) {
	fmt.Println("getAvailableTasks")
	getAvailablePlans, err := c.Tasks.GetAvailableTasks(
		tasks.NewGetAvailableTasksParams().WithEnvironmentID(&environmentID),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *tasks.GetAvailableTasksBadRequest:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return

	}
	page := getAvailablePlans.GetPayload()
	json, err := json.Marshal(page)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)

}

func createTask(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, body *models.TaskV21CreateRequest) *models.TaskV21 {
	fmt.Println("createTask")
	createTaskCreated, err := c.Tasks.CreateTask(
		tasks.NewCreateTaskParams().WithBody(body),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *tasks.CreateTaskBadRequest:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return nil

	}
	task := createTaskCreated.GetPayload()
	json, err := json.Marshal(task)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
	return task
}

func getArtifact(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, artifactID string) *models.Artifact {
	fmt.Println("getArtifact")
	getArtifact, err := c.Artifacts.GetArtifact(
		artifacts.NewGetArtifactParams().WithID(artifactID),
		authInfo,
	)
	if err != nil {
		// switch err.(type) {
		// case *artifacts.GetArtifactBadRequest:
		// 	json, err := json.Marshal(err.(*artifacts.GetArtifactBadRequest).Payload)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	fmt.Printf("%s\n", json)
		// }
		switch err := err.(type) {
		case *artifacts.GetArtifactBadRequest:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return nil
	}
	artifact := getArtifact.GetPayload()
	json, err := json.Marshal(artifact)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)

	return artifact
}

func getArtifactsAvailable(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, environmentID string, workspaceID string) {
	fmt.Println("getArtifactsAvailable")
	getArtifact, err := c.Artifacts.GetArtifactsAvailable(
		artifacts.NewGetArtifactsAvailableParams().
			WithEnvironmentID(&environmentID).
			WithWorkspaceID(&workspaceID),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *artifacts.GetArtifactsAvailableBadRequest:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return
	}
	artifact := getArtifact.GetPayload()
	json, err := json.Marshal(artifact)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}

func getRemoteEngines(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, query string) {
	fmt.Println("getRemoteEngines", "query", query)
	remoteEngines, err := c.RuntimeEngines.GetRemoteEngines(
		runtime_engines.NewGetRemoteEnginesParams().WithQuery(&query),
		authInfo,
	)
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(remoteEngines.GetPayload(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)
}

func createRemoteEngine(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, body *models.EngineRequest) *models.Engine {
	fmt.Println("createRemoteEngine")
	remoteEngines, err := c.RuntimeEngines.CreateRemoteEngine(
		runtime_engines.NewCreateRemoteEngineParams().
			WithBody(body),
		authInfo,
	)
	if err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(remoteEngines.GetPayload(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)

	return remoteEngines.GetPayload()
}

func deleteRemoteEngine(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, engineID string) {
	fmt.Println("createRemoteEngine")
	_, err := c.RuntimeEngines.DeleteRemoteEngine(
		runtime_engines.NewDeleteRemoteEngineParams().WithID(engineID),
		authInfo,
	)
	if err != nil {
		panic(err)
	}
}

func createEnvironment(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, body *models.CreateEnvironmentRequest) *models.EnvironmentInfo {
	fmt.Println("createEnvironment")
	environment, err := c.Environments.CreateEnvironment(
		environments.NewCreateEnvironmentParams().WithBody(body),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *environments.CreateEnvironmentBadRequest:
			fmt.Printf("%v\n", err.Payload)
		case *environments.CreateEnvironmentConflict:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return nil
	}

	json, err := json.MarshalIndent(environment.GetPayload(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", json)

	return environment.GetPayload()
}

func getEnvironments(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, query string) []*models.EnvironmentInfo {
	fmt.Println("getEnvironments", "query", query)
	environments, err := c.Environments.GetEnvironments(
		environments.NewGetEnvironmentsParams().WithQuery(&query),
		authInfo,
	)
	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(environments.GetPayload(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)

	return environments.GetPayload()
}

func deleteEnvironment(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, environmentID string) {
	fmt.Println("deleteEnvironment")
	_, err := c.Environments.DeleteEnvironment(
		environments.NewDeleteEnvironmentParams().WithEnvironmentID(environmentID),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *environments.CreateEnvironmentBadRequest:
			fmt.Printf("%v\n", err.Payload)
		case *environments.CreateEnvironmentConflict:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
	}
}

func getWorkspaces(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, query string) []*models.Workspaceinfo {
	fmt.Println("getWorkspaces", "query", query)
	getWorkspacesOK, err := c.Workspaces.GetWorkspaces(
		workspaces.NewGetWorkspacesParams().WithQuery(&query),
		authInfo,
	)
	if err != nil {
		panic(err)
	}

	bytes, err := json.MarshalIndent(getWorkspacesOK.GetPayload(), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", bytes)

	return getWorkspacesOK.GetPayload()
}

func newWorkspace(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, body *models.CreateWorkspaceRequest) *models.Workspace {
	fmt.Println("newWorkspace")
	createCustomWorkspaceOK, createCustomWorkspaceCreated, err := c.Workspaces.CreateCustomWorkspace(
		workspaces.NewCreateCustomWorkspaceParams().
			WithBody(body),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *workspaces.CreateCustomWorkspaceBadRequest:
			fmt.Printf("%v\n", err.Payload)
		case *workspaces.CreateCustomWorkspaceConflict:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
		return nil
	}
	fmt.Println(createCustomWorkspaceOK)
	fmt.Println(createCustomWorkspaceCreated)
	return createCustomWorkspaceCreated.GetPayload()
}

func updateWorkspace(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, workspaceID string, body *models.UpdateWorkspaceRequest) {
	fmt.Println("updateWorkspace")
	updateCustomWorkspaceNoContent, err := c.Workspaces.UpdateCustomWorkspace(
		workspaces.NewUpdateCustomWorkspaceParams().
			WithWorkspaceID(workspaceID).
			WithBody(body),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *workspaces.UpdateCustomWorkspaceBadRequest:
			fmt.Printf("%v\n", err.Payload)
		default:
			fmt.Printf("%s\n", err)
		}
	} else {
		fmt.Println(updateCustomWorkspaceNoContent)
	}
}

func deleteWorkspace(c *client.TalendManagementConsolePublicAPI, authInfo runtime.ClientAuthInfoWriterFunc, workspaceID string) {
	fmt.Println("deleteWorkspace")
	err := c.Workspaces.DeleteWorkspace(
		workspaces.NewDeleteWorkspaceParams().WithWorkspaceID(workspaceID),
		authInfo,
	)
	if err != nil {
		switch err := err.(type) {
		case *workspaces.UpdateCustomWorkspaceBadRequest:
			fmt.Printf("%v\n", err.Payload)
		// case *runtime.APIError:
		// 	resp := err.Response.(http.Response)
		// 	defer resp.Body.Close()
		// 	body, err := ioutil.ReadAll(resp.Body)
		// 	fmt.Println("%s\n", body)
		default:
			fmt.Printf("%s\n", err)
		}
	}
}
