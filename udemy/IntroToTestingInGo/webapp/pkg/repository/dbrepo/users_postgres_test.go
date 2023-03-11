package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbName   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB

func TestMain(m *testing.M) {
	// connect to docker; fail if docker is not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker, is it running? : %v", err)
	}
	pool = p

	// set up our docker options, specifying the image we want to use
	opt := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	// get a resource (docker image)
	resource, err = pool.RunWithOptions(&opt)
	if err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Could not start resource: %v", err)
	}

	// start the image and wait for it to be ready
	reTryFunc := func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error: ", err)
			return err
		}
		return testDB.Ping()
	}
	if err := pool.Retry(reTryFunc); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Could not connect to database: %v", err)
	}

	// populate the test database with empty tables
	err = createTables()
	if err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	// run the tests
	code := m.Run()

	// clean up after ourselves
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %v", err)
	}

	os.Exit(code)
}

func createTables() error {
	tableSQL, err := os.ReadFile("./testdata/users.sql")
	if err != nil {
		fmt.Println("Error SQL reading file: ", err)
		return err
	}

	_, err = testDB.Exec(string(tableSQL))
	if err != nil {
		fmt.Println("Error SQL creating tables: ", err)
		return err
	}
	return nil
}

func Test_pingDB(t *testing.T) {
	err := testDB.Ping()
	if err != nil {
		t.Error(err)
	}
}
