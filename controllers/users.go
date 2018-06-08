package controllers

import (
	"minego/models/mgo"
	"minego/models"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type UsersControllers struct {
	BaseController
}

func (this *UsersControllers)Get()  {
	var page, limit int
	page, _ = this.GetInt("_page", page)
	limit, _ =this.GetInt("_limit", limit)

	users, err := mgo.GetUsers(page, limit)
	if err != nil {
		this.ResJson(users)
	}

	count, _ := mgo.Count()
	this.Ctx.Output.Header("x-total-count", strconv.Itoa(count))
	this.ResJson(users)
}

func (this *UsersControllers)Post()  {
	user := models.User{ID: bson.NewObjectId()}
	err := this.ParseBody(&user)
	if err != nil {
		beego.Debug(err)
	}

	mgo.PostUser(user)
	this.Get()
}

func (this *UsersControllers)Delete()  {
	id := this.Ctx.Input.Param(":id")
	mgo.DeleteUser(id)

	this.Get()
}

func (this *UsersControllers)Patch()  {
	id := this.Ctx.Input.Param(":id")
	user := models.User{ID: bson.ObjectIdHex(id)}
	err := this.ParseBody(&user)
	if err != nil {
		beego.Debug(err)
	}
	mgo.PatchUser(user)

	this.Get()
}