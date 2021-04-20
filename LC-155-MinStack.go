package LC_Go

import "fmt"

type MinStack struct {
	valStk []int
	minStk []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var minStack MinStack
	return minStack
}

func (this *MinStack) Push(val int) {
	this.valStk = append(this.valStk, val)
	if len(this.minStk) == 0 || val < this.minStk[len(this.minStk)-1] {
		this.minStk = append(this.minStk, val)
	} else {
		this.minStk = append(this.minStk, this.minStk[len(this.minStk)-1])
	}
}

func (this *MinStack) Pop() {
	if len(this.valStk) == 0 {
		fmt.Println("MinStack is empty")
		return
	}
	// valStk.pop()
	this.valStk = this.valStk[:len(this.valStk)-1]
	// minStk.pop()
	this.minStk = this.minStk[:len(this.minStk)-1]
}

func (this *MinStack) Top() int {
	if len(this.valStk) == 0 {
		fmt.Println("MinStack is empty")
		return -1
	}
	return this.valStk[len(this.valStk)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStk) == 0 {
		fmt.Println("MinStack is empty")
		return -1
	}
	return this.minStk[len(this.minStk)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
