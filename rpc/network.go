package rpc

import (
	"fmt"
	"github.com/D-CDC/cdc-backend/car"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/D-CDC/cdc-backend/contract"
	"github.com/D-CDC/cdc-backend/crypto"
	"github.com/D-CDC/cdc-backend/logic"
	"github.com/D-CDC/cdc-backend/parse"
	"strings"
)

func StartRpcServer(method string, params string) {
	if strings.Contains(method, common.CmdIPFSAdd) {
		car.CreateUserInfo(params)
		response, statusCode, _ := logic.Upload(params)
		result := parse.ParseResponse(response, statusCode)
		contract.SendTransaction(method, result)

	} else if strings.Contains(method, common.CmdIPFSDownload) {
		result := contract.SendTransaction(method, params)
		data := crypto.AESCbCDecrypt([]byte(result), []byte(common.CipherKey))
		_ = logic.Download(fmt.Sprintf("%s", data), common.IPFSFileName)
	}
}
