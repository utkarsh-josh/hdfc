package workerpool

// WorkerPool is a interface to call workerpool methods
type WorkerPool interface {
	Run()
	AddTask(task func())
}

type workerPool struct {
	maxWorker int
	task      chan func()
}

// NewWorkerPool returns workerpool layer handler
func NewWorkerPool(maxWorkerCount int) WorkerPool {
	return &workerPool{
		maxWorker: maxWorkerCount,
		task:      make(chan func()),
	}
}

func (wp *workerPool) Run() {
	for i := 0; i < wp.maxWorker; i++ {
		go func(workerID int) {
			for {
				select {
				case task := <-wp.task:
					task()
				}
			}
		}(i + 1)
	}
}

func (wp *workerPool) AddTask(task func()) {
	wp.task <- task
}
