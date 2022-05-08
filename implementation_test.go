package lab2

import (
	"errors"
	"fmt"
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var suit = Suite(&MySuite{})

func (s *MySuite) TestPostfixToInfix(c *C) {
	res, _ := PostfixToInfix("4 2 - 3 * 5 + ")
	c.Check(res, Equals, "4-2*3+5")
}

func (s *MySuite) TestBigPostfixToInfix(c *C) {
	res, _ := PostfixToInfix("4 2 - 3 * 5 + 6 + 2 * 1 /")
	c.Check(res, Equals, "4-2*3+5+6*2/1")
}

func (s *MySuite) TestErrorPostfixToInfix(c *C) {
	_, err := PostfixToInfix("")
	c.Check(err, Equals, errors.New("Пустая строчка"))
}

func (s *MySuite) ExamplePostfixToInfix(c *C) {
	res, _ := PostfixToInfix("42+")
	c.Check(res, Equals, "4+2")
	fmt.Println(res)
}

func ExamplePrefixToPostfix() {
	res, _ := PostfixToInfix("2 2 +")
	fmt.Println(res)
	//Output:
	// 2+2
}
