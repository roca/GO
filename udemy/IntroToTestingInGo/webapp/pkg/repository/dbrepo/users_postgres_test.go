package dbrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"webapp/pkg/data"
	"webapp/pkg/repository"

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
var testRepo repository.DatabaseRepo

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

	testRepo = &PostgresDBRepo{DB: testDB}

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

func TestPostgresDBRepoInsertUser(t *testing.T) {
	// create a new user
	testUser := data.User{
		Email:     "admin@example.com",
		FirstName: "Admin",
		LastName:  "User",
		Password:  "secret",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	id, err := testRepo.InsertUser(testUser)
	if err != nil {
		t.Errorf("insert user returned an error: %s", err)
	}

	if id != 1 {
		t.Errorf("insert user returned wrone id; expected 1, but got %d", id)
	}
}

func TestPostgresDBRepoAllUser(t *testing.T) {
	users, err := testRepo.AllUsers()
	if err != nil {
		t.Errorf("all user returned an error: %s", err)
	}

	if len(users) != 1 {
		t.Errorf("all user returned wrong number of users; expected 1, but got %d", len(users))
	}

	// create a new user
	testUser := data.User{
		Email:     "jsmith@example.com",
		FirstName: "Jack",
		LastName:  "Smith",
		Password:  "secretAgentSmith",
		IsAdmin:   1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, _ = testRepo.InsertUser(testUser)

	users, err = testRepo.AllUsers()
	if err != nil {
		t.Errorf("all user returned an error: %s", err)
	}

	if len(users) != 2 {
		t.Errorf("all user returned wrong number of users after insert; expected 2, but got %d", len(users))
	}
}

func TestPostgresDBRepoGetUser(t *testing.T) {
	user, err := testRepo.GetUser(1)
	if err != nil {
		t.Errorf("get user returned an error: %s", err)
	}

	if user.Email != "admin@example.com" {
		t.Errorf("wrong email returned by GetUser; expected admin@example.com but got %s",user.Email)
	}

	user, err = testRepo.GetUser(3)
	if err == nil {
		t.Errorf("no error reported when getting a nonexisting user: %s", err)
	}
}

func TestPostgresDBRepoGetUserByEmail(t *testing.T) {
	user, err := testRepo.GetUserByEmail("jsmith@example.com")
	if err != nil {
		t.Errorf("get user by email returned an error: %s", err)
	}

	if user.ID != 2 {
		t.Errorf("wrong id returned by GetUserByEmail; expected 2 but got %d",user.ID)
	}
}

func TestPostgresDBRepoUpdateUser(t *testing.T) {
	user, _ := testRepo.GetUser(2)
	user.Email = "janesmith@example.com"
	user.FirstName = "Jane"
	err := testRepo.UpdateUser(*user)
	if err != nil {
		t.Errorf("update user returned an error: %s", err)
	}

	user, _ = testRepo.GetUser(2)
	if user.FirstName != "Jane" || user.Email != "janesmith@example.com" {
		t.Errorf("update user failed to update user")
	}
}

func TestPostgresDBRepoDeleteUser(t *testing.T) {
	err := testRepo.DeleteUser(2)
	if err != nil {
		t.Errorf("delete user returned an error: %s", err)
	}

	_, err = testRepo.GetUser(2)
	if err == nil {
		t.Errorf("no error reported when deleting a nonexisting user: %s", err)
	}

}

func TestPostgresDBRepoResetPassword(t *testing.T) {
	err := testRepo.ResetPassword(1, "newPassword")
	if err != nil {
		t.Errorf("reset password returned an error: %s", err)
	}

	user, _ := testRepo.GetUser(1)
	matches, err := user.PasswordMatches("newPassword")
	if err != nil {
		t.Errorf("password match returned an error: %s", err)
	}

	if !matches {
		t.Errorf("password match returned false for valid password")
	}

}
