package workergroup

const channelBufferSize int = 10

type Job interface {
	Do()
}

type WorkerGroup interface {
	Name() string
	Add(job Job)
	Start()
	Stop()
}

func consume(num int, jobchan <-chan Job) {
	start := 0
	for job := range jobchan {
		start++
		job.Do()
		//fmt.Println("Done in ", num)
	}
}

var NewWorkerGroup = func(workerCount int) WorkerGroup {
	return newWorkerGroup(workerCount, channelBufferSize)
}

var NewMultiChannelWorkerGroup = func(workerCount int) WorkerGroup {
	return newMultiChannelWorkerGroup(workerCount, channelBufferSize)
}

var NewNoChannelWorkerGroup = func(workerCount int) WorkerGroup {
	return newNoChannelWorkerGroup()
}
