package rpc

import (
	"github.com/D-CDC/cdc-backend/car"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/D-CDC/cdc-backend/contract"
	"github.com/D-CDC/cdc-backend/logic"
	"strings"
)

func StartRpcServer(method string, hash string) {
	if strings.Contains(method, common.CmdIPFSAdd) {
		car.CreateUserInfo(hash)
		response, statusCode, _ := logic.Upload(hash)
		contract.SendTransaction(response, statusCode)

	} else if strings.Contains(method, common.CmdIPFSDownload) {
		logic.Download(hash, common.IPFSFileName)
	}
}
