package LC_Go

//// 方法一：位运算
func singleNumber(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}

	return ret
}

//// 方法二：和之差
func singleNumber2(nums []int) int {
	m := make(map[int]bool)
	setSum := 0
	totalSum := 0
	for _, num := range nums {
		_, ok := m[num]
		if !ok {
			setSum += num
			m[num] = true
		}
		totalSum += num
	}

	return 2*setSum - totalSum
}
