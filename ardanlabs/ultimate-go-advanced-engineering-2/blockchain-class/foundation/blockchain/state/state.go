// Package state is the core API for the blockchain and implements all the
// business rules and processing.
package state

import (
	"blockchain/foundation/blockchain/database"
	"blockchain/foundation/blockchain/genesis"
	"blockchain/foundation/blockchain/mempool"
	"blockchain/foundation/blockchain/peer"
	"sync"
)

// EventHandler defines a function that is called when events
// occur in the processing of persisting blocks.
type EventHandler func(v string, args ...any)

// Worker interface represents the behavior required to be implemented by any
// package providing support for mining, peer updates, and transaction sharing.
type Worker interface {
	Shutdown()
	Sync()
	SignalStartMining()
	SignalCancelMining()
	SignalShareTx(blockTx database.BlockTx)
}

// =============================================================================

// Config represents the configuration required to start
// the blockchain node.
type Config struct {
	BeneficiaryID  database.AccountID
	Host           string
	Storage        database.Storage
	Genesis        genesis.Genesis
	KnownPeers     *peer.PeerSet
	SelectStrategy string
	EvHandler      EventHandler
}

// State manages the blockchain database.
type State struct {
	mu sync.RWMutex

	beneficiaryID database.AccountID
	evHandler     EventHandler
	host          string

	knownPeers *peer.PeerSet
	storage    database.Storage
	genesis    genesis.Genesis
	mempool    *mempool.Mempool
	db         *database.Database

	Worker Worker
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
	db, err := database.New(cfg.Genesis, cfg.Storage, ev)
	if err != nil {
		return nil, err
	}

	// Construct a mempool with the specified sort strategy.
	mempool, err := mempool.NewWithStrategy(cfg.SelectStrategy)
	if err != nil {
		return nil, err
	}

	state := State{
		beneficiaryID: cfg.BeneficiaryID,
		storage:       cfg.Storage,
		evHandler:     ev,
		host:          cfg.Host,

		knownPeers: cfg.KnownPeers,
		genesis:    cfg.Genesis,
		mempool:    mempool,
		db:         db,
	}

	// The Worker is not set here. The call to worker.Run will assign itself
	// and start everything up and running for the node.

	return &state, nil
}

func (s *State) Shutdown() error {
	s.evHandler("state: shutdown: started")
	defer s.evHandler("state: shutdown: completed")

	// Make sure the database file is properly closed.
	// defer func() {
	// 	s.db.Close()
	// }()

	// Stop all blockchain writing activity.
	s.Worker.Shutdown()

	return nil
}

// Host returns a copy of host information.
func (s *State) Host() string {
	return s.host
}

// LatestBlock returns a copy the current latest block.
func (s *State) LatestBlock() database.Block {
	return s.db.LatestBlock()
}

// MempoolLength returns the current length of the mempool.
func (s *State) MempoolLength() int {
	return s.mempool.Count()
}

// Mempool returns a copy of the mempool.
func (s *State) Mempool() []database.BlockTx {
	return s.mempool.PickBest()
}

// UpsertMempool adds a new transaction to the mempool.
func (s *State) UpsertMempool(tx database.BlockTx) error {
	return s.mempool.Upsert(tx)
}

// Accounts returns a copy of the database accounts.
func (s *State) Accounts() map[database.AccountID]database.Account {
	return s.db.Copy()
}

// Genesis returns a copy of the genesis information.
func (s *State) Genesis() genesis.Genesis {
	return s.genesis
}

// AddKnownPeer provides the ability to add a new peer to
// the known peer list.
func (s *State) AddKnownPeer(peer peer.Peer) bool {
	return s.knownPeers.Add(peer)
}

// RemoveKnownPeer provides the ability to remove a peer from
// the known peer list.
func (s *State) RemoveKnownPeer(peer peer.Peer) {
	s.knownPeers.Remove(peer)
}

// KnownExternalPeers retrieves a copy of the known peer list without
// including this node.
func (s *State) KnownExternalPeers() []peer.Peer {
	return s.knownPeers.Copy(s.host)
}

// KnownPeers retrieves a copy of the full known peer list which includes
// this node as well. Used by the PoA selection algorithm.
func (s *State) KnownPeers() []peer.Peer {
	return s.knownPeers.Copy("")
}