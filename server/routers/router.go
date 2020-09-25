package routers

import (
	"example/server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/fuzz", &controllers.FuzzController{}, "*:GetData")
	beego.Router("/fuzz/fuzzstop", &controllers.FuzzStopController{}, "*:Get")
	beego.Router("/fuzz/fuzzdata", &controllers.FuzzDataController{}, "*:Get")
}
