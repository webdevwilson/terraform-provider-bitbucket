variable "username" {}

variable "password" {}

# configure the resource
provider "bitbucket" {
  username = "${var.username}"
  password = "${var.password}"
}

# the minimum we need to create a group
resource "bitbucket_group" "minimal" {
  name                      = "tf_provider_readonly"
  owner                     = "${var.username}"
  permission                = "read"
  auto_add                  = false
  email_forwarding_disabled = false
}

# this will add the calling user to the group
resource "bitbucket_group_membership" "minimal_member" {
  name          = "${bitbucket_group.minimal.name}"
  owner         = "${bitbucket_group.minimal.owner}"
  email_address = "${var.username}"
}
