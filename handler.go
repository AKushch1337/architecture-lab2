package lab2

import (
	"bytes"
	"io"
)

// The ComputeHandler struct defines an input and output for computing expressions.
type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

// The Compute function reads and computes an expression, then writes the result to the output writer.
func (ch *ComputeHandler) Compute() error {
	inputBuffer := make([]byte, 256)
	_, err := ch.Input.Read(inputBuffer)
	if err != nil {
		return err
	}

	trimmedInput := bytes.Trim(inputBuffer, "\x00")

	prefixExpression, err := PostfixToPrefix(string(trimmedInput))
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(prefixExpression))
	if err != nil {
		return err
	}

	return nil
}
