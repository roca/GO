package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/abadojack/whatlanggo"
	"github.com/jdkato/prose/chunk"
	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
	"github.com/roca/GO/udemy/DataScienceNLP/files"
	"github.com/rylans/getlang"

	cregex "github.com/mingrammer/commonregex"
)

func main() {
	// languageDetectionExample01()
	//languageDetectionExample02()
	//filesExmaples()
	tokenizationExample02()
}

func tokenizationExample01() { // From Scratch

	// Method 1: strings.Split
	var myText string = "Paul wasn't coding at all"
	tokens := strings.Split(myText, " ")
	fmt.Println(tokens)

	// Method 2: Rule Based (Regex)
	r := regexp.MustCompile(`\w+`)
	tokens2 := r.FindAllString(myText, -1)
	fmt.Println(tokens2)

	// Method 3: Regex + Split
	r2 := regexp.MustCompile(`\s+`)
	tokens3 := r2.Split(myText, -1)
	fmt.Println(tokens3)

	// Method 4: Using pros
	tokenizer := tokenize.NewTreebankWordTokenizer()
	tokens4 := tokenizer.Tokenize(myText)
	for _, tok := range tokens4 {
		fmt.Println(tok)
	}
	fmt.Println(tokens4)
}

func tokenizationExample02() { // Using prose
	myText := "Jesse was going to fish a fish at the bank in London"

	// Tokens
	tokenizer := tokenize.NewTreebankWordTokenizer()
	tokens := tokenizer.Tokenize(myText)
	fmt.Println(tokens)

	// Tags
	postagger := tag.NewPerceptronTagger()
	tags := postagger.Tag(tokens)
	for _, token := range tags {
		fmt.Println(token.Text, token.Tag)
	}

	fmt.Println("Noun Chunks::", getChunks(myText, "NN"))
	fmt.Println("Verb Chunks::", getChunks(myText, "V"))

	
	regex := chunk.TreebankNamedEntities
	// Loop: tag + reg == Named Entity Chunks
	for _, entity := range chunk.Chunk(postagger.Tag(tokens), regex) {
		fmt.Println(entity)
	}

}

func getChunks(text string, tagName string) []string {
	// Tokenize
	tokens := tokenize.NewTreebankWordTokenizer().Tokenize(text)

	// Tags
	tags := tag.NewPerceptronTagger().Tag(tokens)

	// if tag ==  requested tagName
	chunks := []string{}
	for _, token := range tags {
		if strings.HasPrefix(token.Tag, tagName) {
			chunks = append(chunks, token.Text)
		}
	}
	return chunks
}

func languageDetectionExample02() { // Using github.com/abadojack/whatlanggo
	var mydocx string = "Hello world of Go"
	lang := whatlanggo.Detect(mydocx)
	fmt.Println("Text:", mydocx)
	fmt.Println("whatlango: ", lang.Lang.String()) // Language name
	fmt.Println("whatlango: ", lang.Confidence)    // Confidence/Accuracy of prediction
}

func languageDetectionExample01() { // Using github.com/rylans/getlang
	var mystr string = "Hello world of Go"
	// var mystrfr string = "Bonjour a tous"

	lang := getlang.FromString(mystr)
	// lang2 := getlang.FromString(mystrfr)

	fmt.Println("Text:", mystr)
	fmt.Println("getlang: ", lang.LanguageCode()) // Language code
	fmt.Println("getlang: ", lang.Confidence())   // Confidence/Accuracy of prediction

	// fmt.Println("Text:", mystrfr)
	// fmt.Println("getlang: ",lang2.LanguageCode()) // Language code
	// fmt.Println("getlang: ",lang2.Confidence())   // Confidence/Accuracy of prediction
}

