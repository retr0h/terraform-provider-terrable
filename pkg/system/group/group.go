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

package group

import (
	"fmt"
	"os/user"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
)

const (
	LinuxGroupAddCommand    = "/usr/sbin/groupadd"
	LinuxGroupDeleteCommand = "/usr/sbin/groupdel"
)

type Group struct {
	Name string
	GID  string
}

// Lookup looks up a group by name. If the user cannot be found,
// the returned error is of type UnknownGroupError.
func Lookup(groupName string) (*user.Group, error) {
	u, err := user.LookupGroup(groupName)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Add a group. If the group cannot be added the returned error contains
// the stderr of the command executed.
func (g Group) Add(i interface{}) error {
	exec := i.(exec.CommanderDelegate)
	gid := g.GID

	cmd := LinuxGroupAddCommand
	cmdArgs := []string{}

	if gid != "" {
		cmdArgs = append(cmdArgs, "-g", gid)
	}

	cmdArgs = append(cmdArgs, g.Name)

	out, err := exec.Run(cmd, cmdArgs...)
	if err != nil {
		return fmt.Errorf("%w\n\nOut:\n%s", err, out)
	}

	return nil
}

// Delete a group. If the group cannot be deleted the error contains
// the stderr of the command executed.
func (g Group) Delete(i interface{}) error {
	exec := i.(exec.CommanderDelegate)
	cmd := LinuxGroupDeleteCommand
	cmdArgs := []string{
		g.Name,
	}

	out, err := exec.Run(cmd, cmdArgs...)
	if err != nil {
		return fmt.Errorf("%w\n\nOut:\n%s", err, out)
	}

	return nil
}
