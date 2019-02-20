package main

import (
	"fmt"
	"time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	 pb "github.com/hyperledger/fabric/protos/peer"
)

type recurringLottery struct{

	
}
type Round struct {
    endBlock string `json:"endBlock"`
	drawBlock string  `json:"drawBlock"`
	//uint Entry[] entries 
	uint totalQuantity;
	Winner string `json:"Winner"`
	dispalyAll string `json:"displayAll"`
	betAgain string `json:"betAgain"`
	buyTicket string `json:"buyTicket"`
}

func (r * RecurringlotteryChaincode) Init(stub shim.ChaincodeStubInterface)rc.Response{
	return shim.sucess(nil)
}

func (r * RecurringloterryChaincode) Invoke (stub shim.ChaincodeStubInterface)rc.Response{

	function, args := APIstub.GetFunctionAndParameters()
	if function == "initLedger"{
		return rc.initLedger(APIstub)
	}else if function == "queryLottery"{
		return rc.queryLottery(APIStub)
	}else if function == "drawWinner"{
		return rc.drawWinner(APIstub)
	}else if function == "getHistory"{
		return rc.getHistory(APIstub)
	}else if function == "displayAll"{
		return rc.displayRound(APIStub)	
	}else if function =="buyTicket"{
		return rc.buyTicket(APIstub)
	}
	return shim.Error("Invalid Smart Contract function name.")
} 

func (r * RecurringloterryChaincode) initLedger (APIstub shim.ChaincodeStubInterface) rc.Response{
	lottery = []Lottery {
		Lottery{BilledTo: "PCSO", InvoiceData: "2/20/2019", buyTicket:"200", Ticketnumber: "001"},
	}

	i := 0
	for i < len(lottery){
		fmt.Println("i is ", i)
		lotteryAsBytes, _ := json.Marshal(lottery[i])
		APIstub.PutState("INV"+strconv.Itoa(i), lotteryAsBytes)
		fmt.Println("Added", lottery[i])
		i = i + 1
	}
	return shim.Success(nil)
}
func (r *RecurringloterryChaincode) Total(APIstub shim.ChaincodeStubInterface) rc.Response {

	startKey := "INV0"
	endKey := "INV999"
	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":") 
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}\n")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	
	fmt.Printf("- displayAllStake:\n%s\n", buffer.String())


	return shim.Success(buffer.Bytes())


	func (r *RecurringloterryChaincode) reEntry(APIstub shim.ChaincodeStubInterface, args []string) rc.Response {

		if len(args) != 2 {
			return shim.Error("Incorrect number of arguments. Expecting 2")
		}
	
		lotteryAsBytes, _ := APIstub.GetState(args[0])
		lottery := Lottery{}
	
		json.Unmarshal(invoiceAsBytes, &invoice)
		lottery.ReEntry = args[1]
	
		invoiceAsBytes, _ = json.Marshal(invoice)
		APIstub.PutState(args[0], lotteryAsBytes)

		return shim.Success(nil)
}



