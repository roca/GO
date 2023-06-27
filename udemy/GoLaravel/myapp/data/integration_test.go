//go:build integration

// run test with this command: go test . --tags integration --count=1

package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "secret"
	dbName   = "celeritas_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var dummtUser = User{
	FirstName: "Some",
	LastName:  "Guy",
	Email:     "me@here.com",
	Active:    1,
	Password:  "password",
}

var models Models
var testDB *sql.DB
var resourse *dockertest.Resource
var pool *dockertest.Pool

func setup() {
	os.Setenv("DATABASE_TYPE", "postgres")

	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	pool = p

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "13.4",
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

	resourse, err = pool.RunWithOptions(&opts)
	if err != nil {
		_ = pool.Purge(resourse)
		log.Fatalf("Could not start resource: %s", err)
	}

	if err := pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbName))
		if err != nil {
			log.Println("Error:", err)
			return err
		}

		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resourse)
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Create tables
	err = createTables(testDB)
	if err != nil {
		_ = pool.Purge(resourse)
		log.Fatalf("Could not create tables: %s", err)
	}

	models = New(testDB)

	return
}

func createTables(db *sql.DB) error {
	// Read SQl text from file
	bytes, err := ioutil.ReadFile("./users.sql")
	if err != nil {
		return fmt.Errorf("Could  not read users.sql file: %s", err)
	}

	// Execute SQL text string
	_, err = db.Exec(string(bytes))
	if err != nil {
		return fmt.Errorf("Could not create tables: %s", err)
	}

	return nil
}

func teardown() {
	return
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestUser_Table(t *testing.T) {
	s := models.Users.Table()
	if s != "users" {
		t.Errorf("Wrong table name returned. Expected 'users', got '%s'", s)
	}
}
