package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="content=type" content="text/html; charset=utf-8">
<title>Markdown Preview</title>
</head>
<body>
`

	footer = `
</body>
</html>
`
)

func main() {
	filename := flag.String("file", "", "Markdown file to preview")
	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func run(filename string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	htmlData, err := parseContent(input)
	if err != nil {
		return err
	}

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))
	fmt.Println(outName)

	return saveHTML(outName, htmlData)
}

func parseContent(input []byte) ([]byte, error) {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer

	_, err := buffer.WriteString(header)
	if err != nil {
		return nil, err
	}
	_, err = buffer.Write(body)
	if err != nil {
		return nil, err
	}
	buffer.WriteString(footer)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

func saveHTML(outFname string, data []byte) error {
	return ioutil.WriteFile(outFname, data, 0644)
}
