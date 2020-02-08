package workergroup

import "sync"

type multiChannelWorkerGroup struct {
	cs           []chan Job
	currentIndex int
	workerCount  int
	bufferSize   int
	locker       sync.Mutex
}

func (worker *multiChannelWorkerGroup) Name() string {
	return "MultiChannelWorkerGroup"
}

func (worker *multiChannelWorkerGroup) Add(job Job) {
	worker.locker.Lock()
	worker.cs[worker.currentIndex] <- job
	worker.currentIndex = (worker.currentIndex + 1) % worker.workerCount
	worker.locker.Unlock()
}

func (worker *multiChannelWorkerGroup) Start() {
	for count := 0; count < worker.workerCount; count++ {
		go consume(count, worker.cs[count])
	}
}

func (worker *multiChannelWorkerGroup) Stop() {
	for _, c := range worker.cs {
		close(c)
	}
}
func newMultiChannelWorkerGroup(workerCount int, bufferSize int) *multiChannelWorkerGroup {
	w := new(multiChannelWorkerGroup)
	w.locker = sync.Mutex{}
	w.currentIndex = 0
	w.workerCount = workerCount
	w.bufferSize = bufferSize
	chans := make([]chan Job, workerCount)
	for i := range chans {
		chans[i] = make(chan Job)
	}
	w.cs = chans
	return w
}
