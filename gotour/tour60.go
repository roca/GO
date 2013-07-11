package main

//import (
//	"io"
//	"os"
//	"strings"
//)

//type rot13Reader struct {
//	r io.Reader
//}

//func rot13(b byte) byte {
//	var a, z byte
//	switch {
//	case 'a' <= b && b <= 'z':
//		a, z = 'a', 'z'
//	case 'A' <= b && b <= 'Z':
//		a, z = 'A', 'Z'
//	default:
//		return b
//	}
//	return (b-a+13)%(z-a+1) + a
//}

//func (rdr *rot13Reader) Read(b []byte) (n int, err error) {

//	n, err = rdr.r.Read(b)

//	for i, _ := range b {
//		b[i] = rot13(b[i])

//	}

//	return n, err

//}

//func main() {
//	s := strings.NewReader(
//		"Lbh penpxrq gur pbqr!")
//	r := rot13Reader{s}
//	io.Copy(os.Stdout, &r)
//}
