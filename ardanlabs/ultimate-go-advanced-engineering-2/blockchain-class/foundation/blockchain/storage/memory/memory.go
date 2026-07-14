// Package memory implements the ability to read and write blocks to memory
// using a slice.
package memory

import (
	"blockchain/foundation/blockchain/database"
	"sync"
)

// Memory represents the serialization implementation for reading and storing
// blocks in memory using a slice. This implements the database.Storage
// interface.
type Memory struct {
	mu     sync.RWMutex
	blocks []database.BlockData
}
