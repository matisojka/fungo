package fungo

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) {
	TestingT(t)
}

type TestSuite struct{}

var _ = Suite(&TestSuite{})

var elems []int
var strings []string

func (s *TestSuite) SetUpTest(c *C) {
	elems = []int{1, 2, 3}
	strings = []string{"one", "long", "string"}
}

