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
    directory = "/home/tomcat"
    groups = ["sudo"]
}
```
