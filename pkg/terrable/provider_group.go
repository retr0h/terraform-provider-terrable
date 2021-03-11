// Copyright (c) 2021 John Dewey <john@dewey.ws>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package terrable

import (
	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
	log "github.com/retr0h/terraform-provider-terrable/pkg/logging"
	"github.com/retr0h/terraform-provider-terrable/pkg/system/group"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Delete: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultCreateTimeout),
			Delete: schema.DefaultTimeout(defaultDeleteTimeout),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Description:  "The name of the group, also acts as it's unique ID",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"gid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Use GID for the new group",
			},
		},
	}
}

func resourceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Info().
		Str("name", name).
		Str("func", "resourceGroupCreate").
		Msg("Terrable")

	if err := createGroup(d); err != nil {
		return err
	}

	d.SetId(name)

	return resourceGroupRead(d, meta)
}

func resourceGroupRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Info().
		Str("name", name).
		Str("func", "resourceGroupRead").
		Msg("Terrable")

	_, err := group.Lookup(name)
	if err != nil {
		d.SetId("")

		return err
	}

	d.SetId(name)

	return nil
}

func resourceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceGroupRead(d, meta)
}

func resourceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	log.Info().
		Str("func", "resourceGroupDelete").
		Msg("Terrable")

	if err := deleteGroup(d); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteGroup(d *schema.ResourceData) error {
	id := d.Id()

	log.Info().
		Str("id", id).
		Str("func", "deleteGroup").
		Msg("Terrable")

	commander := &exec.Commander{}
	g := group.Group{
		Name: id,
	}

	if err := g.Delete(commander); err != nil {
		log.Error().
			Str("func", "deleteGroup").
			Err(err).
			Msg("Terrable")
		return err
	}

	return nil
}

func createGroup(d *schema.ResourceData) error {
	name := d.Get("name").(string)
	gid := d.Get("gid").(string)

	log.Info().
		Str("name", name).
		Str("func", "createGroup").
		Msg("Terrable")

	commander := &exec.Commander{}
	g := group.Group{
		Name: name,
		GID:  gid,
	}

	if err := g.Add(commander); err != nil {
		log.Error().
			Str("func", "createGroup").
			Err(err).
			Msg("Terrable")
		return err
	}

	return nil
}
