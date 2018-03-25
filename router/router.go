package router

import (
	"fmt"
	"github.com/akkagao/martiniStudy/service"
	"github.com/go-martini/martini"
)

func Router(m *martini.ClassicMartini) {
	//测试根目录
	m.Get("/", func() string {
		return "hello martini"
	})
	m.Get("/test", service.Test)

	m.Get("/test/showurl", service.TestShowUrl)

	// 用户模块
	m.Group("/user", func(r martini.Router) {
		// 测试返回数据
		r.Group("/test", func(router martini.Router) {
			r.Get("/getUser/:name", service.GetUserTest)
			r.Post("/addUser/:name/:tel/:mail", service.AddUserTest)
		})

		// 从数据库读取数据
		r.Group("/db", func(router martini.Router) {
			r.Get("/getUser/:id", service.GetUser)
		})
	})

	m.NotFound(func() {
		fmt.Println("not fount")
	})
}
