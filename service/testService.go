package service

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
)

/**
测试路由
*/
func Test() string {
	return "hello test"
}

/**
测试返回状态码 和 json
*/
func TestShowUrl() (int, string) {
	var user struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	user.Id = 1
	user.Name = "test"
	j, _ := ffjson.Marshal(user)
	fmt.Println(string(j))
	return 200, string(j)
}
