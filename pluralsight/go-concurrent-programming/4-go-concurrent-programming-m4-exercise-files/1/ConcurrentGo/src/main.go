package main 
 
import ( 
	"fmt"
	"runtime" 
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	
	f, _ := os.Create("./log.txt")
	f.Close()
	
	logCh := make(chan string, 50)
	
	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()
	
	for i:=1; i < 10;i++ { 
		for j:=1; j<10;j++ {
			go func(i, j int) { 
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}
	
	fmt.Scanln()
}

/* 4.1.1 - initial demo code with sync.Mutex
package main 
 
import ( 
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	
	mutex := new(sync.Mutex)
	
	
	for i:=1; i < 10;i++ { 
		for j:=1; j<10;j++ {
			mutex.Lock()
			go func() { 
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				mutex.Unlock()
			}()
		}
	}
	
	fmt.Scanln()
}
*/

/* 2.4.2 - Update to use channel to force sequential access
package main 
 
import (  
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	
	mutex := make(chan bool, 1)
	
	
	for i:=1; i < 10;i++ { 
		for j:=1; j<10;j++ {
			mutex <- true
			go func() { 
				fmt.Printf("%d + %d = %d\n", i, j, i+j)
				<-mutex
			}()
		}
	}
	
	fmt.Scanln()
}
*/

/* 2.4.3 - async logger with channels
package main 
 
import (  
	"fmt"
	"runtime"
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	f, _ := os.Create("./log.txt")
	f.Close()
	
	logCh := make(chan string, 50)
	
	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				
				f.Close()
			} else {
				break
			}
		}
	}()
	
	mutex := make(chan bool, 1)
	
	 
	for i:=1; i < 10;i++ { 
		for j:=1; j<10;j++ {
			mutex <- true
			go func() { 
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
				<-mutex
			}()
		}
	}
	
	fmt.Scanln()
}
*/

/* 2.4.4 - remove mutex and pass parameters into goroutine
package main 
 
import ( 
	"fmt"
	"runtime" 
	"os"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)
	
	f, _ := os.Create("./log.txt")
	f.Close()
	
	logCh := make(chan string, 50)
	
	go func() {
		for {
			msg, ok := <- logCh
			if ok {
				f, _ := os.OpenFile("./log.txt", os.O_APPEND, os.ModeAppend)
				
				logTime := time.Now().Format(time.RFC3339)
				f.WriteString(logTime + " - " + msg)
				f.Close()
			} else {
				break
			}
		}
	}()
	
	for i:=1; i < 10;i++ { 
		for j:=1; j<10;j++ {
			go func(i, j int) { 
				msg := fmt.Sprintf("%d + %d = %d\n", i, j, i+j)
				logCh <- msg
				fmt.Print(msg)
			}(i, j)
		}
	}
	
	fmt.Scanln()
}
*/
