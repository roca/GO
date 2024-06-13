package streamer

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

// type videoWorker
// See https://tleyden.github.io/blog/2013/11/23/understanding-chan-chans-in-go/
type videoWorker struct {
	id         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// newVideoWorker
func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// start()
func (w videoWorker) start() {
	go func() {
		for {
			// Add jobQueue to the worker pool
			w.workerPool <- w.jobQueue

			// Wait for a job to come in.
			job := <-w.jobQueue

			// Process the job
			w.processVideoJob(job.Video)
		}
	}()
}

// Run()

// dispatch()
func (vd *VideoDispatcher) dispatch() {
	for {
		// Wait for a job to come in.
		job := <-vd.jobQueue

		go func() {
			// Wait for a worker to be available.
			workerJobQueue := <-vd.WorkerPool

			// Send the job to the worker.
			workerJobQueue <- job
		}()

	}
}

// processVideoJob()
func (w videoWorker) processVideoJob(video Video) {
	video.encode()
}
