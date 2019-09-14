package car

import (
	"fmt"
	"github.com/D-CDC/cdc-backend/common"
	"github.com/D-CDC/cdc-backend/crypto"
	"os"
	"strconv"
)

type InfoCar struct {
	model    string
	age      int64
	position string
	speed    int64
}

func CreateUserInfo(fileName string) {
	cipherText := crypto.AESCbCEncrypt([]byte(readUserInfo().toString()), []byte(common.CipherKey))
	createFileAndWrite(fileName, cipherText)
}

func createFileAndWrite(fileName string, cipher []byte) {
	f, err := os.Create(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		_, err = f.Write(cipher)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func readUserInfo() *InfoCar {
	var model string
	var age int64
	var position string
	var speed int64
	var info = &InfoCar{}
	fmt.Println("please input car model age")
	_, _ = fmt.Scanf("%s %d %f", &model, &age)
	fmt.Println("please input position speed")
	_, _ = fmt.Scanf("%s %d %f", &position, &speed)
	info.model, info.age, info.position, info.speed = model, age, position, speed
	return info
}

func (f InfoCar) toString() string {
	return f.model + strconv.FormatInt(f.age, 10) + f.position + strconv.FormatInt(f.speed, 10)
}
