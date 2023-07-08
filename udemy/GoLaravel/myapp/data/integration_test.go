//go:build integration

// run test with this command: go test . --tags integration --count=1

package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

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

var dummyUser = User{
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

func truncateTables(db *sql.DB) error {
	// Read SQl text from file
	bytes, err := ioutil.ReadFile("./truncate.sql")
	if err != nil {
		return fmt.Errorf("Could  not read truncate.sql file: %s", err)
	}

	// Execute SQL text string
	_, err = db.Exec(string(bytes))
	if err != nil {
		return fmt.Errorf("Could not truncate tables: %s", err)
	}

	return nil
}

func teardown() {
	err := pool.Purge(resourse)
	if err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
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

func TestUser_Insert(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	if id == 0 {
		t.Errorf("No id returned, Zero id returned")
	}
}

// Test getting a User
func TestUser_Get(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.Users.Get(id)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.ID != id {
		t.Errorf("Wrong user returned. Expected id %d, got %d", id, user.ID)
	}
}

// Test getting all users
func TestUser_GetAll(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	users, err := models.Users.GetAll()
	if err != nil {
		t.Errorf("Error getting users: %s", err)
	}

	if len(users) != 1 {
		t.Errorf("Wrong number of users returned. Expected 1, got %d", len(users))
	}
}

// Test getting a user by email
func TestUser_GetByEmail(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.Email != dummyUser.Email {
		t.Errorf("Wrong user returned. Expected email %s, got %s", dummyUser.Email, user.Email)
	}
}

// Test updating a user
func TestUser_Update(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	user.LastName = "Smith"
	err = user.Update(*user)
	if err != nil {
		t.Errorf("Error updating user: %s", err)
	}

	user, err = models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.LastName != "Smith" {
		t.Errorf("User LastName not updated. Expected last name 'Smith', got '%s'", user.LastName)
	}
}

// Test that password matches
func TestUser_PasswordMatches(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	matches, err := user.PasswordMatches(dummyUser.Password)
	if err != nil {
		t.Errorf("Error matching password: %s", err)
	}

	if !matches {
		t.Errorf("Password does not match")
	}

	matches, err = user.PasswordMatches("wrongpassword")
	if err != nil {
		t.Errorf("Error matching password: %s", err)
	}

	if matches {
		t.Errorf("Password matches")
	}
}

// Test resetting the users password
func TestUser_ResetPassword(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	err = models.Users.ResetPassword(id, "newpassword")
	if err != nil {
		t.Errorf("Error resetting password: %s", err)
	}

	err = models.Users.ResetPassword(id+1, "newpassword")
	if err == nil {
		t.Errorf("did not get error resetting password for non-existing user: %s", err)
	}

	matches, err := models.Users.PasswordMatches("newpassword")
	if err != nil {
		t.Errorf("Error matching password: %s", err)
	}

	if !matches {
		t.Errorf("Password does not match")
	}
}

// Test deleting a user
func TestUser_Delete(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	err = models.Users.Delete(id)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}

	_, err = models.Users.Get(id)
	if err == nil {
		t.Errorf("User not deleted")
	}
}

// Test Table token
func TestToken_Table(t *testing.T) {
	token := models.Tokens
	if token.Table() != "tokens" {
		t.Errorf("Wrong table name returned. Expected 'tokens', got '%s'", token.Table())
	}
}

// Test Generating a token
func TestToken_GenerateToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	id, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(id, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	if token.UserID != id {
		t.Errorf("Wrong user id in token. Expected %d, got %d", id, token.UserID)
	}
}

// Test inserting a token
func TestToken_Insert(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}
}

func TestToken_GetUserForToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	user, err = models.Tokens.GetUserForToken(token.PlainText)
	if err != nil {
		t.Errorf("Error getting user for token: %s", err)
	}

	if user.Email != dummyUser.Email {
		t.Errorf("Wrong user returned. Expected %s, got %s", dummyUser.Email, user.Email)
	}

	_, err = models.Tokens.GetUserForToken("wrongToken")
	if err == nil {
		t.Errorf("Did not get error for getting a user with a wrong token")
	}
}

