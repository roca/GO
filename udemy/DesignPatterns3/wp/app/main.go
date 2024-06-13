package main

import (
	"fmt"
	"wp/streamer"
)

func main() {
	// Define number of workers and jobs
	const numJobs = 4
	const numWorkers = 1

	// Create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, numWorkers)
	fmt.Println("wp:", wp)

	// Start the worker pool.
	wp.Run()

	// Create a video to send to the worker pool.
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// Send the videos to the worker pool.
	videoQueue <- streamer.VideoProcessingJob{Video: video}


	// Print out the results
	

}
