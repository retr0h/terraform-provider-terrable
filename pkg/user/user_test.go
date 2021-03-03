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
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

type fakeCommander struct {
	testCommand []string
}

func (c *fakeCommander) Run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	c.testCommand = cmd.Args

	return []byte(""), nil
}

func TestAdd(t *testing.T) {
	fakeCommander := &fakeCommander{}
	u := &User{
		Name:      "fake-name",
		Directory: "fake-dir",
		Shell:     "fake-shell",
		Commander: fakeCommander,
	}
	err := u.Add()

	want := []string{
		"useradd",
		"-d", "fake-dir",
		"-s", "fake-shell",
		"-m",
		"fake-name",
	}
	got := fakeCommander.testCommand

	assert.Equal(t, got, want)
	assert.NoError(t, err)
}
