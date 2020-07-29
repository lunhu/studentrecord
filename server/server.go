package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"studentrecord/route"
	"time"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func NewWebServer(listenAddr string) *http.Server {

	// 创建mux路由器
	mx := mux.NewRouter()

	// 初始化路由
	route.InitRoutes(mx)

	// 使用negroni中间件
	n := negroni.New()
	r := negroni.NewRecovery()
	r.PrintStack = false
	n.Use(r)

	// 启用logger
	logger := negroni.NewLogger()
	n.Use(logger)

	// 启用JWT验证中间件
	//n.Use(util.NewAuth())

	// 注册路由器
	n.UseHandler(mx)

	// http 服务配置:
	serv := &http.Server{
		Addr:    listenAddr,
		Handler: n,
		//ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return serv
}

// 关闭server
// quit: 接收关闭信号
// done: 发出已关闭信号
func Shutdown(serv *http.Server, logger *log.Logger, quit <-chan os.Signal, done chan<- struct{}) {
	// 等待接收到退出信号
	<-quit
	logger.Println("Sever is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	serv.SetKeepAlivesEnabled(false)
	err := serv.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("could not gracefully shutdown the server: %v \n", err)
	}

	// do somthing :

	fmt.Println("do something start ...", time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("do something end ...", time.Now())

	close(done)
}
