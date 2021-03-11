<!--ts-->
* [Resources](#resources)
    * [System](#system)
        * [User](#user)
<!--te-->

# Resources

## System

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
