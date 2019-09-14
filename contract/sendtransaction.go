package contract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/truechain/truechain-engineering-code/rpc"
)

type IPFSResult struct {
	Name string `json:"name"      gencodec:"required"`
	Hash string `json:"hash" 	  gencodec:"required"`
	Size string `json:"size"      gencodec:"required"`
}

func SendTransaction(response string, statusCode int) {

	var result = IPFSResult{}
	json.Unmarshal(bytes.NewBufferString(response).Bytes(), &result)

	fmt.Println("response ", response, " status_code ", statusCode, "result ", result.Hash)

	client, err := rpc.Dial(common.TrueDialAddress)

	defer client.Close()

	if err != nil {
		fmt.Println("Dail:", err.Error())
		return
	}

	_, err = unlockAccount(client, common.ADDRESS, common.PASSWORD, 9000000)

	sendRawTransaction(client, common.ADDRESS, common.CONTRACTADDRESSS, "0x3000000")
}

func sendRawTransaction(client *rpc.Client, from string, to string, value string) (string, error) {

	mapData := make(map[string]interface{})

	mapData[common.TXFrom] = from
	mapData[common.TXTo] = to
	mapData[common.TXInput] = "0xa9059cbb00000000000000000000000061549f53bb59b7eafffbab82ff24addd54a9b8060000000000000000000000000000000000000000000000056bc75e2d63100000"
	var result string
	err := client.Call(&result, common.ETRUESendTransaction, mapData)
	fmt.Println("result ", result, " err ", err)
	return result, err
}

func combineInput(method string, hash string) string {
	data := fmt.Sprintf("%x", hash)
	return method + data
}

func unlockAccount(client *rpc.Client, account string, password string, time int) (bool, error) {
	var reBool bool
	err := client.Call(&reBool, common.ETRUEPersonalUnlockAccount, account, password, time)
	fmt.Println("personal_unlockAccount Ok", reBool)
	return reBool, err
}
