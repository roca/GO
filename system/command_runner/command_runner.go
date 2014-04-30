// command_runner
package main

import (
	//"crypto/tls"
	//"crypto/x509"
	"database/sql"
	"fmt"
	//"github.com/astaxie/beedb"
	_ "github.com/go-sql-driver/mysql"
	//"io/ioutil"
	"log"
	"os/exec"
	"time"
)

type Command struct {
	Id   int
	Path string
	Dir  string
}

type Runner interface {
	run(done chan bool)
}

func (command *Command) run(done chan bool) {
	//fmt.Println("--------------------------------Working on command: ", *command)

	cmd := exec.Command(command.Path)
	cmd.Dir = command.Dir

	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Output from '%s' is\n%s\n", command.Path, out)

	time.Sleep(5000 * time.Millisecond)
	//fmt.Println("--------------------------------Done with ", *command)
	done <- true
}

func main() {

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

	db, err := sql.Open("mysql", "commander:cody@tcp(localhost:3306)/commandlogs_dev?charset=utf8")
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
	rows, err := db.Query("SELECT id,path,dir FROM commands")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int
		var path, dir string
		if err := rows.Scan(&id, &path, &dir); err != nil {
			log.Fatal(err)
		}
		commands = append(commands, Command{id, path, dir})
		//fmt.Printf("%s\n", path)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return commands

}
