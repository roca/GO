package streamer

import "fmt"

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
	fmt.Println("newVideoWorker(): creating video worker id:", id)
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// start()
func (w videoWorker) start() {
	fmt.Println("w.start(): starting worker id:", w.id)
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
func (vd *VideoDispatcher) Run() {
	fmt.Println("vd.Run(): starting worker pool by running workers")
	for i := 0; i < vd.maxWorkers; i++ {
		fmt.Println("vd.Run(): starting worker id:", i+1)
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}
	go vd.dispatch()
}

// dispatch()
func (vd *VideoDispatcher) dispatch() {
	for {
		// Wait for a job to come in.
		job := <-vd.jobQueue
		fmt.Println("vd.dispatch(): sending job", job.Video.ID, "to worker job queue")
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
	fmt.Println("w.processVideoJob(): staring encode on video", video.ID)
	video.encode()
}
