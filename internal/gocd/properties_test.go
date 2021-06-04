package gocd

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestProperties(t *testing.T) {
	setup()
	defer teardown()

	t.Run("List", testPropertiesList)
	t.Run("Get", testPropertiesGet)
	t.Run("ListHistorical", testPropertiesListHistorical)
	t.Run("Create", testPropertiesCreate)
	//t.Run("Pipelines", testPipelineTemplatePipelines)
	//t.Run("Update", testPipelineTemplateUpdate)
	//t.Run("StageContainerI", testPipelineTemplateStageContainer)
}

func testPropertiesList(t *testing.T) {
	mux.HandleFunc("/properties/test-pipeline/5/test-stage/3/test-job", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")

		j, _ := ioutil.ReadFile("test/resources/properties.0.csv")
		fmt.Fprint(w, string(j))
	})

	p, _, err := client.Properties.List(context.Background(), &PropertyRequest{
		Pipeline:        "test-pipeline",
		PipelineCounter: 5,
		Stage:           "test-stage",
		StageCounter:    3,
		Job:             "test-job",
	})

	assert.Nil(t, err)
	assert.Equal(t, []string{"cruise_agent", "cruise_timestamp_01_scheduled", "cruise_timestamp_02_assigned"}, p.Header)
	assert.Equal(t, []string{"myLocalAgent", "2015-07-09T11:59:08+05:30", "2015-07-09T11:59:16+05:30"}, p.DataFrame[0])
}

func testPropertiesGet(t *testing.T) {
	mux.HandleFunc("/properties/test-pipeline/5/test-stage/3/test-job/cruise_agent", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")
		fmt.Fprint(w, `cruise_agent
myLocalAgent`)
	})

	p, _, err := client.Properties.Get(context.Background(), "cruise_agent", &PropertyRequest{
		Pipeline:        "test-pipeline",
		PipelineCounter: 5,
		Stage:           "test-stage",
		StageCounter:    3,
		Job:             "test-job",
	})

	assert.Nil(t, err)
	assert.Equal(t, []string{"cruise_agent"}, p.Header)
	assert.Equal(t, []string{"myLocalAgent"}, p.DataFrame[0])
}

func testPropertiesListHistorical(t *testing.T) {
	mux.HandleFunc("/properties/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "GET", "Unexpected HTTP method")

		q := r.URL.Query()
		assert.Equal(t, "PipelineName", q.Get("pipelineName"))
		assert.Equal(t, "StageName", q.Get("stageName"))
		assert.Equal(t, "JobName", q.Get("jobName"))
		assert.Equal(t, "latest", q.Get("limitPipeline"))
		assert.Equal(t, "2", q.Get("limitCount"))

		j, _ := ioutil.ReadFile("test/resources/properties.1.csv")
		fmt.Fprint(w, string(j))
	})
	p, _, err := client.Properties.ListHistorical(context.Background(), &PropertyRequest{
		Pipeline:      "PipelineName",
		Stage:         "StageName",
		Job:           "JobName",
		LimitPipeline: "latest",
		Limit:         2,
	})

	assert.Nil(t, err)
	assert.Equal(t, []string{"cruise_agent", "cruise_job_duration", "cruise_job_id", "cruise_job_result", "cruise_pipeline_counter", "cruise_pipeline_label", "cruise_stage_counter", "cruise_timestamp_01_scheduled", "cruise_timestamp_02_assigned", "cruise_timestamp_03_preparing", "cruise_timestamp_04_building", "cruise_timestamp_05_completing", "cruise_timestamp_06_completed"}, p.Header)
	assert.Equal(t, []string{"myLocalAgent", "0", "13", "Passed", "8", "4f9e580347b2e259fe030a775771359cdc984346", "1", "2015-07-07T09:44:27+05:30", "2015-07-07T09:44:34+05:30", "2015-07-07T09:44:44+05:30", "2015-07-07T09:44:46+05:30", "2015-07-07T09:44:46+05:30", "2015-07-07T09:44:46+05:30"}, p.DataFrame[0])
	assert.Equal(t, []string{"myLocalAgent", "0", "14", "Passed", "9", "4f9e580347b2e259fe030a775771359cdc984346", "1", "2015-07-07T10:17:37+05:30", "2015-07-07T10:17:45+05:30", "2015-07-07T10:17:55+05:30", "2015-07-07T10:17:56+05:30", "2015-07-07T10:17:56+05:30", "2015-07-07T10:17:56+05:30"}, p.DataFrame[1])
}

func testPropertiesCreate(t *testing.T) {
	mux.HandleFunc("/properties/PipelineName/541/StageName/1/JobName/PropertyName", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, "POST", "Unexpected HTTP method")
		fmt.Fprint(w, "Property 'PropertyName' created with value 'PropertyValue'")
	})

	r, _, err := client.Properties.Create(context.Background(), "PropertyName", "PropertyValue", &PropertyRequest{
		Pipeline:        "PipelineName",
		PipelineCounter: 541,
		Stage:           "StageName",
		StageCounter:    1,
		Job:             "JobName",
	})

	assert.Nil(t, err)
	assert.True(t, r)
}
