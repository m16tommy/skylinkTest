package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type Test struct {
}
type MsgObj struct {
	MsgNo       string
	MsgType     string
	MsgAmount   string
	MsgCreateBy string
	Sender      string
	Receiver    string
	TimeStamp   string // This is the time stamp
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

func main() {
	args := []string{
		"1",
		"850",
		"1000",
		"tommy",
	}
	fmt.Print(args)
	msgObj := argsToMsgObj(args)

	jStr, err := json.Marshal(msgObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jStr))
}
