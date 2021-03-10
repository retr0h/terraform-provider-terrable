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

resource "terrable_user" "foo" {
  name  = "foo"
  shell = "/bin/bash"
}

resource "terrable_user" "bar" {
  name  = "bar"
  shell = "/bin/bash"
}

resource "terrable_user" "tomcat" {
  name      = "tomcat"
  shell     = "/bin/bash"
  directory = "/usr/local/tomcat"
  groups    = ["foo", "bar"]

  depends_on = [terrable_user.foo, terrable_user.bar]
}
