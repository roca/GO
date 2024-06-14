package streamer

import (
	"fmt"
	"os/exec"

	"github.com/xfrr/goffmpeg/transcoder"
)

type Encoder interface {
	EncodeToMP4(v *Video, baseFileName string) error
	EncodeToHLS(v *Video, baseFileName string) error
}

type VideoEncoder struct{}

func (ve *VideoEncoder) EncodeToMP4(v *Video, baseFileName string) error {
	// Create a transcoder.
	trans := new(transcoder.Transcoder)

	// Build the output path
	outputPath := fmt.Sprintf("%s/%s.mp4", v.OutputDir, baseFileName)

	// Initialize the transcoder
	err := trans.Initialize(v.InputFile, outputPath)
	if err != nil {
		return err
	}

	// Set codec
	trans.MediaFile().SetVideoCodec("libx264")

	// Start transcoder process
	done := trans.Run(false)

	err = <-done
	if err != nil {
		return err
	}

	return nil
}

func (ve *VideoEncoder) EncodeToHLS(v *Video, baseFileName string) error {
	// Create a channel to get results
	result := make(chan error)

	// Spawn a goroutine to do the encode
	go func(result chan error) {
		ffmpegCmd := exec.Command(
			"ffmpeg",
			"-i", v.InputFile,
			"-map", "0:v:0",
			"-map", "0:a:0",
		)
	}(result)

	// Listen to the result channel
	err := <-result
	if err != nil {
		return err
	}
	
	// Return the result
	return nil
}
