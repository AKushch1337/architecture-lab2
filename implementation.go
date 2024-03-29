package lab2

import (
	"errors"
	"fmt"
	"strings"
)

// Type Stack is a stack of strings used for postfix to prefix conversion.
type Stack []string

// PostfixToPrefix converts a postfix expression to prefix notation.
func PostfixToPrefix(postfix string) (string, error) {
	if len(postfix) == 0 {
		return "", errors.New("expression is invalid")
	}
	stack := Stack{}
	tokens := strings.Split(postfix, " ")
	for _, token := range tokens {
		switch {
		case IsOperator(token):
			operand2, err := stack.Pop()
			operand1, err := stack.Pop()
			if err != nil {
				return "", errors.New("not enough operands for operator")
			}
			stack.Push(fmt.Sprintf("%s %s %s", token, operand1, operand2))
		default:
			stack.Push(token)
		}
	}
	result, err := stack.Pop()
	if err != nil || len(stack) > 0 {
		return "", errors.New("expression is invalid")
	}
	return result, nil

}

// IsOperator returns true if the given token is a supported operator.
// The supported operators are: +, -, *, /, and ^.
func IsOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/", "^":
		return true
	default:
		return false
	}
}

// Push adds a new string onto the top of the stack.
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop removes and returns the string from the top of the stack.
func (s *Stack) Pop() (string, error) {
	index := len(*s) - 1
	if index < 0 {
		return "", errors.New("stack is empty")
	}
	val := (*s)[index]
	*s = (*s)[:index]
	return val, nil
}
