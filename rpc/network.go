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

func StartRpcServer(method string, params string) {
	if strings.Contains(method, common.CmdIPFSAdd) {
		switch params {
		case common.ADDCar:
			car.CreateUserInfo(params)
			response, statusCode, _ := logic.Upload(params, nil)
			result := parse.ParseResponse(response, statusCode)
			contract.SendTransaction(method, result)
			break
		case common.ADDDrive:
			data := car.CreateUserDrive(params + common.SuffixJpg)
			response, statusCode, _ := logic.Upload(params, data)
			result := parse.ParseResponse(response, statusCode)
			contract.SendTransaction(method, result)
			break
		}
	} else if strings.Contains(method, common.CmdIPFSDownload) {
		switch params {
		case common.ADDCar:
			result := contract.SendTransaction(method, params)
			data, _ := logic.Download(result, common.IPFSFileName)
			dataOrigin := crypto.AESCbCDecrypt(data, []byte(common.CipherKey))
			car.ConvertToFile(dataOrigin, params+common.SuffixTxt)
			break
		case common.ADDDrive:
			result := contract.SendTransaction(method, params)
			data, _ := logic.Download(result, common.IPFSFileName)
			dataOrigin := crypto.AESCbCDecrypt(data, []byte(common.CipherKey))
			car.ConvertToFile(dataOrigin, params+common.SuffixJpg)
		}
	}
}
