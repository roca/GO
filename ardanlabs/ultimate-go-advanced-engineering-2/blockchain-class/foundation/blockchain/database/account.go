package database

// AccountID represents an account id that is used to sign transactions and is
// associated with transactions on the blockchain. This will be the last 20
// bytes of the public key.
type AccountID string
