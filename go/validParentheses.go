package main

type stack []byte

func (s stack) Push(c byte) stack {
	return append(s, c)
}

func (s stack) Pop() (stack, byte) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func isValid(s string) bool {
	length := len(s)

	if length == 0 {
		return true
	}

	if length%2 == 1 {
		return false
	}

	var sk stack
	for i := 0; i < length; i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			sk = sk.Push(s[i])
			continue
		}

		if len(sk) == 0 {
			return false
		}

		skk, cur := sk.Pop()
		if s[i] == ')' && cur != '(' {
			return false
		}
		if s[i] == '}' && cur != '{' {
			return false
		}
		if s[i] == ']' && cur != '[' {
			return false
		}
		sk = skk
	}

	return len(sk) == 0
}
