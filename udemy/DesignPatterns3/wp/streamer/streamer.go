package streamer

type ProcessingMessage struct {
	ID         int
	Success    bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct{}

type Processor struct {}


func New(jobQueue chan VideoProcessingJob, maxWorkers int) *VideoDispatcher {
	return &VideoDispatcher{

	}
}
