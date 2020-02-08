package workergroup

type singleChannelWorkerGroup struct {
	c1          chan Job
	workerCount int
}

func (worker *singleChannelWorkerGroup) Name() string {
	return "SingleChannelWorkerGroup"
}

func (worker *singleChannelWorkerGroup) Add(job Job) {
	worker.c1 <- job
}

func (worker *singleChannelWorkerGroup) Start() {
	for count := 0; count < worker.workerCount; count++ {
		go consume(count, worker.c1)
	}
}

func (worker *singleChannelWorkerGroup) Stop() {
	close(worker.c1)
}

func newWorkerGroup(workerCount int, bufferSize int) *singleChannelWorkerGroup {
	w := new(singleChannelWorkerGroup)
	w.workerCount = workerCount
	w.c1 = make(chan Job, bufferSize)
	return w
}
