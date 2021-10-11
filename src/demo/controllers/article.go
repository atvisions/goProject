package controllers

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Data["title"] = "新闻列表"
	c.TplName = "article.tpl"
}

func (c *ArticleController) ArticleAdd() {
	c.Data["title"] = "添加新闻"
	c.TplName = "article_add.tpl"
}

func (c *ArticleController) ArticleEdit() {
	id := c.GetString("id")

	c.Data["title"] = "编辑新闻" + id
	c.TplName = "article_edit.tpl"
}