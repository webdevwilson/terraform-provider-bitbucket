# Summary

bitbucket provider for Terraform

# Configure Provider

```
provider "bitbucket" {
  username = "foo"
  password = "bar"
}
```

# Create Group

```
resource "bitbucket_group" "group" {
  name       = "admin_group"
  permission = "admin"
  auto_add   = false
}
```
