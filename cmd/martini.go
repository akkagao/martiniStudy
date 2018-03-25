// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/go-martini/martini"
	"github.com/spf13/cobra"

	"github.com/akkagao/martiniStudy/router"
	"github.com/akkagao/martiniStudy/dao"
)

// martiniCmd represents the martini command
var martiniCmd = &cobra.Command{
	Use:   "martini",
	Short: "martini 框架简单使用",
	Long: `martini 框架的简单使用
	路由方式、请求参数解析等`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("martini start ...")
		initMartini()
	},
}

/**
用Martini 初始化启动web服务
*/
func initMartini() {
	m := martini.Classic()
	//http.Handle("/", m)
	m.Handlers(martini.Recovery())
	router.Router(m)
	println("test")
	dao.InitGorm()
	m.Run()
	defer dao.DB.Close()
}

func init() {
	rootCmd.AddCommand(martiniCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// martiniCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// martiniCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
