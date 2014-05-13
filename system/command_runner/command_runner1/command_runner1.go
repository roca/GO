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
	"strings"
	"time"
)

type Group struct {
	Id int
}

type Command struct {
	Id    int
	Path  string
	Dir   string
	Async int
}

type Groups []*Group

func (s Groups) Len() int      { return len(s) }
func (s Groups) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type ById struct{ Groups }

func (s ById) Less(i, j int) bool { return s.Groups[i].Id < s.Groups[j].Id }

type Runner interface {
	run(done chan bool)
}

func (command *Command) run(done chan bool) {
	//fmt.Println("--------------------------------Working on command: ", *command)

	command_arguments := strings.Split(command.Path, " ")

	cmd := exec.Command(command_arguments[0], command_arguments[1:]...)
	cmd.Dir = command.Dir

	out, err := cmd.Output()
	fmt.Printf("Output from '%s' is\n%s\n", command.Path, out)
	if err != nil {
		log.Fatal(err)
	}

	done <- true
}

func main() {

	db := getDatabaseConnection("commander", "cody")
	groups := queryGroups(db)
	for _, group := range groups {
		commands := queryCommands(db, group)
		runCommands(commands)
		fmt.Println("group ", group.Id)
	}

	fmt.Println("all done !")
}

func runCommands(commands []Command) {
	done := make(chan bool)
	defer func() { close(done) }()

	for i := range commands {
		if commands[i].Async == 0 {
			commands[i].run(done)
		} else {
			go commands[i].run(done)
		}
	}

	all_done := len(commands)
	doneCount := 0
	timeout := time.After(3600 * time.Second)

	for doneCount != all_done {

		select {
		case <-done:
			doneCount++
		case <-timeout:
			fmt.Println("timed out!")
			doneCount = all_done
		}

	}

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

	db, err := sql.Open("mysql", username+":"+password+"@tcp(tarritdb01t:3306)/commandlogs_dev?charset=utf8")
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

func queryGroups(db *sql.DB) []Group {
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
func queryCommands(db *sql.DB, group Group) []Command {
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
