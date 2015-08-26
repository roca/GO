package main 
 
import (
	"strings"
	"fmt"
)

func main() { 
	phrase := "These are the times that try men's souls\n"
	
	words := strings.Split(phrase, " ")
	
	ch := make(chan string, len(words))
	
	for _, word := range words {
		ch <- word
	}
	
	close(ch)
	
	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
	
	ch <- "test"
	
}