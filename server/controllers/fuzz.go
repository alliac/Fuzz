package controllers

import (
	"os/exec"
	//"strings"
	"io/ioutil"
	"os"

	"github.com/astaxie/beego"
)

type FuzzController struct {
	beego.Controller
}

func (c *FuzzController) GetData() {

	InitialTestcase := "`" + c.GetString("InitialCase") + "`"
	upload := c.GetString("upload")
	Version := c.GetString("Version")
	gopath := os.Getenv("GOPATH")
	var FuzzURL string
	if Version == "2.0" {
		FuzzURL = gopath + "/src/github.com/hyperledger/fabric/integration/chaincode/aab"
		_ = exeSysCommand("cd " + gopath + "/src/github.com/hyperledger;sudo mv fabric fabric1.1")
		_ = exeSysCommand("cd " + gopath + "/src/github.com/hyperledger;sudo mv fabric2.0 fabric")
	} else if Version == "1.1" {
		FuzzURL = gopath + "/src/example/server/models/fabric1.1"
	}
	caseURL := FuzzURL + "/gen/case"
	writeToFile2(InitialTestcase, caseURL)
	SourceURL := FuzzURL + "/source.go"
	writeToFile2(upload, SourceURL)
	fuzzing(c, FuzzURL)

}
func exeSysCommand(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	opBytes, _ := cmd.Output()

	return string(opBytes)
}
func writeToFile2(msg string, URL string) {
	if err := ioutil.WriteFile(URL, []byte(msg), 777); err != nil {
		//os.Exit(111)
		//fmt.Println(err.Error())
	}
}
func fuzzing(c *FuzzController, FuzzURL string) {
	c.Ctx.WriteString("fuzzing start...")
	genURL := FuzzURL + "/gen"
	_ = exeSysCommand("cd " + genURL + ";go run main.go -out ../corpus/")
	_ = exeSysCommand("cd " + FuzzURL + ";export GO111MODULE=auto;go-fuzz-build;nohup go-fuzz > output.log 2>&1 &")
	c.Ctx.WriteString("fuzzing finished!")
}
