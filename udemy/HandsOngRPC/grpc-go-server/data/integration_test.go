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
	dbName   = "bank_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var dummyBankAccount = BankAccount{
	AccountNumber:  "1234567890",
	AccountName:    "John Doe",
	Currency:       "USD",
	CurrentBalance: 1000.00,
}

var models Models
var testDB *sql.DB
var resourse *dockertest.Resource
var pool *dockertest.Pool
var sqlUpMigrations = []string{
	"../migrations/1700064420603920_bank_accounts.postgres.up.psql",
	"../migrations/1700069746098377_bank_tranfers.postgres.up.psql",
	"../migrations/1700070206630394_bank_transactions.postgres.up.psql",
	"../migrations/1700070540906734_bank_exchange_rates.postgres.up.psql",
}
var sqlDownMigrations = []string{
	"../migrations/1700070540906734_bank_exchange_rates.postgres.down.psql",
	"../migrations/1700070206630394_bank_transactions.postgres.down.psql",
	"../migrations/1700069746098377_bank_tranfers.postgres.down.psql",
	"../migrations/1700064420603920_bank_accounts.postgres.down.psql",
}

func setup() {
	os.Setenv("DATABASE_TYPE", "postgres")
	os.Setenv("UPPER_DB_LOG", "ERROR")

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
	for _, sqlUpMigration := range sqlUpMigrations {
		// Read SQl text from file
		bytes, err := ioutil.ReadFile(sqlUpMigration)
		if err != nil {
			return fmt.Errorf("Could  not read sql migration file: %s", err)
		}

		// Execute SQL text string
		_, err = db.Exec(string(bytes))
		if err != nil {
			return fmt.Errorf("Could not tables: %s", err)
		}
	}
	return nil
}

