package contract

import (
	"fmt"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/truechain/truechain-engineering-code/rpc"
	"strings"
)

func SendTransaction(method string, hash string) string {

	client, err := rpc.Dial(common.TrueDialAddress)

	defer client.Close()

	if err != nil {
		fmt.Println("Dail:", err.Error())
		return ""
	}
	_, err = unlockAccount(client, common.ADDRESS, common.PASSWORD, 9000000)

	var result string
	if strings.Contains(method, common.CmdIPFSAdd) {
		result, _ = sendRawTransactionUpload(client, common.ADDRESS, common.CONTRACTADDRESSS)
	} else {
		result, _ = sendTransactionDownload(client, common.ADDRESS, common.CONTRACTADDRESSS, "0x3f2")
	}
	return result
}

func sendRawTransactionUpload(client *rpc.Client, from string, to string) (string, error) {

	mapData := make(map[string]interface{})

	mapData[common.TXFrom] = from
	mapData[common.TXTo] = to
	mapData[common.TXInput] = common.ContractUpload
	mapData[common.TxGasPrice] = common.GasValue
	mapData[common.TxLimit] = common.GasLimit
	var result string
	err := client.Call(&result, common.ETRUESendTransaction, mapData)
	fmt.Println("result ", result, " err ", err)
	return result, err
}

func sendTransactionDownload(client *rpc.Client, from string, to string, value string) (string, error) {
	mapData := make(map[string]interface{})

	mapData[common.TXFrom] = from
	mapData[common.TXTo] = to
	mapData[common.TXValue] = value
	mapData[common.TXInput] = common.ContractDownload
	var result string
	err := client.Call(&result, common.ETRUESendTransaction, mapData)
	fmt.Println("result ", result, " err ", err)
	_, _ = hexutil.Decode(result)
	result = common.DownLoadValue
	return result, err
}

func unlockAccount(client *rpc.Client, account string, password string, time int) (bool, error) {
	var reBool bool
	err := client.Call(&reBool, common.ETRUEPersonalUnlockAccount, account, password, time)
	fmt.Println("personal_unlockAccount Ok", reBool)
	return reBool, err
}
