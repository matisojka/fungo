package fungo

import (
	. "launchpad.net/gocheck"
)

// Helper functions

func longerThanTwo(s string) bool {
	if len(s) > 2 {
		return true
	}
	return false
}

func longerThanThree(s string) bool {
	if len(s) > 3 {
		return true
	}
	return false
}

func longerThanSix(s string) bool {
  if len(s) > 6 {
    return true
  }
  return false
}

func withPrefix(s string) string {
  return "pre" + s
}

// Tests

func (s *TestSuite) TestStringFilter(c *C) {
	res_strings := StringFilter(strings, longerThanThree)
	c.Assert(res_strings, DeepEquals, []string{"long", "string"})
}

func (s *TestSuite) TestStringMap(c *C) {
  res_strings := StringMap(strings, withPrefix)
  c.Assert(res_strings, DeepEquals, []string{"preone", "prelong", "prestring"})
}

func (s *TestSuite) TestStringFindTrue(c *C) {
  long_string, ok := StringFind(strings, longerThanThree)
  c.Assert(ok, Equals, true)
  c.Assert(long_string, Equals, "long")
}

func (s *TestSuite) TestStringFindFalse(c *C) {
  _, ok := StringFind(strings, longerThanSix)
  c.Assert(ok, Equals, false)
}

func (s *TestSuite) TestStringReject(c *C) {
  res_strings := StringReject(strings, longerThanThree)
  c.Assert(res_strings, DeepEquals, []string{"one"})
}

func (s *TestSuite) TestStringEveryTrue(c *C) {
  every := StringEvery(strings, longerThanTwo)
  c.Assert(every, Equals, true)
}

func (s *TestSuite) TestStringEveryFalse(c *C) {
  every := StringEvery(strings, longerThanThree)
  c.Assert(every, Equals, false)
}

func (s *TestSuite) TestStringContainTrue(c *C) {
  contains := StringContain(strings, "one")
  c.Assert(contains, Equals, true)
}

func (s *TestSuite) TestStringContainFalse(c *C) {
  contains := StringContain(strings, "two")
  c.Assert(contains, Equals, false)
}

func (s *TestSuite) TestStringWithout(c *C) {
  banned_strings := []string{"one", "long"}
  allowed := StringWithout(strings, banned_strings)
  c.Assert(allowed, DeepEquals, []string{"string"})
}
