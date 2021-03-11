terraform {
  required_version = ">= 0.13.0"
  required_providers {
    terrable = {
      source  = "github.com/retr0h/terrable"
      version = "1.0"
    }
  }
}

provider "terrable" {}

resource "terrable_user" "tomcat" {
  name  = "tomcat"
  shell = "/bin/bash"
  uid   = "1099"
  # TODO(retr0h): Test gid once `terrable_group` written
  # gid = "1099"
}
