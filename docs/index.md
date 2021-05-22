---
page_title: "GoCD Provider"
subcategory: ""
description: |-
  
---

# GoCD Provider



## Example Usage

```terraform
provider "gocd" {
  # example configuration here
}
```

## Authentication

The GoCD provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static Credentials

!> **Warning:** Hard-coded credentials are not recommended in any Terraform
configuration and risks secret leakage should this file ever be committed to a
public version control system.

Static credentials can be provided by adding an `username` and `password`
in-line in the AWS provider block:

Usage:

```terraform
provider "gocd" {
  baseurl     = "http://gocd.local/go"
  username = "my-username"
  password = "my-password"
}
```

### Environment Variables

You can provide your credentials via the `GOCD_USERNAME` and
`GOCD_PASSWORD`, environment variables, representing your GoCD
Username and Password, respectively:

```terraform
provider "gocd" {}
```

Usage:

```sh
$ export GOCD_URL="http://gocd.local/go"
$ export GOCD_USERNAME="my-username"
$ export GOCD_PASSWORD="my-password"
$ export GOCD_SKIP_SSL_CHECK="true"
$ terraform plan
```

## Schema

### Required

- **baseurl** (String)

### Optional

- **password** (String) Password for User for GoCD API interaction.
- **skip_ssl_check** (Boolean)
- **username** (String) User to interact with the GoCD API with.
