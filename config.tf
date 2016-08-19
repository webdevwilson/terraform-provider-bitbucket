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
  owner                     = "self"
  permission                = "read"
  auto_add                  = false
  email_forwarding_disabled = false

  /*members = ["${var.username}"]*/
}
