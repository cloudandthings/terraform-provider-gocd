---
page_title: "gocd_environment Resource - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Resource `gocd_environment`



## Example Usage

```terraform
resource "gocd_environment" "test-environment" {
  name = "test-environment"
}

resource "gocd_pipeline" "test-pipeline" {
  name  = "test-pipeline"
  group = "test-group"

  materials = [
    {
      type = "git"

      attributes = {
        name   = "gocd-src"
        url    = "git@github.com:gocd/gocd"
        branch = "master"
      }
    },
  ]

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


resource "gocd_environment_association" "test-environment-association" {
  environment = "${gocd_environment.test-environment.name}"
  pipeline    = "${gocd_pipeline.test-pipeline.name}"
}
```

## Schema

### Required

- **name** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **version** (String)


