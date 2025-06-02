package state

import "blockchain/foundation/blockchain/database"

// QueryLastest represents to query the latest block in the chain.
const QueryLastest = ^uint64(0) >> 1

// QueryAccount returns a copy of the account from the database.
func (s *State) QueryAccount(account database.AccountID) (database.Account, error) {
	return s.db.Query(account)
}

// QueryBlocksByNumber returns the set of blocks based on block numbers. This
// function reads the blockchain from disk first.
func (s *State) QueryBlocksByNumber(from uint64, to uint64) []database.Block {
	if from == QueryLastest {
		from = s.db.LatestBlock().Header.Number
		to = from
	}
	if to == QueryLastest {
		to = s.db.LatestBlock().Header.Number
	}

	var out []database.Block
	for i := from; i <= to; i++ {
		block, err := s.db.GetBlock(i)
		if err != nil {
			s.evHandler("state: getblock: ERROR: %s", err)
			return nil
		}
		out = append(out, block)
	}

	return out
}