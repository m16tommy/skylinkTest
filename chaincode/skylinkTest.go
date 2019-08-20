package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type MsgObj struct {
	MsgNo       string
	MsgType     string
	MsgAmount   string
	MsgCreateBy string
	Sender      string
	Receiver    string
	TimeStamp   string // This is the time stamp
}
type SimpleChaincode struct {
}

func main() {

	// maximize CPU usage for maximum performance
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("Starting Item Auction Application chaincode BlueMix ver 21 Dated 2016-07-02 09.45.00: ")

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("Error starting Item Fun Application chaincode: %s", err)
	}

}
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("[Trade and Auction Application] Init")
	fmt.Println("\nInit() Initialization Complete ")
	return shim.Success(nil)
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()
	fmt.Println("==========================================================")
	fmt.Println("BEGIN Function ====> ", function)
	if function == "insert" {
		return t.insert(stub, args)
	} else if function == "delete" {
		return t.delete(stub, args)
	} else if function == "query" {
		return t.query(stub, args)
	} else if function == "update" {
		return t.update(stub, args)
	}

	fmt.Println("==========================================================")

	return shim.Error("Invoke: Invalid Function Name - function names begin with a q or i")

}
func argsToMsgObj(args []string) MsgObj {
	msgObj := MsgObj{}
	fields := reflect.TypeOf(msgObj)
	values := reflect.ValueOf(msgObj)

	for i := 0; i < len(args); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		obj := reflect.Indirect(reflect.ValueOf(&msgObj))
		obj.FieldByName(field.Name).SetString(args[i])
		fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")
	}
	now := time.Now()
	msgObj.TimeStamp = now.String()
	return msgObj
}

func (t *SimpleChaincode) insert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	msgObj := argsToMsgObj(args)

	jStr, err := json.Marshal(msgObj)
	if err != nil {
		fmt.Println(err)
	}

	err = stub.PutState(msgObj.MsgNo, []byte(jStr))
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("insert jsonStr : " + string(jStr))

	return shim.Success(nil)
}

func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	msgObj := argsToMsgObj(args)
	err := stub.DelState(msgObj.MsgNo)

	if err != nil {
		return shim.Error("Failed to delete state")
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	msgObj := argsToMsgObj(args)
	msgbytes, err := stub.GetState(msgObj.MsgNo)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + msgObj.MsgNo + "\"}"
		return shim.Error(jsonResp)
	}

	fmt.Println("query jsonStr : " + string(msgbytes))
	return shim.Success(msgbytes)

}

func (t *SimpleChaincode) update(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success(nil)
}
