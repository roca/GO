package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// LogMrtgApps ...
func LogMrtgApps() {
	start := time.Now()

	OpenFile(FileName)
	defer CloseFile()

	var c = make(chan MrtgType)
	var b = make(chan bool)

	f := []interface{}{logApprovedApplications, logRejectedApplications}

	for i := range f {
		// 'chan<-' send-only channel
		go f[i].(func(chan bool, chan<- MrtgType))(b, c)
	}

	go closeChannels(b, c)

	for val := range c {
		fmt.Println(val)
	}

	fmt.Printf("\n** Number of Approved Applications: %d\n", len(AppprovedIdx))
	fmt.Printf("** Number of Rejected Applications: %d\n", len(RejectedIdx))

	elapsed := time.Since(start)
	fmt.Printf("\n= = = = = = = = = = = = = = = = = = = = = = = = = = =\n")
	fmt.Printf("= LogMrtgApps() - time elapsed: %s \n", elapsed)
	fmt.Printf("= = = = = = = = = = = = = = = = = = = = = = = = = = =\n\n")
}

func logApprovedApplications(b chan bool, c chan<- MrtgType) {

	for _, v := range AppprovedIdx {

		mrtgSummary := MrtgType{
			strconv.FormatUint(uint64(Mrtgs[v].CID), 10),
			strings.ToUpper(Mrtgs[v].FirstName),
			strings.ToUpper(Mrtgs[v].LastName),
			strconv.FormatFloat(float64(Mrtgs[v].Salary), 'g', 1000, 32),
			strconv.FormatFloat(float64(Mrtgs[v].Downpayment), 'g', 1000, 32),
			"APPROVED"}

		c <- mrtgSummary
		WriteToFile(mrtgSummary)
	}
	b <- true
}
func logRejectedApplications(b chan bool, c chan<- MrtgType) {

	for _, v := range RejectedIdx {

		mrtgSummary := MrtgType{
			strconv.FormatUint(uint64(Mrtgs[v].CID), 10),
			strings.ToLower(Mrtgs[v].FirstName),
			strings.ToLower(Mrtgs[v].LastName),
			strconv.FormatFloat(float64(Mrtgs[v].Salary), 'g', 1000, 32),
			strconv.FormatFloat(float64(Mrtgs[v].Downpayment), 'g', 1000, 32),
			"**REJECTED**"}

		c <- mrtgSummary
		WriteToFile(mrtgSummary)
	}
	b <- true
}

func closeChannels(b chan bool, c chan MrtgType) {
	<-b
	<-b
	close(c)
	close(b)
}
