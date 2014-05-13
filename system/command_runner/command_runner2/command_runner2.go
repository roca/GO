// Copyright © 2011-12 Qtrac Ltd.
//
// This program or package and any associated files are licensed under the
// Apache License, Version 2.0 (the "License"); you may not use these files
// except in compliance with the License. You can get a copy of the License
// at: http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The approach taken here was inspired by an example on the gonuts mailing
// list by Roger Peppe.

package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"
)



func (command Command) execute() []byte {
	command_arguments := strings.Split(command.path, " ")

	cmd := exec.Command(command_arguments[0], command_arguments[1:]...)
	cmd.Dir = command.dir
	//time.Sleep(1000 * time.Millisecond)

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return out

}

func (job Job) Do() {

	out := job.execute()
	job.results <- Result{fmt.Sprintf("Output from '%s' is\n%s\n", job.path, out)}

}

func main() {

	var async = flag.Bool("async", false, "async flag")

	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	fmt.Println("The number of CPU on this server is", runtime.NumCPU())

	commands := make([]Command, 10000)

	for i := range commands {
		commands[i] = Command{i, "date", "", 1}
	}

	t0 := time.Now()
	if *async {
		executeCommands(commands)
		t1 := time.Now()
		fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	} else {
		for i := range commands {
			fmt.Printf(fmt.Sprintf("%d:Output from '%s' is\n%s\n", i, commands[i].path, commands[i].execute()))
		}
		t1 := time.Now()
		fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
	}

	//db := getDatabaseConnection("commander", "cody")
	//groups := queryForGroups(db)
	//for _, group := range groups {
	//	commands := queryForCommands(db, group)
	//	executeCommands(commands)
	//	time.Sleep(5000 * time.Millisecond)
	//}

}

func executeCommands(commands []Command) {
	jobs := make(chan Job, workers)
	results := make(chan Result, minimum(1000, len(commands)))
	done := make(chan struct{}, workers)

	go addJobs(jobs, commands, results) // Executes in its own goroutine
	for i := 0; i < workers; i++ {
		go doJobs(done, jobs) // Each executes in its own goroutine
	}
	go awaitCompletion(done, results) // Executes in its own goroutine
	processResults(results)           // Blocks until the work is done
}

func addJobs(jobs chan<- Job, commands []Command, results chan<- Result) {
	for _, command := range commands {
		jobs <- Job{command, results}
	}
	close(jobs)
}

func doJobs(done chan<- struct{}, jobs <-chan Job) {
	for job := range jobs {
		job.Do()
	}
	done <- struct{}{}
}

func awaitCompletion(done <-chan struct{}, results chan Result) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(results)
}

func processResults(results <-chan Result) {
	i := 0
	for result := range results {
		i++
		fmt.Printf("%d: %s\n", i,result.output)
	}
}

func minimum(x int, ys ...int) int {
	for _, y := range ys {
		if y < x {
			x = y
		}
	}
	return x
}

