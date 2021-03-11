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

resource "terrable_group" "test_group" {
  name = "test_group"
  gid = "2000"
}

resource "terrable_user" "test_user" {
  name  = "test_user"
  shell = "/bin/bash"
  uid   = "1099"
  gid = "2000"

  depends_on = [terrable_group.test_group]
}
