package main

import (
	"fmt"

	// // The shim package - needed for shim
	"github.com/hyperledger/fabric/core/chaincode/shim"

	// ChaincodeProposalPayload buffer is defined in this
	"github.com/hyperledger/fabric/protos/peer"

	// Needed for common proto buffers unmarshalling
	"github.com/golang/protobuf/proto"

	// Multiple common proto header
	"github.com/hyperledger/fabric/protos/common"
)

// PrintSignedProposalInfo prints the info to console
func PrintSignedProposalInfo(stub shim.ChaincodeStubInterface) {

	fmt.Println("PrintSignedProposalInfo() executed ")

	// Get the SignedProposal
	// SignedProposal has 2 parts
	// 1. ProposalBytes
	// 2. Signature
	signedProposal, _ := stub.GetSignedProposal()
	data := signedProposal.GetProposalBytes()
	proposal := &peer.Proposal{}
	proto.Unmarshal(data, proposal)

	// Proposal has 2 parts
	// 1. Header
	// 2. Payload - the structure for this depends on the type in the ChannelHeader
	header := &common.Header{}
	proto.Unmarshal(proposal.GetHeader(), header)

	// Header has 2 parts
	// 1. ChannelHeader
	// 2. SignatureHeader
	channelHeader := &common.ChannelHeader{}
	proto.Unmarshal(header.GetChannelHeader(), channelHeader)
	fmt.Println("channelHeader.GetType() => ", common.HeaderType(channelHeader.GetType()))
	fmt.Println("channelHeader.GetChannelId() => ", channelHeader.GetChannelId())
}

// Prints the creator information
func PrintCreatorInfo(stub shim.ChaincodeStubInterface) {
	fmt.Println("PrintCreatorInfo() executed ")

	byteData, _ := stub.GetCreator()

	fmt.Println("PrintCreatorInfo => ", string(byteData))
}
