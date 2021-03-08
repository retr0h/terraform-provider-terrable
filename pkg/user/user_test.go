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
	"testing"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
	"github.com/stretchr/testify/assert"
)

type fakeCommander struct {
	command []string
}

func (c *fakeCommander) Run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	c.command = cmd.Args

	return []byte(""), nil
}

func (fc *fakeCommander) About() []string {
	return fc.command
}

type fakeUnsuccessfulCommander struct {
	command []string
}

func (c *fakeUnsuccessfulCommander) Run(name string, args ...string) ([]byte, error) {
	err := fmt.Errorf("faked run error")

	return []byte(nil), err
}

func (c *fakeUnsuccessfulCommander) Delete() error {
	return fmt.Errorf("faked delete error")
}

func (fuc *fakeUnsuccessfulCommander) About() []string {
	return fuc.command
}

func TestLocate(t *testing.T) {
	userName := "root"
	u, err := Lookup(userName)
	assert.Equal(t, u.Username, "root")

	assert.NoError(t, err)
}

func TestLocateErrorOnInvalidUser(t *testing.T) {
	userName := "invalid"
	_, err := Lookup(userName)

	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	fc := &fakeCommander{}
	fuc := &fakeUnsuccessfulCommander{}
	cases := []struct {
		Name      string
		User      *User
		Want      []string
		Err       bool
		Commander exec.CommanderDelegate
	}{
		{
			Name: "Default",
			User: &User{
				Name:      "fake-name",
				Directory: "fake-dir",
				Shell:     "fake-shell",
			},
			Want: []string{
				"/usr/sbin/useradd",
				"-s", "fake-shell",
				"-m",
				"-d", "fake-dir",
				"fake-name",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Without directory field",
			User: &User{
				Name:  "fake-name",
				Shell: "fake-shell",
			},
			Want: []string{
				"/usr/sbin/useradd",
				"-s", "fake-shell",
				"-m",
				"-d", "/home/fake-name",
				"fake-name",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Returns an error",
			User: &User{
				Name:  "fake-name",
				Shell: "fake-shell",
			},
			Want:      []string(nil),
			Err:       true,
			Commander: fuc,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			u := tc.User
			err := u.Add(tc.Commander)

			got := tc.Commander.About()
			assert.Equal(t, tc.Want, got)

			if tc.Err == false {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	fc := &fakeCommander{}
	fuc := &fakeUnsuccessfulCommander{}
	cases := []struct {
		Name      string
		User      *User
		Want      []string
		Err       bool
		Commander exec.CommanderDelegate
	}{
		{
			Name: "Default",
			User: &User{
				Name: "fake-user",
			},
			Want: []string{
				"/usr/sbin/userdel",
				"fake-user",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Returns an error",
			User: &User{
				Name: "fake-name",
			},
			Want:      []string(nil),
			Err:       true,
			Commander: fuc,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			u := tc.User
			err := u.Delete(tc.Commander)

			got := tc.Commander.About()
			assert.Equal(t, tc.Want, got)

			if tc.Err == false {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
