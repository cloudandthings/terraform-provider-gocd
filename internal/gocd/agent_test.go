package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAgent(t *testing.T) {
	setup()
	defer teardown()

	t.Run("JobRunHistory", testAgentJobRunHistory)
	t.Run("BulkUpdate", testAgentBulkUpdate)
	t.Run("Delete", testAgentDelete)
	t.Run("Get", testAgentGet)
	t.Run("Update", testAgentUpdate)
	t.Run("RemoveLinks", testAgentRemoveLinks)
	t.Run("List", testAgentList)
}

func testAgentJobRunHistory(t *testing.T) {
	mux.HandleFunc("/api/agents/testAgentJobRunHistory-e6d3-4299-9120-7faff6e6030b/job_run_history", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		bdy, err := ioutil.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Empty(t, string(bdy))

		j, _ := ioutil.ReadFile("test/resources/agents-job-history.0.json")
		fmt.Fprint(w, string(j))
	})
	jobs, _, err := client.Agents.JobRunHistory(context.Background(), "testAgentJobRunHistory-e6d3-4299-9120-7faff6e6030b")
	assert.Nil(t, err)
	assert.NotEmpty(t, jobs)
	assert.Len(t, jobs, 1)

	job := jobs[0]
	assert.Equal(t, "5c5c318f-e6d3-4299-9120-7faff6e6030b", job.AgentUUID)
	assert.Equal(t, "upload", job.Name)
	assert.Equal(t, 1435631497131, job.ScheduledDate)
	assert.Empty(t, job.OriginalJobID)
	assert.Equal(t, 251, job.PipelineCounter)
	assert.Equal(t, false, job.Rerun)
	assert.Equal(t, "distributions-all", job.PipelineName)
	assert.Equal(t, "Passed", job.Result)
	assert.Equal(t, "Completed", job.State)
	assert.Equal(t, 100129, job.ID)
	assert.Equal(t, "1", job.StageCounter)
	assert.Equal(t, "upload-installers", job.StageName)
	assert.Len(t, job.JobStateTransitions, 1)

	transition := job.JobStateTransitions[0]
	assert.Equal(t, 1435631497131, transition.StateChangeTime)
	assert.Equal(t, 539906, transition.ID)
	assert.Equal(t, JobStateTransitionScheduled, transition.State)

}

func testAgentBulkUpdate(t *testing.T) {

	mux.HandleFunc("/api/agents", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "PATCH", "Unexpected HTTP method")
		bdy, err := ioutil.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, `{
  "uuids": [
    "adb9540a-b954-4571-9d9b-2f330739d4da"
  ],
  "operations": {
    "environments": {
      "add": [
        "new-env",
        "new-env1"
      ]
    }
  }
}
`, string(bdy))

		j, _ := ioutil.ReadFile("test/resources/agents.2.json")
		fmt.Fprint(w, string(j))
	})
	bulkUpdate := AgentBulkUpdate{
		Uuids: []string{"adb9540a-b954-4571-9d9b-2f330739d4da"},
		Operations: &AgentBulkOperationsUpdate{
			Environments: &AgentBulkOperationUpdate{
				Add: []string{"new-env", "new-env1"},
			},
		},
	}
	message, _, err := client.Agents.BulkUpdate(context.Background(), bulkUpdate)
	assert.Nil(t, err)
	assert.NotEmpty(t, message)
	assert.Equal(t, "Updated agent(s) with uuid(s): [adb9540a-b954-4571-9d9b-2f330739d4da, adb528b2-b954-1234-9d9b-b27ag4h568e1].", message)
}

func testAgentDelete(t *testing.T) {

	mux.HandleFunc("/api/agents/testAgentDelete-b954-4571-9d9b-2f330739d4da", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "DELETE", "Unexpected HTTP method")
		fmt.Fprint(w, `{"message":"Deleted this resource"}`)
	})

	message, _, err := client.Agents.Delete(context.Background(), "testAgentDelete-b954-4571-9d9b-2f330739d4da")
	assert.Nil(t, err)
	assert.NotEmpty(t, message)

	assert.Equal(t, message, "Deleted this resource")
}

func testAgentRemoveLinks(t *testing.T) {
	a := Agent{
		Links: &HALLinks{},
	}

	assert.NotNil(t, a.Links)
	a.RemoveLinks()
	assert.Nil(t, a.Links)
}

func testAgentUpdate(t *testing.T) {

	mux.HandleFunc("/api/agents/testAgentUpdate-b954-4571-9d9b-2f330739d4da", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "PATCH", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/agent.1.json")
		fmt.Fprint(w, string(j))
	})

	agentUpdate := Agent{
		Resources: []string{"other"},
	}
	agent, _, err := client.Agents.Update(context.Background(), "testAgentUpdate-b954-4571-9d9b-2f330739d4da", &agentUpdate)
	assert.Nil(t, err)
	assert.NotNil(t, *agent)

	assert.Equal(t, agent.Resources[0], "other")
}

func testAgentGet(t *testing.T) {

	mux.HandleFunc("/api/agents/testAgentGet-b954-4571-9d9b-2f330739d4da", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/agent.0.json")
		fmt.Fprint(w, string(j))
	})

	agent, _, err := client.Agents.Get(context.Background(), "testAgentGet-b954-4571-9d9b-2f330739d4da")
	assert.Nil(t, err)

	for _, attribute := range []EqualityTest{
		{agent.BuildDetails.Links.Get("Job").URL.String(), "https://ci.example.com/go/tab/build/detail/up42/1/up42_stage/1/up42_job"},
		{agent.BuildDetails.Links.Get("Stage").URL.String(), "https://ci.example.com/go/pipelines/up42/1/up42_stage/1"},
		{agent.BuildDetails.Links.Get("Pipeline").URL.String(), "https://ci.example.com/go/tab/pipeline/history/up42"},
	} {
		assert.Equal(t, attribute.wanted, attribute.got)
	}

	assert.NotNil(t, agent.BuildDetails)
	testAgent(t, agent)
}

func testAgentList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/api/agents", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		j, _ := ioutil.ReadFile("test/resources/agents.1.json")
		fmt.Fprint(w, string(j))
	})

	agents, _, err := client.Agents.List(context.Background())

	assert.Nil(t, err)
	assert.Len(t, agents, 1)

	testAgent(t, agents[0])
}

func testAgent(t *testing.T, agent *Agent) {

	for _, attribute := range []EqualityTest{
		{agent.Links.Get("Self").URL.String(), "https://ci.example.com/go/api/agents/adb9540a-b954-4571-9d9b-2f330739d4da"},
		{agent.Links.Get("Doc").URL.String(), "https://api.gocd.org/#agents"},
		{agent.Links.Get("Find").URL.String(), "https://ci.example.com/go/api/agents/:uuid"},
		{agent.UUID, "adb9540a-b954-4571-9d9b-2f330739d4da"},
		{agent.Hostname, "agent01.example.com"},
		{agent.IPAddress, "10.12.20.47"},
		{agent.Sandbox, "/Users/ketanpadegaonkar/projects/gocd/gocd/agent"},
		{agent.OperatingSystem, "Mac OS X"},
		{agent.AgentConfigState, "Enabled"},
		{agent.AgentState, "Idle"},
		{agent.Resources[0], "java"},
		{agent.Resources[1], "linux"},
		{agent.Resources[2], "firefox"},
		{agent.Environments[0], "perf"},
		{agent.Environments[1], "UAT"},
		{agent.BuildState, "Idle"},
	} {
		assert.Equal(t, attribute.wanted, attribute.got)
	}

	assert.Equal(t, 84983328768, agent.FreeSpace)
}
