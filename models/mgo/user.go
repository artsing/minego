package mgo

import (
	"minego/models"
	"gopkg.in/mgo.v2/bson"
)

const USERS = "users"

func Count() (int, error) {
	return mHandler.C(USERS).Count()
}

func GetUsers(page int, limit int) (users []models.User, err error) {
	err = mHandler.C(USERS).Find(nil).Skip((page-1) * limit).Limit(limit).All(&users)
	return
}

func PostUser(user models.User) error {
	return mHandler.C(USERS).Insert(user)
}

func PatchUser(user models.User) error {
	return mHandler.C(USERS).UpdateId(user.ID, user)
}

func DeleteUser(id string) error {
	return mHandler.C(USERS).RemoveId(bson.ObjectIdHex(id))
}
