package main

import (
	//"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/devteam/lab-2"
)

var (
	inputExpression = flag.Bool("e", false, "Expression to compute")
	inputFile       = flag.Bool("f", false, "File with expression")
	outputResult    = flag.Bool("o", true, "Write result to file")
)

type Reader interface {
	Read(str string)
}

type Writer interface {
	Write(str string, outFile string)
}

type Readall struct{}

type WriteAll struct{}

func (rd Readall) Read(inputStr string) string {
	if strings.Contains(inputStr, ".txt") {
		file, err := os.Open(inputStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		data := make([]byte, 64)
		str := make([]string, 0)

		for {
			n, err := file.Read(data)
			if err == io.EOF { // если конец файла
				break // выходим из цикла
			}
			str = append(str, string(data[:n]))
		}
		result := strings.Join(str, " ")
		return result
	}
	return inputStr
}

func (wr WriteAll) Write(str string, outFile string) {
	f, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte(str))
}

var writer = WriteAll{}
var reader = Readall{}

func main() {
	flag.Parse()
	in := os.Args[2]

	if *inputExpression && *outputResult {

		handler := &lab2.ComputeHandler{
			Wr:     writer,
			Rd:     reader,
			Input:  in,
			Output: "default",
		}
		handler.Compute()
		fmt.Println(inputExpression)
	} else if *inputExpression && !*outputResult {
		handler := &lab2.ComputeHandler{
			Wr:     writer,
			Rd:     reader,
			Input:  os.Args[2],
			Output: os.Args[4],
		}
		handler.Compute()
	} else if *inputFile && !*outputResult {
		handler := &lab2.ComputeHandler{
			Wr:     writer,
			Rd:     reader,
			Input:  os.Args[2],
			Output: "default",
		}
		fmt.Println(handler.Compute())
	} else if *inputFile && *outputResult {
		//out := os.Args[4]
		handler := &lab2.ComputeHandler{
			Wr:     writer,
			Rd:     reader,
			Input:  os.Args[2],
			Output: os.Args[4],
		}
		handler.Compute()
	} else if !*inputExpression {
		//err := errors.New("Input expretion")
	}

}
