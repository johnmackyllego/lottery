package main

import (
	"encoding/json"
	"fmt"
	/*"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"*/)

type PowerballChaincode struct {
}

type Tickets struct {
	TicketNo    int64   `json:"ticketno"`
	RoundNumber int64   `json:"roundnumbers"`
	Address     string  `json:"address"`
	Ticket      Numbers `json:"numbers"`
}

type Numbers struct {
	TicketNumbers int64      `json:"ticketnumbers"`
	Ticket        [6][]int64 `json:"ticket"`
}

type Rounds struct {
	RoundNo        int64    `json:"roundno"`
	Endtime        int64    `json:"endtime"`
	DrawBlock      int64    `json:"drawblock"`
	WinningNumbers [6]int64 `json:"winningnumbers"`
}

type InitDetails struct {
	ticketPrice        int64
	maxNumber          int64
	maxPowerballNumber int64
	roundLength        int64
}

var ticketPrice int64
var maxNumber int64
var maxPowerballNumber int64
var roundLength int64
var round int64
var ticketNo int64

func (p *PowerballChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	ticketPrice = 1
	maxNumber = 69
	maxPowerballNumber = 26
	roundLength = 15
	round = 1
	ticketNo = 0

	return shim.Success(nil)
}

func (p *PowerballChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fc, args := stub.GetFunctionAndParameters()
	if fc == "displayDetails" {
		return p.displayDetails(stub)
	} else if fc == "displayRound" {
		return p.round(stub)
	} else if fc == "rounds" {
		return p.rounds(stub, args)
	} else if fc == "ticketsFor" {
		return p.ticketsFor(stub, args)
	} else if fc == "winningNumbersFor" {
		return p.winningNumbersFor(stub, args)
	} else if fc == "drawNumbers" {
		return p.drawNumbers(stub, args)
	} else if fc == "claim" {
		return p.claim(stub, args)
	} else if fc == "buy" {
		return p.buy(stub, args)
	}

	return shim.Error("Invalid function name" + fc)
}

//to be tested
//this function return all the default details
func (p *PowerballChaincode) displayDetails(stub shim.ChaincodeStubInterface) pb.Response {

	details := InitDetails{ticketPrice, maxNumber, maxPowerballNumber, roundLength}
	detailsAsBytes, _ := json.Marshal(details)

	return shim.Success(detailsAsBytes)
}

//to be tested
//this function returns current round
func (p *PowerballChaincode) round(stub shim.ChaincodeStubInterface) pb.Response {

	roundAsBytes, _ := json.Marshal(round)

	return shim.Success(roundAsBytes)

}

//unfinished function
//function should return details about the round input
func (p *PowerballChaincode) rounds(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args != 1) {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	r := args[0]

	return shim.Success(r)

}

//function to associate tickets to a user address
//function should return Tickets struct
func (p *PowerballChaincode) ticketsFor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args != 2) {
		return shim.Error("Incorrect number of arguments. Expecting s2")
	}

	var ticket = Tickets{

		TicketNo:    ticketNo,
		RoundNumber: args[0],
		Address:     args[1],
		Ticket:      null,
	}

	ticketAsBytes, _ := json.Marshal(ticket)
	if stub.PutState(ticketNo, ticketAsBytes) == true {
		ticketNo++
	}

	return shim.Success(nil)

}

//function should return winning numbers based on round number
func (p *PowerballChaincode) winningNumbersFor(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args != 1) {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(PowerballChaincode))
	if err != nil {
		fmt.Printf("Error creating smartcontract ", err.Error())
	}
}
