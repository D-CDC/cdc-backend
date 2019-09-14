package rpc

import (
	"github.com/D-CDC/cdc-backend/logic"
	"strings"
)

func StartRpcServer(method string, hash string) {
	if strings.Contains(method, "add") {
		logic.Upload(hash)
	} else if strings.Contains(method, "download") {
		logic.Download(hash, "name")
	}
}
