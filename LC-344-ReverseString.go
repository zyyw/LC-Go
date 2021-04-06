package LC_Go

func reverseString(s []byte)  {
	if len(s) <= 1 {
		return
	}
	start := 0
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}