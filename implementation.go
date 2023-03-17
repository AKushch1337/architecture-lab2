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
	stack := Stack{}
	tokens := strings.Split(postfix, " ")
	for _, token := range tokens {
		if isOperator(token) {
			operand2, err := stack.pop()
			if err != nil {
				return "", errors.New("not enough operands for operator")
			}
			operand1, err := stack.pop()
			if err != nil {
				return "", errors.New("not enough operands for operator")
			}
			stack.push(fmt.Sprintf("%s %s %s", token, operand1, operand2))
		} else {
			stack.push(token)
		}
	}
	result, err := stack.pop()
	if err != nil {
		return "", errors.New("expression is invalid")
	}
	if len(stack) > 0 {
		return "", errors.New("expression is invalid")
	}
	return result, nil
}

// isOperator returns true if the given token is a supported operator.
// The supported operators are: +, -, *, /, and ^.
func isOperator(token string) bool {
	switch token {
	case "+", "-", "*", "/", "^":
		return true
	}
	return false
}

// push adds a new string onto the top of the stack.
func (s *Stack) push(str string) {
	*s = append(*s, str)
}

// pop removes and returns the string from the top of the stack.
func (s *Stack) pop() (string, error) {
	index := len(*s) - 1
	if index < 0 {
		return "", errors.New("stack is empty")
	}
	val := (*s)[index]
	*s = (*s)[:index]
	return val, nil
}
