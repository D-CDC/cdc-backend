package main

import (
	"fmt"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/D-CDC/cdc-backend/rpc"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("invalid args : %s [method] [file/hash] [\"port\"]\n", os.Args[0])
		return
	}

	checkIP(os.Args)

	rpc.StartRpcServer(os.Args[1], os.Args[2])
}

func checkIP(params []string) error {
	ip := common.IPFSDialAddress
	if len(params) == 4 {
		ip = ip + os.Args[6]
	} else {
		ip = ip + common.IPFSDefaultPort
	}
	common.DefaultSocketConfig.RpcAddress = ip
	return nil
}
