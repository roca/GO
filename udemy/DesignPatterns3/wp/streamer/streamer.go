package streamer

type ProcessingMessage struct {
	ID         int
	Success    bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct{
	Video Video
}

type Processor struct{}

type Video struct {
	ID           int
	InputFile    string
	OutputDir    string
	EncodingType string
	NotifyChan   chan ProcessingMessage
	// Options *VideoOptions
	Encoder Processor
}

func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	workPool := make(chan chan VideoProcessingJob, maxWorkers)

	// TODO: implement processor logic
	p := Processor{}

	return &VideoDispatcher{
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		WorkerPool: workPool,
		Processor:  p,
	}
}