func textCleaningExamples01() { // Using github.com/mingrammer/commonregex
	// Textcleaning using Regex & more
	// var mystr string = "Hello GoDev my email is jharis@gmail.com"
	// Multi line large text string literal ``
	// 	var docx string = `
	// 	Golang was designed at Google by Robert Griesemer, Rob Pike,
	//  and Ken Thompson. Ken called Rob on 519-555-7765 which was redirected to +44 22 777 555.
	// Jesse sent an email to jc.@gmail.com which he found on the website http://jcharistech.com.
	// Golang was publicly announced in November 2009 and version 1.0 was released in March 2012.
	// Go is widely used in production at Google USA and in many other organizations and open-source projects.
	// In November 2016, the Go and Go Mono fonts were released by type designers Charles Bigelow and Kris Holmes specifically for use by the Go project. Go is a humanist sans-serif which resembles Lucida Grande and Go Mono is monospaced. Each of the fonts adhere to the WGL4 character set and were designed to be legible with a large x-height and distinct letterforms. Both Go and Go Mono adhere to the DIN 1450 standard by having a slashed zero, lowercase l with a tail, and an uppercase I with serifs.
	// I have been coding since 4:00 AM this morning.Accra is big but not bigger as London.
	// john.smith@yahoo.com
	// 	`

	/*
		Reading text from a file
		os, ioutil, bufio
	*/
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	docx2 := string(content)
	fmt.Println("Sample:::", docx2)

	// Text Preprocessing
	// Normalizing: uniform case, removing unicode chars
	// fmt.Println(strings.ToLower(mystr))
	// fmt.Println(docx)
	// Remove noise [ special chars, eamils, phone #s]
	// Lemma/Stemming
	// Tokenization

	/*
	  Extra Emails
	  Method 1: Standard Library Regexp
	*/
	// Pattern
	// p := regexp.MustCompile(`GoDev`)
	// Find/Replace
	// fmt.Println(p.FindAllString(mystr, 1))
	// fmt.Println(p.ReplaceAllString(mystr, "REPLACED"))

	/*
	  Extra Emails
	  Method 2: External Library commonregex similar
	  to python library neatregex
	*/
	// fmt.Println(cregex.Emails(mystr))

	// Exercise 1.
	fmt.Println("Emails: ", cregex.Emails(docx2))

	// Remove/Replace : Document Redaction/Text Cleaning
	p := cregex.EmailRegex
	fmt.Println("ALTERED::: ", p.ReplaceAllString(docx2, "REPLACED"))

}
func stringExamples02() {
	var mystr string = "hello Go"
	fmt.Printf("Value: '%s'\n", mystr)
	fmt.Printf("Type: %T\n", mystr)
	fmt.Printf("Length: %d\n", len(mystr))
	fmt.Printf("Uppercase: %s\n", strings.ToUpper(mystr))
	fmt.Printf("Lowercase: %s\n", strings.ToLower(mystr))
	fmt.Printf("Titlecase: %s\n", strings.Title(mystr))
	fmt.Printf("Count 'l' occurrences: %d\n", strings.Count(mystr, "l"))
	fmt.Printf("Contains 'Go': %t\n", strings.Contains(mystr, "Go"))
	fmt.Printf("Split on ' ': %q\n", strings.Split(mystr, " "))
	fmt.Printf("Split after 'hel': %q\n", strings.SplitAfter(mystr, "hel"))
	fmt.Printf("Replace 'hello': %s\n", strings.ReplaceAll(mystr, "hello", "I love"))

	s := strings.Split(strings.ReplaceAll(mystr, "hello", "N.L.P programing"), " ")
	ss := strings.Join(s, " using ")
	fmt.Printf("Split and Join : %s\n", ss)

}

func stringExamples01() {
	fmt.Println("Method1: Create Characters with type ")
	var char byte = 'A'
	var char2 rune = 'A'
	fmt.Printf("Char as Byte %d %T\n", char, char)
	fmt.Printf("Char as Rune %d %T\n", char2, char2)

	fmt.Println("Method2: Create Characters with Method")
	charA := byte('A')
	charB := rune('A')
	fmt.Printf("Char as Byte:Fxn %d %T\n", charA, charA)
	fmt.Printf("Char as Rune:Fxn %d %T\n", charB, charB)

	fmt.Println("Actual Representation")
	str1 := string(char)
	str2 := string(char2)
	fmt.Printf("Char as String %s %T\n", str1, str1)
	fmt.Printf("Char as String %s %T\n", str2, str2)

	fmt.Println("Representation method 2 using Printf")
	fmt.Printf("Char as Byte %c %T\n", char, char)
	fmt.Printf("Char as Rune %c %T\n", char2, char2)
}
func filesExmaples() {
	args := os.Args[1:]
	if len(args) == 2 {
		switch args[1] {
		case "csv":
			files.OpenCSVFile(args[0])
		case "pdf":
			files.OpenPDFFile(args[0])
		case "txt":
			files.OpenTextFile(args[0])
		default:
			fmt.Println("Usage: file type[csv|txt|pdf]")
		}
	} else {
		fmt.Print("Usage: file type[csv|txt|pdf]")
	}
}
