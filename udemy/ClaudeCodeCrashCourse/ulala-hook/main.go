package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/wav"
)

func main() {
	// Open the audio file
	f, err := os.Open("/Users/romel.campbell/GOCODE/udemy/ClaudeCodeCrashCourse/ulala-hook/ulala.wav")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Decode the WAV file
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	// Initialize the speaker with the sample rate of the audio file
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Playing ulala.wav...")

	// Play the audio and wait for it to finish
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	fmt.Println("Playback complete!")
}
