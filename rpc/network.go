package rpc

import (
	"github.com/D-CDC/cdc-backend/car"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/D-CDC/cdc-backend/contract"
	"github.com/D-CDC/cdc-backend/crypto"
	"github.com/D-CDC/cdc-backend/logic"
	"github.com/D-CDC/cdc-backend/parse"
	"strings"
)

func StartRpcServer(method string, hash string) {
	if strings.Contains(method, common.CmdIPFSAdd) {
		car.CreateUserInfo(hash)
		response, statusCode, _ := logic.Upload(hash)
		result := parse.ParseResponse(response, statusCode)
		contract.SendTransaction(method, result)

	} else if strings.Contains(method, common.CmdIPFSDownload) {
		//contract.SendTransaction(cipherText)
		crypto.AESCbCDecrypt([]byte(hash), []byte(common.CipherKey))
		logic.Download(hash, common.IPFSFileName)
	}
}
