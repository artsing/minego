package mgo

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"mine/mongo"
)

var (
	mHandler *mongo.Handler
)

func init() {
	url := beego.AppConfig.String("MongoURL")
	if url == ""{
		url = "mongodb://localhost:27017/mine?maxPoolSize=512"
	}
	handler, err := mongo.InitHandler(url)
	if err != nil {
		panic(err)
	}
	mHandler = handler
}

func HexObjectID(objID bson.ObjectId) string {
	return objID.Hex()
}
