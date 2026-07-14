package main

import (
	"fmt"
	"wp/streamer"
)

func main() {
	// Define number of workers and jobs
	const numJobs = 4
	const numWorkers = 4

	// Create channels for work and results
	notifyChan := make(chan streamer.ProcessingMessage, numJobs)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, numJobs)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, numWorkers)

	// Start the worker pool.
	wp.Run()
	fmt.Println("Worker pool started. Press any key to continue.")
	fmt.Scanln()

	// Create a video that converts mp4 to web ready format.
	video1 := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// Create a second video that should fail
	video2 := wp.NewVideo(2, "./input/bad.txt", "./output", "mp4", notifyChan, nil)

	// Create a video to send to the worker pool.
	ops := &streamer.VideoOptions{
		RenameOutput: true,
		SegmentDuration: 10,
		MaxRate1080p:   "12000k",
		MaxRate720p: "600k",
		MaxRate480p: "400k",
	}
	// create a third video that converts mp4 to hls format.
	video3 := wp.NewVideo(3, "./input/puppy2.mp4", "./output", "hls", notifyChan, ops)

	// Create a forth video that converts mp4 to web ready format.
	video4 := wp.NewVideo(4, "./input/puppy2.mp4", "./output", "mp4", notifyChan, nil)


	// Send the videos to the worker pool.
	videoQueue <- streamer.VideoProcessingJob{Video: video1}
	videoQueue <- streamer.VideoProcessingJob{Video: video2}
	videoQueue <- streamer.VideoProcessingJob{Video: video3}
	videoQueue <- streamer.VideoProcessingJob{Video: video4}

	// Print out the results
	for i := 1; i <= numJobs; i++ {
		msg := <-notifyChan
		fmt.Println("i:", i, "msg:", msg)
	}

	fmt.Println("Done")

}
