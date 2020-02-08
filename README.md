# go-workergroup

In java, we have many chances to create a thread pool to share the load of jobs.
This poc is to try to do the samething by go routine.

There are 3 implementations:
1. SingleChannelWorkerGroup: use one channel to produce and retrieve jobs.
2. MultiChannelWorkerGroup: use one channel for each woker and dispatch the job in round robin manner.
3. NoWorkerGroup: No channel at all, execute the job directly by goroutine and not limited by a number of workers.

## Setting
```
var jobCounts = 1000  // jobs to run each run  
var warmupRun = 10    // do warmup before actually calculate the time  
var totalRun = 50       
```
## Result (ms)
```
MultiChannelWorkerGroup  workers: 2  total time:  1425  avg time: 28
MultiChannelWorkerGroup  workers: 4  total time:  754  avg time: 15
MultiChannelWorkerGroup  workers: 8  total time:  392  avg time: 7
MultiChannelWorkerGroup  workers: 16  total time:  238  avg time: 4
MultiChannelWorkerGroup  workers: 32  total time:  120  avg time: 2
SingleChannelWorkerGroup  workers: 2  total time:  1466  avg time: 29
SingleChannelWorkerGroup  workers: 4  total time:  773  avg time: 15
SingleChannelWorkerGroup  workers: 8  total time:  363  avg time: 7
SingleChannelWorkerGroup  workers: 16  total time:  215  avg time: 4
SingleChannelWorkerGroup  workers: 32  total time:  137  avg time: 2
NoWorkerGroup  workers: 2  total time:  90  avg time: 1
NoWorkerGroup  workers: 4  total time:  91  avg time: 1
NoWorkerGroup  workers: 8  total time:  93  avg time: 1
NoWorkerGroup  workers: 16  total time:  88  avg time: 1
NoWorkerGroup  workers: 32  total time:  97  avg time: 1
```
## Conculsion
1. SingleChannelWorkerGroup and MultiChannelWorkerGroup does not have big differences.
2. If it's not specifically like to limit the number of workers, just run with go routine.
