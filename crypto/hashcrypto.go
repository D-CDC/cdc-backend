package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"github.com/D-CDC/cdc-backend/common"
)

//plainText fill
func Padding(plainText []byte, blockSize int) []byte {
	//cal len
	n := blockSize - len(plainText)%blockSize
	//fil n
	temp := bytes.Repeat([]byte{byte(n)}, n)
	plainText = append(plainText, temp...)
	return plainText
}

//delete fill
func UnPadding(cipherText []byte) []byte {
	//get last one byte
	end := cipherText[len(cipherText)-1]
	//delete fil
	cipherText = cipherText[:len(cipherText)-int(end)]
	return cipherText
}

//AEC Encrypt（CBC mode）
func AESCbCEncrypt(plainText []byte, key []byte) []byte {
	//return AES Block interface
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//fill
	plainText = Padding(plainText, block.BlockSize())
	//assign vector vi,len and accordance block
	iv := []byte(common.VECVI)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText
}
func AESCbCDecrypt(cipherText []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := []byte(common.VECVI)
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = UnPadding(plainText)
	return plainText
}
