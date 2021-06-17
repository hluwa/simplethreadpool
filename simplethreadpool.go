package simplethreadpool

type SimpleThreadPool struct {
	workList    []func()
	workerCount int
}

func (pool *SimpleThreadPool) Sync() {
	poolChannel := make(chan int, pool.workerCount)
	waitChannel := make(chan int)
	for i, f := range pool.workList {
		go func(i int, f func()) {
			defer func() { waitChannel <- i }()
			poolChannel <- i
			defer func() { <-poolChannel }()
			if f != nil {
				f()
			}
		}(i, f)
	}

	for i := 0; i < len(pool.workList); i++ {
		<-waitChannel
	}
}

func (pool *SimpleThreadPool) Put(f func()) {
	pool.workList = append(pool.workList, f)
}

func NewSimpleThreadPool(workerCount int) *SimpleThreadPool {
	return &SimpleThreadPool{
		workerCount: workerCount,
	}
}
