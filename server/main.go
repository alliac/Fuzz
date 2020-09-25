package main

import (
	_ "example/server/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
