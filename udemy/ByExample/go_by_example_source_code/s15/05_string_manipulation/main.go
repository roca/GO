// File name: ...\s15\05_string_manipulation\main.go
// Course Name: Go (Golang) Programming by Example (by Kam Hojati)

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	testCase1() //call testCase1() to testCase5()
}

func testCase1() {
	fmt.Println(strings.Compare("at", "at"))
	fmt.Println(strings.Compare("at", "amber"))
	fmt.Println(strings.Compare("amber", "at"))

	fmt.Println(bytes.Compare([]byte("amber"), []byte("at")))

	fmt.Println(strings.Contains("team player", "play"))
	fmt.Println(bytes.Contains([]byte("team player"), []byte("play")))

	fmt.Println(strings.Count("interesting", "e"))
	fmt.Println(strings.Count("redredred", "re"))
	fmt.Println(strings.Count("aaaaaaaa", "aa"))
}
func testCase2() {
	fmt.Printf("Fields are: %q \n", strings.Fields("  foo bar  baz   "))

	fmt.Printf("Fields are: %q \n", bytes.Fields([]byte("  foo bar  baz   ")))

	fmt.Println(strings.HasPrefix("Golang", "Go"))

	fmt.Println(strings.HasSuffix("friend", "end"))

	fmt.Println(strings.Index("booklet", "let"))
	fmt.Println(strings.Index("booklet", "lets"))
}
func testCase3() {
	s := []string{"Mon", "Tue", "Wed"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Println(strings.Index("go leafs go", "go"))
	fmt.Println(strings.LastIndex("go leafs go", "go"))

	shiftOne := func(r rune) rune {
		return r + 1
	}
	fmt.Println(strings.Map(shiftOne, "ABCDabcd I al going to be converted."))

	fmt.Println(strings.Repeat("la", 3))
}
func testCase4() {
	fmt.Println(strings.Replace("blah blah blah", "h", "ck", 2))
	fmt.Println(strings.Replace("blah blah blah", "h", "ck", 4))
	fmt.Println(strings.Replace("blah blah blah", "blah", "ye", -1))

	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" I am ", ""))
	fmt.Printf("%q\n", strings.Split("", "Rob Travis"))
}
func testCase5() {
	fmt.Println(strings.ToLower("GoLang"))

	fmt.Println(strings.ToUpper("GoLang"))

	fmt.Println(strings.Trim("¿¿¿??Cómo Estás?¿??", "¿?"))
	fmt.Println(strings.Trim("   Cómo Estás  ", " "))

	fmt.Println(strings.TrimLeft("¿¿¿??Cómo Estás?¿??", "¿?"))

	fmt.Println(strings.TrimRight("¿¿¿??Cómo Estás?¿??", "¿?"))

	bnry := []byte("ABCD")
	str := string([]byte{'A', 'B', 'C', 'D'})
	fmt.Println(bnry, str)
}
