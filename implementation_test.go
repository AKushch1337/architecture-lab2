package lab2

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToPrefix(t *testing.T) {
	tests := []struct {
		name    string
		postfix string
		prefix  string
		err     error
	}{
		{
			name:    "valid expression with 3 operands",
			postfix: "4 9 + 2 *",
			prefix:  "* + 4 9 2",
			err:     nil,
		},
		{
			name:    "valid expression with 4 operands",
			postfix: "4 2 - 3 * 5 +",
			prefix:  "+ * - 4 2 3 5",
			err:     nil,
		},
		{
			name:    "valid expression with 4 operands",
			postfix: "40 3 5 * 6 / +",
			prefix:  "+ 40 / * 3 5 6",
			err:     nil,
		},
		{
			name:    "valid expression with 7 operands",
			postfix: "5 3 2 3 ^ 5 - 7 -3 * + * -",
			prefix:  "- 5 * 3 + - ^ 2 3 5 * 7 -3",
			err:     nil,
		},
		{
			name:    "valid expression with large number of operands",
			postfix: "2 20 * 2 / 3 4 + 3 2 ^ * + 6 - 15 +",
			prefix:  "+ - + / * 2 20 2 * + 3 4 ^ 3 2 6 15",
			err:     nil,
		},
		{
			name:    "invalid expression with too few operands",
			postfix: "2 +",
			prefix:  "",
			err:     errors.New("not enough operands for operator"),
		},
		{
			name:    "invalid expression with too many operands",
			postfix: "2 3 + 4",
			prefix:  "",
			err:     errors.New("expression is invalid"),
		},
		{
			name:    "invalid expression with invalid operator",
			postfix: "2 3 $",
			prefix:  "",
			err:     errors.New("expression is invalid"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := PostfixToPrefix(test.postfix)
			assert.Equal(t, test.prefix, actual)
			assert.Equal(t, test.err, err)
		})
	}
}

func ExamplePostfixToPrefix() {
	res, _ := PostfixToPrefix("2 2 +")
	fmt.Println(res)

	// Output:
	// + 2 2
}
