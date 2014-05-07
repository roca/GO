package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


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
