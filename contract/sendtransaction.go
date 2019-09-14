package contract

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/truechain/truechain-engineering-code/rpc"
)

type IPFSResult struct {
	Name string `json:"name"    gencodec:"required"`
	Hash string `json:"hash" gencodec:"required"`
	Size string `json:"size"      gencodec:"required"`
}

func SendTransaction(response string, statusCode int) {

	var result = IPFSResult{}
	json.Unmarshal(bytes.NewBufferString(response).Bytes(), &result)
	fmt.Println("response ", response, " status_code ", statusCode, "result ", result.Hash)

	client, err := rpc.Dial("http://39.100.35.164:8888")

	defer client.Close()

	if err != nil {
		fmt.Println("Dail:", err.Error())
		return
	}

	//json, err := ioutil.ReadAll(strings.NewReader(key))
	//if err != nil {
	//    return
	//}
	//key, err := keystore.DecryptKey(json, "19870101zby")
	//if err != nil {
	//    return
	//}

	_, err = unlockAccount(client, "0x3e21a8702212fd571e23a21086a76b9592d8100a", "123456", 9000000, "main")

	sendRawTransaction(client, "0x3e21a8702212fd571e23a21086a76b9592d8100a", "0x189cF5e2079686ff909EfdA350D5c66d28B6857a", "0x3000000")
}

func sendRawTransaction(client *rpc.Client, from string, to string, value string) (string, error) {

	mapData := make(map[string]interface{})

	mapData["from"] = from
	mapData["to"] = to
	mapData["input"] = "0xa9059cbb00000000000000000000000061549f53bb59b7eafffbab82ff24addd54a9b8060000000000000000000000000000000000000000000000056bc75e2d63100000"
	var result string
	err := client.Call(&result, "etrue_sendTransaction", mapData)
	fmt.Println("result ", result, " err ", err)
	return result, err
}

func unlockAccount(client *rpc.Client, account string, password string, time int, name string) (bool, error) {
	var reBool bool
	err := client.Call(&reBool, "personal_unlockAccount", account, password, time)
	fmt.Println(name, " personal_unlockAccount Ok", reBool)
	return reBool, err
}
