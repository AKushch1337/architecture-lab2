package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/AKushch1337/architecture-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	readFile        = flag.String("f", "", "File that contains the input expression")
	writeFile       = flag.String("o", "", "File that will contain the result")
)

func main() {
	flag.Parse()

	if *inputExpression == "" && *readFile == "" {
		fmt.Fprintf(os.Stderr, "Either input expression or input file must be provided\n")
		os.Exit(1)
	}

	if *inputExpression != "" && *readFile != "" {
		fmt.Fprintf(os.Stderr, "Cannot provide both input expression and input file\n")
		os.Exit(1)
	}

	var reader io.Reader
	if *inputExpression != "" {
		reader = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*readFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open input file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	var writer io.Writer
	if *writeFile != "" {
		file, err := os.Create(*writeFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create output file: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		writer = file
	} else {
		writer = os.Stdout
	}

	handler := &lab2.ComputeHandler{Input: reader, Output: writer}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to compute expression: %s\n", err)
		os.Exit(1)
	}

}
