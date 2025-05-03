package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Tx struct {
	FromID string `json:"from"`
	ToID   string `json:"to"`
	Value  uint64 `json:"value"`
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	privateKey, err := crypto.LoadECDSA("zblock/accounts/kennedy.ecdsa")
	if err != nil {
		return fmt.Errorf("unable to load private key for node: %w", err)
	}

	tx := Tx{
		FromID: "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		ToID:   "Aaron",
		Value:  1000,
	}

	data, err := json.Marshal(tx)
	if err != nil {
		return fmt.Errorf("unable to marshal transaction: %w", err)
	}

	stamp := []byte(fmt.Sprintf("\x19Ardan Signed Message:\n%d", len(data)))

	v := crypto.Keccak256(stamp, data)

	sig, err := crypto.Sign(v, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign transaction: %w", err)
	}

	fmt.Println("Signature:", hexutil.Encode(sig))

	// ===========================================================================================
	// OVER THE WIRE

	publicKey, err := crypto.SigToPub(v, sig)
	if err != nil {
		return fmt.Errorf("unable to recover public key from signature: %w", err)
	}

	extractedAddress := crypto.PubkeyToAddress(*publicKey).String()

	fmt.Println("Public key2:", extractedAddress)

	// ===========================================================================================

	tx2 := Tx{
		FromID: "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		ToID:   "Frank",
		Value:  250,
	}

	data2, err := json.Marshal(tx2)
	if err != nil {
		return fmt.Errorf("unable to marshal transaction: %w", err)
	}

	stamp2 := []byte(fmt.Sprintf("\x19Ardan Signed Message:\n%d", len(data2)))

	v2 := crypto.Keccak256(stamp2, data2)

	sig2, err := crypto.Sign(v2, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign transaction: %w", err)
	}

	fmt.Println("Signature2:", hexutil.Encode(sig2))

	// ===========================================================================================
	// OVER THE WIRE

	tx3 := Tx{
		FromID: "0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		ToID:   "Frank",
		Value:  1000,
	}

	data3, err := json.Marshal(tx3)
	if err != nil {
		return fmt.Errorf("unable to marshal transaction: %w", err)
	}

	stamp3 := []byte(fmt.Sprintf("\x19Ardan Signed Message:\n%d", len(data3)))

	v3 := crypto.Keccak256(stamp3, data3)

	sig3, err := crypto.Sign(v3, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign transaction: %w", err)
	}

	publicKey2, err := crypto.SigToPub(v3, sig3)
	if err != nil {
		return fmt.Errorf("unable to recover public key from signature: %w", err)
	}

	extractedAddress = crypto.PubkeyToAddress(*publicKey2).String()

	fmt.Println("Public key3:", extractedAddress)

	if extractedAddress != tx3.FromID {
		return fmt.Errorf("extracted address does not match sender address")
	}

	return nil
}