func truncateTables(db *sql.DB) error {
	for _, sqlDownMigration := range sqlDownMigrations {
		// Read SQl text from file
		bytes, err := ioutil.ReadFile(sqlDownMigration)
		if err != nil {
			return fmt.Errorf("Could  not read sql migration file: %s", err)
		}

		// Execute SQL text string
		_, err = db.Exec(string(bytes))
		if err != nil {
			return fmt.Errorf("Could not tables: %s", err)
		}
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

func TestBankAccount_Table(t *testing.T) {
	s := models.BankAccount.Table()
	if s != "bank_accounts" {
		t.Errorf("Wrong table name returned. Expected 'bank_accounts', got '%s'", s)
	}
}

func TestBankAccount_Insert(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	id, err := models.BankAccount.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	if id == 0 {
		t.Errorf("No id returned, Zero id returned")
	}
}

/*

// Test getting a BankAccount
func TestBankAccount_Get(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	id, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.bank_accounts.Get(id)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.ID != id {
		t.Errorf("Wrong user returned. Expected id %d, got %d", id, user.ID)
	}
}

// Test getting all bank_accounts
func TestBankAccount_GetAll(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	bank_accounts, err := models.bank_accounts.GetAll()
	if err != nil {
		t.Errorf("Error getting bank_accounts: %s", err)
	}

	if len(bank_accounts) != 1 {
		t.Errorf("Wrong number of bank_accounts returned. Expected 1, got %d", len(bank_accounts))
	}
}

// Test getting a user by email
func TestBankAccount_GetByEmail(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.Email != dummyBankAccount.Email {
		t.Errorf("Wrong user returned. Expected email %s, got %s", dummyBankAccount.Email, user.Email)
	}
}

// Test updating a user
func TestBankAccount_Update(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	user.LastName = "Smith"
	err = user.Update(*user)
	if err != nil {
		t.Errorf("Error updating user: %s", err)
	}

	user, err = models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	if user.LastName != "Smith" {
		t.Errorf("BankAccount LastName not updated. Expected last name 'Smith', got '%s'", user.LastName)
	}
}

// Test that password matches
func TestBankAccount_PasswordMatches(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}
	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	matches, err := user.PasswordMatches(dummyBankAccount.Password)
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

// Test resetting the bank_accounts password
func TestBankAccount_ResetPassword(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	id, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	err = models.bank_accounts.ResetPassword(id, "newpassword")
	if err != nil {
		t.Errorf("Error resetting password: %s", err)
	}

	err = models.bank_accounts.ResetPassword(id+1, "newpassword")
	if err == nil {
		t.Errorf("did not get error resetting password for non-existing user: %s", err)
	}

	matches, err := models.bank_accounts.PasswordMatches("newpassword")
	if err != nil {
		t.Errorf("Error matching password: %s", err)
	}

	if !matches {
		t.Errorf("Password does not match")
	}
}

// Test deleting a user
func TestBankAccount_Delete(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	id, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	err = models.bank_accounts.Delete(id)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
	}

	_, err = models.bank_accounts.Get(id)
	if err == nil {
		t.Errorf("BankAccount not deleted")
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

	id, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(id, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	if token.BankAccountID != id {
		t.Errorf("Wrong user id in token. Expected %d, got %d", id, token.BankAccountID)
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

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}
}

func TestToken_GetBankAccountForToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	user, err = models.Tokens.GetBankAccountForToken(token.PlainText)
	if err != nil {
		t.Errorf("Error getting user for token: %s", err)
	}

	if user.Email != dummyBankAccount.Email {
		t.Errorf("Wrong user returned. Expected %s, got %s", dummyBankAccount.Email, user.Email)
	}

	_, err = models.Tokens.GetBankAccountForToken("wrongToken")
	if err == nil {
		t.Errorf("Did not get error for getting a user with a wrong token")
	}
}

func TestToken_GetTokensForBankAccount(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	tokens, err := models.Tokens.GetTokensForBankAccount(1)
	if err != nil {
		t.Errorf("Did not get error for getting tokens for a non existing user: %s", err)
	}

	if len(tokens) != 0 {
		t.Errorf("Got tokens for a non existing user, length %d", len(tokens))
	}

	_, err = models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	tokens, err = models.Tokens.GetTokensForBankAccount(user.ID)
	if err != nil {
		t.Errorf("Error getting tokens for user: %s", err)
	}

	if len(tokens) == 0 {
		t.Errorf("No tokens returned")
	}

	if tokens[0].BankAccountID != user.ID {
		t.Errorf("Wrong user id returned. Expected %d, got %d", user.ID, tokens[0].BankAccountID)
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

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	tokenID, err := models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	tt, err := models.Tokens.Get(10)
	if err == nil {
		t.Errorf("Did not get error for getting a non existing token")
	}

	tt, err = models.Tokens.Get(tokenID)
	if err != nil {
		t.Errorf("Error getting token: %s", err)
	}

	if tt.ID != tokenID {
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

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	tokenID, err := models.Tokens.Insert(*token, *user)
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

	if tt.ID != tokenID {
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
	{"no_header", "", "me@here,com", true, "no header, but user is valid"},
	{"bad_header", "", "me@here,com", true, "bad header, but user is valid"},
	{"valid", "", "me@here.com", false, "valid token reported as invalid"},
	{"valid_user_deleted", "", "me@here.com", true, "valid token but user was deleted"},
}

// AuthenticateToken
func TestToken_AuthenticateToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	for _, tt := range authData {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/", nil)
			if tt.email == dummyBankAccount.Email {
				req.Header.Add("Authorization", "Bearer "+token.PlainText)
			} else {
				req.Header.Add("Authorization", "Bearer "+tt.token)
				if tt.name == "no_header" {
					req.Header.Del("Authorization")
				}

				if tt.name == "bad_header" { // Removed space after Bearer
					req.Header.Set("Authorization", "Bearer"+tt.token)
				}

			}

			if tt.name == "valid_user_deleted" {
				err = models.bank_accounts.Delete(user.ID)
				if err != nil {
					t.Errorf("Error deleting user: %s", err)
				}
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

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	okay, err := models.Tokens.ValidToken(token.PlainText)
	if err != nil {
		t.Errorf("Error validating token: %s", err)
	}

	if !okay {
		t.Errorf("Token should be valid")
	}

	okay, err = models.Tokens.ValidToken("abcdefghijklmnopqrstuvwxyz")
	if err == nil {
		t.Errorf("Error validating token: %s", err)
	}

	if okay {
		t.Errorf("Token should not be valid")
	}


	err = models.Tokens.Delete(user.Token.ID)
	if err != nil {
		t.Errorf("Error deleting token: %s", err)
	}

	okay, err = models.Tokens.ValidToken(user.Token.PlainText)
	if err == nil {
		t.Errorf("Error validating token: %s", err)
	}

	if okay {
		t.Errorf("no error reported for non-existing token")
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

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	err = models.Tokens.Delete(token.ID)
	if err != nil {
		t.Errorf("Error deleting token: %s", err)
	}
}

// ExpiredToken
func TestToken_ExpiredToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := -time.Hour
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	b, err := models.Tokens.ValidToken(token.PlainText)
	if err == nil {
		t.Errorf("Expired token error should be returned: %s", err)
	}

	if b {
		t.Errorf("Token should not be valid")
	}

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Authorization", "Bearer "+token.PlainText)

	_, err = models.Tokens.AuthenticateToken(req)
	if err == nil {
		t.Errorf("Failed to catch expired token: %s", err)
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

	_, err := models.bank_accounts.Insert(dummyBankAccount)
	if err != nil {
		t.Errorf("Error inserting user: %s", err)
	}

	user, err := models.bank_accounts.GetByEmail(dummyBankAccount.Email)
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}

	time_duration := time.Hour * 24 * 365
	token, err := models.Tokens.GenerateToken(user.ID, time_duration)
	if err != nil {
		t.Errorf("Error generating token: %s", err)
	}

	_, err = models.Tokens.Insert(*token, *user)
	if err != nil {
		t.Errorf("Error inserting token: %s", err)
	}

	err = models.Tokens.DeleteByToken(token.PlainText)
	if err != nil {
		t.Errorf("Error deleting token: %s", err)
	}
}

func TestToken_DeleteNonExistentToken(t *testing.T) {
	defer func() { // Truncate tables after test
		err := truncateTables(testDB)
		if err != nil {
			t.Errorf("Error truncating tables: %s", err)
		}
	}()
	err := models.Tokens.DeleteByToken("non-existent-token")
	if err != nil {
		t.Errorf("Error should be returned")
	}
}
*/
