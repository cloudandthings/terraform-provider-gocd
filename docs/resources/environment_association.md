---
page_title: "gocd_environment_association Resource - terraform-provider-gocd"
subcategory: ""
description: |-
  
---

# Resource `gocd_environment_association`



## Example Usage

```terraform
resource "gocd_environment" "test-environment" {
  name = "test-environment"
}
```

## Schema

### Required

- **environment** (String)
- **pipeline** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-only

- **version** (String)


