package LC_Go

/*
	stack + map
	map:
		')': '(',
		']': '[',
		'}': '{',
*/
func isValid(s string) bool {
	stk := make([]rune, 0)

	m := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stk = append(stk, v)
		} else {
			if len(stk) == 0 || stk[len(stk)-1] != m[v] {
				return false
			} else {
				// stk.pop()
				stk = stk[:len(stk)-1]
			}
		}
	}

	return len(stk) == 0
}