package main

import (
	//"errors"
	"flag"
	"fmt"

	"os"

	lab2 "github.com/devteam/lab-2"
)

var (
	inputExpression = flag.Bool("e", false, "Expression to compute")
	inputFile       = flag.Bool("f", false, "File with expression")
	outputResult    = flag.Bool("o", true, "Write result to file")
)

func main() {
	flag.Parse()
	in := os.Args[2]

	if *inputExpression && *outputResult {

		handler := &lab2.ComputeHandler{
			Input:  in,
			Output: os.Args[4],
		}
		handler.Compute()
		fmt.Println(inputExpression)
	} else if *inputExpression && !*outputResult {
		handler := &lab2.ComputeHandler{
			Input:  os.Args[2],
			Output: "default",
		}
		handler.Compute()
	} else if *inputFile && !*outputResult {
		handler := &lab2.ComputeHandler{
			Input:  os.Args[2],
			Output: os.Args[4],
		}
		fmt.Println(handler.Compute())
	} else if *inputFile && *outputResult {
		//out := os.Args[4]
		handler := &lab2.ComputeHandler{
			Input:  os.Args[2],
			Output: "default",
		}
		handler.Compute()
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
