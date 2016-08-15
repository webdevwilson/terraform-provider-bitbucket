package resource

import "github.com/hashicorp/terraform/helper/schema"

// GroupResource returns the schema for the group resource
func GroupResource() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createFunc,
		Read:          readFunc,
		Update:        updateFunc,
		Delete:        deleteFunc,
		Schema: map[string]*schema.Schema{ // List of supported configuration fields for your resource
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the group",
			},
			"permission": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "read, write, or admin",
			},
			"auto_add": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    false,
				Default:     false,
				Description: "Should all new users be added to this group?",
			},
		},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	return nil
}
