package lab2

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type Reader interface {
	Read(str string)
}

type Writer interface {
	Write(str string)
}

type ComputeHandler struct {
	// TODO: Add necessary fields.
	Input  string
	Output string
}

func (ch ComputeHandler) Read() string {
	if strings.Contains(ch.Input, ".txt") {
		file, err := os.Open(ch.Input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		data := make([]byte, 64)

		for {
			n, err := file.Read(data)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
			fmt.Print(string(data[:n]))
		}
		return string(data)
	}
	return ch.Input
}

func (ch ComputeHandler) Write(str string) {
	f, err := os.OpenFile(ch.Output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte(str))
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.

	ex := ch.Read()
	res, _ := PostfixToInfix(ex)
	if ch.Output == "default" {
		fmt.Println(res)
	} else {
		ch.Write(res)
	}
	return nil
}

var st1 = ComputeHandler{"4 3 -", "default"}

var st2 = st1.Compute()

func main() {
	fmt.Println(st2)
}
