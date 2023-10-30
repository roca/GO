package main

import (
	"fmt"
	"log"
	"project01/basic"
	"time"
)

type logWriter struct{}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Printf(time.Now().Format("15:04:04") + " " + string(bytes))
}

func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// basic.BasicHello()
	// basic.BasicUser()
	// basic.ProtoToJsonUser()
	// basic.JsonToProtoUser()
	//basic.BasicUserGroup()
	//jobsearch.JobSearchSoftware()
	//jobsearch.JobSearchCandidate()
	// basic.BasicUnMarshalAynKnown()
	//basic. WriteToFileSample()
	//basic.ReadFromFileSample()
	//basic. WriteToFileJSONSample()
	//basic.ReadFromFileJSONSample()

	// basic.BasicWriteUserContentV1()
	// basic.BasicReadUserContentV1()

	//basic.BasicWriteUserContentV2()
	//basic.BasicReadUserContentV2()

	//basic.BasicWriteUserContentV3()
	//basic.BasicReadUserContentV3()

	// basic.BasicWriteUserContentV4()
	basic.BasicReadUserContentV4()
}
