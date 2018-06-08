package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

const (
	Error       = 0
	SUCCESS     = 1
)

// ResResult 响应数据
type ResResult struct {
	Status  int         `json:"status"` // 响应状态码
	Data    interface{} `json:"data"`   // 响应数据
	Message string      `json:"msg"`    // 响应消息
}

// BaseController 提供基础操作
type BaseController struct {
	beego.Controller
}

// ParseBody 解析请求的body数据
func (this *BaseController) ParseBody(v interface{}) error {
	return json.Unmarshal(this.Ctx.Input.RequestBody, v)
}

// Header 设置key, value
func (this *BaseController)Header(key, value string)  {
	this.Ctx.Output.Header(key, value)
}


// ResJson 响应json数据
func (this *BaseController) ResJson(v interface{}) {
	this.Data["json"] = v
	this.ServeJSON()
}

// Success 响应成功数据
func (this *BaseController) Success(data interface{}, msg string) {
	result := ResResult{
		Status:  SUCCESS,
		Data:    data,
		Message: msg,
	}
	this.ResJson(result)
}

// Error 响应异常数据
func (this *BaseController) Error(status int, msg string, err error) {
	result := ResResult{
		Status:  status,
		Message: msg,
	}
	beego.Error(fmt.Sprintf("%s:%v", msg, err))
	this.ResJson(result)
}
