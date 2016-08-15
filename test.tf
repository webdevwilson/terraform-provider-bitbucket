variable "username" {}

variable "password" {}

provider "bitbucket" {
  username = "${var.username}"
  password = "${var.password}"
}
