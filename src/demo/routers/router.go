package routers

import (
	"demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/add", &controllers.ArticleController{},"get:ArticleAdd")
	beego.Router("/article/edit", &controllers.ArticleController{},"get:ArticleEdit")
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{},"get:UserAdd")
	beego.Router("/user/addDo", &controllers.UserController{},"post:UserDo")
}
