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
			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The url slug for the group",
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
	name := d.Get("name").(string)

	// if the owner is self, use the current owner
	if owner == "self" {
		current, err := bitbucket.Users.Current()
		if err != nil {
			return err
		}
		owner = current.User.Username
	}

	// make the API call
	group, err := bitbucket.Groups.Create(owner, name)
	if err != nil {
		return err
	}

	d.SetId(group.ResourceURI)
	d.Set("resource_uri", group.ResourceURI)
	d.Set("slug", group.Slug)

	// now update with the rest of the values
	// group.AutoAdd = d.Get("auto_add").(bool)
	// group.EmailForwardingDisabled = d.Get("email_forwarding_disabled").(bool)
	// group.Permission = d.Get("permission").(string)
	//
	// bitbucket.Groups.Update(owner, name)

	return readFunc(d, meta)
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	bitbucket := meta.(*bitbucket.Client)

	if bitbucket == nil {
		return errors.New("bitbucket client not found")
	}

	owner, err := getOwner(d, meta)
	if err != nil {
		d.SetId("")
		return err
	}

	slug := d.Get("slug").(string)

	group, err := bitbucket.Groups.Find(owner, slug)
	if err != nil {
		return err
	}

	if group == nil {
		return errors.New("Error reading group")
	}

	d.Set("auto_add", group.AutoAdd)
	d.Set("email_forwarding_disabled", group.EmailForwardingDisabled)
	d.Set("name", group.Name)
	d.Set("slug", group.Slug)
	d.Set("owner", group.Owner.Username)
	d.Set("permission", group.Permission)
	d.Set("resource_uri", group.ResourceURI)

	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	return errors.New("updateFunc called")
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	bitbucket := meta.(*bitbucket.Client)

	if bitbucket == nil {
		return errors.New("bitbucket client not found")
	}

	owner, err := getOwner(d, meta)
	if err != nil {
		return err
	}
	slug := d.Get("slug").(string)

	return bitbucket.Groups.Delete(owner, slug)
}

func getOwner(d *schema.ResourceData, meta interface{}) (string, error) {
	bitbucket := meta.(*bitbucket.Client)
	owner := d.Get("owner").(string)
	if owner == "self" {
		current, err := bitbucket.Users.Current()
		if err != nil {
			return "", err
		}
		owner = current.User.Username
	}
	return owner, nil
}
