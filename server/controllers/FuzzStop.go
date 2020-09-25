package controllers

import (
	"os"
	"time"

	"github.com/astaxie/beego"
)

type FuzzStopController struct {
	beego.Controller
}

func (c *FuzzStopController) Get() {
	Version := c.GetString("Version")
	gopath := os.Getenv("GOPATH")
	var FuzzURL string
	if Version == "1.1" {
		FuzzURL = gopath + "/src/example/server/models/fabric1.1"
	} else {
		FuzzURL = gopath + "/src/github.com/hyperledger/fabric/integration/chaincode/aab"
	}
	output := "mv " + FuzzURL + "/output.log result/"
	corpus := "mv " + FuzzURL + "/corpus result/"
	crashers := "mv " + FuzzURL + "/crashers result/"
	suppressions := "mv " + FuzzURL + "/suppressions result/"
	zip := "zip -q -r result.zip result"
	res := "mkdir result" + ";" + output + ";" + corpus + ";" + crashers + ";" + suppressions + ";" + zip
	_ = exeSysCommand(res)
	c.Ctx.Output.Download("result.zip")
	time.Sleep(30 * time.Second)
	Clear(Version)
}

func DownloadFile(c *FuzzStopController, name string) {
	id := c.Ctx.Input.Param(":" + name)
	//fmt.Println("*****download file: ", id)
	c.Ctx.Output.Download("./models/fabric1.1" + id)
}
func Clear(Version string) {
	var FuzzURL string
	gopath := os.Getenv("GOPATH")
	if Version == "1.1" {
		FuzzURL = gopath + "/src/example/server/models/fabric1.1"
	} else {
		FuzzURL = gopath + "/src/github.com/hyperledger/fabric2.0/integration/chaincode/aab"
		_ = exeSysCommand("cd " + gopath + "/src/github.com/hyperledger;sudo mv fabric fabric2.0;sudo mv fabric1.1 fabric")
	}
	_ = exeSysCommand("rm -rf result;rm -rf result.zip;" + "rm -rf " + FuzzURL + "/test-fuzz.zip")
}
