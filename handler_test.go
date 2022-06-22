package lab2

import (
	"fmt"
	"io"
	"os"
	"strings"
	//. "gopkg.in/check.v1"
)

func ExampleTestHandler() {
	handler := &ComputeHandler{
		Wr:     writer,
		Rd:     reader,
		Input:  "4 2 - 3 *",
		Output: "default",
	}
	handler.Compute()
	//Output:
	// (4-2)*3
}

func ExampleTestHandler1() {
	handler := &ComputeHandler{
		Wr:     writer,
		Rd:     reader,
		Input:  "a a a",
		Output: "default",
	}
	handler.Compute()
	//Output:
	// Недопустимые символы
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
