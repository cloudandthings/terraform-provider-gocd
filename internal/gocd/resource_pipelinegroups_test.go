package gocd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testResourcePipelineGroups(t *testing.T) {
	pgs := PipelineGroups{}
	pg1 := &PipelineGroup{
		Name: "test-group1",
		Pipelines: []*Pipeline{
			{Name: "pipeline1"},
			{Name: "pipeline2"},
		},
	}
	pg2 := &PipelineGroup{
		Name: "test-group2",
		Pipelines: []*Pipeline{
			{Name: "pipeline3"},
			{Name: "pipeline4"},
		},
	}
	pgs = append(pgs, pg1)
	pgs = append(pgs, pg2)

	p := pgs.GetGroupByPipelineName("pipeline3")
	assert.Equal(t, "test-group2", p.Name)

	p = pgs.GetGroupByPipelineName("pipeline5")
	assert.Nil(t, p)

	p = pgs.GetGroupByPipeline(&Pipeline{
		Name: "pipeline1",
	})
	assert.Equal(t, "test-group1", p.Name)
}
