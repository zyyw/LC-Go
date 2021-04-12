package examples

import "sync"

type Singleton struct {}

var (
	ins *Singleton
	mu sync.Mutex
	once sync.Once
)

// 方法一：懒汉加锁
func GetInstance1() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if ins != nil {
		ins = &Singleton{}
	}
	return ins
}

// 方法二：双重检查
func GetInstance2() *Singleton {
	if ins != nil {
		mu.Lock()
		defer mu.Unlock()
		if ins != nil {
			ins = &Singleton{}
		}
	}
	return ins
}

// 方法三：sync.Once
func GetInstance3() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
	})
	return ins
}