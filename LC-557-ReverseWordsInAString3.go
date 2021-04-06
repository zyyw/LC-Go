package LC_Go

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	runeArr := []rune(s)
	start := 0
	end := 0
	for end < len(runeArr) {
		for end < len(runeArr) && runeArr[end] != ' ' {
			end++
		}
		reverse(runeArr, start, end - 1)
		end++
		start = end
	}

	return string(runeArr)
}

func reverse(r []rune, start int, end int) {
	if len(r) == 0 || start >= end {
		return
	}
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}