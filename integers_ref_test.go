package fungo

import (
	. "launchpad.net/gocheck"
)

func (s *TestSuite) TestIntFilterR(c *C) {
	IntFilterR(&elems, isOdd)
	c.Assert(elems, DeepEquals, []int{1, 3})
}

func (s *TestSuite) TestIntMapR(c *C) {
	double := func(a int) int {
		return a * 2
	}

	IntMapR(&elems, double)
	c.Assert(elems, DeepEquals, []int{2, 4, 6})
}

func (s *TestSuite) TestIntRejectR(c *C) {
	IntRejectR(&elems, isOdd)
	c.Assert(elems, DeepEquals, []int{2})
}

func (s *TestSuite) TestIntWithoutR(c *C) {
	without_elems := []int{1, 2}

	IntWithoutR(&elems, &without_elems)
	c.Assert(elems, DeepEquals, []int{3})
}

func (s *TestSuite) TestIntWithoutR2(c *C) {
	elems = []int{1, 2, 2}
	without_elems := []int{2}

	IntWithoutR(&elems, &without_elems)
	c.Assert(elems, DeepEquals, []int{1})
}
