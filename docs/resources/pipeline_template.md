---
page_title: "gocd_pipeline_template Resource - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Resource `gocd_pipeline_template`



## Example Usage

```terraform
resource "gocd_pipeline_template" "test-pipeline" {
  name   = "template0-terraform"
  stages = ["${data.gocd_stage_definition.test-stage.json}"]
}

data "gocd_stage_definition" "test-stage" {
  name = "test"
  jobs = [
    "${data.gocd_job_definition.test-job.json}"
  ]
}

data "gocd_job_definition" "test-job" {
  name = "test"
  tasks = [
    "${data.gocd_task_definition.test.json}"
  ]
}
data "gocd_task_definition" "test" {
  type    = "exec"
  command = "echo"
  arguments = [
    "hello",
    "world",
  ]
}
```

## Schema

### Required

- **name** (String)
- **stages** (List of String)

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **version** (String)


