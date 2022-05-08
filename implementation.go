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
	if lenght == 0 {
		strerr = "Пустая строчка"
	}

	if strings.Contains(s, `^[a-zA-Z]+$`) {
		strerr = "Недопустимые символы"
	}

	for i := 0; i < lenght; i++ {
		if str[i] == "+" || str[i] == "-" || str[i] == "*" || str[i] == "/" || str[i] == "^" {
			prev := str[i-1]
			symbol := str[i]
			str[i-1] = symbol
			str[i] = prev
		}

	}
	sentence := strings.Join(str, "")
	//fmt.Println(sentence)
	return sentence, fmt.Errorf(strerr)
}

func main() {
	//PostfixToInfix(a)
}
