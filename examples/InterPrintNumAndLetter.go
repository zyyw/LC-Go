package examples

import (
	"fmt"
	"sync"
)

/*
问题描述：
使用两个 goroutine 交替打印序列，一个 goroutine 打印数字， 另外一 个 goroutine 打印字母， 最终效果如下:
output: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

follow up, output: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ
*/

func PrintNumAndLetter() {
	letter, number := make(chan bool, 1), make(chan bool)
	wg := sync.WaitGroup{}

	// goroutine 1 打印数字
	wg.Add(1)
	go func() {
		val := 1
		for i := 0; i <= 13; i++ {
			select {
			case <-number:
				fmt.Print(val)
				val++
				fmt.Print(val)
				val++
				letter <- true
			}
		}
		wg.Done()
	}()

	// goroutine 2 打印字母
	go func() {
		char := 'A'
		for i := 0; i < 13; i++ {
			select {
			case <-letter:
				fmt.Print(string(char))
				char++
				fmt.Print(string(char))
				char++
				number <- true
			}
		}
	}()

	number <- true
	wg.Wait()
}

// 不用 WaitGroup
func PrintNumAndLetter2() {
	letter, number := make(chan bool, 1), make(chan bool)
	done := make(chan struct{})

	// goroutine 1 打印数字
	go func() {
		val := 1
		for i := 0; i < 13; i++ {
			select {
			case <-number:
				fmt.Print(val)
				val++
				fmt.Print(val)
				val++
				letter <- true
			}
		}
		done <- struct{}{}
	}()

	// goroutine 2 打印字母
	go func() {
		char := 'A'
		for i := 0; i < 13; i++ {
			select {
			case <-letter:
				fmt.Print(string(char))
				char++
				fmt.Print(string(char))
				char++
				number <- true
			}
		}
	}()

	number <- true
	<-done
}
