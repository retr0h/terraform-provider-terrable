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
	"os/user"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
)

type User struct {
	Name      string
	Directory string
	Shell     string
	// Commander is swapped for tests.
	Commander exec.CommanderDelegate
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

// Add a user. If the user cannot be added the returned error is of
// type *exec.ExitError.
func (u User) Add() error {
	cmd := "useradd"
	cmdArgs := []string{
		"-d", u.Directory,
		"-s", u.Shell,
		"-m",   // create home
		u.Name, // account
	}

	_, err := u.Commander.Run(cmd, cmdArgs...)

	return err
}

func Add(userName string, directory string, shell string) error {
	cmd := "/usr/sbin/useradd"
	cmdArgs := []string{
		//"-d", directory,
		"-s", shell,
		"-m",     // create home
		userName, // account
	}

	c := exec.Commander{}

	_, err := c.Run(cmd, cmdArgs...)

	return err
}

func Delete(userName string) error {
	cmd := "/usr/sbin/userdel"
	cmdArgs := []string{
		userName, // account
	}

	c := exec.Commander{}

	_, err := c.Run(cmd, cmdArgs...)

	return err
}
