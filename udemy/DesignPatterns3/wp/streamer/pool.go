package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor Processor
}

// type videoWorker
// See http://tleyden.github.io/blog/2013/11/23/understanding-channel-in-go/
type videoWorker struct {
	id        int
	jobQueue  chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// newVideoWorker 
func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	return videoWorker{
		id: id,
		jobQueue: make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// start()

// Run()

// dispatch()

// processVideoJob()
