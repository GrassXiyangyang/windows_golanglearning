package routers

import (
	"crawl_movie/crawl_movie/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
