package api

import (
	"fmt"
	"github.com/kalifun/gin-template/global"
	"github.com/kalifun/gin-template/middleware/config"
	"github.com/kalifun/gin-template/middleware/logs"
	"github.com/kalifun/gin-template/router"
	"github.com/spf13/cobra"
	"net/http"
)

var Api = &cobra.Command{
	Use:     "server",
	Short:   "Start Server",
	Long:    "Start Api Server",
	Example: "gin-template server",
	Run: func(cmd *cobra.Command, args []string) {
		RunServer()
	},
}

func RunServer() {
	// 初始化中间件
	config.Init()
	logs.InitLog()
	router := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", global.ConfigSvr.System.Port),
		Handler:        router,
		//ReadTimeout:    conf.ReadTimeout,
		//WriteTimeout:   conf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
