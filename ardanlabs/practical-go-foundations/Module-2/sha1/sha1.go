package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	fmt.Println(SHA1Sig("http.log.gz"))
	fmt.Println(SHA1Sig("sha1.go"))

}

// SHA1Sig returns SHA1 signature of uncompressed file
// Exercise: Decompress onl;y if file name ends with ".gz"
// cat http.log.gz | gunzip | sha1sum
func SHA1Sig(fileName string) (string, error) {

	// cat http.log.gz
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var r io.Reader = file

	if strings.HasSuffix(fileName, ".gz") {
		// | gunzip
		gz, err := gzip.NewReader(file)
		if err != nil {
			return "", fmt.Errorf("%q - gzip.NewReader(): %w", fileName, err)
		}
		defer gz.Close()
		r = gz
	}

	// | sha1sum
	w := sha1.New()
	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - io.Copy(): %w", fileName, err)
	}

	sig := w.Sum(nil)
	return fmt.Sprintf("%x", sig), nil
}
