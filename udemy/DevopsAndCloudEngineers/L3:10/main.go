package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type MySlowReader struct {
	contents string
	pos      int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.contents) {
		n = copy(p, m.contents[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {

	mySlowReaderInstance := &MySlowReader{contents: "hello world!"}

	out, err := ioutil.ReadAll(mySlowReaderInstance)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output: %s\n", out)

	// if response.StatusCode != 200 {
	// 	fmt.Printf("Invalid output (HTTP Code %d): %sn", response.StatusCode, out)
	// 	os.Exit(1)
	// }

	// var words Words
	// if err := json.Unmarshal(out, &words); err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("JSON Parsed\nPage: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))
}
