---
page_title: "gocd_task_definition Data Source - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Data Source `gocd_task_definition`



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
```

## Schema

### Required

- **type** (String)

### Optional

- **arguments** (List of String)
- **artifact_origin** (String)
- **build_file** (String)
- **command** (String)
- **configuration** (List of Map of String)
- **destination** (String)
- **id** (String) The ID of this resource.
- **is_source_a_file** (Boolean)
- **job** (String)
- **nant_path** (String)
- **pipeline** (String)
- **plugin_id** (String)
- **plugin_version** (String)
- **run_if** (List of String)
- **source** (String)
- **stage** (String)
- **target** (String)
- **working_directory** (String)

### Read-only

- **json** (String)


