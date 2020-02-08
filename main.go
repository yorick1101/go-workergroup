package main

import (
	"fmt"
	"go-kafka/workergroup"
	"math"
	"sync"
	"time"
)

type TestJob struct {
	number int
	lock   *sync.WaitGroup
}

func (t *TestJob) Do() {
	//fmt.Println("Do Job", t.number)
	time.Sleep(10)
	t.lock.Done()
}

func main() {

	var jobCounts = 1000
	var exponential = 5
	var warmupRun = 10
	var totalRun = 50

	builders := getWorkerGroupBuilders()
	for _, builder := range builders {
		for currentExponential := 1; currentExponential <= exponential; currentExponential++ {
			workers := int(math.Pow(2, float64(currentExponential)))
			wg := builder(workers)
			var totalTime time.Duration = 0

			for run := 0; run < warmupRun; run++ {
				test(wg, jobCounts)
			}

			for run := 0; run < totalRun; run++ {
				starting := time.Now().UTC()
				test(wg, jobCounts)
				ending := time.Now().UTC()
				totalTime += ending.Sub(starting)
			}

			wg.Stop()
			fmt.Println(wg.Name(), " workers:", workers, " time:", totalTime.Milliseconds(), " ", totalTime.Milliseconds()/int64(totalRun))
		}
	}
}

type WorkerGroupBuilder func(workers int) workergroup.WorkerGroup

func getWorkerGroupBuilders() []WorkerGroupBuilder {
	wgs := make([]WorkerGroupBuilder, 0)
	wgs = append(wgs, workergroup.NewMultiChannelWorkerGroup)
	wgs = append(wgs, workergroup.NewWorkerGroup)
	wgs = append(wgs, workergroup.NewNoChannelWorkerGroup)
	return wgs
}

func test(wg workergroup.WorkerGroup, jobsCount int) {
	lock := new(sync.WaitGroup)
	jobs := createJobs(jobsCount, lock)
	wg.Start()
	for _, job := range jobs {
		wg.Add(job)
	}
	lock.Wait()
}

func createJobs(count int, lock *sync.WaitGroup) []*TestJob {
	end := count
	jobs := make([]*TestJob, count)
	for count := 0; count < end; count++ {
		lock.Add(1)
		j := new(TestJob)
		j.number = count
		j.lock = lock
		jobs[count] = j
	}
	return jobs
}
