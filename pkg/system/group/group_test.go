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
	"testing"

	"github.com/retr0h/terraform-provider-terrable/pkg/exec"
	"github.com/stretchr/testify/assert"
)

func TestLocate(t *testing.T) {
	u, _ := user.Current()
	g, _ := user.LookupGroupId(u.Gid)

	cases := []struct {
		Name      string
		GroupName string
		Err       bool
	}{
		{
			Name:      "Existing group name",
			GroupName: g.Name,
			Err:       false,
		},
		{
			Name:      "Non-existing group name",
			GroupName: "invalid",
			Err:       true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			u, err := Lookup(tc.GroupName)

			if tc.Err == false {
				assert.Equal(t, u.Name, tc.GroupName)
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

		})
	}
}

func TestAdd(t *testing.T) {
	fc := &exec.FakeCommander{}
	fuc := &exec.FakeUnsuccessfulCommander{}
	cases := []struct {
		Name      string
		Group     *Group
		Want      []string
		Err       bool
		Commander exec.CommanderDelegate
	}{
		{
			Name: "All fields",
			Group: &Group{
				Name: "fake-group",
				GID:  "1099",
			},
			Want: []string{
				"/usr/sbin/groupadd",
				"-g", "1099",
				"fake-group",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Without optional fields",
			Group: &Group{
				Name: "fake-group",
			},
			Want: []string{
				"/usr/sbin/groupadd",
				"fake-group",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Returns an error",
			Group: &Group{
				Name: "fake-name",
			},
			Want:      []string(nil),
			Err:       true,
			Commander: fuc,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			g := tc.Group
			err := g.Add(tc.Commander)

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
	fc := &exec.FakeCommander{}
	fuc := &exec.FakeUnsuccessfulCommander{}
	cases := []struct {
		Name      string
		Group     *Group
		Want      []string
		Err       bool
		Commander exec.CommanderDelegate
	}{
		{
			Name: "Default",
			Group: &Group{
				Name: "fake-group",
			},
			Want: []string{
				"/usr/sbin/groupdel",
				"fake-group",
			},
			Err:       false,
			Commander: fc,
		},
		{
			Name: "Returns an error",
			Group: &Group{
				Name: "fake-group",
			},
			Want:      []string(nil),
			Err:       true,
			Commander: fuc,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			g := tc.Group
			err := g.Delete(tc.Commander)

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
