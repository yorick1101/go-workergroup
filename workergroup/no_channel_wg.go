package workergroup

type noChannelWorkerGroup struct {
}

func (worker *noChannelWorkerGroup) Name() string {
	return "NoWorkerGroup"
}

func (worker *noChannelWorkerGroup) Add(job Job) {
	go job.Do()
}

func (worker *noChannelWorkerGroup) Start() {

}

func (worker *noChannelWorkerGroup) Stop() {

}
func newNoChannelWorkerGroup() *noChannelWorkerGroup {
	w := new(noChannelWorkerGroup)
	return w
}
