package parse

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type IPFSResult struct {
	Name string `json:"name"      gencodec:"required"`
	Hash string `json:"hash" 	  gencodec:"required"`
	Size string `json:"size"      gencodec:"required"`
}

func ParseResponse(response string, statusCode int) string {
	var result = IPFSResult{}
	err := json.Unmarshal(bytes.NewBufferString(response).Bytes(), &result)
	if err != nil {
		println("error ", err.Error())
	}
	fmt.Println("response ", response, " status_code ", statusCode, "result ", result.Hash)
	return result.Hash
}
