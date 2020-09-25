package controllers

import (
	"bufio"
	"os"

	"github.com/astaxie/beego"
)

type FuzzDataController struct {
	beego.Controller
}

func (c *FuzzDataController) Get() {
	FuzzStop()
	Version := c.GetString("Version")
	gopath := os.Getenv("GOPATH")
	var FuzzURL string
	if Version == "1.1" {
		FuzzURL = gopath + "/src/example/server/models/fabric1.1"
	} else {
		FuzzURL = gopath + "/src/github.com/hyperledger/fabric/integration/chaincode/aab"
	}
	fileURL := FuzzURL + "/output.log"
	data := ReadFile(fileURL)
	c.Ctx.WriteString(data)

}
func FuzzStop() {
	a := exeSysCommand("ps -ef|grep go-fuzz|awk 'NR==1{print $2}'")
	_ = exeSysCommand("kill " + a)
}

func ReadFile(file_name string) (info string) {
	file, err := os.Open(file_name)
	if err != nil {
		return "error"
	}
	defer file.Close()
	var lineText string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText = scanner.Text()
	}

	return string(lineText)
}
