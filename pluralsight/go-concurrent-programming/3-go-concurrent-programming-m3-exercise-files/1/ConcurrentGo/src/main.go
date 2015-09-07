package main 
 
import (
	"fmt" 
)

func main() { 
	ch := make(chan string,1)
	
	ch <- "Hello"
	
	fmt.Println(<-ch)
	
	ch <- "Go"
	
	fmt.Println(<-ch)
}

/*
3.1.1 - basic channel with no buffer
package main 
 
import (
	"fmt" 
)

func main() { 
	ch := make(chan string)
	
	ch <- "Hello"
	
	fmt.Println(<-ch)
}
*/

/*
3.1.2 - add buffer to channel to make non-blocking on one message
package main 
 
import (
	"fmt" 
)

func main() { 
	ch := make(chan string,1)
	
	ch <- "Hello"
	
	fmt.Println(<-ch)
}
*/

/*3.1.3 - send and receive second message in channel

package main 
 
import (
	"fmt" 
)

func main() { 
	ch := make(chan string,1)
	
	ch <- "Hello"
	
	fmt.Println(<-ch)
	
	ch <- "Go"
	
	fmt.Println(<-ch)
}
*/