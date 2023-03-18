package lab2

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestComputeHandler_Compute(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "valid expression with 2 operands",
			input:  "2 3 +",
			output: "+ 2 3",
			err:    nil,
		},
		{
			name:   "valid expression with 3 operands",
			input:  "5 6 + 7 *",
			output: "* + 5 6 7",
			err:    nil,
		},
		{
			name:   "invalid expression with not enough operands",
			input:  "1 +",
			output: "",
			err:    errors.New("not enough operands for operator"),
		},
		{
			name:   "invalid expression with wrong operator",
			input:  "1 2 =",
			output: "",
			err:    errors.New("expression is invalid"),
		},
	}

	// Run tests.
	for _, testCases := range tests {
		t.Run(testCases.name, func(t *testing.T) {
			input := strings.NewReader(testCases.input)
			output := new(bytes.Buffer)

			handler := ComputeHandler{
				Input:  input,
				Output: output,
			}

			err := handler.Compute()

			assert.Equal(t, testCases.err, err)

			outputString := output.String()

			assert.Equal(t, testCases.output, outputString)
		})
	}
}
