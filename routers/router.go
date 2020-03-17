package routers

import (
	"testing-golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/template", &controllers.TemplateController{})
}
