package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["title"] = "user"
	c.TplName = "user.tpl"
}

func (c *UserController) UserAdd(){
	c.TplName = "user_add.tpl"
}

func (c *UserController) UserDo(){
	id,err := c.GetInt("id")
	if err != nil{
		c.Ctx.WriteString("ID必须为数字")
		return
	}
	username := c.GetString("username")
	password := c.GetString("password")

	c.Ctx.WriteString("表单数据" + strconv.Itoa(id) + username + password)

}

func (c *UserController) UserEdit(){
	c.TplName = "user_edit.tpl"
}
