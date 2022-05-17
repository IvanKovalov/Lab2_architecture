package lab2

import (
	"fmt"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type Reader interface {
	Read(inputStr string) string
}

type Writer interface {
	Write(str string, outFile string)
}

type ComputeHandler struct {
	// TODO: Add necessary fields.
	Rd     Reader
	Wr     Writer
	Input  string
	Output string
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.

	ex := ch.Rd.Read(ch.Input)
	res, err := PostfixToInfix(ex)
	if err != nil {
		fmt.Println("Недопустимый ввод")
		return err
	}
	if ch.Output == "default" {
		fmt.Println(res)
	} else {
		ch.Wr.Write(res, ch.Output)
	}
	return nil
}
