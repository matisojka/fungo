package fungo

import (
	. "launchpad.net/gocheck"
)

// Helper functions

var isOdd = func(a int) bool {
	if a%2 != 0 {
		return true
	}
	return false
}

var sum = func(a int, b int) int {
	return a + b
}

var isThree = func(a int) bool {
	if a == 3 {
		return true
	}
	return false
}

var isFour = func(a int) bool {
	if a == 4 {
		return true
	}
	return false
}

// Tests

func (s *TestSuite) TestIntFilter(c *C) {
	res_elems := IntFilter(elems, isOdd)
	c.Assert(res_elems, DeepEquals, []int{1, 3})
}

func (s *TestSuite) TestIntMap(c *C) {
	double := func(a int) int {
		return a * 2
	}

	res_elems := IntMap(elems, double)
	c.Assert(res_elems, DeepEquals, []int{2, 4, 6})
}

func (s *TestSuite) TestIntReduce(c *C) {
	memo := 0

	res := IntReduce(elems, sum, memo)
	c.Assert(res, Equals, 6)
}

func (s *TestSuite) TestIntReduceMemo(c *C) {
	memo := 5

	res := IntReduce(elems, sum, memo)
	c.Assert(res, Equals, 11)
}

func (s *TestSuite) TestIntFindFound(c *C) {
	res, ok := IntFind(elems, isThree)
	c.Assert(ok, Equals, true)
	c.Assert(res, Equals, 3)
}

func (s *TestSuite) TestIntFindNotFound(c *C) {
	elems = []int{1, 2}

	_, ok := IntFind(elems, isThree)
	c.Assert(ok, Equals, false)
}

func (s *TestSuite) TestIntReject(c *C) {
	res_elems := IntReject(elems, isOdd)
	c.Assert(res_elems, DeepEquals, []int{2})
}

func (s *TestSuite) TestIntEveryTrue(c *C) {
	smallerThanFour := func(a int) bool {
		if a < 4 {
			return true
		}

		return false
	}

	every := IntEvery(elems, smallerThanFour)
	c.Assert(every, Equals, true)
}

func (s *TestSuite) TestIntEveryFalse(c *C) {
	smallerThanThree := func(a int) bool {
		if a < 3 {
			return true
		}

		return false
	}

	every := IntEvery(elems, smallerThanThree)
	c.Assert(every, Equals, false)
}

func (s *TestSuite) TestIntSomeTrue(c *C) {
	some := IntSome(elems, isThree)
	c.Assert(some, Equals, true)
}

func (s *TestSuite) TestIntSomeFalse(c *C) {
	some := IntSome(elems, isFour)
	c.Assert(some, Equals, false)
}

func (s *TestSuite) TestIntContainTrue(c *C) {
	contains := IntContain(elems, 3)
	c.Assert(contains, Equals, true)
}

func (s *TestSuite) TestIntContainFalse(c *C) {
	contains := IntContain(elems, 4)
	c.Assert(contains, Equals, false)
}

func (s *TestSuite) TestIntWithout(c *C) {
	without_elems := []int{1, 2}

	ret_elems := IntWithout(elems, without_elems)
	c.Assert(ret_elems, DeepEquals, []int{3})
}

func (s *TestSuite) TestIntWithout2(c *C) {
	elems = []int{1, 2, 2}
	without_elems := []int{2}

	ret_elems := IntWithout(elems, without_elems)
	c.Assert(ret_elems, DeepEquals, []int{1})
}
