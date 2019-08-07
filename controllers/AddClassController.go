package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"github.com/lvchay00/myblog2/models"
)

type AddClassController struct {
	beego.Controller
}

func (c *AddClassController) Index() {

	c.Data["PageTitle"] = "增加分类"
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.html"
	c.TplName = "AddClass.tpl"
}
func (c *AddClassController) Add() {

	c.Data["PageTitle"] = "Add分类"
	name := c.Input().Get("name")
	sections := models.Section{Name: name}
	models.InsertSection(&sections)
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.html"
	c.TplName = "AddClass.tpl"
}
func (c *AddClassController) Delete() {

	c.Data["PageTitle"] = "Delete分类"
	id, _ := strconv.Atoi(c.Input().Get("sid"))
	sections := models.Section{Id: id}
	models.DeleteSection(&sections)
	c.Data["Sections"] = models.FindAllSection()
	c.Layout = "layout/layout.html"
	c.TplName = "AddClass.tpl"
}
