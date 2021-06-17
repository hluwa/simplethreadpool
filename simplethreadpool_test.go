package simplethreadpool

import (
	"sync"
	"testing"
	"time"
)

func makeFunc(_ int) func() {
	return func() {
		time.Sleep(time.Second)
		panic("111")
	}
}

func TestSimpleThreadPool_Sync(t *testing.T) {
	pool := NewSimpleThreadPool(20)
	s1 := time.Now().Second()
	for i := 0; i < 100; i++ {
		pool.Put(makeFunc(i))
	}

	var mu sync.Mutex
	excepted := 0
	pool.OnException(func(index int, workFunc func(), panic interface{}) {
		mu.Lock()
		defer mu.Unlock()
		excepted += 1
	})
	pool.Sync()
	s2 := time.Now().Second()
	if s2-s1 != 5 {
		t.Fatal("more time than expected")
	}
	if excepted != 100 {
		t.Fatal("less panic than expected")
	}
}
