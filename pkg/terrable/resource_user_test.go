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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestAccResourceUser(t *testing.T) {
	rName := acctest.RandString(32)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceUserConfig(rName),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func testAccResourceUserConfig(rName string) string {
	return fmt.Sprintf(`
resource "terrable_user" "tomcat" {
  name  = "%s"
  shell = "/bin/bash"
}
`, rName)
}

func TestValidateName(t *testing.T) {
	cases := []struct {
		Name     string
		Username string
		Err      bool
	}{
		{
			Name: "",
			Err:  true, // less than 1
		},
		{
			Name: "a",
			Err:  false,
		},
		{
			Name: "valid-user",
			Err:  false,
		},
		{
			Name: "valid_user",
			Err:  false,
		},
		{
			Name: "valid-",
			Err:  false,
		},
		{
			Name: "Invalid",
			Err:  true,
		},
		{
			Name: "1invalid",
			Err:  true,
		},
		{
			Name: `in
		valid`,
			Err: true,
		},
		{
			Name: `in valid`,
			Err:  true,
		},
		{
			Name: `in	valid`,
			Err: true,
		},
		{
			Name: `in/valid`,
			Err:  true,
		},
		{
			Name: "invaliduserabcdefghijklmnopqrstuv", // greater than 32
			Err:  true,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d-%s", i, tc.Name), func(t *testing.T) {
			_, err := validateName(tc.Name, "name")

			if tc.Err == false {
				assert.Empty(t, err)
			} else {
				assert.Error(t, err[0])
			}
		})
	}
}
