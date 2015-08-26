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
	
	for msg := range ch {
		fmt.Print(msg + " ")
	}
	
}

/*3.4.1 - before ranging over a channel
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
	
	for i:=0; i < len(words); i++ {
		fmt.Print(<-ch + " ")
	}
	
}
*/

/*3.4.2 - use if test to range over loop
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
	
	for {
		if msg, ok := <- ch; ok {
			fmt.Print(msg + " ")
		} else {
			break
		}
	}
	
}
*/

/*3.4.3 - use range keyword to range over channel 
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
	
	for msg := range ch {
		fmt.Print(msg + " ")
	}
	
}
*/