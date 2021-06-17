package simplethreadpool

type SimpleThreadPool struct {
	workList    []func()
	workerCount int

	exceptionHandler func(index int, workFunc func(), panic interface{})
}

func (pool *SimpleThreadPool) Sync() {
	poolChannel := make(chan int, pool.workerCount)
	waitChannel := make(chan int)
	for i, f := range pool.workList {
		go func(i int, f func()) {
			defer func() {
				if e := recover(); e != nil && pool.exceptionHandler != nil {
					pool.exceptionHandler(i, f, e)
				}
			}()
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

func (pool *SimpleThreadPool) OnException(handler func(index int, workFunc func(), panic interface{})) {
	pool.exceptionHandler = handler
}

func NewSimpleThreadPool(workerCount int) *SimpleThreadPool {
	return &SimpleThreadPool{
		workerCount: workerCount,
	}
}
