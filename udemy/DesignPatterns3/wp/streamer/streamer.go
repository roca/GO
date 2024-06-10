package streamer

type ProcessingMessage struct {
	ID         int
	Success    bool
	Message    string
	OutputFile string
}

type VideoProcessingJob struct{}