func TestToken_GetTokensForUser(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	tokens, err := models.Tokens.GetTokensForUser(1)
	if err != nil {
		t.Errorf("Did not get error for getting tokens for a non existing user: %s", err)
	}

	if len(tokens) != 0 {
		t.Errorf("Got tokens for a non existing user, length %d", len(tokens))
	}

	_, err = models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	tokens, err = models.Tokens.GetTokensForUser(user.ID)
	if err != nil {
		t.Errorf("Error getting tokens for user: %s", err)
	}

	if len(tokens) == 0 {
		t.Errorf("No tokens returned")
	}

	if tokens[0].UserID != user.ID {
		t.Errorf("Wrong user id returned. Expected %d, got %d", user.ID, tokens[0].UserID)
	}
}

// Get token by ID
func TestToken_Get(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	tt, err := models.Tokens.Get(10)
	if err == nil {
		t.Errorf("Did not get error for getting a non existing token")
	}

	tt, err = models.Tokens.Get(token.ID)
	if err != nil {
		t.Errorf("Error getting token: %s", err)
	}

	if tt.ID != token.ID {
		t.Errorf("Wrong token returned. Expected %d, got %d", token.ID, tt.ID)
	}
}

// Get token by PlainText
func TestToken_GetByToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	tt, err := models.Tokens.GetByToken("10")
	if err == nil {
		t.Errorf("Did not get error for getting a non existing token")
	}

	tt, err = models.Tokens.GetByToken(token.PlainText)
	if err != nil {
		t.Errorf("Error getting token: %s", err)
	}

	if tt.ID != token.ID {
		t.Errorf("Wrong token returned. Expected %d, got %d", token.ID, tt.ID)
	}
}

var authData = []struct {
	name        string
	token       string
	email       string
	errExpected bool
	message     string
}{
	{"invalid token", "abcdefghijklmnopqrstuvwxyz", "a@here,com", true, "invalid token accepted as valid"},
	{"invalid_length", "abcdefghijklmnopqrstuvwxy", "a@here,com", true, "token of wrong length token accepted as valid"},
	{"no_user", "abcdefghijklmnopqrstuvwxyz", "a@here,com", true, "no user, but token accepted as valid"},
	{"no_token", "", "me@here,com", true, "no token, but user is valid"},
	{"valid", "", "me@here.com", false, "valid token reported as invalid"},
}

// AuthenticateToken
func TestToken_AuthenticateToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	for _, tt := range authData {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			if tt.email == dummyUser.Email {
				req.Header.Add("Authorization", "Bearer "+token.PlainText)
			} else {
				req.Header.Add("Authorization", "Bearer "+tt.token)
			}

			_, err := models.Tokens.AuthenticateToken(req)
			if tt.errExpected && err == nil {
				t.Errorf("%s: %s", tt.name, tt.message)
			} else if !tt.errExpected && err != nil {
				t.Errorf("%s: %s - %s", tt.name, tt.message, err)
			}
		})
	}

}

// ValidToken
func TestToken_ValidToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	b, err := models.Tokens.ValidToken(token.PlainText)
	if err != nil {
		t.Errorf("Error validating token: %s", err)
	}

	if !b {
		t.Errorf("Token should be valid")
	}

}

// Delete token by id
func TestToken_Delete(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	err = models.Tokens.Delete(token.ID)
	if err != nil {
		t.Errorf("Error deleting token: %s", err)
	}
}


// Delete token by plain text
func TestToken_DeleteByToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.Users.Insert(dummyUser)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.Users.GetByEmail(dummyUser.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	err = models.Tokens.DeleteByToken(token.PlainText)
	if err != nil {
		t.Errorf("Error deleting token: %s", err)
	}
}

