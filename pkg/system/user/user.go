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

package user

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
)

const (
	LinuxUserAddCommand    = "/usr/sbin/useradd"
	LinuxUserDeleteCommand = "/usr/sbin/userdel"
)

type User struct {
	Name      string
	Directory string
	Shell     string
	Groups    []string
}

// Lookup looks up a user by username. If the user cannot be found,
// the returned error is of type UnknownUserError.
func Lookup(userName string) (*user.User, error) {
	u, err := user.Lookup(userName)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Add a user. If the user cannot be added the returned error contains
// the stderr of the command executed.
func (u User) Add(i interface{}) error {
	exec := i.(exec.CommanderDelegate)
	directory := u.Directory
	groups := u.Groups

	cmd := LinuxUserAddCommand
	cmdArgs := []string{
		"-s", u.Shell,
		"-m", // create home
	}

	if directory != "" {
		cmdArgs = append(cmdArgs, "-d", directory)
	} else {
		directory = fmt.Sprintf("/home/%s", u.Name)
		cmdArgs = append(cmdArgs, "-d", directory)
	}

	if len(groups) > 0 {
		joinedGroups := strings.Join(groups, ",")
		cmdArgs = append(cmdArgs, "-G", joinedGroups)
	}

	cmdArgs = append(cmdArgs, u.Name)

	out, err := exec.Run(cmd, cmdArgs...)
	if err != nil {
		return fmt.Errorf("%w\n\nOut:\n%s", err, out)
	}

	return nil
}

// Delete a user. If the user cannot be deleted the error contains
// the stderr of the command executed.
func (u User) Delete(i interface{}) error {
	exec := i.(exec.CommanderDelegate)
	cmd := LinuxUserDeleteCommand
	cmdArgs := []string{
		u.Name, // account
	}

	out, err := exec.Run(cmd, cmdArgs...)
	if err != nil {
		return fmt.Errorf("%w\n\nOut:\n%s", err, out)
	}

	return nil
}
