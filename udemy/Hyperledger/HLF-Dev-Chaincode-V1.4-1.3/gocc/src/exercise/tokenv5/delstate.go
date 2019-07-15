package		main

// Solution for the exercise
import (
	"fmt"

	// // The shim package - needed for shim
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// ChaincodeProposalPayload buffer is defined in this
	"github.com/hyperledger/fabric/protos/peer"

)

// DeleteToken deletes the token from the database
// V5
// Returns 0 if successful
func  DeleteToken(stub shim.ChaincodeStubInterface) peer.Response {

	// Check if the key exists - if not then return false
	value, _ := stub.GetState("MyToken")
	if value == nil {
		return shim.Success([]byte("false"))
	}
	// Delete the key
	if err := stub.DelState("MyToken"); err != nil {
		fmt.Println("Delete Failed!!! ", err.Error())
		return shim.Error(("Delete Failed!! "+err.Error()+"!!!"))
	}

	return shim.Success([]byte("true"))
}