package service

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/pquerna/ffjson/ffjson"
)

func GetUser(p martini.Params) (int, string) {
	fmt.Println(p)
	return 200, p["name"]
}

func AddUser(p martini.Params) (int, string) {
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
