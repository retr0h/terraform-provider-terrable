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
	log "github.com/retr0h/terraform-provider-terrable/pkg/logging"
	"github.com/retr0h/terraform-provider-terrable/pkg/user"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		// Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(defaultCreateTimeout),
			// Update: schema.DefaultTimeout(defaultUpdateTimeout),
			Delete: schema.DefaultTimeout(defaultDeleteTimeout),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the user, also acts as it's unique ID",
				Required:    true,
				ForceNew:    true,
				// ValidateFunc: validateName,
			},
			"shell": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Login shell of the new account",
			},
			// "servers_running": {
			//     Type:        schema.TypeString,
			//     Computed:    true,
			//     Description: "Number of running servers",
			// },
		},
	}
}

// func validateName(val interface{}, key string) (warns []string, errs []error) {
//     name := val.(string)

//     if err := cluster.CheckName(name); err != nil {
//         errs = append(errs, fmt.Errorf("%s", err))
//     }

//     return
// }

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	log.Info().
		Msg("Something in create")

	if err := createUser(d); err != nil {
		return err
	}

	name := d.Get("name").(string)
	d.SetId(name)

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Info().
		Str("name", name).
		Msg("Something in user read fucked")

	_, err := user.Lookup(name)
	if err != nil {
		d.SetId("")

		return err
	}

	d.SetId(name)

	return nil
}

func resourceUserUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceUserRead(d, meta)
}

func resourceUserDelete(d *schema.ResourceData, meta interface{}) error {
	log.Info().
		Msg("Something in delete")

	if err := deleteUser(d); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteUser(d *schema.ResourceData) error {
	id := d.Id()

	if err := user.Delete(id); err != nil {
		log.Error().
			Err(err).
			Msg("Something in delete user fucked up")
		return err
	}

	return nil
}

func createUser(d *schema.ResourceData) error {
	name := d.Get("name").(string)
	shell := d.Get("shell").(string)

	log.Info().
		Msg("Something in createUser")

	if err := user.Add(name, "foo", shell); err != nil {
		log.Error().
			Err(err).
			Msg("Something in create user fucked up")
		return err
	}

	return nil
}

func listUser(d *schema.ResourceData) ([]byte, error) {
	// id := d.Id()

	// log.Info().
	//     Str("id", id).
	//     Msg("List User")

	// args := []string{"cluster", "list", id, "--no-headers"}
	// cmd := exec.Command("k3d", args...)
	// out, err := cmd.CombinedOutput()

	// if err != nil {
	//     return out, fmt.Errorf("Reading cluster: '%s'\n\n%s", id, string(out))
	// }

	// return out, nil
	return []byte(""), nil
}
