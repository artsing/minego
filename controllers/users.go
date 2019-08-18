package controllers

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"minego/models"
	"minego/models/mgo"
	"strconv"
)

type UsersControllers struct {
	BaseController
}

func (c *UsersControllers)Get()  {
	var page, limit int
	page, _ = c.GetInt("_page", page)
	limit, _ =c.GetInt("_limit", limit)

	users, err := mgo.GetUsers(page, limit)
	if err != nil {
		c.ResJson(users)
	}

	count, _ := mgo.Count()
	c.Header("x-total-count", strconv.Itoa(count))
	c.ResJson(users)
}

func (c *UsersControllers)Post()  {
	user := models.User{ID: bson.NewObjectId()}
	err := c.ParseBody(&user)
	if err != nil {
		beego.Debug(err)
	}

	err = mgo.PostUser(user)
	c.Get()
}

func (c *UsersControllers)Delete()  {
	id := c.Ctx.Input.Param(":id")
	_ = mgo.DeleteUser(id)

	c.Get()
}

func (c *UsersControllers)Patch()  {
	id := c.Ctx.Input.Param(":id")
	user := models.User{ID: bson.ObjectIdHex(id)}
	err := c.ParseBody(&user)
	if err != nil {
		beego.Debug(err)
	}
	_ = mgo.PatchUser(user)

	c.Get()
}