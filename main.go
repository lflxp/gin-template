package main

import (
	"flag"
	"os"

	"github.com/lflxp/gin-template/router"
	log "github.com/sirupsen/logrus"
)

var (
	host string
	port string
)

func init() {
	flag.StringVar(&host, "host", "0.0.0.0", "start server host")
	flag.StringVar(&port, "port", "8888", "port number")

	// 设置日志格式为json格式
	log.SetFormatter(&log.TextFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(log.DebugLevel)
}

// @title Gin Template
// @version 1.1
// @description Gin API 接口模板服务
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	flag.Parse()
	router.Run(host, port)
}
