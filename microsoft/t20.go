package microsoft

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append([]rune{c}, stack...)
		} else {
			if len(stack) == 0 || (c == ')' && stack[0] != '(') || (c == ']' && stack[0] != '[') || (c == '}' && stack[0] != '{') {
				return false
			}
			stack = stack[1:]
		}
	}
	return len(stack) == 0
}
