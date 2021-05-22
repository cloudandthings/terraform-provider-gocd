---
page_title: "gocd_stage_definition Data Source - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Data Source `gocd_stage_definition`



## Example Usage

```terraform
data "gocd_task_definition" "test" {
  type = "exec"

  run_if = [
    "passed",
  ]

  command = "/usr/local/bin/terraform"

  arguments = [
    "-debug",
    "version",
  ]

  working_directory = "tmp/"
}

data "gocd_job_definition" "test" {
  name = "job-name"

  tasks = [
    "${data.gocd_task_definition.test.json}",
  ]
}

data "gocd_stage_definition" "test" {
  name = "stage-name"

  jobs = [
    "${data.gocd_job_definition.test.json}",
  ]

  approval {
    type = "manual"
  }
}
```

## Schema

### Required

- **jobs** (List of String)
- **name** (String)

### Optional

- **approval** (Block Set, Max: 1) (see [below for nested schema](#nestedblock--approval))
- **clean_working_directory** (Boolean)
- **environment_variables** (Block List) (see [below for nested schema](#nestedblock--environment_variables))
- **fetch_materials** (Boolean)
- **id** (String) The ID of this resource.
- **never_cleanup_artifacts** (Boolean)
- **pipeline** (String)
- **pipeline_template** (String)

### Read-only

- **json** (String)

<a id="nestedblock--approval"></a>
### Nested Schema for `approval`

Required:

- **type** (String)

Optional:

- **authorization** (Block Set, Max: 1) (see [below for nested schema](#nestedblock--approval--authorization))

<a id="nestedblock--approval--authorization"></a>
### Nested Schema for `approval.authorization`

Optional:

- **roles** (Set of String)
- **users** (Set of String)



<a id="nestedblock--environment_variables"></a>
### Nested Schema for `environment_variables`

Required:

- **name** (String)

Optional:

- **encrypted_value** (String)
- **secure** (Boolean)
- **value** (String)


