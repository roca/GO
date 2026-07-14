package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

// Event ...
type Event struct {
}

var mesages = []string{
	"Hello World",
	"Hello Serverless",
	"It's a great day today",
	"Yay, I'm learning something new today",
	"On cloud nine!",
	"Over the moon",
	"Shooting for the stars!",
	"On top of the World",
	"World at my feet!",
	"Doing everthing I love!",
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initLoggers(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func handler(event Event) (string, error) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	message := mesages[r.Intn(9)]

	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")

	return fmt.Sprintf(message), nil
}

func main() {
	initLoggers(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	lambda.Start(handler)
}
