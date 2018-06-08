package mgo

import (
	"minego/models"
)

func ListUsers(page int, limit int) (users []models.User, err error) {
	err = mHandler.C("users").Find(nil).Skip((page-1) * limit).Limit(limit).All(&users)
	return
}
