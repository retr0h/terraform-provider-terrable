<!--ts-->
* [Resources](#resources)
    * [File](#file)
        * [template](#template)
    * [System](#system)
        * [group](#group)
        * [user](#user)
<!--te-->

# Resources

## File

### template

Rather than building our own implementation, use Terraform's
[templatefile function][].

[templatefile function]: https://www.terraform.io/docs/language/functions/templatefile.html

```hcl
locals {
  content = templatefile("${path.module}/policy.tpl", {
    name = "foo"
  })
}

resource "local_file" "policy" {
  content  = local.content
  filename = "${path.module}/policy.ini"
}
```

## System

### Group

```hcl
provider "terrable" {}

resource "terrable_group" "sudo" {
  name  = "sudo"
}
```

#### Arguments

* `name` - The name of the group (string)
* `gid` (Optional) - Use GID for the new group (string)

### User

```hcl
provider "terrable" {}

resource "terrable_user" "tomcat" {
  name  = "tomcat"
  shell = "/bin/zsh"
}
```

#### Arguments

* `name` - The name of the user (string)
* `shell` - Login shell of the new account (string)
* `directory` (Optional) - Home directory of the new account (string)
* `groups` (Optional) - List of supplementary groups of the new account (list)
* `system` (Optional) - Create a system account (bool)
* `uid` (Optional) - User ID of the new account (string)
* `gid` (Optional) - Name or ID of the primary group of the new account (string)
