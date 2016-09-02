package resources

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
)

// GroupMembershipResource provides the schema for the bitbucket_group_membership resource
func GroupMembershipResource() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		Create:        createGroupMembership,
		Read:          readGroupMembership,
		Update:        updateGroupMembership,
		Delete:        deleteGroupMembership,
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
			"email_address": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The email address of the person to add to the group",
			},
		},
	}
}

func createGroupMembership(d *schema.ResourceData, meta interface{}) error {
	bitbucket, err := bitbucketClient(meta)
	if err != nil {
		return err
	}

	owner, err := getOwner(d, meta)
	if err != nil {
		return err
	}

	group := d.Get("group").(string)
	email := d.Get("email_address").(string)

	_, err = bitbucket.Groups.AddMember(owner, group, email)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", owner, group, email))
	return nil
}

func readGroupMembership(d *schema.ResourceData, meta interface{}) error {
	_, err := bitbucketClient(meta)
	if err != nil {
		return err
	}
	return errors.New("readGroupMembership called")
}

func updateGroupMembership(d *schema.ResourceData, meta interface{}) error {
	_, err := bitbucketClient(meta)
	if err != nil {
		return err
	}
	return errors.New("updateGroupMembership called")
}

func deleteGroupMembership(d *schema.ResourceData, meta interface{}) error {
	_, err := bitbucketClient(meta)
	if err != nil {
		return err
	}
	return errors.New("deleteGroupMembership called")
}
