package lab2

import (
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var suit = Suite(&MySuite{})

func (s *MySuite) TestPostfixToInfix(c *C) {
	res, _ := PostfixToInfix("7 8 + 3 2 + /")
	c.Check(res, Equals, "(7+8)(3+2)")
}

func (s *MySuite) TestErrorPostfixToInfix(c *C) {
	_, err := PostfixToInfix("")
	c.Check(err, Not(Equals), fmt.Errorf("Пустая строчка"))
}

func (s *MySuite) ExamplePostfixToInfix(c *C) {
	res, _ := PostfixToInfix("4 2 +")
	c.Check(res, Equals, "4+2")
	fmt.Println(res)
}
