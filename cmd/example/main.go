package main

import (
	//"errors"
	"flag"
	"fmt"

	"os"

	lab2 "github.com/devteam/lab-2"
)

var (
	inputExpression = flag.Bool("e", true, "Expression to compute")
	inputFile       = flag.Bool("f", true, "File with expression")
	outputResult    = flag.Bool("o", true, "Write result to file")
)

func main() {
	flag.Parse()
	in := os.Args[1]
	out := os.Args[2]
	if *inputExpression || *inputFile && *outputResult {
		handler := &lab2.ComputeHandler{
			Input:  in,
			Output: out,
		}
		fmt.Println(handler.Compute())
	} else if *inputExpression || *inputFile && !*outputResult {
		handler := &lab2.ComputeHandler{
			Input:  os.Args[1],
			Output: "default",
		}
		fmt.Println(handler.Compute())
	} else if !*inputExpression {
		//err := errors.New("Input expretion")
	}

	// TODO: Change this to accept input from the command line arguments as described in the task and
	//       output the results using the ComputeHandler instance.
	//       handler := &lab2.ComputeHandler{
	//           Input: {construct io.Reader according the command line parameters},
	//           Output: {construct io.Writer according the command line parameters},
	//       }
	//       err := handler.Compute()

}
