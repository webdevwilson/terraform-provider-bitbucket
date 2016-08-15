default: build

deps:
				go install github.com/hashicorp/terraform

build:
				go build -o terraform-provider-bitbucket .
