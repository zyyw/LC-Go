package arrays

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	m := make(map[byte]struct{})
	left, right := 0, 0
	ret := 0
	for left <= right && right < len(s) {
		_, ok := m[s[right]]
		if !ok {
			// s[right] 不在 m 里
			m[s[right]] = struct{}{}
			size := right - left + 1
			if size > ret {
				ret = size
			}
			right++
		} else {
			delete(m, s[left])
			left++
		}
	}

	return ret
}
