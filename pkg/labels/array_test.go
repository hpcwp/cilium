// Copyright 2016-2017 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package labels

import (
	"github.com/cilium/cilium/pkg/comparator"

	. "gopkg.in/check.v1"
)

var _ = Suite(&LabelsSuite{})

func (s *LabelsSuite) TestMatches(c *C) {
	a := LabelArray{
		NewLabel("1", "1", "1"),
		NewLabel("2", "2", "1"),
		NewLabel("3", "3", "1"),
	}
	b := LabelArray{
		NewLabel("1", "1", "1"),
		NewLabel("2", "2", "1"),
	}
	empty := LabelArray{}

	c.Assert(a.Contains(b), Equals, true)      // b is in a
	c.Assert(b.Contains(a), Equals, false)     // a is NOT in b
	c.Assert(a.Contains(empty), Equals, true)  // empty is in a
	c.Assert(b.Contains(empty), Equals, true)  // empty is in b
	c.Assert(empty.Contains(a), Equals, false) // a is NOT in empty
}

func (s *LabelsSuite) TestParse(c *C) {
	c.Assert(ParseLabelArray(), comparator.DeepEquals, LabelArray{})
	c.Assert(ParseLabelArray("magic"), comparator.DeepEquals, LabelArray{ParseLabel("magic")})
	c.Assert(ParseLabelArray("a", "b", "c"), comparator.DeepEquals,
		LabelArray{ParseLabel("a"), ParseLabel("b"), ParseLabel("c")})
}

func (s *LabelsSuite) TestHas(c *C) {
	lbls := LabelArray{
		NewLabel("env", "devel", LabelSourceAny),
		NewLabel("user", "bob", LabelSourceContainer),
	}
	var hasTests = []struct {
		input    string // input
		expected bool   // expected result
	}{
		{"", false},
		{"any", false},
		{"env", true},
		{"container.env", false},
		{"container:env", false},
		{"any:env", false},
		{"any.env", true},
		{"any:user", false},
		{"any.user", true},
		{"user", true},
		{"container.user", true},
		{"container:bob", false},
	}
	for _, tt := range hasTests {
		c.Logf("has %q?", tt.input)
		c.Assert(lbls.Has(tt.input), Equals, tt.expected)
	}
}
