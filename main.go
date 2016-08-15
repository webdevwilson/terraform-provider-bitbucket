/*
The MIT License (MIT)
Copyright (c) <year> <copyright holders>

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	opts := plugin.ServeOpts{
		ProviderFunc: Provider,
	}
	plugin.Serve(&opts)
}

// Provider returns the schema for the Terraform provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{ // Source https://github.com/hashicorp/terraform/blob/v0.6.6/helper/schema/provider.go#L20-L43
		Schema: providerSchema(),
		ResourcesMap: map[string]*schema.Resource{
			"group": resources.GroupResource(),
		},
		ConfigureFunc: providerConfigure,
	}
}

/*
  provider "bitbucket" {
    username = "foo"
		password = "bar"
  }
*/
func providerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"username": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The username to authenticate with",
		},
		"password": &schema.Schema{
			Type:        schema.TypeString,
			Required:    true,
			Description: "The password for authentication",
		},
	}
}

// Initializes the Bitbucket HTTP Client
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	client := client.New(username, password)
	return &client, nil
}

//
// // The methods defined below will get called for each resource that needs to
// // get created (createFunc), read (readFunc), updated (updateFunc) and deleted (deleteFunc).
// // For example, if 10 resources need to be created then `createFunc`
// // will get called 10 times every time with the information for the proper
// // resource that is being mapped.
// //
// // If at some point any of these functions returns an error, Terraform will
// // imply that something went wrong with the modification of the resource and it
// // will prevent the execution of further calls that depend on that resource
// // that failed to be created/updated/deleted.
//
