// Package state is the core API for the blockchain and implements all the
// business rules and processing.
package state

import (
	"blockchain/foundation/blockchain/database"
	"blockchain/foundation/blockchain/genesis"
	"sync"
)

// EventHandler defines a function that is called when events
// occur in the processing of persisting blocks.
type EventHandler func(v string, args ...any)

// =============================================================================

// Config represents the configuration required to start
// the blockchain node.
type Config struct {
	BeneficiaryID database.AccountID
	Genesis       genesis.Genesis
	EvHandler     EventHandler
}

// State manages the blockchain database.
type State struct {
	mu sync.RWMutex

	beneficiaryID database.AccountID
	evHandler     EventHandler

	genesis genesis.Genesis
	db      *database.Database
}

// New constructs a new blockchain for data management.
func New(cfg Config) (*State, error) {

	// Build a safe event handler function for use.
	ev := func(v string, args ...any) {
		if cfg.EvHandler != nil {
			cfg.EvHandler(v, args...)
		}
	}

	// Access the storage for the blockchain.
	db, err := database.New(cfg.Genesis, ev)
	if err != nil {
		return nil, err
	}

	state := State{
		beneficiaryID: cfg.BeneficiaryID,
		evHandler:     ev,

		genesis: cfg.Genesis,
		db:      db,
	}

	return &state, nil
}

func (s *State) Shutdown() error {
	s.evHandler("state: shutdown: started")
	defer s.evHandler("state: shutdown: completed")

	return nil
}
