package jobsearch

import (
	"log"
	"project01/proto/basic/protogen/basic"
	"project01/proto/dummy/protogen/dummy"
	"project01/proto/jobsearch/protogen/jobsearch"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js := jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "The Amazing Proto",
			Platforms: []string{"Mac", "Windows", "Linux"},
		},
	}

	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Software  :", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId: 887,
			ApplicantFullName: "shazam",
			Phone: "555-555-5555",
			Email: "shazam@dc.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("Candidate  :", string(jsonBytes))
}
