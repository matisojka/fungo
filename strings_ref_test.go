package fungo

import (
	. "launchpad.net/gocheck"
)

func (s *TestSuite) TestStringFilterR(c *C) {
	StringFilterR(&strings, longerThanThree)
	c.Assert(strings, DeepEquals, []string{"long", "string"})
}

func (s *TestSuite) TestStringMapR(c *C) {
  StringMapR(&strings, withPrefix)
  c.Assert(strings, DeepEquals, []string{"preone", "prelong", "prestring"})
}

func (s *TestSuite) TestStringRejectR(c *C) {
  StringRejectR(&strings, longerThanThree)
  c.Assert(strings, DeepEquals, []string{"one"})
}

func (s *TestSuite) TestStringWithoutR(c *C) {
  banned_strings := []string{"one", "long"}
  StringWithoutR(&strings, &banned_strings)
  c.Assert(strings, DeepEquals, []string{"string"})
}
