package main

import (
	"car-ipfs/common"
	"fmt"
	"github.com/D-CDC/cdc-backend/rpc"
	"os"
)

//Count send complete
var Count int64

// get par
func main() {
	if len(os.Args) < 3 {
		fmt.Printf("invalid args : %s [method] [file/hash] [\"port\"]\n", os.Args[0])
		return
	}

	ip := "127.0.0.1:"
	if len(os.Args) == 4 {
		ip = ip + os.Args[6]
	} else {
		ip = ip + "5001"
	}
	common.DefaultSocketConfig.RpcAddress = ip

	rpc.StartRpcServer(os.Args[1], os.Args[2])
}
