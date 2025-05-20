// Package database handles all the lower level support for maintaining the
// blockchain in storage and maintaining an in-memory databse of account information.
package database

import (
	"blockchain/foundation/blockchain/genesis"
	"blockchain/foundation/blockchain/signature"
	"errors"
	"sort"
	"sync"
)

// =============================================================================

// Database manages data related to accounts who have transacted on the blockchain.
type Database struct {
	mu          sync.RWMutex
	genesis     genesis.Genesis
	latestBlock Block
	accounts    map[AccountID]Account
}

// New constructs a new database and applies account genesis information and
// reads/writes the blockchain database on disk if a dbPath is provided.
func New(genesis genesis.Genesis, evHandler func(v string, args ...any)) (*Database, error) {
	db := Database{
		genesis:  genesis,
		accounts: make(map[AccountID]Account),
	}

	// Update the database with account balance information from genesis.
	for accountStr, balance := range genesis.Balances {
		accountID, err := ToAccountID(accountStr)
		if err != nil {
			return nil, err
		}
		db.accounts[accountID] = newAccount(accountID, balance)

		evHandler("Account: %s, Balance: %d", accountID, balance)
	}

	return &db, nil
}

// Remove deletes an account from the database.
func (db *Database) Remove(accountID AccountID) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.accounts, accountID)
}

// Query retrieves an account from the database.
func (db *Database) Query(accountID AccountID) (Account, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	acount, exists := db.accounts[accountID]
	if !exists {
		return Account{}, errors.New("account does not exist")
	}

	return acount, nil
}

// Copy makes a copy of the current accounts in the database.
func (db *Database) Copy() map[AccountID]Account {
	db.mu.RLock()
	defer db.mu.RUnlock()

	accounts := make(map[AccountID]Account)
	for accountID, account := range db.accounts {
		accounts[accountID] = account
	}
	return accounts
}

// UpdateLatestBlock provides safe access to update the latest block.
func (db *Database) UpdateLatestBlock(block Block) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.latestBlock = block
}

// LatestBlock returns the latest block.
func (db *Database) LatestBlock() Block {
	db.mu.RLock()
	defer db.mu.RUnlock()

	return db.latestBlock
}

// HashState returns a hash based on the contents of the accounts and
// their balances. This is added to each block and checked by peers.
func (db *Database) HashState() string {
	accounts := make([]Account, 0, len(db.accounts))
	db.mu.RLock()
	{
		for _, account := range db.accounts {
			accounts = append(accounts, account)
		}
	}
	db.mu.RUnlock()

	sort.Sort(byAccount(accounts))
	return signature.Hash(accounts)
}
