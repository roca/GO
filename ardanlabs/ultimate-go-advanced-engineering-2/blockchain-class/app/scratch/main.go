package main

import (
	"blockchain/foundation/blockchain/database"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/big"

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

	v1 := crypto.Keccak256(stamp, data)

	sig, err := crypto.Sign(v1, privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign transaction: %w", err)
	}

	fmt.Println("Signature:", hexutil.Encode(sig))

	// ===========================================================================================
	// OVER THE WIRE

	publicKey, err := crypto.SigToPub(v1, sig)
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

	v, r, s, err := ToVRSFromHexSignature(hexutil.Encode(sig3))
	if err != nil {
		return fmt.Errorf("unable to convert hex signature to v, r, s: %w", err)
	}

	fmt.Println("V|R|S:", v, r, s)

	fmt.Println("+++++++++++++++++++++++++++++++TX+++++++++++++++++++++++++++++++")

	billTx, err := database.NewTx(1, 1,
		"0xF01813E4B85e178A83e29B8E7bF26BD830a25f32",
		"0xbEE6ACE826eC3DE1B6349888B9151B92522F7F76",
		1000,
		0,
		nil,
	)
	if err != nil {
		return fmt.Errorf("unable to create new transaction: %w", err)
	}
	signedTx, err := billTx.Sign(privateKey)
	if err != nil {
		return fmt.Errorf("unable to sign transaction: %w", err)
	}
	fmt.Println("SignedTx:", signedTx)

	return nil
}

// ToVRSFromHexSignature converts a hex representation of the signature into
// its R, S and V parts.
func ToVRSFromHexSignature(sigStr string) (v, r, s *big.Int, err error) {
	sig, err := hex.DecodeString(sigStr[2:])
	if err != nil {
		return nil, nil, nil, err
	}

	r = big.NewInt(0).SetBytes(sig[:32])
	s = big.NewInt(0).SetBytes(sig[32:64])
	v = big.NewInt(0).SetBytes([]byte{sig[64]})

	return v, r, s, nil
}
