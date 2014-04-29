// command_runner
package main

import (
	"database/sql"
	"fmt"
	//"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
	"log"
	"os/exec"
	"time"
)

type Command string

type Runner interface {
	run(done chan bool)
}

func (command *Command) run(done chan bool) {
	//fmt.Println("--------------------------------Working on command: ", *command)

	out, err := exec.Command(string(*command)).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Output from '%s' is\n%s\n", *command, out)

	//time.Sleep(5000 * time.Millisecond)
	//fmt.Println("--------------------------------Done with ", *command)
	done <- true
}

func main() {

	db, err := sql.Open("mymysql", "commandlogs_dev/commander/cody")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	conn_err := db.Ping()
	if conn_err != nil {
		panic(conn_err)
	}

	commands := queryCommands(db)

	//commands := make([]Command, 10000)

	//for i := range commands {
	//	commands[i] = Command(fmt.Sprintf("command %d", i))
	//}

	done := make(chan bool)
	defer func() { close(done) }()

	for i := range commands {
		go commands[i].run(done)
	}

	all_done := len(commands)
	doneCount := 0
	timeout := time.After(6 * time.Second)

	for doneCount != all_done {

		select {
		case <-done:
			doneCount++
		case <-timeout:
			fmt.Println("timed out!")
			doneCount = all_done
		}

	}

	fmt.Println("all done !")
}

func queryCommands(db *sql.DB) []Command {
	commands := []Command{}
	rows, err := db.Query("SELECT path FROM commands")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var path string
		if err := rows.Scan(&path); err != nil {
			log.Fatal(err)
		}
		commands = append(commands, Command(path))
		//fmt.Printf("%s\n", path)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return commands

}
