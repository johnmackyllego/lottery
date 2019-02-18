package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PowerballChaincode struct {
}

type Round struct {
	Endtime        int64    `json:"endtime"`
	DrawBlock      int64    `json:"drawblock"`
	WinningNumbers [6]int64 `json:"winningnumbers"`
	Tickets        Ticket   `json:"tickets"`
}

type Ticket struct {
	TicketNo [6][]int64 `json:"ticketno"`
}

func (p *PowerballChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// var ticketPrice int64 = 1
	// var maxNumber int64 = 69
	// var maxPowerballNumber int64 = 26
	// var roundLength int64 = 3
	return shim.Success(nil)
}

func (p *PowerballChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fc, args := stub.GetFunctionAndParameters()
	if fc == "buy" {
		return p.buy(stub, args)
	} else if fc == "drawNumbers" {
		return p.drawNumbers(stub, args)
	} else if fc == "claim" {
		return p.claim(stub, args)
	} else if fc == "ticketsFor" {
		return p.ticketsFor(stub, args)
	} else if fc == "winningNumbersFor" {
		return p.winningNumbersFor(stub, args)
	}

	return shim.Error("Invalid function name" + fc)
}

func (p *PowerballChaincode) buy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("Success buy!"))
}

func (p *PowerballChaincode) drawNumbers(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("Success draw numbers!"))
}

func (p *PowerballChaincode) claim(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("Success claim!"))
}

func (p *PowerballChaincode) ticketsFor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("Success tickets for~!"))
}

func (p *PowerballChaincode) winningNumbersFor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("Success winning numbers for ~!"))
}

func main() {
	err := shim.Start(new(PowerballChaincode))
	if err != nil {
		fmt.Printf("Error creating smartcontract ", err.Error())
	}
}
