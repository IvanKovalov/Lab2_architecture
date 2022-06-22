package lab2

import (
	"fmt"
	"strings"
)

// TODO: document this function.
// PrefixToPostfix converts
var a string = "4 2 - 3 * 5 +"

func PostfixToInfix(s string) (string, error) {

	var strerr string
	str := strings.Split(s, " ")
	lenght := len(str)
	if len(s) == 0 {
		strerr = "Пустая строчка"
		return "Пустая строчка", fmt.Errorf(strerr)
	}

	if strings.Contains(s, `^[a-zA-Z]+$`) {
		strerr = "Недопустимые символы"
		return "Недопустимые символы", fmt.Errorf(strerr)
	}

	stack := make([]string, 0)

	for i := 0; i < lenght; i++ {
		if strings.ContainsAny("0123456789", str[i]) {
			stack = append(stack, str[i])
		} else if len(stack) >= 2 {
			if i == lenght-1 {
				opperand2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				opperand1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, ""+opperand1+str[i]+opperand2+"")
			} else {
				opperand2 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				opperand1 := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, "("+opperand1+str[i]+opperand2+")")
			}
		}

	}

	/*for i := 0; i < lenght; i++ {
		if str[i] == "+" || str[i] == "-" || str[i] == "*" || str[i] == "/" || str[i] == "^" {
			prev := str[i-1]
			symbol := str[i]
			str[i-1] = symbol
			str[i] = prev
		}

	}*/
	result := strings.Join(stack, "")
	return result, fmt.Errorf(strerr)
}
