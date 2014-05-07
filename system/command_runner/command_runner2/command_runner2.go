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
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os/exec"
	"runtime"
	"strings"
	//"time"
)

var workers = runtime.NumCPU()

type Group struct {
	Id int
}

type Result struct {
	output string
}

type Job struct {
	Command
	results chan<- Result
}

type Command struct {
	Id    int
	path  string
	dir   string
	async int
}

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
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all the machine's cores
	fmt.Println("The number of CPU on this server is", runtime.NumCPU())

	commands := make([]Command, 10000)

	for i := range commands {
		commands[i] = Command{i, "date", "", 1}
		fmt.Printf(fmt.Sprintf("%d:Output from '%s' is\n%s\n", i, commands[i].path, commands[i].execute()))

	}

	//executeCommands(commands)

	//db := getDatabaseConnection("commander", "cody")
	//groups := queryForGroups(db)
	//for _, group := range groups {
	//	commands := queryForCommands(db, group)
	//	executeCommands(commands)
	//	time.Sleep(5000 * time.Millisecond)
	//}

}

func queryForGroups(db *sql.DB) []Group {
	groups := []Group{}
	rows, err := db.Query("SELECT distinct(group_id) FROM commands order by group_id")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			log.Fatal(err)
		}
		groups = append(groups, Group{id})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return groups

}

func queryForCommands(db *sql.DB, group Group) []Command {
	commands := []Command{}
	sql := fmt.Sprintf("SELECT id,path,dir,async FROM commands where group_id = %d", group.Id)
	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id, async int
		var path, dir string
		if err := rows.Scan(&id, &path, &dir, &async); err != nil {
			log.Fatal(err)
		}
		commands = append(commands, Command{id, path, dir, async})

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return commands

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
	for result := range results {
		fmt.Printf("%s\n", result.output)
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

func getDatabaseConnection(username, password string) *sql.DB {
	//rootCertPool := x509.NewCertPool()
	//pem, err := ioutil.ReadFile("/etc/ssl/tarritdb01p/regeneronCA.cert")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
	//	log.Fatal("Failed to append PEM.")
	//}
	//clientCert := make([]tls.Certificate, 0, 1)
	//certs, err := tls.LoadX509KeyPair("/etc/ssl/tarritdb01p/tarritdb01p.cer", "/etc/ssl/tarritdb01p/tarritdb01p.key")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//clientCert = append(clientCert, certs)
	//mysql.RegisterTLSConfig("custom", &tls.Config{
	//	RootCAs:      rootCertPool,
	//	Certificates: clientCert,
	//})

	//db, err := sql.Open("mysql", "commander:cody@tcp(tarritdb01t:3306)/commandlogs_dev?charset=utf8&tls=custom")

	//db, err := sql.Open("mysql", username+":"+password+"@tcp(tarritdb01t:3306)/commandlogs_dev?charset=utf8")
	db, err := sql.Open("mysql", username+":"+password+"@/commandlogs_dev?charset=utf8")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	conn_err := db.Ping()
	if conn_err != nil {
		panic(conn_err)
	}

	return db

}
