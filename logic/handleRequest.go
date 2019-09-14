package logic

import (
	"bytes"
	"fmt"
	"github.com/D-CDC/cdc-backend/common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const key = `{
    "version": 3,
    "id": "1df6bf61-859d-43e7-b287-f46c5aa46232",
    "address": "4a83d2b8a7ae9e4c4e5efc805b0b016d55873e91",
    "crypto": {
        "ciphertext": "94aff10e5992018ff68333381d5426ee713eef506db539af06fe7464ec3abc8b",
        "cipherparams": {
            "iv": "47f6fb4beef15edf940afb5d6c952d8c"
        },
        "cipher": "aes-128-ctr",
        "kdf": "scrypt",
        "kdfparams": {
            "dklen": 32,
            "salt": "f4c72118fd0c9354fcc77006c5d3a9f9e0a3ba9e7da83f342ddfce07ba18a558",
            "n": 8192,
            "r": 8,
            "p": 1
        },
        "mac": "dbfa7fa5e162a3afdedb3ce3b0b043241596f2d03fd15076e5e7d5624f88e970"
    }
}`

//upload file
func Upload(filename string) (response string, statusCode int, err error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	fileWriter, err := w.CreateFormFile(common.FormFile, filename)
	if err != nil {
		panic(err)
	}

	fh, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		panic(err)
	}
	w.Close()

	request, err := http.NewRequest(common.HttpPost, common.IpiIPFSAdd, buf)
	if err != nil {
		panic(err)
	}

	request.Header.Set(common.HttpHeadType, w.FormDataContentType())
	var client http.Client
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	resbuf := new(bytes.Buffer)
	resbuf.ReadFrom(res.Body)
	response = resbuf.String()
	return
}

//pin file
func PinAdd(hash string) (response string, err error) {
	buf := new(bytes.Buffer)
	r := multipart.NewWriter(buf)
	defer r.Close()

	request, err := http.NewRequest("POST", "http://localhost:5001/api/v0/pin/add?arg="+hash, buf)
	if err != nil {
		panic(err)
	}
	var client http.Client
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	resbuf := new(bytes.Buffer)
	resbuf.ReadFrom(res.Body)
	response = resbuf.String()
	return
}

//download file
func Download(hash string, filepath string) (err error) {
	// Create buffer
	buf := new(bytes.Buffer) // caveat IMO dont use this for large files, \
	// create a tmpfile and assemble your multipart from there (not tested)
	r := multipart.NewWriter(buf)

	defer r.Close()
	req, err := http.NewRequest("POST", "http://localhost:5001/api/v0/cat?arg="+hash, buf)
	if err != nil {
		panic(err)
	}

	var client http.Client
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(file, res.Body) // Replace this with Status.Code check
	fmt.Println("response ", " statusCode ", res.Status)
	return err
}

//get node info
func Id() (response string, err error) {
	buf := new(bytes.Buffer)
	r := multipart.NewWriter(buf)
	defer r.Close()

	request, err := http.NewRequest("POST", "http://localhost:5001/api/v0/id", buf)
	if err != nil {
		panic(err)
	}

	var client http.Client
	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	resbuf := new(bytes.Buffer)
	resbuf.ReadFrom(res.Body)
	response = resbuf.String()
	return
}
