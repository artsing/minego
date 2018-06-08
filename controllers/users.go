package controllers

import (
	"minego/models/mgo"
	"github.com/astaxie/beego"
)

type UsersControllers struct {
	BaseController
}

func (this *UsersControllers)Get()  {
	var page, limit int
	page, _ = this.GetInt("_page", page)
	limit, _ =this.GetInt("_limit", limit)
	users, err := mgo.ListUsers(page, limit)
	if err != nil {
		this.ResJson(users)
	}
	this.ResJson(users)
}

func (this *UsersControllers)Post()  {
	var page, limit int
	page, _ = this.GetInt("_page", page)
	limit, _ =this.GetInt("_limit", limit)
	users, err := mgo.ListUsers(page, limit)
	if err != nil {
		this.ResJson(users)
	}
	this.ResJson(users)
}

func (this *UsersControllers)Delete()  {
	id := this.Ctx.Input.Param(":id")
	beego.Debug(id)
	var page, limit int
	page, _ = this.GetInt("_page", page)
	limit, _ =this.GetInt("_limit", limit)
	users, err := mgo.ListUsers(page, limit)
	if err != nil {
		this.ResJson(users)
	}
	this.ResJson(users)
}

func (this *UsersControllers)Patch()  {
	var page, limit int
	page, _ = this.GetInt("_page", page)
	limit, _ =this.GetInt("_limit", limit)
	users, err := mgo.ListUsers(page, limit)
	if err != nil {
		this.ResJson(users)
	}
	this.ResJson(users)
}