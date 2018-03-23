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

	m.Group("/user", func(r martini.Router) {
		r.Get("/getUser", service.GetUser)
		r.Post("/addUser/:name/:tel/:mail", service.AddUser)
	})

	m.NotFound(func() {
		fmt.Println("not fount")
	})
}
