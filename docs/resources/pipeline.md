---
page_title: "gocd_pipeline Resource - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Resource `gocd_pipeline`



## Example Usage

```terraform
resource "gocd_pipeline" "test-pipeline3-upstream" {
  name           = "test-pipeline3-upstream"
  group          = "testing"
  label_template = "$${COUNT}"

  materials = [{
    type = "git"

    attributes = {
      url    = "https://github.com/beamly/terraform-provider-gocd.git"
      branch = "master"
    }
  }]

  stages = ["${data.gocd_stage_definition.test.json}"]
}

resource "gocd_pipeline" "test-pipeline3" {
  name           = "test-pipeline3"
  group          = "testing"
  label_template = "$${COUNT}"

  materials = [{
    type = "git"

    attributes = {
      url    = "https://github.com/beamly/terraform-provider-gocd.git"
      branch = "master"
    }
  }]

  stages = ["${data.gocd_stage_definition.test.json}"]
}

# CMD terraform import gocd_pipeline_stage.test "test"
data "gocd_stage_definition" "test" {
  name            = "test"
  fetch_materials = true

  jobs = [
    "${data.gocd_job_definition.test.json}",
  ]
}

data "gocd_job_definition" "test" {
  name = "test"

  tasks = [
    "${data.gocd_task_definition.test-pipeline3_test_test_1.json}",
  ]
}

data "gocd_task_definition" "test-pipeline3_test_test_1" {
  type    = "exec"
  run_if  = ["passed"]
  command = "echo"

  arguments = [
    "test",
  ]
}
```

## Schema

### Required

- **group** (String)
- **materials** (Block List, Min: 1) (see [below for nested schema](#nestedblock--materials))
- **name** (String)

### Optional

- **enable_pipeline_locking** (Boolean, Deprecated)
- **environment_variables** (Block List) (see [below for nested schema](#nestedblock--environment_variables))
- **id** (String) The ID of this resource.
- **label_template** (String)
- **lock_behavior** (String)
- **parameters** (Map of String)
- **stages** (List of String)
- **template** (String)

### Read-only

- **version** (String)

<a id="nestedblock--materials"></a>
### Nested Schema for `materials`

Required:

- **attributes** (Block List, Min: 1, Max: 1) (see [below for nested schema](#nestedblock--materials--attributes))

Optional:

- **type** (String)

<a id="nestedblock--materials--attributes"></a>
### Nested Schema for `materials.attributes`

Optional:

- **auto_update** (Boolean)
- **branch** (String)
- **destination** (String)
- **filter** (List of String)
- **invert_filter** (Boolean)
- **name** (String)
- **pipeline** (String)
- **shallow_clone** (Boolean)
- **stage** (String)
- **submodule_folder** (String)
- **url** (String)



<a id="nestedblock--environment_variables"></a>
### Nested Schema for `environment_variables`

Required:

- **name** (String)

Optional:

- **encrypted_value** (String)
- **secure** (Boolean)
- **value** (String)


