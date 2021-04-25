package LC_Go

func isHappy(n int) bool {
	m := make(map[int]struct{})
	for {
		sum := 0
		for n > 0 {
			num := n % 10
			sum += num * num
			n = n / 10
		}
		if sum == 1 {
			return true
		}
		_, ok := m[sum]
		if ok {
			return false
		}
		m[sum] = struct{}{}
		n = sum
	}
}
