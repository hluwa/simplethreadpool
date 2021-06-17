package simplethreadpool

import (
	"fmt"
	"testing"
	"time"
)

func makeFunc(i int) func() {
	return func() {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func TestSimpleThreadPool_Sync(t *testing.T) {
	pool := NewSimpleThreadPool(20)
	s1 := time.Now().Second()
	for i := 0; i < 100; i++ {
		pool.Put(makeFunc(i))
	}
	pool.Sync()
	s2 := time.Now().Second()
	if s2-s1 != 5 {
		t.Fatal("more time than expected")
	}
}
