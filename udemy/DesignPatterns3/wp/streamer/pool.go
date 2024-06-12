package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor Processor
}

// type videoWorker
// See https://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
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
