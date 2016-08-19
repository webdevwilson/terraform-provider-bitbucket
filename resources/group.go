package resources

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/webdevwilson/go-bitbucket/bitbucket"
)

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
			"owner": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "self",
				Description: "what account owns this group? default: self",
			},
			"permission": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "read",
				Description: "read, write, or admin",
			},
			"email_forwarding_disabled": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "whether or not email forwarding is disabled for this group",
			},
			"auto_add": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Should all new users be added to this group?",
			},
			"resource_uri": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The uri path to the group",
			},
		},
	}
}

func createFunc(d *schema.ResourceData, meta interface{}) error {
	bitbucket := meta.(*bitbucket.Client)

	if bitbucket == nil {
		return errors.New("bitbucket client not found")
	}

	// get the owner
	owner := d.Get("owner").(string)
	if owner == "self" {
		current, err := bitbucket.Users.Current()
		if err != nil {
			return err
		}
		owner = current.User.Username
	}

	name := d.Get("name").(string)

	bitbucket.Groups.Create(owner, name)

	// make the API call
	_, err := bitbucket.Groups.Create(owner, name)
	if err != nil {
		return err
	}

	// now update with the rest of the values
	// group.AutoAdd = d.Get("auto_add").(bool)
	// group.EmailForwardingDisabled = d.Get("email_forwarding_disabled").(bool)
	// group.Permission = d.Get("permission").(string)
	//
	// bitbucket.Groups.Update(owner, name)

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
