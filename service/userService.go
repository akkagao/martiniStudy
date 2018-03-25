package service

import (
	"fmt"
	"github.com/akkagao/martiniStudy/dao"
	"github.com/akkagao/martiniStudy/util"
	"github.com/go-martini/martini"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/quexer/utee"
	"strconv"
)

func GetUserTest(p martini.Params) (int, string) {
	fmt.Println(p)
	return 200, p["name"]
}

func GetUser(p martini.Params) (int, string) {
	id, err := strconv.Atoi(p["id"])
	utee.Chk(err)
	var user struct {
		ID   int `json:"id"`
		Name string `json:"name"`
	}
	if err = dao.DB.Raw("select t.id,t.name from t_user t where t.id = ?", id).Scan(&user).Error; err != nil {
		return 500, "select user fail..."
	}
	result, err := util.ToJson(user)
	utee.Chk(err)
	return 200, result
}

func AddUserTest(p martini.Params) (int, string) {
	var user struct {
		Name string `json:"name"`
		Tel  string `json:"tel"`
		Mail string `json:"mail"`
	}
	user.Name = p["name"]
	user.Tel = p["tel"]
	user.Mail = p["mail"]

	jsonByte, err := ffjson.Marshal(user)
	if err != nil {
		return 500, "系统异常请稍后再试。。。"
	}
	return 200, string(jsonByte)
}
