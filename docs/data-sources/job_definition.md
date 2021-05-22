---
page_title: "gocd_job_definition Data Source - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Data Source `gocd_job_definition`



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
  name               = "job-name"
  run_instance_count = 3
  timeout            = 6

  environment_variables {
    name  = "USERNAME"
    value = "myusername"
  }

  environment_variables = [
    {
      name            = "PASSWORD"
      encrypted_value = "$R*YN:LDFIOH"
    },
    {
      name   = "HIDDEN"
      value  = "shown"
      secure = true
    },
  ]

  resources = [
    "alpha",
    "beta",
  ]

  tasks = [
    "${data.gocd_task_definition.test.json}",
  ]

  tabs = [
    {
      name = "Report"
      path = "report1.html"
    },
  ]

  artifacts = [
    {
      type   = "build"
      source = "web.war"
    },
  ]

  properties {
    name   = "coverage.class"
    source = "target/emma/coverage.xml"
    xpath  = "substring-before(//report/data/all/coverage[starts-with(@type,'class')]/@value, '%')"
  }
}
```

## Schema

### Required

- **name** (String)
- **tasks** (List of String)

### Optional

- **artifacts** (Block List) (see [below for nested schema](#nestedblock--artifacts))
- **elastic_profile_id** (String)
- **environment_variables** (Block List) (see [below for nested schema](#nestedblock--environment_variables))
- **id** (String) The ID of this resource.
- **properties** (Block List) (see [below for nested schema](#nestedblock--properties))
- **resources** (Set of String)
- **run_instance_count** (Number)
- **tabs** (Block List) (see [below for nested schema](#nestedblock--tabs))
- **timeout** (Number)

### Read-only

- **json** (String)

<a id="nestedblock--artifacts"></a>
### Nested Schema for `artifacts`

Required:

- **source** (String)
- **type** (String)

Optional:

- **destination** (String)


<a id="nestedblock--environment_variables"></a>
### Nested Schema for `environment_variables`

Required:

- **name** (String)

Optional:

- **encrypted_value** (String)
- **secure** (Boolean)
- **value** (String)


<a id="nestedblock--properties"></a>
### Nested Schema for `properties`

Required:

- **name** (String)
- **source** (String)
- **xpath** (String)


<a id="nestedblock--tabs"></a>
### Nested Schema for `tabs`

Required:

- **name** (String)
- **path** (String)


