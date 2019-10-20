package arcade

func reverseInParentheses(s string) string {
	stack, tmp := make([]rune, 0, len(s)), make([]rune, 0, len(s))
	for _, c := range s {
		if c == ')' {
			i := len(stack) - 1
			tmp = tmp[:0]
			for ; stack[i] != '('; i-- {
				tmp = append(tmp, stack[i])
			}
			stack = append(stack[:i], tmp...)
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
