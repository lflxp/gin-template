package router

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	_ "github.com/lflxp/gin-template/docs"
	log "github.com/sirupsen/logrus"
)

func Run(ip, port string) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	// 注册路由
	PreGinServe(router)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", ip, port),
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	log.Infof("Listening and serving HTTPS on %s:%s", ip, port)
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
