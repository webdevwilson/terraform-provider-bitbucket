package resources

import (
	"errors"
	"fmt"

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
				Required:    true,
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
	bitbucket, err := bitbucketClient(meta)
	if err != nil {
		return err
	}

	// get the owner
	owner := d.Get("owner").(string)
	name := d.Get("name").(string)

	// make the API call
	group, err := bitbucket.Groups.Create(owner, name)
	if err != nil {
		return err
	}

	// set computed values
	d.SetId(fmt.Sprintf("%s/%s", group.Owner.Username, group.Slug))
	d.Set("auto_add", group.AutoAdd)
	d.Set("email_forwarding_disabled", group.EmailForwardingDisabled)
	d.Set("name", group.Name)
	d.Set("slug", group.Slug)
	d.Set("owner", group.Owner.Username)
	d.Set("permission", group.Permission)

	return nil
}

func readFunc(d *schema.ResourceData, meta interface{}) error {
	bitbucket, err := bitbucketClient(meta)
	if err != nil {
		return err
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
	d.Set("owner", group.Owner.Username)
	d.Set("slug", group.Slug)
	d.Set("permission", group.Permission)

	return nil
}

func updateFunc(d *schema.ResourceData, meta interface{}) error {
	_, err := bitbucketClient(meta)
	if err != nil {
		return err
	}

	return errors.New("updateFunc called")
}

func deleteFunc(d *schema.ResourceData, meta interface{}) error {
	bitbucket, err := bitbucketClient(meta)
	if err != nil {
		return err
	}

	owner, err := getOwner(d, meta)
	if err != nil {
		return err
	}
	slug := d.Get("slug").(string)

	return bitbucket.Groups.Delete(owner, slug)
}

func bitbucketClient(meta interface{}) (*bitbucket.Client, error) {
	client := meta.(*bitbucket.Client)
	if client == nil {
		return nil, errors.New("bitbucket client not found")
	}
	return client, nil
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

func readOwner(bitbucket *bitbucket.Client, desired string, owner string) (string, error) {
	current, err := bitbucket.Users.Current()
	if err == nil {
		return "", err
	}
	currentUser := current.User.Username

	if desired == "self" && owner == currentUser {
		return desired, nil
	}

	return owner, nil
}
