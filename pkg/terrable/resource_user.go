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
	"fmt"
	"regexp"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
	log "github.com/retr0h/terraform-provider-terrable/pkg/logging"
	"github.com/retr0h/terraform-provider-terrable/pkg/system/user"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	LinuxUserNameMinLength = 1
	LinuxUserNameMaxLength = 32
)

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Delete: resourceUserDelete,
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
				Description:  "The name of the user, also acts as it's unique ID",
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateName,
			},
			"shell": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Login shell of the new account",
			},
			"directory": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Home directory of the new account",
			},
			"groups": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "List of supplementary groups of the new account",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"system": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Create a system account",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "User ID of the new account",
			},
			"gid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name or ID of the primary group of the new account",
			},
		},
	}
}

// validateName validates a username against rules defined by Debians's
// USERADD(8).  If the validation fails the returned error is of type Error.
//
// It is usually recommended to only use usernames that begin with a lower case
// letter or an underscore, followed by lower case letters, digits, underscores,
// or dashes. They can end with a dollar sign.
//   In regular expression terms: [a-z_][a-z0-9_-]*[$]?
//
// On Debian, the only constraints are that usernames must neither start with a
// dash ('-') nor contain a colon (':') or a whitespace (space: '', end of line:
// '\n', tabulation: '	', etc.).  Note that using a slash ('/') may break the
// default algorithm for the definition of the user's home directory.
//
// Usernames may only be up to 32 characters long.
func validateName(val interface{}, key string) (warns []string, errs []error) {
	name := val.(string)
	regex := fmt.Sprintf("^[a-z][a-z0-9_-]{%d,%d}$",
		(LinuxUserNameMinLength - 1), // The regexp enforces at least one character.
		(LinuxUserNameMaxLength - 1))

	match, _ := regexp.MatchString(regex, name)
	if !match {
		errs = append(errs, fmt.Errorf("Username invalid, must match: '%s'", regex))
	}

	return
}

func resourceUserCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Info().
		Str("name", name).
		Str("func", "resourceUserCreate").
		Msg("Terrable")

	if err := createUser(d); err != nil {
		return err
	}

	d.SetId(name)

	return resourceUserRead(d, meta)
}

func resourceUserRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)

	log.Info().
		Str("name", name).
		Str("func", "resourceUserRead").
		Msg("Terrable")

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
		Str("func", "resourceUserDelete").
		Msg("Terrable")

	if err := deleteUser(d); err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func deleteUser(d *schema.ResourceData) error {
	id := d.Id()

	log.Info().
		Str("id", id).
		Str("func", "deleteUser").
		Msg("Terrable")

	commander := &exec.Commander{}
	u := user.User{
		Name: id,
	}

	if err := u.Delete(commander); err != nil {
		log.Error().
			Str("func", "deleteUser").
			Err(err).
			Msg("Terrable")
		return err
	}

	return nil
}

func createUser(d *schema.ResourceData) error {
	name := d.Get("name").(string)
	shell := d.Get("shell").(string)
	directory := d.Get("directory").(string)
	inputGroups := d.Get("groups").([]interface{})
	system := d.Get("system").(bool)
	uid := d.Get("uid").(string)
	gid := d.Get("gid").(string)

	groups := make([]string, len(inputGroups))
	for i, v := range inputGroups {
		groups[i] = fmt.Sprint(v)
	}

	log.Info().
		Str("name", name).
		Str("func", "createUser").
		Msg("Terrable")

	commander := &exec.Commander{}
	u := user.User{
		Name:      name,
		Shell:     shell,
		Directory: directory,
		Groups:    groups,
		System:    system,
		UID:       uid,
		GID:       gid,
	}

	if err := u.Add(commander); err != nil {
		log.Error().
			Str("func", "createUser").
			Err(err).
			Msg("Terrable")
		return err
	}

	return nil
}
