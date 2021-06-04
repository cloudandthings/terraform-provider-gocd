package gocd

import (
	"context"
	"github.com/hashicorp/go-version"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testPipelineServiceUnPause(t *testing.T) {
	if runIntegrationTest(t) {

		ctx := context.Background()
		pipelineName := "test-pipeline-un-pause"

		stages := buildMockPipelineStages()
		mockPipeline := &Pipeline{
			Name:                 pipelineName,
			Group:                mockTestingGroup,
			LabelTemplate:        "${COUNT}",
			LockBehavior:         "none",
			Parameters:           make([]*Parameter, 0),
			EnvironmentVariables: make([]*EnvironmentVariable, 0),
			Materials: []Material{{
				Type: "git",
				Attributes: &MaterialAttributesGit{
					URL:         "git@github.com:sample_repo/example.git",
					Destination: "dest",
					Branch:      "master",
					AutoUpdate:  true,
				},
			}},
			Stages: stages,
		}

		pausePipeline, _, err := intClient.PipelineConfigs.Create(ctx, mockTestingGroup, mockPipeline)
		assert.NoError(t, err)
		pausePipeline.Links = nil
		pausePipeline.Version = ""

		// Make sure version-specific defaults are properly set
		apiVersion, err := client.getAPIVersion(ctx, "admin/pipelines/:pipeline_name")
		assert.NoError(t, err)
		switch apiVersion {
		case apiV6, apiV7, apiV8, apiV9, apiV10, apiV11:
			mockPipeline.Origin = &PipelineConfigOrigin{Type: "gocd"}
			fallthrough
		case apiV5:
			mockPipeline.LockBehavior = "none"
		}

		apiVersion, err = client.getAPIVersion(ctx, "pipelines/:pipeline_name/unlock")
		assert.NoError(t, err)
		releaseLockErrorMessage := "Received HTTP Status '406 Not Acceptable'"
		switch apiVersion {
		case apiV1:
			releaseLockErrorMessage = "Received HTTP Status '409 Conflict': {\n  \"message\": \"No lock exists within the pipeline configuration for test-pipeline-un-pause\"\n}"
		}

		assert.Equal(t, mockPipeline, pausePipeline)

		// From 18.8.0 onwards pipelines are no-longer created paused
		v, _, err := client.ServerVersion.Get(context.Background())

		pausedBeforeVersion, _ := version.NewVersion("18.8.0")
		if v.VersionParts.LessThan(pausedBeforeVersion) {
			pp, _, err := intClient.Pipelines.Unpause(ctx, pipelineName)
			assert.NoError(t, err)
			assert.True(t, pp)
		}

		pp, _, err := intClient.Pipelines.Pause(ctx, pipelineName)
		assert.NoError(t, err)
		assert.True(t, pp)

		pp, _, err = intClient.Pipelines.Unpause(ctx, pipelineName)
		assert.NoError(t, err)
		assert.True(t, pp)

		pp, _, err = intClient.Pipelines.ReleaseLock(ctx, pipelineName)
		assert.EqualError(t, err, releaseLockErrorMessage)
		assert.False(t, pp)

		deleteResponse, _, err := intClient.PipelineConfigs.Delete(ctx, pipelineName)
		assert.Contains(t, deleteResponse, "'test-pipeline-un-pause' was deleted successfully")
	}

}
