package main

import (
	"github.com/golang-collections/collections/stack"
)

func isValid(s string) bool {
	length := len(s)

	if length == 0 {
		return true
	}

	if length%2 == 1 {
		return false
	}

	var sk = stack.New()

	for i := 0; i < length; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			sk.Push(s[i])
			continue
		}

		if sk.Len() == 0 {
			return false
		}

		cur := sk.Peek()
		sk.Pop()
		if s[i] == ')' && cur != '(' {
			return false
		}
		if s[i] == '}' && cur != '{' {
			return false
		}
		if s[i] == ']' && cur != '[' {
			return false
		}
	}

	return sk.Len() == 0
}
