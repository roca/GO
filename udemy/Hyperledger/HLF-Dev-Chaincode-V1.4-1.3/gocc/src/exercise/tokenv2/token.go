package main

/**
 * tokenv2
 * Shows the
 *    A) Use of Logger
 **/
import (
	"fmt"
	"time"

	// The shim package
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// peer.Response is in the peer package
	"github.com/hyperledger/fabric/protos/peer"
)

// TokenChaincode Represents our chaincode object
type TokenChaincode struct {
}

// V2

// ChaincodeName - Create an instance of the Logger
const ChaincodeName = "tokenv2"

var logger = shim.NewLogger(ChaincodeName)

// Init Implements the Init method
func (token *TokenChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// Simply print a message
	// fmt.Println("Init executed")

	logger.Debug("Init executed v2 - DEBUG")

	// Return success
	return shim.Success(nil)
}

// Invoke method
func (token *TokenChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("Invoke executed ")

	fmt.Printf("GetTxTID() %s\n", stub.GetTxID())
	timestamp, _ := stub.GetTxTimestamp()
	// if err != nil {
	// 	fmt.Errorf("GetTxTimestamp Error: %s\n", err.Error())
	// }
	fmt.Printf("GetTxTimestamp() %s\n", time.Unix(timestamp.GetSeconds(), 0))
	fmt.Printf("GetTxTID() %s\n", stub.GetChannelID())

	// V3
	// Extract the information from proposal
	PrintSignedProposalInfo(stub)

	// V3
	// Will receieve empty map unless client set the transient data in Tx Proposal
	// transientData, _ := stub.GetTransient()
	// fmt.Println("GetTransient() =>", transientData)

	PrintCreatorInfo(stub)

	return shim.Success(nil)
}

// Chaincode registers with the Shim on startup
func main() {
	fmt.Printf("Started Chaincode. token/v2")
	err := shim.Start(new(TokenChaincode))
	if err != nil {
		//fmt.Printf("Error starting chaincode: %s", err)
		// V2
		logger.Error("Error starting chaincode: %s", err)
	}
}
